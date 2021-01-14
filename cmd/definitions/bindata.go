// Code generated by go-bindata. DO NOT EDIT.
// sources:
// tmpl/function.tmpl (332B)
// tmpl/info.tmpl (1.802kB)
// tmpl/object.tmpl (1.909kB)
// tmpl/operation.tmpl (1.109kB)
// tmpl/pair.tmpl (588B)
// tmpl/service.tmpl (8.909kB)

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

var _cmdDefinitionsTmplFunctionTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xcf\x4a\xc4\x30\x10\xc6\xef\x7d\x8a\x8f\x25\x87\xee\xb2\xe6\x01\x04\x4f\x45\x41\x10\x59\xd4\xbb\x0c\xd9\x74\x0d\x36\x93\x92\x4c\xeb\x42\xcc\xbb\x4b\x5a\x75\xf1\xe0\x29\x21\xbf\xc9\xf7\x67\x72\xbe\x42\x24\x3e\x59\xa8\xd7\x3d\xd4\x8c\xeb\x1b\x28\x7d\x37\xb1\x49\x28\xa5\xa9\xd8\xf5\xe0\x20\x50\xb3\xbe\xf7\xe3\x60\xbd\x65\xb1\xc7\x1f\xa8\x7a\x7e\x5f\xfe\xcc\xfa\x91\xbc\xc5\x27\x24\x74\xe4\xed\x50\x07\xfa\x89\x0d\xda\x84\x5d\xce\x50\x17\x7e\xa0\x64\xa8\x0e\x6c\x91\x73\x15\x28\xa5\x35\x72\x86\x09\x2c\xf6\x2c\xba\x5b\xcf\x7d\xa5\xb3\x3e\x50\x24\x9f\xf4\x4b\x74\xfe\x81\x92\xe8\x67\x89\x8e\x4f\xb7\x7c\x4c\x1f\x4e\xde\xba\xe0\x3d\x95\x82\x30\x0a\x76\x23\xb9\xf8\x8f\x55\x7d\xae\x51\xff\xfa\xb7\x8b\xc3\x93\x4d\xd3\x20\xe9\x5b\x79\x89\xd5\x00\xc0\x48\xec\x4c\xbb\xa9\xe5\xdd\xa5\xf9\x66\xdb\xac\xd5\x2d\xff\x6e\x61\xbd\x7e\x05\x00\x00\xff\xff\x3b\x1f\xad\x5f\x4c\x01\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd1, 0x53, 0x97, 0x24, 0x27, 0xb1, 0x48, 0xc, 0x9a, 0xff, 0xbf, 0xcc, 0x50, 0x35, 0x3e, 0xf2, 0xfd, 0x26, 0xfd, 0x33, 0xa9, 0x5, 0x3f, 0xa2, 0x4c, 0x9b, 0x7, 0xfa, 0x1a, 0x11, 0xb7, 0xe0}}
	return a, nil
}

var _cmdDefinitionsTmplInfoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x94\x51\x6f\xd3\x3e\x14\xc5\xdf\xfd\x29\xce\xbf\xaa\xfe\x6a\x50\xd7\x6c\x12\xe2\x01\xd6\xa7\x6d\xa0\x09\x6d\x43\xda\xc4\x03\x08\x21\x37\xb9\xa9\x4c\x63\x3b\xb2\x9d\x68\x25\xcd\x77\x47\x76\xb2\xb6\x29\x5d\xa1\x88\x17\xde\xec\xd8\xbe\xe7\xe7\xe3\x73\x13\xc7\xb8\xd0\x29\x61\x4e\x8a\x0c\x77\x94\x62\xb6\xc4\x5c\xaf\xe7\x10\xca\x91\x51\x3c\x8f\x13\x99\xbe\xc1\xe5\x1d\x6e\xef\x1e\x70\x75\x79\xfd\x30\x61\x05\x4f\x16\x7c\x4e\x70\xcb\x82\x2c\x63\x42\x16\xda\x38\x8c\x18\x00\x0c\x32\xe9\x06\x2c\x62\xac\xae\x4f\x60\xb8\x9a\x13\x86\x8b\x31\x86\x42\x65\xda\xe2\xf5\x14\x93\x6b\x3f\xba\xe1\x05\x9a\x86\xd5\x35\x86\x96\x4c\x25\x12\xba\xe5\x92\xfc\xfa\x70\x81\x15\x9c\xbe\xe0\x92\x72\xbf\x85\xc5\x31\xde\x0a\xca\x53\x08\x95\xd2\x23\x84\x42\x5d\x6f\x1f\x6a\x1a\xcc\x84\x63\x89\x56\xd6\x43\xec\xe8\x56\xa1\x66\xab\xde\x34\x01\x71\xf7\xf8\xb5\xaf\xeb\x49\xaa\x49\x80\xf0\xf2\x1f\xb8\x4d\xb8\xd7\xc7\x14\x67\xe7\xe7\x7e\x75\xd1\x02\x9f\x80\x54\xea\x87\x11\x63\xde\x00\xec\xde\xa1\x7f\xdc\x3a\x53\x26\x0e\x75\xa7\xbc\x66\xfb\xfa\x1c\x9b\xc7\x78\x58\x16\x6d\xad\xa6\xd9\xfa\xb2\xd9\xb3\x66\x08\xf3\x38\xf6\x06\xa0\xb4\x94\x82\x5b\x70\x3f\x93\xbc\x40\xa6\x0d\xf4\xec\x1b\x25\x0e\x15\xcf\x4b\x1a\xe3\x14\x92\xb8\xb2\x50\xda\xc1\x92\x1b\xe3\xac\xfb\x60\xc9\x85\x52\xa1\x8e\x50\xee\xd5\xcb\x30\x95\x90\xbc\xf8\x6c\x9d\x11\x6a\xfe\x25\x04\x22\xe3\x09\xd5\x0d\xeb\x94\x0f\x7b\xed\x57\x45\xe6\xe9\xaf\x1e\x43\x42\x9a\x86\x65\xa5\x4a\x30\x92\x78\x71\xd0\xb5\x08\xef\xc8\xb5\x17\xbf\x14\xb6\xc8\xf9\xb2\x73\x63\x14\xf5\xfd\xe8\x7c\x35\xe4\x4a\xa3\x20\x27\x3f\xd9\xe7\x49\x7f\x57\xf3\xfe\x19\xcd\xaa\xaf\x19\xfd\xa2\x50\xc7\xb4\x07\x06\x53\x54\x3d\x5e\xd6\x05\x2a\xb7\x81\xf5\x2f\xb8\x33\xea\xa1\x8e\x31\xd3\x3a\x8f\x3a\x22\x91\x41\x4e\xfc\x0b\xff\x7f\x64\x0b\xfc\x37\xc5\x69\x57\xe3\xb0\xdb\x63\x38\x53\x52\xd8\xb8\x89\x6a\x1b\x81\x4f\x64\xf4\x47\x9f\xc2\xa7\x14\x6f\x55\x6a\xeb\x6c\xef\x18\x23\xe3\xb9\xa5\x4d\xdc\x3b\x87\xf6\x1e\x0c\x97\x5d\xe1\xfb\xa1\xf3\x6d\xbb\x1c\x61\xf1\x4d\x69\xdd\x71\x21\xfc\x63\x7b\xa7\x7d\x7b\x0b\xae\x44\x32\xca\xa4\x9b\xdc\x17\x46\x28\x97\x8d\x06\xfb\x58\xdf\xd3\x8c\xcf\x36\x3f\x88\xa7\x88\x89\x75\x77\x0f\xa2\x68\xeb\x25\xfe\xb5\x16\x69\xad\x5c\x4d\x8f\xf3\x72\x6f\x77\xb5\x6f\xbf\x13\x85\xcd\xf0\x47\x00\x00\x00\xff\xff\xc9\xe7\x0b\xe2\x0a\x07\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd, 0xa2, 0xdc, 0xf4, 0x79, 0x9b, 0xb4, 0x94, 0xfb, 0x8a, 0xe3, 0xb3, 0xca, 0x75, 0x5b, 0x39, 0x8, 0x1a, 0xc4, 0x69, 0x74, 0x63, 0x42, 0x82, 0xec, 0xa5, 0xc6, 0xc, 0xe0, 0x35, 0x8a, 0x7d}}
	return a, nil
}

var _cmdDefinitionsTmplObjectTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\x4d\x6f\xe3\x36\x10\x3d\x9b\xbf\xe2\xd5\x08\x0a\xa9\x70\xa5\xa4\x2d\x7a\x68\xe3\x53\x92\x16\x39\x24\x29\x90\xa0\x87\x16\x7b\xa0\xa5\x91\xc3\x8d\x44\x0a\xe4\xc8\x6b\xaf\xa2\xff\xbe\x20\x25\x7f\x25\xde\xc0\xc1\x5e\xf6\x64\x0e\x39\xf3\xe6\xe3\xbd\xb1\xd2\x14\x17\x26\x27\xcc\x49\x93\x95\x4c\x39\x66\x2b\xcc\xcd\xc6\x86\xd2\x4c\x56\xcb\x32\xcd\xaa\xfc\x4f\x5c\xde\xe1\xf6\xee\x01\x57\x97\xd7\x0f\x89\xa8\x65\xf6\x24\xe7\x04\x5e\xd5\xe4\x84\x50\x55\x6d\x2c\x23\x12\x00\x30\x2e\x2a\x1e\xf7\x27\x56\x15\x0d\x47\xb7\xd2\xd9\x58\xc4\x42\xa4\x29\xfe\x52\x54\xe6\x50\x3a\xa7\x25\x94\x86\x99\x7d\xa4\x8c\x31\x53\x2c\x32\xa3\x9d\xc7\x69\xdb\x9f\x61\xa5\x9e\x13\x4e\x9e\x26\x38\x59\xe0\x8f\x29\x92\xbb\xe0\x77\x43\x2c\xd1\x75\x01\xb5\x8f\xbc\xf6\x40\x6d\x8b\x93\x45\x72\x2b\x2b\xc2\x33\xd8\xfc\x23\x5d\x26\x4b\x74\x1d\x1a\xa5\xf9\xf7\xdf\x30\xc5\xd9\xf9\xb9\x77\x7a\xf2\xc1\x1e\x9f\x74\xee\x8f\xb1\x10\xbe\x0b\xf4\xe8\x70\x6c\x9b\x8c\xd1\x1e\x5b\x82\x77\x53\x85\xcf\x7d\x61\xaa\x8a\x34\xaf\x1f\x90\xa6\xe8\x8b\x7a\xf1\xb0\x93\xbb\x37\xbd\xcf\xc3\xaa\xa6\x50\x7c\xd7\xed\xdc\xbc\x28\x55\x8c\xd2\x14\x59\xa9\x3c\x98\x72\xe0\x47\xda\x58\x1a\x9f\x1e\x55\xf6\xb8\xee\x42\x39\xc8\x52\x2d\x28\x11\xa3\xc1\xe3\x9e\x8d\x95\x73\xb2\x01\xa3\x82\x63\x63\xc9\x85\x1f\xcf\xa3\xa5\x32\x28\xa0\x22\x96\xb9\x64\x99\x84\xd2\xbc\x85\x4a\xd6\xff\x3b\xb6\x4a\xcf\x3f\x04\x41\x14\x32\xa3\xb6\x13\x62\x68\x71\xa6\x18\x8d\xa3\x1c\xd2\x41\x7a\xab\x92\x35\x0a\x63\xd7\xac\x2e\x64\xd9\xd0\x04\xa7\xa8\x48\x6a\x07\x6d\x18\x8e\x78\x82\xb3\xe1\xc2\x11\x07\xa8\x80\x13\xa8\x12\xa3\xdc\x68\x0a\xc6\xaf\xbf\x88\x51\xe5\x5f\xbd\x7a\x92\x9b\x86\x69\x29\x3a\x21\x8e\xe1\x66\xcb\xcb\xd5\x32\x88\xb3\xeb\x44\xd1\xe8\x0c\x91\xc1\x4f\xbd\x6b\x8c\xbf\x89\xfb\x61\x5f\x2a\x57\x97\x72\x35\x30\x10\xc5\xfb\x1c\xa0\x0d\x25\x5a\xe2\xc6\x6a\x98\xe4\x15\x65\xbe\xa8\x57\xe0\xf7\x5f\x01\x5f\xec\x83\xc7\xeb\x88\x21\xcb\x01\x78\x4c\xb1\xd8\xab\x40\x0c\xb2\x28\x5d\xc8\xfe\x9e\xc6\xa2\xbd\xe4\x13\xcc\x8c\x29\xe3\x4d\x6a\xc7\x92\xa3\xb8\x27\x57\x15\x30\x89\xa7\xe5\xc7\xa3\xf6\xec\x87\x29\x4e\x07\x9c\xb7\x87\x35\x01\xdb\x86\x82\xe3\x20\xa3\x2d\x57\xff\x91\x35\xff\x7a\xc5\x6c\xb6\x68\x0b\xd5\x03\xed\x7a\x4c\x50\xc8\xd2\xd1\x76\xab\x86\x79\x1c\x0c\x0c\x1d\x3f\xe3\xf3\x5b\xf1\xfd\x9a\x1d\x1a\xe8\x4d\xe3\xf8\x7d\x6a\xf9\xa6\x61\x4e\xf7\x87\x59\x4b\xad\xb2\xa8\xa8\x38\xb9\xaf\xad\xd2\x5c\x44\xe3\x61\xbd\x76\x10\xba\xce\xaf\xfd\xb0\x60\xe3\x38\x1e\x26\xfc\x7d\x28\xb7\x6f\xfe\x79\x7a\x4c\xf7\x07\xa5\xde\x53\xb3\xfb\x67\xf8\xaa\xec\xac\x34\x9a\xa2\xe5\xce\x4d\xbb\xa1\xf6\x98\x8f\xc9\xc1\xf2\x97\x87\x86\xf6\x42\x30\x43\x78\xf8\xbb\x0c\x11\xfe\x24\x76\xfa\x0e\x97\xfe\xfb\xd6\x7d\x09\x00\x00\xff\xff\x83\x24\x9a\x18\x75\x07\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x77, 0x12, 0x6b, 0xe1, 0xaa, 0xe9, 0xfb, 0xb4, 0x95, 0xa8, 0x1d, 0x54, 0x53, 0x2d, 0xc1, 0x96, 0x91, 0xbb, 0x1a, 0x6d, 0x4c, 0x12, 0x34, 0xf7, 0xe0, 0xb, 0x36, 0x6e, 0x95, 0xe4, 0x20, 0x9f}}
	return a, nil
}

var _cmdDefinitionsTmplOperationTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x52\x4f\xeb\xd4\x30\x10\xbd\xe7\x53\x0c\xa5\x87\x16\xd6\xf6\x2e\x78\x10\xff\x80\x17\x2d\x7a\xf0\x28\xd9\xec\xb4\x0e\xb6\x4d\x4c\xb2\xcb\x2e\x31\xdf\x5d\x92\xa6\xa5\x5b\xba\xac\xc8\xef\x96\xcc\x9b\xf7\xe6\x65\xf2\x14\x17\xbf\x78\x87\x60\x6f\x0a\x0d\x63\x34\x28\xa9\x2d\x14\x0c\x00\x20\x13\x72\xb4\x78\xb5\xd9\x74\x23\x99\xb1\x92\x31\xe7\x5e\x81\xe6\x63\x87\x90\xff\x38\x40\x4e\xf0\xfa\x0d\x54\x9f\x46\x8b\xba\xe5\x02\x0d\x78\xcf\x9c\x83\x9c\xaa\xf7\x68\x84\x26\x65\x49\x8e\xa1\x18\x26\x40\x42\xc8\xa8\x9e\xdf\x3e\xf3\x01\xc1\x7b\xa0\x99\x0c\x2e\x4e\x0a\x13\xa8\x05\xa9\xa1\xc0\xdf\xa1\x3f\x36\x66\x06\xf5\x85\x04\xea\xac\xdc\xd4\xad\xd4\xbc\x0b\x75\xef\x23\xff\x9b\xd5\x34\x76\x45\x09\x26\x1e\x16\x4d\x1c\x4f\xc1\xc8\x72\x5f\xbd\x02\x87\x23\x9e\xc2\x4b\x72\xaa\x3e\xc4\x73\xd2\x9a\x7a\x53\xc3\xc6\xf8\x03\xe1\xb5\xae\x54\x49\xf4\x8b\x32\x33\xa3\xae\xa3\xa2\x54\x93\xff\x3f\x60\x65\xc3\x8d\xe0\x7d\xd8\x45\x42\x36\xbb\x9b\x5d\xec\x72\x8a\x84\x7c\x94\x7a\xe0\xb6\xe1\x9a\x0f\x61\x56\x09\xf7\xc0\x57\x34\xe7\xde\x9a\xef\x64\x7f\x36\xd3\xa7\xdf\xad\x2f\x0b\x94\xa7\xfe\x02\xfb\xdd\x94\x8a\xff\xb1\xba\xa2\x17\xc2\x5e\x21\x05\xac\x4a\xb5\xc3\x0b\xbf\xc4\xb9\xf9\x6b\x3c\x5b\xff\x53\x8c\x62\xc3\x49\x37\xb2\x27\x71\x0b\x39\x39\x0b\x9b\xd2\xf7\xb6\xef\xe1\x28\x65\xff\x8f\x41\xa7\x16\x46\xdc\x09\x69\x80\xeb\x1a\x14\x27\x6d\xa0\x95\x7a\x9d\x72\xb7\xf4\xdf\x6d\x67\x3b\x71\x2f\x3c\x8f\x93\xb3\x6b\xfa\x12\x15\xa4\xaa\x9a\xe8\xe3\x99\x48\x00\x2e\x1b\xdd\x45\x36\x6d\xef\xe9\xd1\xb3\xbf\x01\x00\x00\xff\xff\xce\x1e\x6e\xf5\x55\x04\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd3, 0xc, 0xf7, 0x3d, 0x63, 0x22, 0xe7, 0x8f, 0x9e, 0xdb, 0xf8, 0x26, 0x3c, 0x34, 0xd7, 0xca, 0xba, 0x28, 0x89, 0xe5, 0xc3, 0x64, 0xb3, 0x43, 0x96, 0x9d, 0xdd, 0x5, 0x25, 0x36, 0xfa, 0x6c}}
	return a, nil
}

var _cmdDefinitionsTmplPairTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\x31\xeb\xdb\x30\x10\xc5\x77\x7d\x8a\x87\xf1\x90\x40\x62\x41\xc7\x94\x4e\x4d\x87\x52\x48\x32\x84\x76\x2c\x8a\x7c\x55\x44\x64\x49\xc8\x67\xb7\xc6\xf5\x77\x2f\xb2\x93\x40\xf9\x2f\xff\x68\xd2\xfd\xee\xde\xd3\x3b\x49\x89\xcf\xa1\x26\x18\xf2\x94\x14\x53\x8d\xcb\x00\x13\x9e\x35\xac\x67\x4a\x5e\x39\xa9\x9b\xfa\x23\xf6\x47\x1c\x8e\x67\x7c\xd9\x7f\x3d\x57\x22\x2a\x7d\x53\x86\x10\x95\x4d\xad\x10\xb6\x89\x21\x31\x56\x02\x00\x0a\x1d\x3c\xd3\x1f\x2e\xc4\x52\x1a\xcb\xd7\xee\x52\xe9\xd0\x48\x15\xda\x6d\x4d\xbd\x34\x61\xdb\x72\x48\xca\x90\xec\x3f\xc8\x78\x33\x92\x7c\x1d\x83\xf5\x5c\xbc\xa0\xd1\x89\x6a\xf2\x6c\x95\x7b\x45\x75\x65\x8e\xda\x59\x7a\xbc\x55\xbd\x43\xc7\x43\xa4\xb6\x10\x6b\x21\xc6\x71\x8b\xa4\xbc\x21\x94\x3f\x37\x28\x7b\xec\x3e\xa1\x3a\xe5\x5f\xc0\x34\xcd\xdd\x32\x7a\xd5\x50\xe6\x65\x5f\x1d\xf2\xf5\x2f\x38\x9c\x54\xab\x95\xcb\x33\x52\xe2\x87\xe5\xeb\x38\x3e\x26\xa7\x09\xbf\xad\x73\x50\x31\xba\x01\x99\xdf\x75\xd3\x84\x5e\xb9\x8e\xc0\x01\xc7\xc8\x36\xf8\x56\x2c\xed\x3d\xb5\x3a\xd9\x19\x65\xcb\x5f\x9d\xd7\x6f\x4c\x57\xfd\xdd\xeb\x3c\xc4\x5c\xaf\x91\x63\x62\x9c\x97\x4e\xc4\x5d\xf2\x33\x59\x40\x3e\xdf\x68\xd8\xa1\xf8\x2f\x40\xb1\x79\x76\xbf\xe7\x28\x3b\xf4\x0b\x99\xc4\xb2\x2d\xf9\x3a\x27\xf8\x17\x00\x00\xff\xff\x08\xa2\x3c\x2b\x4c\x02\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd8, 0x5e, 0xa0, 0x39, 0x1, 0x59, 0x67, 0x13, 0x40, 0x8c, 0x71, 0x6, 0x57, 0xa0, 0x25, 0x10, 0x6d, 0x91, 0xa4, 0xef, 0xc6, 0xa0, 0x70, 0x5d, 0x3b, 0x5f, 0x54, 0xed, 0x8e, 0xe, 0xaa, 0x99}}
	return a, nil
}

var _cmdDefinitionsTmplServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6d\x6f\xdb\xb8\x1d\x7f\xef\x4f\xf1\xaf\xe0\x0e\x52\xa1\x48\xdb\x5e\x66\xf0\x8b\x2d\xed\xb2\xa2\x43\x12\x2c\xd9\x15\xb8\xeb\x21\x60\x24\xca\xe6\x45\x22\x55\x92\x56\xe2\x73\xf4\xdd\x0f\x7c\x90\x44\xc9\x92\x63\x17\xc5\xf5\x0e\x4d\x5e\x04\x32\xf9\x7f\x7e\xe2\x4f\x54\x1c\xc3\x19\x4b\x31\x2c\x31\xc5\x1c\x49\x9c\xc2\xdd\x06\x96\xac\xfd\x0d\x15\x41\x40\xa8\xc4\x9c\xa2\x3c\x4e\x8a\x34\x16\x98\x57\x24\xc1\xff\x80\xb7\x97\x70\x71\x79\x03\xef\xde\xbe\xbf\x89\x66\x25\x4a\xee\xd1\x12\xc3\x76\x0b\xd1\x05\x2a\x30\xd4\xf5\x6c\x46\x8a\x92\x71\x09\xfe\x0c\x00\xc0\x4b\x18\x95\xf8\x51\x7a\xe6\x17\x61\xde\xcc\x3c\x2d\x89\x5c\xad\xef\xa2\x84\x15\x31\x62\xe2\x24\xc5\x55\xbc\x64\x27\x42\x32\x8e\x96\x38\xae\xfe\x1e\x97\xf7\xcb\x18\xd3\xb4\x64\x84\x36\xdc\x07\xf1\x24\x1c\xa7\x98\x4a\x82\xf2\x63\xb8\x56\x52\x96\x49\x4e\xf0\xe1\xba\x6c\x44\x84\xa1\x8f\x0e\xe0\x90\x9b\x52\x91\x07\xb3\x59\x85\x38\xdc\x42\x67\x69\x74\xc5\x59\x45\x52\xcc\xed\x4e\xe3\xf7\x70\xfd\xda\x08\x6b\x7e\x36\x26\x44\xd7\xe6\xe1\x1d\xe7\xac\xd9\xeb\x1c\x8a\x2e\x4b\x49\x18\x15\xb3\x59\x1c\xc3\xcd\xa6\xc4\x40\x04\xc8\x15\x06\x65\x0f\x64\x8c\xf7\xd2\x97\x30\x2a\xa4\x21\x5b\x80\xe7\xec\x78\x9a\xdf\x6a\x02\x54\x21\x92\xa3\xbb\x1c\x43\x89\x08\x17\x91\xe5\xf3\x67\xdb\xed\x09\x70\x44\x97\x18\xe6\xb7\x21\xcc\x2b\x38\x5d\x40\x74\xa5\x68\x94\x74\x15\x2a\x45\x41\x32\xa0\x4c\xc2\xbc\x8a\xce\x73\x76\x87\xf2\x6e\x4f\xad\xbd\xc5\x22\xe1\x44\x5b\xdd\xdf\xf8\xf7\x3a\xcf\xad\x39\xd6\xba\x79\x63\xde\xad\xa1\x68\x8d\x6d\x34\x61\x9a\x2a\x19\xce\x63\x30\x7b\xc6\xc8\x29\x03\xd5\xfa\xbc\xa4\x4a\xc3\xe9\xa2\x55\xf6\x04\x92\x5d\x21\x91\x18\x9a\x38\x86\x8f\x44\xae\x94\x31\x86\xb2\xae\xe1\x81\xe4\x39\xa0\xb2\xcc\x37\xd0\x33\x12\x2a\x94\xaf\x31\x48\x06\x4d\x86\x46\xdd\xcf\xd6\x34\xd9\x11\xea\x57\x56\x96\xce\x54\x5d\x07\xa0\xcc\x87\xad\xf6\x5b\xfd\x71\x2c\xd7\x9c\xea\xd5\x6e\xf1\x03\xde\x9c\xee\x86\x32\x6c\xf7\x7f\x50\x06\x9d\x42\x65\x56\xea\x59\x2f\x6e\xce\xe3\x44\x04\xdf\xd3\x8c\x1d\x9c\xe6\x67\x02\xe9\x88\x50\x21\x21\xa2\xcc\xd1\xa6\x29\xd2\xc6\xde\x4e\xce\x62\x82\xca\x31\xba\x61\xb0\x5d\x73\x61\xd5\x97\x9c\x50\x99\x81\xf7\x5a\xbc\x16\x1e\xf8\xf3\x2a\xba\x4e\x58\xe9\x9a\x13\xe8\xd5\x33\x24\xf1\x92\xf1\x4d\x6f\xa3\x6f\x28\xfe\x0c\x2d\xbb\xc7\xee\x7e\xc1\x89\xf4\x86\xe6\xba\xda\x17\xe0\x5d\xf6\xa9\x06\xe6\xc6\x31\x9c\x63\xb9\x5b\x4c\x4b\x2c\x47\x4b\x29\xe3\xac\x80\x02\x4b\x94\x22\x89\x22\x2d\x42\x57\xcf\x40\x88\x5f\xc0\x9b\xa1\x2d\xaa\x84\xfc\x5e\x49\x85\x70\xc7\x58\x1e\x80\x2d\xaa\x2a\x04\x76\xaf\x02\x56\x44\xe7\x58\xfa\xbd\xe6\x3b\xe9\x37\x5f\xa0\x19\x48\x06\xaf\xd8\xbd\xe5\xee\x52\xf9\x23\xe6\x4c\xd7\x99\x1b\x19\x5b\xad\x46\x8c\x4b\x11\x42\x86\x72\x81\xbb\xe0\xe4\x62\x9a\x51\x1b\xfe\x04\xbf\xee\xe3\xef\x82\x6b\xfe\x5b\x01\x55\xd4\x77\x3e\x08\x41\xf2\x35\xb6\x84\x4d\x32\x84\x89\x63\x17\x76\x9d\x0d\x31\x91\x0d\x42\x25\x1b\xcb\x86\x38\x24\x1b\x21\xec\x74\xf8\x08\x95\x8d\xae\x75\xa2\x88\xae\x9f\xc9\x4c\x08\x55\xe0\x38\x3f\xd1\xdd\xa3\xcd\xad\x44\x88\x12\x25\xb8\xd7\xe1\x12\x17\x65\xae\x10\x83\xa7\xce\x81\x5b\x8a\x1f\x3c\x28\xd0\x3d\xbe\xce\xd5\x29\xe1\x8f\x35\x77\xa0\x17\xf1\xc3\xb4\x18\xe1\xca\x68\x24\xe8\x99\x45\x93\x09\xed\x2a\xb0\x7a\x86\x3e\xc7\x3a\xf4\x33\xc5\x19\xa1\x9d\x5e\x47\xf8\xbc\xa4\xca\x71\x42\x53\xfc\x08\x11\xfc\x75\x62\x40\xcd\x55\x76\x5d\xc2\xbf\x69\xd9\xcd\xf6\x30\x8e\x86\xbc\x37\x14\x86\xee\xf7\x5c\x28\x29\xcc\xab\xdd\xe9\x60\x1f\xc7\xbc\x30\x49\xe8\xce\xab\xbe\x13\xcd\x7a\x46\x77\x6d\xb6\x1b\xba\xcb\xe7\x19\x1d\x1b\xcb\xea\x80\x53\x4a\x4c\x01\x43\x5d\xab\x07\xc5\x52\xd7\x0d\xac\x28\x11\x17\x38\x05\x21\xf9\x3a\x91\x33\x0d\x32\x06\x1c\x8a\xa1\xae\x2d\x85\x2d\x61\x1d\x7f\xf8\xe9\x67\x75\x60\xb5\xfd\xf6\x3f\xfc\x79\x4d\x38\x4e\xcd\xee\x58\x4c\xd5\x46\x63\x6e\x4b\x6d\xc3\xf5\x1f\x24\xb4\x52\x44\xf8\x88\x2b\x00\x7a\xbe\x75\xe7\xd1\x24\x59\xb7\x6d\x5b\x71\x62\x54\x9b\x53\x1c\xe5\x87\x59\xdb\x52\x7f\x6d\x6b\x0f\xb6\xf7\xbc\x45\xff\x07\x19\xdc\x91\xff\xde\xf1\x6d\xca\x8e\x0b\x7c\x35\x51\x7b\x7a\x10\x6b\x0a\x03\x84\x84\xee\x1e\x3d\x81\xdf\x4c\xd4\xab\x01\x56\xfb\xa4\xfa\xac\x94\x4d\x4d\x06\xe0\x4f\x09\x0a\x01\x2b\xe0\x1d\xb4\xc3\x58\xac\x73\xa9\xa2\xf6\x97\x09\x86\x0e\x90\xe9\xc0\x9f\x82\xd2\x13\xba\x67\x8d\x42\xe5\xb7\xea\x08\x38\x5d\xd8\x74\x68\x53\x3a\x46\xf1\x40\x64\xb2\x82\x2a\xfa\x80\x37\xce\xf2\x78\xcf\x1c\xd9\x37\xea\x2f\x41\x02\x77\x79\x71\xc0\xe2\x69\x4b\xd2\xf9\x1a\x3d\x53\x09\x8b\xee\x2c\x1d\x30\xee\xe7\xaa\x22\x7d\x8c\x9b\xa3\xd9\x2d\x90\xa0\xe7\x96\x53\xd5\xd3\x9d\x78\x64\x37\xfe\xf9\x63\x30\xd6\xdd\xc7\x76\xf8\x11\x51\xd0\x90\xa7\x81\x88\x09\xa3\x15\xe6\xd2\xa9\xf8\x69\x1f\xfd\x69\x0f\x83\x9e\x06\x92\x41\x0b\x27\xbf\x4a\xfc\x0f\xcf\x81\xb2\xaf\xc7\x5a\x1b\x20\xba\x6b\xcc\xc1\xd6\x74\xb0\xb4\x95\xb9\x2f\xa1\xb5\x3b\x1d\x8e\xe8\x65\x05\xc2\x0f\x33\x69\xe7\xd5\x91\x92\x3c\xec\xae\x18\x2e\xf0\x83\x1a\x83\x8d\x7c\x7d\xd5\xe0\x1b\x84\xd1\xab\x8b\x93\xa6\x32\x77\xa6\xb8\x8b\x54\x8d\x49\xa1\xd2\x31\xdb\x8f\xc9\xfe\xd8\x48\xa6\xe5\x78\xc1\x32\x2f\x58\xe6\x9b\x63\x19\x5f\x98\x77\x44\xbd\x19\xbc\x20\x9b\x17\x64\xf3\x82\x6c\xbe\x32\xb2\x11\xd1\x77\x8c\x6d\xb6\x5b\x92\x01\x35\x97\x12\x9e\xfd\x12\xd0\xbb\xe0\x74\xdc\x11\x91\xd2\x79\xc5\x72\x92\x6c\xa2\x7f\xe6\x39\x3c\x3d\xf5\xd7\x9c\x81\x37\xbd\x75\x18\x60\xe9\x3b\x3d\x0d\x5e\xfe\x4f\xc5\xba\x2c\x19\x97\x0d\x7e\xa9\x82\x1d\x19\xbb\xce\x6c\xb7\x2a\x1e\x23\x5e\x7e\x39\xd4\xb3\x62\xf5\xfd\xcd\xb8\xf2\x91\xca\x4f\x71\x86\xd6\xb9\xec\x17\xe7\x41\x29\x39\x32\x1d\x63\x65\xf8\xe5\x41\xad\x67\x07\x05\x33\x61\x54\x12\x3a\x28\xc5\x61\x8c\xbe\x4b\x14\xec\x5c\x6d\x7e\xeb\xdb\xc9\xc6\x94\x23\x6e\x28\x27\xdd\xd9\xef\x8d\xbb\x37\x06\xf0\xc1\x42\xad\x67\x30\x3e\x98\x13\xc3\x2d\x6e\xfd\x38\xf6\x89\x31\x8e\x1b\xf2\x9b\x15\x11\xd0\xd8\x69\xe0\x58\xc2\xb1\x8a\x00\x02\xfb\x3d\x1b\xee\x36\x4d\x43\x3a\x37\xfb\x06\x80\xcd\x4b\x5a\xd7\x81\xa3\xd4\xd7\x57\x9e\xd1\x15\xe2\xa8\x10\xd1\xb5\xe4\x84\x2e\xed\xf7\x16\x53\xad\xaa\x0c\xdc\x8d\xae\x0a\x13\xf9\x68\x5f\xa8\x95\xd6\xe8\x5f\x28\xb9\x5f\x72\xb6\xa6\xa9\x1f\x0c\x2b\x55\x38\x3d\xfc\x91\xc8\xd5\x99\xe1\xf1\x13\xf9\x18\x42\xcf\x82\x33\x94\xe7\x98\x77\x25\xba\x1b\x26\x87\x7f\x4f\xc4\xf6\xf9\x3c\xb0\xa0\xf5\xc0\xae\x0d\x2c\x3a\x2a\x26\x29\xce\x30\xd7\xca\xfd\x60\x3b\x98\x17\xba\x13\xe5\xaa\xa9\x09\x2b\xfe\x0a\xc9\x95\x71\x7a\x38\x1a\xed\x77\x29\x44\x53\xf0\xf1\x67\xcb\xec\x79\x81\xfd\xe5\x0c\xd5\x60\xec\xa0\x33\xa3\x44\xae\x60\x01\x5e\xf8\xc9\xfb\xe4\xed\xcc\xde\x91\x41\xae\xfe\x30\xe7\xa0\xd0\x44\xc6\x78\x81\xa4\x99\x23\x9e\x71\x5d\xd5\xb1\xfe\x54\xa3\x68\x5a\x05\x75\x0d\x5d\xc6\x6b\x27\xfb\x15\xe2\x0a\x42\x4f\xbf\x1a\x34\x84\xac\x94\x61\xab\x77\xef\xcb\x81\x86\x55\x9d\x06\x92\x69\xb6\x57\x0b\x35\xa8\x06\xa7\x83\xa9\x3d\x67\x40\x8f\x56\x65\xbf\x3f\xcf\x50\x81\x35\x3a\xda\xad\xcc\x1b\x4e\x8a\xff\x22\x21\x6d\x89\xbe\xa3\xa9\x7a\x1d\x58\x9d\xb1\xa2\x40\x75\xad\x5c\x68\x6a\xb6\x8b\xeb\x6f\x01\x00\x00\xff\xff\xe2\x22\x7c\x5a\xcd\x22\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2a, 0x33, 0xd7, 0xab, 0x92, 0x24, 0x5, 0xbb, 0xce, 0x6b, 0x4c, 0x18, 0x82, 0xd9, 0x28, 0x65, 0xe2, 0x9e, 0xef, 0x24, 0x29, 0x5a, 0xf1, 0x77, 0xd5, 0x2a, 0x7, 0xcb, 0x67, 0xec, 0x97, 0x52}}
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
				"function.tmpl":  {cmdDefinitionsTmplFunctionTmpl, map[string]*bintree{}},
				"info.tmpl":      {cmdDefinitionsTmplInfoTmpl, map[string]*bintree{}},
				"object.tmpl":    {cmdDefinitionsTmplObjectTmpl, map[string]*bintree{}},
				"operation.tmpl": {cmdDefinitionsTmplOperationTmpl, map[string]*bintree{}},
				"pair.tmpl":      {cmdDefinitionsTmplPairTmpl, map[string]*bintree{}},
				"service.tmpl":   {cmdDefinitionsTmplServiceTmpl, map[string]*bintree{}},
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
