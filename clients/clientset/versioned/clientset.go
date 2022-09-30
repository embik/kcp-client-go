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

package kubernetes

import (
	"fmt"
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	"k8s.io/client-go/discovery"
	client "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"

	admissionregistrationv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/admissionregistration/v1"
	admissionregistrationv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/admissionregistration/v1beta1"
	internalv1alpha1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/apiserverinternal/v1alpha1"
	appsv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/apps/v1"
	appsv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/apps/v1beta1"
	appsv1beta2 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/apps/v1beta2"
	authenticationv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/authentication/v1"
	authenticationv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/authentication/v1beta1"
	authorizationv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/authorization/v1"
	authorizationv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/authorization/v1beta1"
	autoscalingv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/autoscaling/v1"
	autoscalingv2 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/autoscaling/v2"
	autoscalingv2beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/autoscaling/v2beta1"
	autoscalingv2beta2 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/autoscaling/v2beta2"
	batchv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/batch/v1"
	batchv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/batch/v1beta1"
	certificatesv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/certificates/v1"
	certificatesv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/certificates/v1beta1"
	coordinationv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/coordination/v1"
	coordinationv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/coordination/v1beta1"
	corev1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/core/v1"
	discoveryv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/discovery/v1"
	discoveryv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/discovery/v1beta1"
	eventsv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/events/v1"
	eventsv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/events/v1beta1"
	extensionsv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/extensions/v1beta1"
	flowcontrolv1alpha1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/flowcontrol/v1alpha1"
	flowcontrolv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/flowcontrol/v1beta1"
	flowcontrolv1beta2 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/flowcontrol/v1beta2"
	networkingv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/networking/v1"
	networkingv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/networking/v1beta1"
	nodev1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/node/v1"
	nodev1alpha1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/node/v1alpha1"
	nodev1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/node/v1beta1"
	policyv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/policy/v1"
	policyv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/policy/v1beta1"
	rbacv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/rbac/v1"
	rbacv1alpha1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/rbac/v1alpha1"
	rbacv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/rbac/v1beta1"
	schedulingv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/scheduling/v1"
	schedulingv1alpha1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/scheduling/v1alpha1"
	schedulingv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/scheduling/v1beta1"
	storagev1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/storage/v1"
	storagev1alpha1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/storage/v1alpha1"
	storagev1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/storage/v1beta1"
)

type ClusterInterface interface {
	Cluster(logicalcluster.Name) client.Interface
	Discovery() discovery.DiscoveryInterface
	AdmissionregistrationV1() admissionregistrationv1.AdmissionregistrationV1ClusterInterface
	AdmissionregistrationV1beta1() admissionregistrationv1beta1.AdmissionregistrationV1beta1ClusterInterface
	AppsV1() appsv1.AppsV1ClusterInterface
	AppsV1beta1() appsv1beta1.AppsV1beta1ClusterInterface
	AppsV1beta2() appsv1beta2.AppsV1beta2ClusterInterface
	AuthenticationV1() authenticationv1.AuthenticationV1ClusterInterface
	AuthenticationV1beta1() authenticationv1beta1.AuthenticationV1beta1ClusterInterface
	AuthorizationV1() authorizationv1.AuthorizationV1ClusterInterface
	AuthorizationV1beta1() authorizationv1beta1.AuthorizationV1beta1ClusterInterface
	AutoscalingV1() autoscalingv1.AutoscalingV1ClusterInterface
	AutoscalingV2() autoscalingv2.AutoscalingV2ClusterInterface
	AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1ClusterInterface
	AutoscalingV2beta2() autoscalingv2beta2.AutoscalingV2beta2ClusterInterface
	BatchV1() batchv1.BatchV1ClusterInterface
	BatchV1beta1() batchv1beta1.BatchV1beta1ClusterInterface
	CertificatesV1() certificatesv1.CertificatesV1ClusterInterface
	CertificatesV1beta1() certificatesv1beta1.CertificatesV1beta1ClusterInterface
	CoordinationV1() coordinationv1.CoordinationV1ClusterInterface
	CoordinationV1beta1() coordinationv1beta1.CoordinationV1beta1ClusterInterface
	CoreV1() corev1.CoreV1ClusterInterface
	DiscoveryV1() discoveryv1.DiscoveryV1ClusterInterface
	DiscoveryV1beta1() discoveryv1beta1.DiscoveryV1beta1ClusterInterface
	EventsV1() eventsv1.EventsV1ClusterInterface
	EventsV1beta1() eventsv1beta1.EventsV1beta1ClusterInterface
	ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1ClusterInterface
	FlowcontrolV1alpha1() flowcontrolv1alpha1.FlowcontrolV1alpha1ClusterInterface
	FlowcontrolV1beta1() flowcontrolv1beta1.FlowcontrolV1beta1ClusterInterface
	FlowcontrolV1beta2() flowcontrolv1beta2.FlowcontrolV1beta2ClusterInterface
	InternalV1alpha1() internalv1alpha1.InternalV1alpha1ClusterInterface
	NetworkingV1() networkingv1.NetworkingV1ClusterInterface
	NetworkingV1beta1() networkingv1beta1.NetworkingV1beta1ClusterInterface
	NodeV1() nodev1.NodeV1ClusterInterface
	NodeV1alpha1() nodev1alpha1.NodeV1alpha1ClusterInterface
	NodeV1beta1() nodev1beta1.NodeV1beta1ClusterInterface
	PolicyV1() policyv1.PolicyV1ClusterInterface
	PolicyV1beta1() policyv1beta1.PolicyV1beta1ClusterInterface
	RbacV1() rbacv1.RbacV1ClusterInterface
	RbacV1alpha1() rbacv1alpha1.RbacV1alpha1ClusterInterface
	RbacV1beta1() rbacv1beta1.RbacV1beta1ClusterInterface
	SchedulingV1() schedulingv1.SchedulingV1ClusterInterface
	SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1ClusterInterface
	SchedulingV1beta1() schedulingv1beta1.SchedulingV1beta1ClusterInterface
	StorageV1() storagev1.StorageV1ClusterInterface
	StorageV1alpha1() storagev1alpha1.StorageV1alpha1ClusterInterface
	StorageV1beta1() storagev1beta1.StorageV1beta1ClusterInterface
}

// ClusterClientset contains the clients for groups.
type ClusterClientset struct {
	*discovery.DiscoveryClient
	clientCache                  kcpclient.Cache[*client.Clientset]
	admissionregistrationV1      *admissionregistrationv1.AdmissionregistrationV1ClusterClient
	admissionregistrationV1beta1 *admissionregistrationv1beta1.AdmissionregistrationV1beta1ClusterClient
	appsV1                       *appsv1.AppsV1ClusterClient
	appsV1beta1                  *appsv1beta1.AppsV1beta1ClusterClient
	appsV1beta2                  *appsv1beta2.AppsV1beta2ClusterClient
	authenticationV1             *authenticationv1.AuthenticationV1ClusterClient
	authenticationV1beta1        *authenticationv1beta1.AuthenticationV1beta1ClusterClient
	authorizationV1              *authorizationv1.AuthorizationV1ClusterClient
	authorizationV1beta1         *authorizationv1beta1.AuthorizationV1beta1ClusterClient
	autoscalingV1                *autoscalingv1.AutoscalingV1ClusterClient
	autoscalingV2                *autoscalingv2.AutoscalingV2ClusterClient
	autoscalingV2beta1           *autoscalingv2beta1.AutoscalingV2beta1ClusterClient
	autoscalingV2beta2           *autoscalingv2beta2.AutoscalingV2beta2ClusterClient
	batchV1                      *batchv1.BatchV1ClusterClient
	batchV1beta1                 *batchv1beta1.BatchV1beta1ClusterClient
	certificatesV1               *certificatesv1.CertificatesV1ClusterClient
	certificatesV1beta1          *certificatesv1beta1.CertificatesV1beta1ClusterClient
	coordinationV1               *coordinationv1.CoordinationV1ClusterClient
	coordinationV1beta1          *coordinationv1beta1.CoordinationV1beta1ClusterClient
	coreV1                       *corev1.CoreV1ClusterClient
	discoveryV1                  *discoveryv1.DiscoveryV1ClusterClient
	discoveryV1beta1             *discoveryv1beta1.DiscoveryV1beta1ClusterClient
	eventsV1                     *eventsv1.EventsV1ClusterClient
	eventsV1beta1                *eventsv1beta1.EventsV1beta1ClusterClient
	extensionsV1beta1            *extensionsv1beta1.ExtensionsV1beta1ClusterClient
	flowcontrolV1alpha1          *flowcontrolv1alpha1.FlowcontrolV1alpha1ClusterClient
	flowcontrolV1beta1           *flowcontrolv1beta1.FlowcontrolV1beta1ClusterClient
	flowcontrolV1beta2           *flowcontrolv1beta2.FlowcontrolV1beta2ClusterClient
	internalV1alpha1             *internalv1alpha1.InternalV1alpha1ClusterClient
	networkingV1                 *networkingv1.NetworkingV1ClusterClient
	networkingV1beta1            *networkingv1beta1.NetworkingV1beta1ClusterClient
	nodeV1                       *nodev1.NodeV1ClusterClient
	nodeV1alpha1                 *nodev1alpha1.NodeV1alpha1ClusterClient
	nodeV1beta1                  *nodev1beta1.NodeV1beta1ClusterClient
	policyV1                     *policyv1.PolicyV1ClusterClient
	policyV1beta1                *policyv1beta1.PolicyV1beta1ClusterClient
	rbacV1                       *rbacv1.RbacV1ClusterClient
	rbacV1alpha1                 *rbacv1alpha1.RbacV1alpha1ClusterClient
	rbacV1beta1                  *rbacv1beta1.RbacV1beta1ClusterClient
	schedulingV1                 *schedulingv1.SchedulingV1ClusterClient
	schedulingV1alpha1           *schedulingv1alpha1.SchedulingV1alpha1ClusterClient
	schedulingV1beta1            *schedulingv1beta1.SchedulingV1beta1ClusterClient
	storageV1                    *storagev1.StorageV1ClusterClient
	storageV1alpha1              *storagev1alpha1.StorageV1alpha1ClusterClient
	storageV1beta1               *storagev1beta1.StorageV1beta1ClusterClient
}

// Discovery retrieves the DiscoveryClient
func (c *ClusterClientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// AdmissionregistrationV1 retrieves the AdmissionregistrationV1ClusterClient.
func (c *ClusterClientset) AdmissionregistrationV1() admissionregistrationv1.AdmissionregistrationV1ClusterInterface {
	return c.admissionregistrationV1
}

// AdmissionregistrationV1beta1 retrieves the AdmissionregistrationV1beta1ClusterClient.
func (c *ClusterClientset) AdmissionregistrationV1beta1() admissionregistrationv1beta1.AdmissionregistrationV1beta1ClusterInterface {
	return c.admissionregistrationV1beta1
}

// AppsV1 retrieves the AppsV1ClusterClient.
func (c *ClusterClientset) AppsV1() appsv1.AppsV1ClusterInterface {
	return c.appsV1
}

// AppsV1beta1 retrieves the AppsV1beta1ClusterClient.
func (c *ClusterClientset) AppsV1beta1() appsv1beta1.AppsV1beta1ClusterInterface {
	return c.appsV1beta1
}

// AppsV1beta2 retrieves the AppsV1beta2ClusterClient.
func (c *ClusterClientset) AppsV1beta2() appsv1beta2.AppsV1beta2ClusterInterface {
	return c.appsV1beta2
}

// AuthenticationV1 retrieves the AuthenticationV1ClusterClient.
func (c *ClusterClientset) AuthenticationV1() authenticationv1.AuthenticationV1ClusterInterface {
	return c.authenticationV1
}

// AuthenticationV1beta1 retrieves the AuthenticationV1beta1ClusterClient.
func (c *ClusterClientset) AuthenticationV1beta1() authenticationv1beta1.AuthenticationV1beta1ClusterInterface {
	return c.authenticationV1beta1
}

// AuthorizationV1 retrieves the AuthorizationV1ClusterClient.
func (c *ClusterClientset) AuthorizationV1() authorizationv1.AuthorizationV1ClusterInterface {
	return c.authorizationV1
}

// AuthorizationV1beta1 retrieves the AuthorizationV1beta1ClusterClient.
func (c *ClusterClientset) AuthorizationV1beta1() authorizationv1beta1.AuthorizationV1beta1ClusterInterface {
	return c.authorizationV1beta1
}

// AutoscalingV1 retrieves the AutoscalingV1ClusterClient.
func (c *ClusterClientset) AutoscalingV1() autoscalingv1.AutoscalingV1ClusterInterface {
	return c.autoscalingV1
}

// AutoscalingV2 retrieves the AutoscalingV2ClusterClient.
func (c *ClusterClientset) AutoscalingV2() autoscalingv2.AutoscalingV2ClusterInterface {
	return c.autoscalingV2
}

// AutoscalingV2beta1 retrieves the AutoscalingV2beta1ClusterClient.
func (c *ClusterClientset) AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1ClusterInterface {
	return c.autoscalingV2beta1
}

// AutoscalingV2beta2 retrieves the AutoscalingV2beta2ClusterClient.
func (c *ClusterClientset) AutoscalingV2beta2() autoscalingv2beta2.AutoscalingV2beta2ClusterInterface {
	return c.autoscalingV2beta2
}

// BatchV1 retrieves the BatchV1ClusterClient.
func (c *ClusterClientset) BatchV1() batchv1.BatchV1ClusterInterface {
	return c.batchV1
}

// BatchV1beta1 retrieves the BatchV1beta1ClusterClient.
func (c *ClusterClientset) BatchV1beta1() batchv1beta1.BatchV1beta1ClusterInterface {
	return c.batchV1beta1
}

// CertificatesV1 retrieves the CertificatesV1ClusterClient.
func (c *ClusterClientset) CertificatesV1() certificatesv1.CertificatesV1ClusterInterface {
	return c.certificatesV1
}

// CertificatesV1beta1 retrieves the CertificatesV1beta1ClusterClient.
func (c *ClusterClientset) CertificatesV1beta1() certificatesv1beta1.CertificatesV1beta1ClusterInterface {
	return c.certificatesV1beta1
}

// CoordinationV1 retrieves the CoordinationV1ClusterClient.
func (c *ClusterClientset) CoordinationV1() coordinationv1.CoordinationV1ClusterInterface {
	return c.coordinationV1
}

// CoordinationV1beta1 retrieves the CoordinationV1beta1ClusterClient.
func (c *ClusterClientset) CoordinationV1beta1() coordinationv1beta1.CoordinationV1beta1ClusterInterface {
	return c.coordinationV1beta1
}

// CoreV1 retrieves the CoreV1ClusterClient.
func (c *ClusterClientset) CoreV1() corev1.CoreV1ClusterInterface {
	return c.coreV1
}

// DiscoveryV1 retrieves the DiscoveryV1ClusterClient.
func (c *ClusterClientset) DiscoveryV1() discoveryv1.DiscoveryV1ClusterInterface {
	return c.discoveryV1
}

// DiscoveryV1beta1 retrieves the DiscoveryV1beta1ClusterClient.
func (c *ClusterClientset) DiscoveryV1beta1() discoveryv1beta1.DiscoveryV1beta1ClusterInterface {
	return c.discoveryV1beta1
}

// EventsV1 retrieves the EventsV1ClusterClient.
func (c *ClusterClientset) EventsV1() eventsv1.EventsV1ClusterInterface {
	return c.eventsV1
}

// EventsV1beta1 retrieves the EventsV1beta1ClusterClient.
func (c *ClusterClientset) EventsV1beta1() eventsv1beta1.EventsV1beta1ClusterInterface {
	return c.eventsV1beta1
}

// ExtensionsV1beta1 retrieves the ExtensionsV1beta1ClusterClient.
func (c *ClusterClientset) ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1ClusterInterface {
	return c.extensionsV1beta1
}

// FlowcontrolV1alpha1 retrieves the FlowcontrolV1alpha1ClusterClient.
func (c *ClusterClientset) FlowcontrolV1alpha1() flowcontrolv1alpha1.FlowcontrolV1alpha1ClusterInterface {
	return c.flowcontrolV1alpha1
}

// FlowcontrolV1beta1 retrieves the FlowcontrolV1beta1ClusterClient.
func (c *ClusterClientset) FlowcontrolV1beta1() flowcontrolv1beta1.FlowcontrolV1beta1ClusterInterface {
	return c.flowcontrolV1beta1
}

// FlowcontrolV1beta2 retrieves the FlowcontrolV1beta2ClusterClient.
func (c *ClusterClientset) FlowcontrolV1beta2() flowcontrolv1beta2.FlowcontrolV1beta2ClusterInterface {
	return c.flowcontrolV1beta2
}

// InternalV1alpha1 retrieves the InternalV1alpha1ClusterClient.
func (c *ClusterClientset) InternalV1alpha1() internalv1alpha1.InternalV1alpha1ClusterInterface {
	return c.internalV1alpha1
}

// NetworkingV1 retrieves the NetworkingV1ClusterClient.
func (c *ClusterClientset) NetworkingV1() networkingv1.NetworkingV1ClusterInterface {
	return c.networkingV1
}

// NetworkingV1beta1 retrieves the NetworkingV1beta1ClusterClient.
func (c *ClusterClientset) NetworkingV1beta1() networkingv1beta1.NetworkingV1beta1ClusterInterface {
	return c.networkingV1beta1
}

// NodeV1 retrieves the NodeV1ClusterClient.
func (c *ClusterClientset) NodeV1() nodev1.NodeV1ClusterInterface {
	return c.nodeV1
}

// NodeV1alpha1 retrieves the NodeV1alpha1ClusterClient.
func (c *ClusterClientset) NodeV1alpha1() nodev1alpha1.NodeV1alpha1ClusterInterface {
	return c.nodeV1alpha1
}

// NodeV1beta1 retrieves the NodeV1beta1ClusterClient.
func (c *ClusterClientset) NodeV1beta1() nodev1beta1.NodeV1beta1ClusterInterface {
	return c.nodeV1beta1
}

// PolicyV1 retrieves the PolicyV1ClusterClient.
func (c *ClusterClientset) PolicyV1() policyv1.PolicyV1ClusterInterface {
	return c.policyV1
}

// PolicyV1beta1 retrieves the PolicyV1beta1ClusterClient.
func (c *ClusterClientset) PolicyV1beta1() policyv1beta1.PolicyV1beta1ClusterInterface {
	return c.policyV1beta1
}

// RbacV1 retrieves the RbacV1ClusterClient.
func (c *ClusterClientset) RbacV1() rbacv1.RbacV1ClusterInterface {
	return c.rbacV1
}

// RbacV1alpha1 retrieves the RbacV1alpha1ClusterClient.
func (c *ClusterClientset) RbacV1alpha1() rbacv1alpha1.RbacV1alpha1ClusterInterface {
	return c.rbacV1alpha1
}

// RbacV1beta1 retrieves the RbacV1beta1ClusterClient.
func (c *ClusterClientset) RbacV1beta1() rbacv1beta1.RbacV1beta1ClusterInterface {
	return c.rbacV1beta1
}

// SchedulingV1 retrieves the SchedulingV1ClusterClient.
func (c *ClusterClientset) SchedulingV1() schedulingv1.SchedulingV1ClusterInterface {
	return c.schedulingV1
}

// SchedulingV1alpha1 retrieves the SchedulingV1alpha1ClusterClient.
func (c *ClusterClientset) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1ClusterInterface {
	return c.schedulingV1alpha1
}

// SchedulingV1beta1 retrieves the SchedulingV1beta1ClusterClient.
func (c *ClusterClientset) SchedulingV1beta1() schedulingv1beta1.SchedulingV1beta1ClusterInterface {
	return c.schedulingV1beta1
}

// StorageV1 retrieves the StorageV1ClusterClient.
func (c *ClusterClientset) StorageV1() storagev1.StorageV1ClusterInterface {
	return c.storageV1
}

// StorageV1alpha1 retrieves the StorageV1alpha1ClusterClient.
func (c *ClusterClientset) StorageV1alpha1() storagev1alpha1.StorageV1alpha1ClusterInterface {
	return c.storageV1alpha1
}

// StorageV1beta1 retrieves the StorageV1beta1ClusterClient.
func (c *ClusterClientset) StorageV1beta1() storagev1beta1.StorageV1beta1ClusterInterface {
	return c.storageV1beta1
}

// Cluster scopes this clientset to one cluster.
func (c *ClusterClientset) Cluster(name logicalcluster.Name) client.Interface {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return c.clientCache.ClusterOrDie(name)
}

// NewForConfig creates a new ClusterClientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ClusterClientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new ClusterClientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*ClusterClientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	cache := kcpclient.NewCache(c, httpClient, &kcpclient.Constructor[*client.Clientset]{
		NewForConfigAndClient: client.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.New("root")); err != nil {
		return nil, err
	}

	var cs ClusterClientset
	cs.clientCache = cache
	var err error
	cs.admissionregistrationV1, err = admissionregistrationv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.admissionregistrationV1beta1, err = admissionregistrationv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.appsV1, err = appsv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.appsV1beta1, err = appsv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.appsV1beta2, err = appsv1beta2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.authenticationV1, err = authenticationv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.authenticationV1beta1, err = authenticationv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.authorizationV1, err = authorizationv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.authorizationV1beta1, err = authorizationv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV1, err = autoscalingv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV2, err = autoscalingv2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV2beta1, err = autoscalingv2beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV2beta2, err = autoscalingv2beta2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.batchV1, err = batchv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.batchV1beta1, err = batchv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.certificatesV1, err = certificatesv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.certificatesV1beta1, err = certificatesv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.coordinationV1, err = coordinationv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.coordinationV1beta1, err = coordinationv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.coreV1, err = corev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.discoveryV1, err = discoveryv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.discoveryV1beta1, err = discoveryv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.eventsV1, err = eventsv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.eventsV1beta1, err = eventsv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.extensionsV1beta1, err = extensionsv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.flowcontrolV1alpha1, err = flowcontrolv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.flowcontrolV1beta1, err = flowcontrolv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.flowcontrolV1beta2, err = flowcontrolv1beta2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.internalV1alpha1, err = internalv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.networkingV1, err = networkingv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.networkingV1beta1, err = networkingv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.nodeV1, err = nodev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.nodeV1alpha1, err = nodev1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.nodeV1beta1, err = nodev1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.policyV1, err = policyv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.policyV1beta1, err = policyv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.rbacV1, err = rbacv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.rbacV1alpha1, err = rbacv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.rbacV1beta1, err = rbacv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.schedulingV1, err = schedulingv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.schedulingV1alpha1, err = schedulingv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.schedulingV1beta1, err = schedulingv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.storageV1, err = storagev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.storageV1alpha1, err = storagev1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.storageV1beta1, err = storagev1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new ClusterClientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ClusterClientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}
