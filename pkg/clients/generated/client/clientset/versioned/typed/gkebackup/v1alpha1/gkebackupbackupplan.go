// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/gkebackup/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GKEBackupBackupPlansGetter has a method to return a GKEBackupBackupPlanInterface.
// A group's client should implement this interface.
type GKEBackupBackupPlansGetter interface {
	GKEBackupBackupPlans(namespace string) GKEBackupBackupPlanInterface
}

// GKEBackupBackupPlanInterface has methods to work with GKEBackupBackupPlan resources.
type GKEBackupBackupPlanInterface interface {
	Create(ctx context.Context, gKEBackupBackupPlan *v1alpha1.GKEBackupBackupPlan, opts v1.CreateOptions) (*v1alpha1.GKEBackupBackupPlan, error)
	Update(ctx context.Context, gKEBackupBackupPlan *v1alpha1.GKEBackupBackupPlan, opts v1.UpdateOptions) (*v1alpha1.GKEBackupBackupPlan, error)
	UpdateStatus(ctx context.Context, gKEBackupBackupPlan *v1alpha1.GKEBackupBackupPlan, opts v1.UpdateOptions) (*v1alpha1.GKEBackupBackupPlan, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.GKEBackupBackupPlan, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.GKEBackupBackupPlanList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GKEBackupBackupPlan, err error)
	GKEBackupBackupPlanExpansion
}

// gKEBackupBackupPlans implements GKEBackupBackupPlanInterface
type gKEBackupBackupPlans struct {
	client rest.Interface
	ns     string
}

// newGKEBackupBackupPlans returns a GKEBackupBackupPlans
func newGKEBackupBackupPlans(c *GkebackupV1alpha1Client, namespace string) *gKEBackupBackupPlans {
	return &gKEBackupBackupPlans{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gKEBackupBackupPlan, and returns the corresponding gKEBackupBackupPlan object, and an error if there is any.
func (c *gKEBackupBackupPlans) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GKEBackupBackupPlan, err error) {
	result = &v1alpha1.GKEBackupBackupPlan{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gkebackupbackupplans").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GKEBackupBackupPlans that match those selectors.
func (c *gKEBackupBackupPlans) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GKEBackupBackupPlanList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GKEBackupBackupPlanList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gkebackupbackupplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gKEBackupBackupPlans.
func (c *gKEBackupBackupPlans) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gkebackupbackupplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a gKEBackupBackupPlan and creates it.  Returns the server's representation of the gKEBackupBackupPlan, and an error, if there is any.
func (c *gKEBackupBackupPlans) Create(ctx context.Context, gKEBackupBackupPlan *v1alpha1.GKEBackupBackupPlan, opts v1.CreateOptions) (result *v1alpha1.GKEBackupBackupPlan, err error) {
	result = &v1alpha1.GKEBackupBackupPlan{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gkebackupbackupplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gKEBackupBackupPlan).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a gKEBackupBackupPlan and updates it. Returns the server's representation of the gKEBackupBackupPlan, and an error, if there is any.
func (c *gKEBackupBackupPlans) Update(ctx context.Context, gKEBackupBackupPlan *v1alpha1.GKEBackupBackupPlan, opts v1.UpdateOptions) (result *v1alpha1.GKEBackupBackupPlan, err error) {
	result = &v1alpha1.GKEBackupBackupPlan{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gkebackupbackupplans").
		Name(gKEBackupBackupPlan.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gKEBackupBackupPlan).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *gKEBackupBackupPlans) UpdateStatus(ctx context.Context, gKEBackupBackupPlan *v1alpha1.GKEBackupBackupPlan, opts v1.UpdateOptions) (result *v1alpha1.GKEBackupBackupPlan, err error) {
	result = &v1alpha1.GKEBackupBackupPlan{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gkebackupbackupplans").
		Name(gKEBackupBackupPlan.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gKEBackupBackupPlan).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the gKEBackupBackupPlan and deletes it. Returns an error if one occurs.
func (c *gKEBackupBackupPlans) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gkebackupbackupplans").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gKEBackupBackupPlans) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gkebackupbackupplans").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched gKEBackupBackupPlan.
func (c *gKEBackupBackupPlans) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GKEBackupBackupPlan, err error) {
	result = &v1alpha1.GKEBackupBackupPlan{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gkebackupbackupplans").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}