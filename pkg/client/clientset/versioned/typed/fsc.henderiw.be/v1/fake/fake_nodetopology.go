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

// FakeNodeTopologies implements NodeTopologyInterface
type FakeNodeTopologies struct {
	Fake *FakeFscV1
	ns   string
}

var nodetopologiesResource = schema.GroupVersionResource{Group: "fsc.henderiw.be", Version: "v1", Resource: "nodetopologies"}

var nodetopologiesKind = schema.GroupVersionKind{Group: "fsc.henderiw.be", Version: "v1", Kind: "NodeTopology"}

// Get takes name of the nodeTopology, and returns the corresponding nodeTopology object, and an error if there is any.
func (c *FakeNodeTopologies) Get(ctx context.Context, name string, options v1.GetOptions) (result *fschenderiwbev1.NodeTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nodetopologiesResource, c.ns, name), &fschenderiwbev1.NodeTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fschenderiwbev1.NodeTopology), err
}

// List takes label and field selectors, and returns the list of NodeTopologies that match those selectors.
func (c *FakeNodeTopologies) List(ctx context.Context, opts v1.ListOptions) (result *fschenderiwbev1.NodeTopologyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nodetopologiesResource, nodetopologiesKind, c.ns, opts), &fschenderiwbev1.NodeTopologyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &fschenderiwbev1.NodeTopologyList{ListMeta: obj.(*fschenderiwbev1.NodeTopologyList).ListMeta}
	for _, item := range obj.(*fschenderiwbev1.NodeTopologyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeTopologies.
func (c *FakeNodeTopologies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nodetopologiesResource, c.ns, opts))

}

// Create takes the representation of a nodeTopology and creates it.  Returns the server's representation of the nodeTopology, and an error, if there is any.
func (c *FakeNodeTopologies) Create(ctx context.Context, nodeTopology *fschenderiwbev1.NodeTopology, opts v1.CreateOptions) (result *fschenderiwbev1.NodeTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nodetopologiesResource, c.ns, nodeTopology), &fschenderiwbev1.NodeTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fschenderiwbev1.NodeTopology), err
}

// Update takes the representation of a nodeTopology and updates it. Returns the server's representation of the nodeTopology, and an error, if there is any.
func (c *FakeNodeTopologies) Update(ctx context.Context, nodeTopology *fschenderiwbev1.NodeTopology, opts v1.UpdateOptions) (result *fschenderiwbev1.NodeTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nodetopologiesResource, c.ns, nodeTopology), &fschenderiwbev1.NodeTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fschenderiwbev1.NodeTopology), err
}

// Delete takes name of the nodeTopology and deletes it. Returns an error if one occurs.
func (c *FakeNodeTopologies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nodetopologiesResource, c.ns, name), &fschenderiwbev1.NodeTopology{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeTopologies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nodetopologiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &fschenderiwbev1.NodeTopologyList{})
	return err
}

// Patch applies the patch and returns the patched nodeTopology.
func (c *FakeNodeTopologies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *fschenderiwbev1.NodeTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodetopologiesResource, c.ns, name, pt, data, subresources...), &fschenderiwbev1.NodeTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fschenderiwbev1.NodeTopology), err
}
