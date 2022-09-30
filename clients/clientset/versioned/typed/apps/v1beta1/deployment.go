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

package v1beta1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	appsv1beta1 "k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	appsv1beta1client "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
)

// DeploymentsClusterGetter has a method to return a DeploymentsClusterInterface.
// A group's cluster client should implement this interface.
type DeploymentsClusterGetter interface {
	Deployments() DeploymentsClusterInterface
}

// DeploymentsClusterInterface can operate on Deployments across all clusters,
// or scope down to one cluster and return a DeploymentsNamespacer.
type DeploymentsClusterInterface interface {
	Cluster(logicalcluster.Name) DeploymentsNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*appsv1beta1.DeploymentList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type deploymentsClusterInterface struct {
	clientCache kcpclient.Cache[*appsv1beta1client.AppsV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *deploymentsClusterInterface) Cluster(name logicalcluster.Name) DeploymentsNamespacer {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &deploymentsNamespacer{clientCache: c.clientCache, name: name}
}

// List returns the entire collection of all Deployments across all clusters.
func (c *deploymentsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*appsv1beta1.DeploymentList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Deployments(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all Deployments across all clusters.
func (c *deploymentsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Deployments(metav1.NamespaceAll).Watch(ctx, opts)
}

// DeploymentsNamespacer can scope to objects within a namespace, returning a appsv1beta1client.DeploymentInterface.
type DeploymentsNamespacer interface {
	Namespace(string) appsv1beta1client.DeploymentInterface
}

type deploymentsNamespacer struct {
	clientCache kcpclient.Cache[*appsv1beta1client.AppsV1beta1Client]
	name        logicalcluster.Name
}

func (n *deploymentsNamespacer) Namespace(namespace string) appsv1beta1client.DeploymentInterface {
	return n.clientCache.ClusterOrDie(n.name).Deployments(namespace)
}
