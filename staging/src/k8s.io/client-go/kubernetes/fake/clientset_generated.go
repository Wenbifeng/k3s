/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	clientset "k8s.io/client-go/kubernetes"
	admissionregistrationv1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	fakeadmissionregistrationv1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1/fake"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	fakeappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1/fake"
	appsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	fakeappsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1/fake"
	appsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	fakeappsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2/fake"
	authenticationv1 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	fakeauthenticationv1 "k8s.io/client-go/kubernetes/typed/authentication/v1/fake"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	fakeauthorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1/fake"
	autoscalingv1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	fakeautoscalingv1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1/fake"
	autoscalingv2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	fakeautoscalingv2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1/fake"
	autoscalingv2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	fakeautoscalingv2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2/fake"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	fakebatchv1 "k8s.io/client-go/kubernetes/typed/batch/v1/fake"
	batchv1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	fakebatchv1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1/fake"
	batchv2alpha1 "k8s.io/client-go/kubernetes/typed/batch/v2alpha1"
	fakebatchv2alpha1 "k8s.io/client-go/kubernetes/typed/batch/v2alpha1/fake"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	fakecorev1 "k8s.io/client-go/kubernetes/typed/core/v1/fake"
	extensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	fakeextensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1/fake"
	networkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	fakenetworkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1/fake"
	policyv1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	fakepolicyv1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1/fake"
	rbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	fakerbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1/fake"
	schedulingv1beta1 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1"
	fakeschedulingv1beta1 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1/fake"
	storagev1 "k8s.io/client-go/kubernetes/typed/storage/v1"
	fakestoragev1 "k8s.io/client-go/kubernetes/typed/storage/v1/fake"
	storagev1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	fakestoragev1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// AdmissionregistrationV1beta1 retrieves the AdmissionregistrationV1beta1Client
func (c *Clientset) AdmissionregistrationV1beta1() admissionregistrationv1beta1.AdmissionregistrationV1beta1Interface {
	return &fakeadmissionregistrationv1beta1.FakeAdmissionregistrationV1beta1{Fake: &c.Fake}
}

// Admissionregistration retrieves the AdmissionregistrationV1beta1Client
func (c *Clientset) Admissionregistration() admissionregistrationv1beta1.AdmissionregistrationV1beta1Interface {
	return &fakeadmissionregistrationv1beta1.FakeAdmissionregistrationV1beta1{Fake: &c.Fake}
}

// AppsV1beta1 retrieves the AppsV1beta1Client
func (c *Clientset) AppsV1beta1() appsv1beta1.AppsV1beta1Interface {
	return &fakeappsv1beta1.FakeAppsV1beta1{Fake: &c.Fake}
}

// AppsV1beta2 retrieves the AppsV1beta2Client
func (c *Clientset) AppsV1beta2() appsv1beta2.AppsV1beta2Interface {
	return &fakeappsv1beta2.FakeAppsV1beta2{Fake: &c.Fake}
}

// AppsV1 retrieves the AppsV1Client
func (c *Clientset) AppsV1() appsv1.AppsV1Interface {
	return &fakeappsv1.FakeAppsV1{Fake: &c.Fake}
}

// Apps retrieves the AppsV1Client
func (c *Clientset) Apps() appsv1.AppsV1Interface {
	return &fakeappsv1.FakeAppsV1{Fake: &c.Fake}
}

// AuthenticationV1 retrieves the AuthenticationV1Client
func (c *Clientset) AuthenticationV1() authenticationv1.AuthenticationV1Interface {
	return &fakeauthenticationv1.FakeAuthenticationV1{Fake: &c.Fake}
}

// Authentication retrieves the AuthenticationV1Client
func (c *Clientset) Authentication() authenticationv1.AuthenticationV1Interface {
	return &fakeauthenticationv1.FakeAuthenticationV1{Fake: &c.Fake}
}

// AuthorizationV1 retrieves the AuthorizationV1Client
func (c *Clientset) AuthorizationV1() authorizationv1.AuthorizationV1Interface {
	return &fakeauthorizationv1.FakeAuthorizationV1{Fake: &c.Fake}
}

// Authorization retrieves the AuthorizationV1Client
func (c *Clientset) Authorization() authorizationv1.AuthorizationV1Interface {
	return &fakeauthorizationv1.FakeAuthorizationV1{Fake: &c.Fake}
}

// AutoscalingV1 retrieves the AutoscalingV1Client
func (c *Clientset) AutoscalingV1() autoscalingv1.AutoscalingV1Interface {
	return &fakeautoscalingv1.FakeAutoscalingV1{Fake: &c.Fake}
}

// Autoscaling retrieves the AutoscalingV1Client
func (c *Clientset) Autoscaling() autoscalingv1.AutoscalingV1Interface {
	return &fakeautoscalingv1.FakeAutoscalingV1{Fake: &c.Fake}
}

// AutoscalingV2beta1 retrieves the AutoscalingV2beta1Client
func (c *Clientset) AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1Interface {
	return &fakeautoscalingv2beta1.FakeAutoscalingV2beta1{Fake: &c.Fake}
}

// AutoscalingV2beta2 retrieves the AutoscalingV2beta2Client
func (c *Clientset) AutoscalingV2beta2() autoscalingv2beta2.AutoscalingV2beta2Interface {
	return &fakeautoscalingv2beta2.FakeAutoscalingV2beta2{Fake: &c.Fake}
}

// BatchV1 retrieves the BatchV1Client
func (c *Clientset) BatchV1() batchv1.BatchV1Interface {
	return &fakebatchv1.FakeBatchV1{Fake: &c.Fake}
}

// Batch retrieves the BatchV1Client
func (c *Clientset) Batch() batchv1.BatchV1Interface {
	return &fakebatchv1.FakeBatchV1{Fake: &c.Fake}
}

// BatchV1beta1 retrieves the BatchV1beta1Client
func (c *Clientset) BatchV1beta1() batchv1beta1.BatchV1beta1Interface {
	return &fakebatchv1beta1.FakeBatchV1beta1{Fake: &c.Fake}
}

// BatchV2alpha1 retrieves the BatchV2alpha1Client
func (c *Clientset) BatchV2alpha1() batchv2alpha1.BatchV2alpha1Interface {
	return &fakebatchv2alpha1.FakeBatchV2alpha1{Fake: &c.Fake}
}

// CoreV1 retrieves the CoreV1Client
func (c *Clientset) CoreV1() corev1.CoreV1Interface {
	return &fakecorev1.FakeCoreV1{Fake: &c.Fake}
}

// Core retrieves the CoreV1Client
func (c *Clientset) Core() corev1.CoreV1Interface {
	return &fakecorev1.FakeCoreV1{Fake: &c.Fake}
}

// ExtensionsV1beta1 retrieves the ExtensionsV1beta1Client
func (c *Clientset) ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1Interface {
	return &fakeextensionsv1beta1.FakeExtensionsV1beta1{Fake: &c.Fake}
}

// Extensions retrieves the ExtensionsV1beta1Client
func (c *Clientset) Extensions() extensionsv1beta1.ExtensionsV1beta1Interface {
	return &fakeextensionsv1beta1.FakeExtensionsV1beta1{Fake: &c.Fake}
}

// NetworkingV1 retrieves the NetworkingV1Client
func (c *Clientset) NetworkingV1() networkingv1.NetworkingV1Interface {
	return &fakenetworkingv1.FakeNetworkingV1{Fake: &c.Fake}
}

// Networking retrieves the NetworkingV1Client
func (c *Clientset) Networking() networkingv1.NetworkingV1Interface {
	return &fakenetworkingv1.FakeNetworkingV1{Fake: &c.Fake}
}

// PolicyV1beta1 retrieves the PolicyV1beta1Client
func (c *Clientset) PolicyV1beta1() policyv1beta1.PolicyV1beta1Interface {
	return &fakepolicyv1beta1.FakePolicyV1beta1{Fake: &c.Fake}
}

// Policy retrieves the PolicyV1beta1Client
func (c *Clientset) Policy() policyv1beta1.PolicyV1beta1Interface {
	return &fakepolicyv1beta1.FakePolicyV1beta1{Fake: &c.Fake}
}

// RbacV1 retrieves the RbacV1Client
func (c *Clientset) RbacV1() rbacv1.RbacV1Interface {
	return &fakerbacv1.FakeRbacV1{Fake: &c.Fake}
}

// Rbac retrieves the RbacV1Client
func (c *Clientset) Rbac() rbacv1.RbacV1Interface {
	return &fakerbacv1.FakeRbacV1{Fake: &c.Fake}
}

// SchedulingV1beta1 retrieves the SchedulingV1beta1Client
func (c *Clientset) SchedulingV1beta1() schedulingv1beta1.SchedulingV1beta1Interface {
	return &fakeschedulingv1beta1.FakeSchedulingV1beta1{Fake: &c.Fake}
}

// Scheduling retrieves the SchedulingV1beta1Client
func (c *Clientset) Scheduling() schedulingv1beta1.SchedulingV1beta1Interface {
	return &fakeschedulingv1beta1.FakeSchedulingV1beta1{Fake: &c.Fake}
}

// StorageV1beta1 retrieves the StorageV1beta1Client
func (c *Clientset) StorageV1beta1() storagev1beta1.StorageV1beta1Interface {
	return &fakestoragev1beta1.FakeStorageV1beta1{Fake: &c.Fake}
}

// StorageV1 retrieves the StorageV1Client
func (c *Clientset) StorageV1() storagev1.StorageV1Interface {
	return &fakestoragev1.FakeStorageV1{Fake: &c.Fake}
}

// Storage retrieves the StorageV1Client
func (c *Clientset) Storage() storagev1.StorageV1Interface {
	return &fakestoragev1.FakeStorageV1{Fake: &c.Fake}
}
