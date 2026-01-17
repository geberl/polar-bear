package shared

import (
	"fmt"
	"strings"

	"github.com/a-h/templ"
)

func registryURL(name string) (string, error) {
	// GitHub Container Registry
	nameNoPrefix, found := strings.CutPrefix(name, "ghcr.io/")
	if found {
		url := fmt.Sprintf("https://github.com/%s", nameNoPrefix)
		return url, nil
	}

	// Codeberg project
	nameNoPrefix, found = strings.CutPrefix(name, "codeberg.org/")
	if found {
		url := fmt.Sprintf("https://codeberg.org/%s/releases", nameNoPrefix)
		return url, nil
	}

	// Quay project
	nameNoPrefix, found = strings.CutPrefix(name, "quay.io/")
	if found {
		url := fmt.Sprintf("https://quay.io/repository/%s", nameNoPrefix)
		return url, nil
	}

	// Kubernetes project
	if strings.HasPrefix(name, "registry.k8s.io/") {
		parts := strings.Split(name, "/")
		if len(parts) >= 1 {
			url := fmt.Sprintf("https://github.com/kubernetes-sigs/%s", parts[1])
			return url, nil
		}
		return "", fmt.Errorf("no registry URL found")
	}

	// Docker Hub named user, new version
	nameNoPrefix, found = strings.CutPrefix(name, "docker.io/")
	if found {
		url := fmt.Sprintf("https://hub.docker.com/r/%s", nameNoPrefix)
		return url, nil
	}

	// Docker Hub
	seps := strings.Count(name, "/")
	if seps == 0 {
		// Official image
		url := fmt.Sprintf("https://hub.docker.com/_/%s", name)
		return url, nil
	}
	if seps == 1 {
		// Named user, classical version
		url := fmt.Sprintf("https://hub.docker.com/r/%s", name)
		return url, nil
	}

	return "", fmt.Errorf("no registry URL found")
}

func RegistryLink(name string) templ.SafeURL {
	url, err := registryURL(name)
	if err != nil {
		return templ.URL("https://hub.docker.com/")
	}
	return templ.URL(url)
}

func ClusterLink(item string) templ.SafeURL {
	return templ.URL(fmt.Sprintf("/%s", item))
}

func NamespacesLink() templ.SafeURL {
	return templ.URL("/ns")
}

func NamespaceLink(name string) templ.SafeURL {
	return templ.URL(fmt.Sprintf("/ns/%s", name))
}

func NodesLink() templ.SafeURL {
	return templ.URL("/no")
}

func NodeLink(name string) templ.SafeURL {
	return templ.URL(fmt.Sprintf("/no/%s", name))
}

func PodsLink(name string) templ.SafeURL {
	return templ.URL(fmt.Sprintf("/ns/%s/pd", name))
}

func PodLink(ns string, name string) templ.SafeURL {
	return templ.URL(fmt.Sprintf("/ns/%s/pd/%s", ns, name))
}

func DeploymentsLink(ns string) templ.SafeURL {
	return templ.URL(fmt.Sprintf("/ns/%s/deploy", ns))
}

func DeploymentLink(ns string, name string) templ.SafeURL {
	return templ.URL(fmt.Sprintf("/ns/%s/deploy/%s", ns, name))
}

func ReplicaSetsLink(ns string) templ.SafeURL {
	return templ.URL(fmt.Sprintf("/ns/%s/rs", ns))
}

func ReplicaSetLink(ns string, name string) templ.SafeURL {
	return templ.URL(fmt.Sprintf("/ns/%s/rs/%s", ns, name))
}

func StatefulSetsLink(ns string) templ.SafeURL {
	return templ.URL(fmt.Sprintf("/ns/%s/sts", ns))
}

func StatefulSetLink(ns string, name string) templ.SafeURL {
	return templ.URL(fmt.Sprintf("/ns/%s/sts/%s", ns, name))
}
