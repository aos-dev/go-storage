package posixfs

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"syscall"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/Xuanwo/storage/pkg/iterator"
	"github.com/Xuanwo/storage/types"
)

func TestNewClient(t *testing.T) {
	c := NewClient()
	assert.NotNil(t, c)
}

func TestClient_Init(t *testing.T) {
	t.Run("without options", func(t *testing.T) {
		client := Client{}
		err := client.Init()
		assert.Error(t, err)
		assert.Equal(t, "", client.base)
	})

	t.Run("with base", func(t *testing.T) {
		client := Client{}
		err := client.Init(types.WithBase("test"))
		assert.NoError(t, err)
		assert.Equal(t, "test", client.base)
	})
}

func TestClient_Capable(t *testing.T) {
	client := Client{}
	assert.True(t, client.Capable("read"))
	assert.True(t, client.Capable("list_dir", types.Recursive))
	assert.False(t, client.Capable("list_dir", types.Expire))
	assert.False(t, client.Capable("reach"))
	assert.False(t, client.Capable("read", types.Expire))
}

func TestClient_Metadata(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	{
		client := Client{}

		m, err := client.Metadata()
		assert.NoError(t, err)
		assert.Equal(t, 0, len(m))
	}
}

type fileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (f fileInfo) Name() string {
	return f.name
}

func (f fileInfo) Size() int64 {
	return f.size
}

func (f fileInfo) Mode() os.FileMode {
	return f.mode
}

func (f fileInfo) ModTime() time.Time {
	return f.modTime
}

func (f fileInfo) IsDir() bool {
	return f.mode.IsDir()
}

func (f fileInfo) Sys() interface{} {
	return f
}

func TestClient_Stat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name string
		err  error
		file fileInfo

		object *types.Object
	}{
		{
			"regular file",
			nil,
			fileInfo{
				name:    "regular file",
				size:    1234,
				mode:    0777,
				modTime: time.Now(),
			},
			&types.Object{
				Name: "regular file",
				Type: types.ObjectTypeFile,
				Metadata: types.Metadata{
					types.Size: int64(1234),
				},
			},
		},
		{
			"dir",
			nil,
			fileInfo{
				name:    "dir",
				size:    0,
				mode:    os.ModeDir | 0777,
				modTime: time.Now(),
			},
			&types.Object{
				Name:     "dir",
				Type:     types.ObjectTypeDir,
				Metadata: make(types.Metadata),
			},
		},
		{
			"stream",
			nil,
			fileInfo{
				name:    "stream",
				size:    0,
				mode:    os.ModeDevice | 0777,
				modTime: time.Now(),
			},
			&types.Object{
				Name:     "stream",
				Type:     types.ObjectTypeStream,
				Metadata: make(types.Metadata),
			},
		},
		{
			"invalid",
			nil,
			fileInfo{
				name:    "invalid",
				size:    0,
				mode:    os.ModeIrregular | 0777,
				modTime: time.Now(),
			},
			&types.Object{
				Name:     "invalid",
				Type:     types.ObjectTypeInvalid,
				Metadata: make(types.Metadata),
			},
		},
		{
			"error",
			&os.PathError{
				Op:   "stat",
				Path: "invalid",
				Err:  os.ErrPermission,
			},
			fileInfo{}, nil,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			client := Client{
				osStat: func(name string) (os.FileInfo, error) {
					assert.Equal(t, v.name, name)
					return v.file, v.err
				},
			}
			o, err := client.Stat(v.name)
			assert.Equal(t, v.err == nil, err == nil)
			if v.object != nil {
				assert.NotNil(t, o)
				assert.EqualValues(t, v.object, o)
			} else {
				assert.Nil(t, o)
			}
		})
	}
}

func TestClient_WriteStream(t *testing.T) {
	err := os.Remove("/tmp/test")
	var e *os.PathError
	if errors.As(err, &e) {
		t.Logf("%#v", e)
	}
}

func TestClient_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name      string
		err       error
		recursive bool
	}{
		{"delete file", nil, false},
		{"delete dir", nil, true},
		{"delete nonempty dir", &os.PathError{
			Op:   "remove",
			Path: "delete nonempty dir",
			Err:  syscall.ENOTEMPTY,
		}, false},
	}

	for _, v := range tests {
		v := v

		t.Run(v.name, func(t *testing.T) {

			client := Client{
				osRemove: func(name string) error {
					assert.Equal(t, v.name, name)
					assert.False(t, v.recursive)
					return v.err
				},
				osRemoveAll: func(name string) error {
					assert.Equal(t, v.name, name)
					assert.True(t, v.recursive)
					return v.err
				},
			}
			pairs := make([]*types.Pair, 0)
			if v.recursive {
				pairs = append(pairs, types.WithRecursive(true))
			}
			err := client.Delete(v.name, pairs...)
			assert.Equal(t, v.err == nil, err == nil)
		})
	}
}

func TestClient_Copy(t *testing.T) {
	t.Run("Failed at open source file", func(t *testing.T) {
		srcName := uuid.New().String()
		dstName := uuid.New().String()
		client := Client{
			osOpen: func(name string) (file *os.File, e error) {
				assert.Equal(t, srcName, name)
				return nil, &os.PathError{
					Op:  "open",
					Err: syscall.ENONET,
				}
			},
		}

		err := client.Copy(srcName, dstName)
		assert.Error(t, err)
	})

	t.Run("Failed at open dst file", func(t *testing.T) {
		srcName := uuid.New().String()
		dstName := uuid.New().String()
		client := Client{
			osOpen: func(name string) (file *os.File, e error) {
				assert.Equal(t, srcName, name)
				return nil, nil
			},
			osCreate: func(name string) (file *os.File, e error) {
				assert.Equal(t, dstName, name)
				return nil, &os.PathError{
					Op:  "open",
					Err: syscall.EEXIST,
				}
			},
		}

		err := client.Copy(srcName, dstName)
		assert.Error(t, err)
	})

	t.Run("Failed at io.CopyBuffer", func(t *testing.T) {
		srcName := uuid.New().String()
		dstName := uuid.New().String()
		client := Client{
			osOpen: func(name string) (file *os.File, e error) {
				assert.Equal(t, srcName, name)
				return nil, nil
			},
			osCreate: func(name string) (file *os.File, e error) {
				assert.Equal(t, dstName, name)
				return nil, nil
			},
			ioCopyBuffer: func(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) {
				return 0, io.ErrShortWrite
			},
		}

		err := client.Copy(srcName, dstName)
		assert.Error(t, err)
	})

	t.Run("All successful", func(t *testing.T) {
		fakeFile := &os.File{}
		// Monkey patch the file's Close.
		monkey.PatchInstanceMethod(reflect.TypeOf(fakeFile), "Close",
			func(f *os.File) error {
				return nil
			})

		srcName := uuid.New().String()
		dstName := uuid.New().String()
		client := Client{
			osOpen: func(name string) (file *os.File, e error) {
				assert.Equal(t, srcName, name)
				return fakeFile, nil
			},
			osCreate: func(name string) (file *os.File, e error) {
				assert.Equal(t, dstName, name)
				return fakeFile, nil
			},
			ioCopyBuffer: func(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) {
				return 0, nil
			},
		}

		err := client.Copy(srcName, dstName)
		assert.NoError(t, err)
	})
}

func TestClient_Move(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		srcName := uuid.New().String()
		dstName := uuid.New().String()

		client := Client{
			osRename: func(oldpath, newpath string) error {
				assert.Equal(t, srcName, oldpath)
				assert.Equal(t, dstName, newpath)
				return &os.LinkError{
					Op:  "rename",
					Old: oldpath,
					New: newpath,
					Err: syscall.EISDIR,
				}
			},
		}

		err := client.Move(srcName, dstName)
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		srcName := uuid.New().String()
		dstName := uuid.New().String()

		client := Client{
			osRename: func(oldpath, newpath string) error {
				assert.Equal(t, srcName, oldpath)
				assert.Equal(t, dstName, newpath)
				return nil
			},
		}

		err := client.Move(srcName, dstName)
		assert.NoError(t, err)
	})
}

func TestClient_Reach(t *testing.T) {
	client := Client{}

	assert.Panics(t, func() {
		_, _ = client.Reach(uuid.New().String())
	})
}

func TestClient_CreateDir(t *testing.T) {
	paths := make([]string, 10)
	for k := range paths {
		paths[k] = uuid.New().String()
	}
	tests := []struct {
		name string
		err  error
	}{
		{
			"error",
			&os.PathError{Op: "mkdir", Path: paths[0], Err: syscall.ENOTDIR},
		},
		{
			"success",
			nil,
		},
	}

	for k, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			client := Client{
				osMkdirAll: func(path string, perm os.FileMode) error {
					assert.Equal(t, paths[k], path)
					assert.Equal(t, os.FileMode(0755), perm)
					return v.err
				},
			}

			err := client.CreateDir(paths[k])
			assert.Equal(t, v.err == nil, err == nil)
		})
	}
}

func TestClient_ListDir(t *testing.T) {
	paths := make([]string, 100)
	for k := range paths {
		paths[k] = uuid.New().String()
	}

	tests := []struct {
		name  string
		pairs []*types.Pair
		fi    []os.FileInfo
		items []*types.Object
		err   error
	}{
		{
			"success file",
			nil,
			[]os.FileInfo{
				fileInfo{
					name:    "test_file",
					size:    1234,
					mode:    0644,
					modTime: time.Unix(1, 0),
				},
			},
			[]*types.Object{
				{
					Name: filepath.Join(paths[0], "test_file"),
					Type: types.ObjectTypeFile,
					Metadata: types.Metadata{
						types.Size:      int64(1234),
						types.UpdatedAt: time.Unix(1, 0),
					},
				},
			},
			nil,
		},
		{
			"success file recursively",
			[]*types.Pair{
				types.WithRecursive(true),
			},
			[]os.FileInfo{
				fileInfo{
					name:    "test_file",
					size:    1234,
					mode:    0644,
					modTime: time.Unix(1, 0),
				},
			},
			[]*types.Object{
				{
					Name: filepath.Join(paths[1], "test_file"),
					Type: types.ObjectTypeFile,
					Metadata: types.Metadata{
						types.Size:      int64(1234),
						types.UpdatedAt: time.Unix(1, 0),
					},
				},
			},
			nil,
		},
		{
			"success dir",
			nil,
			[]os.FileInfo{
				fileInfo{
					name:    "test_dir",
					size:    0,
					mode:    os.ModeDir | 0755,
					modTime: time.Unix(1, 0),
				},
			},
			[]*types.Object{
				{
					Name: filepath.Join(paths[2], "test_dir"),
					Type: types.ObjectTypeDir,
					Metadata: types.Metadata{
						types.Size:      int64(0),
						types.UpdatedAt: time.Unix(1, 0),
					},
				},
			},
			nil,
		},
		{
			"success dir recursively",
			[]*types.Pair{
				types.WithRecursive(true),
			},
			[]os.FileInfo{
				fileInfo{
					name:    "test_dir",
					size:    0,
					mode:    os.ModeDir | 0755,
					modTime: time.Unix(1, 0),
				},
			},
			nil,
			nil,
		},
		{
			"os error",
			[]*types.Pair{
				types.WithRecursive(true),
			},
			nil,
			nil,
			&os.PathError{Op: "readdir", Path: "", Err: syscall.ENOTDIR},
		},
		{
			"os error",
			nil,
			nil,
			nil,
			&os.PathError{Op: "readdir", Path: "", Err: syscall.ENOTDIR},
		},
		{
			"done",
			nil,
			nil,
			nil,
			nil,
		},
	}

	for k, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			called := false
			client := Client{
				ioutilReadDir: func(dirname string) (infos []os.FileInfo, e error) {
					if called {
						return nil, nil
					}
					called = true
					assert.Equal(t, paths[k], dirname)
					return v.fi, v.err
				},
			}

			x := client.ListDir(paths[k], v.pairs...)
			for _, expectItem := range v.items {
				item, err := x.Next()
				if v.err != nil {
					assert.Error(t, err)
					assert.True(t, errors.Is(err, v.err))
				}
				assert.NotNil(t, item)
				assert.EqualValues(t, expectItem, item)
			}
			if len(v.items) == 0 {
				item, err := x.Next()
				if v.err != nil {
					assert.Error(t, err)
				} else {
					assert.True(t, errors.Is(err, iterator.ErrDone))
				}
				assert.Nil(t, item)
			}
		})
	}
}

func TestClient_Read(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		pairs   []*types.Pair
		isNil   bool
		openErr error
		seekErr error
	}{
		{
			"success",
			"test_success",
			nil,
			false,
			nil,
			nil,
		},
		{
			"error",
			"test_error",
			nil,
			true,
			&os.PathError{Op: "readdir", Path: "", Err: syscall.ENOTDIR},
			nil,
		},
		{
			"stdin",
			"-",
			nil,
			false,
			nil,
			nil,
		},
		{
			"stdin with size",
			"-",
			[]*types.Pair{
				types.WithSize(100),
			},
			false,
			nil,
			nil,
		},
		{
			"success with size",
			"test_success",
			[]*types.Pair{
				types.WithSize(100),
			},
			false,
			nil,
			nil,
		},
		{
			"success with offset",
			"test_success",
			[]*types.Pair{
				types.WithOffset(10),
			},
			false,
			nil,
			nil,
		},
		{
			"error with offset",
			"test_success",
			[]*types.Pair{
				types.WithOffset(10),
			},
			true,
			nil,
			io.ErrUnexpectedEOF,
		},
		{
			"success with and size offset",
			"test_success",
			[]*types.Pair{
				types.WithSize(100),
				types.WithOffset(10),
			},
			false,
			nil,
			nil,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			fakeFile := &os.File{}
			monkey.PatchInstanceMethod(reflect.TypeOf(fakeFile), "Seek", func(f *os.File, offset int64, whence int) (ret int64, err error) {
				t.Logf("Seek has been called.")
				assert.Equal(t, int64(10), offset)
				assert.Equal(t, 0, whence)
				return 0, v.seekErr
			})

			client := Client{
				osOpen: func(name string) (file *os.File, e error) {
					assert.Equal(t, v.path, name)
					return fakeFile, v.openErr
				},
			}

			o, err := client.Read(v.path, v.pairs...)
			assert.Equal(t, v.openErr == nil && v.seekErr == nil, err == nil)
			assert.Equal(t, v.isNil, o == nil)
		})
	}
}

func TestClient_Write(t *testing.T) {
	paths := make([]string, 10)
	for k := range paths {
		paths[k] = uuid.New().String()
	}

	tests := []struct {
		name         string
		osCreate     func(name string) (*os.File, error)
		ioCopyN      func(dst io.Writer, src io.Reader, n int64) (written int64, err error)
		ioCopyBuffer func(dst io.Writer, src io.Reader, buf []byte) (written int64, err error)
		hasErr       bool
	}{
		{
			"failed os create",
			func(name string) (file *os.File, e error) {
				assert.Equal(t, paths[0], name)
				return nil, &os.PathError{
					Op:   "open",
					Path: "",
					Err:  os.ErrNotExist,
				}
			},
			nil,
			nil,
			true,
		},
		{
			"failed io copyn",
			func(name string) (file *os.File, e error) {
				assert.Equal(t, paths[1], name)
				return &os.File{}, nil
			},
			func(dst io.Writer, src io.Reader, n int64) (written int64, err error) {
				return 0, io.EOF
			},
			nil,
			true,
		},
		{
			"failed io copy buffer",
			nil,
			nil,
			func(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) {
				return 0, io.EOF
			},
			true,
		},
		{
			"success with size",
			func(name string) (file *os.File, e error) {
				assert.Equal(t, paths[3], name)
				return &os.File{}, nil
			},
			func(dst io.Writer, src io.Reader, n int64) (written int64, err error) {
				assert.Equal(t, int64(1234), n)
				return 0, nil
			},
			nil,
			false,
		},
		{
			"success with stdout",
			nil,
			func(dst io.Writer, src io.Reader, n int64) (written int64, err error) {
				assert.Equal(t, int64(1234), n)
				return 0, nil
			},
			nil,
			false,
		},
	}

	for k, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			client := Client{
				osCreate:     v.osCreate,
				ioCopyN:      v.ioCopyN,
				ioCopyBuffer: v.ioCopyBuffer,
			}

			var pairs []*types.Pair
			if v.ioCopyN != nil {
				pairs = append(pairs, types.WithSize(1234))
			}

			var err error
			if v.osCreate == nil {
				err = client.Write("-", nil, pairs...)
			} else {
				err = client.Write(paths[k], nil, pairs...)
			}
			assert.Equal(t, v.hasErr, err != nil)
		})
	}
}

func TestGetAbsPath(t *testing.T) {
	paths := make([]string, 10)
	for k := range paths {
		paths[k] = uuid.New().String()
	}

	cases := []struct {
		base string
	}{
		{paths[0]},
		{paths[1]},
	}

	for _, tt := range cases {
		client := Client{
			base: tt.base,
		}

		absPath := client.getAbsPath(paths[9])
		assert.True(t, strings.HasPrefix(absPath, tt.base))
		assert.Equal(t, tt.base, filepath.Dir(absPath))
	}
}
