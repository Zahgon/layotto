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

package hdfs

import (
	"context"
	"errors"
	"io"
	"sync"

	"mosn.io/layotto/components/file"
	"mosn.io/layotto/components/pkg/actuators"

	"go.beyondstorage.io/v5/types"
)

const (
	endpointKey   = "endpoint"
	fileSize      = "filesize"
	componentName = "file-hdfs"
)

var (
	ErrMissingEndPoint    error = errors.New("missing endpoint info in metadata")
	ErrClientNotExist     error = errors.New("specific client not exist")
	ErrInvalidConfig      error = errors.New("invalid hdfs config")
	ErrNotSpecifyEndpoint error = errors.New("other error happend in metadata")
	ErrHdfsListFail       error = errors.New("hdfs list opt failed")
	ErrInitFailed         error = errors.New("hdfs client init failed")
	once                  sync.Once
	readinessIndicator    *actuators.HealthIndicator
	livenessIndicator     *actuators.HealthIndicator
)

func init() {
	readinessIndicator = actuators.NewHealthIndicator()
	livenessIndicator = actuators.NewHealthIndicator()
}

type hdfs struct {
	client map[string]types.Storager
	meta   map[string]*HdfsMetaData
}

type HdfsMetaData struct {
	EndPoint string `json:"endpoint"`
}

func NewHdfs() file.File { _ = "STUB: not implemented"; return *new(file.File) }

func (h *hdfs) Init(ctx context.Context, config *file.FileConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *hdfs) Put(ctx context.Context, stu *file.PutFileStu) error {
	_ = "STUB: not implemented"
	return nil
}

//It depends on OS HDFS XML ???

func (h *hdfs) Get(ctx context.Context, stu *file.GetFileStu) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (h *hdfs) List(ctx context.Context, request *file.ListRequest) (*file.ListResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *hdfs) Del(ctx context.Context, request *file.DelRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *hdfs) Stat(ctx context.Context, request *file.FileMetaRequest) (*file.FileMetaResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *hdfs) selectClient(meta map[string]string) (client types.Storager, err error) {
	_ = "STUB: not implemented"
	return *new(types.Storager), nil
}

//endpoint not invaild

//May be not use?
//Because BeyondStorage implemented storage type cannot be assigned a value

func (h *hdfs) createHdfsClient(meta *HdfsMetaData) (types.Storager, error) {
	_ = "STUB: not implemented"
	return *new(types.Storager), nil
}

// ishdfsMetaValid check if the metadata is valid
func (hm *HdfsMetaData) isHdfsMetaValid() bool { _ = "STUB: not implemented"; return false }
