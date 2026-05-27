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

package persistence

import (
	"math/rand"
	"sync"

	"mosn.io/layotto/pkg/filter/network/tcpcopy/model"
)

func init() {
	dumpWorkPoolInstance = NewDefaultWorkPool(20)
}

var dumpWorkPoolInstance *DefaultWorkPool

type WorkGoroutine struct {
	tasks *sync.Map
}

func NewWorkGoroutine() *WorkGoroutine { _ = "STUB: not implemented"; return nil }

func (g *WorkGoroutine) AddTask(key string, data *model.DumpUploadDynamicConfig) {
	_ = "STUB: not implemented"
	return
}

func (g *WorkGoroutine) Start() { _ = "STUB: not implemented"; return }

func (g *WorkGoroutine) work() { _ = "STUB: not implemented"; return }

type DefaultWorkPool struct {
	size           int64
	workers        *sync.Map
	randomInstance *rand.Rand
	lock           *sync.Mutex
}

func NewDefaultWorkPool(size int64) *DefaultWorkPool { _ = "STUB: not implemented"; return nil }

func GetDumpWorkPoolInstance() *DefaultWorkPool { _ = "STUB: not implemented"; return nil }

func (w *DefaultWorkPool) random() int64 { _ = "STUB: not implemented"; return 0 }

func (w *DefaultWorkPool) Schedule(data *model.DumpUploadDynamicConfig) {
	_ = "STUB: not implemented"
	return
}
