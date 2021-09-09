package storage

import (
	"bytes"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/ravielze/oculi/request"
)

type (
	S3 interface {
		// List buckets available
		ListBuckets() ([]BucketInfo, error)

		// Check if bucket with specified name is exist
		BucketExists(bucketName string) (bool, error)

		// Create bucket if not exists, return the same bucket if exists.
		// The bucket can be retrieved again from GetBucket method.
		InitBucket(bucketName string) (Bucket, error)

		// Get bucket, return nil if not exists, always nil if bucketName never initiated.
		GetBucket(bucketName string) Bucket
	}

	Bucket interface {
		// Perform a deletion on bucket, needs to empty the bucket.
		Delete(ctx request.ReqContext) error

		// File put object
		FPutObject(ctx request.ReqContext, objectName, filePath string) error
		// io.Reader put object
		PutObject(ctx request.ReqContext, objectName string, content io.ReadSeeker) error

		// Get object into bytes
		GetObject(ctx request.ReqContext, objectName string) (*bytes.Buffer, error)
		// Get object into file
		FGetObject(ctx request.ReqContext, objectName, filePath string) error

		// Remove object by name
		RemoveObject(ctx request.ReqContext, objectName string) error
		// Remove multiple objects by filter
		RemoveFilteredObjects(ctx request.ReqContext, filter func(o minio.ObjectInfo) bool, limit int) (int, error)

		// List available objects in the bucket
		ListObjects(ctx request.ReqContext, prefix string) ([]ObjectInfo, error)
		// List and filter available objects in the bucket
		FilteredListObjects(ctx request.ReqContext, filter func(o minio.ObjectInfo) bool) ([]ObjectInfo, error)

		// Get object info
		StatObject(ctx request.ReqContext, objectName string) (ObjectInfo, error)
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
