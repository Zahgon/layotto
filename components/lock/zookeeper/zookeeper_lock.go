//
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

package zookeeper

import (
	"context"
	"sync"
	"time"

	"mosn.io/layotto/kit/logger"

	"mosn.io/layotto/components/lock"
	"mosn.io/layotto/components/pkg/actuators"
	"mosn.io/layotto/components/pkg/utils"
)

var (
	closeConn = func(conn utils.ZKConnection, expireInSecond int32) {
		//can also
		//time.Sleep(time.Second * time.Duration(expireInSecond))
		<-time.After(time.Second * time.Duration(expireInSecond))
		// make sure close connecion
		conn.Close()
	}
	once               sync.Once
	readinessIndicator *actuators.HealthIndicator
	livenessIndicator  *actuators.HealthIndicator
)

const (
	componentName = "lock-zookeeper"
)

func init() {
	readinessIndicator = actuators.NewHealthIndicator()
	livenessIndicator = actuators.NewHealthIndicator()
}

// ZookeeperLock lock store
type ZookeeperLock struct {
	//trylock reestablish connection  every time
	factory utils.ConnectionFactory
	//unlock reuse this conneciton
	unlockConn utils.ZKConnection
	metadata   utils.ZookeeperMetadata
	logger     logger.Logger
}

// NewZookeeperLock Create ZookeeperLock
func NewZookeeperLock() *ZookeeperLock { _ = "STUB: not implemented"; return nil }

// OnLogLevelChanged change log level
func (p *ZookeeperLock) OnLogLevelChanged(level logger.LogLevel) { _ = "STUB: not implemented"; return }

// Init ZookeeperLock
func (p *ZookeeperLock) Init(metadata lock.Metadata) error { _ = "STUB: not implemented"; return nil }

//init unlock connection

// Features is to get ZookeeperLock's features
func (p *ZookeeperLock) Features() []lock.Feature {
	_ = "STUB: not implemented"

	// LockKeepAlive try to renewal lease
	return nil
}

func (p *ZookeeperLock) LockKeepAlive(ctx context.Context, request *lock.LockKeepAliveRequest) (*lock.LockKeepAliveResponse, error) {
	_ = "STUB: not implemented"
	//TODO: implemnt function
	return nil, nil
}

// TryLock Node tries to acquire a zookeeper lock
func (p *ZookeeperLock) TryLock(ctx context.Context, req *lock.TryLockRequest) (*lock.TryLockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//1.create zk ephemeral node

//2.1 create node fail ,indicates lock fail

//the node exists,lock fail

//other err

//2.2 create node success, asyn  to make sure zkclient alive for need time

// Unlock Node tries to release a zookeeper lock
func (p *ZookeeperLock) Unlock(ctx context.Context, req *lock.UnlockRequest) (*lock.UnlockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//node does not exist, indicates this lock has expired

//other err

//node exist ,but owner not this, indicates this lock has occupied or wrong unlock

//owner is this, but delete fail

// delete no node , indicates this lock has expired

// delete version error , indicates this lock has occupied by others

//other error

//delete success, unlock success
