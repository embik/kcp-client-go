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

	flowcontrolv1beta1 "k8s.io/api/flowcontrol/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	flowcontrolv1beta1listers "k8s.io/client-go/listers/flowcontrol/v1beta1"
	"k8s.io/client-go/tools/cache"
)

// PriorityLevelConfigurationClusterLister can list PriorityLevelConfigurations across all workspaces, or scope down to a PriorityLevelConfigurationLister for one workspace.
type PriorityLevelConfigurationClusterLister interface {
	List(selector labels.Selector) (ret []*flowcontrolv1beta1.PriorityLevelConfiguration, err error)
	Cluster(cluster logicalcluster.Name) flowcontrolv1beta1listers.PriorityLevelConfigurationLister
}

type priorityLevelConfigurationClusterLister struct {
	indexer cache.Indexer
}

// NewPriorityLevelConfigurationClusterLister returns a new PriorityLevelConfigurationClusterLister.
func NewPriorityLevelConfigurationClusterLister(indexer cache.Indexer) *priorityLevelConfigurationClusterLister {
	return &priorityLevelConfigurationClusterLister{indexer: indexer}
}

// List lists all PriorityLevelConfigurations in the indexer across all workspaces.
func (s *priorityLevelConfigurationClusterLister) List(selector labels.Selector) (ret []*flowcontrolv1beta1.PriorityLevelConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*flowcontrolv1beta1.PriorityLevelConfiguration))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get PriorityLevelConfigurations.
func (s *priorityLevelConfigurationClusterLister) Cluster(cluster logicalcluster.Name) flowcontrolv1beta1listers.PriorityLevelConfigurationLister {
	return &priorityLevelConfigurationLister{indexer: s.indexer, cluster: cluster}
}

// priorityLevelConfigurationLister implements the flowcontrolv1beta1listers.PriorityLevelConfigurationLister interface.
type priorityLevelConfigurationLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all PriorityLevelConfigurations in the indexer for a workspace.
func (s *priorityLevelConfigurationLister) List(selector labels.Selector) (ret []*flowcontrolv1beta1.PriorityLevelConfiguration, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*flowcontrolv1beta1.PriorityLevelConfiguration)
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

// Get retrieves the PriorityLevelConfiguration from the indexer for a given workspace and name.
func (s *priorityLevelConfigurationLister) Get(name string) (*flowcontrolv1beta1.PriorityLevelConfiguration, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(flowcontrolv1beta1.Resource("PriorityLevelConfiguration"), name)
	}
	return obj.(*flowcontrolv1beta1.PriorityLevelConfiguration), nil
}
