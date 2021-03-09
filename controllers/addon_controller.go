/*
Copyright 2020.

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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	addonv1alpha1 "github.com/apahim/addon-operator/api/v1alpha1"
)

// AddOnReconciler reconciles an AddOn object
type AddOnReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=addon.example.com,resources=addons,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=addon.example.com,resources=addons/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=addon.example.com,resources=addons/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;create;update;patch;delete;

func (r *AddOnReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("addon", req.NamespacedName)

	addon := &addonv1alpha1.AddOn{}
	err := r.Get(ctx, req.NamespacedName, addon)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Info("Addon resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		log.Error(err, "Failed to get Addon")
		return ctrl.Result{}, err
	}

	allowedNamespaces := []string{
		"addon-operator",
	}
	if !stringInSlice(addon.Namespace, allowedNamespaces) {
		log.Info("kind 'Addon' not in an allowed namespace. Ignoring.", "Addon.Namespace", addon.Namespace)
		return ctrl.Result{}, nil
	}

	found := &corev1.Namespace{}
	err = r.Get(ctx, types.NamespacedName{Name: addon.Spec.TargetNamespace}, found)
	if err != nil && errors.IsNotFound(err) {
		ns := r.namespaceForAddon(addon)
		log.Info("Creating a new Namespace", "Namespace.Name", ns.Name)

		// Fake messages
		log.Info("Creating the CatalogSource")
		log.Info("Creating the OperatorGroup")
		log.Info("Creating the Subscription")
		log.Info("Creating the Parameters ConfigMap")
		log.Info("Creating the PD and DMS Secret")
		// end of fake messages

		err = r.Create(ctx, ns)
		if err != nil {
			log.Error(err, "Failed to create new Namespace", "Namespace.Name", ns.Name)
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		log.Error(err, "Failed to get Namespace")
		return ctrl.Result{}, err
	}

	if addon.Status.NamespaceStatus != found.Name {
		addon.Status.NamespaceStatus = string(found.Status.Phase)
		err := r.Status().Update(ctx, addon)
		if err != nil {
			log.Error(err, "Failed to update Addon status")
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (r *AddOnReconciler) namespaceForAddon(m *addonv1alpha1.AddOn) *corev1.Namespace {
	namespaceName := m.Spec.TargetNamespace

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespaceName,
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: m.APIVersion,
					Kind:       m.Kind,
					Name:       m.Name,
					UID:        m.UID,
				},
			},
		},
	}

	ctrl.SetControllerReference(m, ns, r.Scheme)
	return ns
}

func (r *AddOnReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&addonv1alpha1.AddOn{}).
		Owns(&corev1.Namespace{}).
		Complete(r)
}
