package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Xuanwo/templateutils"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

const (
	pairPath     = "pairs.hcl"
	metadataPath = "metadata.hcl"
)

func parse() (data *Data) {
	data = &Data{}

	// Parse pairs
	content, err := ioutil.ReadFile(pairPath)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}
	err = parseHCL(content, pairPath, data)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}
	for _, v := range data.Pairs {
		v.GeneratedDescription = strings.ReplaceAll(v.Description, "\n", "\n//")
	}

	data.TypeMap = make(map[string]string)
	for _, v := range data.Pairs {
		data.TypeMap[v.Name] = v.Type
	}

	// Parse metadata
	content, err = ioutil.ReadFile(metadataPath)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}
	err = parseHCL(content, metadataPath, data)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	// Parse service
	files, err := ioutil.ReadDir("services")
	if err != nil {
		log.Fatalf("parse: %v", err)
	}
	for _, v := range files {
		if !strings.HasSuffix(v.Name(), ".hcl") {
			continue
		}

		srv := parseService("services/" + v.Name())

		// Create type map
		srv.TypeMap = data.TypeMap

		data.Service = append(data.Service, srv)
	}

	data.Sort()
	return
}

func injectReadCallbackFunc(ops *Ops) {
	if ops == nil {
		return
	}

	for _, op := range ops.Op {
		fn := op.Func
		if fn == nil {
			continue
		}
		if strings.Contains(fn.Params, "io.Reader") ||
			strings.Contains(fn.Returns, "io.ReadCloser") {
			op.Generated = append(op.Generated, "read_callback_func")
		}
	}
}

func injectContext(ops *Ops) {
	if ops == nil {
		return
	}

	for _, op := range ops.Op {
		op.Generated = append(op.Generated, "context")
	}
}

func parseHCL(src []byte, filename string, in interface{}) (err error) {
	var diags hcl.Diagnostics
	defer func() {
		if err != nil {
			err = fmt.Errorf("parse hcl: %w", err)
		}
	}()

	file, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return diags
	}

	diags = gohcl.DecodeBody(file.Body, nil, in)
	if diags.HasErrors() {
		return diags
	}

	return nil
}

func parseService(filePath string) *Service {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	srv := &Service{}
	err = parseHCL(content, filePath, srv)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	servicerFuncs, storagerFuncs := parseFunc(srv.Name, "servicer"), parseFunc(srv.Name, "storager")

	// Register funcs into service
	if srv.Service != nil {
		for _, v := range srv.Service.Op {
			if fn, ok := servicerFuncs[v.Op]; ok {
				fn.hasPair = true
				v.Func = fn
			}
		}

		// Add missing pairs into service
		for k, v := range servicerFuncs {
			if v.hasPair {
				continue
			}
			srv.Service.Op = append(srv.Service.Op, &Op{
				Op:   k,
				Func: v,
			})
		}
	}

	for _, v := range srv.Storage.Op {
		if fn, ok := storagerFuncs[v.Op]; ok {
			fn.hasPair = true
			v.Func = fn
		}
	}
	for k, v := range storagerFuncs {
		if v.hasPair {
			continue
		}
		srv.Storage.Op = append(srv.Storage.Op, &Op{
			Op:   k,
			Func: v,
		})
	}

	// Inject paris
	injectReadCallbackFunc(srv.Storage)
	injectContext(srv.Service)
	injectContext(srv.Storage)

	return srv
}

func parseFunc(service, name string) map[string]*Func {
	data := make(map[string]*Func)
	filename := fmt.Sprintf("../services/%s/%s.go", service, name)

	content, err := ioutil.ReadFile(filename)
	if os.IsNotExist(err) {
		return data
	}
	if err != nil {
		log.Fatalf("read file failed: %v", err)
	}

	f, err := parser.ParseFile(token.NewFileSet(), filename, string(content), 0)
	if err != nil {
		log.Fatalf("decorator parse failed: %v", err)
	}

	for _, fn := range f.Decls {
		fndecl, ok := fn.(*ast.FuncDecl)
		// Ignore Non-FuncDecl node.
		if !ok {
			continue
		}
		// Ignore non-exported funcs.
		if !fndecl.Name.IsExported() {
			continue
		}
		// Ignore some not needed functions.
		if fndecl.Name.Name == "String" {
			continue
		}

		opName := templateutils.ToSnack(fndecl.Name.Name)

		data[opName] = &Func{
			Parent:   fndecl.Name.Name,
			Receiver: getReceiver(fndecl),
			Returns:  getReturns(fndecl),
			Caller:   getCaller(fndecl),
			Params:   getParams(fndecl),
		}

		if fndecl.Recv != nil {
			data[opName].HasContext = true
		}
	}
	return data
}

func getReceiver(fn *ast.FuncDecl) string {
	if fn.Recv == nil {
		return ""
	}
	return fmt.Sprintf("s %s", formatExpr(fn.Recv.List[0].Type))
}

func getParams(fn *ast.FuncDecl) string {
	parms := []string{}
	for _, v := range fn.Type.Params.List {
		parms = append(parms, formatField(v))
	}
	ans := fmt.Sprintf("%s", strings.Join(parms, ","))
	return ans
}

func getReturns(fn *ast.FuncDecl) string {
	results := []string{}
	for _, v := range fn.Type.Results.List {
		results = append(results, formatField(v))
	}
	ans := fmt.Sprintf("%s", strings.Join(results, ","))
	return ans
}

func getCaller(fn *ast.FuncDecl) string {
	parms := []string{}
	for _, v := range fn.Type.Params.List {
		if _, ok := v.Type.(*ast.Ellipsis); ok {
			parms = append(parms, v.Names[0].Name+"...")
			continue
		}
		for _, name := range v.Names {
			parms = append(parms, name.Name)
		}
	}
	ans := fmt.Sprintf("%s", strings.Join(parms, ","))
	return ans
}

func formatField(f *ast.Field) string {
	s := []string{}
	for _, name := range f.Names {
		s = append(s, name.Name)
	}
	return strings.Join(s, ",") + " " + formatExpr(f.Type)
}

func formatExpr(t ast.Expr) string {
	switch v := t.(type) {
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", formatExpr(v.X), v.Sel.Name)
	case *ast.Ident:
		return v.Name
	case *ast.StarExpr:
		return "*" + formatExpr(v.X)
	case *ast.Ellipsis:
		return "..." + formatExpr(v.Elt)
	default:
		println(fmt.Sprintf("not handled type %+#v", v))
		return ""
	}
}