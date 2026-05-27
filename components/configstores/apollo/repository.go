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
	"github.com/apolloconfig/agollo/v4"
	agolloConfig "github.com/apolloconfig/agollo/v4/env/config"

	"mosn.io/layotto/kit/logger"
)

// An interface to abstract different apollo sdks,also making it easier to write unit tests.
type Repository interface {
	SetConfig(r *repoConfig)
	Connect() error
	// subscribe
	AddChangeListener(listener *changeListener)
	// query
	Get(namespace string, key string) (interface{}, error)
	//	process every items under the namespace
	Range(namespace string, f func(key, value interface{}) bool) error
}

type repoConfig struct {
	addr          string
	appId         string
	storeName     string
	env           string
	cluster       string
	namespaceName string
	// whether backup config after fetch config from apollo
	isBackupConfig bool
	secret         string
	logger         logger.Logger
}

// Implement Repository interface
type AgolloRepository struct {
	client agollo.Client
	cfg    *repoConfig
}

func (a *AgolloRepository) Connect() error { _ = "STUB: not implemented"; return nil }

func (a *AgolloRepository) SetConfig(r *repoConfig) { _ = "STUB: not implemented"; return }

func repoConfig2AgolloConfig(r *repoConfig) *agolloConfig.AppConfig {
	_ = "STUB: not implemented"
	return nil
}

func newAgolloRepository() Repository { _ = "STUB: not implemented"; return *new(Repository) }

func (a *AgolloRepository) Get(namespace string, key string) (interface{}, error) {
	_ = "STUB: not implemented"
	// 1. get cache
	return nil, nil
}

// 2. query value

func (a *AgolloRepository) Range(namespace string, f func(key interface{}, value interface{}) bool) error {
	_ = "STUB: not implemented"
	// 1. get cache
	return nil
}

// 2. loop process

func (a *AgolloRepository) AddChangeListener(listener *changeListener) {
	_ = "STUB: not implemented"
	return
}
