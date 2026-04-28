package storage

import (
	"context"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var Client *minio.Client

func Init(endpoint, accessKey, secretKey string, useSSL bool) error {
	var err error
	Client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	return err
}

func EnsureBucket(ctx context.Context, bucket, region string) error {
	exists, err := Client.BucketExists(ctx, bucket)
	if err != nil {
		return err
	}
	if !exists {
		return Client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{Region: region})
	}
	return nil
}

func Upload(ctx context.Context, bucket, objectName, filePath, contentType string) error {
	_, err := Client.FPutObject(ctx, bucket, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

func UploadReader(ctx context.Context, bucket, objectName string, reader io.Reader, size int64, contentType string) error {
	_, err := Client.PutObject(ctx, bucket, objectName, reader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

func GetURL(ctx context.Context, bucket, objectName string, expiry time.Duration) (string, error) {
	u, err := Client.PresignedGetObject(ctx, bucket, objectName, expiry, nil)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func Delete(ctx context.Context, bucket, objectName string) error {
	return Client.RemoveObject(ctx, bucket, objectName, minio.RemoveObjectOptions{})
}

func Exists(ctx context.Context, bucket, objectName string) (bool, error) {
	_, err := Client.StatObject(ctx, bucket, objectName, minio.StatObjectOptions{})
	if err != nil {
		resp := minio.ToErrorResponse(err)
		if resp.Code == "NoSuchKey" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
