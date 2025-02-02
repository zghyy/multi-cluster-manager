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
	v1alpha1 "harmonycloud.cn/multi-cluster-manager/pkg/apis/multicluster/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MultiClusterResourceLister helps list MultiClusterResources.
// All objects returned here must be treated as read-only.
type MultiClusterResourceLister interface {
	// List lists all MultiClusterResources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MultiClusterResource, err error)
	// MultiClusterResources returns an object that can list and get MultiClusterResources.
	MultiClusterResources(namespace string) MultiClusterResourceNamespaceLister
	MultiClusterResourceListerExpansion
}

// multiClusterResourceLister implements the MultiClusterResourceLister interface.
type multiClusterResourceLister struct {
	indexer cache.Indexer
}

// NewMultiClusterResourceLister returns a new MultiClusterResourceLister.
func NewMultiClusterResourceLister(indexer cache.Indexer) MultiClusterResourceLister {
	return &multiClusterResourceLister{indexer: indexer}
}

// List lists all MultiClusterResources in the indexer.
func (s *multiClusterResourceLister) List(selector labels.Selector) (ret []*v1alpha1.MultiClusterResource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MultiClusterResource))
	})
	return ret, err
}

// MultiClusterResources returns an object that can list and get MultiClusterResources.
func (s *multiClusterResourceLister) MultiClusterResources(namespace string) MultiClusterResourceNamespaceLister {
	return multiClusterResourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MultiClusterResourceNamespaceLister helps list and get MultiClusterResources.
// All objects returned here must be treated as read-only.
type MultiClusterResourceNamespaceLister interface {
	// List lists all MultiClusterResources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MultiClusterResource, err error)
	// Get retrieves the MultiClusterResource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MultiClusterResource, error)
	MultiClusterResourceNamespaceListerExpansion
}

// multiClusterResourceNamespaceLister implements the MultiClusterResourceNamespaceLister
// interface.
type multiClusterResourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MultiClusterResources in the indexer for a given namespace.
func (s multiClusterResourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MultiClusterResource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MultiClusterResource))
	})
	return ret, err
}

// Get retrieves the MultiClusterResource from the indexer for a given namespace and name.
func (s multiClusterResourceNamespaceLister) Get(name string) (*v1alpha1.MultiClusterResource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("multiclusterresource"), name)
	}
	return obj.(*v1alpha1.MultiClusterResource), nil
}
