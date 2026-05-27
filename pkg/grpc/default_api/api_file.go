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

package default_api

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	runtimev1pb "mosn.io/layotto/spec/proto/runtime/v1"
)

func (a *api) GetFile(req *runtimev1pb.GetFileRequest, stream runtimev1pb.Runtime_GetFileServer) error {
	_ = "STUB: not implemented"
	return nil
}

type putObjectStreamReader struct {
	data   []byte
	server runtimev1pb.Runtime_PutFileServer
}

func newPutObjectStreamReader(data []byte, server runtimev1pb.Runtime_PutFileServer) *putObjectStreamReader {
	_ = "STUB: not implemented"
	return nil
}

func (r *putObjectStreamReader) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (a *api) PutFile(stream runtimev1pb.Runtime_PutFileServer) error {
	_ = "STUB: not implemented"
	return nil
}

//if client send eof error directly, return nil

// ListFile list all files
func (a *api) ListFile(ctx context.Context, in *runtimev1pb.ListFileRequest) (*runtimev1pb.ListFileResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DelFile delete specific file
func (a *api) DelFile(ctx context.Context, in *runtimev1pb.DelFileRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFileMeta get meta of file
func (a *api) GetFileMeta(ctx context.Context, in *runtimev1pb.GetFileMetaRequest) (*runtimev1pb.GetFileMetaResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
