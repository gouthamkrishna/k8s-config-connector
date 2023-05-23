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

package fake

import (
	"context"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dns/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDNSResponsePolicyRules implements DNSResponsePolicyRuleInterface
type FakeDNSResponsePolicyRules struct {
	Fake *FakeDnsV1alpha1
	ns   string
}

var dnsresponsepolicyrulesResource = schema.GroupVersionResource{Group: "dns.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "dnsresponsepolicyrules"}

var dnsresponsepolicyrulesKind = schema.GroupVersionKind{Group: "dns.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DNSResponsePolicyRule"}

// Get takes name of the dNSResponsePolicyRule, and returns the corresponding dNSResponsePolicyRule object, and an error if there is any.
func (c *FakeDNSResponsePolicyRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DNSResponsePolicyRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dnsresponsepolicyrulesResource, c.ns, name), &v1alpha1.DNSResponsePolicyRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSResponsePolicyRule), err
}

// List takes label and field selectors, and returns the list of DNSResponsePolicyRules that match those selectors.
func (c *FakeDNSResponsePolicyRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DNSResponsePolicyRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dnsresponsepolicyrulesResource, dnsresponsepolicyrulesKind, c.ns, opts), &v1alpha1.DNSResponsePolicyRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DNSResponsePolicyRuleList{ListMeta: obj.(*v1alpha1.DNSResponsePolicyRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.DNSResponsePolicyRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dNSResponsePolicyRules.
func (c *FakeDNSResponsePolicyRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dnsresponsepolicyrulesResource, c.ns, opts))

}

// Create takes the representation of a dNSResponsePolicyRule and creates it.  Returns the server's representation of the dNSResponsePolicyRule, and an error, if there is any.
func (c *FakeDNSResponsePolicyRules) Create(ctx context.Context, dNSResponsePolicyRule *v1alpha1.DNSResponsePolicyRule, opts v1.CreateOptions) (result *v1alpha1.DNSResponsePolicyRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dnsresponsepolicyrulesResource, c.ns, dNSResponsePolicyRule), &v1alpha1.DNSResponsePolicyRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSResponsePolicyRule), err
}

// Update takes the representation of a dNSResponsePolicyRule and updates it. Returns the server's representation of the dNSResponsePolicyRule, and an error, if there is any.
func (c *FakeDNSResponsePolicyRules) Update(ctx context.Context, dNSResponsePolicyRule *v1alpha1.DNSResponsePolicyRule, opts v1.UpdateOptions) (result *v1alpha1.DNSResponsePolicyRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dnsresponsepolicyrulesResource, c.ns, dNSResponsePolicyRule), &v1alpha1.DNSResponsePolicyRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSResponsePolicyRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDNSResponsePolicyRules) UpdateStatus(ctx context.Context, dNSResponsePolicyRule *v1alpha1.DNSResponsePolicyRule, opts v1.UpdateOptions) (*v1alpha1.DNSResponsePolicyRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dnsresponsepolicyrulesResource, "status", c.ns, dNSResponsePolicyRule), &v1alpha1.DNSResponsePolicyRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSResponsePolicyRule), err
}

// Delete takes name of the dNSResponsePolicyRule and deletes it. Returns an error if one occurs.
func (c *FakeDNSResponsePolicyRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dnsresponsepolicyrulesResource, c.ns, name, opts), &v1alpha1.DNSResponsePolicyRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDNSResponsePolicyRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dnsresponsepolicyrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DNSResponsePolicyRuleList{})
	return err
}

// Patch applies the patch and returns the patched dNSResponsePolicyRule.
func (c *FakeDNSResponsePolicyRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DNSResponsePolicyRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dnsresponsepolicyrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DNSResponsePolicyRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSResponsePolicyRule), err
}