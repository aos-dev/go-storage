// Code generated by go-bindata. DO NOT EDIT.
// sources:
// cmd/definitions/tmpl/function.tmpl (567B)
// cmd/definitions/tmpl/info.tmpl (1.702kB)
// cmd/definitions/tmpl/object.tmpl (2.192kB)
// cmd/definitions/tmpl/operation.tmpl (1.698kB)
// cmd/definitions/tmpl/pair.tmpl (502B)
// cmd/definitions/tmpl/service.tmpl (13.988kB)

// +build tools

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _cmdDefinitionsTmplFunctionTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x91\x41\x4b\xc4\x30\x10\x85\xef\xfd\x15\x8f\x25\x87\xae\xec\xe6\x07\x08\x9e\x8a\x82\xb0\xc8\xa2\xde\x25\x64\xa7\x6b\xb0\x99\x94\x66\x5a\x17\x62\xfe\xbb\xa4\x55\x17\x11\x0f\xde\x3c\x25\xcc\xcc\x7b\xf3\x3e\x26\xa5\x2d\x06\xc3\x47\x82\x7a\xda\x40\x4d\xb8\xbc\x82\xd2\x37\x23\xdb\x88\x9c\xab\xd2\x76\x2d\x38\x08\xd4\xa4\x6f\x7d\xdf\x91\x27\x16\x3a\x7c\x36\x55\xcb\x2f\xb3\x66\xd2\x77\xc6\x13\xde\x20\xa1\x31\x9e\xba\x65\xa0\x88\xd5\xa4\x77\xc1\x9a\xb9\xd2\x8e\x6c\x51\x47\x5c\xa4\x04\x75\x56\xec\x4d\x5c\x06\xd6\x48\xa9\x58\xe6\x5c\xa7\xa4\x26\xbd\x37\x83\xf1\x51\x3f\x0e\xce\xef\x4c\x14\xfd\x20\x83\xe3\xe3\x35\x1f\xe2\xab\x93\xe7\x26\x78\x6f\x72\x46\xe8\x05\xbd\x71\xc3\x2f\xa6\xa5\x5c\x62\x7e\xdf\xb4\x2c\xb8\xa7\x38\x76\x12\x3f\x8c\xe7\x00\x15\x00\xf4\x86\x9d\xad\x57\x05\xdc\x9d\xa9\x57\xeb\x6a\xa6\xa2\x2e\xd2\x1f\x71\xac\x9c\x60\x03\x0b\x9d\x44\x37\xcb\xbb\xc1\x3f\x66\xdc\x82\xf8\xeb\xca\x3f\xbe\xef\x01\x00\x00\xff\xff\x41\xc8\xae\xbf\x37\x02\x00\x00")

func cmdDefinitionsTmplFunctionTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplFunctionTmpl,
		"cmd/definitions/tmpl/function.tmpl",
	)
}

func cmdDefinitionsTmplFunctionTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplFunctionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/function.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc1, 0xfc, 0xb, 0xf4, 0xfb, 0xf, 0x94, 0xca, 0x64, 0x7a, 0xb8, 0x48, 0xf9, 0x69, 0xf1, 0x88, 0x46, 0x59, 0xd3, 0xbc, 0x0, 0x4e, 0xe7, 0x8, 0x1e, 0xdf, 0x27, 0x59, 0x81, 0x75, 0x52, 0xa2}}
	return a, nil
}

var _cmdDefinitionsTmplInfoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x94\x5f\x6f\xd3\x30\x14\xc5\xdf\xfd\x29\x0e\x55\x85\x1a\xd4\x35\x9b\x84\x78\x80\xe5\x69\x1b\x68\x42\xdb\x90\x36\xf1\x82\x10\x72\x92\x9b\xca\x34\xb6\x23\xdb\x89\x56\x32\x7f\x77\xe4\x24\x74\x4d\xf7\x07\x8a\x78\xe1\xcd\x37\xf6\xbd\xe7\x77\x8f\xaf\x13\xc7\x38\xd1\x39\x61\x49\x8a\x0c\x77\x94\x23\x5d\x63\xa9\x37\x31\x32\x99\xc7\x39\x15\x42\x09\x27\xb4\xb2\xef\x70\x7a\x85\xcb\xab\x1b\x9c\x9d\x9e\xdf\x2c\x58\xc5\xb3\x15\x5f\x12\xdc\xba\x22\xcb\x98\x90\x95\x36\x0e\x33\x06\x00\x93\x42\xba\x09\x8b\x18\x6b\xdb\x03\x18\xae\x96\x84\xe9\x6a\x8e\xa9\x50\x85\xb6\x78\x9b\x60\x71\x1e\x56\x17\xbc\x82\xf7\xac\x6d\x31\xb5\x64\x1a\x91\xd1\x25\x97\x14\xf6\xa7\x2b\xdc\xc1\xe9\x13\x2e\xa9\x0c\x47\x58\x1c\xe3\xbd\xa0\x32\x87\x50\x39\xdd\x42\x28\xb4\xed\x76\x92\xf7\x48\x85\x63\x99\x56\x36\x40\xec\xe8\x36\x5d\xcd\x5e\xdd\xfb\x0e\x71\x37\xfd\x3c\xd4\x0d\x24\xcd\xa2\x83\x08\xf2\x9f\xb8\xcd\x78\xd0\x47\x82\xa3\xe3\xe3\xb0\xbb\xea\x81\x0f\x40\x2a\x0f\xcb\x88\xb1\x60\x00\x76\x7b\x18\xa7\x5b\x67\xea\xcc\xa1\x1d\x94\x37\x6c\xdf\x9e\x62\x0b\x18\x37\xeb\xaa\xaf\xe5\xfd\xd6\x97\xfb\x33\x1b\x86\x2e\x8e\xe3\x60\x00\x6a\x4b\x39\xb8\x05\x0f\x91\xe4\x15\x0a\x6d\xa0\xd3\xef\x94\x39\x34\xbc\xac\x69\x8e\x43\x48\xe2\xca\x42\x69\x07\x4b\x6e\x8e\xa3\xe1\x83\x25\xd7\x95\xea\xea\x08\xe5\xde\xbc\xee\x42\x09\xc9\xab\x2f\xd6\x19\xa1\x96\x5f\x85\x72\x64\x0a\x9e\x51\xeb\xd9\xa0\xfc\xbc\xd7\x61\x57\x14\x81\xfe\xec\xb6\x9b\x10\xef\x59\x51\xab\x0c\x33\x89\x57\xcf\xba\x16\xe1\x03\xb9\xbe\xf1\x53\x61\xab\x92\xaf\x07\x37\x66\xd1\xd8\x8f\xc1\x57\x43\xae\x36\x0a\x72\xf1\xc0\xbe\x40\xfa\xa7\x9a\xd7\x4f\x68\x36\x63\xcd\xe8\x37\x85\x06\xa6\x47\x60\x90\xa0\x19\xf1\xb2\x61\xa0\x4a\xdb\xb1\xfe\x03\x77\x66\x23\xd4\x39\x52\xad\xcb\x68\x20\x12\x05\xe4\x22\xdc\xf0\xcb\x3d\x9f\xc0\x8b\x04\x87\x43\x8d\xe7\xdd\x9e\xc3\x99\x9a\xba\x83\x7e\xbb\xd1\x2d\xa8\x3b\xfc\x20\xa3\x3f\x87\x79\xec\x12\x0a\x5e\x5a\xda\xe7\x96\x2e\x6a\xeb\xf6\x9b\x8e\xbf\xee\x3b\x19\xf7\x5d\x71\x25\xb2\x59\x21\xdd\xe2\xba\x32\x42\xb9\x62\x36\x79\x8c\xf5\x23\xa5\x3c\xbd\x7f\xb9\xbf\xee\x5e\x6c\x9e\xdd\x24\x8a\x1e\x5a\xf4\xdf\xcc\x6e\x6f\xe5\x5d\xb2\x9f\x97\x8f\x8e\x7d\xff\x0f\xdb\xf9\xa5\xdd\x2f\x7f\x06\x00\x00\xff\xff\x26\xf5\xd0\xb3\xa6\x06\x00\x00")

func cmdDefinitionsTmplInfoTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplInfoTmpl,
		"cmd/definitions/tmpl/info.tmpl",
	)
}

func cmdDefinitionsTmplInfoTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplInfoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/info.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa6, 0x9d, 0xdc, 0x29, 0x17, 0xc8, 0x8a, 0x14, 0x25, 0x13, 0xa5, 0x85, 0x13, 0x6, 0xf, 0xb3, 0x4a, 0x21, 0xd1, 0x20, 0x4e, 0x6f, 0x92, 0x60, 0xe2, 0x6e, 0x29, 0x96, 0x4a, 0xee, 0x99, 0x66}}
	return a, nil
}

var _cmdDefinitionsTmplObjectTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x95\x51\x73\xe3\x34\x10\xc7\x9f\xab\x4f\xf1\x27\xd3\x81\x98\x49\xed\x3b\x60\x78\x38\x2e\x0f\xcc\xb5\xc0\xcd\xd0\x86\x99\x14\xde\x15\x79\x9d\x8a\xda\x92\x47\x5a\xe7\x12\x5c\x7f\x77\x46\xb2\xdb\x73\xda\xe4\xc8\x0d\x03\x6f\x92\xb5\xfb\xd7\xee\x6f\x77\xe5\x2c\xc3\x3b\x9b\x13\xd6\x64\xc8\x49\xa6\x1c\xab\x1d\xd6\xf6\x69\x0f\x55\xe5\x59\x4e\x85\x36\x9a\xb5\x35\xfe\x07\x5c\x2e\x70\xb3\xb8\xc5\xd5\xe5\xfb\xdb\x54\xd4\x52\xdd\xcb\x35\x81\x77\x35\x79\x21\x74\x55\x5b\xc7\x98\x0a\x00\x98\x14\x15\x4f\xfa\x15\xeb\x8a\x86\xa5\xdf\x19\x35\x11\x89\x10\x59\x86\x9f\x34\x95\x39\xb4\xc9\x69\x0b\x6d\x60\x57\x7f\x92\x62\xac\x34\x0b\x65\x8d\x0f\x3a\x6d\x7b\x01\x27\xcd\x9a\x70\x7e\x3f\xc3\xf9\x06\x6f\xe6\x48\x17\xd1\xee\x9a\x58\xa2\xeb\xa2\x6a\xef\xf9\x3e\x08\xb5\x2d\xce\x37\xe9\x8d\xac\x08\x0f\x60\xfb\x9b\xf4\x4a\x96\xe8\x3a\x34\xda\xf0\xf7\xdf\x61\x8e\xd7\x6f\xdf\x06\xa3\xfb\xe0\x1c\xf4\xc9\xe4\x61\xd9\x87\xd4\x6b\x43\x7b\xf0\x1d\xc1\x57\xb2\x2c\xc9\x33\x1a\xa3\x39\x84\xb8\xb6\x17\x9e\xad\x93\x6b\x4a\x45\x96\x05\x87\x9b\xc5\xed\xd5\xf2\x4d\x58\x01\x17\x83\xfb\x57\x1e\x45\x48\xcd\x63\xf9\xcb\xe2\xf7\x5f\x2f\x61\x2c\x63\x45\x50\x77\x21\x95\x1c\xb6\x61\xaf\x73\x82\x27\xb7\xd1\x8a\x7c\xba\xef\x8e\x77\x3f\xde\x04\xc4\xc1\xc3\xd6\x9a\xf2\x67\xc7\xda\x43\x59\xa3\x1a\xe7\xc8\x30\xbc\x2c\x28\x15\xa1\x00\x8f\xe7\x9e\x5d\xa3\x18\xed\xa9\xf4\x82\x99\x2e\x02\xb6\x4b\xf2\xca\xe9\x3a\x54\xfa\xe3\xe1\xd1\x83\x27\x74\x1f\xed\x6e\x77\x35\x45\xf6\x5d\x37\xfa\xf2\x8c\xb4\x38\xcb\x32\xa8\x52\x87\xe8\x07\xd0\x8f\x3b\x83\x0f\x77\x5a\xdd\x8d\x32\x95\xa5\xde\x50\x2a\xce\x06\x8b\x65\x4f\xdf\x89\x78\x69\x96\x85\x6e\x41\xe3\x29\x87\xf4\x90\x61\x57\xc9\x1a\x85\x75\x8f\xdd\xb4\x91\x65\x43\x33\xbc\x42\x45\xd2\xf8\x58\x09\x4f\x3c\xc3\xeb\xe1\x83\x27\x8e\x52\x51\x27\xb6\x88\x38\xcb\xad\xa1\xb8\xf9\xf6\x1b\x71\x56\x85\xd3\xd0\xb5\xe9\x75\xc3\xb4\x15\x9d\x10\xa7\x80\x6d\xdb\x81\xe9\xd5\x36\xce\x44\xd7\x89\xa2\x31\x0a\x53\x8b\xaf\x7b\xcb\x04\x3f\x13\x0f\x78\xb5\xaf\x4b\xb9\x1b\xc8\x4d\x93\x7d\x76\x68\x63\x84\x8e\xb8\x71\x06\x36\x7d\x81\x3a\xc4\xf4\x42\x7c\x79\x44\x7c\xb3\x2f\x9e\x3c\x7a\x0c\xb7\x1c\x90\xc7\x1c\x9b\xbd\x08\xc4\x50\xce\xd2\xc7\xdb\x43\x7b\x1e\x49\x05\x1f\x74\x59\x62\x4d\x8c\x83\xa7\x85\xb3\xd5\x50\xeb\x38\x4f\x07\x9b\xed\x33\xb0\x4d\xf7\x52\x9b\x61\x65\x6d\x99\x3c\x25\xe6\x59\xf2\x34\xe9\x3b\x47\x17\xb0\x69\xa8\xf9\x97\x27\x3d\x1e\x5f\xcc\xf1\x6a\xd0\xf9\x74\x29\x66\x60\xd7\x50\x34\xec\xc4\x18\xda\x28\xb2\x07\xfc\x45\xce\xfe\x11\x1a\x33\x7a\x14\xb2\xf4\x24\x7a\x8e\xd7\x8d\xe7\xff\x93\xe5\x27\xee\x3b\xd2\x86\xff\x8a\xe3\x7c\x9f\x63\x2d\x8d\x56\xd3\xa2\xe2\x74\x59\x3b\x6d\xb8\x98\x4e\x86\xb1\x1d\x29\x74\x5d\x78\x07\x86\xc1\x9d\x24\xc9\x00\xf7\x9f\x47\x22\xcb\x8e\x0d\x41\x8f\xd2\x1f\x43\xa9\x0d\xdb\xcf\x47\xf9\x1f\x0c\x5c\x8f\xf6\x61\x7e\x0a\xdb\x83\x13\xda\x3f\xb8\xe3\xb7\xf7\x45\xd8\xaa\xb4\x86\xa6\xdb\xd1\x97\xf6\xe9\x7d\x3f\xe5\xd7\x7b\x30\xfc\xed\xa1\x92\x3c\xfb\x6b\x8c\x32\x8c\x0e\xe1\xbf\xdf\xfd\x1d\x00\x00\xff\xff\x97\x36\xb5\x8b\x90\x08\x00\x00")

func cmdDefinitionsTmplObjectTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplObjectTmpl,
		"cmd/definitions/tmpl/object.tmpl",
	)
}

func cmdDefinitionsTmplObjectTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplObjectTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/object.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb5, 0x55, 0xb8, 0xc, 0xe1, 0xc5, 0x5, 0xdc, 0x2, 0xe1, 0xf5, 0x81, 0x85, 0x80, 0x6d, 0x2c, 0x6, 0x2d, 0x35, 0xa2, 0x78, 0xa1, 0xc3, 0x18, 0x1f, 0x9e, 0x87, 0x5a, 0xbc, 0x61, 0x54, 0xd5}}
	return a, nil
}

var _cmdDefinitionsTmplOperationTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x54\x4d\x8f\xd3\x30\x10\xbd\xfb\x57\x8c\xac\x1e\x12\x69\x49\xef\x48\x7b\x82\x45\x5a\x09\x75\x2b\x10\xe2\x88\x5c\x67\xda\xb5\x48\x6c\x33\x9e\xec\x87\x82\xff\x3b\xb2\x93\x76\xd3\xaa\x10\x15\x10\xe2\x16\x7b\x66\xde\x9b\xf7\xc6\x13\xaf\xf4\x57\xb5\x43\xe0\x67\x8f\x41\x08\xd3\x7a\x47\x0c\x85\x00\x00\x90\xda\x59\xc6\x27\x96\xc3\xc9\x38\x29\x4a\x21\xfa\xfe\x15\x90\xb2\x3b\x84\xc5\x97\x2b\x58\x18\x78\x7d\x0d\xd5\xad\x65\xa4\xad\xd2\x18\x20\x46\xd1\xf7\xb0\x30\xd5\x5b\x0c\x9a\x8c\x67\xe3\x6c\xba\x4c\x0c\x30\x46\x4c\xf0\x8d\x7a\x5e\xa9\x16\x21\x46\x30\xfb\x62\xe8\x33\x53\x62\x30\x5b\x70\x04\x05\x7e\x4b\xf9\x39\x51\x06\xa4\x07\xa3\x91\x64\x79\x72\xcf\x8e\xd4\x2e\xdd\xc7\x98\xeb\x3f\x32\x19\xbb\x2b\x4a\x08\xf9\xe3\x80\x89\xb6\x4e\x8d\x8c\xe7\xa9\x08\xe7\x93\x8a\x85\xa9\xee\x7c\x16\x90\x32\x96\xcb\xdc\xad\xf3\x03\xcd\x77\x60\xb7\x56\x41\xab\x26\xb5\x3c\x46\x4e\x24\x8e\xc0\xe7\x6b\x8a\x31\xf2\xce\x51\xab\x78\xad\x48\xb5\x89\xab\x84\xe3\xc0\x07\x0c\x5d\xc3\xe1\xb3\xe1\xfb\xf5\x30\x9b\x23\x95\x32\x95\x4c\x5d\xb2\x8e\x73\xf5\x7b\x37\xd0\xcc\xf6\x9e\x90\xdf\x0c\x83\xfd\x1d\x19\x93\xf2\x42\xf3\x13\x8c\x6f\xa4\x1a\xef\xae\xfe\xba\xca\xfd\xd8\x8e\x4f\xf9\xd8\x76\x81\x6f\xda\x0d\xd6\x9f\xac\x69\x7d\x83\x2d\x5a\xc6\xfa\xdc\x1b\x2b\x4a\x11\x85\x58\x2e\x61\x36\x33\x83\xc2\x06\x01\x13\x70\x8d\x35\xb0\x83\x7b\xf5\x80\xb0\x75\xf4\xa8\xa8\x06\xed\x5a\xaf\xd8\x6c\x1a\x84\x03\x96\x4a\xde\x85\x6a\x78\xe4\xf3\x1c\x81\xa9\xd3\x0c\x7d\x14\x62\xdb\x59\x0d\x45\x98\x2f\x2a\x2f\x92\x7b\x29\xf6\xc9\xce\x8c\x9b\x48\xc8\x1d\x59\x90\xb3\x00\x32\xd9\xfb\xeb\x9d\xba\xa0\x99\x7f\xb2\x44\xc7\x3f\x9b\xb3\x6b\x84\x44\x70\x0d\x2b\x7c\xbc\xf3\x48\x79\xc4\x2b\xc7\xb7\x2f\xdd\xdf\x10\x39\x2a\xe4\xb4\xdb\x18\x65\x79\xfa\xc3\x79\x71\x52\x44\xf1\x33\xbe\x0b\xec\xf9\x0f\x76\x73\x34\xef\x4f\x0c\x9a\x38\xb2\x37\x6a\xfa\x75\x70\xef\x47\x00\x00\x00\xff\xff\xbd\xe8\x37\x32\xa2\x06\x00\x00")

func cmdDefinitionsTmplOperationTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplOperationTmpl,
		"cmd/definitions/tmpl/operation.tmpl",
	)
}

func cmdDefinitionsTmplOperationTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplOperationTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/operation.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc2, 0xaf, 0x9a, 0xb8, 0xf7, 0x9f, 0x1b, 0x53, 0xaa, 0x36, 0x95, 0x6d, 0xe6, 0x56, 0x7d, 0xb9, 0x8e, 0xd7, 0x4, 0x74, 0xce, 0x5b, 0x9, 0x4f, 0xb6, 0x92, 0xc, 0x1e, 0x49, 0x5, 0x4c, 0x4c}}
	return a, nil
}

var _cmdDefinitionsTmplPairTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x41\x6b\xe3\x30\x10\x85\xef\xfa\x15\x0f\xe3\x43\x02\x89\x75\xd9\x53\x96\x3d\x6d\xf6\xb0\x14\x92\x1c\x42\x7b\x2c\x8a\x3c\x51\x44\x6c\x49\xc8\x63\xb7\xc6\xf5\x7f\x2f\xb2\x93\x40\xe9\xa5\x3a\xcd\x7c\x33\xef\xf1\x46\x52\xe2\xaf\x2f\x09\x86\x1c\x45\xc5\x54\xe2\xd4\xc3\xf8\x47\x0f\x5d\x97\xb2\xa4\xb3\x75\x96\xad\x77\xcd\x6f\x6c\xf7\xd8\xed\x8f\xf8\xb7\xfd\x7f\x2c\x44\x50\xfa\xaa\x0c\x21\x28\x1b\x1b\x21\x6c\x1d\x7c\x64\x2c\x04\x00\x64\xda\x3b\xa6\x77\xce\xc4\xdc\x1a\xcb\x97\xf6\x54\x68\x5f\xcb\x13\xf5\xde\x95\x0d\xfb\xa8\x0c\x49\xe3\xd7\xf7\xb2\xfb\x25\xc3\xd5\xc8\x0b\x73\xd0\x95\x25\xc7\xd9\xa4\x2d\x7e\xac\xe6\x3e\x50\x93\x09\xb1\x14\x62\x18\xd6\x88\xca\x19\x42\xfe\xba\x42\xde\x61\xf3\x07\xc5\x21\x05\xc5\x38\x4e\xd3\x3c\x38\x55\x53\xe2\x79\x57\xec\x52\xf9\x01\xf6\x07\xd5\x68\x55\xa5\x1d\x29\xf1\x62\xf9\x32\x0c\xf7\xcd\x71\xc4\x9b\xad\x2a\xa8\x10\xaa\x1e\x89\xdf\x74\xe3\x88\x4e\x55\x2d\x81\x3d\xf6\x61\xfa\xa9\x42\x48\x29\xe6\x95\x2d\x35\x3a\xda\x09\x27\xdb\x73\xeb\xf4\x37\xe3\x45\x77\xf3\x3b\xf6\x21\xf5\x4b\xa4\xa8\x18\xa6\xfb\x23\x71\x1b\xdd\x44\x66\x90\xde\x13\xf5\x1b\x64\x5f\x42\x64\xab\xc7\xf4\x39\xc5\xd9\xa0\x9b\xc9\x28\xe6\x8b\xc9\x95\x29\xc1\x67\x00\x00\x00\xff\xff\x81\x9d\x53\x8d\xf6\x01\x00\x00")

func cmdDefinitionsTmplPairTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplPairTmpl,
		"cmd/definitions/tmpl/pair.tmpl",
	)
}

func cmdDefinitionsTmplPairTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplPairTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/pair.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1b, 0x1a, 0xbf, 0xf0, 0xe9, 0x47, 0x35, 0xe6, 0xa7, 0xc, 0xe6, 0xba, 0xc0, 0x5e, 0x39, 0xda, 0x2d, 0xb9, 0xd0, 0xee, 0xbc, 0x1, 0xc0, 0xd1, 0x5e, 0xa2, 0xbf, 0x4d, 0xcd, 0x4, 0xe9, 0xc0}}
	return a, nil
}

var _cmdDefinitionsTmplServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4b\x6f\x1c\xb9\x11\xbe\xcf\xaf\xa8\x6d\x08\x8b\x19\x63\xb6\x3b\x01\x72\x52\xa0\xc3\x46\xf6\x3a\xc2\xfa\x21\x58\xde\xec\xc1\xbb\x10\xa8\xee\xea\x19\x46\xdd\x64\x9b\xe4\x8c\x2c\x8c\xfb\xbf\x07\x7c\xf4\x9b\x9c\x87\xd7\x59\x3b\x81\x7d\xf1\x88\x2c\x16\x8b\xf5\xf8\xaa\x8a\xcd\x24\x81\x4b\x9e\x21\xac\x90\xa1\x20\x0a\x33\xb8\x7b\x84\x15\x6f\xff\x86\x2d\x25\x90\x96\x59\x92\x61\x4e\x19\x55\x94\x33\xf9\x77\x78\xfa\x1a\x5e\xbd\x7e\x0b\xcf\x9e\x5e\xbd\x8d\x67\x15\x49\xef\xc9\x0a\x61\xb7\x83\xf8\x15\x29\x11\xea\x7a\x36\xa3\x65\xc5\x85\x82\xf9\x0c\x00\x20\x4a\x39\x53\xf8\x41\x45\xf6\x2f\xca\xa3\x99\xfd\xb5\xa2\x6a\xbd\xb9\x8b\x53\x5e\x26\x77\xf8\xc8\x59\x26\x15\x17\x64\x85\xc9\x8a\xff\xd0\xfc\xdc\xfe\x2d\xa9\xee\x57\x09\xb2\xac\xe2\x94\x35\x3c\x4e\x58\x99\x0a\xcc\x90\x29\x4a\x8a\xd3\xd7\xae\x95\xaa\xd2\x82\xe2\xa9\xfb\x4a\x14\x5b\x9a\xa2\xb4\xab\xe2\xa3\xd7\xa9\xc7\x4a\x2f\x5a\xcc\x66\x5b\x22\xe0\x16\x3a\xd9\xe3\x6b\xc1\xb7\x34\x43\xe1\x66\x1a\x7d\xc4\xff\x22\xc5\x06\xdd\xe0\x8d\xe5\xd4\xd0\x34\x52\xc4\x37\xf6\xc7\x33\x21\x78\x33\xd7\x9d\x2c\x7e\x5d\x19\xb3\xce\x66\x49\x02\x6f\x1f\x2b\x04\x2a\x41\xad\x11\xb4\x30\x90\x73\x31\xb0\x6c\xca\x99\x54\x96\xec\x02\xa2\xde\x4c\x64\xd6\xbf\xbe\xfb\x37\xa6\xea\x25\x2a\x92\x11\x45\x40\x1f\x0d\x65\x23\x08\x94\xcd\xb8\xe6\xca\x0d\x69\x3c\x4b\x12\xbd\xf0\x29\x56\x02\x53\xed\x81\xe7\xf0\x8b\x44\xc7\xe8\xe6\x51\x2a\x2c\x5b\x76\x94\x49\x85\x24\x8b\x67\x46\xb4\xc9\x5e\x62\x93\x2a\xd8\xcd\x76\xbb\x1f\x40\x10\xb6\x42\x38\xbb\x5d\xc2\xd9\x16\xce\x2f\x20\xbe\x62\x39\x97\xfa\x00\xda\x20\x9a\x82\xe6\x80\xef\xe1\x6c\x1b\xdf\xa4\xbc\x42\x88\xac\x38\xd1\x88\x84\x71\xa5\x69\x9e\x17\xfc\x8e\x14\xfd\xb9\xb3\x8a\xe9\x73\x9f\x5f\xe8\x69\xa3\x82\x8f\xa0\xf8\x35\x91\xe9\x90\x8e\xe6\x9a\xe0\x29\x95\x55\x41\x1e\x1b\x25\x82\xfb\xd7\x63\x74\x11\x20\xd3\x24\xc8\xb2\xee\x4f\x43\x87\x32\x15\xd4\x98\xad\x3f\x61\x39\xd5\xb5\xa3\x32\x46\x0a\xb1\x69\xff\xec\xfd\xac\x8d\x09\x9f\xa3\x1a\x69\xf6\x81\x16\x05\xac\x50\x8d\x35\x9e\x0b\x5e\xba\xb1\xc6\x8c\x3f\xc0\xdb\x35\x95\x90\x6f\x58\x6a\xa4\x93\x6b\xbe\x29\x32\xa3\xc6\x3b\x84\x94\x14\x85\x85\x98\xc6\x23\x68\x59\x15\x58\x22\x53\x28\xe2\x66\x3d\x82\x40\xb5\x11\x8c\xb2\xd5\x78\x47\x2a\x41\x20\xc9\x80\xb3\xe2\x11\x08\xcb\x46\xfc\x4b\x9e\xd1\x9c\x62\x16\x72\xaa\xf6\x68\x21\xbf\xd2\x72\x4f\x15\x30\xe7\xf0\xc4\x8e\x2c\xc6\x02\xed\x8c\x3a\x79\xb9\x04\x7e\xaf\x9d\x81\xc7\xcf\x51\xb9\x70\x6b\x97\x2f\x0c\x11\xcd\x35\xcd\xae\x35\xbe\x3d\x24\xf0\x32\x9e\x0f\x99\x5a\x72\x6b\x28\x47\x34\x24\xd8\x35\xa6\x92\x7e\x53\xc9\xa9\xa9\x28\x53\xfc\x38\x53\x19\xd5\x76\xb6\xe2\x2c\xc5\x25\x54\x05\x12\x89\x50\x92\x7b\x04\xb9\x11\x08\xa4\x28\xc0\x30\x5e\x13\x09\x77\x88\x0c\x1e\x04\x55\x0a\x19\xdc\x61\xce\x05\x6a\x19\x42\x56\x90\x47\x59\x61\x72\xb6\xce\x0a\x4b\x68\xfd\xae\xd5\x59\x63\x89\xf8\x66\xaa\x7f\x5e\x2e\x9c\xc2\xbc\xfb\x36\x20\x65\x46\xfd\x18\xd5\xc3\x9b\xc9\xe2\x6f\xa8\xf3\x59\x51\x67\xa4\xdf\x11\xf6\x8c\x66\xff\x34\x04\x1a\xb9\xd4\xf1\x40\x34\x44\x94\xa1\xf8\x53\x5c\x19\x1d\xcf\xfa\xb4\x1c\xa1\xcb\x90\xc7\x5e\x70\x91\x2d\xb8\x0c\x57\x05\x21\x66\x48\x36\x05\x1a\x9f\x75\x64\xc8\x3a\x5f\x06\x74\x86\xe8\x11\xd2\xf8\x12\x64\xe9\x95\x7a\x88\x24\xc3\xd5\xb2\x05\x12\x57\x69\x1d\x8f\x24\xae\xc8\x33\x83\x0e\x4f\x42\x3c\xfe\x10\xa0\xb8\x7d\xa2\x71\xc0\x87\x50\xc5\x03\x0a\x7b\xd1\x65\xc4\x33\x8c\x32\x27\xa0\x4d\x9f\xb4\x07\x17\xdd\x70\x00\x7d\xfa\x04\x7b\x50\xe8\x74\x24\xf2\x5b\xa6\x85\x22\x1f\x08\xb9\x25\x2f\x8d\x79\x3b\x5f\xef\x23\x89\x9f\xeb\xc9\x48\xe2\x65\x33\x97\xf0\xa4\x27\xc2\x22\xb0\xd9\x18\x50\xe4\xa7\x00\x8a\x97\xb5\x07\x51\xbc\x74\x7d\x48\xd9\xa3\x66\x39\x55\xb3\x41\x93\x80\x9a\xff\x24\x48\x39\x46\xf5\x06\x58\xfc\x2a\x6a\xd4\x1f\x46\x16\x6f\xcc\x5f\x13\x2a\x64\xe3\xaa\xbe\x38\x3e\x26\x76\xbb\x78\xbd\xe4\x2c\x2f\x68\xaa\xf4\xa8\xa7\x3a\xab\xf4\x6e\xf1\xaf\x54\xad\x07\x61\xd5\x16\x67\x49\xd2\x0f\x99\x24\x81\x09\xa9\x31\x20\xa9\xaa\xe2\xd1\x05\xa3\x8b\x75\xd8\xea\x56\x15\x74\x4e\xb0\x1d\xa7\x63\xe6\x89\x6e\xa3\xf0\x31\xe3\xf9\x76\x18\xdc\x0b\xd0\x9a\x99\xfa\xa8\x1e\xed\x06\x7f\xc6\xc7\x73\xd3\xa8\x76\x82\x44\xcb\x76\xd6\x74\xcf\xe7\xb0\x5d\x3a\xf7\x1d\xe0\x41\xef\xa7\xe9\x9a\xb5\x6a\x5e\x92\x0a\x2e\xa0\x24\xd5\x3b\xa9\x04\x65\xab\xdf\xed\x7f\x4e\x8a\xfd\xf6\x3b\x05\x67\x47\x12\xb7\x47\x70\x67\x77\x47\x18\x82\x97\x77\x77\xcd\x41\x56\x24\xc5\x81\x08\x0a\xcb\xaa\x20\x0a\x21\xa2\xba\xf2\xc9\xf5\x7c\x64\xe2\xe2\xa6\xd0\x55\x51\xb3\xf5\xd9\x36\xbe\x6a\x09\xfc\x0c\x72\x24\x6a\x23\x82\xcb\x7f\xda\xb0\x34\xb0\x52\xeb\xf3\x96\xe1\x43\x7f\xe5\xdc\xa7\x93\x85\x19\xc4\x87\x30\x9b\x4f\xd8\xbd\x01\x8c\x83\x4b\xfb\x6e\xa0\x7f\x9b\xab\xb0\x91\xe6\xba\x28\xd4\x4a\xa7\x2c\xc3\x0f\x10\xc3\x5f\xda\x71\x43\x2b\xfb\x73\x7f\xd5\x73\xda\xab\xe6\x41\xcf\x69\x56\x39\xf1\x6f\xfb\xe1\x34\xf0\x18\xb8\x80\xef\x6d\xa4\x0c\xc7\x77\x93\x6c\xb7\x08\x9e\xa6\x33\x63\xe8\x2c\x53\x40\x39\xcb\xd9\xf4\x4c\xb6\xb8\x71\xe2\xd4\xf5\x4f\x8e\x6f\x57\xd8\x68\x91\x5e\x70\x2e\xf1\x75\x85\x82\x68\x13\xfc\x58\x14\x70\xc7\x79\x11\xd6\x84\xde\xc8\xa9\x61\xb8\x34\xa8\x92\x01\xbf\xe6\xc0\x26\xe4\xa9\x50\x1b\x52\x1c\xb5\x79\xce\xc6\xbb\x37\x38\x9a\xb3\xf8\x86\x96\x9b\xc2\xdc\x99\xba\xb9\x31\x67\x2d\x5b\xce\x8e\x15\x6e\x8f\xac\x1a\x41\x3e\x49\xcc\x02\x99\x11\xc1\xf1\x19\xca\x69\x70\x32\x24\xe2\x64\xbb\x91\x94\x1e\x0b\x35\xbb\xc0\x09\xdb\x9c\x64\xbe\x09\xe2\x85\x7c\xb9\x03\x96\x43\x71\x99\xb3\xa9\xff\xba\x89\xfb\xe6\x58\x1e\x84\xd6\x79\xaf\x72\xe7\x32\x5e\x6e\x0f\x78\x6f\x32\xa5\xbd\x3d\xad\x88\x90\x98\x39\xaf\xb7\x31\x31\x5a\xa1\x17\xd4\xf5\x30\x2e\x0c\x96\xc1\xbb\xdf\xb5\xd2\xac\x0b\x24\x09\xbc\xc1\xf7\x1b\x2a\x30\xb3\xb3\x3e\x13\xe8\x89\x46\xdc\x96\xda\x69\xec\x9f\x44\x9a\x4d\x09\x15\x3e\x45\x43\x5f\xd5\xfb\xc8\xba\xe9\x3d\x2d\x7e\x92\xb8\xe4\x4e\x8a\xe3\xa4\x6d\xa9\x3f\xb7\xb4\x87\xe5\x6d\xcc\x28\x24\x5e\x07\x6c\x69\x4a\x19\x43\x61\x6b\x0d\x69\x92\x84\xa9\x44\x9f\x04\xec\x6f\x6b\x97\x7d\x5c\xe7\xbc\x52\x8d\x8d\x17\x30\x0f\xf0\x59\x02\x0a\xc1\x45\x53\x34\x0a\x94\x9b\x42\x69\xad\x05\xe8\xbb\x8a\xc7\xe8\xfd\x1c\xf4\x2e\x4d\x59\x63\xfe\xd3\x6d\xe8\xed\x12\x4c\xb8\x5a\x6b\x18\x41\xba\x85\xf2\x81\xaa\x74\x0d\xdb\xf8\x67\x7c\xec\x0d\xfb\x3d\xf0\x44\x2f\xd4\xff\x52\x5d\x7b\x47\x43\xbb\xe9\xda\x66\xd0\xcc\xd1\xdc\x9d\x35\x3e\xe0\x0a\xbb\x49\xab\x99\x72\xa6\x28\xdb\xe0\x60\x62\xd8\x2a\x1e\xc7\xfa\x02\x94\x18\xb1\x71\x0b\xf7\xaf\xda\xda\xaf\x31\xf1\x7c\xec\x7a\x8b\x50\x23\x1a\x8e\x99\x13\xe3\xe6\x9b\x7e\x83\xfa\xad\xfb\x41\x70\x82\xcb\xd2\x1c\xbe\x3b\x55\x55\xae\x05\x09\xc5\x68\xbd\xec\x3e\xca\xe9\xf0\x6f\xb6\x34\x5f\xe6\x76\xba\x57\x91\xe7\xef\x5c\x4b\xb1\xf3\x98\x12\xea\xf6\x34\xe0\x29\x17\xdc\xee\x56\xe8\x25\x30\x5a\x1c\x48\x92\x32\x1a\xf5\x25\xfb\x2a\xbe\x96\xcc\x34\xcc\x93\xb4\xe9\x5c\xf9\x29\xe6\x64\x53\xa8\xf6\xf4\xb6\xfb\xa1\x52\xef\xab\x27\x5c\x8a\x33\x77\x62\x15\xa6\x34\xa7\x29\x10\x53\x88\x1b\x0e\x26\x51\xfa\x79\x0c\x12\xa5\xb7\x06\x31\x82\x0d\x6e\x7d\x02\xd5\x85\x4b\xb0\x1e\x87\xf1\xbb\x49\x88\xfd\xb8\x0d\x19\xb4\x12\x15\x83\xb3\xed\xde\x0b\xa7\x89\x3d\xbe\xee\x82\xa5\x5d\x71\xa0\x64\xf9\x12\xc5\xc9\x71\xe5\xc9\xd7\x59\x88\x1c\x90\x6c\x54\x56\xff\x9f\x54\x48\x73\x09\x4f\xda\xc9\xc5\xff\x64\xbd\x94\x24\x40\xe5\x2f\x4c\x6e\xaa\x8a\x0b\x85\x99\x51\x80\xc0\x94\x8b\x4c\xc2\xc3\x1a\xd5\x1a\x05\xa4\x1b\x21\x90\x59\xd8\xd3\xa1\xb6\xe9\xe8\xe3\x96\xd3\x94\xcd\xf9\x05\xe4\xa4\x90\x38\x3b\x54\x9d\x7d\xab\xc1\x4e\xae\x11\xbc\x72\x7a\x0a\x87\x6f\xe5\xd7\x97\x50\xed\x08\xee\xbe\x12\xcd\x26\x09\x5c\xe5\xb0\x91\x28\x00\x19\xb9\x2b\xd0\x26\xcd\xad\x13\xd6\x1c\xc0\x5d\xa1\x2d\xe1\x01\x21\x25\xba\x0a\x94\x1e\x2a\x83\x8d\x54\xc5\x63\xe1\x65\xdc\x5c\xc1\xc5\xa3\x6b\x9f\x8f\x1f\x03\x93\x1d\x9a\x1d\x49\x76\x9a\x5a\x3e\xd9\x9b\x3e\xb7\x47\x1d\x69\xa3\x29\x8a\x8e\x64\xf3\x38\xa2\xab\x4a\xcf\x4f\x63\x54\x77\xb0\xac\xbb\x84\x29\xfd\x6e\x7f\x44\xf4\xd6\x07\x1c\xab\xe0\x5c\xea\x84\xe3\xae\x12\x07\xbe\x65\x72\x2e\x5d\x31\x2e\x6c\xd2\xed\x6d\x6e\x1a\x89\x78\xe6\x77\xab\xe9\xa5\xeb\xd0\x65\xa6\x37\xab\x2e\xc9\x1f\x3c\x4d\x67\xf4\x93\xda\x9f\xb1\xe0\x3b\x3d\x78\xbe\x1d\x54\xe2\x49\x02\x97\x6b\x4c\xef\x41\x0c\xee\x20\xe2\x6f\xdd\xdc\xf0\xfa\xbe\xfb\x9a\xf1\x59\x3b\xba\xa3\x1b\x21\xfb\x71\xf1\x05\x1f\xbd\x15\x18\xf6\x48\x85\x9e\xbe\x6d\x64\x0d\x76\x4b\xad\x36\x0a\x39\xf9\x8a\x3f\xfd\x80\x73\x98\xcd\xa1\x0f\xff\x5e\x4d\x7e\x6a\x2f\x76\x4c\x2b\xd6\x8f\x2d\x7b\x35\x3e\xfa\x00\x6a\x3f\x6d\x0f\x3f\x6c\x9b\xb0\x4f\x05\xea\xc3\x13\x70\x6f\xaa\xe1\xee\xb1\xc1\xb0\xb8\x5f\x58\x9f\x55\xac\xae\x17\xbd\x8d\xe6\xe6\xc2\x39\xbe\x26\x82\x94\x32\xbe\x31\x4e\xa9\x29\xdc\xf8\x1b\xe3\x65\xfd\x09\x1b\x06\xa9\xfa\xa0\xcf\xe2\x76\x8b\xff\x41\xd2\xfb\x95\xe0\x1b\x96\xb9\x77\x02\xcd\xbb\x80\xb8\xdb\xe9\x57\xaa\xd6\x97\x96\x7e\x9e\xaa\x0f\x4b\x18\xec\x7c\x49\x8a\x02\x85\x46\xf8\xb1\x2a\x7a\xeb\x02\x5a\xd9\x77\xbe\xd1\xae\xad\xc4\x6e\x6c\x24\xc5\xd1\xe7\xcf\x30\x47\x61\x6c\x30\x5f\x0c\x8b\xee\xb3\x8a\xa8\x75\x63\x67\xc7\xf6\x9a\xa8\xb5\x3d\xa0\x27\x3c\x08\xcb\x60\x8e\xef\xdd\xc2\x28\x5a\xb8\xbf\x18\x44\xee\x79\x59\xb4\xf0\x3e\x6d\xd1\xe4\x17\x10\x2d\x7f\x8b\x7e\x8b\x26\x0f\x7b\x46\xf9\x0c\x85\x4e\x56\x32\xce\xb9\x28\x89\x32\x48\x34\x8f\xec\x11\xb5\x2f\xd6\x75\x64\x1a\xa5\x8e\x71\x5d\x83\x7b\xbb\x31\x5f\x74\x21\x1f\xfa\xc4\x65\xb3\x55\x53\xcf\x4c\xd2\x93\x6e\x70\x18\x57\x2e\x93\x65\x5d\xb6\x22\xa9\x82\x82\xde\x23\x28\xed\xd2\xdd\x3a\x47\xdf\x3d\xc2\x73\x4d\x91\x46\xeb\x69\x31\x33\xc8\x5d\xdf\x7f\xbf\x97\xc4\x9b\xbe\xac\x76\x5e\xe1\x43\x4b\xf6\x8a\xab\xab\x6e\xf3\x56\x5f\x5d\xf0\xd6\x75\xb4\x18\x25\x84\x3d\x98\xdd\x3c\xd1\x40\xdb\x25\xd8\x37\xaa\x3c\x43\x88\x86\xb8\x5c\xea\xb1\xa6\x97\xe8\xa8\xa6\xf0\xac\x15\xc1\x63\x3d\x19\x5f\x99\xac\x65\x56\xd6\xf5\x7c\x31\x39\x56\x9b\x8b\x3a\x86\x57\x6c\x4b\x0a\xea\x12\xd2\xb3\x0f\x15\xa6\xe6\x9d\x87\x9e\xea\xf1\x5a\xc2\x8f\xa9\xd6\xdd\x39\xd8\x9d\xc6\xf9\x7c\xcf\x79\xed\xe5\xcf\x05\x90\xaa\x42\x96\x99\xde\x5c\x2e\x41\xc6\x0e\x92\xcc\xc5\x5d\x0f\x1a\xe2\x38\xb6\xca\xdc\x12\xa1\xbb\xe9\x50\x6a\xb5\xcc\x79\xa5\x96\xad\x43\xef\xbd\x21\x30\xfb\xb6\x0f\x97\xf4\x92\xef\x2e\x74\xc2\x9c\xe4\xf2\x7e\x65\x31\x82\xae\x0e\xab\x2f\x49\x89\xda\x00\x1e\xf4\x7a\x2b\x68\xf9\x82\x48\xe5\x60\xec\x19\xcb\x74\x5b\xbe\xbe\xe4\x65\x49\xea\x5a\x4b\xbc\xd8\x93\xa4\xc7\xe9\x6f\x5f\xa6\xee\xcf\xf9\x12\x4d\x83\x0e\x07\x92\x8d\x0b\xda\x83\x09\xc7\xd2\x35\xe4\xa7\x26\x1e\xbd\xec\xbf\x90\x7c\xfe\x88\x87\x1d\xed\x65\x0d\xac\xd9\x7a\xda\x5c\x22\xc1\xc3\x9a\x16\x08\x6b\xc2\xb2\x82\xb2\x15\x18\xbb\xe9\x03\xba\x67\x4b\xcd\x32\xe3\xa0\xb7\x47\xbb\xe7\xb8\xae\x34\x72\x07\x5c\xef\x64\xaf\xb3\x8e\xdd\xf7\x3c\x63\x10\xca\xa8\x6a\xb1\xe2\xc8\x07\x42\xfa\x5f\x0b\x26\x6f\x70\x45\xa5\x42\x11\xba\x4b\x17\x73\xdd\xb2\x2d\x35\xa8\x06\x49\x16\xbe\x1b\xcf\xc9\x0e\x37\xe9\x1a\x4b\xe2\xd8\xb9\xb7\x56\x3a\x9a\xfe\x13\x00\x00\xff\xff\x03\x8f\x85\x75\xa4\x36\x00\x00")

func cmdDefinitionsTmplServiceTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplServiceTmpl,
		"cmd/definitions/tmpl/service.tmpl",
	)
}

func cmdDefinitionsTmplServiceTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplServiceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/service.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd3, 0x65, 0x44, 0x4a, 0x31, 0xc6, 0x2c, 0x89, 0x1c, 0xdd, 0x27, 0x50, 0xa9, 0xa9, 0x2e, 0x63, 0xf0, 0xf7, 0x6b, 0x93, 0x55, 0x7b, 0x35, 0x6c, 0x5a, 0xaa, 0xfd, 0x5e, 0x9d, 0xbe, 0xcb, 0x90}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"cmd/definitions/tmpl/function.tmpl":  cmdDefinitionsTmplFunctionTmpl,
	"cmd/definitions/tmpl/info.tmpl":      cmdDefinitionsTmplInfoTmpl,
	"cmd/definitions/tmpl/object.tmpl":    cmdDefinitionsTmplObjectTmpl,
	"cmd/definitions/tmpl/operation.tmpl": cmdDefinitionsTmplOperationTmpl,
	"cmd/definitions/tmpl/pair.tmpl":      cmdDefinitionsTmplPairTmpl,
	"cmd/definitions/tmpl/service.tmpl":   cmdDefinitionsTmplServiceTmpl,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"cmd": {nil, map[string]*bintree{
		"definitions": {nil, map[string]*bintree{
			"tmpl": {nil, map[string]*bintree{
				"function.tmpl": {cmdDefinitionsTmplFunctionTmpl, map[string]*bintree{}},
				"info.tmpl": {cmdDefinitionsTmplInfoTmpl, map[string]*bintree{}},
				"object.tmpl": {cmdDefinitionsTmplObjectTmpl, map[string]*bintree{}},
				"operation.tmpl": {cmdDefinitionsTmplOperationTmpl, map[string]*bintree{}},
				"pair.tmpl": {cmdDefinitionsTmplPairTmpl, map[string]*bintree{}},
				"service.tmpl": {cmdDefinitionsTmplServiceTmpl, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
