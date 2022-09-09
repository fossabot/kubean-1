// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubean.io/api/apis/kubeanclusterconfig/v1alpha1"
)

// FakeKubeanClusterConfigs implements KubeanClusterConfigInterface
type FakeKubeanClusterConfigs struct {
	Fake *FakeKubeanV1alpha1
}

var kubeanclusterconfigsResource = schema.GroupVersionResource{Group: "kubean.io", Version: "v1alpha1", Resource: "kubeanclusterconfigs"}

var kubeanclusterconfigsKind = schema.GroupVersionKind{Group: "kubean.io", Version: "v1alpha1", Kind: "KubeanClusterConfig"}

// Get takes name of the kubeanClusterConfig, and returns the corresponding kubeanClusterConfig object, and an error if there is any.
func (c *FakeKubeanClusterConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KubeanClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kubeanclusterconfigsResource, name), &v1alpha1.KubeanClusterConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeanClusterConfig), err
}

// List takes label and field selectors, and returns the list of KubeanClusterConfigs that match those selectors.
func (c *FakeKubeanClusterConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KubeanClusterConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kubeanclusterconfigsResource, kubeanclusterconfigsKind, opts), &v1alpha1.KubeanClusterConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KubeanClusterConfigList{ListMeta: obj.(*v1alpha1.KubeanClusterConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.KubeanClusterConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeanClusterConfigs.
func (c *FakeKubeanClusterConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kubeanclusterconfigsResource, opts))
}

// Create takes the representation of a kubeanClusterConfig and creates it.  Returns the server's representation of the kubeanClusterConfig, and an error, if there is any.
func (c *FakeKubeanClusterConfigs) Create(ctx context.Context, kubeanClusterConfig *v1alpha1.KubeanClusterConfig, opts v1.CreateOptions) (result *v1alpha1.KubeanClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kubeanclusterconfigsResource, kubeanClusterConfig), &v1alpha1.KubeanClusterConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeanClusterConfig), err
}

// Update takes the representation of a kubeanClusterConfig and updates it. Returns the server's representation of the kubeanClusterConfig, and an error, if there is any.
func (c *FakeKubeanClusterConfigs) Update(ctx context.Context, kubeanClusterConfig *v1alpha1.KubeanClusterConfig, opts v1.UpdateOptions) (result *v1alpha1.KubeanClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kubeanclusterconfigsResource, kubeanClusterConfig), &v1alpha1.KubeanClusterConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeanClusterConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeanClusterConfigs) UpdateStatus(ctx context.Context, kubeanClusterConfig *v1alpha1.KubeanClusterConfig, opts v1.UpdateOptions) (*v1alpha1.KubeanClusterConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(kubeanclusterconfigsResource, "status", kubeanClusterConfig), &v1alpha1.KubeanClusterConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeanClusterConfig), err
}

// Delete takes name of the kubeanClusterConfig and deletes it. Returns an error if one occurs.
func (c *FakeKubeanClusterConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(kubeanclusterconfigsResource, name, opts), &v1alpha1.KubeanClusterConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeanClusterConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kubeanclusterconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KubeanClusterConfigList{})
	return err
}

// Patch applies the patch and returns the patched kubeanClusterConfig.
func (c *FakeKubeanClusterConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KubeanClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kubeanclusterconfigsResource, name, pt, data, subresources...), &v1alpha1.KubeanClusterConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubeanClusterConfig), err
}