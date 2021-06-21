// Code generated by go generate cmd/definitions; DO NOT EDIT.
package types

import (
	"fmt"
)

// Field index in storageMeta bit
const (
	storageMetaIndexAppendNumberMaximum    = 1 << 0
	storageMetaIndexAppendSizeMaximum      = 1 << 1
	storageMetaIndexAppendTotalSizeMaximum = 1 << 2
	storageMetaIndexCopySizeMaximum        = 1 << 3
	storageMetaIndexFetchSizeMaximum       = 1 << 4
	storageMetaIndexLocation               = 1 << 5
	storageMetaIndexMoveSizeMaximum        = 1 << 6
	storageMetaIndexMultipartNumberMaximum = 1 << 7
	storageMetaIndexMultipartSizeMaximum   = 1 << 8
	storageMetaIndexMultipartSizeMinimum   = 1 << 9
	storageMetaIndexName                   = 1 << 10
	storageMetaIndexWorkDir                = 1 << 11
	storageMetaIndexWriteSizeMaximum       = 1 << 12
)

type StorageMeta struct {
	appendNumberMaximum    int
	appendSizeMaximum      int64
	appendTotalSizeMaximum int64
	copySizeMaximum        int64
	fetchSizeMaximum       int64
	location               string
	moveSizeMaximum        int64
	multipartNumberMaximum int
	multipartSizeMaximum   int64
	multipartSizeMinimum   int64
	Name                   string
	WorkDir                string
	writeSizeMaximum       int64

	// bit used as a bitmap for object value, 0 means not set, 1 means set
	bit uint64
	m   map[string]interface{}
}

func (m *StorageMeta) GetAppendNumberMaximum() (int, bool) {
	if m.bit&storageMetaIndexAppendNumberMaximum != 0 {
		return m.appendNumberMaximum, true
	}
	return 0, false
}

func (m *StorageMeta) MustGetAppendNumberMaximum() int {
	if m.bit&storageMetaIndexAppendNumberMaximum == 0 {
		panic(fmt.Sprintf("storage-meta append-number-maximum is not set"))
	}
	return m.appendNumberMaximum
}

func (m *StorageMeta) SetAppendNumberMaximum(v int) *StorageMeta {
	m.appendNumberMaximum = v
	m.bit |= storageMetaIndexAppendNumberMaximum
	return m
}

func (m *StorageMeta) GetAppendSizeMaximum() (int64, bool) {
	if m.bit&storageMetaIndexAppendSizeMaximum != 0 {
		return m.appendSizeMaximum, true
	}
	return 0, false
}

func (m *StorageMeta) MustGetAppendSizeMaximum() int64 {
	if m.bit&storageMetaIndexAppendSizeMaximum == 0 {
		panic(fmt.Sprintf("storage-meta append-size-maximum is not set"))
	}
	return m.appendSizeMaximum
}

func (m *StorageMeta) SetAppendSizeMaximum(v int64) *StorageMeta {
	m.appendSizeMaximum = v
	m.bit |= storageMetaIndexAppendSizeMaximum
	return m
}

func (m *StorageMeta) GetAppendTotalSizeMaximum() (int64, bool) {
	if m.bit&storageMetaIndexAppendTotalSizeMaximum != 0 {
		return m.appendTotalSizeMaximum, true
	}
	return 0, false
}

func (m *StorageMeta) MustGetAppendTotalSizeMaximum() int64 {
	if m.bit&storageMetaIndexAppendTotalSizeMaximum == 0 {
		panic(fmt.Sprintf("storage-meta append-total-size-maximum is not set"))
	}
	return m.appendTotalSizeMaximum
}

func (m *StorageMeta) SetAppendTotalSizeMaximum(v int64) *StorageMeta {
	m.appendTotalSizeMaximum = v
	m.bit |= storageMetaIndexAppendTotalSizeMaximum
	return m
}

func (m *StorageMeta) GetCopySizeMaximum() (int64, bool) {
	if m.bit&storageMetaIndexCopySizeMaximum != 0 {
		return m.copySizeMaximum, true
	}
	return 0, false
}

func (m *StorageMeta) MustGetCopySizeMaximum() int64 {
	if m.bit&storageMetaIndexCopySizeMaximum == 0 {
		panic(fmt.Sprintf("storage-meta copy-size-maximum is not set"))
	}
	return m.copySizeMaximum
}

func (m *StorageMeta) SetCopySizeMaximum(v int64) *StorageMeta {
	m.copySizeMaximum = v
	m.bit |= storageMetaIndexCopySizeMaximum
	return m
}

func (m *StorageMeta) GetFetchSizeMaximum() (int64, bool) {
	if m.bit&storageMetaIndexFetchSizeMaximum != 0 {
		return m.fetchSizeMaximum, true
	}
	return 0, false
}

func (m *StorageMeta) MustGetFetchSizeMaximum() int64 {
	if m.bit&storageMetaIndexFetchSizeMaximum == 0 {
		panic(fmt.Sprintf("storage-meta fetch-size-maximum is not set"))
	}
	return m.fetchSizeMaximum
}

func (m *StorageMeta) SetFetchSizeMaximum(v int64) *StorageMeta {
	m.fetchSizeMaximum = v
	m.bit |= storageMetaIndexFetchSizeMaximum
	return m
}

func (m *StorageMeta) GetLocation() (string, bool) {
	if m.bit&storageMetaIndexLocation != 0 {
		return m.location, true
	}
	return "", false
}

func (m *StorageMeta) MustGetLocation() string {
	if m.bit&storageMetaIndexLocation == 0 {
		panic(fmt.Sprintf("storage-meta location is not set"))
	}
	return m.location
}

func (m *StorageMeta) SetLocation(v string) *StorageMeta {
	m.location = v
	m.bit |= storageMetaIndexLocation
	return m
}

func (m *StorageMeta) GetMoveSizeMaximum() (int64, bool) {
	if m.bit&storageMetaIndexMoveSizeMaximum != 0 {
		return m.moveSizeMaximum, true
	}
	return 0, false
}

func (m *StorageMeta) MustGetMoveSizeMaximum() int64 {
	if m.bit&storageMetaIndexMoveSizeMaximum == 0 {
		panic(fmt.Sprintf("storage-meta move-size-maximum is not set"))
	}
	return m.moveSizeMaximum
}

func (m *StorageMeta) SetMoveSizeMaximum(v int64) *StorageMeta {
	m.moveSizeMaximum = v
	m.bit |= storageMetaIndexMoveSizeMaximum
	return m
}

func (m *StorageMeta) GetMultipartNumberMaximum() (int, bool) {
	if m.bit&storageMetaIndexMultipartNumberMaximum != 0 {
		return m.multipartNumberMaximum, true
	}
	return 0, false
}

func (m *StorageMeta) MustGetMultipartNumberMaximum() int {
	if m.bit&storageMetaIndexMultipartNumberMaximum == 0 {
		panic(fmt.Sprintf("storage-meta multipart-number-maximum is not set"))
	}
	return m.multipartNumberMaximum
}

func (m *StorageMeta) SetMultipartNumberMaximum(v int) *StorageMeta {
	m.multipartNumberMaximum = v
	m.bit |= storageMetaIndexMultipartNumberMaximum
	return m
}

func (m *StorageMeta) GetMultipartSizeMaximum() (int64, bool) {
	if m.bit&storageMetaIndexMultipartSizeMaximum != 0 {
		return m.multipartSizeMaximum, true
	}
	return 0, false
}

func (m *StorageMeta) MustGetMultipartSizeMaximum() int64 {
	if m.bit&storageMetaIndexMultipartSizeMaximum == 0 {
		panic(fmt.Sprintf("storage-meta multipart-size-maximum is not set"))
	}
	return m.multipartSizeMaximum
}

func (m *StorageMeta) SetMultipartSizeMaximum(v int64) *StorageMeta {
	m.multipartSizeMaximum = v
	m.bit |= storageMetaIndexMultipartSizeMaximum
	return m
}

func (m *StorageMeta) GetMultipartSizeMinimum() (int64, bool) {
	if m.bit&storageMetaIndexMultipartSizeMinimum != 0 {
		return m.multipartSizeMinimum, true
	}
	return 0, false
}

func (m *StorageMeta) MustGetMultipartSizeMinimum() int64 {
	if m.bit&storageMetaIndexMultipartSizeMinimum == 0 {
		panic(fmt.Sprintf("storage-meta multipart-size-minimum is not set"))
	}
	return m.multipartSizeMinimum
}

func (m *StorageMeta) SetMultipartSizeMinimum(v int64) *StorageMeta {
	m.multipartSizeMinimum = v
	m.bit |= storageMetaIndexMultipartSizeMinimum
	return m
}
func (m *StorageMeta) GetName() string {
	return m.Name
}

func (m *StorageMeta) SetName(v string) *StorageMeta {
	m.Name = v
	return m
}
func (m *StorageMeta) GetWorkDir() string {
	return m.WorkDir
}

func (m *StorageMeta) SetWorkDir(v string) *StorageMeta {
	m.WorkDir = v
	return m
}

func (m *StorageMeta) GetWriteSizeMaximum() (int64, bool) {
	if m.bit&storageMetaIndexWriteSizeMaximum != 0 {
		return m.writeSizeMaximum, true
	}
	return 0, false
}

func (m *StorageMeta) MustGetWriteSizeMaximum() int64 {
	if m.bit&storageMetaIndexWriteSizeMaximum == 0 {
		panic(fmt.Sprintf("storage-meta write-size-maximum is not set"))
	}
	return m.writeSizeMaximum
}

func (m *StorageMeta) SetWriteSizeMaximum(v int64) *StorageMeta {
	m.writeSizeMaximum = v
	m.bit |= storageMetaIndexWriteSizeMaximum
	return m
}
