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

package local

import (
	"context"
	"io"
	"sync"

	"mosn.io/layotto/components/file"
	"mosn.io/layotto/components/pkg/actuators"
)

const (
	FileMode      = "FileMode"
	FileFlag      = "FileFlag"
	FileIsDir     = "IsDir"
	componentName = "file-local"
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

type LocalStore struct {
}

func NewLocalStore() file.File { _ = "STUB: not implemented"; return *new(file.File) }

func (lf *LocalStore) Init(ctx context.Context, f *file.FileConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (lf *LocalStore) Put(ctx context.Context, f *file.PutFileStu) error {
	_ = "STUB: not implemented"
	return nil
}

func (lf *LocalStore) Get(ctx context.Context, f *file.GetFileStu) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (lf *LocalStore) List(ctx context.Context, f *file.ListRequest) (*file.ListResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lf *LocalStore) Del(ctx context.Context, f *file.DelRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (lf *LocalStore) Stat(ctx context.Context, f *file.FileMetaRequest) (*file.FileMetaResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
