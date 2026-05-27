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

package consul

import (
	"context"
	"sync"

	msync "mosn.io/mosn/pkg/sync"

	log "mosn.io/layotto/kit/logger"

	"mosn.io/layotto/components/lock"
	"mosn.io/layotto/components/pkg/actuators"
	"mosn.io/layotto/components/pkg/utils"
)

const (
	componentName = "lock-consul"
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

type ConsulLock struct {
	metadata       utils.ConsulMetadata
	log            log.Logger
	client         utils.ConsulClient
	sessionFactory utils.SessionFactory
	kv             utils.ConsulKV
	sMap           sync.Map
	workPool       msync.WorkerPool
}

func NewConsulLock() *ConsulLock { _ = "STUB: not implemented"; return nil }

func (c *ConsulLock) OnLogLevelChanged(outputLevel log.LogLevel) { _ = "STUB: not implemented"; return }

func (c *ConsulLock) Init(metadata lock.Metadata) error { _ = "STUB: not implemented"; return nil }

func (c *ConsulLock) Features() []lock.Feature {
	_ = "STUB: not implemented"

	// LockKeepAlive try to renewal lease
	return nil
}

func (c *ConsulLock) LockKeepAlive(ctx context.Context, request *lock.LockKeepAliveRequest) (*lock.LockKeepAliveResponse, error) {
	_ = "STUB: not implemented"
	//TODO: implemnt function
	return nil, nil
}

func getTTL(expire int32) string {
	_ = "STUB: not implemented"
	// session TTL must be between [10s=24h0m0s]
	return ""
}

func (c *ConsulLock) TryLock(ctx context.Context, req *lock.TryLockRequest) (*lock.TryLockResponse, error) {
	_ = "STUB: not implemented"

	// create a session TTL
	return nil, nil
}

//Controls the behavior to delete when a session is invalidated.

// put a new KV pair with ttl session

//acquire lock

//bind lockOwner+resourceId and session

func (c *ConsulLock) Unlock(ctx context.Context, req *lock.UnlockRequest) (*lock.UnlockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// put a new KV pair with ttl session

//release lock
