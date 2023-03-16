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

package fake

import (
	"github.com/kcp-dev/logicalcluster/v3"

	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	"k8s.io/client-go/rest"

	kcpbatchv1 "github.com/kcp-dev/client-go/kubernetes/typed/batch/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpbatchv1.BatchV1ClusterInterface = (*BatchV1ClusterClient)(nil)

type BatchV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *BatchV1ClusterClient) Cluster(clusterPath logicalcluster.Path) batchv1.BatchV1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &BatchV1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *BatchV1ClusterClient) Jobs() kcpbatchv1.JobClusterInterface {
	return &jobsClusterClient{Fake: c.Fake}
}

func (c *BatchV1ClusterClient) CronJobs() kcpbatchv1.CronJobClusterInterface {
	return &cronJobsClusterClient{Fake: c.Fake}
}

var _ batchv1.BatchV1Interface = (*BatchV1Client)(nil)

type BatchV1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *BatchV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *BatchV1Client) Jobs(namespace string) batchv1.JobInterface {
	return &jobsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *BatchV1Client) CronJobs(namespace string) batchv1.CronJobInterface {
	return &cronJobsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
