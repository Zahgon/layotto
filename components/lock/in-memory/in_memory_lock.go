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

package in_memory

import (
	"context"
	"sync"
	"time"

	"mosn.io/layotto/components/lock"
	"mosn.io/layotto/components/pkg/actuators"
)

const (
	componentName = "lock-memory"
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

type InMemoryLock struct {
	features []lock.Feature
	data     *lockMap
}

// memoryLock is a lock holder
type memoryLock struct {
	key        string
	owner      string
	expireTime time.Time
	lock       int
}

type lockMap struct {
	sync.Mutex
	locks map[string]*memoryLock
}

func NewInMemoryLock() *InMemoryLock { _ = "STUB: not implemented"; return nil }

func (s *InMemoryLock) Init(_ lock.Metadata) error { _ = "STUB: not implemented"; return nil }

// LockKeepAlive try to renewal lease
func (s *InMemoryLock) LockKeepAlive(ctx context.Context, request *lock.LockKeepAliveRequest) (*lock.LockKeepAliveResponse, error) {
	_ = "STUB: not implemented"
	//TODO: implemnt function
	return nil, nil
}

func (s *InMemoryLock) Features() []lock.Feature {
	_ = "STUB: not implemented"

	// Try to add a lock. Currently this is a non-reentrant lock
	return nil
}

func (s *InMemoryLock) TryLock(ctx context.Context, req *lock.TryLockRequest) (*lock.TryLockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. Find the memoryLock for this resourceId

//0 unlock, 1 lock

// 2. Construct a new one if the lockData has expired
//check expire

// 3. Check if it has been locked by others.
// Currently this is a non-reentrant lock

//lock failed

// 4. Update owner information

func (s *InMemoryLock) Unlock(ctx context.Context, req *lock.UnlockRequest) (*lock.UnlockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. Find the memoryLock for this resourceId

// 2. check the owner information

// 3. unlock and reset the owner information
