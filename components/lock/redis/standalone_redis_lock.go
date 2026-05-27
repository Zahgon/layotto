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

	"github.com/go-redis/redis/v8"

	"mosn.io/layotto/kit/logger"

	"mosn.io/layotto/components/lock"
	"mosn.io/layotto/components/pkg/actuators"
	"mosn.io/layotto/components/pkg/utils"
)

func init() {
	readinessIndicator = actuators.NewHealthIndicator()
	livenessIndicator = actuators.NewHealthIndicator()
}

// Standalone Redis lock store.Any fail-over related features are not supported,such as Sentinel and Redis Cluster.
type StandaloneRedisLock struct {
	client   *redis.Client
	metadata utils.RedisMetadata

	features []lock.Feature
	logger   logger.Logger

	ctx    context.Context
	cancel context.CancelFunc
}

// NewStandaloneRedisLock returns a new redis lock store
func NewStandaloneRedisLock() *StandaloneRedisLock { _ = "STUB: not implemented"; return nil }

func (p *StandaloneRedisLock) OnLogLevelChanged(outputLevel logger.LogLevel) {
	_ = "STUB: not implemented"
	return
}

// Init StandaloneRedisLock
func (p *StandaloneRedisLock) Init(metadata lock.Metadata) error {
	_ = "STUB: not implemented"
	// 1. parse config
	return nil
}

// 2. construct client

// 3. connect to redis

// Features is to get StandaloneRedisLock's features
func (p *StandaloneRedisLock) Features() []lock.Feature {
	_ = "STUB: not implemented"

	// LockKeepAlive try to renewal lease
	return nil
}

func (p *StandaloneRedisLock) LockKeepAlive(ctx context.Context, request *lock.LockKeepAliveRequest) (*lock.LockKeepAliveResponse, error) {
	_ = "STUB: not implemented"
	//TODO: implemnt function
	return nil, nil
}

// Node tries to acquire a redis lock
func (p *StandaloneRedisLock) TryLock(ctx context.Context, req *lock.TryLockRequest) (*lock.TryLockResponse, error) {
	_ = "STUB: not implemented"
	// 1.Setting redis expiration time
	return nil, nil
}

// 2. check error

const unlockScript = "local v = redis.call(\"get\",KEYS[1]); if v==false then return -1 end; if v~=ARGV[1] then return -2 else return redis.call(\"del\",KEYS[1]) end"

// Node tries to release a redis lock
func (p *StandaloneRedisLock) Unlock(ctx context.Context, req *lock.UnlockRequest) (*lock.UnlockResponse, error) {
	_ = "STUB: not implemented"
	// 1. delegate to client.eval lua script
	return nil, nil
}

// 2. check error

// 3. parse result

// newInternalErrorUnlockResponse is to return lock release error
func newInternalErrorUnlockResponse() *lock.UnlockResponse { _ = "STUB: not implemented"; return nil }

// Close shuts down the client's redis connections.
func (p *StandaloneRedisLock) Close() error { _ = "STUB: not implemented"; return nil }
