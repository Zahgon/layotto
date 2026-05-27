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

package bindings

import (
	"github.com/dapr/components-contrib/bindings"

	"mosn.io/layotto/components/pkg/info"
)

const (
	ServiceName = "bindings"
)

type Registry interface {
	RegisterOutputBinding(fs ...*OutputBindingFactory)
	RegisterInputBinding(fs ...*InputBindingFactory)
	CreateOutputBinding(compType string) (bindings.OutputBinding, error)
	CreateInputBinding(compType string) (bindings.InputBinding, error)
}

type bindingsRegistry struct {
	outputBindingStores map[string]func() bindings.OutputBinding
	inputBindingStores  map[string]func() bindings.InputBinding
	info                *info.RuntimeInfo
}

func NewRegistry(info *info.RuntimeInfo) Registry { _ = "STUB: not implemented"; return *new(Registry) }

func (r *bindingsRegistry) RegisterOutputBinding(fs ...*OutputBindingFactory) {
	_ = "STUB: not implemented"
	return
}

func (r *bindingsRegistry) RegisterInputBinding(fs ...*InputBindingFactory) {
	_ = "STUB: not implemented"
	return
}

func (r *bindingsRegistry) CreateOutputBinding(compType string) (bindings.OutputBinding, error) {
	_ = "STUB: not implemented"
	return *new(bindings.OutputBinding), nil
}

func (r *bindingsRegistry) CreateInputBinding(compType string) (bindings.InputBinding, error) {
	_ = "STUB: not implemented"
	return *new(bindings.InputBinding), nil
}
