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

package patcher

import (
	"reflect"

	corev1 "k8s.io/api/core/v1"
)

// SidecarConfig contains the configuration for the sidecar container.
// Its parameters can be read from annotations on a pod.
type SidecarConfig struct {
	SidecarAPIGRPCPort int32 `default:"34904"`

	Namespace       string
	ImagePullPolicy corev1.PullPolicy

	SidecarInject  bool   `annotation:"layotto/sidecar-inject"`
	SidecarImage   string `annotation:"layotto/sidecar-image"`
	ConfigVolume   string `annotation:"layotto/config-volume"`
	VolumeMounts   string `annotation:"layotto/volume-mounts"`
	VolumeMountsRW string `annotation:"layotto/volume-mounts-rw"`

	pod *corev1.Pod
}

// NewSidecarConfig returns a ContainerConfig object for a pod.
func NewSidecarConfig(pod *corev1.Pod) *SidecarConfig { _ = "STUB: not implemented"; return nil }

func (c *SidecarConfig) setDefaultValues() {
	_ = "STUB: not implemented"
	// Iterate through the fields using reflection
	return
}

// Assign the default value

func (c *SidecarConfig) SetFromPodAnnotations() { _ = "STUB: not implemented"; return }

// setFromAnnotations updates the object with properties from an annotation map.
func (c *SidecarConfig) setFromAnnotations(an map[string]string) {
	_ = "STUB: not implemented"
	// Iterate through the fields using reflection
	return
}

// Skip annotations that are not defined or which have an empty value

// Assign the value

func setValueFromString(rt reflect.Type, rv reflect.Value, val string, key string) bool {
	_ = "STUB: not implemented"
	return false
}
