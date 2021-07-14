package annotations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	apiv1 "k8s.io/api/core/v1"
)

func TestHasRecommendationAnnotations(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		want        bool
	}{
		{
			name: "empty annotations should return false",
			want: false,
		},
		{
			name: "no annotations should return false",
			annotations: map[string]string{
				"annotation1":           "somevalue",
				"other-annotation/type": "othervalue",
			},
			want: false,
		},
		{
			name: "one annotation should return true",
			annotations: map[string]string{
				"container.recommendations.com/cpu": "100m",
			},
			want: true,
		},
		{
			name: "multiple annotations should return true",
			annotations: map[string]string{
				"container1.recommendations.com/cpu":    "100m",
				"container1.recommendations.com/memory": "100000000",
				"container2.recommendations.com/cpu":    "2",
				"container3.recommendations.com/memory": "60Gi",
			},
			want: true,
		},
		{
			name: "mixed annotation should return true",
			annotations: map[string]string{
				"container1.recommendations.com/cpu": "100m",
				"annotation1":                        "somevalue",
			},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := HasRecommendationAnnotations(tc.annotations)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestGetResourceRequestAnnotations(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		want        map[string]map[apiv1.ResourceName]string
	}{
		{
			name: "empty annotation should return empty map",
			want: map[string]map[apiv1.ResourceName]string{},
		},
		{
			name: "no annotation should return empty map",
			annotations: map[string]string{
				"annotation1":           "somevalue",
				"other-annotation/type": "othervalue",
			},
			want: map[string]map[apiv1.ResourceName]string{},
		},
		{
			name: "annotation should return non-empty map",
			annotations: map[string]string{
				"container1.recommendations.com/cpu":    "100m",
				"container1.recommendations.com/memory": "100000000",
			},
			want: map[string]map[apiv1.ResourceName]string{
				"container1": {
					apiv1.ResourceCPU:    "100m",
					apiv1.ResourceMemory: "100000000",
				},
			},
		},
		{
			name: "annotations with multiple containers should return non-empty map",
			annotations: map[string]string{
				"container1.recommendations.com/cpu":    "100m",
				"container1.recommendations.com/memory": "100000000",
				"container2.recommendations.com/cpu":    "2",
				"container3.recommendations.com/memory": "60Gi",
			},
			want: map[string]map[apiv1.ResourceName]string{
				"container1": {
					apiv1.ResourceCPU:    "100m",
					apiv1.ResourceMemory: "100000000",
				},
				"container2": {
					apiv1.ResourceCPU: "2",
				},
				"container3": {
					apiv1.ResourceMemory: "60Gi",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetResourceRequestAnnotations(tc.annotations)
			assert.Equal(t, tc.want, got)
		})
	}
}
