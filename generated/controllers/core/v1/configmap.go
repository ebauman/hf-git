/*

© 2022 The HobbyFarm Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/
// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/generic"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type ConfigMapHandler func(string, *v1.ConfigMap) (*v1.ConfigMap, error)

type ConfigMapController interface {
	generic.ControllerMeta
	ConfigMapClient

	OnChange(ctx context.Context, name string, sync ConfigMapHandler)
	OnRemove(ctx context.Context, name string, sync ConfigMapHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() ConfigMapCache
}

type ConfigMapClient interface {
	Create(*v1.ConfigMap) (*v1.ConfigMap, error)
	Update(*v1.ConfigMap) (*v1.ConfigMap, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.ConfigMap, error)
	List(namespace string, opts metav1.ListOptions) (*v1.ConfigMapList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ConfigMap, err error)
}

type ConfigMapCache interface {
	Get(namespace, name string) (*v1.ConfigMap, error)
	List(namespace string, selector labels.Selector) ([]*v1.ConfigMap, error)

	AddIndexer(indexName string, indexer ConfigMapIndexer)
	GetByIndex(indexName, key string) ([]*v1.ConfigMap, error)
}

type ConfigMapIndexer func(obj *v1.ConfigMap) ([]string, error)

type configMapController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewConfigMapController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) ConfigMapController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &configMapController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromConfigMapHandlerToHandler(sync ConfigMapHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.ConfigMap
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.ConfigMap))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *configMapController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.ConfigMap))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateConfigMapDeepCopyOnChange(client ConfigMapClient, obj *v1.ConfigMap, handler func(obj *v1.ConfigMap) (*v1.ConfigMap, error)) (*v1.ConfigMap, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *configMapController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *configMapController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *configMapController) OnChange(ctx context.Context, name string, sync ConfigMapHandler) {
	c.AddGenericHandler(ctx, name, FromConfigMapHandlerToHandler(sync))
}

func (c *configMapController) OnRemove(ctx context.Context, name string, sync ConfigMapHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromConfigMapHandlerToHandler(sync)))
}

func (c *configMapController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *configMapController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *configMapController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *configMapController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *configMapController) Cache() ConfigMapCache {
	return &configMapCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *configMapController) Create(obj *v1.ConfigMap) (*v1.ConfigMap, error) {
	result := &v1.ConfigMap{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *configMapController) Update(obj *v1.ConfigMap) (*v1.ConfigMap, error) {
	result := &v1.ConfigMap{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *configMapController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *configMapController) Get(namespace, name string, options metav1.GetOptions) (*v1.ConfigMap, error) {
	result := &v1.ConfigMap{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *configMapController) List(namespace string, opts metav1.ListOptions) (*v1.ConfigMapList, error) {
	result := &v1.ConfigMapList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *configMapController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *configMapController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.ConfigMap, error) {
	result := &v1.ConfigMap{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type configMapCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *configMapCache) Get(namespace, name string) (*v1.ConfigMap, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.ConfigMap), nil
}

func (c *configMapCache) List(namespace string, selector labels.Selector) (ret []*v1.ConfigMap, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ConfigMap))
	})

	return ret, err
}

func (c *configMapCache) AddIndexer(indexName string, indexer ConfigMapIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.ConfigMap))
		},
	}))
}

func (c *configMapCache) GetByIndex(indexName, key string) (result []*v1.ConfigMap, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.ConfigMap, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.ConfigMap))
	}
	return result, nil
}
