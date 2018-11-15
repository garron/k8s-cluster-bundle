// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-cluster-bundle/pkg/apis/bundle/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComponentSets implements ComponentSetInterface
type FakeComponentSets struct {
	Fake *FakeBundleV1alpha1
	ns   string
}

var componentsetsResource = schema.GroupVersionResource{Group: "bundle.gke.io", Version: "v1alpha1", Resource: "componentsets"}

var componentsetsKind = schema.GroupVersionKind{Group: "bundle.gke.io", Version: "v1alpha1", Kind: "ComponentSet"}

// Get takes name of the componentSet, and returns the corresponding componentSet object, and an error if there is any.
func (c *FakeComponentSets) Get(name string, options v1.GetOptions) (result *v1alpha1.ComponentSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(componentsetsResource, c.ns, name), &v1alpha1.ComponentSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComponentSet), err
}

// List takes label and field selectors, and returns the list of ComponentSets that match those selectors.
func (c *FakeComponentSets) List(opts v1.ListOptions) (result *v1alpha1.ComponentSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(componentsetsResource, componentsetsKind, c.ns, opts), &v1alpha1.ComponentSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComponentSetList{}
	for _, item := range obj.(*v1alpha1.ComponentSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested componentSets.
func (c *FakeComponentSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(componentsetsResource, c.ns, opts))

}

// Create takes the representation of a componentSet and creates it.  Returns the server's representation of the componentSet, and an error, if there is any.
func (c *FakeComponentSets) Create(componentSet *v1alpha1.ComponentSet) (result *v1alpha1.ComponentSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(componentsetsResource, c.ns, componentSet), &v1alpha1.ComponentSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComponentSet), err
}

// Update takes the representation of a componentSet and updates it. Returns the server's representation of the componentSet, and an error, if there is any.
func (c *FakeComponentSets) Update(componentSet *v1alpha1.ComponentSet) (result *v1alpha1.ComponentSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(componentsetsResource, c.ns, componentSet), &v1alpha1.ComponentSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComponentSet), err
}

// Delete takes name of the componentSet and deletes it. Returns an error if one occurs.
func (c *FakeComponentSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(componentsetsResource, c.ns, name), &v1alpha1.ComponentSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComponentSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(componentsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComponentSetList{})
	return err
}

// Patch applies the patch and returns the patched componentSet.
func (c *FakeComponentSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComponentSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(componentsetsResource, c.ns, name, data, subresources...), &v1alpha1.ComponentSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComponentSet), err
}