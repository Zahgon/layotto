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
	"encoding/json"
	"sync"

	"mosn.io/layotto/components/pkg/actuators"

	"mosn.io/layotto/components/oss"

	"github.com/aws/aws-sdk-go-v2/service/s3"

	"mosn.io/layotto/kit/logger"
)

const (
	componentName = "oss-aws"
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

type AwsOss struct {
	client    *s3.Client
	basicConf json.RawMessage
	logger    logger.Logger
}

func NewAwsOss() oss.Oss { _ = "STUB: not implemented"; return *new(oss.Oss) }

func (a *AwsOss) OnLogLevelChanged(level logger.LogLevel) { _ = "STUB: not implemented"; return }

func (a *AwsOss) Init(ctx context.Context, config *oss.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AwsOss) GetObject(ctx context.Context, req *oss.GetObjectInput) (*oss.GetObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) PutObject(ctx context.Context, req *oss.PutObjectInput) (*oss.PutObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) DeleteObject(ctx context.Context, req *oss.DeleteObjectInput) (*oss.DeleteObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) PutObjectTagging(ctx context.Context, req *oss.PutObjectTaggingInput) (*oss.PutObjectTaggingOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) DeleteObjectTagging(ctx context.Context, req *oss.DeleteObjectTaggingInput) (*oss.DeleteObjectTaggingOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) GetObjectTagging(ctx context.Context, req *oss.GetObjectTaggingInput) (*oss.GetObjectTaggingOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) CopyObject(ctx context.Context, req *oss.CopyObjectInput) (*oss.CopyObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//TODO: should support objects accessed through access points

func (a *AwsOss) DeleteObjects(ctx context.Context, req *oss.DeleteObjectsInput) (*oss.DeleteObjectsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) ListObjects(ctx context.Context, req *oss.ListObjectsInput) (*oss.ListObjectsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) GetObjectCannedAcl(ctx context.Context, req *oss.GetObjectCannedAclInput) (*oss.GetObjectCannedAclOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) PutObjectCannedAcl(ctx context.Context, req *oss.PutObjectCannedAclInput) (*oss.PutObjectCannedAclOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) RestoreObject(ctx context.Context, req *oss.RestoreObjectInput) (*oss.RestoreObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) CreateMultipartUpload(ctx context.Context, req *oss.CreateMultipartUploadInput) (*oss.CreateMultipartUploadOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) UploadPart(ctx context.Context, req *oss.UploadPartInput) (*oss.UploadPartOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) UploadPartCopy(ctx context.Context, req *oss.UploadPartCopyInput) (*oss.UploadPartCopyOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//TODO: should support objects accessed through access points

func (a *AwsOss) CompleteMultipartUpload(ctx context.Context, req *oss.CompleteMultipartUploadInput) (*oss.CompleteMultipartUploadOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) AbortMultipartUpload(ctx context.Context, req *oss.AbortMultipartUploadInput) (*oss.AbortMultipartUploadOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) ListMultipartUploads(ctx context.Context, req *oss.ListMultipartUploadsInput) (*oss.ListMultipartUploadsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) ListObjectVersions(ctx context.Context, req *oss.ListObjectVersionsInput) (*oss.ListObjectVersionsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) HeadObject(ctx context.Context, req *oss.HeadObjectInput) (*oss.HeadObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) IsObjectExist(ctx context.Context, req *oss.IsObjectExistInput) (*oss.IsObjectExistOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) SignURL(ctx context.Context, req *oss.SignURLInput) (*oss.SignURLOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) UpdateDownloadBandwidthRateLimit(ctx context.Context, req *oss.UpdateBandwidthRateLimitInput) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AwsOss) UpdateUploadBandwidthRateLimit(ctx context.Context, req *oss.UpdateBandwidthRateLimitInput) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AwsOss) AppendObject(ctx context.Context, req *oss.AppendObjectInput) (*oss.AppendObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) ListParts(ctx context.Context, req *oss.ListPartsInput) (*oss.ListPartsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AwsOss) getClient() (*s3.Client, error) { _ = "STUB: not implemented"; return nil, nil }
