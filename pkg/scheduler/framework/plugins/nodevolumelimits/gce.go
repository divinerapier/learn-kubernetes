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

package nodevolumelimits

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/divinerapier/learn-kubernetes/pkg/scheduler/algorithm/predicates"
	"github.com/divinerapier/learn-kubernetes/pkg/scheduler/framework/plugins/migration"
	framework "github.com/divinerapier/learn-kubernetes/pkg/scheduler/framework/v1alpha1"
	"github.com/divinerapier/learn-kubernetes/pkg/scheduler/nodeinfo"
)

// GCEPDLimits is a plugin that checks node volume limits.
type GCEPDLimits struct {
	predicate predicates.FitPredicate
}

var _ framework.FilterPlugin = &GCEPDLimits{}

// GCEPDName is the name of the plugin used in the plugin registry and configurations.
const GCEPDName = "GCEPDLimits"

// Name returns name of the plugin. It is used in logs, etc.
func (pl *GCEPDLimits) Name() string {
	return GCEPDName
}

// Filter invoked at the filter extension point.
func (pl *GCEPDLimits) Filter(ctx context.Context, _ *framework.CycleState, pod *v1.Pod, nodeInfo *nodeinfo.NodeInfo) *framework.Status {
	// metadata is not needed
	_, reasons, err := pl.predicate(pod, nil, nodeInfo)
	return migration.PredicateResultToFrameworkStatus(reasons, err)
}

// NewGCEPD returns function that initializes a new plugin and returns it.
func NewGCEPD(_ *runtime.Unknown, handle framework.FrameworkHandle) (framework.Plugin, error) {
	informerFactory := handle.SharedInformerFactory()
	csiNodeInfo := &predicates.CachedCSINodeInfo{
		CSINodeLister: informerFactory.Storage().V1beta1().CSINodes().Lister(),
	}
	pvInfo := &predicates.CachedPersistentVolumeInfo{
		PersistentVolumeLister: informerFactory.Core().V1().PersistentVolumes().Lister(),
	}
	pvcInfo := &predicates.CachedPersistentVolumeClaimInfo{
		PersistentVolumeClaimLister: informerFactory.Core().V1().PersistentVolumeClaims().Lister(),
	}
	classInfo := &predicates.CachedStorageClassInfo{
		StorageClassLister: informerFactory.Storage().V1().StorageClasses().Lister(),
	}

	return &GCEPDLimits{
		predicate: predicates.NewMaxPDVolumeCountPredicate(predicates.GCEPDVolumeFilterType, csiNodeInfo, classInfo, pvInfo, pvcInfo),
	}, nil
}
