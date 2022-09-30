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
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/third_party/informers"

	certificatesv1 "k8s.io/api/certificates/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	clientset "github.com/kcp-dev/client-go/clients/clientset/versioned"
	"github.com/kcp-dev/client-go/clients/informers/internalinterfaces"
	certificatesv1listers "github.com/kcp-dev/client-go/clients/listers/certificates/v1"
)

// CertificateSigningRequestClusterInformer provides access to a shared informer and lister for
// CertificateSigningRequests.
type CertificateSigningRequestClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() certificatesv1listers.CertificateSigningRequestClusterLister
}

type certificateSigningRequestClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCertificateSigningRequestClusterInformer constructs a new informer for CertificateSigningRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCertificateSigningRequestClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCertificateSigningRequestClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCertificateSigningRequestClusterInformer constructs a new informer for CertificateSigningRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCertificateSigningRequestClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertificatesV1().CertificateSigningRequests().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertificatesV1().CertificateSigningRequests().Watch(context.TODO(), options)
			},
		},
		&certificatesv1.CertificateSigningRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *certificateSigningRequestClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCertificateSigningRequestClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *certificateSigningRequestClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&certificatesv1.CertificateSigningRequest{}, f.defaultInformer)
}

func (f *certificateSigningRequestClusterInformer) Lister() certificatesv1listers.CertificateSigningRequestClusterLister {
	return certificatesv1listers.NewCertificateSigningRequestClusterLister(f.Informer().GetIndexer())
}
