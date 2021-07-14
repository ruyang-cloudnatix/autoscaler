package annotations

import (
	"regexp"

	apiv1 "k8s.io/api/core/v1"
)

const (
	// recommendCPU is the annotation key for CPU recommendations.
	recommendCPU = "%s.recommendations.com/cpu"
	// cnatixMem is the annotation key for memory recommendations.
	recommendMem = "%s.recommendations.com/memory"
	// recommendRegex is the regex for matching recommend annotations.
	// The first part is to match container names, but the expression is
	// less strict than allowable container names.
	recommendRegex = `([a-z0-9-]+).recommendations.com/(cpu|memory)$`
)

// GetResourceRequestAnnotations returns the corresponding resource requests for
// `%s.recommendation.com/cpu` or `%s.recommendation.com/memory` or both.
func GetResourceRequestAnnotations(annotations map[string]string) map[string]map[apiv1.ResourceName]string {
	re := regexp.MustCompile(recommendRegex)
	requested := make(map[string]map[apiv1.ResourceName]string, 0)

	for key, value := range annotations {
		matches := re.FindStringSubmatch(key)
		if matches == nil {
			continue
		}

		container := matches[1]
		resourceName := matches[2]
		if _, ok := requested[container]; !ok {
			requested[container] = map[apiv1.ResourceName]string{}
		}
		requested[container][apiv1.ResourceName(resourceName)] = value
	}
	return requested
}

// HasRecommendationAnnotations checks if the recommendation annotations are set.
func HasRecommendationAnnotations(annotations map[string]string) bool {
	re := regexp.MustCompile(recommendRegex)
	for key := range annotations {
		if re.MatchString(key) {
			return true
		}
	}
	return false
}
