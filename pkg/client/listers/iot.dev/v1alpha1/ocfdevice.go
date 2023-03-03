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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/vitu1234/iot-operator/pkg/apis/iot.dev/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OCFDeviceLister helps list OCFDevices.
// All objects returned here must be treated as read-only.
type OCFDeviceLister interface {
	// List lists all OCFDevices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OCFDevice, err error)
	// OCFDevices returns an object that can list and get OCFDevices.
	OCFDevices(namespace string) OCFDeviceNamespaceLister
	OCFDeviceListerExpansion
}

// oCFDeviceLister implements the OCFDeviceLister interface.
type oCFDeviceLister struct {
	indexer cache.Indexer
}

// NewOCFDeviceLister returns a new OCFDeviceLister.
func NewOCFDeviceLister(indexer cache.Indexer) OCFDeviceLister {
	return &oCFDeviceLister{indexer: indexer}
}

// List lists all OCFDevices in the indexer.
func (s *oCFDeviceLister) List(selector labels.Selector) (ret []*v1alpha1.OCFDevice, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OCFDevice))
	})
	return ret, err
}

// OCFDevices returns an object that can list and get OCFDevices.
func (s *oCFDeviceLister) OCFDevices(namespace string) OCFDeviceNamespaceLister {
	return oCFDeviceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OCFDeviceNamespaceLister helps list and get OCFDevices.
// All objects returned here must be treated as read-only.
type OCFDeviceNamespaceLister interface {
	// List lists all OCFDevices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OCFDevice, err error)
	// Get retrieves the OCFDevice from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OCFDevice, error)
	OCFDeviceNamespaceListerExpansion
}

// oCFDeviceNamespaceLister implements the OCFDeviceNamespaceLister
// interface.
type oCFDeviceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OCFDevices in the indexer for a given namespace.
func (s oCFDeviceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OCFDevice, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OCFDevice))
	})
	return ret, err
}

// Get retrieves the OCFDevice from the indexer for a given namespace and name.
func (s oCFDeviceNamespaceLister) Get(name string) (*v1alpha1.OCFDevice, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ocfdevice"), name)
	}
	return obj.(*v1alpha1.OCFDevice), nil
}