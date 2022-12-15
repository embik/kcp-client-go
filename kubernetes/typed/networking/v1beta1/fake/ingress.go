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

package v1beta1

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsnetworkingv1beta1 "k8s.io/client-go/applyconfigurations/networking/v1beta1"
	networkingv1beta1client "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	"k8s.io/client-go/testing"

	kcpnetworkingv1beta1 "github.com/kcp-dev/client-go/kubernetes/typed/networking/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var ingressesResource = schema.GroupVersionResource{Group: "networking.k8s.io", Version: "v1beta1", Resource: "ingresses"}
var ingressesKind = schema.GroupVersionKind{Group: "networking.k8s.io", Version: "v1beta1", Kind: "Ingress"}

type ingressesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *ingressesClusterClient) Cluster(clusterPath logicalcluster.Path) kcpnetworkingv1beta1.IngressesNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &ingressesNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of Ingresses that match those selectors across all clusters.
func (c *ingressesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*networkingv1beta1.IngressList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(ingressesResource, ingressesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &networkingv1beta1.IngressList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkingv1beta1.IngressList{ListMeta: obj.(*networkingv1beta1.IngressList).ListMeta}
	for _, item := range obj.(*networkingv1beta1.IngressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested Ingresses across all clusters.
func (c *ingressesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(ingressesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type ingressesNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *ingressesNamespacer) Namespace(namespace string) networkingv1beta1client.IngressInterface {
	return &ingressesClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type ingressesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *ingressesClient) Create(ctx context.Context, ingress *networkingv1beta1.Ingress, opts metav1.CreateOptions) (*networkingv1beta1.Ingress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(ingressesResource, c.ClusterPath, c.Namespace, ingress), &networkingv1beta1.Ingress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1beta1.Ingress), err
}

func (c *ingressesClient) Update(ctx context.Context, ingress *networkingv1beta1.Ingress, opts metav1.UpdateOptions) (*networkingv1beta1.Ingress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(ingressesResource, c.ClusterPath, c.Namespace, ingress), &networkingv1beta1.Ingress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1beta1.Ingress), err
}

func (c *ingressesClient) UpdateStatus(ctx context.Context, ingress *networkingv1beta1.Ingress, opts metav1.UpdateOptions) (*networkingv1beta1.Ingress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(ingressesResource, c.ClusterPath, "status", c.Namespace, ingress), &networkingv1beta1.Ingress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1beta1.Ingress), err
}

func (c *ingressesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(ingressesResource, c.ClusterPath, c.Namespace, name, opts), &networkingv1beta1.Ingress{})
	return err
}

func (c *ingressesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(ingressesResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &networkingv1beta1.IngressList{})
	return err
}

func (c *ingressesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*networkingv1beta1.Ingress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(ingressesResource, c.ClusterPath, c.Namespace, name), &networkingv1beta1.Ingress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1beta1.Ingress), err
}

// List takes label and field selectors, and returns the list of Ingresses that match those selectors.
func (c *ingressesClient) List(ctx context.Context, opts metav1.ListOptions) (*networkingv1beta1.IngressList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(ingressesResource, ingressesKind, c.ClusterPath, c.Namespace, opts), &networkingv1beta1.IngressList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkingv1beta1.IngressList{ListMeta: obj.(*networkingv1beta1.IngressList).ListMeta}
	for _, item := range obj.(*networkingv1beta1.IngressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *ingressesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(ingressesResource, c.ClusterPath, c.Namespace, opts))
}

func (c *ingressesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*networkingv1beta1.Ingress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(ingressesResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &networkingv1beta1.Ingress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1beta1.Ingress), err
}

func (c *ingressesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsnetworkingv1beta1.IngressApplyConfiguration, opts metav1.ApplyOptions) (*networkingv1beta1.Ingress, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(ingressesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &networkingv1beta1.Ingress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1beta1.Ingress), err
}

func (c *ingressesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsnetworkingv1beta1.IngressApplyConfiguration, opts metav1.ApplyOptions) (*networkingv1beta1.Ingress, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(ingressesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &networkingv1beta1.Ingress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1beta1.Ingress), err
}
