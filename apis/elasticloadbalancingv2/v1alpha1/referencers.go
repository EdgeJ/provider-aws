/*
Copyright 2021 The Crossplane Authors.
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

package v1alpha1

import (
	"context"
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Listener) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	// resolve loadbalancer ARN reference
	rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancerARN),
		Reference:    mg.Spec.ForProvider.LoadBalancerARNRef,
		Selector:     mg.Spec.ForProvider.LoadBalancerARNSelector,
		To:           reference.To{Managed: &LoadBalancer{}, List: &LoadBalancerList{}},
		Extract:      reference.ExternalName(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.forProvider.loadBalancerArn")
	}
	mg.Spec.ForProvider.LoadBalancerARN = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerARNRef = rsp.ResolvedReference

	for i := range mg.Spec.ForProvider.DefaultActions {
		// resolve single target group ARN references for each default action
		rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultActions[i].TargetGroupARN),
			Reference:    mg.Spec.ForProvider.DefaultActions[i].TargetGroupARNRef,
			Selector:     mg.Spec.ForProvider.DefaultActions[i].TargetGroupARNSelector,
			To:           reference.To{Managed: &TargetGroup{}, List: &TargetGroupList{}},
			Extract:      reference.ExternalName(),
		})
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("spec.forProvider.DefaultActions[%d].targetGroupArn", i))
		}

		mg.Spec.ForProvider.DefaultActions[i].TargetGroupARN = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DefaultActions[i].TargetGroupARNRef = rsp.ResolvedReference

		// resolve target group ARN references in forwardconfig if there are any
		if mg.Spec.ForProvider.DefaultActions[i].ForwardConfig != nil {
			for j := range mg.Spec.ForProvider.DefaultActions[i].ForwardConfig.TargetGroups {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultActions[i].ForwardConfig.TargetGroups[j].TargetGroupARN),
					Reference:    mg.Spec.ForProvider.DefaultActions[i].ForwardConfig.TargetGroups[j].TargetGroupARNRef,
					Selector:     mg.Spec.ForProvider.DefaultActions[i].ForwardConfig.TargetGroups[j].TargetGroupARNSelector,
					To:           reference.To{Managed: &TargetGroup{}, List: &TargetGroupList{}},
					Extract:      reference.ExternalName(),
				})
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("spec.forProvider.DefaultActions[%d].forwardConfig.targetGroups[%d]", i, j))
				}

				mg.Spec.ForProvider.DefaultActions[i].ForwardConfig.TargetGroups[j].TargetGroupARN = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.DefaultActions[i].ForwardConfig.TargetGroups[j].TargetGroupARNRef = rsp.ResolvedReference
			}
		}
	}

	return nil
}

func (mg *LoadBalancer) ResolveReferences(ctx context.Context, c client.Reader) error {
	return nil
}

func (mg *TargetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	return nil
}
