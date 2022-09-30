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
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	schedulingv1alpha1 "k8s.io/api/scheduling/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	schedulingv1alpha1listers "k8s.io/client-go/listers/scheduling/v1alpha1"
	"k8s.io/client-go/tools/cache"
)

// PriorityClassClusterLister can list PriorityClasses across all workspaces, or scope down to a PriorityClassLister for one workspace.
type PriorityClassClusterLister interface {
	List(selector labels.Selector) (ret []*schedulingv1alpha1.PriorityClass, err error)
	Cluster(cluster logicalcluster.Name) schedulingv1alpha1listers.PriorityClassLister
}

type priorityClassClusterLister struct {
	indexer cache.Indexer
}

// NewPriorityClassClusterLister returns a new PriorityClassClusterLister.
func NewPriorityClassClusterLister(indexer cache.Indexer) *priorityClassClusterLister {
	return &priorityClassClusterLister{indexer: indexer}
}

// List lists all PriorityClasses in the indexer across all workspaces.
func (s *priorityClassClusterLister) List(selector labels.Selector) (ret []*schedulingv1alpha1.PriorityClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*schedulingv1alpha1.PriorityClass))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get PriorityClasses.
func (s *priorityClassClusterLister) Cluster(cluster logicalcluster.Name) schedulingv1alpha1listers.PriorityClassLister {
	return &priorityClassLister{indexer: s.indexer, cluster: cluster}
}

// priorityClassLister implements the schedulingv1alpha1listers.PriorityClassLister interface.
type priorityClassLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all PriorityClasses in the indexer for a workspace.
func (s *priorityClassLister) List(selector labels.Selector) (ret []*schedulingv1alpha1.PriorityClass, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*schedulingv1alpha1.PriorityClass)
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

// Get retrieves the PriorityClass from the indexer for a given workspace and name.
func (s *priorityClassLister) Get(name string) (*schedulingv1alpha1.PriorityClass, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schedulingv1alpha1.Resource("PriorityClass"), name)
	}
	return obj.(*schedulingv1alpha1.PriorityClass), nil
}
