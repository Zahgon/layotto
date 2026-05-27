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

package service

import (
	"net/http"

	"k8s.io/apimachinery/pkg/runtime"
)

const (
	port = 8443
)

// Injector is the interface for the layotto runtime sidecar injection component.
type Injector interface {
	Run() error
}

type injector struct {
	config       Config
	deserializer runtime.Decoder
	server       *http.Server
}

// Run implements Injector.
func (i *injector) Run() error { _ = "STUB: not implemented"; return nil }

// NewInjector returns a new Injector instance.
func NewInjector(config Config) (Injector, error) {
	_ = "STUB: not implemented"
	return *new(Injector), nil
}
