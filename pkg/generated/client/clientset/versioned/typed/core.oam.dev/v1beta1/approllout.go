/*
Copyright 2021 The KubeVela Authors.

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

// AppRolloutsGetter has a method to return a AppRolloutInterface.
// A group's client should implement this interface.
type AppRolloutsGetter interface {
	AppRollouts(namespace string) AppRolloutInterface
}

// AppRolloutInterface has methods to work with AppRollout resources.
type AppRolloutInterface interface {
	Create(ctx context.Context, appRollout *v1beta1.AppRollout, opts v1.CreateOptions) (*v1beta1.AppRollout, error)
	Update(ctx context.Context, appRollout *v1beta1.AppRollout, opts v1.UpdateOptions) (*v1beta1.AppRollout, error)
	UpdateStatus(ctx context.Context, appRollout *v1beta1.AppRollout, opts v1.UpdateOptions) (*v1beta1.AppRollout, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.AppRollout, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.AppRolloutList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AppRollout, err error)
	AppRolloutExpansion
}

// appRollouts implements AppRolloutInterface
type appRollouts struct {
	client rest.Interface
	ns     string
}

// newAppRollouts returns a AppRollouts
func newAppRollouts(c *CoreV1beta1Client, namespace string) *appRollouts {
	return &appRollouts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the appRollout, and returns the corresponding appRollout object, and an error if there is any.
func (c *appRollouts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.AppRollout, err error) {
	result = &v1beta1.AppRollout{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("approllouts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppRollouts that match those selectors.
func (c *appRollouts) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.AppRolloutList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.AppRolloutList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("approllouts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appRollouts.
func (c *appRollouts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("approllouts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a appRollout and creates it.  Returns the server's representation of the appRollout, and an error, if there is any.
func (c *appRollouts) Create(ctx context.Context, appRollout *v1beta1.AppRollout, opts v1.CreateOptions) (result *v1beta1.AppRollout, err error) {
	result = &v1beta1.AppRollout{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("approllouts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(appRollout).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a appRollout and updates it. Returns the server's representation of the appRollout, and an error, if there is any.
func (c *appRollouts) Update(ctx context.Context, appRollout *v1beta1.AppRollout, opts v1.UpdateOptions) (result *v1beta1.AppRollout, err error) {
	result = &v1beta1.AppRollout{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("approllouts").
		Name(appRollout.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(appRollout).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *appRollouts) UpdateStatus(ctx context.Context, appRollout *v1beta1.AppRollout, opts v1.UpdateOptions) (result *v1beta1.AppRollout, err error) {
	result = &v1beta1.AppRollout{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("approllouts").
		Name(appRollout.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(appRollout).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the appRollout and deletes it. Returns an error if one occurs.
func (c *appRollouts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("approllouts").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *appRollouts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("approllouts").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched appRollout.
func (c *appRollouts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AppRollout, err error) {
	result = &v1beta1.AppRollout{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("approllouts").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}