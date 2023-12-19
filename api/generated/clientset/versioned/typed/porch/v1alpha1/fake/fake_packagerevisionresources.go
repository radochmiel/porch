// Copyright 2023 The kpt and Nephio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/nephio-project/porch/api/porch/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePackageRevisionResources implements PackageRevisionResourcesInterface
type FakePackageRevisionResources struct {
	Fake *FakePorchV1alpha1
	ns   string
}

var packagerevisionresourcesResource = schema.GroupVersionResource{Group: "porch.kpt.dev", Version: "v1alpha1", Resource: "packagerevisionresources"}

var packagerevisionresourcesKind = schema.GroupVersionKind{Group: "porch.kpt.dev", Version: "v1alpha1", Kind: "PackageRevisionResources"}

// Get takes name of the packageRevisionResources, and returns the corresponding packageRevisionResources object, and an error if there is any.
func (c *FakePackageRevisionResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PackageRevisionResources, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(packagerevisionresourcesResource, c.ns, name), &v1alpha1.PackageRevisionResources{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRevisionResources), err
}

// List takes label and field selectors, and returns the list of PackageRevisionResources that match those selectors.
func (c *FakePackageRevisionResources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PackageRevisionResourcesList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(packagerevisionresourcesResource, packagerevisionresourcesKind, c.ns, opts), &v1alpha1.PackageRevisionResourcesList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PackageRevisionResourcesList{ListMeta: obj.(*v1alpha1.PackageRevisionResourcesList).ListMeta}
	for _, item := range obj.(*v1alpha1.PackageRevisionResourcesList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested packageRevisionResources.
func (c *FakePackageRevisionResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(packagerevisionresourcesResource, c.ns, opts))

}

// Create takes the representation of a packageRevisionResources and creates it.  Returns the server's representation of the packageRevisionResources, and an error, if there is any.
func (c *FakePackageRevisionResources) Create(ctx context.Context, packageRevisionResources *v1alpha1.PackageRevisionResources, opts v1.CreateOptions) (result *v1alpha1.PackageRevisionResources, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(packagerevisionresourcesResource, c.ns, packageRevisionResources), &v1alpha1.PackageRevisionResources{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRevisionResources), err
}

// Update takes the representation of a packageRevisionResources and updates it. Returns the server's representation of the packageRevisionResources, and an error, if there is any.
func (c *FakePackageRevisionResources) Update(ctx context.Context, packageRevisionResources *v1alpha1.PackageRevisionResources, opts v1.UpdateOptions) (result *v1alpha1.PackageRevisionResources, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(packagerevisionresourcesResource, c.ns, packageRevisionResources), &v1alpha1.PackageRevisionResources{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRevisionResources), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePackageRevisionResources) UpdateStatus(ctx context.Context, packageRevisionResources *v1alpha1.PackageRevisionResources, opts v1.UpdateOptions) (*v1alpha1.PackageRevisionResources, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(packagerevisionresourcesResource, "status", c.ns, packageRevisionResources), &v1alpha1.PackageRevisionResources{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRevisionResources), err
}

// Delete takes name of the packageRevisionResources and deletes it. Returns an error if one occurs.
func (c *FakePackageRevisionResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(packagerevisionresourcesResource, c.ns, name, opts), &v1alpha1.PackageRevisionResources{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePackageRevisionResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(packagerevisionresourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PackageRevisionResourcesList{})
	return err
}

// Patch applies the patch and returns the patched packageRevisionResources.
func (c *FakePackageRevisionResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PackageRevisionResources, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(packagerevisionresourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.PackageRevisionResources{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRevisionResources), err
}
