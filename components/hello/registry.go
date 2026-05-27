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

package hello

import (
	"mosn.io/layotto/components/pkg/info"
)

type Registry interface {
	Register(fs ...*HelloFactory)
	Create(compType string) (HelloService, error)
}

type HelloFactory struct {
	CompType      string
	FactoryMethod func() HelloService
}

func NewHelloFactory(compType string, f func() HelloService) *HelloFactory {
	_ = "STUB: not implemented"
	return nil
}

type helloRegistry struct {
	stores map[string]func() HelloService
	info   *info.RuntimeInfo
}

func NewRegistry(info *info.RuntimeInfo) Registry {
	_ = "STUB: not implemented"
	return *
	// 添加服务信息
	new(Registry)
}

func (r *helloRegistry) Register(fs ...*HelloFactory) { _ = "STUB: not implemented"; return }

// 注册组件信息

func (r *helloRegistry) Create(compType string) (HelloService, error) {
	_ = "STUB: not implemented"
	return *new(HelloService), nil
}

// 加载组件信息
