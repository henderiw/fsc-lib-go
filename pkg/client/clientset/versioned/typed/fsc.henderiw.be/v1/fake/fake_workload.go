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

	fschenderiwbev1 "github.com/fsc-demo-wim/fsc-lib-go/pkg/apis/fsc.henderiw.be/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWorkLoads implements WorkLoadInterface
type FakeWorkLoads struct {
	Fake *FakeFscV1
	ns   string
}

var workloadsResource = schema.GroupVersionResource{Group: "fsc.henderiw.be", Version: "v1", Resource: "workloads"}

var workloadsKind = schema.GroupVersionKind{Group: "fsc.henderiw.be", Version: "v1", Kind: "WorkLoad"}

// Get takes name of the workLoad, and returns the corresponding workLoad object, and an error if there is any.
func (c *FakeWorkLoads) Get(ctx context.Context, name string, options v1.GetOptions) (result *fschenderiwbev1.WorkLoad, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(workloadsResource, c.ns, name), &fschenderiwbev1.WorkLoad{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fschenderiwbev1.WorkLoad), err
}

// List takes label and field selectors, and returns the list of WorkLoads that match those selectors.
func (c *FakeWorkLoads) List(ctx context.Context, opts v1.ListOptions) (result *fschenderiwbev1.WorkLoadList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(workloadsResource, workloadsKind, c.ns, opts), &fschenderiwbev1.WorkLoadList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &fschenderiwbev1.WorkLoadList{ListMeta: obj.(*fschenderiwbev1.WorkLoadList).ListMeta}
	for _, item := range obj.(*fschenderiwbev1.WorkLoadList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workLoads.
func (c *FakeWorkLoads) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(workloadsResource, c.ns, opts))

}

// Create takes the representation of a workLoad and creates it.  Returns the server's representation of the workLoad, and an error, if there is any.
func (c *FakeWorkLoads) Create(ctx context.Context, workLoad *fschenderiwbev1.WorkLoad, opts v1.CreateOptions) (result *fschenderiwbev1.WorkLoad, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(workloadsResource, c.ns, workLoad), &fschenderiwbev1.WorkLoad{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fschenderiwbev1.WorkLoad), err
}

// Update takes the representation of a workLoad and updates it. Returns the server's representation of the workLoad, and an error, if there is any.
func (c *FakeWorkLoads) Update(ctx context.Context, workLoad *fschenderiwbev1.WorkLoad, opts v1.UpdateOptions) (result *fschenderiwbev1.WorkLoad, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(workloadsResource, c.ns, workLoad), &fschenderiwbev1.WorkLoad{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fschenderiwbev1.WorkLoad), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWorkLoads) UpdateStatus(ctx context.Context, workLoad *fschenderiwbev1.WorkLoad, opts v1.UpdateOptions) (*fschenderiwbev1.WorkLoad, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(workloadsResource, "status", c.ns, workLoad), &fschenderiwbev1.WorkLoad{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fschenderiwbev1.WorkLoad), err
}

// Delete takes name of the workLoad and deletes it. Returns an error if one occurs.
func (c *FakeWorkLoads) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(workloadsResource, c.ns, name), &fschenderiwbev1.WorkLoad{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkLoads) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(workloadsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &fschenderiwbev1.WorkLoadList{})
	return err
}

// Patch applies the patch and returns the patched workLoad.
func (c *FakeWorkLoads) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *fschenderiwbev1.WorkLoad, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(workloadsResource, c.ns, name, pt, data, subresources...), &fschenderiwbev1.WorkLoad{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fschenderiwbev1.WorkLoad), err
}
