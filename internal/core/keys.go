package core

import (
	"fmt"
)

func ResourceKey(kind string, ns string, name string) ([]byte, error) {
	// Non-namespaced resource kinds
	switch kind {
	case "namespace":
		return fmt.Appendf(nil, "namespace/%s", name), nil
	case "node":
		return fmt.Appendf(nil, "node/%s", name), nil
	case "persistentvolume":
		return fmt.Appendf(nil, "pv/%s", name), nil
	case "storageclass":
		return fmt.Appendf(nil, "sc/%s", name), nil
	case "clusterrole":
		return fmt.Appendf(nil, "clusterrole/%s", name), nil
	case "clusterrolebinding":
		return fmt.Appendf(nil, "clusterrolebinding/%s", name), nil
	case "customresourcedefinition":
		return fmt.Appendf(nil, "crd/%s", name), nil
	case "mutatingadmissionconfiguration":
		return fmt.Appendf(nil, "mutatingadmissionconfig/%s", name), nil
	case "validatingadmissionconfiguration":
		return fmt.Appendf(nil, "validatingadmissionconfig/%s", name), nil
	case "priorityclass":
		return fmt.Appendf(nil, "priorityclass/%s", name), nil
	case "runtimeclass":
		return fmt.Appendf(nil, "runtimeclass/%s", name), nil
	case "volumeattachment":
		return fmt.Appendf(nil, "volumeattachment/%s", name), nil
	case "csidriver":
		return fmt.Appendf(nil, "csidriver/%s", name), nil
	case "csinode":
		return fmt.Appendf(nil, "csinode/%s", name), nil
	case "csistoragecapacity":
		return fmt.Appendf(nil, "csistoragecapacity/%s", name), nil
	}

	// Namespaced resource kinds
	if ns == "" {
		return []byte(""), fmt.Errorf("got empty namespace for namespaced resource")
	}

	switch kind {
	// Workload resources
	case "pod":
		return fmt.Appendf(nil, "ns/%s/pd/%s", ns, name), nil
	case "deployment":
		return fmt.Appendf(nil, "ns/%s/deploy/%s", ns, name), nil
	case "statefulset":
		return fmt.Appendf(nil, "ns/%s/st/%s", ns, name), nil
	case "daemonset":
		return fmt.Appendf(nil, "ns/%s/ds/%s", ns, name), nil
	case "replicaset":
		return fmt.Appendf(nil, "ns/%s/rs/%s", ns, name), nil
	case "job":
		return fmt.Appendf(nil, "ns/%s/job/%s", ns, name), nil
	case "cronjob":
		return fmt.Appendf(nil, "ns/%s/cronjob/%s", ns, name), nil
	case "replicationcontroller":
		return fmt.Appendf(nil, "ns/%s/rc/%s", ns, name), nil

	// Service and networking resources
	case "service":
		return fmt.Appendf(nil, "ns/%s/svc/%s", ns, name), nil
	case "endpoints":
		return fmt.Appendf(nil, "ns/%s/ep/%s", ns, name), nil
	case "endpointslice":
		return fmt.Appendf(nil, "ns/%s/epslice/%s", ns, name), nil
	case "ingress":
		return fmt.Appendf(nil, "ns/%s/ing/%s", ns, name), nil
	case "networkpolicy":
		return fmt.Appendf(nil, "ns/%s/netpol/%s", ns, name), nil

	// Configuration and storage resources
	case "configmap":
		return fmt.Appendf(nil, "ns/%s/cm/%s", ns, name), nil
	case "secret":
		return fmt.Appendf(nil, "ns/%s/secret/%s", ns, name), nil
	case "persistentvolumeclaim":
		return fmt.Appendf(nil, "ns/%s/pvc/%s", ns, name), nil

	// Authorization resources
	case "serviceaccount":
		return fmt.Appendf(nil, "ns/%s/sa/%s", ns, name), nil
	case "role":
		return fmt.Appendf(nil, "ns/%s/role/%s", ns, name), nil
	case "rolebinding":
		return fmt.Appendf(nil, "ns/%s/rolebinding/%s", ns, name), nil

	// Autoscaling resources
	case "horizontalpodautoscaler":
		return fmt.Appendf(nil, "ns/%s/hpa/%s", ns, name), nil
	case "verticalpodautoscaler":
		return fmt.Appendf(nil, "ns/%s/vpa/%s", ns, name), nil
	case "poddisruptionbudget":
		return fmt.Appendf(nil, "ns/%s/pdb/%s", ns, name), nil

	// Policy resources
	case "podsecuritypolicy":
		return fmt.Appendf(nil, "ns/%s/psp/%s", ns, name), nil
	case "resourcequota":
		return fmt.Appendf(nil, "ns/%s/quota/%s", ns, name), nil
	case "limitrange":
		return fmt.Appendf(nil, "ns/%s/limitrange/%s", ns, name), nil

	// Event and monitoring resources
	case "event":
		return fmt.Appendf(nil, "ns/%s/event/%s", ns, name), nil
	case "lease":
		return fmt.Appendf(nil, "ns/%s/lease/%s", ns, name), nil

	// Extension resources
	case "mutatingwebhookconfiguration":
		return fmt.Appendf(nil, "ns/%s/mutatingwebhook/%s", ns, name), nil
	case "validatingwebhookconfiguration":
		return fmt.Appendf(nil, "ns/%s/validatingwebhook/%s", ns, name), nil

	// API machinery resources
	case "controllerrevision":
		return fmt.Appendf(nil, "ns/%s/controllerrevision/%s", ns, name), nil

	// Batch resources
	case "podtemplate":
		return fmt.Appendf(nil, "ns/%s/podtemplate/%s", ns, name), nil
	}

	// Unsupported resource kinds
	return []byte(""), fmt.Errorf("got unsupported resource kind")
}
