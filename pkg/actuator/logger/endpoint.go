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

package info

import (
	"context"

	"mosn.io/layotto/pkg/actuator"
	"mosn.io/layotto/pkg/filter/stream/common/http"
)

// init info Endpoint.
func init() {
	actuator.GetDefault().AddEndpoint("logger", NewEndpoint())
}

type Endpoint struct {
}

type LoggerLevelChangedRequest struct {
	Component string `json:"component"`
	Level     string `json:"level"`
}

func NewEndpoint() *Endpoint { _ = "STUB: not implemented"; return nil }

func (e *Endpoint) Handle(ctx context.Context, params http.ParamsScanner) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handle the infoContributors
