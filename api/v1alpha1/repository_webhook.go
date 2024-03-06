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
	"context"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	pacv1alpha1 "github.com/openshift-pipelines/pipelines-as-code/pkg/apis/pipelinesascode/v1alpha1"
)

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-pipelinesascode-tekton-dev-appstudio-redhat-com-v1alpha1-repository,mutating=false,failurePolicy=fail,sideEffects=None,groups=pipelinesascode.tekton.dev.appstudio.redhat.com,resources=repositories,verbs=create;update,versions=v1alpha1,name=vrepository.kb.io,admissionReviewVersions=v1

type URLValidator struct {
	URLPrefixAllowList []string
}

type URLValidationFailedError struct{}

func (u *URLValidationFailedError) Error() string {
	return "failed to validate url error"
}

func (u *URLValidator) Validate(url string) (admission.Warnings, error) {
	for _, urlPrefix := range u.URLPrefixAllowList {
		if strings.HasPrefix(url, urlPrefix) {
			return nil, nil
		}
	}

	warning := admission.Warnings{
		fmt.Sprintf(
			"URL %s is not in the allowed list. URL must start with one of: %v",
			url,
			u.URLPrefixAllowList,
		),
	}

	return warning, &URLValidationFailedError{}
}

var _ webhook.CustomValidator = &RepositoryValidator{}

type RepositoryValidator struct {
	UrlValidator URLValidator
	Logger       logr.Logger
}

// ValidateCreate implements admission.CustomValidator.
func (r *RepositoryValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (warnings admission.Warnings, err error) {
	repo, ok := obj.(*pacv1alpha1.Repository)
	if !ok {
		gvk := obj.GetObjectKind().GroupVersionKind().String()
		r.Logger.Info("Got $s, while expecting to get a Repository, skipping validation", gvk)
		return nil, nil
	}
	return r.UrlValidator.Validate(repo.Spec.URL)
}

// ValidateDelete implements admission.CustomValidator.
func (r *RepositoryValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (warnings admission.Warnings, err error) {
	return nil, nil
}

// ValidateUpdate implements admission.CustomValidator.
func (r *RepositoryValidator) ValidateUpdate(ctx context.Context, oldObj runtime.Object, newObj runtime.Object) (warnings admission.Warnings, err error) {
	repo, ok := newObj.(*pacv1alpha1.Repository)
	if !ok {
		gvk := newObj.GetObjectKind().GroupVersionKind().String()
		r.Logger.Info("Got $s, while expecting to get a Repository, skipping validation", gvk)
		return nil, nil
	}
	return r.UrlValidator.Validate(repo.Spec.URL)
}

// SetupWebhookWithManager will setup the manager to manage the webhooks
func SetupWebhookWithManager(mgr ctrl.Manager, validator webhook.CustomValidator) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(&pacv1alpha1.Repository{}).
		WithValidator(validator).
		Complete()
}
