/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resources

import (
	corev1 "k8s.io/api/core/v1"
)

var (
	pvcDir = "/pvc"
)

func getPvcMount(name string) corev1.VolumeMount {
	return corev1.VolumeMount{
		Name:      name,   // taskrun pvc name
		MountPath: pvcDir, // nothing should be mounted here
	}
}

// GetPVCVolume gets pipelinerun pvc volume
func GetPVCVolume(name string) corev1.Volume {
	return corev1.Volume{
		Name: name,
		VolumeSource: corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{ClaimName: name},
		},
	}
}
