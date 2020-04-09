// Code generated by go generate via internal/cmd/service; DO NOT EDIT.
package fs

import (
	"context"
	"io"

	"github.com/opentracing/opentracing-go"

	"github.com/Xuanwo/storage"
	"github.com/Xuanwo/storage/pkg/credential"
	"github.com/Xuanwo/storage/pkg/endpoint"
	"github.com/Xuanwo/storage/pkg/segment"
	"github.com/Xuanwo/storage/pkg/storageclass"
	"github.com/Xuanwo/storage/services"
	"github.com/Xuanwo/storage/types"
	"github.com/Xuanwo/storage/types/metadata"
	ps "github.com/Xuanwo/storage/types/pairs"
)

var _ credential.Provider
var _ endpoint.Provider
var _ segment.Segment
var _ storage.Storager
var _ storageclass.Type
var _ services.ServiceError

// Type is the type for fs
const Type = "fs"

var pairServiceNewMap = map[string]struct{}{}

type pairServiceNew struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func parseServicePairNew(opts ...*types.Pair) (*pairServiceNew, error) {
	result := &pairServiceNew{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairServiceNewMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs

	return result, nil
}

var pairStorageCopyMap = map[string]struct{}{}

type pairStorageCopy struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func (s *Storage) parsePairCopy(opts ...*types.Pair) (*pairStorageCopy, error) {
	result := &pairStorageCopy{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageCopyMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs

	return result, nil
}

var pairStorageDeleteMap = map[string]struct{}{}

type pairStorageDelete struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func (s *Storage) parsePairDelete(opts ...*types.Pair) (*pairStorageDelete, error) {
	result := &pairStorageDelete{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageDeleteMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs

	return result, nil
}

var pairStorageListDirMap = map[string]struct{}{
	"dir_func":  struct{}{},
	"file_func": struct{}{},
}

type pairStorageListDir struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasDirFunc  bool
	DirFunc     types.ObjectFunc
	HasFileFunc bool
	FileFunc    types.ObjectFunc
}

func (s *Storage) parsePairListDir(opts ...*types.Pair) (*pairStorageListDir, error) {
	result := &pairStorageListDir{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageListDirMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs
	v, ok = values[ps.DirFunc]
	if ok {
		result.HasDirFunc = true
		result.DirFunc = v.(types.ObjectFunc)
	}
	v, ok = values[ps.FileFunc]
	if ok {
		result.HasFileFunc = true
		result.FileFunc = v.(types.ObjectFunc)
	}

	return result, nil
}

var pairStorageMetadataMap = map[string]struct{}{}

type pairStorageMetadata struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func (s *Storage) parsePairMetadata(opts ...*types.Pair) (*pairStorageMetadata, error) {
	result := &pairStorageMetadata{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageMetadataMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs

	return result, nil
}

var pairStorageMoveMap = map[string]struct{}{}

type pairStorageMove struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func (s *Storage) parsePairMove(opts ...*types.Pair) (*pairStorageMove, error) {
	result := &pairStorageMove{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageMoveMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs

	return result, nil
}

var pairStorageNewMap = map[string]struct{}{
	"work_dir": struct{}{},
}

type pairStorageNew struct {
	// Pre-defined pairs

	// Meta-defined pairs
	HasWorkDir bool
	WorkDir    string
}

func parseStoragePairNew(opts ...*types.Pair) (*pairStorageNew, error) {
	result := &pairStorageNew{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageNewMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs

	// Parse meta-defined pairs
	v, ok = values[ps.WorkDir]
	if ok {
		result.HasWorkDir = true
		result.WorkDir = v.(string)
	}

	return result, nil
}

var pairStorageReadMap = map[string]struct{}{
	"offset":             struct{}{},
	"read_callback_func": struct{}{},
	"size":               struct{}{},
}

type pairStorageRead struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasOffset           bool
	Offset              int64
	HasReadCallbackFunc bool
	ReadCallbackFunc    func([]byte)
	HasSize             bool
	Size                int64
}

func (s *Storage) parsePairRead(opts ...*types.Pair) (*pairStorageRead, error) {
	result := &pairStorageRead{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageReadMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs
	v, ok = values[ps.Offset]
	if ok {
		result.HasOffset = true
		result.Offset = v.(int64)
	}
	v, ok = values[ps.ReadCallbackFunc]
	if ok {
		result.HasReadCallbackFunc = true
		result.ReadCallbackFunc = v.(func([]byte))
	}
	v, ok = values[ps.Size]
	if ok {
		result.HasSize = true
		result.Size = v.(int64)
	}

	return result, nil
}

var pairStorageStatMap = map[string]struct{}{}

type pairStorageStat struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func (s *Storage) parsePairStat(opts ...*types.Pair) (*pairStorageStat, error) {
	result := &pairStorageStat{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageStatMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs

	return result, nil
}

var pairStorageWriteMap = map[string]struct{}{
	"read_callback_func": struct{}{},
	"size":               struct{}{},
}

type pairStorageWrite struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasReadCallbackFunc bool
	ReadCallbackFunc    func([]byte)
	HasSize             bool
	Size                int64
}

func (s *Storage) parsePairWrite(opts ...*types.Pair) (*pairStorageWrite, error) {
	result := &pairStorageWrite{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageWriteMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs
	v, ok = values[ps.ReadCallbackFunc]
	if ok {
		result.HasReadCallbackFunc = true
		result.ReadCallbackFunc = v.(func([]byte))
	}
	v, ok = values[ps.Size]
	if ok {
		result.HasSize = true
		result.Size = v.(int64)
	}

	return result, nil
}

// CopyWithContext adds context support for Copy.
func (s *Storage) CopyWithContext(ctx context.Context, src, dst string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/fs.storage.Copy")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Copy(src, dst, pairs...)
}

// DeleteWithContext adds context support for Delete.
func (s *Storage) DeleteWithContext(ctx context.Context, path string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/fs.storage.Delete")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Delete(path, pairs...)
}

// ListDirWithContext adds context support for ListDir.
func (s *Storage) ListDirWithContext(ctx context.Context, path string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/fs.storage.ListDir")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.ListDir(path, pairs...)
}

// MetadataWithContext adds context support for Metadata.
func (s *Storage) MetadataWithContext(ctx context.Context, pairs ...*types.Pair) (m metadata.StorageMeta, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/fs.storage.Metadata")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Metadata(pairs...)
}

// MoveWithContext adds context support for Move.
func (s *Storage) MoveWithContext(ctx context.Context, src, dst string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/fs.storage.Move")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Move(src, dst, pairs...)
}

// ReadWithContext adds context support for Read.
func (s *Storage) ReadWithContext(ctx context.Context, path string, pairs ...*types.Pair) (r io.ReadCloser, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/fs.storage.Read")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Read(path, pairs...)
}

// StatWithContext adds context support for Stat.
func (s *Storage) StatWithContext(ctx context.Context, path string, pairs ...*types.Pair) (o *types.Object, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/fs.storage.Stat")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Stat(path, pairs...)
}

// WriteWithContext adds context support for Write.
func (s *Storage) WriteWithContext(ctx context.Context, path string, r io.Reader, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/fs.storage.Write")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Write(path, r, pairs...)
}
