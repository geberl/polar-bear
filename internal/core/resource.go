package core

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"reflect"
	"sort"
	"strings"

	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	batchv1 "k8s.io/api/batch/v1"
	coordinationv1 "k8s.io/api/coordination/v1"
	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	eventsv1 "k8s.io/api/events/v1"
	networkingv1 "k8s.io/api/networking/v1"
	nodev1 "k8s.io/api/node/v1"
	policyv1 "k8s.io/api/policy/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	schedulingv1 "k8s.io/api/scheduling/v1"
	storagev1 "k8s.io/api/storage/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"polar-bear/internal/store"
)

type KubernetesResource interface {
	// Non-namespaced resources
	*corev1.Namespace |
		*corev1.Node |
		*corev1.PersistentVolume |
		*storagev1.StorageClass |
		*rbacv1.ClusterRole |
		*rbacv1.ClusterRoleBinding |
		*apiextensionsv1.CustomResourceDefinition |
		*admissionregistrationv1.MutatingWebhookConfiguration |
		*admissionregistrationv1.ValidatingWebhookConfiguration |
		*schedulingv1.PriorityClass |
		*nodev1.RuntimeClass |
		*storagev1.VolumeAttachment |
		*storagev1.CSIDriver |
		*storagev1.CSINode |
		*storagev1.CSIStorageCapacity |

		// Namespaced resources - Workloads
		*corev1.Pod |
		*appsv1.Deployment |
		*appsv1.StatefulSet |
		*appsv1.DaemonSet |
		*appsv1.ReplicaSet |
		*batchv1.Job |
		*batchv1.CronJob |
		*corev1.ReplicationController |

		// Namespaced resources - Services and Networking
		*corev1.Service |
		*discoveryv1.EndpointSlice |
		*networkingv1.Ingress |
		*networkingv1.NetworkPolicy |

		// Namespaced resources - Configuration and Storage
		*corev1.ConfigMap |
		*corev1.Secret |
		*corev1.PersistentVolumeClaim |

		// Namespaced resources - Authorization
		*corev1.ServiceAccount |
		*rbacv1.Role |
		*rbacv1.RoleBinding |

		// Namespaced resources - Autoscaling and Policy
		*autoscalingv2.HorizontalPodAutoscaler |
		*policyv1.PodDisruptionBudget |
		*corev1.ResourceQuota |
		*corev1.LimitRange |

		// Namespaced resources - Events and Coordination
		*eventsv1.Event |
		*coordinationv1.Lease |

		// Namespaced resources - API Machinery
		*appsv1.ControllerRevision |
		*corev1.PodTemplate
}

func getResourceKind[T KubernetesResource]() string {
	var zero T
	t := reflect.TypeOf(zero)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return strings.ToLower(t.Name())
}

func GetResource[T KubernetesResource](
	store store.Store,
	ns string,
	name string,
) T {
	resourceKind := getResourceKind[T]()
	logger := slog.With("component", fmt.Sprintf("core-%s", resourceKind))

	dbKey, err := ResourceKey(resourceKind, ns, name)
	if err != nil {
		logger.Error(
			"unable to get resource key",
			"kind", resourceKind,
			"namespace", ns,
			"name", name,
			"error", err,
		)
		var zero T
		return zero
	}

	keyVal, err := store.Get(dbKey)
	if err != nil {
		logger.Error(
			"error on get from db",
			"key", string(dbKey),
			"error", err,
		)
		var zero T
		return zero
	}

	var res T
	err = json.Unmarshal(keyVal, &res)
	if err != nil {
		logger.Error(
			"error on unmarshal from json",
			"namespace", ns,
			"name", name,
			"error", err,
		)
	}

	return res
}

func GetResources[T KubernetesResource](
	store store.Store,
	ns string,
) []T {
	resourceKind := getResourceKind[T]()
	logger := slog.With("component", fmt.Sprintf("core-%s", resourceKind))

	dbKey, err := ResourceKey(resourceKind, ns, "")
	if err != nil {
		logger.Error(
			"unable to get resource key",
			"kind", resourceKind,
			"namespace", ns,
			"error", err,
		)
		return nil
	}

	keyVals, err := store.GetAll(dbKey)
	if err != nil {
		logger.Error(
			"error on get all from db",
			"key", string(dbKey),
			"error", err,
		)
		return nil
	}

	names := make([]string, 0, len(keyVals))
	for key := range keyVals {
		names = append(names, key)
	}
	sort.Strings(names)

	resources := make([]T, 0, len(names))
	for _, name := range names {
		var resource T
		err := json.Unmarshal(keyVals[name], &resource)
		if err != nil {
			logger.Error(
				"error on unmarshal from json",
				"namespace", ns,
				"name", name,
				"error", err,
			)
			continue
		}
		resources = append(resources, resource)
	}

	return resources
}

func CountResources[T KubernetesResource](
	store store.Store,
	ns string,
) uint {
	resourceKind := getResourceKind[T]()
	logger := slog.With("component", fmt.Sprintf("core-%s", resourceKind))

	dbKey, err := ResourceKey(resourceKind, ns, "")
	if err != nil {
		logger.Error(
			"unable to get resource key",
			"kind", resourceKind,
			"namespace", ns,
			"error", err,
		)
		return uint(0)
	}

	count, err := store.Count(dbKey)
	if err != nil {
		logger.Error(
			"error on get from db",
			"key", string(dbKey),
			"error", err,
		)
		return uint(0)
	}

	return count
}
