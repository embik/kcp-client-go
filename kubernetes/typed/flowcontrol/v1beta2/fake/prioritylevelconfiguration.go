//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package fake

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	flowcontrolv1beta2 "k8s.io/api/flowcontrol/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsflowcontrolv1beta2 "k8s.io/client-go/applyconfigurations/flowcontrol/v1beta2"
	flowcontrolv1beta2client "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var priorityLevelConfigurationsResource = schema.GroupVersionResource{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta2", Resource: "prioritylevelconfigurations"}
var priorityLevelConfigurationsKind = schema.GroupVersionKind{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta2", Kind: "PriorityLevelConfiguration"}

type priorityLevelConfigurationsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *priorityLevelConfigurationsClusterClient) Cluster(clusterPath logicalcluster.Path) flowcontrolv1beta2client.PriorityLevelConfigurationInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &priorityLevelConfigurationsClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of PriorityLevelConfigurations that match those selectors across all clusters.
func (c *priorityLevelConfigurationsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*flowcontrolv1beta2.PriorityLevelConfigurationList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(priorityLevelConfigurationsResource, priorityLevelConfigurationsKind, logicalcluster.Wildcard, opts), &flowcontrolv1beta2.PriorityLevelConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &flowcontrolv1beta2.PriorityLevelConfigurationList{ListMeta: obj.(*flowcontrolv1beta2.PriorityLevelConfigurationList).ListMeta}
	for _, item := range obj.(*flowcontrolv1beta2.PriorityLevelConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested PriorityLevelConfigurations across all clusters.
func (c *priorityLevelConfigurationsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(priorityLevelConfigurationsResource, logicalcluster.Wildcard, opts))
}

type priorityLevelConfigurationsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *priorityLevelConfigurationsClient) Create(ctx context.Context, priorityLevelConfiguration *flowcontrolv1beta2.PriorityLevelConfiguration, opts metav1.CreateOptions) (*flowcontrolv1beta2.PriorityLevelConfiguration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(priorityLevelConfigurationsResource, c.ClusterPath, priorityLevelConfiguration), &flowcontrolv1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*flowcontrolv1beta2.PriorityLevelConfiguration), err
}

func (c *priorityLevelConfigurationsClient) Update(ctx context.Context, priorityLevelConfiguration *flowcontrolv1beta2.PriorityLevelConfiguration, opts metav1.UpdateOptions) (*flowcontrolv1beta2.PriorityLevelConfiguration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(priorityLevelConfigurationsResource, c.ClusterPath, priorityLevelConfiguration), &flowcontrolv1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*flowcontrolv1beta2.PriorityLevelConfiguration), err
}

func (c *priorityLevelConfigurationsClient) UpdateStatus(ctx context.Context, priorityLevelConfiguration *flowcontrolv1beta2.PriorityLevelConfiguration, opts metav1.UpdateOptions) (*flowcontrolv1beta2.PriorityLevelConfiguration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(priorityLevelConfigurationsResource, c.ClusterPath, "status", priorityLevelConfiguration), &flowcontrolv1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*flowcontrolv1beta2.PriorityLevelConfiguration), err
}

func (c *priorityLevelConfigurationsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(priorityLevelConfigurationsResource, c.ClusterPath, name, opts), &flowcontrolv1beta2.PriorityLevelConfiguration{})
	return err
}

func (c *priorityLevelConfigurationsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(priorityLevelConfigurationsResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &flowcontrolv1beta2.PriorityLevelConfigurationList{})
	return err
}

func (c *priorityLevelConfigurationsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*flowcontrolv1beta2.PriorityLevelConfiguration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(priorityLevelConfigurationsResource, c.ClusterPath, name), &flowcontrolv1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*flowcontrolv1beta2.PriorityLevelConfiguration), err
}

// List takes label and field selectors, and returns the list of PriorityLevelConfigurations that match those selectors.
func (c *priorityLevelConfigurationsClient) List(ctx context.Context, opts metav1.ListOptions) (*flowcontrolv1beta2.PriorityLevelConfigurationList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(priorityLevelConfigurationsResource, priorityLevelConfigurationsKind, c.ClusterPath, opts), &flowcontrolv1beta2.PriorityLevelConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &flowcontrolv1beta2.PriorityLevelConfigurationList{ListMeta: obj.(*flowcontrolv1beta2.PriorityLevelConfigurationList).ListMeta}
	for _, item := range obj.(*flowcontrolv1beta2.PriorityLevelConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *priorityLevelConfigurationsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(priorityLevelConfigurationsResource, c.ClusterPath, opts))
}

func (c *priorityLevelConfigurationsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*flowcontrolv1beta2.PriorityLevelConfiguration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(priorityLevelConfigurationsResource, c.ClusterPath, name, pt, data, subresources...), &flowcontrolv1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*flowcontrolv1beta2.PriorityLevelConfiguration), err
}

func (c *priorityLevelConfigurationsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsflowcontrolv1beta2.PriorityLevelConfigurationApplyConfiguration, opts metav1.ApplyOptions) (*flowcontrolv1beta2.PriorityLevelConfiguration, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(priorityLevelConfigurationsResource, c.ClusterPath, *name, types.ApplyPatchType, data), &flowcontrolv1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*flowcontrolv1beta2.PriorityLevelConfiguration), err
}

func (c *priorityLevelConfigurationsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsflowcontrolv1beta2.PriorityLevelConfigurationApplyConfiguration, opts metav1.ApplyOptions) (*flowcontrolv1beta2.PriorityLevelConfiguration, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(priorityLevelConfigurationsResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &flowcontrolv1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*flowcontrolv1beta2.PriorityLevelConfiguration), err
}
