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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "harmonycloud.cn/multi-cluster-manager/pkg/apis/multicluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMultiClusterResources implements MultiClusterResourceInterface
type FakeMultiClusterResources struct {
	Fake *FakeMulticlusterV1alpha1
	ns   string
}

var multiclusterresourcesResource = schema.GroupVersionResource{Group: "multicluster.harmonycloud.cn", Version: "v1alpha1", Resource: "multiclusterresources"}

var multiclusterresourcesKind = schema.GroupVersionKind{Group: "multicluster.harmonycloud.cn", Version: "v1alpha1", Kind: "MultiClusterResource"}

// Get takes name of the multiClusterResource, and returns the corresponding multiClusterResource object, and an error if there is any.
func (c *FakeMultiClusterResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MultiClusterResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(multiclusterresourcesResource, c.ns, name), &v1alpha1.MultiClusterResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterResource), err
}

// List takes label and field selectors, and returns the list of MultiClusterResources that match those selectors.
func (c *FakeMultiClusterResources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MultiClusterResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(multiclusterresourcesResource, multiclusterresourcesKind, c.ns, opts), &v1alpha1.MultiClusterResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MultiClusterResourceList{ListMeta: obj.(*v1alpha1.MultiClusterResourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.MultiClusterResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested multiClusterResources.
func (c *FakeMultiClusterResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(multiclusterresourcesResource, c.ns, opts))

}

// Create takes the representation of a multiClusterResource and creates it.  Returns the server's representation of the multiClusterResource, and an error, if there is any.
func (c *FakeMultiClusterResources) Create(ctx context.Context, multiClusterResource *v1alpha1.MultiClusterResource, opts v1.CreateOptions) (result *v1alpha1.MultiClusterResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(multiclusterresourcesResource, c.ns, multiClusterResource), &v1alpha1.MultiClusterResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterResource), err
}

// Update takes the representation of a multiClusterResource and updates it. Returns the server's representation of the multiClusterResource, and an error, if there is any.
func (c *FakeMultiClusterResources) Update(ctx context.Context, multiClusterResource *v1alpha1.MultiClusterResource, opts v1.UpdateOptions) (result *v1alpha1.MultiClusterResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(multiclusterresourcesResource, c.ns, multiClusterResource), &v1alpha1.MultiClusterResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterResource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMultiClusterResources) UpdateStatus(ctx context.Context, multiClusterResource *v1alpha1.MultiClusterResource, opts v1.UpdateOptions) (*v1alpha1.MultiClusterResource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(multiclusterresourcesResource, "status", c.ns, multiClusterResource), &v1alpha1.MultiClusterResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterResource), err
}

// Delete takes name of the multiClusterResource and deletes it. Returns an error if one occurs.
func (c *FakeMultiClusterResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(multiclusterresourcesResource, c.ns, name), &v1alpha1.MultiClusterResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMultiClusterResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(multiclusterresourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MultiClusterResourceList{})
	return err
}

// Patch applies the patch and returns the patched multiClusterResource.
func (c *FakeMultiClusterResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MultiClusterResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(multiclusterresourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MultiClusterResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterResource), err
}
