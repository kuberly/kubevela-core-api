/*
Copyright 2023 The KubeVela Authors.

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

	v1beta1 "github.com/kuberly/kubevela-core-api/apis/core.oam.dev/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePolicyDefinitions implements PolicyDefinitionInterface
type FakePolicyDefinitions struct {
	Fake *FakeCoreV1beta1
	ns   string
}

var policydefinitionsResource = schema.GroupVersionResource{Group: "core.oam.dev", Version: "v1beta1", Resource: "policydefinitions"}

var policydefinitionsKind = schema.GroupVersionKind{Group: "core.oam.dev", Version: "v1beta1", Kind: "PolicyDefinition"}

// Get takes name of the policyDefinition, and returns the corresponding policyDefinition object, and an error if there is any.
func (c *FakePolicyDefinitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.PolicyDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(policydefinitionsResource, c.ns, name), &v1beta1.PolicyDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PolicyDefinition), err
}

// List takes label and field selectors, and returns the list of PolicyDefinitions that match those selectors.
func (c *FakePolicyDefinitions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PolicyDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(policydefinitionsResource, policydefinitionsKind, c.ns, opts), &v1beta1.PolicyDefinitionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PolicyDefinitionList{ListMeta: obj.(*v1beta1.PolicyDefinitionList).ListMeta}
	for _, item := range obj.(*v1beta1.PolicyDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested policyDefinitions.
func (c *FakePolicyDefinitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(policydefinitionsResource, c.ns, opts))

}

// Create takes the representation of a policyDefinition and creates it.  Returns the server's representation of the policyDefinition, and an error, if there is any.
func (c *FakePolicyDefinitions) Create(ctx context.Context, policyDefinition *v1beta1.PolicyDefinition, opts v1.CreateOptions) (result *v1beta1.PolicyDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(policydefinitionsResource, c.ns, policyDefinition), &v1beta1.PolicyDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PolicyDefinition), err
}

// Update takes the representation of a policyDefinition and updates it. Returns the server's representation of the policyDefinition, and an error, if there is any.
func (c *FakePolicyDefinitions) Update(ctx context.Context, policyDefinition *v1beta1.PolicyDefinition, opts v1.UpdateOptions) (result *v1beta1.PolicyDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(policydefinitionsResource, c.ns, policyDefinition), &v1beta1.PolicyDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PolicyDefinition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePolicyDefinitions) UpdateStatus(ctx context.Context, policyDefinition *v1beta1.PolicyDefinition, opts v1.UpdateOptions) (*v1beta1.PolicyDefinition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(policydefinitionsResource, "status", c.ns, policyDefinition), &v1beta1.PolicyDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PolicyDefinition), err
}

// Delete takes name of the policyDefinition and deletes it. Returns an error if one occurs.
func (c *FakePolicyDefinitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(policydefinitionsResource, c.ns, name, opts), &v1beta1.PolicyDefinition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePolicyDefinitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(policydefinitionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.PolicyDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched policyDefinition.
func (c *FakePolicyDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PolicyDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(policydefinitionsResource, c.ns, name, pt, data, subresources...), &v1beta1.PolicyDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PolicyDefinition), err
}
