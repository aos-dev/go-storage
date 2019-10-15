package qingstor

import (
	"fmt"
	"io"
	"strings"

	"github.com/pengsrc/go-shared/convert"
	iface "github.com/yunify/qingstor-sdk-go/v3/interface"
	"github.com/yunify/qingstor-sdk-go/v3/service"

	"github.com/Xuanwo/storage/pkg/iterator"
	"github.com/Xuanwo/storage/pkg/segment"
	"github.com/Xuanwo/storage/types"
)

// Client is the qingstor object storage client.
//
//go:generate go run ../../internal/cmd/meta_gen/main.go
//go:generate mockgen -package qingstor -destination mock_test.go github.com/yunify/qingstor-sdk-go/v3/interface Service,Bucket
type Client struct {
	bucket iface.Bucket

	segments map[string]*segment.Segment
}

// newClient will create a new client.
func newClient(bucket iface.Bucket) *Client {
	return &Client{
		bucket:   bucket,
		segments: make(map[string]*segment.Segment),
	}
}

// Metadata implements Storager.Metadata
func (c *Client) Metadata() (m types.Metadata, err error) {
	errorMessage := "qingstor Metadata: %w"

	output, err := c.bucket.GetStatistics()
	if err != nil {
		err = handleQingStorError(err)
		return nil, fmt.Errorf(errorMessage, err)
	}

	m = make(types.Metadata)
	if output.Name != nil {
		m.SetName(*output.Name)
	}
	if output.Location != nil {
		m.SetLocation(*output.Location)
	}
	if output.Size != nil {
		m.SetSize(*output.Size)
	}
	if output.Count != nil {
		m.SetCount(*output.Count)
	}
	return m, nil
}

// Stat implements Storager.Stat
func (c *Client) Stat(path string, pairs ...*types.Pair) (o *types.Object, err error) {
	errorMessage := "qingstor Stat: %w"

	input := &service.HeadObjectInput{}

	output, err := c.bucket.HeadObject(path, input)
	if err != nil {
		err = handleQingStorError(err)
		return nil, fmt.Errorf(errorMessage, err)
	}

	o = &types.Object{
		Name:     path,
		Type:     types.ObjectTypeFile,
		Metadata: make(types.Metadata),
	}
	o.SetType(service.StringValue(output.ContentType))
	o.SetSize(*output.ContentLength)
	o.SetChecksum(service.StringValue(output.ETag))
	o.SetStorageClass(service.StringValue(output.XQSStorageClass))
	return o, nil
}

// Delete implements Storager.Delete
func (c *Client) Delete(path string, pairs ...*types.Pair) (err error) {
	errorMessage := "qingstor Delete: %w"

	// TODO: support delete dir.

	_, err = c.bucket.DeleteObject(path)
	if err != nil {
		err = handleQingStorError(err)
		return fmt.Errorf(errorMessage, err)
	}
	return nil
}

// Copy implements Storager.Copy
func (c *Client) Copy(src, dst string, pairs ...*types.Pair) (err error) {
	errorMessage := "qingstor Copy: %w"

	_, err = c.bucket.PutObject(dst, &service.PutObjectInput{
		XQSCopySource: &src,
	})
	if err != nil {
		err = handleQingStorError(err)
		return fmt.Errorf(errorMessage, err)
	}
	return nil
}

// Move implements Storager.Move
func (c *Client) Move(src, dst string, pairs ...*types.Pair) (err error) {
	errorMessage := "qingstor Move: %w"

	_, err = c.bucket.PutObject(dst, &service.PutObjectInput{
		XQSMoveSource: &src,
	})
	if err != nil {
		err = handleQingStorError(err)
		return fmt.Errorf(errorMessage, err)
	}
	return nil
}

// Reach implements Storager.Reach
func (c *Client) Reach(path string, pairs ...*types.Pair) (url string, err error) {
	errorMessage := "qingstor Reach: %w"

	opt, err := parseStoragePairReach(pairs...)
	if err != nil {
		return "", fmt.Errorf(errorMessage, err)
	}

	// FIXME: sdk should export GetObjectRequest as interface too?
	bucket := c.bucket.(*service.Bucket)

	r, _, err := bucket.GetObjectRequest(path, nil)
	if err != nil {
		err = handleQingStorError(err)
		return "", fmt.Errorf(errorMessage, err)
	}
	if err = r.Build(); err != nil {
		err = handleQingStorError(err)
		return "", fmt.Errorf(errorMessage, err)
	}

	expire := 3600
	if opt.HasExpire {
		expire = opt.Expire
	}
	if err = r.SignQuery(expire); err != nil {
		err = handleQingStorError(err)
		return "", fmt.Errorf(errorMessage, err)
	}
	return r.HTTPRequest.URL.String(), nil
}

// CreateDir implements Storager.CreateDir
func (c *Client) CreateDir(path string, option ...*types.Pair) (err error) {
	panic("not supported")
}

// ListDir implements Storager.ListDir
func (c *Client) ListDir(path string, pairs ...*types.Pair) (it iterator.Iterator) {
	errorMessage := "qingstor ListDir: %w"

	opt, _ := parseStoragePairListDir(pairs...)

	marker := ""
	limit := 200
	delimiter := ""
	if opt.HasDelimiter {
		delimiter = opt.Delimiter
	}

	var output *service.ListObjectsOutput
	var err error

	fn := iterator.NextFunc(func(objects *[]*types.Object) error {
		idx := 0
		buf := make([]*types.Object, limit)

		output, err = c.bucket.ListObjects(&service.ListObjectsInput{
			Limit:     &limit,
			Marker:    &marker,
			Prefix:    &path,
			Delimiter: &delimiter,
		})
		if err != nil {
			err = handleQingStorError(err)
			return fmt.Errorf(errorMessage, err)
		}

		for _, v := range output.CommonPrefixes {
			o := &types.Object{
				Name: *v,
				Type: types.ObjectTypeDir,
			}

			buf[idx] = o
			idx++
		}

		for _, v := range output.Keys {
			o := &types.Object{
				Name:     *v.Key,
				Metadata: make(types.Metadata),
			}

			// If Key end with delimiter or key's MimeType == DirectoryMIMEType,
			// we should treat this key as a Dir Object.
			if (delimiter != "" && strings.HasSuffix(*v.Key, delimiter)) ||
				service.StringValue(v.MimeType) == DirectoryMIMEType {
				o.Type = types.ObjectTypeDir
				o.SetType(service.StringValue(v.MimeType))
			} else {
				o.Type = types.ObjectTypeFile
				o.SetType(service.StringValue(v.MimeType))
				o.SetStorageClass(service.StringValue(v.StorageClass))
				o.SetChecksum(service.StringValue(v.Etag))
				o.SetSize(service.Int64Value(v.Size))
			}

			buf[idx] = o
			idx++
		}

		// Set input objects
		*objects = buf[:idx]

		marker = convert.StringValue(output.NextMarker)
		if marker == "" {
			return iterator.ErrDone
		}
		if output.HasMore != nil && !*output.HasMore {
			return iterator.ErrDone
		}
		if len(output.Keys) == 0 {
			return iterator.ErrDone
		}
		return nil
	})

	it = iterator.NewPrefixBasedIterator(fn)
	return
}

// Read implements Storager.Read
func (c *Client) Read(path string, pairs ...*types.Pair) (r io.ReadCloser, err error) {
	errorMessage := "qingstor ReadFile: %w"

	input := &service.GetObjectInput{}

	output, err := c.bucket.GetObject(path, input)
	if err != nil {
		err = handleQingStorError(err)
		return nil, fmt.Errorf(errorMessage, err)
	}
	return output.Body, nil
}

// WriteFile implements Storager.WriteFile
func (c *Client) WriteFile(path string, size int64, r io.Reader, pairs ...*types.Pair) (err error) {
	errorMessage := "qingstor WriteFile for path %s: %w"

	opts, err := parseStoragePairWriteFile(pairs...)
	if err != nil {
		return fmt.Errorf(errorMessage, path, err)
	}

	input := &service.PutObjectInput{
		ContentLength: &size,
		Body:          r,
	}
	if opts.HasChecksum {
		input.ContentMD5 = &opts.Checksum
	}
	if opts.HasStorageClass {
		input.XQSStorageClass = &opts.StorageClass
	}

	_, err = c.bucket.PutObject(path, input)
	if err != nil {
		err = handleQingStorError(err)
		return fmt.Errorf(errorMessage, path, err)
	}
	return nil
}

// WriteStream implements Storager.WriteStream
func (c *Client) WriteStream(path string, r io.Reader, option ...*types.Pair) (err error) {
	panic("not supported")
}

// InitSegment implements Storager.InitSegment
func (c *Client) InitSegment(path string, pairs ...*types.Pair) (err error) {
	errorMessage := "qingstor InitSegment for path %s: %w"

	if _, ok := c.segments[path]; ok {
		return fmt.Errorf(errorMessage, path, segment.ErrSegmentAlreadyInitiated)
	}

	input := &service.InitiateMultipartUploadInput{}

	output, err := c.bucket.InitiateMultipartUpload(path, input)
	if err != nil {
		err = handleQingStorError(err)
		return fmt.Errorf(errorMessage, path, err)
	}

	c.segments[path] = &segment.Segment{
		ID:    *output.UploadID,
		Parts: make([]*segment.Part, 0),
	}
	return
}

// ReadSegment implements Storager.ReadSegment
func (c *Client) ReadSegment(path string, offset, size int64, option ...*types.Pair) (r io.ReadCloser, err error) {
	panic("implement me")
}

// WriteSegment implements Storager.WriteSegment
func (c *Client) WriteSegment(path string, offset, size int64, r io.Reader, pairs ...*types.Pair) (err error) {
	errorMessage := "qingstor WriteSegment for path %s: %w"

	s, ok := c.segments[path]
	if !ok {
		return fmt.Errorf(errorMessage, path, segment.ErrSegmentNotInitiated)
	}

	p := &segment.Part{
		Offset: offset,
		Size:   size,
	}

	partNumber, err := s.GetPartIndex(p)
	if err != nil {
		return fmt.Errorf(errorMessage, path, err)
	}

	_, err = c.bucket.UploadMultipart(path, &service.UploadMultipartInput{
		PartNumber:    &partNumber,
		UploadID:      &s.ID,
		ContentLength: &size,
		Body:          r,
	})
	if err != nil {
		err = handleQingStorError(err)
		return fmt.Errorf(errorMessage, path, err)
	}

	err = s.InsertPart(p)
	if err != nil {
		return fmt.Errorf(errorMessage, path, err)
	}
	return
}

// CompleteSegment implements Storager.CompleteSegment
func (c *Client) CompleteSegment(path string, pairs ...*types.Pair) (err error) {
	errorMessage := "qingstor CompleteSegment for path %s: %w"

	s, ok := c.segments[path]
	if !ok {
		return fmt.Errorf(errorMessage, path, segment.ErrSegmentNotInitiated)
	}

	err = s.ValidateParts()
	if err != nil {
		return fmt.Errorf(errorMessage, path, err)
	}

	objectParts := make([]*service.ObjectPartType, len(s.Parts))
	for k, v := range s.Parts {
		partNumber := k
		objectParts[k] = &service.ObjectPartType{
			PartNumber: &partNumber,
			Size:       &v.Size,
		}
	}

	_, err = c.bucket.CompleteMultipartUpload(path, &service.CompleteMultipartUploadInput{
		UploadID:    &s.ID,
		ObjectParts: objectParts,
	})
	if err != nil {
		err = handleQingStorError(err)
		return fmt.Errorf(errorMessage, path, err)
	}

	delete(c.segments, path)
	return
}

// AbortSegment implements Storager.AbortSegment
func (c *Client) AbortSegment(path string, pairs ...*types.Pair) (err error) {
	errorMessage := "qingstor AbortSegment for path %s: %w"

	s, ok := c.segments[path]
	if !ok {
		return fmt.Errorf(errorMessage, path, segment.ErrSegmentNotInitiated)
	}

	_, err = c.bucket.AbortMultipartUpload(path, &service.AbortMultipartUploadInput{
		UploadID: &s.ID,
	})
	if err != nil {
		err = handleQingStorError(err)
		return fmt.Errorf(errorMessage, path, err)
	}

	delete(c.segments, path)
	return
}
