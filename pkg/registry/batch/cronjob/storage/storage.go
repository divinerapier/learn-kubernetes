/*
Copyright 2016 The Kubernetes Authors.

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

package storage

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"github.com/divinerapier/learn-kubernetes/pkg/apis/batch"
	"github.com/divinerapier/learn-kubernetes/pkg/printers"
	printersinternal "github.com/divinerapier/learn-kubernetes/pkg/printers/internalversion"
	printerstorage "github.com/divinerapier/learn-kubernetes/pkg/printers/storage"
	"github.com/divinerapier/learn-kubernetes/pkg/registry/batch/cronjob"
)

// REST implements a RESTStorage for scheduled jobs against etcd
type REST struct {
	*genericregistry.Store
}

// NewREST returns a RESTStorage object that will work against CronJobs.
func NewREST(optsGetter generic.RESTOptionsGetter) (*REST, *StatusREST, error) {
	store := &genericregistry.Store{
		NewFunc:                  func() runtime.Object { return &batch.CronJob{} },
		NewListFunc:              func() runtime.Object { return &batch.CronJobList{} },
		DefaultQualifiedResource: batch.Resource("cronjobs"),

		CreateStrategy: cronjob.Strategy,
		UpdateStrategy: cronjob.Strategy,
		DeleteStrategy: cronjob.Strategy,

		TableConvertor: printerstorage.TableConvertor{TableGenerator: printers.NewTableGenerator().With(printersinternal.AddHandlers)},
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, nil, err
	}

	statusStore := *store
	statusStore.UpdateStrategy = cronjob.StatusStrategy

	return &REST{store}, &StatusREST{store: &statusStore}, nil
}

var _ rest.CategoriesProvider = &REST{}
var _ rest.ShortNamesProvider = &REST{}

// Categories implements the CategoriesProvider interface. Returns a list of categories a resource is part of.
func (r *REST) Categories() []string {
	return []string{"all"}
}

// ShortNames implements the ShortNamesProvider interface. Returns a list of short names for a resource.
func (r *REST) ShortNames() []string {
	return []string{"cj"}
}

// StatusREST implements the REST endpoint for changing the status of a resourcequota.
type StatusREST struct {
	store *genericregistry.Store
}

// New creates a new CronJob object.
func (r *StatusREST) New() runtime.Object {
	return &batch.CronJob{}
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return r.store.Get(ctx, name, options)
}

// Update alters the status subset of an object.
func (r *StatusREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	// We are explicitly setting forceAllowCreate to false in the call to the underlying storage because
	// subresources should never allow create on update.
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, false, options)
}
