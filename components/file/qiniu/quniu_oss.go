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
	"errors"
	"io"
	"sync"

	"mosn.io/layotto/components/file"
	"mosn.io/layotto/components/pkg/actuators"
)

const (
	endpointKey   = "endpoint"
	fileSizeKey   = "filesize"
	componentName = "file-qiniu"
)

var (
	ErrClientNotExist     = errors.New("specific client not exist")
	ErrNotSpecifyEndPoint = errors.New("not specify endpoint in metadata")
	once                  sync.Once
	readinessIndicator    *actuators.HealthIndicator
	livenessIndicator     *actuators.HealthIndicator
)

func init() {
	readinessIndicator = actuators.NewHealthIndicator()
	livenessIndicator = actuators.NewHealthIndicator()
}

type QiniuOSS struct {
	metadata map[string]*OssMetadata
	client   map[string]*QiniuOSSClient
}

type OssMetadata struct {
	Endpoint        string `json:"endpoint"`        // bucket url
	AccessKeyID     string `json:"accessKeyID"`     // SecretID
	AccessKeySecret string `json:"accessKeySecret"` // SecretKey
	Bucket          string `json:"bucket"`
	Private         bool   `json:"private"`
	UseHTTPS        bool   `json:"useHTTPS"`
	UseCdnDomains   bool   `json:"useCdnDomains"`
}

func NewQiniuOSS() file.File { _ = "STUB: not implemented"; return *new(file.File) }

func (q *QiniuOSS) Init(ctx context.Context, metadata *file.FileConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *OssMetadata) checkMetadata() bool { _ = "STUB: not implemented"; return false }

func (q *QiniuOSS) selectClient(meta map[string]string) (*QiniuOSSClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *QiniuOSS) Put(ctx context.Context, st *file.PutFileStu) error {
	_ = "STUB: not implemented"
	return nil
}

func (q *QiniuOSS) Get(ctx context.Context, st *file.GetFileStu) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (q *QiniuOSS) List(ctx context.Context, st *file.ListRequest) (*file.ListResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *QiniuOSS) Del(ctx context.Context, st *file.DelRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (q *QiniuOSS) Stat(ctx context.Context, st *file.FileMetaRequest) (*file.FileMetaResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
