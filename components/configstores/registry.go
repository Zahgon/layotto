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

package configstores

import (
	"mosn.io/layotto/components/pkg/info"
)

type Registry interface {
	Register(fs ...*StoreFactory)
	Create(compType string) (Store, error)
}

type StoreFactory struct {
	CompType      string
	FactoryMethod func() Store
}

func NewStoreFactory(compType string, f func() Store) *StoreFactory {
	_ = "STUB: not implemented"
	return nil
}

type StoreRegistry struct {
	stores map[string]func() Store
	info   *info.RuntimeInfo
}

func NewRegistry(info *info.RuntimeInfo) Registry { _ = "STUB: not implemented"; return *new(Registry) }

func (r *StoreRegistry) Register(fs ...*StoreFactory) { _ = "STUB: not implemented"; return }

func (r *StoreRegistry) Create(compType string) (Store, error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}
