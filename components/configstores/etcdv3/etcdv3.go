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

package etcdv3

import (
	"context"
	"sync"

	"mosn.io/layotto/components/pkg/actuators"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"

	log "mosn.io/layotto/kit/logger"

	"mosn.io/layotto/components/configstores"
)

const (
	defaultGroup  = "default"
	defaultLabel  = "default"
	componentName = "configstore-etcdv3"
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

type EtcdV3ConfigStore struct {
	client *clientv3.Client
	sync.RWMutex
	subscribeKey map[string]string
	appIdKey     string
	storeName    string
	// cancel is the func, call cancel will stop watching on the appIdKey
	cancel       context.CancelFunc
	watchStarted bool
	watchRespCh  chan *configstores.SubscribeResp
	log          log.Logger
}

func (c *EtcdV3ConfigStore) GetDefaultGroup() string { _ = "STUB: not implemented"; return "" }

func (c *EtcdV3ConfigStore) GetDefaultLabel() string { _ = "STUB: not implemented"; return "" }

func (c *EtcdV3ConfigStore) OnLogLevelChanged(outputLevel log.LogLevel) {
	_ = "STUB: not implemented"
	return
}

func NewStore() configstores.Store { _ = "STUB: not implemented"; return *new(configstores.Store) }

// Init init the configuration store.
func (c *EtcdV3ConfigStore) Init(config *configstores.StoreConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *EtcdV3ConfigStore) GetPrimaryKeyWithoutTag(s string) string {
	_ = "STUB: not implemented"
	//key no tag
	return ""
}

func (c *EtcdV3ConfigStore) GetItemsFromAllKeys(kvs []*mvccpb.KeyValue, targetString []string) []*configstores.ConfigurationItem {
	_ = "STUB: not implemented"
	return nil
}

// Get gets configuration from configuration store.
func (c *EtcdV3ConfigStore) Get(ctx context.Context, req *configstores.GetRequest) ([]*configstores.ConfigurationItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//TODO: the imp read all keys under app, then do match operation, should change later.

// Set saves configuration into configuration store.
func (c *EtcdV3ConfigStore) Set(ctx context.Context, req *configstores.SetRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete deletes configuration from configuration store.
func (c *EtcdV3ConfigStore) Delete(ctx context.Context, req *configstores.DeleteRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *EtcdV3ConfigStore) processWatchResponse(resp *clientv3.WatchResponse) {
	_ = "STUB: not implemented"
	return
}

func (c *EtcdV3ConfigStore) watch() {
	_ = "STUB: not implemented"
	// Add watch for propertyKey from lastUpdatedRevision updated after Initializing
	return
}

// Subscribe gets configuration from configuration store and subscribe the updates.
func (c *EtcdV3ConfigStore) Subscribe(req *configstores.SubscribeReq, ch chan *configstores.SubscribeResp) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *EtcdV3ConfigStore) StopSubscribe() { _ = "STUB: not implemented"; return }

func (c *EtcdV3ConfigStore) ParseKey(appId string, req *configstores.ConfigurationItem) []string {
	_ = "STUB: not implemented"
	return nil
}
