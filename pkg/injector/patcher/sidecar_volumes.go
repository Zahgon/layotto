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
	corev1 "k8s.io/api/core/v1"
)

// getVolumeMounts returns the list of VolumeMount's for the sidecar container.
func (c *SidecarConfig) getVolumeMounts() []corev1.VolumeMount {
	_ = "STUB: not implemented"
	return nil
}

func podContainsVolume(pod *corev1.Pod, name string) bool { _ = "STUB: not implemented"; return false }

// parseVolumeMountsString parses the annotation and returns volume mounts.
// The format of the annotation is: "mountPath1:hostPath1,mountPath2:hostPath2"
// The readOnly parameter applies to all mounts.
func parseVolumeMountsString(volumeMountStr string, readOnly bool) []corev1.VolumeMount {
	_ = "STUB: not implemented"
	return nil
}

// getConfigVolumeMount returns the layotto config volume mount.
// Currently, the path of the Layotto configuration file is "/runtime/configs/config.json"
func getConfigVolumeMount(configVolumeName string, readOnly bool) corev1.VolumeMount {
	_ = "STUB: not implemented"
	return *new(corev1.VolumeMount)
}
