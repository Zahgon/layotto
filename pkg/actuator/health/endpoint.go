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

package health

import (
	"context"
	"errors"

	"mosn.io/layotto/pkg/actuator"
	"mosn.io/layotto/pkg/filter/stream/common/http"
)

const (
	health_key     = "health"
	liveness_key   = "liveness"
	readiness_key  = "readiness"
	status_key     = "status"
	components_key = "components"
)

var (
	invalidTypeError = errors.New("health type invalid")
	serviceDownError = errors.New("service unavailable")
	serviceInitError = errors.New("service is initializing")
)

// init health Endpoint.
func init() {
	actuator.GetDefault().AddEndpoint(health_key, NewEndpoint())
}

var type2Indicators = make(map[string]map[string]Indicator)

type Endpoint struct {
}

func NewEndpoint() *Endpoint {
	_ = "STUB: not implemented"

	// Handle will check health status.The structure of the returned map is like:
	//
	//	{
	//	 "status": "DOWN",
	//	 "components": {
	//	   "readinessProbe": {
	//	     "status": "DOWN",
	//	     "details": {}
	//	   }
	//	 }
	//	}
	return nil
}

func (e *Endpoint) Handle(ctx context.Context, params http.ParamsScanner) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. validate params

// 2. traverse the indicator chain

// AddLivenessIndicator register health.Indicator for liveness check.It's not concurrent-safe,so please invoke it ONLY in init method.
func AddLivenessIndicator(name string, idc Indicator) { _ = "STUB: not implemented"; return }

// AddLivenessIndicatorFunc register health.Indicator for liveness check.Indicator.It's not concurrent-safe,so please invoke it ONLY in init method.
func AddLivenessIndicatorFunc(name string, f func() (string, map[string]interface{})) {
	_ = "STUB: not implemented"
	return
}

// AddReadinessIndicator register health.Indicator for readiness check.Indicator.It's not concurrent-safe,so please invoke it ONLY in init method.
func AddReadinessIndicator(name string, idc Indicator) { _ = "STUB: not implemented"; return }

// AddReadinessIndicatorFunc register health.Indicator for readiness check.Indicator.It's not concurrent-safe,so please invoke it ONLY in init method.
func AddReadinessIndicatorFunc(name string, f func() (string, map[string]interface{})) {
	_ = "STUB: not implemented"
	return
}

func addIndicator(indicatorType string, name string, idc Indicator) {
	_ = "STUB: not implemented"
	return
}
