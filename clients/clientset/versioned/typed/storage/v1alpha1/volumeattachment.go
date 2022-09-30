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

package v1alpha1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	storagev1alpha1 "k8s.io/api/storage/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	storagev1alpha1client "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
)

// VolumeAttachmentsClusterGetter has a method to return a VolumeAttachmentsClusterInterface.
// A group's cluster client should implement this interface.
type VolumeAttachmentsClusterGetter interface {
	VolumeAttachments() VolumeAttachmentsClusterInterface
}

// VolumeAttachmentsClusterInterface can operate on VolumeAttachments across all clusters,
// or scope down to one cluster and return a storagev1alpha1client.VolumeAttachmentInterface.
type VolumeAttachmentsClusterInterface interface {
	Cluster(logicalcluster.Name) storagev1alpha1client.VolumeAttachmentInterface
	List(ctx context.Context, opts metav1.ListOptions) (*storagev1alpha1.VolumeAttachmentList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type volumeAttachmentsClusterInterface struct {
	clientCache kcpclient.Cache[*storagev1alpha1client.StorageV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *volumeAttachmentsClusterInterface) Cluster(name logicalcluster.Name) storagev1alpha1client.VolumeAttachmentInterface {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(name).VolumeAttachments()
}

// List returns the entire collection of all VolumeAttachments across all clusters.
func (c *volumeAttachmentsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*storagev1alpha1.VolumeAttachmentList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).VolumeAttachments().List(ctx, opts)
}

// Watch begins to watch all VolumeAttachments across all clusters.
func (c *volumeAttachmentsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).VolumeAttachments().Watch(ctx, opts)
}
