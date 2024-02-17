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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	pacv1alpha1 "github.com/openshift-pipelines/pipelines-as-code/pkg/apis/pipelinesascode/v1alpha1"
)

type Repository struct {
	*pacv1alpha1.Repository
}

// log is for logging in this package.
var repositorylog = logf.Log.WithName("repository-resource")

// SetupWebhookWithManager will setup the manager to manage the webhooks
func (r *Repository) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-pipelinesascode-tekton-dev-appstudio-redhat-com-v1alpha1-repository,mutating=false,failurePolicy=fail,sideEffects=None,groups=pipelinesascode.tekton.dev.appstudio.redhat.com,resources=repositories,verbs=create;update,versions=v1alpha1,name=vrepository.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Repository{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Repository) ValidateCreate() (admission.Warnings, error) {
	repositorylog.Info("validate create", "name", r.Name)
	// r.Spec.URL

	// TODO(user): fill in your validation logic upon object creation.
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Repository) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	repositorylog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Repository) ValidateDelete() (admission.Warnings, error) {
	repositorylog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
