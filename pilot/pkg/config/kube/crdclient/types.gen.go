// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by pilot/pkg/config/kube/crd/codegen/types.go. DO NOT EDIT!

package crdclient

// This file contains Go definitions for Custom Resource Definition kinds
// to adhere to the idiomatic use of k8s API machinery.
// These definitions are synthesized from Istio configuration type descriptors
// as declared in the Istio config model.

import (
	"context"
	"fmt"

	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	serviceapisclient "sigs.k8s.io/service-apis/pkg/client/clientset/versioned"

	"istio.io/istio/pkg/config"
	"istio.io/istio/pkg/config/schema/collections"

	networkingv1alpha3 "istio.io/api/networking/v1alpha3"
	securityv1beta1 "istio.io/api/security/v1beta1"
	clientnetworkingv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	clientsecurityv1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"

	servicev1alpha1 "sigs.k8s.io/service-apis/apis/v1alpha1"
)

func create(ic versionedclient.Interface, sc serviceapisclient.Interface, cfg config.Config, objMeta metav1.ObjectMeta) (metav1.Object, error) {
	switch cfg.GroupVersionKind {
	case collections.IstioNetworkingV1Alpha3Destinationrules.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().DestinationRules(cfg.Namespace).Create(context.TODO(), &clientnetworkingv1alpha3.DestinationRule{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.DestinationRule)),
		}, metav1.CreateOptions{})
	case collections.IstioNetworkingV1Alpha3Envoyfilters.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().EnvoyFilters(cfg.Namespace).Create(context.TODO(), &clientnetworkingv1alpha3.EnvoyFilter{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.EnvoyFilter)),
		}, metav1.CreateOptions{})
	case collections.IstioNetworkingV1Alpha3Gateways.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().Gateways(cfg.Namespace).Create(context.TODO(), &clientnetworkingv1alpha3.Gateway{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.Gateway)),
		}, metav1.CreateOptions{})
	case collections.IstioNetworkingV1Alpha3Serviceentries.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().ServiceEntries(cfg.Namespace).Create(context.TODO(), &clientnetworkingv1alpha3.ServiceEntry{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.ServiceEntry)),
		}, metav1.CreateOptions{})
	case collections.IstioNetworkingV1Alpha3Sidecars.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().Sidecars(cfg.Namespace).Create(context.TODO(), &clientnetworkingv1alpha3.Sidecar{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.Sidecar)),
		}, metav1.CreateOptions{})
	case collections.IstioNetworkingV1Alpha3Virtualservices.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().VirtualServices(cfg.Namespace).Create(context.TODO(), &clientnetworkingv1alpha3.VirtualService{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.VirtualService)),
		}, metav1.CreateOptions{})
	case collections.IstioNetworkingV1Alpha3Workloadentries.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().WorkloadEntries(cfg.Namespace).Create(context.TODO(), &clientnetworkingv1alpha3.WorkloadEntry{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.WorkloadEntry)),
		}, metav1.CreateOptions{})
	case collections.IstioNetworkingV1Alpha3Workloadgroups.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().WorkloadGroups(cfg.Namespace).Create(context.TODO(), &clientnetworkingv1alpha3.WorkloadGroup{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.WorkloadGroup)),
		}, metav1.CreateOptions{})
	case collections.IstioSecurityV1Beta1Authorizationpolicies.Resource().GroupVersionKind():
		return ic.SecurityV1beta1().AuthorizationPolicies(cfg.Namespace).Create(context.TODO(), &clientsecurityv1beta1.AuthorizationPolicy{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*securityv1beta1.AuthorizationPolicy)),
		}, metav1.CreateOptions{})
	case collections.IstioSecurityV1Beta1Peerauthentications.Resource().GroupVersionKind():
		return ic.SecurityV1beta1().PeerAuthentications(cfg.Namespace).Create(context.TODO(), &clientsecurityv1beta1.PeerAuthentication{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*securityv1beta1.PeerAuthentication)),
		}, metav1.CreateOptions{})
	case collections.IstioSecurityV1Beta1Requestauthentications.Resource().GroupVersionKind():
		return ic.SecurityV1beta1().RequestAuthentications(cfg.Namespace).Create(context.TODO(), &clientsecurityv1beta1.RequestAuthentication{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*securityv1beta1.RequestAuthentication)),
		}, metav1.CreateOptions{})
	case collections.K8SServiceApisV1Alpha1Gatewayclasses.Resource().GroupVersionKind():
		return sc.NetworkingV1alpha1().GatewayClasses().Create(context.TODO(), &servicev1alpha1.GatewayClass{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*servicev1alpha1.GatewayClassSpec)),
		}, metav1.CreateOptions{})
	case collections.K8SServiceApisV1Alpha1Gateways.Resource().GroupVersionKind():
		return sc.NetworkingV1alpha1().Gateways(cfg.Namespace).Create(context.TODO(), &servicev1alpha1.Gateway{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*servicev1alpha1.GatewaySpec)),
		}, metav1.CreateOptions{})
	case collections.K8SServiceApisV1Alpha1Httproutes.Resource().GroupVersionKind():
		return sc.NetworkingV1alpha1().HTTPRoutes(cfg.Namespace).Create(context.TODO(), &servicev1alpha1.HTTPRoute{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*servicev1alpha1.HTTPRouteSpec)),
		}, metav1.CreateOptions{})
	case collections.K8SServiceApisV1Alpha1Tcproutes.Resource().GroupVersionKind():
		return sc.NetworkingV1alpha1().TCPRoutes(cfg.Namespace).Create(context.TODO(), &servicev1alpha1.TCPRoute{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*servicev1alpha1.TCPRouteSpec)),
		}, metav1.CreateOptions{})
	default:
		return nil, fmt.Errorf("unsupported type: %v", cfg.GroupVersionKind)
	}
}

func update(ic versionedclient.Interface, sc serviceapisclient.Interface, cfg config.Config, objMeta metav1.ObjectMeta) (metav1.Object, error) {
	switch cfg.GroupVersionKind {
	case collections.IstioNetworkingV1Alpha3Destinationrules.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().DestinationRules(cfg.Namespace).Update(context.TODO(), &clientnetworkingv1alpha3.DestinationRule{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.DestinationRule)),
		}, metav1.UpdateOptions{})
	case collections.IstioNetworkingV1Alpha3Envoyfilters.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().EnvoyFilters(cfg.Namespace).Update(context.TODO(), &clientnetworkingv1alpha3.EnvoyFilter{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.EnvoyFilter)),
		}, metav1.UpdateOptions{})
	case collections.IstioNetworkingV1Alpha3Gateways.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().Gateways(cfg.Namespace).Update(context.TODO(), &clientnetworkingv1alpha3.Gateway{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.Gateway)),
		}, metav1.UpdateOptions{})
	case collections.IstioNetworkingV1Alpha3Serviceentries.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().ServiceEntries(cfg.Namespace).Update(context.TODO(), &clientnetworkingv1alpha3.ServiceEntry{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.ServiceEntry)),
		}, metav1.UpdateOptions{})
	case collections.IstioNetworkingV1Alpha3Sidecars.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().Sidecars(cfg.Namespace).Update(context.TODO(), &clientnetworkingv1alpha3.Sidecar{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.Sidecar)),
		}, metav1.UpdateOptions{})
	case collections.IstioNetworkingV1Alpha3Virtualservices.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().VirtualServices(cfg.Namespace).Update(context.TODO(), &clientnetworkingv1alpha3.VirtualService{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.VirtualService)),
		}, metav1.UpdateOptions{})
	case collections.IstioNetworkingV1Alpha3Workloadentries.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().WorkloadEntries(cfg.Namespace).Update(context.TODO(), &clientnetworkingv1alpha3.WorkloadEntry{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.WorkloadEntry)),
		}, metav1.UpdateOptions{})
	case collections.IstioNetworkingV1Alpha3Workloadgroups.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().WorkloadGroups(cfg.Namespace).Update(context.TODO(), &clientnetworkingv1alpha3.WorkloadGroup{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*networkingv1alpha3.WorkloadGroup)),
		}, metav1.UpdateOptions{})
	case collections.IstioSecurityV1Beta1Authorizationpolicies.Resource().GroupVersionKind():
		return ic.SecurityV1beta1().AuthorizationPolicies(cfg.Namespace).Update(context.TODO(), &clientsecurityv1beta1.AuthorizationPolicy{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*securityv1beta1.AuthorizationPolicy)),
		}, metav1.UpdateOptions{})
	case collections.IstioSecurityV1Beta1Peerauthentications.Resource().GroupVersionKind():
		return ic.SecurityV1beta1().PeerAuthentications(cfg.Namespace).Update(context.TODO(), &clientsecurityv1beta1.PeerAuthentication{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*securityv1beta1.PeerAuthentication)),
		}, metav1.UpdateOptions{})
	case collections.IstioSecurityV1Beta1Requestauthentications.Resource().GroupVersionKind():
		return ic.SecurityV1beta1().RequestAuthentications(cfg.Namespace).Update(context.TODO(), &clientsecurityv1beta1.RequestAuthentication{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*securityv1beta1.RequestAuthentication)),
		}, metav1.UpdateOptions{})
	case collections.K8SServiceApisV1Alpha1Gatewayclasses.Resource().GroupVersionKind():
		return sc.NetworkingV1alpha1().GatewayClasses().Update(context.TODO(), &servicev1alpha1.GatewayClass{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*servicev1alpha1.GatewayClassSpec)),
		}, metav1.UpdateOptions{})
	case collections.K8SServiceApisV1Alpha1Gateways.Resource().GroupVersionKind():
		return sc.NetworkingV1alpha1().Gateways(cfg.Namespace).Update(context.TODO(), &servicev1alpha1.Gateway{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*servicev1alpha1.GatewaySpec)),
		}, metav1.UpdateOptions{})
	case collections.K8SServiceApisV1Alpha1Httproutes.Resource().GroupVersionKind():
		return sc.NetworkingV1alpha1().HTTPRoutes(cfg.Namespace).Update(context.TODO(), &servicev1alpha1.HTTPRoute{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*servicev1alpha1.HTTPRouteSpec)),
		}, metav1.UpdateOptions{})
	case collections.K8SServiceApisV1Alpha1Tcproutes.Resource().GroupVersionKind():
		return sc.NetworkingV1alpha1().TCPRoutes(cfg.Namespace).Update(context.TODO(), &servicev1alpha1.TCPRoute{
			ObjectMeta: objMeta,
			Spec:       *(cfg.Spec.(*servicev1alpha1.TCPRouteSpec)),
		}, metav1.UpdateOptions{})
	default:
		return nil, fmt.Errorf("unsupported type: %v", cfg.GroupVersionKind)
	}
}

func patch(ic versionedclient.Interface, sc serviceapisclient.Interface, orig config.Config, origMeta metav1.ObjectMeta, mod config.Config, modMeta metav1.ObjectMeta) (metav1.Object, error) {
	if orig.GroupVersionKind != mod.GroupVersionKind {
		return nil, fmt.Errorf("gvk mismatch: %v, modified: %v", orig.GroupVersionKind, mod.GroupVersionKind)
	}
	// TODO support multiple patch types and setting field manager
	switch orig.GroupVersionKind {
	case collections.IstioNetworkingV1Alpha3Destinationrules.Resource().GroupVersionKind():
		oldRes := &clientnetworkingv1alpha3.DestinationRule{
			ObjectMeta: origMeta,
			Spec:       *(orig.Spec.(*networkingv1alpha3.DestinationRule)),
		}
		modRes := &clientnetworkingv1alpha3.DestinationRule{
			ObjectMeta: modMeta,
			Spec:       *(mod.Spec.(*networkingv1alpha3.DestinationRule)),
		}
		patchBytes, err := genPatchBytes(oldRes, modRes)
		if err != nil {
			return nil, err
		}
		return ic.NetworkingV1alpha3().DestinationRules(orig.Namespace).
			Patch(context.TODO(), orig.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{FieldManager: "pilot-discovery"})
	case collections.IstioNetworkingV1Alpha3Envoyfilters.Resource().GroupVersionKind():
		oldRes := &clientnetworkingv1alpha3.EnvoyFilter{
			ObjectMeta: origMeta,
			Spec:       *(orig.Spec.(*networkingv1alpha3.EnvoyFilter)),
		}
		modRes := &clientnetworkingv1alpha3.EnvoyFilter{
			ObjectMeta: modMeta,
			Spec:       *(mod.Spec.(*networkingv1alpha3.EnvoyFilter)),
		}
		patchBytes, err := genPatchBytes(oldRes, modRes)
		if err != nil {
			return nil, err
		}
		return ic.NetworkingV1alpha3().EnvoyFilters(orig.Namespace).
			Patch(context.TODO(), orig.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{FieldManager: "pilot-discovery"})
	case collections.IstioNetworkingV1Alpha3Gateways.Resource().GroupVersionKind():
		oldRes := &clientnetworkingv1alpha3.Gateway{
			ObjectMeta: origMeta,
			Spec:       *(orig.Spec.(*networkingv1alpha3.Gateway)),
		}
		modRes := &clientnetworkingv1alpha3.Gateway{
			ObjectMeta: modMeta,
			Spec:       *(mod.Spec.(*networkingv1alpha3.Gateway)),
		}
		patchBytes, err := genPatchBytes(oldRes, modRes)
		if err != nil {
			return nil, err
		}
		return ic.NetworkingV1alpha3().Gateways(orig.Namespace).
			Patch(context.TODO(), orig.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{FieldManager: "pilot-discovery"})
	case collections.IstioNetworkingV1Alpha3Serviceentries.Resource().GroupVersionKind():
		oldRes := &clientnetworkingv1alpha3.ServiceEntry{
			ObjectMeta: origMeta,
			Spec:       *(orig.Spec.(*networkingv1alpha3.ServiceEntry)),
		}
		modRes := &clientnetworkingv1alpha3.ServiceEntry{
			ObjectMeta: modMeta,
			Spec:       *(mod.Spec.(*networkingv1alpha3.ServiceEntry)),
		}
		patchBytes, err := genPatchBytes(oldRes, modRes)
		if err != nil {
			return nil, err
		}
		return ic.NetworkingV1alpha3().ServiceEntries(orig.Namespace).
			Patch(context.TODO(), orig.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{FieldManager: "pilot-discovery"})
	case collections.IstioNetworkingV1Alpha3Sidecars.Resource().GroupVersionKind():
		oldRes := &clientnetworkingv1alpha3.Sidecar{
			ObjectMeta: origMeta,
			Spec:       *(orig.Spec.(*networkingv1alpha3.Sidecar)),
		}
		modRes := &clientnetworkingv1alpha3.Sidecar{
			ObjectMeta: modMeta,
			Spec:       *(mod.Spec.(*networkingv1alpha3.Sidecar)),
		}
		patchBytes, err := genPatchBytes(oldRes, modRes)
		if err != nil {
			return nil, err
		}
		return ic.NetworkingV1alpha3().Sidecars(orig.Namespace).
			Patch(context.TODO(), orig.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{FieldManager: "pilot-discovery"})
	case collections.IstioNetworkingV1Alpha3Virtualservices.Resource().GroupVersionKind():
		oldRes := &clientnetworkingv1alpha3.VirtualService{
			ObjectMeta: origMeta,
			Spec:       *(orig.Spec.(*networkingv1alpha3.VirtualService)),
		}
		modRes := &clientnetworkingv1alpha3.VirtualService{
			ObjectMeta: modMeta,
			Spec:       *(mod.Spec.(*networkingv1alpha3.VirtualService)),
		}
		patchBytes, err := genPatchBytes(oldRes, modRes)
		if err != nil {
			return nil, err
		}
		return ic.NetworkingV1alpha3().VirtualServices(orig.Namespace).
			Patch(context.TODO(), orig.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{FieldManager: "pilot-discovery"})
	case collections.IstioNetworkingV1Alpha3Workloadentries.Resource().GroupVersionKind():
		oldRes := &clientnetworkingv1alpha3.WorkloadEntry{
			ObjectMeta: origMeta,
			Spec:       *(orig.Spec.(*networkingv1alpha3.WorkloadEntry)),
		}
		modRes := &clientnetworkingv1alpha3.WorkloadEntry{
			ObjectMeta: modMeta,
			Spec:       *(mod.Spec.(*networkingv1alpha3.WorkloadEntry)),
		}
		patchBytes, err := genPatchBytes(oldRes, modRes)
		if err != nil {
			return nil, err
		}
		return ic.NetworkingV1alpha3().WorkloadEntries(orig.Namespace).
			Patch(context.TODO(), orig.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{FieldManager: "pilot-discovery"})
	case collections.IstioNetworkingV1Alpha3Workloadgroups.Resource().GroupVersionKind():
		oldRes := &clientnetworkingv1alpha3.WorkloadGroup{
			ObjectMeta: origMeta,
			Spec:       *(orig.Spec.(*networkingv1alpha3.WorkloadGroup)),
		}
		modRes := &clientnetworkingv1alpha3.WorkloadGroup{
			ObjectMeta: modMeta,
			Spec:       *(mod.Spec.(*networkingv1alpha3.WorkloadGroup)),
		}
		patchBytes, err := genPatchBytes(oldRes, modRes)
		if err != nil {
			return nil, err
		}
		return ic.NetworkingV1alpha3().WorkloadGroups(orig.Namespace).
			Patch(context.TODO(), orig.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{FieldManager: "pilot-discovery"})
	case collections.IstioSecurityV1Beta1Authorizationpolicies.Resource().GroupVersionKind():
		oldRes := &clientsecurityv1beta1.AuthorizationPolicy{
			ObjectMeta: origMeta,
			Spec:       *(orig.Spec.(*securityv1beta1.AuthorizationPolicy)),
		}
		modRes := &clientsecurityv1beta1.AuthorizationPolicy{
			ObjectMeta: modMeta,
			Spec:       *(mod.Spec.(*securityv1beta1.AuthorizationPolicy)),
		}
		patchBytes, err := genPatchBytes(oldRes, modRes)
		if err != nil {
			return nil, err
		}
		return ic.SecurityV1beta1().AuthorizationPolicies(orig.Namespace).
			Patch(context.TODO(), orig.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{FieldManager: "pilot-discovery"})
	case collections.IstioSecurityV1Beta1Peerauthentications.Resource().GroupVersionKind():
		oldRes := &clientsecurityv1beta1.PeerAuthentication{
			ObjectMeta: origMeta,
			Spec:       *(orig.Spec.(*securityv1beta1.PeerAuthentication)),
		}
		modRes := &clientsecurityv1beta1.PeerAuthentication{
			ObjectMeta: modMeta,
			Spec:       *(mod.Spec.(*securityv1beta1.PeerAuthentication)),
		}
		patchBytes, err := genPatchBytes(oldRes, modRes)
		if err != nil {
			return nil, err
		}
		return ic.SecurityV1beta1().PeerAuthentications(orig.Namespace).
			Patch(context.TODO(), orig.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{FieldManager: "pilot-discovery"})
	case collections.IstioSecurityV1Beta1Requestauthentications.Resource().GroupVersionKind():
		oldRes := &clientsecurityv1beta1.RequestAuthentication{
			ObjectMeta: origMeta,
			Spec:       *(orig.Spec.(*securityv1beta1.RequestAuthentication)),
		}
		modRes := &clientsecurityv1beta1.RequestAuthentication{
			ObjectMeta: modMeta,
			Spec:       *(mod.Spec.(*securityv1beta1.RequestAuthentication)),
		}
		patchBytes, err := genPatchBytes(oldRes, modRes)
		if err != nil {
			return nil, err
		}
		return ic.SecurityV1beta1().RequestAuthentications(orig.Namespace).
			Patch(context.TODO(), orig.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{FieldManager: "pilot-discovery"})
	case collections.K8SServiceApisV1Alpha1Gatewayclasses.Resource().GroupVersionKind():
		oldRes := &servicev1alpha1.GatewayClass{
			ObjectMeta: origMeta,
			Spec:       *(orig.Spec.(*servicev1alpha1.GatewayClassSpec)),
		}
		modRes := &servicev1alpha1.GatewayClass{
			ObjectMeta: modMeta,
			Spec:       *(mod.Spec.(*servicev1alpha1.GatewayClassSpec)),
		}
		patchBytes, err := genPatchBytes(oldRes, modRes)
		if err != nil {
			return nil, err
		}
		return sc.NetworkingV1alpha1().GatewayClasses().
			Patch(context.TODO(), orig.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{FieldManager: "pilot-discovery"})
	case collections.K8SServiceApisV1Alpha1Gateways.Resource().GroupVersionKind():
		oldRes := &servicev1alpha1.Gateway{
			ObjectMeta: origMeta,
			Spec:       *(orig.Spec.(*servicev1alpha1.GatewaySpec)),
		}
		modRes := &servicev1alpha1.Gateway{
			ObjectMeta: modMeta,
			Spec:       *(mod.Spec.(*servicev1alpha1.GatewaySpec)),
		}
		patchBytes, err := genPatchBytes(oldRes, modRes)
		if err != nil {
			return nil, err
		}
		return sc.NetworkingV1alpha1().Gateways(orig.Namespace).
			Patch(context.TODO(), orig.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{FieldManager: "pilot-discovery"})
	case collections.K8SServiceApisV1Alpha1Httproutes.Resource().GroupVersionKind():
		oldRes := &servicev1alpha1.HTTPRoute{
			ObjectMeta: origMeta,
			Spec:       *(orig.Spec.(*servicev1alpha1.HTTPRouteSpec)),
		}
		modRes := &servicev1alpha1.HTTPRoute{
			ObjectMeta: modMeta,
			Spec:       *(mod.Spec.(*servicev1alpha1.HTTPRouteSpec)),
		}
		patchBytes, err := genPatchBytes(oldRes, modRes)
		if err != nil {
			return nil, err
		}
		return sc.NetworkingV1alpha1().HTTPRoutes(orig.Namespace).
			Patch(context.TODO(), orig.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{FieldManager: "pilot-discovery"})
	case collections.K8SServiceApisV1Alpha1Tcproutes.Resource().GroupVersionKind():
		oldRes := &servicev1alpha1.TCPRoute{
			ObjectMeta: origMeta,
			Spec:       *(orig.Spec.(*servicev1alpha1.TCPRouteSpec)),
		}
		modRes := &servicev1alpha1.TCPRoute{
			ObjectMeta: modMeta,
			Spec:       *(mod.Spec.(*servicev1alpha1.TCPRouteSpec)),
		}
		patchBytes, err := genPatchBytes(oldRes, modRes)
		if err != nil {
			return nil, err
		}
		return sc.NetworkingV1alpha1().TCPRoutes(orig.Namespace).
			Patch(context.TODO(), orig.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{FieldManager: "pilot-discovery"})
	default:
		return nil, fmt.Errorf("unsupported type: %v", orig.GroupVersionKind)
	}
}

func delete(ic versionedclient.Interface, sc serviceapisclient.Interface, typ config.GroupVersionKind, name, namespace string) error {
	switch typ {
	case collections.IstioNetworkingV1Alpha3Destinationrules.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().DestinationRules(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	case collections.IstioNetworkingV1Alpha3Envoyfilters.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().EnvoyFilters(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	case collections.IstioNetworkingV1Alpha3Gateways.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().Gateways(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	case collections.IstioNetworkingV1Alpha3Serviceentries.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().ServiceEntries(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	case collections.IstioNetworkingV1Alpha3Sidecars.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().Sidecars(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	case collections.IstioNetworkingV1Alpha3Virtualservices.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().VirtualServices(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	case collections.IstioNetworkingV1Alpha3Workloadentries.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().WorkloadEntries(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	case collections.IstioNetworkingV1Alpha3Workloadgroups.Resource().GroupVersionKind():
		return ic.NetworkingV1alpha3().WorkloadGroups(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	case collections.IstioSecurityV1Beta1Authorizationpolicies.Resource().GroupVersionKind():
		return ic.SecurityV1beta1().AuthorizationPolicies(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	case collections.IstioSecurityV1Beta1Peerauthentications.Resource().GroupVersionKind():
		return ic.SecurityV1beta1().PeerAuthentications(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	case collections.IstioSecurityV1Beta1Requestauthentications.Resource().GroupVersionKind():
		return ic.SecurityV1beta1().RequestAuthentications(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	case collections.K8SServiceApisV1Alpha1Gatewayclasses.Resource().GroupVersionKind():
		return sc.NetworkingV1alpha1().GatewayClasses().Delete(context.TODO(), name, metav1.DeleteOptions{})
	case collections.K8SServiceApisV1Alpha1Gateways.Resource().GroupVersionKind():
		return sc.NetworkingV1alpha1().Gateways(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	case collections.K8SServiceApisV1Alpha1Httproutes.Resource().GroupVersionKind():
		return sc.NetworkingV1alpha1().HTTPRoutes(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	case collections.K8SServiceApisV1Alpha1Tcproutes.Resource().GroupVersionKind():
		return sc.NetworkingV1alpha1().TCPRoutes(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	default:
		return fmt.Errorf("unsupported type: %v", typ)
	}
}

var translationMap = map[config.GroupVersionKind]func(r runtime.Object) *config.Config{
	collections.IstioNetworkingV1Alpha3Destinationrules.Resource().GroupVersionKind(): func(r runtime.Object) *config.Config {
		obj := r.(*clientnetworkingv1alpha3.DestinationRule)
		return &config.Config{
			Meta: config.Meta{
				GroupVersionKind:  collections.IstioNetworkingV1Alpha3Destinationrules.Resource().GroupVersionKind(),
				Name:              obj.Name,
				Namespace:         obj.Namespace,
				Labels:            obj.Labels,
				Annotations:       obj.Annotations,
				ResourceVersion:   obj.ResourceVersion,
				CreationTimestamp: obj.CreationTimestamp.Time,
			},
			Spec: &obj.Spec,
		}
	},
	collections.IstioNetworkingV1Alpha3Envoyfilters.Resource().GroupVersionKind(): func(r runtime.Object) *config.Config {
		obj := r.(*clientnetworkingv1alpha3.EnvoyFilter)
		return &config.Config{
			Meta: config.Meta{
				GroupVersionKind:  collections.IstioNetworkingV1Alpha3Envoyfilters.Resource().GroupVersionKind(),
				Name:              obj.Name,
				Namespace:         obj.Namespace,
				Labels:            obj.Labels,
				Annotations:       obj.Annotations,
				ResourceVersion:   obj.ResourceVersion,
				CreationTimestamp: obj.CreationTimestamp.Time,
			},
			Spec: &obj.Spec,
		}
	},
	collections.IstioNetworkingV1Alpha3Gateways.Resource().GroupVersionKind(): func(r runtime.Object) *config.Config {
		obj := r.(*clientnetworkingv1alpha3.Gateway)
		return &config.Config{
			Meta: config.Meta{
				GroupVersionKind:  collections.IstioNetworkingV1Alpha3Gateways.Resource().GroupVersionKind(),
				Name:              obj.Name,
				Namespace:         obj.Namespace,
				Labels:            obj.Labels,
				Annotations:       obj.Annotations,
				ResourceVersion:   obj.ResourceVersion,
				CreationTimestamp: obj.CreationTimestamp.Time,
			},
			Spec: &obj.Spec,
		}
	},
	collections.IstioNetworkingV1Alpha3Serviceentries.Resource().GroupVersionKind(): func(r runtime.Object) *config.Config {
		obj := r.(*clientnetworkingv1alpha3.ServiceEntry)
		return &config.Config{
			Meta: config.Meta{
				GroupVersionKind:  collections.IstioNetworkingV1Alpha3Serviceentries.Resource().GroupVersionKind(),
				Name:              obj.Name,
				Namespace:         obj.Namespace,
				Labels:            obj.Labels,
				Annotations:       obj.Annotations,
				ResourceVersion:   obj.ResourceVersion,
				CreationTimestamp: obj.CreationTimestamp.Time,
			},
			Spec: &obj.Spec,
		}
	},
	collections.IstioNetworkingV1Alpha3Sidecars.Resource().GroupVersionKind(): func(r runtime.Object) *config.Config {
		obj := r.(*clientnetworkingv1alpha3.Sidecar)
		return &config.Config{
			Meta: config.Meta{
				GroupVersionKind:  collections.IstioNetworkingV1Alpha3Sidecars.Resource().GroupVersionKind(),
				Name:              obj.Name,
				Namespace:         obj.Namespace,
				Labels:            obj.Labels,
				Annotations:       obj.Annotations,
				ResourceVersion:   obj.ResourceVersion,
				CreationTimestamp: obj.CreationTimestamp.Time,
			},
			Spec: &obj.Spec,
		}
	},
	collections.IstioNetworkingV1Alpha3Virtualservices.Resource().GroupVersionKind(): func(r runtime.Object) *config.Config {
		obj := r.(*clientnetworkingv1alpha3.VirtualService)
		return &config.Config{
			Meta: config.Meta{
				GroupVersionKind:  collections.IstioNetworkingV1Alpha3Virtualservices.Resource().GroupVersionKind(),
				Name:              obj.Name,
				Namespace:         obj.Namespace,
				Labels:            obj.Labels,
				Annotations:       obj.Annotations,
				ResourceVersion:   obj.ResourceVersion,
				CreationTimestamp: obj.CreationTimestamp.Time,
			},
			Spec: &obj.Spec,
		}
	},
	collections.IstioNetworkingV1Alpha3Workloadentries.Resource().GroupVersionKind(): func(r runtime.Object) *config.Config {
		obj := r.(*clientnetworkingv1alpha3.WorkloadEntry)
		return &config.Config{
			Meta: config.Meta{
				GroupVersionKind:  collections.IstioNetworkingV1Alpha3Workloadentries.Resource().GroupVersionKind(),
				Name:              obj.Name,
				Namespace:         obj.Namespace,
				Labels:            obj.Labels,
				Annotations:       obj.Annotations,
				ResourceVersion:   obj.ResourceVersion,
				CreationTimestamp: obj.CreationTimestamp.Time,
			},
			Spec: &obj.Spec,
		}
	},
	collections.IstioNetworkingV1Alpha3Workloadgroups.Resource().GroupVersionKind(): func(r runtime.Object) *config.Config {
		obj := r.(*clientnetworkingv1alpha3.WorkloadGroup)
		return &config.Config{
			Meta: config.Meta{
				GroupVersionKind:  collections.IstioNetworkingV1Alpha3Workloadgroups.Resource().GroupVersionKind(),
				Name:              obj.Name,
				Namespace:         obj.Namespace,
				Labels:            obj.Labels,
				Annotations:       obj.Annotations,
				ResourceVersion:   obj.ResourceVersion,
				CreationTimestamp: obj.CreationTimestamp.Time,
			},
			Spec: &obj.Spec,
		}
	},
	collections.IstioSecurityV1Beta1Authorizationpolicies.Resource().GroupVersionKind(): func(r runtime.Object) *config.Config {
		obj := r.(*clientsecurityv1beta1.AuthorizationPolicy)
		return &config.Config{
			Meta: config.Meta{
				GroupVersionKind:  collections.IstioSecurityV1Beta1Authorizationpolicies.Resource().GroupVersionKind(),
				Name:              obj.Name,
				Namespace:         obj.Namespace,
				Labels:            obj.Labels,
				Annotations:       obj.Annotations,
				ResourceVersion:   obj.ResourceVersion,
				CreationTimestamp: obj.CreationTimestamp.Time,
			},
			Spec: &obj.Spec,
		}
	},
	collections.IstioSecurityV1Beta1Peerauthentications.Resource().GroupVersionKind(): func(r runtime.Object) *config.Config {
		obj := r.(*clientsecurityv1beta1.PeerAuthentication)
		return &config.Config{
			Meta: config.Meta{
				GroupVersionKind:  collections.IstioSecurityV1Beta1Peerauthentications.Resource().GroupVersionKind(),
				Name:              obj.Name,
				Namespace:         obj.Namespace,
				Labels:            obj.Labels,
				Annotations:       obj.Annotations,
				ResourceVersion:   obj.ResourceVersion,
				CreationTimestamp: obj.CreationTimestamp.Time,
			},
			Spec: &obj.Spec,
		}
	},
	collections.IstioSecurityV1Beta1Requestauthentications.Resource().GroupVersionKind(): func(r runtime.Object) *config.Config {
		obj := r.(*clientsecurityv1beta1.RequestAuthentication)
		return &config.Config{
			Meta: config.Meta{
				GroupVersionKind:  collections.IstioSecurityV1Beta1Requestauthentications.Resource().GroupVersionKind(),
				Name:              obj.Name,
				Namespace:         obj.Namespace,
				Labels:            obj.Labels,
				Annotations:       obj.Annotations,
				ResourceVersion:   obj.ResourceVersion,
				CreationTimestamp: obj.CreationTimestamp.Time,
			},
			Spec: &obj.Spec,
		}
	},
	collections.K8SServiceApisV1Alpha1Gatewayclasses.Resource().GroupVersionKind(): func(r runtime.Object) *config.Config {
		obj := r.(*servicev1alpha1.GatewayClass)
		return &config.Config{
			Meta: config.Meta{
				GroupVersionKind:  collections.K8SServiceApisV1Alpha1Gatewayclasses.Resource().GroupVersionKind(),
				Name:              obj.Name,
				Namespace:         obj.Namespace,
				Labels:            obj.Labels,
				Annotations:       obj.Annotations,
				ResourceVersion:   obj.ResourceVersion,
				CreationTimestamp: obj.CreationTimestamp.Time,
			},
			Spec: &obj.Spec,
		}
	},
	collections.K8SServiceApisV1Alpha1Gateways.Resource().GroupVersionKind(): func(r runtime.Object) *config.Config {
		obj := r.(*servicev1alpha1.Gateway)
		return &config.Config{
			Meta: config.Meta{
				GroupVersionKind:  collections.K8SServiceApisV1Alpha1Gateways.Resource().GroupVersionKind(),
				Name:              obj.Name,
				Namespace:         obj.Namespace,
				Labels:            obj.Labels,
				Annotations:       obj.Annotations,
				ResourceVersion:   obj.ResourceVersion,
				CreationTimestamp: obj.CreationTimestamp.Time,
			},
			Spec: &obj.Spec,
		}
	},
	collections.K8SServiceApisV1Alpha1Httproutes.Resource().GroupVersionKind(): func(r runtime.Object) *config.Config {
		obj := r.(*servicev1alpha1.HTTPRoute)
		return &config.Config{
			Meta: config.Meta{
				GroupVersionKind:  collections.K8SServiceApisV1Alpha1Httproutes.Resource().GroupVersionKind(),
				Name:              obj.Name,
				Namespace:         obj.Namespace,
				Labels:            obj.Labels,
				Annotations:       obj.Annotations,
				ResourceVersion:   obj.ResourceVersion,
				CreationTimestamp: obj.CreationTimestamp.Time,
			},
			Spec: &obj.Spec,
		}
	},
	collections.K8SServiceApisV1Alpha1Tcproutes.Resource().GroupVersionKind(): func(r runtime.Object) *config.Config {
		obj := r.(*servicev1alpha1.TCPRoute)
		return &config.Config{
			Meta: config.Meta{
				GroupVersionKind:  collections.K8SServiceApisV1Alpha1Tcproutes.Resource().GroupVersionKind(),
				Name:              obj.Name,
				Namespace:         obj.Namespace,
				Labels:            obj.Labels,
				Annotations:       obj.Annotations,
				ResourceVersion:   obj.ResourceVersion,
				CreationTimestamp: obj.CreationTimestamp.Time,
			},
			Spec: &obj.Spec,
		}
	},
}
