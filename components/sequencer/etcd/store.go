// Copyright 2021 Layotto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package etcd

import (
	"context"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"

	"mosn.io/layotto/kit/logger"

	"mosn.io/layotto/components/pkg/actuators"
	"mosn.io/layotto/components/pkg/utils"
	"mosn.io/layotto/components/sequencer"
)

const (
	componentName = "sequencer-etcd"
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

type EtcdSequencer struct {
	client     *clientv3.Client
	metadata   utils.EtcdMetadata
	biggerThan map[string]int64

	logger logger.Logger

	ctx    context.Context
	cancel context.CancelFunc
}

// EtcdSequencer returns a new etcd sequencer
func NewEtcdSequencer() *EtcdSequencer { _ = "STUB: not implemented"; return nil }

func (e *EtcdSequencer) OnLogLevelChanged(level logger.LogLevel) { _ = "STUB: not implemented"; return }

func (e *EtcdSequencer) Init(config sequencer.Configuration) error {
	_ = "STUB: not implemented"
	// 1. parse config
	return nil
}

// 2. construct client

// 3. check biggerThan

// TODO close component?

func (e *EtcdSequencer) GetNextId(req *sequencer.GetNextIdRequest) (*sequencer.GetNextIdResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Create new KV
}

// Create txn

// Commit

func (e *EtcdSequencer) GetSegment(req *sequencer.GetSegmentRequest) (support bool, result *sequencer.GetSegmentResponse, err error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (e *EtcdSequencer) Close() error { _ = "STUB: not implemented"; return nil }

func (e *EtcdSequencer) getKeyInEtcd(key string) string { _ = "STUB: not implemented"; return "" }

func addPathSeparator(p string) string { _ = "STUB: not implemented"; return "" }
