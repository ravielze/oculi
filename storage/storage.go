package storage

import (
	"bytes"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/ravielze/oculi/context"
)

type (
	S3 interface {
		// List buckets available
		ListBuckets(ctx *context.Context) ([]BucketInfo, error)

		// Check if bucket with specified name is exist
		BucketExists(ctx *context.Context, bucketName string) (bool, error)

		// Create bucket if not exists, return the same bucket if exists.
		// The bucket can be retrieved again from GetBucket method.
		InitBucket(ctx *context.Context, bucketName string) (Bucket, error)

		// Get bucket, return nil if not exists, always nil if bucketName never initiated.
		GetBucket(ctx *context.Context, bucketName string) Bucket
	}

	Bucket interface {
		// Perform a deletion on bucket, needs to empty the bucket.
		Delete(ctx *context.Context) error

		// File put object
		FPutObject(ctx *context.Context, objectName, filePath string) error
		// io.Reader put object
		PutObject(ctx *context.Context, objectName string, content io.ReadSeeker) error

		// Get object into bytes
		GetObject(ctx *context.Context, objectName string) (*bytes.Buffer, error)
		// Get object into file
		FGetObject(ctx *context.Context, objectName, filePath string) error

		// Remove object by name
		RemoveObject(ctx *context.Context, objectName string) error
		// Remove multiple objects by filter
		RemoveFilteredObjects(ctx *context.Context, filter func(o minio.ObjectInfo) bool, limit int) (int, error)

		// List available objects in the bucket
		ListObjects(ctx *context.Context, prefix string) ([]ObjectInfo, error)
		// List and filter available objects in the bucket
		FilteredListObjects(ctx *context.Context, filter func(o minio.ObjectInfo) bool) ([]ObjectInfo, error)

		// Get object info
		StatObject(ctx *context.Context, objectName string) (ObjectInfo, error)
	}

	BucketInfo struct {
		Name      string
		CreatedAt time.Time
	}

	ObjectInfo struct {
		Key          string    `json:"name"`
		LastModified time.Time `json:"lastModified"`
		Size         int64     `json:"size"`
		ContentType  string    `json:"contentType"`
	}
)
