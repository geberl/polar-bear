package informer

import (
	"encoding/json"
	"log/slog"

	"k8s.io/client-go/tools/cache"

	"polar-bear/internal/core"
	"polar-bear/internal/event"
	"polar-bear/internal/store"
)

type ResourceInformer[T any] struct {
	logger       *slog.Logger
	stopper      chan struct{}
	inf          cache.SharedIndexInformer
	store        store.Store
	event        event.Distribution
	resourceType string             // string representation of the resource (e.g., "namespace", "node")
	getNamespace func(obj T) string // function that extracts the namespace from the resource
	getName      func(obj T) string // function that extracts the name from the resource
}

func NewResourceInformer[T any](
	stopper chan struct{},
	inf cache.SharedIndexInformer,
	store store.Store,
	ed event.Distribution,
	resourceType string,
	getNamespace func(obj T) string,
	getName func(obj T) string,
) *ResourceInformer[T] {
	return &ResourceInformer[T]{
		logger:       slog.With("component", "informer"),
		stopper:      stopper,
		inf:          inf,
		store:        store,
		event:        ed,
		resourceType: resourceType,
		getNamespace: getNamespace,
		getName:      getName,
	}
}

func (informer *ResourceInformer[T]) Run() error {
	_, err := informer.inf.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			resource := obj.(T)

			dbKey, err := core.ResourceKey(
				informer.resourceType,
				informer.getNamespace(resource),
				informer.getName(resource),
			)
			if err != nil {
				informer.logger.Error(
					"unable to get resource key",
					"kind", informer.resourceType,
					"namespace", informer.getNamespace(resource),
					"name", informer.getName(resource),
					"error", err,
				)
			}

			informer.logger.Info(
				"add resource",
				"kind", informer.resourceType,
				"key", string(dbKey),
			)

			dbVal, err := json.Marshal(resource)
			if err != nil {
				informer.logger.Error(
					"unable to marshal to json",
					"kind", informer.resourceType,
					"key", string(dbKey),
					"error", err,
				)
				return
			}

			err = informer.store.Set(dbKey, dbVal)
			if err != nil {
				informer.logger.Error(
					"unable to add to store",
					"kind", informer.resourceType,
					"key", string(dbKey),
					"error", err,
				)
			}
			informer.event.Send(string(dbKey))
		},
		UpdateFunc: func(_, newObj any) {
			resource := newObj.(T)

			dbKey, err := core.ResourceKey(
				informer.resourceType,
				informer.getNamespace(resource),
				informer.getName(resource),
			)
			if err != nil {
				informer.logger.Error(
					"unable to get resource key",
					"kind", informer.resourceType,
					"namespace", informer.getNamespace(resource),
					"name", informer.getName(resource),
					"error", err,
				)
			}

			informer.logger.Info(
				"update resource",
				"kind", informer.resourceType,
				"key", string(dbKey),
			)

			dbVal, err := json.Marshal(resource)
			if err != nil {
				informer.logger.Error(
					"unable to marshal to json",
					"kind", informer.resourceType,
					"key", string(dbKey),
					"error", err,
				)
				return
			}

			err = informer.store.Set(dbKey, dbVal)
			if err != nil {
				informer.logger.Error(
					"unable to update in store",
					"kind", informer.resourceType,
					"key", string(dbKey),
					"error", err,
				)
			}
			informer.event.Send(string(dbKey))
		},
		DeleteFunc: func(obj any) {
			resource := obj.(T)

			dbKey, err := core.ResourceKey(
				informer.resourceType,
				informer.getNamespace(resource),
				informer.getName(resource),
			)
			if err != nil {
				informer.logger.Error(
					"unable to get resource key",
					"kind", informer.resourceType,
					"namespace", informer.getNamespace(resource),
					"name", informer.getName(resource),
					"error", err,
				)
			}

			informer.logger.Info(
				"delete resource",
				"kind", informer.resourceType,
				"key", string(dbKey),
			)

			err = informer.store.Delete(dbKey)
			if err != nil {
				informer.logger.Error(
					"unable to remove from store",
					"kind", informer.resourceType,
					"key", string(dbKey),
					"error", err,
				)
			}
			informer.event.Send(string(dbKey))
		},
	})
	if err != nil {
		return err
	}
	informer.inf.Run(informer.stopper)
	return nil
}

func (informer *ResourceInformer[T]) Close() error {
	informer.logger.Info("closing namespace informer")
	close(informer.stopper)
	return nil
}

func (informer *ResourceInformer[T]) Kind() string {
	return informer.resourceType
}
