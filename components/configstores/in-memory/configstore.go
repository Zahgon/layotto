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

package in_memory

import (
	"context"
	"sync"

	"mosn.io/layotto/components/configstores"
	"mosn.io/layotto/components/pkg/actuators"
)

var (
	once               sync.Once
	readinessIndicator *actuators.HealthIndicator
	livenessIndicator  *actuators.HealthIndicator
)

const (
	componentName = "configstore-memory"
	defaultGroup  = "default"
	defaultLabel  = "default"
)

func init() {
	readinessIndicator = actuators.NewHealthIndicator()
	livenessIndicator = actuators.NewHealthIndicator()
}

type InMemoryConfigStore struct {
	data      *sync.Map
	listener  *sync.Map
	storeName string
	appId     string
}

func NewStore() configstores.Store { _ = "STUB: not implemented"; return *new(configstores.Store) }

func (m *InMemoryConfigStore) Init(config *configstores.StoreConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Get gets configuration from configuration store.
func (m *InMemoryConfigStore) Get(ctx context.Context, req *configstores.GetRequest) ([]*configstores.ConfigurationItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set saves configuration into configuration store.
func (m *InMemoryConfigStore) Set(ctx context.Context, req *configstores.SetRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete deletes configuration from configuration store.
func (m *InMemoryConfigStore) Delete(ctx context.Context, req *configstores.DeleteRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Subscribe gets configuration from configuration store and subscribe the updates.
func (m *InMemoryConfigStore) Subscribe(request *configstores.SubscribeReq, ch chan *configstores.SubscribeResp) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *InMemoryConfigStore) notifyChanged(item *configstores.ConfigurationItem) {
	_ = "STUB: not implemented"
	return
}

type OnChangeFunc func(group, dataId, data string)

func (m *InMemoryConfigStore) subscribeOnChange(ch chan *configstores.SubscribeResp) OnChangeFunc {
	_ = "STUB: not implemented"
	return *new(OnChangeFunc)
}

func (m *InMemoryConfigStore) StopSubscribe() {
	_ = "STUB: not implemented"
	// stop listening all subscribed configs
	return
}

func (m *InMemoryConfigStore) GetDefaultGroup() string { _ = "STUB: not implemented"; return "" }

func (m *InMemoryConfigStore) GetDefaultLabel() string { _ = "STUB: not implemented"; return "" }
