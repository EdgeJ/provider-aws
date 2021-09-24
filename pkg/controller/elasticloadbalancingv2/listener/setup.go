package listener

import (
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	svcapitypes "github.com/crossplane/provider-aws/apis/elasticloadbalancingv2/v1alpha1"
)

// SetupListener adds a controller that reconciles Listener.
func SetupListener(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter, poll time.Duration) error {
	name := managed.ControllerName(svcapitypes.ListenerGroupKind)
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
		}).
		For(&svcapitypes.Listener{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(svcapitypes.ListenerGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient()}),
			managed.WithPollInterval(poll),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}
