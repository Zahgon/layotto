/*
 * Copyright 2021 Layotto Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package qiniu

import (
	"context"
	"io"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiniuOSSClient struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Domain    string
	Private   bool

	mac *qbox.Mac
	fu  FormUploader
	bm  BucketManager
}

type FormUploader interface {
	Put(ctx context.Context, ret interface{}, uptoken, key string, data io.Reader, size int64, extra *storage.PutExtra) (err error)
}

type BucketManager interface {
	Stat(bucket, key string) (storage.FileInfo, error)
	Delete(bucket, key string) (err error)
	ListFiles(bucket, prefix, delimiter, marker string,
		limit int) (entries []storage.ListItem, commonPrefixes []string, nextMarker string, hasNext bool, err error)
}

func newQiniuOSSClient(ak, sk, bucket, domain string, private bool, useHttps, userCdnDomains bool) *QiniuOSSClient {
	_ = "STUB: not implemented"
	return nil
}

func (s *QiniuOSSClient) put(ctx context.Context, fileName string, data io.Reader, dataSize int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *QiniuOSSClient) get(_ context.Context, fileName string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

//1小时有效期

func (s *QiniuOSSClient) stat(_ context.Context, fileName string) (*storage.FileInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *QiniuOSSClient) del(_ context.Context, fileName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *QiniuOSSClient) list(_ context.Context, prefix string, limit int, marker string) (entries []storage.ListItem, commonPrefixes []string, nextMarker string, hasNext bool, err error) {
	_ = "STUB: not implemented"
	return nil, nil, "", false, nil
}

func (s *QiniuOSSClient) checkFileName(fileName string) error {
	_ = "STUB: not implemented"
	return nil
}
