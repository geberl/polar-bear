package node

import (
	"sort"

	corev1 "k8s.io/api/core/v1"
)

func sortedContainerImages(imgs []corev1.ContainerImage) []corev1.ContainerImage {
	sorted := make([]corev1.ContainerImage, len(imgs))
	copy(sorted, imgs)

	sort.Slice(sorted, func(i, j int) bool {
		nameI := ""
		if len(sorted[i].Names) > 0 {
			nameI = sorted[i].Names[0]
		}
		nameJ := ""
		if len(sorted[j].Names) > 0 {
			nameJ = sorted[j].Names[0]
		}
		return nameI < nameJ
	})

	return sorted
}
