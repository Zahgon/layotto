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

package nacos

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"sync"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"

	log "mosn.io/layotto/kit/logger"

	"mosn.io/layotto/components/configstores"
	"mosn.io/layotto/components/pkg/actuators"
)

const (
	componentName = "configstore-nacos"
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

type ConfigStore struct {
	client      config_client.IConfigClient
	storeName   string
	appId       string
	namespaceId string
	listener    sync.Map
	log         log.Logger
}

func NewStore() configstores.Store { _ = "STUB: not implemented"; return *new(configstores.Store) }

func (n *ConfigStore) OnLogLevelChanged(outputLevel log.LogLevel) {
	_ = "STUB: not implemented"
	return
}

// Init SetConfig the configuration store.
func (n *ConfigStore) Init(config *configstores.StoreConfig) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// store name, required

// parse config metadata

// the nacos's addresses, required if not using acm mode.

// the timeout of connect to nacos, not required

// choose different mode to connect to the nacos server.

// replace nacos sdk log

// Connect to self built nacos services
func (n *ConfigStore) init(address []string, timeoutMs uint64, metadata *Metadata) (config_client.IConfigClient, error) {
	// 1.create ServerConfigs
	serverConfigs := make([]constant.ServerConfig, 0, len(address))
	for _, v := range address {
		// split the addresses to ip and port
		splitAddr := strings.Split(v, ":")
		if len(splitAddr) != 2 {
			return nil, errors.New("configuration illegal: addresses is not in the format of ip:port")
		}

		ip := splitAddr[0]
		port, err := strconv.Atoi(splitAddr[1])
		if err != nil {
			return nil, errors.New("configuration illegal: can't convert port form string to int type")
		}
		// default use http schema and use nacos as the context
		sc := *constant.NewServerConfig(ip, uint64(port))
		serverConfigs = append(serverConfigs, sc)
	}

	// 2.create client config
	clientConfig := *constant.NewClientConfig(
		constant.WithTimeoutMs(timeoutMs),
		constant.WithNamespaceId(metadata.NameSpaceId),
		constant.WithUsername(metadata.Username),
		constant.WithPassword(metadata.Password),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithCacheDir(metadata.CacheDir),
	)

	// 3.create config client
	// it only creates a client instance but not connect to nacos.
	// so if the address is wrong, the client instance will still be created successfully.
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Connect to the nacos service provided by Alibaba Cloud
func (n *ConfigStore) initWithACM(timeoutMs uint64, metadata *Metadata) (config_client.IConfigClient, error) {
	_ = "STUB: not implemented"
	return *new(config_client.IConfigClient), nil
}

// a more graceful way to create config client

func (n *ConfigStore) setupLogger(metadata *Metadata) error { _ = "STUB: not implemented"; return nil }

// Get gets configuration from configuration store.
func (n *ConfigStore) Get(ctx context.Context, request *configstores.GetRequest) ([]*configstores.ConfigurationItem, error) {
	_ = "STUB: not implemented"
	// use the configuration's app_name instead of the app_id in request
	// 0. check if illegal
	return nil, nil
}

// 1. get pagination information

// 2. app level

// 3.group level

// 4.key level

const (
	PageNo   = "PageNo"
	PageSize = "PageSize"
)

type Pagination struct {
	PageNo   int
	PageSize int
}

func (n *ConfigStore) getPagination(metadata map[string]string) *Pagination {
	_ = "STUB: not implemented"
	return nil
}

func (n *ConfigStore) getAllWithAppId(ctx context.Context, pagination *Pagination) ([]*configstores.ConfigurationItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *ConfigStore) getAllWithGroup(ctx context.Context, group string, pagination *Pagination) ([]*configstores.ConfigurationItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *ConfigStore) getAllWithKeys(ctx context.Context, group string, keys []string) ([]*configstores.ConfigurationItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// todo: make more goroutine to search the configurations.

// config is not exist
// nacos does not support an empty content.

func (n *ConfigStore) Set(ctx context.Context, request *configstores.SetRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// If the config does not exist, deleting the config will not result in an error.

func (n *ConfigStore) Delete(ctx context.Context, request *configstores.DeleteRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// remove the config change listening

func (n *ConfigStore) Subscribe(request *configstores.SubscribeReq, ch chan *configstores.SubscribeResp) error {
	_ = "STUB: not implemented"
	return nil
}

// todo: use errgroup to deal with it concurrently.

type subscriberKey struct {
	group string
	key   string
}

func (n *ConfigStore) subscribeKey(item *configstores.ConfigurationItem, ch chan *configstores.SubscribeResp) error {
	_ = "STUB: not implemented"
	return nil
}

type OnChangeFunc func(namespace, group, dataId, data string)

func (n *ConfigStore) subscribeOnChange(ch chan *configstores.SubscribeResp) OnChangeFunc {
	_ = "STUB: not implemented"
	return *new(OnChangeFunc)
}

// package the listening data.

func (n *ConfigStore) StopSubscribe() {
	_ = "STUB: not implemented"
	// stop listening all subscribed configs
	return
}

func (n *ConfigStore) GetDefaultGroup() string { _ = "STUB: not implemented"; return "" }

func (n *ConfigStore) GetDefaultLabel() string { _ = "STUB: not implemented"; return "" }
