package informer

import (
	"os"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Informer interface {
	Run() (err error)
	Close() (err error)
	Kind() string
}

func NewInformerFactory(namespace string) (informers.SharedInformerFactory, error) {
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		return nil, err
	}

	config.QPS = 20
	config.Burst = 100

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	factory := informers.NewSharedInformerFactoryWithOptions(
		client,
		0,
		informers.WithNamespace(namespace), // can be empty string, watches everything it can watch
	)

	return factory, nil
}
