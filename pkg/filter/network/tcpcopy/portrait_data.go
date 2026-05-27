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

package tcpcopy

import (
	"context"
	"sync"

	_type "mosn.io/layotto/pkg/filter/network/tcpcopy/type"
)

var lock sync.Mutex

func isHandle(businessType _type.BusinessType) bool {
	_ = "STUB: not implemented"
	// Determine whether to continue sampling
	return false
}

// The same business type, in the same sampling period, only accept one data report

func getAndSwapDumpBusinessCache(businessType _type.BusinessType, new int) int {
	_ = "STUB: not implemented"
	return 0
}

// 默认为0

// Upload portrait data
func UploadPortraitData(businessType _type.BusinessType, data interface{}, ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Persistent user reported data
