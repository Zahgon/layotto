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

package redis

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
	msync "mosn.io/mosn/pkg/sync"

	"mosn.io/layotto/kit/logger"

	"mosn.io/layotto/components/lock"
	"mosn.io/layotto/components/pkg/actuators"
	"mosn.io/layotto/components/pkg/utils"
)

const (
	componentName = "lock-redis-cluster"
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

// RedLock
// it will be best to use at least 5 hosts
type ClusterRedisLock struct {
	clients  []*redis.Client
	metadata utils.RedisClusterMetadata
	workpool msync.WorkerPool

	features []lock.Feature
	logger   logger.Logger

	ctx    context.Context
	cancel context.CancelFunc
}

// NewClusterRedisLock returns a new redis lock store
func NewClusterRedisLock() *ClusterRedisLock { _ = "STUB: not implemented"; return nil }

func (c *ClusterRedisLock) OnLogLevelChanged(outputLevel logger.LogLevel) {
	_ = "STUB: not implemented"
	return
}

type resultMsg struct {
	error        error
	host         string
	lockStatus   bool
	unlockStatus lock.LockStatus
}

func (c *ClusterRedisLock) Init(metadata lock.Metadata) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterRedisLock) Features() []lock.Feature {
	_ = "STUB: not implemented"

	// LockKeepAlive try to renewal lease
	return nil
}

func (c *ClusterRedisLock) LockKeepAlive(ctx context.Context, request *lock.LockKeepAliveRequest) (*lock.LockKeepAliveResponse, error) {
	_ = "STUB: not implemented"
	//TODO: implemnt function
	return nil, nil
}

func (c *ClusterRedisLock) TryLock(ctx context.Context, req *lock.TryLockRequest) (*lock.TryLockResponse, error) {
	_ = "STUB: not implemented"
	//try to get lock on all redis nodes
	return nil, nil
}

//intervalLimit must be 1/10 of expire time to make sure time of lock far less than expire time

//resultChan will be used to collect results of getting lock

//getting lock concurrently

//make sure time interval of locking far less than expire time

//getting lock on majority of redis cluster will be regarded as locking success

func (c *ClusterRedisLock) Unlock(ctx context.Context, req *lock.UnlockRequest) (*lock.UnlockResponse, error) {
	_ = "STUB: not implemented"
	return nil,

		//err means there were some internal errors,then the status must be INTERNAL_ERROR
		//the LOCK_UNEXIST and LOCK_BELONG_TO_OTHERS status codes can be ignore
		//becauce they means the lock of the current redis
		//returned the status code don't need to be unlocked by current invoking
		nil
}

func (c *ClusterRedisLock) UnlockAllRedis(req *lock.UnlockRequest, wg *sync.WaitGroup) (lock.LockStatus, error) {
	_ = "STUB: not implemented"
	return *new(lock.LockStatus), nil
}

//unlock concurrently

//collect result of unlocking

func (c *ClusterRedisLock) LockSingleRedis(clientIndex int, req *lock.TryLockRequest, wg *sync.WaitGroup, ch chan resultMsg) {
	_ = "STUB: not implemented"
	return
}

func (c *ClusterRedisLock) UnlockSingleRedis(clientIndex int, req *lock.UnlockRequest, wg *sync.WaitGroup, ch chan resultMsg) {
	_ = "STUB: not implemented"
	return
}
