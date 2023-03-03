/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	iotdevv1alpha1 "github.com/vitu1234/iot-operator/pkg/apis/iot.dev/v1alpha1"
	versioned "github.com/vitu1234/iot-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/vitu1234/iot-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/vitu1234/iot-operator/pkg/client/listers/iot.dev/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OCFDeviceInformer provides access to a shared informer and lister for
// OCFDevices.
type OCFDeviceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.OCFDeviceLister
}

type oCFDeviceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewOCFDeviceInformer constructs a new informer for OCFDevice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOCFDeviceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOCFDeviceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredOCFDeviceInformer constructs a new informer for OCFDevice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOCFDeviceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IotV1alpha1().OCFDevices(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IotV1alpha1().OCFDevices(namespace).Watch(context.TODO(), options)
			},
		},
		&iotdevv1alpha1.OCFDevice{},
		resyncPeriod,
		indexers,
	)
}

func (f *oCFDeviceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOCFDeviceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *oCFDeviceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&iotdevv1alpha1.OCFDevice{}, f.defaultInformer)
}

func (f *oCFDeviceInformer) Lister() v1alpha1.OCFDeviceLister {
	return v1alpha1.NewOCFDeviceLister(f.Informer().GetIndexer())
}