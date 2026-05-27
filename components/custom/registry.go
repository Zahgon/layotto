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
package custom

import (
	"mosn.io/layotto/components/pkg/info"
)

type Registry interface {
	Register(kind string, factories ...*Factory)
	Create(kind, compType string) (Component, error)
}

type Factory struct {
	Type          string
	FactoryMethod func() Component
}

func NewComponentFactory(compType string, f func() Component) *Factory {
	_ = "STUB: not implemented"
	return nil
}

type componentRegistry struct {
	stores map[string]map[string]func() Component
	info   *info.RuntimeInfo
}

func NewRegistry(info *info.RuntimeInfo) Registry { _ = "STUB: not implemented"; return *new(Registry) }

func (r *componentRegistry) Register(kind string, fs ...*Factory) {
	_ = "STUB: not implemented"
	return
}

// lazy init

// register FactoryMethod

func (r *componentRegistry) Create(kind, compType string) (Component, error) {
	_ = "STUB: not implemented"
	return *new(Component), nil
}
