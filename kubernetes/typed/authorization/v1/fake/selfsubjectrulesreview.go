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

package v1

import (
	"context"

	"github.com/kcp-dev/logicalcluster/v3"

	authorizationv1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	authorizationv1client "k8s.io/client-go/kubernetes/typed/authorization/v1"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var selfSubjectRulesReviewsResource = schema.GroupVersionResource{Group: "authorization.k8s.io", Version: "v1", Resource: "selfsubjectrulesreviews"}
var selfSubjectRulesReviewsKind = schema.GroupVersionKind{Group: "authorization.k8s.io", Version: "v1", Kind: "SelfSubjectRulesReview"}

type selfSubjectRulesReviewsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *selfSubjectRulesReviewsClusterClient) Cluster(clusterPath logicalcluster.Path) authorizationv1client.SelfSubjectRulesReviewInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &selfSubjectRulesReviewsClient{Fake: c.Fake, ClusterPath: clusterPath}
}

type selfSubjectRulesReviewsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *selfSubjectRulesReviewsClient) Create(ctx context.Context, selfSubjectRulesReview *authorizationv1.SelfSubjectRulesReview, opts metav1.CreateOptions) (*authorizationv1.SelfSubjectRulesReview, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(selfSubjectRulesReviewsResource, c.ClusterPath, selfSubjectRulesReview), &authorizationv1.SelfSubjectRulesReview{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorizationv1.SelfSubjectRulesReview), err
}
