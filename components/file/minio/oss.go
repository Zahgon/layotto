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

package minio

import (
	"context"
	"errors"
	"io"
	"sync"

	"mosn.io/layotto/components/pkg/actuators"

	"github.com/minio/minio-go/v7"

	"mosn.io/layotto/components/file"
)

const (
	endpointKey   = "endpoint"
	fileSize      = "fileSize"
	componentName = "file-monio"
)

var (
	ErrMissingEndPoint    error = errors.New("missing endpoint info in metadata")
	ErrClientNotExist     error = errors.New("specific client not exist")
	ErrInvalidConfig      error = errors.New("invalid minio oss config")
	ErrNotSpecifyEndPoint error = errors.New("not specify endpoint in metadata")
	once                  sync.Once
	readinessIndicator    *actuators.HealthIndicator
	livenessIndicator     *actuators.HealthIndicator
)

func init() {
	readinessIndicator = actuators.NewHealthIndicator()
	livenessIndicator = actuators.NewHealthIndicator()
}

type MinioOss struct {
	client map[string]*minio.Core
	meta   map[string]*MinioMetaData
}

type MinioMetaData struct {
	Region          string `json:"region"`
	EndPoint        string `json:"endpoint"`
	AccessKeyID     string `json:"accessKeyID"`
	AccessKeySecret string `json:"accessKeySecret"`
	SSL             bool   `json:"SSL"`
}

func NewMinioOss() file.File { _ = "STUB: not implemented"; return *new(file.File) }

func (m *MinioOss) Init(ctx context.Context, config *file.FileConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MinioOss) Put(ctx context.Context, st *file.PutFileStu) error {
	_ = "STUB: not implemented"
	return nil
}

// specify file size from metadata, default unknown size is -1

func (m *MinioOss) Get(ctx context.Context, st *file.GetFileStu) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (m *MinioOss) List(ctx context.Context, st *file.ListRequest) (*file.ListResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MinioOss) Del(ctx context.Context, st *file.DelRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MinioOss) Stat(ctx context.Context, st *file.FileMetaRequest) (*file.FileMetaResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MinioOss) createOssClient(meta *MinioMetaData) (*minio.Core, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MinioOss) selectClient(meta map[string]string) (client *minio.Core, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// isMinioMetaValid check if the metadata is valid
func (mm *MinioMetaData) isMinioMetaValid() bool { _ = "STUB: not implemented"; return false }
