/*
Copyright 2019 The Kubernetes Authors.

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
	"github.com/divinerapier/learn-kubernetes/pkg/apis/discovery"
	"github.com/divinerapier/learn-kubernetes/pkg/registry/discovery/endpointslice"
)

// REST implements a RESTStorage for EndpointSlice against etcd
type REST struct {
	*genericregistry.Store
}

// NewREST returns a RESTStorage object that will work against endpoint slices.
func NewREST(optsGetter generic.RESTOptionsGetter) (*REST, error) {
	store := &genericregistry.Store{
		NewFunc:     func() runtime.Object { return &discovery.EndpointSlice{} },
		NewListFunc: func() runtime.Object { return &discovery.EndpointSliceList{} },
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*discovery.EndpointSlice).Name, nil
		},
		DefaultQualifiedResource: discovery.Resource("endpointslices"),

		CreateStrategy: endpointslice.Strategy,
		UpdateStrategy: endpointslice.Strategy,
		DeleteStrategy: endpointslice.Strategy,
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &REST{store}, nil
}
