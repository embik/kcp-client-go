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
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/third_party/informers"

	policyv1beta1 "k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	clientset "github.com/kcp-dev/client-go/clients/clientset/versioned"
	"github.com/kcp-dev/client-go/clients/informers/internalinterfaces"
	policyv1beta1listers "github.com/kcp-dev/client-go/clients/listers/policy/v1beta1"
)

// PodSecurityPolicyClusterInformer provides access to a shared informer and lister for
// PodSecurityPolicies.
type PodSecurityPolicyClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() policyv1beta1listers.PodSecurityPolicyClusterLister
}

type podSecurityPolicyClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPodSecurityPolicyClusterInformer constructs a new informer for PodSecurityPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodSecurityPolicyClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodSecurityPolicyClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPodSecurityPolicyClusterInformer constructs a new informer for PodSecurityPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodSecurityPolicyClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1beta1().PodSecurityPolicies().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1beta1().PodSecurityPolicies().Watch(context.TODO(), options)
			},
		},
		&policyv1beta1.PodSecurityPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *podSecurityPolicyClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodSecurityPolicyClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *podSecurityPolicyClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&policyv1beta1.PodSecurityPolicy{}, f.defaultInformer)
}

func (f *podSecurityPolicyClusterInformer) Lister() policyv1beta1listers.PodSecurityPolicyClusterLister {
	return policyv1beta1listers.NewPodSecurityPolicyClusterLister(f.Informer().GetIndexer())
}
