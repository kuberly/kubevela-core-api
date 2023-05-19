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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	scheme "github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ComponentDefinitionsGetter has a method to return a ComponentDefinitionInterface.
// A group's client should implement this interface.
type ComponentDefinitionsGetter interface {
	ComponentDefinitions(namespace string) ComponentDefinitionInterface
}

// ComponentDefinitionInterface has methods to work with ComponentDefinition resources.
type ComponentDefinitionInterface interface {
	Create(ctx context.Context, componentDefinition *v1beta1.ComponentDefinition, opts v1.CreateOptions) (*v1beta1.ComponentDefinition, error)
	Update(ctx context.Context, componentDefinition *v1beta1.ComponentDefinition, opts v1.UpdateOptions) (*v1beta1.ComponentDefinition, error)
	UpdateStatus(ctx context.Context, componentDefinition *v1beta1.ComponentDefinition, opts v1.UpdateOptions) (*v1beta1.ComponentDefinition, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ComponentDefinition, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ComponentDefinitionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComponentDefinition, err error)
	ComponentDefinitionExpansion
}

// componentDefinitions implements ComponentDefinitionInterface
type componentDefinitions struct {
	client rest.Interface
	ns     string
}

// newComponentDefinitions returns a ComponentDefinitions
func newComponentDefinitions(c *CoreV1beta1Client, namespace string) *componentDefinitions {
	return &componentDefinitions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the componentDefinition, and returns the corresponding componentDefinition object, and an error if there is any.
func (c *componentDefinitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComponentDefinition, err error) {
	result = &v1beta1.ComponentDefinition{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("componentdefinitions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComponentDefinitions that match those selectors.
func (c *componentDefinitions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComponentDefinitionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ComponentDefinitionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("componentdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested componentDefinitions.
func (c *componentDefinitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("componentdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a componentDefinition and creates it.  Returns the server's representation of the componentDefinition, and an error, if there is any.
func (c *componentDefinitions) Create(ctx context.Context, componentDefinition *v1beta1.ComponentDefinition, opts v1.CreateOptions) (result *v1beta1.ComponentDefinition, err error) {
	result = &v1beta1.ComponentDefinition{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("componentdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(componentDefinition).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a componentDefinition and updates it. Returns the server's representation of the componentDefinition, and an error, if there is any.
func (c *componentDefinitions) Update(ctx context.Context, componentDefinition *v1beta1.ComponentDefinition, opts v1.UpdateOptions) (result *v1beta1.ComponentDefinition, err error) {
	result = &v1beta1.ComponentDefinition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("componentdefinitions").
		Name(componentDefinition.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(componentDefinition).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *componentDefinitions) UpdateStatus(ctx context.Context, componentDefinition *v1beta1.ComponentDefinition, opts v1.UpdateOptions) (result *v1beta1.ComponentDefinition, err error) {
	result = &v1beta1.ComponentDefinition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("componentdefinitions").
		Name(componentDefinition.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(componentDefinition).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the componentDefinition and deletes it. Returns an error if one occurs.
func (c *componentDefinitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("componentdefinitions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *componentDefinitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("componentdefinitions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched componentDefinition.
func (c *componentDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComponentDefinition, err error) {
	result = &v1beta1.ComponentDefinition{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("componentdefinitions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
