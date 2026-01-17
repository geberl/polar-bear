package core

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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"polar-bear/internal/store"
)

// =============================================================================
// NON-NAMESPACED RESOURCES
// =============================================================================

// Namespace

func GetNamespace(store store.Store, name string) *corev1.Namespace {
	return GetResource[*corev1.Namespace](store, "", name)
}

func GetNamespaces(store store.Store) []*corev1.Namespace {
	return GetResources[*corev1.Namespace](store, "")
}

// Node

func GetNode(store store.Store, name string) *corev1.Node {
	return GetResource[*corev1.Node](store, "", name)
}

func GetNodes(store store.Store) []*corev1.Node {
	return GetResources[*corev1.Node](store, "")
}

// PersistentVolume

func GetPersistentVolume(store store.Store, name string) *corev1.PersistentVolume {
	return GetResource[*corev1.PersistentVolume](store, "", name)
}

func GetPersistentVolumes(store store.Store) []*corev1.PersistentVolume {
	return GetResources[*corev1.PersistentVolume](store, "")
}

// StorageClass

func GetStorageClass(store store.Store, name string) *storagev1.StorageClass {
	return GetResource[*storagev1.StorageClass](store, "", name)
}

func GetStorageClasses(store store.Store) []*storagev1.StorageClass {
	return GetResources[*storagev1.StorageClass](store, "")
}

// ClusterRole

func GetClusterRole(store store.Store, name string) *rbacv1.ClusterRole {
	return GetResource[*rbacv1.ClusterRole](store, "", name)
}

func GetClusterRoles(store store.Store) []*rbacv1.ClusterRole {
	return GetResources[*rbacv1.ClusterRole](store, "")
}

// ClusterRoleBinding

func GetClusterRoleBinding(store store.Store, name string) *rbacv1.ClusterRoleBinding {
	return GetResource[*rbacv1.ClusterRoleBinding](store, "", name)
}

func GetClusterRoleBindings(store store.Store) []*rbacv1.ClusterRoleBinding {
	return GetResources[*rbacv1.ClusterRoleBinding](store, "")
}

// CustomResourceDefinition

func GetCustomResourceDefinition(store store.Store, name string) *apiextensionsv1.CustomResourceDefinition {
	return GetResource[*apiextensionsv1.CustomResourceDefinition](store, "", name)
}

func GetCustomResourceDefinitions(store store.Store) []*apiextensionsv1.CustomResourceDefinition {
	return GetResources[*apiextensionsv1.CustomResourceDefinition](store, "")
}

// MutatingWebhookConfiguration

func GetMutatingWebhookConfiguration(
	store store.Store,
	name string,
) *admissionregistrationv1.MutatingWebhookConfiguration {
	return GetResource[*admissionregistrationv1.MutatingWebhookConfiguration](store, "", name)
}

func GetMutatingWebhookConfigurations(store store.Store) []*admissionregistrationv1.MutatingWebhookConfiguration {
	return GetResources[*admissionregistrationv1.MutatingWebhookConfiguration](store, "")
}

// ValidatingWebhookConfiguration

func GetValidatingWebhookConfiguration(
	store store.Store,
	name string,
) *admissionregistrationv1.ValidatingWebhookConfiguration {
	return GetResource[*admissionregistrationv1.ValidatingWebhookConfiguration](store, "", name)
}

func GetValidatingWebhookConfigurations(store store.Store) []*admissionregistrationv1.ValidatingWebhookConfiguration {
	return GetResources[*admissionregistrationv1.ValidatingWebhookConfiguration](store, "")
}

// PriorityClass

func GetPriorityClass(store store.Store, name string) *schedulingv1.PriorityClass {
	return GetResource[*schedulingv1.PriorityClass](store, "", name)
}

func GetPriorityClasses(store store.Store) []*schedulingv1.PriorityClass {
	return GetResources[*schedulingv1.PriorityClass](store, "")
}

// RuntimeClass

func GetRuntimeClass(store store.Store, name string) *nodev1.RuntimeClass {
	return GetResource[*nodev1.RuntimeClass](store, "", name)
}

func GetRuntimeClasses(store store.Store) []*nodev1.RuntimeClass {
	return GetResources[*nodev1.RuntimeClass](store, "")
}

// VolumeAttachment

func GetVolumeAttachment(store store.Store, name string) *storagev1.VolumeAttachment {
	return GetResource[*storagev1.VolumeAttachment](store, "", name)
}

func GetVolumeAttachments(store store.Store) []*storagev1.VolumeAttachment {
	return GetResources[*storagev1.VolumeAttachment](store, "")
}

// CSIDriver

func GetCSIDriver(store store.Store, name string) *storagev1.CSIDriver {
	return GetResource[*storagev1.CSIDriver](store, "", name)
}

func GetCSIDrivers(store store.Store) []*storagev1.CSIDriver {
	return GetResources[*storagev1.CSIDriver](store, "")
}

// CSINode

func GetCSINode(store store.Store, name string) *storagev1.CSINode {
	return GetResource[*storagev1.CSINode](store, "", name)
}

func GetCSINodes(store store.Store) []*storagev1.CSINode {
	return GetResources[*storagev1.CSINode](store, "")
}

// CSIStorageCapacity

func GetCSIStorageCapacity(store store.Store, name string) *storagev1.CSIStorageCapacity {
	return GetResource[*storagev1.CSIStorageCapacity](store, "", name)
}

func GetCSIStorageCapacities(store store.Store) []*storagev1.CSIStorageCapacity {
	return GetResources[*storagev1.CSIStorageCapacity](store, "")
}

// =============================================================================
// NAMESPACED RESOURCES
// =============================================================================

// Pod

func GetPod(store store.Store, ns string, name string) *corev1.Pod {
	return GetResource[*corev1.Pod](store, ns, name)
}
func GetPods(store store.Store, ns string) []*corev1.Pod {
	return GetResources[*corev1.Pod](store, ns)
}
func CountPods(store store.Store, ns string) uint {
	return CountResources[*corev1.Pod](store, ns)
}

// Deployment

func GetDeployment(store store.Store, ns string, name string) *appsv1.Deployment {
	return GetResource[*appsv1.Deployment](store, ns, name)
}
func GetDeployments(store store.Store, ns string) []*appsv1.Deployment {
	return GetResources[*appsv1.Deployment](store, ns)
}
func CountDeployments(store store.Store, ns string) uint {
	return CountResources[*appsv1.Deployment](store, ns)
}

// StatefulSet

func GetStatefulSet(store store.Store, ns string, name string) *appsv1.StatefulSet {
	return GetResource[*appsv1.StatefulSet](store, ns, name)
}
func GetStatefulSets(store store.Store, ns string) []*appsv1.StatefulSet {
	return GetResources[*appsv1.StatefulSet](store, ns)
}
func CountStatefulSets(store store.Store, ns string) uint {
	return CountResources[*appsv1.StatefulSet](store, ns)
}

// DaemonSet

func GetDaemonSet(store store.Store, ns string, name string) *appsv1.DaemonSet {
	return GetResource[*appsv1.DaemonSet](store, ns, name)
}
func GetDaemonSets(store store.Store, ns string) []*appsv1.DaemonSet {
	return GetResources[*appsv1.DaemonSet](store, ns)
}
func CountDaemonSets(store store.Store, ns string) uint {
	return CountResources[*appsv1.DaemonSet](store, ns)
}

// ReplicaSet

func GetReplicaSet(store store.Store, ns string, name string) *appsv1.ReplicaSet {
	return GetResource[*appsv1.ReplicaSet](store, ns, name)
}
func GetReplicaSets(store store.Store, ns string) []*appsv1.ReplicaSet {
	return GetResources[*appsv1.ReplicaSet](store, ns)
}
func CountReplicaSets(store store.Store, ns string) uint {
	return CountResources[*appsv1.ReplicaSet](store, ns)
}

// Job

func GetJob(store store.Store, ns string, name string) *batchv1.Job {
	return GetResource[*batchv1.Job](store, ns, name)
}
func GetJobs(store store.Store, ns string) []*batchv1.Job {
	return GetResources[*batchv1.Job](store, ns)
}
func CountJobs(store store.Store, ns string) uint {
	return CountResources[*batchv1.Job](store, ns)
}

// CronJob

func GetCronJob(store store.Store, ns string, name string) *batchv1.CronJob {
	return GetResource[*batchv1.CronJob](store, ns, name)
}
func GetCronJobs(store store.Store, ns string) []*batchv1.CronJob {
	return GetResources[*batchv1.CronJob](store, ns)
}

// ReplicationController

func GetReplicationController(store store.Store, ns string, name string) *corev1.ReplicationController {
	return GetResource[*corev1.ReplicationController](store, ns, name)
}
func GetReplicationControllers(store store.Store, ns string) []*corev1.ReplicationController {
	return GetResources[*corev1.ReplicationController](store, ns)
}

// Service

func GetService(store store.Store, ns string, name string) *corev1.Service {
	return GetResource[*corev1.Service](store, ns, name)
}
func GetServices(store store.Store, ns string) []*corev1.Service {
	return GetResources[*corev1.Service](store, ns)
}
func CountServices(store store.Store, ns string) uint {
	return CountResources[*corev1.Service](store, ns)
}

// Endpoints

func GetEndpointSlice(store store.Store, ns string, name string) *discoveryv1.EndpointSlice {
	return GetResource[*discoveryv1.EndpointSlice](store, ns, name)
}
func GetEndpointSlices(store store.Store, ns string) []*discoveryv1.EndpointSlice {
	return GetResources[*discoveryv1.EndpointSlice](store, ns)
}

// Ingress

func GetIngress(store store.Store, ns string, name string) *networkingv1.Ingress {
	return GetResource[*networkingv1.Ingress](store, ns, name)
}
func GetIngresses(store store.Store, ns string) []*networkingv1.Ingress {
	return GetResources[*networkingv1.Ingress](store, ns)
}
func CountIngresses(store store.Store, ns string) uint {
	return CountResources[*networkingv1.Ingress](store, ns)
}

// NetworkPolicy

func GetNetworkPolicy(store store.Store, ns string, name string) *networkingv1.NetworkPolicy {
	return GetResource[*networkingv1.NetworkPolicy](store, ns, name)
}
func GetNetworkPolicies(store store.Store, ns string) []*networkingv1.NetworkPolicy {
	return GetResources[*networkingv1.NetworkPolicy](store, ns)
}

// ConfigMap

func GetConfigMap(store store.Store, ns string, name string) *corev1.ConfigMap {
	return GetResource[*corev1.ConfigMap](store, ns, name)
}
func GetConfigMaps(store store.Store, ns string) []*corev1.ConfigMap {
	return GetResources[*corev1.ConfigMap](store, ns)
}

// Secret

func GetSecret(store store.Store, ns string, name string) *corev1.Secret {
	return GetResource[*corev1.Secret](store, ns, name)
}
func GetSecrets(store store.Store, ns string) []*corev1.Secret {
	return GetResources[*corev1.Secret](store, ns)
}

// PersistentVolumeClaim

func GetPersistentVolumeClaim(store store.Store, ns string, name string) *corev1.PersistentVolumeClaim {
	return GetResource[*corev1.PersistentVolumeClaim](store, ns, name)
}
func GetPersistentVolumeClaims(store store.Store, ns string) []*corev1.PersistentVolumeClaim {
	return GetResources[*corev1.PersistentVolumeClaim](store, ns)
}

// ServiceAccount

func GetServiceAccount(store store.Store, ns string, name string) *corev1.ServiceAccount {
	return GetResource[*corev1.ServiceAccount](store, ns, name)
}
func GetServiceAccounts(store store.Store, ns string) []*corev1.ServiceAccount {
	return GetResources[*corev1.ServiceAccount](store, ns)
}

// Role

func GetRole(store store.Store, ns string, name string) *rbacv1.Role {
	return GetResource[*rbacv1.Role](store, ns, name)
}
func GetRoles(store store.Store, ns string) []*rbacv1.Role {
	return GetResources[*rbacv1.Role](store, ns)
}

// RoleBinding

func GetRoleBinding(store store.Store, ns string, name string) *rbacv1.RoleBinding {
	return GetResource[*rbacv1.RoleBinding](store, ns, name)
}
func GetRoleBindings(store store.Store, ns string) []*rbacv1.RoleBinding {
	return GetResources[*rbacv1.RoleBinding](store, ns)
}

// HorizontalPodAutoscaler

func GetHorizontalPodAutoscaler(store store.Store, ns string, name string) *autoscalingv2.HorizontalPodAutoscaler {
	return GetResource[*autoscalingv2.HorizontalPodAutoscaler](store, ns, name)
}
func GetHorizontalPodAutoscalers(store store.Store, ns string) []*autoscalingv2.HorizontalPodAutoscaler {
	return GetResources[*autoscalingv2.HorizontalPodAutoscaler](store, ns)
}

// PodDisruptionBudget

func GetPodDisruptionBudget(store store.Store, ns string, name string) *policyv1.PodDisruptionBudget {
	return GetResource[*policyv1.PodDisruptionBudget](store, ns, name)
}
func GetPodDisruptionBudgets(store store.Store, ns string) []*policyv1.PodDisruptionBudget {
	return GetResources[*policyv1.PodDisruptionBudget](store, ns)
}

// ResourceQuota

func GetResourceQuota(store store.Store, ns string, name string) *corev1.ResourceQuota {
	return GetResource[*corev1.ResourceQuota](store, ns, name)
}
func GetResourceQuotas(store store.Store, ns string) []*corev1.ResourceQuota {
	return GetResources[*corev1.ResourceQuota](store, ns)
}

// LimitRange

func GetLimitRange(store store.Store, ns string, name string) *corev1.LimitRange {
	return GetResource[*corev1.LimitRange](store, ns, name)
}
func GetLimitRanges(store store.Store, ns string) []*corev1.LimitRange {
	return GetResources[*corev1.LimitRange](store, ns)
}

// Event

func GetEvent(store store.Store, ns string, name string) *eventsv1.Event {
	return GetResource[*eventsv1.Event](store, ns, name)
}
func GetEvents(store store.Store, ns string) []*eventsv1.Event {
	return GetResources[*eventsv1.Event](store, ns)
}

// Lease

func GetLease(store store.Store, ns string, name string) *coordinationv1.Lease {
	return GetResource[*coordinationv1.Lease](store, ns, name)
}
func GetLeases(store store.Store, ns string) []*coordinationv1.Lease {
	return GetResources[*coordinationv1.Lease](store, ns)
}

// ControllerRevision

func GetControllerRevision(store store.Store, ns string, name string) *appsv1.ControllerRevision {
	return GetResource[*appsv1.ControllerRevision](store, ns, name)
}
func GetControllerRevisions(store store.Store, ns string) []*appsv1.ControllerRevision {
	return GetResources[*appsv1.ControllerRevision](store, ns)
}

// PodTemplate

func GetPodTemplate(store store.Store, ns string, name string) *corev1.PodTemplate {
	return GetResource[*corev1.PodTemplate](store, ns, name)
}
func GetPodTemplates(store store.Store, ns string) []*corev1.PodTemplate {
	return GetResources[*corev1.PodTemplate](store, ns)
}
