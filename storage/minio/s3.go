package storage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/ravielze/oculi/context"
	"github.com/ravielze/oculi/storage"
)

func New(endpoint, username, password string, useSSL bool) (storage.S3, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(username, password, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}
	return &impl{
		cl:      client,
		buckets: make(map[string]storage.Bucket),
	}, nil
}

func (i *impl) ListBuckets(ctx context.Context) ([]storage.BucketInfo, error) {
	var result []storage.BucketInfo
	buckets, err := i.cl.ListBuckets(ctx.Context())
	if err != nil {
		return nil, err
	}
	for _, x := range buckets {
		result = append(result,
			storage.BucketInfo{
				Name:      x.Name,
				CreatedAt: x.CreationDate,
			},
		)
	}
	return result, nil
}

func (i *impl) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	found, err := i.cl.BucketExists(ctx.Context(), bucketName)
	if err != nil {
		return false, err
	}
	return found, nil
}

func (i *impl) GetBucket(ctx context.Context, bucketName string) (storage.Bucket, error) {
	i.mu.Lock()
	defer i.mu.Unlock()
	if i.buckets[bucketName] != nil {
		return i.buckets[bucketName], nil
	}

	exists, err := i.BucketExists(ctx, bucketName)
	if err != nil {
		return nil, err
	}

	if !exists {
		err := i.cl.MakeBucket(ctx.Context(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}
	i.buckets[bucketName] = &bucket{
		cl:     i.cl,
		name:   bucketName,
		parent: i,
	}
	return i.buckets[bucketName], nil
}
