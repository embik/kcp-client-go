//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The KCP Authors.

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

package v1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	schedulingv1 "k8s.io/api/scheduling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	schedulingv1client "k8s.io/client-go/kubernetes/typed/scheduling/v1"
)

// PriorityClassesClusterGetter has a method to return a PriorityClassesClusterInterface.
// A group's cluster client should implement this interface.
type PriorityClassesClusterGetter interface {
	PriorityClasses() PriorityClassesClusterInterface
}

// PriorityClassesClusterInterface can operate on PriorityClasses across all clusters,
// or scope down to one cluster and return a schedulingv1client.PriorityClassInterface.
type PriorityClassesClusterInterface interface {
	Cluster(logicalcluster.Name) schedulingv1client.PriorityClassInterface
	List(ctx context.Context, opts metav1.ListOptions) (*schedulingv1.PriorityClassList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type priorityClassesClusterInterface struct {
	clientCache kcpclient.Cache[*schedulingv1client.SchedulingV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *priorityClassesClusterInterface) Cluster(name logicalcluster.Name) schedulingv1client.PriorityClassInterface {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(name).PriorityClasses()
}

// List returns the entire collection of all PriorityClasses across all clusters.
func (c *priorityClassesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*schedulingv1.PriorityClassList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).PriorityClasses().List(ctx, opts)
}

// Watch begins to watch all PriorityClasses across all clusters.
func (c *priorityClassesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).PriorityClasses().Watch(ctx, opts)
}
