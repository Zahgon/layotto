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

package aws

import (
	"context"
	"io"
	"sync"

	"mosn.io/layotto/components/pkg/actuators"
	"mosn.io/layotto/components/pkg/utils"

	"github.com/aws/aws-sdk-go-v2/service/s3"

	"mosn.io/layotto/components/file"
)

const (
	defaultCredentialsSource = "provider"
	componentName            = "file-aws"
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

// AwsOss is a binding for aws oss storage.
type AwsOss struct {
	client *s3.Client
}

func NewAwsFile() file.File { _ = "STUB: not implemented"; return *new(file.File) }

// Init instance by config.
func (a *AwsOss) Init(ctx context.Context, config *file.FileConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// isAwsMetaValid check if the metadata valid.
func (a *AwsOss) isAwsMetaValid(v *utils.OssMetadata) bool { _ = "STUB: not implemented"; return false }

// createOssClient by input meta info.
func (a *AwsOss) createOssClient(meta *utils.OssMetadata) (*s3.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put file to aws oss.
func (a *AwsOss) Put(ctx context.Context, st *file.PutFileStu) error {
	_ = "STUB: not implemented"
	//var bodySize int64
	return nil
}

func (a *AwsOss) selectClient() (*s3.Client, error) { _ = "STUB: not implemented"; return nil, nil }

// Get object from aws oss.
func (a *AwsOss) Get(ctx context.Context, st *file.GetFileStu) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// List objects from aws oss.
func (a *AwsOss) List(ctx context.Context, st *file.ListRequest) (*file.ListResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Del object in aws oss.
func (a *AwsOss) Del(ctx context.Context, st *file.DelRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AwsOss) Stat(ctx context.Context, st *file.FileMetaRequest) (*file.FileMetaResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
