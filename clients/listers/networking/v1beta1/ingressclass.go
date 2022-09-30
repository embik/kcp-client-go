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
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	networkingv1beta1listers "k8s.io/client-go/listers/networking/v1beta1"
	"k8s.io/client-go/tools/cache"
)

// IngressClassClusterLister can list IngressClasses across all workspaces, or scope down to a IngressClassLister for one workspace.
type IngressClassClusterLister interface {
	List(selector labels.Selector) (ret []*networkingv1beta1.IngressClass, err error)
	Cluster(cluster logicalcluster.Name) networkingv1beta1listers.IngressClassLister
}

type ingressClassClusterLister struct {
	indexer cache.Indexer
}

// NewIngressClassClusterLister returns a new IngressClassClusterLister.
func NewIngressClassClusterLister(indexer cache.Indexer) *ingressClassClusterLister {
	return &ingressClassClusterLister{indexer: indexer}
}

// List lists all IngressClasses in the indexer across all workspaces.
func (s *ingressClassClusterLister) List(selector labels.Selector) (ret []*networkingv1beta1.IngressClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*networkingv1beta1.IngressClass))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get IngressClasses.
func (s *ingressClassClusterLister) Cluster(cluster logicalcluster.Name) networkingv1beta1listers.IngressClassLister {
	return &ingressClassLister{indexer: s.indexer, cluster: cluster}
}

// ingressClassLister implements the networkingv1beta1listers.IngressClassLister interface.
type ingressClassLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all IngressClasses in the indexer for a workspace.
func (s *ingressClassLister) List(selector labels.Selector) (ret []*networkingv1beta1.IngressClass, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*networkingv1beta1.IngressClass)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}

	return ret, err
}

// Get retrieves the IngressClass from the indexer for a given workspace and name.
func (s *ingressClassLister) Get(name string) (*networkingv1beta1.IngressClass, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(networkingv1beta1.Resource("IngressClass"), name)
	}
	return obj.(*networkingv1beta1.IngressClass), nil
}
