package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// CustomRuleParameters includes the custom fields of Rule.
type CustomRuleParameters struct{}

// CustomListenerParameters includes the custom fields of Listener.
type CustomListenerParameters struct {
	// The Amazon Resource Name (ARN) of the load balancer.
	// +immutable
	LoadBalancerARN *string `json:"loadBalancerARN,omitempty"`

	// Ref to loadbalancer ARN
	// +optional
	LoadBalancerARNRef *xpv1.Reference `json:"loadBalancerArnRef,omitempty"`

	// Selector for references to LoadBalancer for LoadBalancerARN
	// +optional
	LoadBalancerARNSelector *xpv1.Selector `json:"loadBalancerArnSelector,omitempty"`
}

// CustomLoadBalancerParameters includes the custom fields of LoadBalancer.
type CustomLoadBalancerParameters struct {
	// The type of load balancer. The default is application.
	Type *string `json:"loadBalancerType,omitempty"`
}

// CustomTargetGroupParameters includes the custom fields of TargetGroup.
type CustomTargetGroupParameters struct{}
