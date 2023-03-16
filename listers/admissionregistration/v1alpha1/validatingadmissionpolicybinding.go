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

package v1alpha1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	admissionregistrationv1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	admissionregistrationv1alpha1listers "k8s.io/client-go/listers/admissionregistration/v1alpha1"
	"k8s.io/client-go/tools/cache"
)

// ValidatingAdmissionPolicyBindingClusterLister can list ValidatingAdmissionPolicyBindings across all workspaces, or scope down to a ValidatingAdmissionPolicyBindingLister for one workspace.
// All objects returned here must be treated as read-only.
type ValidatingAdmissionPolicyBindingClusterLister interface {
	// List lists all ValidatingAdmissionPolicyBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*admissionregistrationv1alpha1.ValidatingAdmissionPolicyBinding, err error)
	// Cluster returns a lister that can list and get ValidatingAdmissionPolicyBindings in one workspace.
	Cluster(clusterName logicalcluster.Name) admissionregistrationv1alpha1listers.ValidatingAdmissionPolicyBindingLister
	ValidatingAdmissionPolicyBindingClusterListerExpansion
}

type validatingAdmissionPolicyBindingClusterLister struct {
	indexer cache.Indexer
}

// NewValidatingAdmissionPolicyBindingClusterLister returns a new ValidatingAdmissionPolicyBindingClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewValidatingAdmissionPolicyBindingClusterLister(indexer cache.Indexer) *validatingAdmissionPolicyBindingClusterLister {
	return &validatingAdmissionPolicyBindingClusterLister{indexer: indexer}
}

// List lists all ValidatingAdmissionPolicyBindings in the indexer across all workspaces.
func (s *validatingAdmissionPolicyBindingClusterLister) List(selector labels.Selector) (ret []*admissionregistrationv1alpha1.ValidatingAdmissionPolicyBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*admissionregistrationv1alpha1.ValidatingAdmissionPolicyBinding))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get ValidatingAdmissionPolicyBindings.
func (s *validatingAdmissionPolicyBindingClusterLister) Cluster(clusterName logicalcluster.Name) admissionregistrationv1alpha1listers.ValidatingAdmissionPolicyBindingLister {
	return &validatingAdmissionPolicyBindingLister{indexer: s.indexer, clusterName: clusterName}
}

// validatingAdmissionPolicyBindingLister implements the admissionregistrationv1alpha1listers.ValidatingAdmissionPolicyBindingLister interface.
type validatingAdmissionPolicyBindingLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all ValidatingAdmissionPolicyBindings in the indexer for a workspace.
func (s *validatingAdmissionPolicyBindingLister) List(selector labels.Selector) (ret []*admissionregistrationv1alpha1.ValidatingAdmissionPolicyBinding, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*admissionregistrationv1alpha1.ValidatingAdmissionPolicyBinding))
	})
	return ret, err
}

// Get retrieves the ValidatingAdmissionPolicyBinding from the indexer for a given workspace and name.
func (s *validatingAdmissionPolicyBindingLister) Get(name string) (*admissionregistrationv1alpha1.ValidatingAdmissionPolicyBinding, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(admissionregistrationv1alpha1.Resource("validatingadmissionpolicybindings"), name)
	}
	return obj.(*admissionregistrationv1alpha1.ValidatingAdmissionPolicyBinding), nil
}
