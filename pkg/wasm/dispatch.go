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
package wasm

type Group struct {
	count   int
	plugins []*WasmPlugin
}

type Router struct {
	routes map[string]*Group
}

// RegisterRoute register a group with id
// unsafe for concurrent
func (route *Router) RegisterRoute(id string, plugin *WasmPlugin) {
	_ = "STUB: not implemented"
	return
}

// RemoveRoute remove group by id
func (route *Router) RemoveRoute(id string) { _ = "STUB: not implemented"; return }

// GetRandomPluginByID Get random plugin with rand id
func (route *Router) GetRandomPluginByID(id string) (*WasmPlugin, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
