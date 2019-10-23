/*
Copyright 2017 The Kubernetes Authors.

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
	"github.com/divinerapier/learn-kubernetes/pkg/apis/apps"
	"github.com/divinerapier/learn-kubernetes/pkg/printers"
	printersinternal "github.com/divinerapier/learn-kubernetes/pkg/printers/internalversion"
	printerstorage "github.com/divinerapier/learn-kubernetes/pkg/printers/storage"
	"github.com/divinerapier/learn-kubernetes/pkg/registry/apps/controllerrevision"
)

// REST implements a RESTStorage for ControllerRevision
type REST struct {
	*genericregistry.Store
}

// NewREST returns a RESTStorage object that will work with ControllerRevision objects.
func NewREST(optsGetter generic.RESTOptionsGetter) (*REST, error) {
	store := &genericregistry.Store{
		NewFunc:                  func() runtime.Object { return &apps.ControllerRevision{} },
		NewListFunc:              func() runtime.Object { return &apps.ControllerRevisionList{} },
		DefaultQualifiedResource: apps.Resource("controllerrevisions"),

		CreateStrategy: controllerrevision.Strategy,
		UpdateStrategy: controllerrevision.Strategy,
		DeleteStrategy: controllerrevision.Strategy,

		TableConvertor: printerstorage.TableConvertor{TableGenerator: printers.NewTableGenerator().With(printersinternal.AddHandlers)},
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &REST{store}, nil
}
