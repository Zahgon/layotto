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

package apollo

import (
	"context"
	"net/http"
	"sync"

	"mosn.io/layotto/components/pkg/actuators"

	"mosn.io/layotto/components/configstores"

	log "mosn.io/layotto/kit/logger"
)

var (
	openAPIClientSingleton = &http.Client{}
	once                   sync.Once
	readinessIndicator     *actuators.HealthIndicator
	livenessIndicator      *actuators.HealthIndicator
)

const (
	defaultGroup  = "application"
	componentName = "apollo"
)

func init() {
	readinessIndicator = actuators.NewHealthIndicator()
	livenessIndicator = actuators.NewHealthIndicator()
}

type ConfigStore struct {
	tagsNamespace  string
	delimiter      string
	openAPIToken   string
	openAPIAddress string
	openAPIUser    string
	env            string
	listener       *changeListener
	kvRepo         Repository
	tagsRepo       Repository
	kvConfig       *repoConfig
	tagsConfig     *repoConfig
	openAPIClient  httpClient
	log            log.Logger
}
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type httpClientImpl struct {
	client *http.Client
}

func (c *httpClientImpl) Do(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConfigStore) GetDefaultGroup() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigStore) GetDefaultLabel() string { _ = "STUB: not implemented"; return "" }

func NewStore() configstores.Store { _ = "STUB: not implemented"; return *new(configstores.Store) }

func (c *ConfigStore) OnLogLevelChanged(outputLevel log.LogLevel) {
	_ = "STUB: not implemented"
	return
}

func registerActuator() { _ = "STUB: not implemented"; return }

func newHttpClient() httpClient { _ = "STUB: not implemented"; return *new(httpClient) }

// Init SetConfig the configuration store.
func (c *ConfigStore) Init(config *configstores.StoreConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigStore) doInit(config *configstores.StoreConfig) error {
	_ = "STUB: not implemented"
	// 1. validate and parse config
	return nil
}

// Metadata,required

// Address,required

// is_backup_config,not required
// whether backup config after fetch config from apollo

// app_id,required

// open_api_token,required

// open_api_address,not required

// open_api_user,required

// TODO make 'env' configurable
// 2. SetConfig client

// secret,not required

// 3. SetConfig client for tags query

// 4. SetConfig listener

func (c *ConfigStore) GetAppId() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigStore) GetStoreName() string { _ = "STUB: not implemented"; return "" }

// Get gets configuration from configuration store.
func (c *ConfigStore) Get(ctx context.Context, req *configstores.GetRequest) ([]*configstores.ConfigurationItem, error) {
	_ = "STUB: not implemented"
	// TODO forced pagination
	// 0. check if illegal
	return nil, nil
}

// 1. app level

// 2. group level

// 3. group+key+label level

func (c *ConfigStore) getAllTags(group string, keyWithLabel string) (tags map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil

	//	1. concatenate group+key+label
}

//	2. query

//	it means no tag

//	3. convert

// Set saves configuration into configuration store.
func (c *ConfigStore) Set(ctx context.Context, req *configstores.SetRequest) error {
	_ = "STUB: not implemented"
	// 1. check params
	return nil
}

// 2. loop set

// 2.1. set kv

// 2.2. set tags

// set tags value

// 3. commit tagsNamespace

// 4. commit kv namespace

// TODO 5. write cache

// Delete deletes configuration from configuration store.
func (c *ConfigStore) Delete(ctx context.Context, req *configstores.DeleteRequest) error {
	_ = "STUB: not implemented"
	// 1. check params
	return nil
}

// 2. loop delete

// 2.1. delete item

//	2.2. delete tags

// 3. commit tagsNamespace

// 4. commit kv namespace

// TODO 5. write cache

// Subscribe gets configuration from configuration store and subscribe the updates.
func (c *ConfigStore) Subscribe(req *configstores.SubscribeReq, ch chan *configstores.SubscribeResp) error {
	_ = "STUB: not implemented"
	// 0. check if illegal
	return nil
}

// 1. app level

// loop every namespace in config

// 2. group level

// 3. key level

func (c *ConfigStore) StopSubscribe() {
	_ = "STUB: not implemented"
	// TODO  Now the api layer only supports single connection and does not support multi-connection.
	//
	//	If it supports multiple connections in the future, we can use a context to cancel specific connections
	return
}

func (c *ConfigStore) getKeys(group string, keys []string, label string) ([]*configstores.ConfigurationItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. prepare suffix

// 2. loop query

//query value

//log error and ignore this key

// query tags

func (c *ConfigStore) getAllWithAppId() ([]*configstores.ConfigurationItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// loop every namespace in config

func (c *ConfigStore) getAllWithNamespace(group string) ([]*configstores.ConfigurationItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. loop query

// 1.1. convert

//	never happen

// 1.2. query tags

// 1.3. append result.

//continue

func (c *ConfigStore) setItem(appId string, item *configstores.ConfigurationItem) error {
	_ = "STUB: not implemented"
	// 1. put request
	return nil
}

// add body

// add params

// add headers

// do put request

// 2. parse

func (c *ConfigStore) addHeaderForOpenAPI(req *http.Request) {
	_ = "STUB: not implemented"
	// https://www.apolloconfig.com/#/zh/usage/apollo-open-api-platform?id=_3211-%e4%bf%ae%e6%94%b9%e9%85%8d%e7%bd%ae%e6%8e%a5%e5%8f%a3
	// Http Header中增加一个Authorization字段，字段值为申请的token
	// Http Header的Content-Type字段需要设置成application/json;charset=UTF-8
	return
}

func (c *ConfigStore) concatenateKey(key, label string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *ConfigStore) concatenateKeyForTag(group, keyWithLabel string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *ConfigStore) splitKey(keyWithLabel string) (key, label string) {
	_ = "STUB: not implemented"
	return "", ""
}

func (c *ConfigStore) commit(env string, appId string, cluster string, namespace string) error {
	_ = "STUB: not implemented"
	// 1. post request
	return nil
}

// add body

// add headers

// do request

// 2. parse

func (c *ConfigStore) deleteItem(env string, appId string, cluster string, group string, key string, label string) error {
	_ = "STUB: not implemented"
	// 1. delete request
	return nil
}

// add params

// add headers

// do request

// 2. parse

func (c *ConfigStore) initTagsClient(tagCfg *repoConfig) error {
	_ = "STUB: not implemented"
	// 1. create if not exist
	return nil
}

// 2. Connect

// refer to https://www.apolloconfig.com/#/zh/usage/apollo-open-api-platform?id=_327-%e5%88%9b%e5%bb%banamespace
func (c *ConfigStore) createNamespace(env string, appId string, cluster string, namespace string) error {
	_ = "STUB: not implemented"
	// 1. request
	return nil
}

// add body

// add headers

// do request

// 2. parse

// 4. commit

// if the namespace already exists, the status code will be 400

// log debug information

// Fail fast and take it as an startup error if the status code is neither 200 nor 400
