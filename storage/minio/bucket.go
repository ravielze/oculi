package storage

import (
	"bytes"
	"io"

	"github.com/gabriel-vasile/mimetype"
	"github.com/minio/minio-go/v7"
	"github.com/ravielze/oculi/context"
	errorUtil "github.com/ravielze/oculi/errors"
	"github.com/ravielze/oculi/storage"
)

func (b *bucket) Delete(ctx context.Context) error {
	b.parent.mu.Lock()
	defer b.parent.mu.Unlock()
	err := b.cl.RemoveBucket(ctx.Context(), b.name)
	if err != nil {
		return err
	}
	b.parent.buckets[b.name] = nil
	return nil
}

func (b *bucket) FPutObject(ctx context.Context, objectName, filePath string) error {
	mtype, err := mimetype.DetectFile(filePath)
	if err != nil {
		return err
	}
	_, err = b.cl.FPutObject(ctx.Context(),
		b.name,
		objectName,
		filePath,
		minio.PutObjectOptions{
			ContentType: mtype.String(),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (b *bucket) PutObject(ctx context.Context, objectName string, content io.Reader) error {
	mtype, err := mimetype.DetectReader(content)
	if err != nil {
		return err
	}
	_, err = b.cl.PutObject(ctx.Context(),
		b.name,
		objectName,
		content,
		-1,
		minio.PutObjectOptions{
			ContentType: mtype.String(),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (b *bucket) GetObject(ctx context.Context, objectName string) (*bytes.Buffer, error) {
	o, err := b.cl.GetObject(ctx.Context(), b.name, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	result := &bytes.Buffer{}
	_, err = io.Copy(result, o)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (b *bucket) FGetObject(ctx context.Context, objectName, filePath string) error {
	return b.cl.FGetObject(
		ctx.Context(), b.name,
		objectName, filePath,
		minio.GetObjectOptions{},
	)
}

func (b *bucket) RemoveObject(ctx context.Context, objectName string) error {
	return b.cl.RemoveObject(ctx.Context(), b.name, objectName, minio.RemoveObjectOptions{})
}

func (b *bucket) RemoveFilteredObjects(ctx context.Context, filter func(o minio.ObjectInfo) bool, limit int) (int, error) {
	objectInfo, err := b.listObjects(ctx, "", filter)
	if err != nil {
		return 0, err
	}

	ch := make(chan minio.ObjectInfo)
	count := 0
	go func() {
		defer close(ch)
		for _, obj := range objectInfo {
			ch <- obj
			count++
		}
	}()
	errCh := b.cl.RemoveObjects(ctx.Context(), b.name, ch, minio.RemoveObjectsOptions{})
	var errors []string
	for x := range errCh {
		if x.Err != nil {
			errors = append(errors, x.Err.Error())
		}
	}
	if len(errors) != 0 {
		return count - len(errors), errorUtil.NewDetailedErrors("failed remove objects", errors)
	}
	return count, nil
}

func (b *bucket) listObjects(ctx context.Context, prefix string, filter func(o minio.ObjectInfo) bool) ([]minio.ObjectInfo, error) {
	ch := b.cl.ListObjects(
		ctx.Context(), b.name,
		minio.ListObjectsOptions{
			Prefix: prefix,
		},
	)
	var result []minio.ObjectInfo
	for x := range ch {
		if x.Err != nil {
			return nil, x.Err
		}
		if filter == nil || filter(x) {
			result = append(result, x)
		}
	}
	return result, nil
}

func (b *bucket) ListObjects(ctx context.Context, prefix string) ([]storage.ObjectInfo, error) {
	objectInfo, err := b.listObjects(ctx, prefix, nil)
	if err != nil {
		return nil, err
	}
	var result []storage.ObjectInfo
	for _, x := range objectInfo {
		result = append(result,
			storage.ObjectInfo{
				Key:          x.Key,
				LastModified: x.LastModified,
				Size:         x.Size,
				ContentType:  x.ContentType,
			},
		)
	}
	return result, nil
}

func (b *bucket) FilteredListObjects(ctx context.Context, filter func(o minio.ObjectInfo) bool) ([]storage.ObjectInfo, error) {
	objectInfo, err := b.listObjects(ctx, "", filter)
	if err != nil {
		return nil, err
	}
	var result []storage.ObjectInfo
	for _, x := range objectInfo {
		result = append(result,
			storage.ObjectInfo{
				Key:          x.Key,
				LastModified: x.LastModified,
				Size:         x.Size,
				ContentType:  x.ContentType,
			},
		)
	}
	return result, nil
}

func (b *bucket) StatObject(ctx context.Context, objectName string) (storage.ObjectInfo, error) {
	info, err := b.cl.StatObject(ctx.Context(), b.name, objectName, minio.StatObjectOptions{})
	if err != nil {
		return storage.ObjectInfo{}, err
	}
	return storage.ObjectInfo{
		Key:          info.Key,
		LastModified: info.LastModified,
		Size:         info.Size,
		ContentType:  info.ContentType,
	}, nil
}
