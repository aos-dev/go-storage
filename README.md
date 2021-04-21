# storage

[![Build Status](https://github.com/aos-dev/go-storage/workflows/Unit%20Test/badge.svg?branch=master)](https://github.com/aos-dev/go-storage/actions?query=workflow%3A%22Unit+Test%22)
[![Go dev](https://pkg.go.dev/badge/github.com/aos-dev/go-storage/v3)](https://pkg.go.dev/github.com/aos-dev/go-storage/v3)
[![License](https://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/aos-dev/go-storage/v3/blob/master/LICENSE)
[![go storage dev](https://img.shields.io/matrix/go-storage:aos.dev.svg?server_fqdn=chat.aos.dev&label=%23go-storage%3Aaos.dev&logo=matrix)](https://matrix.to/#/#go-storage:aos.dev)

An application-oriented unified storage layer for Golang.

## Goal

- Production ready
- High performance
- Vendor agnostic

## Features

### Widely services support

- [azblob](https://github.com/aos-dev/go-service-azblob/): [Azure Blob storage](https://docs.microsoft.com/en-us/azure/storage/blobs/)
- [cos](https://github.com/aos-dev/go-service-cos/): [Tencent Cloud Object Storage](https://cloud.tencent.com/product/cos)
- [dropbox](https://github.com/aos-dev/go-service-dropbox/): [Dropbox](https://www.dropbox.com)
- [fs](https://github.com/aos-dev/go-service-fs/): Local file system
- [gcs](https://github.com/aos-dev/go-service-gcs/): [Google Cloud Storage](https://cloud.google.com/storage/)
- [kodo](https://github.com/aos-dev/go-service-kodo/): [qiniu kodo](https://www.qiniu.com/products/kodo)
- [oss](https://github.com/aos-dev/go-service-oss/): [Aliyun Object Storage](https://www.aliyun.com/product/oss)
- [qingstor](https://github.com/aos-dev/go-service-qingstor/): [QingStor Object Storage](https://www.qingcloud.com/products/qingstor/)
- [s3](https://github.com/aos-dev/go-service-s3/): [Amazon S3](https://aws.amazon.com/s3/)
- [uss](https://github.com/aos-dev/go-service-uss/): [UPYUN Storage Service](https://www.upyun.com/products/file-storage)

### Servicer operation support

- List: list all Storager in service
- Get: get a Storager via name
- Create: create a Storager
- Delete: delete a Storager

### Storager operation support

Basic operations

- Metadata: get storager metadata
- Read: read file content
- Write: write content into file
- Stat: get file's metadata
- Delete: delete a file or directory
- List: list file in prefix or dir styles

Extended operations

- Copy: copy a file inside storager
- Move: move a file inside storager
- Reach: generate a public accessible url

Multi object modes support

- Multipart: allow doing multipart uploads
- Append: allow appending to an object
- Block: allow combining an object with block ids.
- Page: allow doing random writes

### Object metadata support

Common metadata

- `id`: unique key in service
- `name`: relative path towards service's work dir
- `type`: object type cloud be `file`, `dir`, `link` or `unknown`

Optional metadata

- `size`: object's content size.
- `updated-at`: object's last updated time.
- `content-md5`: md5 digest as defined in [rfc2616](https://tools.ietf.org/html/rfc2616#section-14.15)
- `content-type`: media type as defined in [rfc2616](https://tools.ietf.org/html/rfc2616#section-14.17)
- `etag`: entity tag as defined in [rfc2616](https://tools.ietf.org/html/rfc2616#section-14.19)
- `storage-class`: object's storage class as defined
  in [storage proposal](https://github.com/aos-dev/specs/tree/master/rfcs/8-normalize-metadata-storage-class.md)

## Quick Start

```go
package main

import (
	"bytes"
	"log"

	"github.com/aos-dev/go-storage/v3/pairs"
	"github.com/aos-dev/go-service-fs/v2"
)

func main() {
	// Init a service.
	store, err := fs.NewStorager(pairs.WithWorkDir("/tmp"))
	if err != nil {
		log.Fatalf("service init failed: %v", err)
	}

	content := []byte("Hello, world!")
	length := int64(len(content))
	r := bytes.NewReader(content)

	_, err = store.Write("hello", r, length)
	if err != nil {
		log.Fatalf("write failed: %v", err)
	}

	var buf bytes.Buffer

	_, err = store.Read("hello", &buf)
	if err != nil {
		log.Fatalf("storager read: %v", err)
	}

	log.Printf("%s", buf.String())
}
```

## Sponsor

<a href="https://vercel.com?utm_source=aos-dev&utm_campaign=oss">
    <img src="./docs/images/vercel.svg">
</a>
