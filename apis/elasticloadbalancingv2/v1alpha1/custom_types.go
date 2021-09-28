package v1alpha1

// CustomRuleParameters includes the custom fields of Rule.
type CustomRuleParameters struct{}

// CustomListenerParameters includes the custom fields of Listener.
type CustomListenerParameters struct{}

// CustomLoadBalancerParameters includes the custom fields of LoadBalancer.
type CustomLoadBalancerParameters struct {
	// The type of load balancer. The default is application.
	Type *string `json:"loadBalancerType,omitempty"`
}

// CustomTargetGroupParameters includes the custom fields of TargetGroup.
type CustomTargetGroupParameters struct{}
