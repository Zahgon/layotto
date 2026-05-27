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

package aliyun

import (
	"context"
	"io"
	"sync"

	"mosn.io/layotto/components/pkg/actuators"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"mosn.io/layotto/components/file"
	"mosn.io/layotto/components/pkg/utils"
)

const (
	storageTypeKey = "storageType"
	componentName  = "file-aliyun"
)

var (
	once               sync.Once
	readinessIndicator *actuators.HealthIndicator
	livenessIndicator  *actuators.HealthIndicator
)

func init() {
	readinessIndicator = actuators.NewHealthIndicator()
	livenessIndicator = actuators.NewHealthIndicator()
}

// AliyunFile is a binding for an AliCloud OSS storage bucketKey
type AliyunFile struct {
	client *oss.Client
}

func NewAliyunFile() file.File { _ = "STUB: not implemented"; return *new(file.File) }

// Init does metadata parsing and connection creation
func (s *AliyunFile) Init(ctx context.Context, metadata *file.FileConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *AliyunFile) Put(ctx context.Context, st *file.PutFileStu) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *AliyunFile) Get(ctx context.Context, st *file.GetFileStu) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (s *AliyunFile) List(ctx context.Context, request *file.ListRequest) (*file.ListResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//last object is marker

func (s *AliyunFile) Del(ctx context.Context, request *file.DelRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *AliyunFile) Stat(ctx context.Context, request *file.FileMetaRequest) (*file.FileMetaResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *AliyunFile) checkMetadata(m *utils.OssMetadata) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *AliyunFile) getBucket(fileName string, metaData map[string]string) (*oss.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get oss bucket

func (s *AliyunFile) getClient() (*oss.Client, error) { _ = "STUB: not implemented"; return nil, nil }
