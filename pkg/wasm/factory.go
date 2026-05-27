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

package wasm

import (
	"context"

	"mosn.io/api"
	v2 "mosn.io/mosn/pkg/config/v2"
	"mosn.io/mosn/pkg/types"
)

const LayottoWasm = "Layotto"

func init() {
	api.RegisterStream(LayottoWasm, createProxyWasmFilterFactory)
}

// FilterConfigFactory contains multi wasm-plugin configs
// its pointer implement api.StreamFilterChainFactory
type FilterConfigFactory struct {
	LayottoHandler

	config        []*filterConfigItem // contains multi wasm config
	RootContextID int32

	// map[pluginName]*WasmPlugin
	plugins map[string]*WasmPlugin
	router  *Router
}

var factory = &FilterConfigFactory{
	config:        make([]*filterConfigItem, 0),
	RootContextID: 1,
	plugins:       make(map[string]*WasmPlugin),
	router:        &Router{routes: make(map[string]*Group)},
}

var _ api.StreamFilterChainFactory = &FilterConfigFactory{}

func GetFactory() *FilterConfigFactory {
	_ = "STUB: not implemented"

	// Create a proxy factory for WasmFilter
	return nil
}

func createProxyWasmFilterFactory(confs map[string]interface{}) (api.StreamFilterChainFactory, error) {
	_ = "STUB: not implemented"
	return *new(api.StreamFilterChainFactory), nil
}

// Create the FilterChain
var filterChain *Filter

func (f *FilterConfigFactory) CreateFilterChain(context context.Context, callbacks api.StreamFilterChainFactoryCallbacks) {
	_ = "STUB: not implemented"
	return
}

func (f *FilterConfigFactory) IsRegister(id string) bool { _ = "STUB: not implemented"; return false }

func (f *FilterConfigFactory) Install(conf map[string]interface{}, manager types.WasmManager) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FilterConfigFactory) UpdateInstanceNum(id string, instanceNum int, manager types.WasmManager) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FilterConfigFactory) UnInstall(id string, manager types.WasmManager) error {
	_ = "STUB: not implemented"
	return nil
}

// GetRootContextID Get RootContext's ID
func (f *FilterConfigFactory) GetRootContextID() int32 { _ = "STUB: not implemented"; return 0 }

// FilterConfigFactory implement types.WasmPluginHandler
// for `pw.RegisterPluginHandler(factory)`
var _ types.WasmPluginHandler = &FilterConfigFactory{}

// OnConfigUpdate Update config of FilterConfigFactory
func (f *FilterConfigFactory) OnConfigUpdate(config v2.WasmPluginConfig) {
	_ = "STUB: not implemented"
	return
}

// OnPluginStart Execute the plugin of FilterConfigFactory
func (f *FilterConfigFactory) OnPluginStart(plugin types.WasmPlugin) {
	_ = "STUB: not implemented"
	return
}

// get the ID of wasm, register route

// OnPluginDestroy Destroy the plugin of FilterConfigFactory
func (f *FilterConfigFactory) OnPluginDestroy(types.WasmPlugin) { _ = "STUB: not implemented"; return }
