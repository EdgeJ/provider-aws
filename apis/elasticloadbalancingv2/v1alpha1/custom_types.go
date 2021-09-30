package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// CustomRuleParameters includes the custom fields of Rule.
type CustomRuleParameters struct{}

type CustomTargetGroupTuple struct { // inject refs and selectors into TargetGroupTuple
	TargetGroupTuple `json:",inline"`

	// Reference to TargetGroupARN used to set TargetGroupARN
	// +optional
	TargetGroupARNRef *xpv1.Reference `json:"targetGroupARNRef,omitempty"`

	// Selector for references to TargetGroup for TargetGroupARN
	// +optional
	TargetGroupARNSelector *xpv1.Selector `json:"targetGroupARNSelector,omitempty"`
}

type CustomForwardActionConfig struct {
	// Information about the target group stickiness for a rule.
	TargetGroupStickinessConfig *TargetGroupStickinessConfig `json:"targetGroupStickinessConfig,omitempty"`

	TargetGroups []*CustomTargetGroupTuple `json:"targetGroups,omitempty"`
}

type CustomAction struct {
	// Request parameters to use when integrating with Amazon Cognito to authenticate
	// users.
	AuthenticateCognitoConfig *AuthenticateCognitoActionConfig `json:"authenticateCognitoConfig,omitempty"`
	// Request parameters when using an identity provider (IdP) that is compliant
	// with OpenID Connect (OIDC) to authenticate users.
	AuthenticateOidcConfig *AuthenticateOidcActionConfig `json:"authenticateOidcConfig,omitempty"`
	// Information about an action that returns a custom HTTP response.
	FixedResponseConfig *FixedResponseActionConfig `json:"fixedResponseConfig,omitempty"`
	// Information about a forward action.
	ForwardConfig *CustomForwardActionConfig `json:"forwardConfig,omitempty"`

	Order *int64 `json:"order,omitempty"`
	// Information about a redirect action.
	//
	// A URI consists of the following components: protocol://hostname:port/path?query.
	// You must modify at least one of the following components to avoid a redirect
	// loop: protocol, hostname, port, or path. Any components that you do not modify
	// retain their original values.
	//
	// You can reuse URI components using the following reserved keywords:
	//
	//    * #{protocol}
	//
	//    * #{host}
	//
	//    * #{port}
	//
	//    * #{path} (the leading "/" is removed)
	//
	//    * #{query}
	//
	// For example, you can change the path to "/new/#{path}", the hostname to "example.#{host}",
	// or the query to "#{query}&value=xyz".
	RedirectConfig *RedirectActionConfig `json:"redirectConfig,omitempty"`

	// The Amazon Resource Name (ARN) of the target group. Specify only when
	// actionType is forward and you want to route to a single target group.
	// To route to one or more target groups, use ForwardConfig instead.
	// +optional
	TargetGroupARN *string `json:"targetGroupARN,omitempty"`

	// Reference to TargetGroupARN used to set TargetGroupARN
	// +optional
	TargetGroupARNRef *xpv1.Reference `json:"targetGroupARNRef,omitempty"`

	// Selector for references to TargetGroups for TargetGroupARNs
	// +optional
	TargetGroupARNSelector *xpv1.Selector `json:"targetGroupARNSelector,omitempty"`

	Type *string `json:"actionType,omitempty"` // renamed json tag from "type_"
}

// CustomListenerParameters includes the custom fields of Listener.
type CustomListenerParameters struct {
	// The actions for the default rule.
	// +kubebuilder:validation:Required
	DefaultActions []*CustomAction `json:"defaultActions"`

	// The Amazon Resource Name (ARN) of the load balancer.
	// +immutable
	LoadBalancerARN *string `json:"loadBalancerARN,omitempty"`

	// Ref to loadbalancer ARN
	// +optional
	LoadBalancerARNRef *xpv1.Reference `json:"loadBalancerARNRef,omitempty"`

	// Selector for references to LoadBalancer for LoadBalancerARN
	// +optional
	LoadBalancerARNSelector *xpv1.Selector `json:"loadBalancerARNSelector,omitempty"`
}

// CustomLoadBalancerParameters includes the custom fields of LoadBalancer.
type CustomLoadBalancerParameters struct {
	// The type of load balancer. The default is application.
	Type *string `json:"loadBalancerType,omitempty"`
}

// CustomTargetGroupParameters includes the custom fields of TargetGroup.
type CustomTargetGroupParameters struct{}
