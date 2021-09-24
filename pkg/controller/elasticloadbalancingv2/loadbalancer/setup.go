package loadbalancer

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/elbv2"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	svcapitypes "github.com/crossplane/provider-aws/apis/elasticloadbalancingv2/v1alpha1"
)

// SetupLoadBalancer adds a controller that reconciles LoadBalancer.
func SetupLoadBalancer(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter, poll time.Duration) error {
	name := managed.ControllerName(svcapitypes.LoadBalancerGroupKind)
	opts := []option{
		func(e *external) {
			e.preObserve = preObserve
			e.postObserve = postObserve
			e.preCreate = preCreate
			// e.postCreate = postCreate
			e.preDelete = preDelete
		},
	}
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
		}).
		For(&svcapitypes.LoadBalancer{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(svcapitypes.LoadBalancerGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient(), opts: opts}),
			managed.WithPollInterval(poll),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

func preObserve(_ context.Context, cr *svcapitypes.LoadBalancer, obj *svcsdk.DescribeLoadBalancersInput) error {
	obj.Names = []*string{aws.String(meta.GetExternalName(cr))}
	return nil
}

func postObserve(_ context.Context, cr *svcapitypes.LoadBalancer, _ *svcsdk.DescribeLoadBalancersOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	if err != nil {
		return managed.ExternalObservation{}, err
	}
	cr.SetConditions(xpv1.Available())
	return obs, nil
}

func preCreate(_ context.Context, cr *svcapitypes.LoadBalancer, obj *svcsdk.CreateLoadBalancerInput) error {
	if obj.Name == nil {
		return errors.New("Loadbalancer Name parameter is empty!")
	}
	return nil
}

func postCreate(_ context.Context, cr *svcapitypes.LoadBalancer, resp *svcsdk.CreateLoadBalancerOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	if err != nil {
		return managed.ExternalCreation{}, err
	}
	if len(resp.LoadBalancers) != 1 {
		return managed.ExternalCreation{}, errors.New("Too many API objects returned")
	}
	meta.SetExternalName(cr, aws.StringValue(resp.LoadBalancers[0].LoadBalancerArn))
	cre.ExternalNameAssigned = true
	return cre, nil
}

func preDelete(_ context.Context, cr *svcapitypes.LoadBalancer, obj *svcsdk.DeleteLoadBalancerInput) (bool, error) {
	obj.LoadBalancerArn = aws.String(meta.GetExternalName(cr))
	return false, nil
}
