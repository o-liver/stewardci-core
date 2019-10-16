/*
#########################
#  SAP Steward-CI       #
#########################

THIS CODE IS GENERATED! DO NOT TOUCH!

Copyright SAP SE.

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
	v1alpha1 "github.com/SAP/stewardci-core/pkg/apis/steward/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTenants implements TenantInterface
type FakeTenants struct {
	Fake *FakeStewardV1alpha1
	ns   string
}

var tenantsResource = schema.GroupVersionResource{Group: "steward.sap.com", Version: "v1alpha1", Resource: "tenants"}

var tenantsKind = schema.GroupVersionKind{Group: "steward.sap.com", Version: "v1alpha1", Kind: "Tenant"}

// Get takes name of the tenant, and returns the corresponding tenant object, and an error if there is any.
func (c *FakeTenants) Get(name string, options v1.GetOptions) (result *v1alpha1.Tenant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tenantsResource, c.ns, name), &v1alpha1.Tenant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tenant), err
}

// List takes label and field selectors, and returns the list of Tenants that match those selectors.
func (c *FakeTenants) List(opts v1.ListOptions) (result *v1alpha1.TenantList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tenantsResource, tenantsKind, c.ns, opts), &v1alpha1.TenantList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TenantList{ListMeta: obj.(*v1alpha1.TenantList).ListMeta}
	for _, item := range obj.(*v1alpha1.TenantList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tenants.
func (c *FakeTenants) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tenantsResource, c.ns, opts))

}

// Create takes the representation of a tenant and creates it.  Returns the server's representation of the tenant, and an error, if there is any.
func (c *FakeTenants) Create(tenant *v1alpha1.Tenant) (result *v1alpha1.Tenant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tenantsResource, c.ns, tenant), &v1alpha1.Tenant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tenant), err
}

// Update takes the representation of a tenant and updates it. Returns the server's representation of the tenant, and an error, if there is any.
func (c *FakeTenants) Update(tenant *v1alpha1.Tenant) (result *v1alpha1.Tenant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tenantsResource, c.ns, tenant), &v1alpha1.Tenant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tenant), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTenants) UpdateStatus(tenant *v1alpha1.Tenant) (*v1alpha1.Tenant, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tenantsResource, "status", c.ns, tenant), &v1alpha1.Tenant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tenant), err
}

// Delete takes name of the tenant and deletes it. Returns an error if one occurs.
func (c *FakeTenants) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tenantsResource, c.ns, name), &v1alpha1.Tenant{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTenants) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tenantsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TenantList{})
	return err
}

// Patch applies the patch and returns the patched tenant.
func (c *FakeTenants) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Tenant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tenantsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Tenant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tenant), err
}
