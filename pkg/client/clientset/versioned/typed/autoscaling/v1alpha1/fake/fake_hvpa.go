/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v1alpha1 "k8s.io/autoscaler/hvpa-controller/pkg/apis/autoscaling/v1alpha1"
	testing "k8s.io/client-go/testing"
)

// FakeHvpas implements HvpaInterface
type FakeHvpas struct {
	Fake *FakeAutoscalingV1alpha1
	ns   string
}

var hvpasResource = schema.GroupVersionResource{Group: "autoscaling.k8s.io", Version: "v1alpha1", Resource: "hvpas"}

var hvpasKind = schema.GroupVersionKind{Group: "autoscaling.k8s.io", Version: "v1alpha1", Kind: "Hvpa"}

// Get takes name of the hvpa, and returns the corresponding hvpa object, and an error if there is any.
func (c *FakeHvpas) Get(name string, options v1.GetOptions) (result *v1alpha1.Hvpa, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hvpasResource, c.ns, name), &v1alpha1.Hvpa{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hvpa), err
}

// List takes label and field selectors, and returns the list of Hvpas that match those selectors.
func (c *FakeHvpas) List(opts v1.ListOptions) (result *v1alpha1.HvpaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hvpasResource, hvpasKind, c.ns, opts), &v1alpha1.HvpaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HvpaList{ListMeta: obj.(*v1alpha1.HvpaList).ListMeta}
	for _, item := range obj.(*v1alpha1.HvpaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hvpas.
func (c *FakeHvpas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hvpasResource, c.ns, opts))

}

// Create takes the representation of a hvpa and creates it.  Returns the server's representation of the hvpa, and an error, if there is any.
func (c *FakeHvpas) Create(hvpa *v1alpha1.Hvpa) (result *v1alpha1.Hvpa, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hvpasResource, c.ns, hvpa), &v1alpha1.Hvpa{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hvpa), err
}

// Update takes the representation of a hvpa and updates it. Returns the server's representation of the hvpa, and an error, if there is any.
func (c *FakeHvpas) Update(hvpa *v1alpha1.Hvpa) (result *v1alpha1.Hvpa, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hvpasResource, c.ns, hvpa), &v1alpha1.Hvpa{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hvpa), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHvpas) UpdateStatus(hvpa *v1alpha1.Hvpa) (*v1alpha1.Hvpa, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hvpasResource, "status", c.ns, hvpa), &v1alpha1.Hvpa{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hvpa), err
}

// Delete takes name of the hvpa and deletes it. Returns an error if one occurs.
func (c *FakeHvpas) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hvpasResource, c.ns, name), &v1alpha1.Hvpa{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHvpas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hvpasResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.HvpaList{})
	return err
}

// Patch applies the patch and returns the patched hvpa.
func (c *FakeHvpas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Hvpa, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hvpasResource, c.ns, name, pt, data, subresources...), &v1alpha1.Hvpa{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hvpa), err
}
