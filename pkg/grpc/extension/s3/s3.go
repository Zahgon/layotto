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

package s3

import (
	"context"
	"sync"

	"mosn.io/layotto/spec/proto/extension/v1/s3"

	l8s3 "mosn.io/layotto/components/oss"

	rawGRPC "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"mosn.io/layotto/pkg/grpc"
)

var (
	s3Instance *S3Server
)

var (
	bytesPool = sync.Pool{
		New: func() interface{} {
			// set size to 100kb
			return new([]byte)
		},
	}
)

type S3Server struct {
	appId       string
	ossInstance map[string]l8s3.Oss
}

func NewS3Server(ac *grpc.ApplicationContext) grpc.GrpcAPI {
	_ = "STUB: not implemented"
	return *new(grpc.GrpcAPI)
}

func (s *S3Server) Init(conn *rawGRPC.ClientConn) error { _ = "STUB: not implemented"; return nil }

func (s *S3Server) Register(rawGrpcServer *rawGRPC.Server) error {
	_ = "STUB: not implemented"
	return nil
}

func transferData(source interface{}, target interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *S3Server) GetObject(req *s3.GetObjectInput, stream s3.ObjectStorageService_GetObjectServer) error {
	_ = "STUB: not implemented"
	// 1. validate
	return nil
}

// 2. convert request

// 3. find the component

// Reduce the number of transmissions to reduce the performance impact caused by the introduction of sidecar

type putObjectStreamReader struct {
	data   []byte
	server s3.ObjectStorageService_PutObjectServer
}

func newPutObjectStreamReader(data []byte, server s3.ObjectStorageService_PutObjectServer) *putObjectStreamReader {
	_ = "STUB: not implemented"
	return nil
}

func (r *putObjectStreamReader) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *S3Server) PutObject(stream s3.ObjectStorageService_PutObjectServer) error {
	_ = "STUB: not implemented"
	return nil
}

//if client send eof error directly, return nil

func (s *S3Server) DeleteObject(ctx context.Context, req *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) PutObjectTagging(ctx context.Context, req *s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) DeleteObjectTagging(ctx context.Context, req *s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) GetObjectTagging(ctx context.Context, req *s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) CopyObject(ctx context.Context, req *s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) DeleteObjects(ctx context.Context, req *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) ListObjects(ctx context.Context, req *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) GetObjectCannedAcl(ctx context.Context, req *s3.GetObjectCannedAclInput) (*s3.GetObjectCannedAclOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) PutObjectCannedAcl(ctx context.Context, req *s3.PutObjectCannedAclInput) (*s3.PutObjectCannedAclOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) RestoreObject(ctx context.Context, req *s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) CreateMultipartUpload(ctx context.Context, req *s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type uploadPartStreamReader struct {
	data   []byte
	server s3.ObjectStorageService_UploadPartServer
}

func newUploadPartStreamReader(data []byte, server s3.ObjectStorageService_UploadPartServer) *uploadPartStreamReader {
	_ = "STUB: not implemented"
	return nil
}

func (r *uploadPartStreamReader) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *S3Server) UploadPart(stream s3.ObjectStorageService_UploadPartServer) error {
	_ = "STUB: not implemented"
	return nil
}

//if client send eof error directly, return nil

func (s *S3Server) UploadPartCopy(ctx context.Context, req *s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) CompleteMultipartUpload(ctx context.Context, req *s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) AbortMultipartUpload(ctx context.Context, req *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) ListMultipartUploads(ctx context.Context, req *s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) ListObjectVersions(ctx context.Context, req *s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) HeadObject(ctx context.Context, req *s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) IsObjectExist(ctx context.Context, req *s3.IsObjectExistInput) (*s3.IsObjectExistOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) SignURL(ctx context.Context, req *s3.SignURLInput) (*s3.SignURLOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) UpdateDownloadBandwidthRateLimit(ctx context.Context, req *s3.UpdateBandwidthRateLimitInput) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *S3Server) UpdateUploadBandwidthRateLimit(ctx context.Context, req *s3.UpdateBandwidthRateLimitInput) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type appendObjectStreamReader struct {
	data   []byte
	server s3.ObjectStorageService_AppendObjectServer
}

func newAppendObjectStreamReader(data []byte, server s3.ObjectStorageService_AppendObjectServer) *appendObjectStreamReader {
	_ = "STUB: not implemented"
	return nil
}

func (r *appendObjectStreamReader) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *S3Server) AppendObject(stream s3.ObjectStorageService_AppendObjectServer) error {
	_ = "STUB: not implemented"
	return nil
}

//if client send eof error directly, return nil

func (s *S3Server) ListParts(ctx context.Context, req *s3.ListPartsInput) (*s3.ListPartsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
