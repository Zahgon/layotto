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
	v2 "mosn.io/mosn/pkg/config/v2"
)

type filterConfigItem struct {
	FromWasmPlugin string            `json:"from_wasm_plugin,omitempty"`
	VmConfig       *v2.WasmVmConfig  `json:"vm_config,omitempty"`
	InstanceNum    int               `json:"instance_num,omitempty"`
	RootContextID  int32             `json:"root_context_id,omitempty"`
	UserData       map[string]string `json:"-"`
	PluginName     string            `json:"-"`
}

// Parse filterConfigItem
func parseFilterConfigItem(cfg map[string]interface{}) (*filterConfigItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default value is 1

// Check VMconfig of filterConfigItem
func checkVmConfig(config *filterConfigItem) error { _ = "STUB: not implemented"; return nil }

// Parse user data
func parseUserData(rawConfigBytes []byte, config *filterConfigItem) error {
	_ = "STUB: not implemented"
	return nil
}

// check rawConfigBytes

// delete all pairs that value type is not string

// add into config
