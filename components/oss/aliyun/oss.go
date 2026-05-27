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
	"encoding/json"
	"sync"

	"mosn.io/layotto/components/pkg/actuators"

	l8oss "mosn.io/layotto/components/oss"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
	connectTimeoutSec   = "connectTimeoutSec"
	readWriteTimeoutSec = "readWriteTimeout"
	componentName       = "oss-aliyun"
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

type AliyunOSS struct {
	client    *oss.Client
	basicConf json.RawMessage
}

func NewAliyunOss() l8oss.Oss { _ = "STUB: not implemented"; return *new(l8oss.Oss) }

func (a *AliyunOSS) Init(ctx context.Context, config *l8oss.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AliyunOSS) GetObject(ctx context.Context, req *l8oss.GetObjectInput) (*l8oss.GetObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//user can use SignedUrl to get file without ak、sk

func (a *AliyunOSS) PutObject(ctx context.Context, req *l8oss.PutObjectInput) (*l8oss.PutObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//user can use SignedUrl to put file without ak、sk

func (a *AliyunOSS) DeleteObject(ctx context.Context, req *l8oss.DeleteObjectInput) (*l8oss.DeleteObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) DeleteObjects(ctx context.Context, req *l8oss.DeleteObjectsInput) (*l8oss.DeleteObjectsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) PutObjectTagging(ctx context.Context, req *l8oss.PutObjectTaggingInput) (*l8oss.PutObjectTaggingOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) DeleteObjectTagging(ctx context.Context, req *l8oss.DeleteObjectTaggingInput) (*l8oss.DeleteObjectTaggingOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) GetObjectTagging(ctx context.Context, req *l8oss.GetObjectTaggingInput) (*l8oss.GetObjectTaggingOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) GetObjectCannedAcl(ctx context.Context, req *l8oss.GetObjectCannedAclInput) (*l8oss.GetObjectCannedAclOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) PutObjectCannedAcl(ctx context.Context, req *l8oss.PutObjectCannedAclInput) (*l8oss.PutObjectCannedAclOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) ListObjects(ctx context.Context, req *l8oss.ListObjectsInput) (*l8oss.ListObjectsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) CopyObject(ctx context.Context, req *l8oss.CopyObjectInput) (*l8oss.CopyObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) CreateMultipartUpload(ctx context.Context, req *l8oss.CreateMultipartUploadInput) (*l8oss.CreateMultipartUploadOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) UploadPart(ctx context.Context, req *l8oss.UploadPartInput) (*l8oss.UploadPartOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) UploadPartCopy(ctx context.Context, req *l8oss.UploadPartCopyInput) (*l8oss.UploadPartCopyOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) CompleteMultipartUpload(ctx context.Context, req *l8oss.CompleteMultipartUploadInput) (*l8oss.CompleteMultipartUploadOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) AbortMultipartUpload(ctx context.Context, req *l8oss.AbortMultipartUploadInput) (*l8oss.AbortMultipartUploadOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) ListMultipartUploads(ctx context.Context, req *l8oss.ListMultipartUploadsInput) (*l8oss.ListMultipartUploadsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) RestoreObject(ctx context.Context, req *l8oss.RestoreObjectInput) (*l8oss.RestoreObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) ListObjectVersions(ctx context.Context, req *l8oss.ListObjectVersionsInput) (*l8oss.ListObjectVersionsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) HeadObject(ctx context.Context, req *l8oss.HeadObjectInput) (*l8oss.HeadObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//if key exist,concatenated with commas

func (a *AliyunOSS) IsObjectExist(ctx context.Context, req *l8oss.IsObjectExistInput) (*l8oss.IsObjectExistOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) SignURL(ctx context.Context, req *l8oss.SignURLInput) (*l8oss.SignURLOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateDownloadBandwidthRateLimit update all client rate
func (a *AliyunOSS) UpdateDownloadBandwidthRateLimit(ctx context.Context, req *l8oss.UpdateBandwidthRateLimitInput) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateUploadBandwidthRateLimit update all client rate
func (a *AliyunOSS) UpdateUploadBandwidthRateLimit(ctx context.Context, req *l8oss.UpdateBandwidthRateLimitInput) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AliyunOSS) AppendObject(ctx context.Context, req *l8oss.AppendObjectInput) (*l8oss.AppendObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) ListParts(ctx context.Context, req *l8oss.ListPartsInput) (*l8oss.ListPartsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AliyunOSS) getClient() (*oss.Client, error) { _ = "STUB: not implemented"; return nil, nil }
