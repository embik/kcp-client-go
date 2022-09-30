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
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	rbacv1listers "k8s.io/client-go/listers/rbac/v1"
	"k8s.io/client-go/tools/cache"
)

// RoleClusterLister can list Roles across all workspaces, or scope down to a RoleLister for one workspace.
type RoleClusterLister interface {
	List(selector labels.Selector) (ret []*rbacv1.Role, err error)
	Cluster(cluster logicalcluster.Name) rbacv1listers.RoleLister
}

type roleClusterLister struct {
	indexer cache.Indexer
}

// NewRoleClusterLister returns a new RoleClusterLister.
func NewRoleClusterLister(indexer cache.Indexer) *roleClusterLister {
	return &roleClusterLister{indexer: indexer}
}

// List lists all Roles in the indexer across all workspaces.
func (s *roleClusterLister) List(selector labels.Selector) (ret []*rbacv1.Role, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*rbacv1.Role))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get Roles.
func (s *roleClusterLister) Cluster(cluster logicalcluster.Name) rbacv1listers.RoleLister {
	return &roleLister{indexer: s.indexer, cluster: cluster}
}

// roleLister implements the rbacv1listers.RoleLister interface.
type roleLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all Roles in the indexer for a workspace.
func (s *roleLister) List(selector labels.Selector) (ret []*rbacv1.Role, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*rbacv1.Role)
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

// Roles returns an object that can list and get Roles in one namespace.
func (s *roleLister) Roles(namespace string) rbacv1listers.RoleNamespaceLister {
	return &roleNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// roleNamespaceLister implements the rbacv1listers.RoleNamespaceLister interface.
type roleNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all Roles in the indexer for a given workspace and namespace.
func (s *roleNamespaceLister) List(selector labels.Selector) (ret []*rbacv1.Role, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterAndNamespaceIndexName, kcpcache.ClusterAndNamespaceIndexKey(s.cluster, s.namespace))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*rbacv1.Role)
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

// Get retrieves the Role from the indexer for a given workspace, namespace and name.
func (s *roleNamespaceLister) Get(name string) (*rbacv1.Role, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(rbacv1.Resource("Role"), name)
	}
	return obj.(*rbacv1.Role), nil
}
