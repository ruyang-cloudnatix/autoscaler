package annotations

import (
	apiv1 "k8s.io/api/core/v1"
	vpa_types "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
)

const (
	// CNatixCPU is the annotation key for CloudNatix CPU recommendations.
	CNatixCPU = "recommendations.cloudnatix.com/cpu"
	// CNatixMEM is the annotation key for CloudNatix Memory recommendations.
	CNatixMEM = "recommendations.cloudnatix.com/memory"
)

// IsAnnotationsAuto returns if the annotated values for CPU and Memory are set in `Auto` mode.
// If `Auto` mode is set, the behaviour of VPA performs the same as the default implementation.
func IsAnnotationsAuto(annotations map[string]string) bool {
	annotationValues := GetResourceRequestAnnotations(annotations)
	cpuValue, hasCPU := annotationValues[apiv1.ResourceCPU]
	memValue, hasMEM := annotationValues[apiv1.ResourceMemory]

	return hasCPU && hasMEM &&
		cpuValue == string(vpa_types.UpdateModeAuto) &&
		memValue == string(vpa_types.UpdateModeAuto)
}

// GetResourceRequestAnnotations returns the corresponding resource requests for
// `cloudnatix.com/recommendation/cpu` or `cloudnatix.com/recommendation/memory` or both and.
func GetResourceRequestAnnotations(annotations map[string]string) map[apiv1.ResourceName]string {
	requests := make(map[apiv1.ResourceName]string)

	requests[apiv1.ResourceCPU] = annotations[CNatixCPU]
	requests[apiv1.ResourceMemory] = annotations[CNatixMEM]
	return requests
}

// HasCloudnatixAnnotations checks if the Cloudnatix annotations are set.
func HasCloudnatixAnnotations(annotations map[string]string) bool {
	_, CPUok := annotations[CNatixCPU]
	_, MEMok := annotations[CNatixMEM]
	return CPUok || MEMok
}
