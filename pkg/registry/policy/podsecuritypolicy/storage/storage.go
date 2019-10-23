/*
Copyright 2014 The Kubernetes Authors.

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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"github.com/divinerapier/learn-kubernetes/pkg/apis/policy"
	"github.com/divinerapier/learn-kubernetes/pkg/printers"
	printersinternal "github.com/divinerapier/learn-kubernetes/pkg/printers/internalversion"
	printerstorage "github.com/divinerapier/learn-kubernetes/pkg/printers/storage"
	"github.com/divinerapier/learn-kubernetes/pkg/registry/policy/podsecuritypolicy"
)

// REST implements a RESTStorage for PodSecurityPolicies.
type REST struct {
	*genericregistry.Store
}

// NewREST returns a RESTStorage object that will work against PodSecurityPolicy objects.
func NewREST(optsGetter generic.RESTOptionsGetter) (*REST, error) {
	store := &genericregistry.Store{
		NewFunc:                  func() runtime.Object { return &policy.PodSecurityPolicy{} },
		NewListFunc:              func() runtime.Object { return &policy.PodSecurityPolicyList{} },
		DefaultQualifiedResource: policy.Resource("podsecuritypolicies"),

		CreateStrategy:      podsecuritypolicy.Strategy,
		UpdateStrategy:      podsecuritypolicy.Strategy,
		DeleteStrategy:      podsecuritypolicy.Strategy,
		ReturnDeletedObject: true,

		TableConvertor: printerstorage.TableConvertor{TableGenerator: printers.NewTableGenerator().With(printersinternal.AddHandlers)},
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &REST{store}, nil
}

// ShortNames implements the ShortNamesProvider interface. Returns a list of short names for a resource.
func (r *REST) ShortNames() []string {
	return []string{"psp"}
}
