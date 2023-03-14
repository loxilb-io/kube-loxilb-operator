/*
Copyright 2023.

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

package controllers

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	kubeloxilbv1alpha1 "loxilb-io/kube-loxilb-operator/api/v1alpha1"
	kubeloxilbdep "loxilb-io/kube-loxilb-operator/util/deployment"
)

// KubeloxilbappReconciler reconciles a Kubeloxilbapp object
type KubeloxilbappReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=kubeloxilb.loxilb.io,resources=kubeloxilbapps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=kubeloxilb.loxilb.io,resources=kubeloxilbapps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=kubeloxilb.loxilb.io,resources=kubeloxilbapps/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps,resources=nodes,verbs=get;list;watch;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Kubeloxilbapp object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *KubeloxilbappReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// TODO(user): your logic here
	k := &kubeloxilbv1alpha1.Kubeloxilbapp{}
	err := r.Client.Get(ctx, req.NamespacedName, k)
	if err != nil {
		// kube-loxilb deployment is deleted.
		if errors.IsNotFound(err) {
			logger.Info("This resource is deleted", "crd", req.NamespacedName)
			deleteDep := &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      req.Name,
					Namespace: req.Namespace,
				},
			}
			if err = r.Delete(ctx, deleteDep); err != nil {
				logger.Info("failed to delete resource", "deployment", req.NamespacedName)
				return ctrl.Result{}, err
			}
			logger.Info("success to delete resource", "deployment", req.NamespacedName)
			return ctrl.Result{}, nil
		}

		logger.Error(err, "Failed to get kube-loxilb (%s:%s)", req.Namespace, req.Name)
		return ctrl.Result{}, err
	}

	// check kube-loxilb deployment
	dep := &appsv1.Deployment{}
	if err := r.Client.Get(ctx, req.NamespacedName, dep); err != nil {
		if errors.IsNotFound(err) {
			//kube-loxilb deployment isn't created yet.
			newDep := kubeloxilbdep.GetKubeLoxilbDeployment(k)
			if err = r.Create(ctx, newDep); err != nil {
				return ctrl.Result{}, err
			}
			return ctrl.Result{Requeue: true}, nil
		} else {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KubeloxilbappReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kubeloxilbv1alpha1.Kubeloxilbapp{}).
		Owns(&appsv1.Deployment{}).
		Complete(r)
}
