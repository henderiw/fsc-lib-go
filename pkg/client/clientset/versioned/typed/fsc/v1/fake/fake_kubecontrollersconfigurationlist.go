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

	fscv1 "github.com/henderiw/fsc-lib-go/pkg/apis/fsc/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubeControllersConfigurationLists implements KubeControllersConfigurationListInterface
type FakeKubeControllersConfigurationLists struct {
	Fake *FakeFscV1
	ns   string
}

var kubecontrollersconfigurationlistsResource = schema.GroupVersionResource{Group: "fsc.henderiw.be", Version: "v1", Resource: "kubecontrollersconfigurationlists"}

var kubecontrollersconfigurationlistsKind = schema.GroupVersionKind{Group: "fsc.henderiw.be", Version: "v1", Kind: "KubeControllersConfigurationList"}

// Get takes name of the kubeControllersConfigurationList, and returns the corresponding kubeControllersConfigurationList object, and an error if there is any.
func (c *FakeKubeControllersConfigurationLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *fscv1.KubeControllersConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubecontrollersconfigurationlistsResource, c.ns, name), &fscv1.KubeControllersConfigurationList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fscv1.KubeControllersConfigurationList), err
}

// List takes label and field selectors, and returns the list of KubeControllersConfigurationLists that match those selectors.
func (c *FakeKubeControllersConfigurationLists) List(ctx context.Context, opts v1.ListOptions) (result *fscv1.KubeControllersConfigurationListList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubecontrollersconfigurationlistsResource, kubecontrollersconfigurationlistsKind, c.ns, opts), &fscv1.KubeControllersConfigurationListList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fscv1.KubeControllersConfigurationListList), err
}

// Watch returns a watch.Interface that watches the requested kubeControllersConfigurationLists.
func (c *FakeKubeControllersConfigurationLists) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubecontrollersconfigurationlistsResource, c.ns, opts))

}

// Create takes the representation of a kubeControllersConfigurationList and creates it.  Returns the server's representation of the kubeControllersConfigurationList, and an error, if there is any.
func (c *FakeKubeControllersConfigurationLists) Create(ctx context.Context, kubeControllersConfigurationList *fscv1.KubeControllersConfigurationList, opts v1.CreateOptions) (result *fscv1.KubeControllersConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubecontrollersconfigurationlistsResource, c.ns, kubeControllersConfigurationList), &fscv1.KubeControllersConfigurationList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fscv1.KubeControllersConfigurationList), err
}

// Update takes the representation of a kubeControllersConfigurationList and updates it. Returns the server's representation of the kubeControllersConfigurationList, and an error, if there is any.
func (c *FakeKubeControllersConfigurationLists) Update(ctx context.Context, kubeControllersConfigurationList *fscv1.KubeControllersConfigurationList, opts v1.UpdateOptions) (result *fscv1.KubeControllersConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubecontrollersconfigurationlistsResource, c.ns, kubeControllersConfigurationList), &fscv1.KubeControllersConfigurationList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fscv1.KubeControllersConfigurationList), err
}

// Delete takes name of the kubeControllersConfigurationList and deletes it. Returns an error if one occurs.
func (c *FakeKubeControllersConfigurationLists) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kubecontrollersconfigurationlistsResource, c.ns, name), &fscv1.KubeControllersConfigurationList{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeControllersConfigurationLists) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubecontrollersconfigurationlistsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &fscv1.KubeControllersConfigurationListList{})
	return err
}

// Patch applies the patch and returns the patched kubeControllersConfigurationList.
func (c *FakeKubeControllersConfigurationLists) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *fscv1.KubeControllersConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubecontrollersconfigurationlistsResource, c.ns, name, pt, data, subresources...), &fscv1.KubeControllersConfigurationList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fscv1.KubeControllersConfigurationList), err
}