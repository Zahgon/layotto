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
	"context"

	jsonpatch "github.com/evanphx/json-patch/v5"
	admissionv1 "k8s.io/api/admission/v1"
)

func (i *injector) getPodPatchOperations(ctx context.Context, ar *admissionv1.AdmissionReview) (patchOps jsonpatch.Patch, err error) {
	_ = "STUB: not implemented"
	return *new(jsonpatch.Patch), nil
}

// Create the sidecar configuration object from the pod

// Default value for the sidecar image, which can be overridden by annotations

// Set the configuration from annotations

// Get the patch to apply to the pod
// Patch may be empty if there's nothing that needs to be done
