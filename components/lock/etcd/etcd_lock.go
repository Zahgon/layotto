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

	"mosn.io/layotto/components/pkg/utils"

	"mosn.io/layotto/kit/logger"

	"mosn.io/layotto/components/lock"
	"mosn.io/layotto/components/pkg/actuators"
)

const (
	componentName = "lock-etcd"
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

// Etcd lock store
type EtcdLock struct {
	client   *clientv3.Client
	metadata utils.EtcdMetadata

	features []lock.Feature

	ctx    context.Context
	cancel context.CancelFunc
	logger logger.Logger
}

// NewEtcdLock returns a new etcd lock
func NewEtcdLock() *EtcdLock { _ = "STUB: not implemented"; return nil }

func (e *EtcdLock) OnLogLevelChanged(outputLevel logger.LogLevel) {
	_ = "STUB: not implemented"
	return
}

// Init EtcdLock
func (e *EtcdLock) Init(metadata lock.Metadata) error {
	_ = "STUB: not implemented"
	// 1. parse config
	return nil
}

// 2. construct client

// LockKeepAlive try to renewal lease
func (e *EtcdLock) LockKeepAlive(ctx context.Context, request *lock.LockKeepAliveRequest) (*lock.LockKeepAliveResponse, error) {
	_ = "STUB: not implemented"
	//TODO: implemnt function
	return nil, nil
}

// Features is to get EtcdLock's features
func (e *EtcdLock) Features() []lock.Feature {
	_ = "STUB: not implemented"

	// Node tries to acquire a etcd lock
	return nil
}

func (e *EtcdLock) TryLock(ctx context.Context, req *lock.TryLockRequest) (*lock.TryLockResponse, error) {
	_ = "STUB: not implemented"
	return nil,

		//1.Create new lease
		nil
}

//2.Create new KV

//3.Create txn

//4.Commit and try get lock

// Node tries to release a etcd lock
func (e *EtcdLock) Unlock(ctx context.Context, req *lock.UnlockRequest) (*lock.UnlockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// 1.Create new KV
}

// 2.Create txn

// 3.Commit and try release lock

// Close shuts down the client's etcd connections.
func (e *EtcdLock) Close() error { _ = "STUB: not implemented"; return nil }

// getkey is to return string of type KeyPrefix + resourceId
func (e *EtcdLock) getKey(resourceId string) string { _ = "STUB: not implemented"; return "" }

// newInternalErrorUnlockResponse is to return lock release error
func newInternalErrorUnlockResponse() *lock.UnlockResponse { _ = "STUB: not implemented"; return nil }
