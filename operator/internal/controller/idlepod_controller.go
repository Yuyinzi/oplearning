/*
Copyright 2024.

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

package controller

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	batchv1 "kube_stuff/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	finalizer = "idlepod.littlemay.io/finalizer"
)

// IdlePodReconciler reconciles a IdlePod object
type IdlePodReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=batch.littlemay.io,resources=idlePods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch.littlemay.io,resources=idlePods/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=batch.littlemay.io,resources=idlePods/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the IdlePod closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the IdlePod object against the actual IdlePod state, and then
// perform operations to make the IdlePod state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.0/pkg/reconcile
func (r *IdlePodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	var IdlePodResource batchv1.IdlePod
	fmt.Println(req.NamespacedName.Name)
	err := r.Get(ctx, req.NamespacedName, &IdlePodResource)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
	}
	// Check if the resource is being deleted
	if !IdlePodResource.ObjectMeta.DeletionTimestamp.IsZero() {
		// Resource is being deleted, handle the deletion process
		if controllerutil.ContainsFinalizer(&IdlePodResource, finalizer) {
			fmt.Println("start cleanup...")
			// do sth
			controllerutil.RemoveFinalizer(&IdlePodResource, finalizer)
			if err := r.Update(ctx, &IdlePodResource); err != nil {
				return ctrl.Result{}, err
			}
		}
		fmt.Println("Delete IdlePod resource")
	}

	// use definition to distinguish a crd temporarily
	if IdlePodResource.Spec.Definition != "hello world" {
		// Create a new Pod
		IdlePod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Labels:      make(map[string]string),
				Annotations: make(map[string]string),
				Name:        req.NamespacedName.Name,
				Namespace:   req.NamespacedName.Namespace,
			},
			Spec: *IdlePodResource.Spec.PodTemplate.Spec.DeepCopy(),
		}
		// when deleting the crd, k8s could delete corresponding pods automatically
		if err := controllerutil.SetControllerReference(&IdlePodResource, IdlePod, r.Scheme); err != nil {
			return ctrl.Result{}, err
		}
		controllerutil.AddFinalizer(&IdlePodResource, finalizer)
		if err := r.Create(ctx, IdlePod); err != nil {
			return ctrl.Result{}, err
		}
		fmt.Println("Create new pod")
		IdlePodResource.Spec.Definition = "hello world"
		if err := r.Update(ctx, &IdlePodResource); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *IdlePodReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1.IdlePod{}).
		Complete(r)
}
