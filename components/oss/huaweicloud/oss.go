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

package huaweicloud

import (
	"context"
	"sync"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"

	"mosn.io/layotto/components/oss"
	"mosn.io/layotto/components/pkg/actuators"
	"mosn.io/layotto/components/pkg/utils"
)

const (
	componentName     = "oss-huaweicloud"
	connectTimeoutSec = "connectTimeoutSec"
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

type HuaweicloudOSS struct {
	client   *obs.ObsClient
	metadata utils.OssMetadata
}

func NewHuaweicloudOSS() oss.Oss { _ = "STUB: not implemented"; return *new(oss.Oss) }

func (h *HuaweicloudOSS) Init(ctx context.Context, config *oss.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *HuaweicloudOSS) GetObject(ctx context.Context, input *oss.GetObjectInput) (*oss.GetObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) PutObject(ctx context.Context, input *oss.PutObjectInput) (*oss.PutObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) DeleteObject(ctx context.Context, input *oss.DeleteObjectInput) (*oss.DeleteObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) PutObjectTagging(ctx context.Context, input *oss.PutObjectTaggingInput) (*oss.PutObjectTaggingOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) DeleteObjectTagging(ctx context.Context, input *oss.DeleteObjectTaggingInput) (*oss.DeleteObjectTaggingOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) GetObjectTagging(ctx context.Context, input *oss.GetObjectTaggingInput) (*oss.GetObjectTaggingOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) CopyObject(ctx context.Context, input *oss.CopyObjectInput) (*oss.CopyObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) DeleteObjects(ctx context.Context, input *oss.DeleteObjectsInput) (*oss.DeleteObjectsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) ListObjects(ctx context.Context, input *oss.ListObjectsInput) (*oss.ListObjectsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) GetObjectCannedAcl(ctx context.Context, input *oss.GetObjectCannedAclInput) (*oss.GetObjectCannedAclOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) PutObjectCannedAcl(ctx context.Context, input *oss.PutObjectCannedAclInput) (*oss.PutObjectCannedAclOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) RestoreObject(ctx context.Context, input *oss.RestoreObjectInput) (*oss.RestoreObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) CreateMultipartUpload(ctx context.Context, input *oss.CreateMultipartUploadInput) (*oss.CreateMultipartUploadOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) UploadPart(ctx context.Context, input *oss.UploadPartInput) (*oss.UploadPartOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) UploadPartCopy(ctx context.Context, input *oss.UploadPartCopyInput) (*oss.UploadPartCopyOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) CompleteMultipartUpload(ctx context.Context, input *oss.CompleteMultipartUploadInput) (*oss.CompleteMultipartUploadOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) AbortMultipartUpload(ctx context.Context, input *oss.AbortMultipartUploadInput) (*oss.AbortMultipartUploadOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) ListMultipartUploads(ctx context.Context, input *oss.ListMultipartUploadsInput) (*oss.ListMultipartUploadsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) ListObjectVersions(ctx context.Context, input *oss.ListObjectVersionsInput) (*oss.ListObjectVersionsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) HeadObject(ctx context.Context, input *oss.HeadObjectInput) (*oss.HeadObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) IsObjectExist(ctx context.Context, input *oss.IsObjectExistInput) (*oss.IsObjectExistOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) SignURL(ctx context.Context, input *oss.SignURLInput) (*oss.SignURLOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) UpdateDownloadBandwidthRateLimit(ctx context.Context, input *oss.UpdateBandwidthRateLimitInput) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *HuaweicloudOSS) UpdateUploadBandwidthRateLimit(ctx context.Context, input *oss.UpdateBandwidthRateLimitInput) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *HuaweicloudOSS) AppendObject(ctx context.Context, input *oss.AppendObjectInput) (*oss.AppendObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// todo 测试异常
func (h *HuaweicloudOSS) ListParts(ctx context.Context, input *oss.ListPartsInput) (*oss.ListPartsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuaweicloudOSS) getClient() (*obs.ObsClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
