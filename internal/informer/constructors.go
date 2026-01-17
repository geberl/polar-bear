package informer

import (
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
	"k8s.io/client-go/informers"
	k8scache "k8s.io/client-go/tools/cache"

	"polar-bear/internal/event"
	"polar-bear/internal/store"
)

func NewTypedInformer[T any](
	informer k8scache.SharedIndexInformer,
	store store.Store,
	ed event.Distribution,
	kind string,
	getNamespace func(obj T) string,
	getName func(obj T) string,
) *ResourceInformer[T] {
	return NewResourceInformer(
		make(chan struct{}),
		informer,
		store,
		ed,
		kind,
		getNamespace,
		getName,
	)
}

// =============================================================================
// NON-NAMESPACED RESOURCES
// =============================================================================

func NewNamespaceInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*corev1.Namespace] {
	return NewTypedInformer(
		factory.Core().V1().Namespaces().Informer(),
		store,
		ed,
		"namespace",
		func(_ *corev1.Namespace) string { return "" },
		func(ns *corev1.Namespace) string { return ns.Name },
	)
}

func NewNodeInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*corev1.Node] {
	return NewTypedInformer(
		factory.Core().V1().Nodes().Informer(),
		store,
		ed,
		"node",
		func(_ *corev1.Node) string { return "" },
		func(no *corev1.Node) string { return no.Name },
	)
}

func NewPersistentVolumeInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*corev1.PersistentVolume] {
	return NewTypedInformer(
		factory.Core().V1().PersistentVolumes().Informer(),
		store,
		ed,
		"persistentvolume",
		func(_ *corev1.PersistentVolume) string { return "" },
		func(pv *corev1.PersistentVolume) string { return pv.Name },
	)
}

func NewStorageClassInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*storagev1.StorageClass] {
	return NewTypedInformer(
		factory.Storage().V1().StorageClasses().Informer(),
		store,
		ed,
		"storageclass",
		func(_ *storagev1.StorageClass) string { return "" },
		func(sc *storagev1.StorageClass) string { return sc.Name },
	)
}

func NewClusterRoleInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*rbacv1.ClusterRole] {
	return NewTypedInformer(
		factory.Rbac().V1().ClusterRoles().Informer(),
		store,
		ed,
		"clusterrole",
		func(_ *rbacv1.ClusterRole) string { return "" },
		func(cr *rbacv1.ClusterRole) string { return cr.Name },
	)
}

func NewClusterRoleBindingInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*rbacv1.ClusterRoleBinding] {
	return NewTypedInformer(
		factory.Rbac().V1().ClusterRoleBindings().Informer(),
		store,
		ed,
		"clusterrolebinding",
		func(_ *rbacv1.ClusterRoleBinding) string { return "" },
		func(crb *rbacv1.ClusterRoleBinding) string { return crb.Name },
	)
}

// This doesn't work, the CRD informer must be handled separately from the main SharedInformerFactory.
// https://chatgpt.com/share/6868f95a-946c-8000-a5cb-61f643958f0d

// func NewCustomResourceDefinitionInformer(
// 	factory informers.SharedInformerFactory,
// 	store store.Store,
// 	ed event.Distribution,
// ) *ResourceInformer[*apiextensionsv1.CustomResourceDefinition] {
// 	return NewTypedInformer(
// 		factory.Apiextensions().V1().CustomResourceDefinitions().Informer(),
// 		store,
// 		ed,
// 		"crd",
// 		func(_ *apiextensionsv1.CustomResourceDefinition) string { return "" },
// 		func(crd *apiextensionsv1.CustomResourceDefinition) string { return crd.Name },
// 	)
// }

func NewMutatingWebhookConfigurationInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*admissionregistrationv1.MutatingWebhookConfiguration] {
	return NewTypedInformer(
		factory.Admissionregistration().V1().MutatingWebhookConfigurations().Informer(),
		store,
		ed,
		"mutatingwebhookconfiguration",
		func(_ *admissionregistrationv1.MutatingWebhookConfiguration) string { return "" },
		func(mwc *admissionregistrationv1.MutatingWebhookConfiguration) string { return mwc.Name },
	)
}

func NewValidatingWebhookConfigurationInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*admissionregistrationv1.ValidatingWebhookConfiguration] {
	return NewTypedInformer(
		factory.Admissionregistration().V1().ValidatingWebhookConfigurations().Informer(),
		store,
		ed,
		"validatingwebhookconfiguration",
		func(_ *admissionregistrationv1.ValidatingWebhookConfiguration) string { return "" },
		func(vwc *admissionregistrationv1.ValidatingWebhookConfiguration) string { return vwc.Name },
	)
}

func NewPriorityClassInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*schedulingv1.PriorityClass] {
	return NewTypedInformer(
		factory.Scheduling().V1().PriorityClasses().Informer(),
		store,
		ed,
		"priorityclass",
		func(_ *schedulingv1.PriorityClass) string { return "" },
		func(pc *schedulingv1.PriorityClass) string { return pc.Name },
	)
}

func NewRuntimeClassInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*nodev1.RuntimeClass] {
	return NewTypedInformer(
		factory.Node().V1().RuntimeClasses().Informer(),
		store,
		ed,
		"runtimeclass",
		func(_ *nodev1.RuntimeClass) string { return "" },
		func(rc *nodev1.RuntimeClass) string { return rc.Name },
	)
}

func NewVolumeAttachmentInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*storagev1.VolumeAttachment] {
	return NewTypedInformer(
		factory.Storage().V1().VolumeAttachments().Informer(),
		store,
		ed,
		"volumeattachment",
		func(_ *storagev1.VolumeAttachment) string { return "" },
		func(va *storagev1.VolumeAttachment) string { return va.Name },
	)
}

func NewCSIDriverInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*storagev1.CSIDriver] {
	return NewTypedInformer(
		factory.Storage().V1().CSIDrivers().Informer(),
		store,
		ed,
		"csidriver",
		func(_ *storagev1.CSIDriver) string { return "" },
		func(csi *storagev1.CSIDriver) string { return csi.Name },
	)
}

func NewCSINodeInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*storagev1.CSINode] {
	return NewTypedInformer(
		factory.Storage().V1().CSINodes().Informer(),
		store,
		ed,
		"csinode",
		func(_ *storagev1.CSINode) string { return "" },
		func(cn *storagev1.CSINode) string { return cn.Name },
	)
}

func NewCSIStorageCapacityInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*storagev1.CSIStorageCapacity] {
	return NewTypedInformer(
		factory.Storage().V1().CSIStorageCapacities().Informer(),
		store,
		ed,
		"csistoragecapacity",
		func(csc *storagev1.CSIStorageCapacity) string {
			if csc.NodeTopology != nil {
				return csc.NodeTopology.String()
			}
			return ""
		},
		func(csc *storagev1.CSIStorageCapacity) string { return csc.Name },
	)
}

// =============================================================================
// NAMESPACED RESOURCES
// =============================================================================

func NewPodInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*corev1.Pod] {
	return NewTypedInformer(
		factory.Core().V1().Pods().Informer(),
		store,
		ed,
		"pod",
		func(p *corev1.Pod) string { return p.Namespace },
		func(p *corev1.Pod) string { return p.Name },
	)
}

func NewDeploymentInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*appsv1.Deployment] {
	return NewTypedInformer(
		factory.Apps().V1().Deployments().Informer(),
		store,
		ed,
		"deployment",
		func(d *appsv1.Deployment) string { return d.Namespace },
		func(d *appsv1.Deployment) string { return d.Name },
	)
}

func NewStatefulSetInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*appsv1.StatefulSet] {
	return NewTypedInformer(
		factory.Apps().V1().StatefulSets().Informer(),
		store,
		ed,
		"statefulset",
		func(ss *appsv1.StatefulSet) string { return ss.Namespace },
		func(ss *appsv1.StatefulSet) string { return ss.Name },
	)
}

func NewDaemonSetInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*appsv1.DaemonSet] {
	return NewTypedInformer(
		factory.Apps().V1().DaemonSets().Informer(),
		store,
		ed,
		"daemonset",
		func(obj *appsv1.DaemonSet) string { return obj.Namespace },
		func(obj *appsv1.DaemonSet) string { return obj.Name },
	)
}

func NewReplicaSetInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*appsv1.ReplicaSet] {
	return NewTypedInformer(
		factory.Apps().V1().ReplicaSets().Informer(),
		store,
		ed,
		"replicaset",
		func(obj *appsv1.ReplicaSet) string { return obj.Namespace },
		func(obj *appsv1.ReplicaSet) string { return obj.Name },
	)
}

func NewJobInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*batchv1.Job] {
	return NewTypedInformer(
		factory.Batch().V1().Jobs().Informer(),
		store,
		ed,
		"job",
		func(obj *batchv1.Job) string { return obj.Namespace },
		func(obj *batchv1.Job) string { return obj.Name },
	)
}

func NewCronJobInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*batchv1.CronJob] {
	return NewTypedInformer(
		factory.Batch().V1().CronJobs().Informer(),
		store,
		ed,
		"cronjob",
		func(obj *batchv1.CronJob) string { return obj.Namespace },
		func(obj *batchv1.CronJob) string { return obj.Name },
	)
}

func NewReplicationControllerInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*corev1.ReplicationController] {
	return NewTypedInformer(
		factory.Core().V1().ReplicationControllers().Informer(),
		store,
		ed,
		"replicationcontroller",
		func(obj *corev1.ReplicationController) string { return obj.Namespace },
		func(obj *corev1.ReplicationController) string { return obj.Name },
	)
}

func NewServiceInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*corev1.Service] {
	return NewTypedInformer(
		factory.Core().V1().Services().Informer(),
		store,
		ed,
		"service",
		func(obj *corev1.Service) string { return obj.Namespace },
		func(obj *corev1.Service) string { return obj.Name },
	)
}

func NewEndpointSliceInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*discoveryv1.EndpointSlice] {
	return NewTypedInformer(
		factory.Discovery().V1().EndpointSlices().Informer(),
		store,
		ed,
		"endpointslice",
		func(obj *discoveryv1.EndpointSlice) string { return obj.Namespace },
		func(obj *discoveryv1.EndpointSlice) string { return obj.Name },
	)
}

func NewIngressInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*networkingv1.Ingress] {
	return NewTypedInformer(
		factory.Networking().V1().Ingresses().Informer(),
		store,
		ed,
		"ingress",
		func(obj *networkingv1.Ingress) string { return obj.Namespace },
		func(obj *networkingv1.Ingress) string { return obj.Name },
	)
}

func NewNetworkPolicyInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*networkingv1.NetworkPolicy] {
	return NewTypedInformer(
		factory.Networking().V1().NetworkPolicies().Informer(),
		store,
		ed,
		"networkpolicy",
		func(obj *networkingv1.NetworkPolicy) string { return obj.Namespace },
		func(obj *networkingv1.NetworkPolicy) string { return obj.Name },
	)
}

func NewConfigMapInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*corev1.ConfigMap] {
	return NewTypedInformer(
		factory.Core().V1().ConfigMaps().Informer(),
		store,
		ed,
		"configmap",
		func(obj *corev1.ConfigMap) string { return obj.Namespace },
		func(obj *corev1.ConfigMap) string { return obj.Name },
	)
}

func NewSecretInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*corev1.Secret] {
	return NewTypedInformer(
		factory.Core().V1().Secrets().Informer(),
		store,
		ed,
		"secret",
		func(obj *corev1.Secret) string { return obj.Namespace },
		func(obj *corev1.Secret) string { return obj.Name },
	)
}

func NewPersistentVolumeClaimInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*corev1.PersistentVolumeClaim] {
	return NewTypedInformer(
		factory.Core().V1().PersistentVolumeClaims().Informer(),
		store,
		ed,
		"pvc",
		func(obj *corev1.PersistentVolumeClaim) string { return obj.Namespace },
		func(obj *corev1.PersistentVolumeClaim) string { return obj.Name },
	)
}

func NewServiceAccountInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*corev1.ServiceAccount] {
	return NewTypedInformer(
		factory.Core().V1().ServiceAccounts().Informer(),
		store,
		ed,
		"serviceaccount",
		func(obj *corev1.ServiceAccount) string { return obj.Namespace },
		func(obj *corev1.ServiceAccount) string { return obj.Name },
	)
}

func NewRoleInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*rbacv1.Role] {
	return NewTypedInformer(
		factory.Rbac().V1().Roles().Informer(),
		store,
		ed,
		"role",
		func(obj *rbacv1.Role) string { return obj.Namespace },
		func(obj *rbacv1.Role) string { return obj.Name },
	)
}

func NewRoleBindingInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*rbacv1.RoleBinding] {
	return NewTypedInformer(
		factory.Rbac().V1().RoleBindings().Informer(),
		store,
		ed,
		"rolebinding",
		func(obj *rbacv1.RoleBinding) string { return obj.Namespace },
		func(obj *rbacv1.RoleBinding) string { return obj.Name },
	)
}

func NewHorizontalPodAutoscalerInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*autoscalingv2.HorizontalPodAutoscaler] {
	return NewTypedInformer(
		factory.Autoscaling().V2().HorizontalPodAutoscalers().Informer(),
		store,
		ed,
		"hpa",
		func(obj *autoscalingv2.HorizontalPodAutoscaler) string { return obj.Namespace },
		func(obj *autoscalingv2.HorizontalPodAutoscaler) string { return obj.Name },
	)
}

func NewPodDisruptionBudgetInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*policyv1.PodDisruptionBudget] {
	return NewTypedInformer(
		factory.Policy().V1().PodDisruptionBudgets().Informer(),
		store,
		ed,
		"pdb",
		func(obj *policyv1.PodDisruptionBudget) string { return obj.Namespace },
		func(obj *policyv1.PodDisruptionBudget) string { return obj.Name },
	)
}

func NewResourceQuotaInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*corev1.ResourceQuota] {
	return NewTypedInformer(
		factory.Core().V1().ResourceQuotas().Informer(),
		store,
		ed,
		"resourcequota",
		func(obj *corev1.ResourceQuota) string { return obj.Namespace },
		func(obj *corev1.ResourceQuota) string { return obj.Name },
	)
}

func NewLimitRangeInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*corev1.LimitRange] {
	return NewTypedInformer(
		factory.Core().V1().LimitRanges().Informer(),
		store,
		ed,
		"limitrange",
		func(obj *corev1.LimitRange) string { return obj.Namespace },
		func(obj *corev1.LimitRange) string { return obj.Name },
	)
}

func NewEventInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*eventsv1.Event] {
	return NewTypedInformer(
		factory.Events().V1().Events().Informer(),
		store,
		ed,
		"event",
		func(obj *eventsv1.Event) string { return obj.Namespace },
		func(obj *eventsv1.Event) string { return obj.Name },
	)
}

func NewLeaseInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*coordinationv1.Lease] {
	return NewTypedInformer(
		factory.Coordination().V1().Leases().Informer(),
		store,
		ed,
		"lease",
		func(obj *coordinationv1.Lease) string { return obj.Namespace },
		func(obj *coordinationv1.Lease) string { return obj.Name },
	)
}

func NewControllerRevisionInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*appsv1.ControllerRevision] {
	return NewTypedInformer(
		factory.Apps().V1().ControllerRevisions().Informer(),
		store,
		ed,
		"controllerrevision",
		func(obj *appsv1.ControllerRevision) string { return obj.Namespace },
		func(obj *appsv1.ControllerRevision) string { return obj.Name },
	)
}

func NewPodTemplateInformer(
	factory informers.SharedInformerFactory,
	store store.Store,
	ed event.Distribution,
) *ResourceInformer[*corev1.PodTemplate] {
	return NewTypedInformer(
		factory.Core().V1().PodTemplates().Informer(),
		store,
		ed,
		"podtemplate",
		func(obj *corev1.PodTemplate) string { return obj.Namespace },
		func(obj *corev1.PodTemplate) string { return obj.Name },
	)
}
