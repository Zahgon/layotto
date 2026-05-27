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

package mongo

import (
	"context"
	"sync"

	"mosn.io/layotto/kit/logger"

	"mosn.io/layotto/components/lock"
	"mosn.io/layotto/components/pkg/actuators"
	"mosn.io/layotto/components/pkg/utils"
)

const (
	TRY_LOCK_SUCCESS        = 1
	TRY_LOCK_FAIL           = 2
	UNLOCK_SUCCESS          = 3
	UNLOCK_UNEXIST          = 4
	UNLOCK_BELONG_TO_OTHERS = 5
	UNLOCK_FAIL             = 6
	componentName           = "lock-mongo"
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

// mongo lock store
type MongoLock struct {
	factory utils.MongoFactory

	client     utils.MongoClient
	session    utils.MongoSession
	collection utils.MongoCollection
	metadata   utils.MongoMetadata

	features []lock.Feature
	logger   logger.Logger

	ctx    context.Context
	cancel context.CancelFunc
}

// NewMongoLock returns a new mongo lock
func NewMongoLock() *MongoLock { _ = "STUB: not implemented"; return nil }

func (e *MongoLock) OnLogLevelChanged(outputLevel logger.LogLevel) {
	_ = "STUB: not implemented"
	return
}

func (e *MongoLock) Init(metadata lock.Metadata) error { _ = "STUB: not implemented"; return nil }

// 1.parse config

// 2. construct client

// Connections Collection

// create exprie time index

// Features is to get MongoLock's features
func (e *MongoLock) Features() []lock.Feature {
	_ = "STUB: not implemented"

	// LockKeepAlive try to renewal lease
	return nil
}

func (e *MongoLock) LockKeepAlive(ctx context.Context, request *lock.LockKeepAliveRequest) (*lock.LockKeepAliveResponse, error) {
	_ = "STUB: not implemented"
	//TODO: implemnt function
	return nil, nil
}

func (e *MongoLock) TryLock(ctx context.Context, req *lock.TryLockRequest) (*lock.TryLockResponse, error) {
	_ = "STUB: not implemented"

	// create mongo session
	return nil, nil
}

// check session

// close mongo session

// start transaction

// set exprie date

// insert mongo lock

// commit and set status

// check lock

func (e *MongoLock) Unlock(ctx context.Context, req *lock.UnlockRequest) (*lock.UnlockResponse, error) {
	_ = "STUB: not implemented"

	// create mongo session
	return nil, nil
}

// check session

// close mongo session

// start transaction

// delete lock

// check delete result

// commit and set status

func newInternalErrorUnlockResponse() *lock.UnlockResponse { _ = "STUB: not implemented"; return nil }

func (e *MongoLock) Close() error { _ = "STUB: not implemented"; return nil }
