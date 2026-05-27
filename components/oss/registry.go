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

package oss

import (
	"mosn.io/layotto/components/pkg/info"
)

type Registry interface {
	Register(fs ...*Factory)
	Create(compType string) (Oss, error)
}

type Factory struct {
	CompType      string
	FactoryMethod func() Oss
}

func NewFactory(compType string, f func() Oss) *Factory { _ = "STUB: not implemented"; return nil }

type registry struct {
	oss  map[string]func() Oss
	info *info.RuntimeInfo
}

func NewRegistry(info *info.RuntimeInfo) Registry { _ = "STUB: not implemented"; return *new(Registry) }

func (r *registry) Register(fs ...*Factory) { _ = "STUB: not implemented"; return }

func (r *registry) Create(compType string) (Oss, error) {
	_ = "STUB: not implemented"
	return *new(Oss), nil
}
