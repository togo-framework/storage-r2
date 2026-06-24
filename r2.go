// Package r2store is a Cloudflare R2 driver for togo storage. R2 is S3-compatible,
// so this uses the AWS S3 SDK pointed at the R2 endpoint. Implements togo.Storage
// and overrides the default filesystem storage when installed.
//
//	togo install togo-framework/storage-r2
//
// Env: R2_ACCOUNT_ID, R2_ACCESS_KEY_ID, R2_SECRET_ACCESS_KEY, R2_BUCKET (all required).
package r2store

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("storage-r2", togo.PriorityService+10, func(k *togo.Kernel) error {
		bucket := os.Getenv("R2_BUCKET")
		account := os.Getenv("R2_ACCOUNT_ID")
		ak := os.Getenv("R2_ACCESS_KEY_ID")
		sk := os.Getenv("R2_SECRET_ACCESS_KEY")
		if bucket == "" || account == "" || ak == "" || sk == "" {
			if k.Log != nil {
				k.Log.Warn("storage-r2: R2_* env not fully set; skipping (using default storage)")
			}
			return nil
		}
		cfg, err := config.LoadDefaultConfig(context.Background(),
			config.WithRegion("auto"),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(ak, sk, "")),
		)
		if err != nil {
			return fmt.Errorf("storage-r2: load config: %w", err)
		}
		endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", account)
		cl := s3.NewFromConfig(cfg, func(o *s3.Options) { o.BaseEndpoint = aws.String(endpoint) })
		k.Storage = &store{cl: cl, bucket: bucket, endpoint: endpoint}
		return nil
	})
}

type store struct {
	cl       *s3.Client
	bucket   string
	endpoint string
}

func (s *store) Put(path string, data []byte) error {
	_, err := s.cl.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(s.bucket), Key: aws.String(key(path)), Body: bytes.NewReader(data),
	})
	return err
}

func (s *store) Get(path string) ([]byte, error) {
	out, err := s.cl.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(s.bucket), Key: aws.String(key(path)),
	})
	if err != nil {
		return nil, err
	}
	defer out.Body.Close()
	return io.ReadAll(out.Body)
}

func (s *store) Delete(path string) error {
	_, err := s.cl.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket), Key: aws.String(key(path)),
	})
	return err
}

func (s *store) Path(path string) string {
	return fmt.Sprintf("%s/%s/%s", s.endpoint, s.bucket, key(path))
}

func key(p string) string { return strings.TrimPrefix(p, "/") }
