/*
Copyright 2022.

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

package hermes

import (
	hermesv1 "Hermes/pkg/adaptor/apis/hermes/v1"
	realtimemprocess "Hermes/pkg/realtimeprocess"
	"context"
	"fmt"
	v12 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	REALTIMESUB            = "SubsRealTime"
	CONDITIONSUB           = "SubsCondition"
	PROMETHEUSRULNAMESPACE = "hypermonitor"
)

// SubscriberRuleReconciler reconciles a SubscriberRule object
type SubscriberRuleReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=hermes.pml.com,resources=subscriberrules,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=hermes.pml.com,resources=subscriberrules/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=hermes.pml.com,resources=subscriberrules/finalizers,verbs=update
//+kubebuilder:rbac:groups=monitoring.coreos.com,resources=prometheusrules,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SubscriberRule object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *SubscriberRuleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	subscribeRule := &hermesv1.SubscriberRule{}
	err := r.Get(ctx, req.NamespacedName, subscribeRule)
	if err != nil {
		if errors.IsNotFound(err) {
			logger.Info(fmt.Sprintf("%s not found, maybe removed", req.Name))
			if realtimemprocess.IsInGlobalSrs(req.Name) {
				realtimemprocess.DeleteGlobalSrs(req.Name)
			} else {
				logger.Info(fmt.Sprintf("Delete PrometheusRule %s", subscribeRule.Name))
				//rule := makePrometheusRule(subscribeRule)
				//err := r.Delete(ctx, &rule)
				//if err != nil {
				//	logger.Error(err, "delete rule fail")
				//}
			}

			return ctrl.Result{}, nil
		}
		logger.Error(err, "unknown error")
	}
	logger.Info("Get object successful")

	if subscribeRule.Spec.SubscribeType == REALTIMESUB {
		realtimemprocess.AddGlobalSrs(subscribeRule.Name, subscribeRule.Spec.SubscriberAddress, subscribeRule.Spec.RealTimeMetricList)
	} else if subscribeRule.Spec.SubscribeType == CONDITIONSUB {
		promRule := &v12.PrometheusRule{}
		rule := makePrometheusRule(subscribeRule)
		err := r.Get(ctx, req.NamespacedName, promRule)
		if err != nil {
			if errors.IsNotFound(err) {
				err = r.Create(ctx, &rule)
				if err != nil {
					logger.Error(err, "Create rule fail")
				}
				logger.Info(fmt.Sprintf("Create PrometheusRule %s success.", subscribeRule.Name))
				return ctrl.Result{}, nil
			}
		}
		rule.Spec.DeepCopyInto(&promRule.Spec)
		err = r.Update(ctx, promRule)
		if err != nil {
			logger.Error(err, "Update rule fail")
		}
	} else {
		logger.Info("Add SubscribeType error")
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}

func makePrometheusRule(sr *hermesv1.SubscriberRule) v12.PrometheusRule {
	rule := v12.PrometheusRule{
		TypeMeta: v1.TypeMeta{
			Kind:       "PrometheusRule",
			APIVersion: "monitoring.coreos.com/v1",
		},
	}
	rule.Annotations = make(map[string]string)
	rule.Annotations["meta.helm.sh/release-name"] = "prometheus"
	rule.Annotations["meta.helm.sh/release-namespace"] = "default"
	rule.Annotations["prometheus-operator-validated"] = "true"
	rule.Labels = make(map[string]string)
	rule.Labels["app"] = "kube-prometheus-stack"
	rule.Labels["app.kubernetes.io/instance"] = "prometheus"
	rule.Labels["app.kubernetes.io/managed-by"] = "Helm"
	rule.Labels["app.kubernetes.io/part-of"] = "kube-prometheus-stack"
	rule.Labels["app.kubernetes.io/version"] = "0.4.0"
	rule.Labels["chart"] = "kube-prometheus-stack-0.4.0"
	rule.Labels["heritage"] = "Helm"
	rule.Labels["release"] = "prometheus"
	rule.Name = sr.Name
	rule.Namespace = PROMETHEUSRULNAMESPACE
	targetSpec := sr.Spec.PrometheusRule
	targetSpec.DeepCopyInto(&rule.Spec)
	rule.ObjectMeta.OwnerReferences = []v1.OwnerReference{
		*v1.NewControllerRef(sr, hermesv1.GroupVersion.WithKind("SubscriberRule")),
	} // TODO add ownerreferences
	return rule
}

// SetupWithManager sets up the controller with the Manager.
func (r *SubscriberRuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&hermesv1.SubscriberRule{}).
		Complete(r)
}
