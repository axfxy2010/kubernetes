// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	configv1 "github.com/openshift/api/config/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFeatureGates implements FeatureGateInterface
type FakeFeatureGates struct {
	Fake *FakeConfigV1
}

var featuregatesResource = schema.GroupVersionResource{Group: "config.openshift.io", Version: "v1", Resource: "featuregates"}

var featuregatesKind = schema.GroupVersionKind{Group: "config.openshift.io", Version: "v1", Kind: "FeatureGate"}

// Get takes name of the featureGate, and returns the corresponding featureGate object, and an error if there is any.
func (c *FakeFeatureGates) Get(name string, options v1.GetOptions) (result *configv1.FeatureGate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(featuregatesResource, name), &configv1.FeatureGate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.FeatureGate), err
}

// List takes label and field selectors, and returns the list of FeatureGates that match those selectors.
func (c *FakeFeatureGates) List(opts v1.ListOptions) (result *configv1.FeatureGateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(featuregatesResource, featuregatesKind, opts), &configv1.FeatureGateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &configv1.FeatureGateList{ListMeta: obj.(*configv1.FeatureGateList).ListMeta}
	for _, item := range obj.(*configv1.FeatureGateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested featureGates.
func (c *FakeFeatureGates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(featuregatesResource, opts))
}

// Create takes the representation of a featureGate and creates it.  Returns the server's representation of the featureGate, and an error, if there is any.
func (c *FakeFeatureGates) Create(featureGate *configv1.FeatureGate) (result *configv1.FeatureGate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(featuregatesResource, featureGate), &configv1.FeatureGate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.FeatureGate), err
}

// Update takes the representation of a featureGate and updates it. Returns the server's representation of the featureGate, and an error, if there is any.
func (c *FakeFeatureGates) Update(featureGate *configv1.FeatureGate) (result *configv1.FeatureGate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(featuregatesResource, featureGate), &configv1.FeatureGate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.FeatureGate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFeatureGates) UpdateStatus(featureGate *configv1.FeatureGate) (*configv1.FeatureGate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(featuregatesResource, "status", featureGate), &configv1.FeatureGate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.FeatureGate), err
}

// Delete takes name of the featureGate and deletes it. Returns an error if one occurs.
func (c *FakeFeatureGates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(featuregatesResource, name), &configv1.FeatureGate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFeatureGates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(featuregatesResource, listOptions)

	_, err := c.Fake.Invokes(action, &configv1.FeatureGateList{})
	return err
}

// Patch applies the patch and returns the patched featureGate.
func (c *FakeFeatureGates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *configv1.FeatureGate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(featuregatesResource, name, pt, data, subresources...), &configv1.FeatureGate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.FeatureGate), err
}
