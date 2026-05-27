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

package tencentcloud

import (
	"context"
	"io"
	"net/url"
	"sync"

	"github.com/pkg/errors"
	"github.com/tencentyun/cos-go-sdk-v5"

	"mosn.io/layotto/components/file"
	"mosn.io/layotto/components/pkg/actuators"
)

const (
	endpointKey    = "endpoint"
	aclKey         = "ACL"
	contentTypeKey = "content-type"
	componentName  = "file-tencentcloud"
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

type TencentCloudOSS struct {
	metadata map[string]*OssMetadata
	client   map[string]*cos.Client
}

type OssMetadata struct {
	Endpoint        string `json:"endpoint"`        // bucket url https://console.cloud.tencent.com/cos/bucket
	AccessKeyID     string `json:"accessKeyID"`     // SecretID https://console.cloud.tencent.com/cam/capi
	AccessKeySecret string `json:"accessKeySecret"` // SecretKey
	Timeout         int    `json:"timeout"`         // timeout in milliseconds
	bucketUrl       *url.URL
}

func NewTencentCloudOSS() file.File { _ = "STUB: not implemented"; return *new(file.File) }

// Init does metadata parsing and connection creation
func (t *TencentCloudOSS) Init(ctx context.Context, metadata *file.FileConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TencentCloudOSS) checkMetadata(m *OssMetadata) bool {
	_ = "STUB: not implemented"
	return false
}

//100s

func (t *TencentCloudOSS) getClient(metadata *OssMetadata) (*cos.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//set timeout

func (t *TencentCloudOSS) Put(ctx context.Context, st *file.PutFileStu) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TencentCloudOSS) Get(ctx context.Context, st *file.GetFileStu) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (t *TencentCloudOSS) List(ctx context.Context, st *file.ListRequest) (*file.ListResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TencentCloudOSS) Del(ctx context.Context, st *file.DelRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TencentCloudOSS) Stat(ctx context.Context, st *file.FileMetaRequest) (*file.FileMetaResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TencentCloudOSS) selectClient(meta map[string]string) (*cos.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TencentCloudOSS) checkFileName(fileName string) error {
	_ = "STUB: not implemented"
	return nil
}
