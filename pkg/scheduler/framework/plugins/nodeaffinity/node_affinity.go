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

package nodeaffinity

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/divinerapier/learn-kubernetes/pkg/scheduler/algorithm/predicates"
	"github.com/divinerapier/learn-kubernetes/pkg/scheduler/algorithm/priorities"
	"github.com/divinerapier/learn-kubernetes/pkg/scheduler/framework/plugins/migration"
	framework "github.com/divinerapier/learn-kubernetes/pkg/scheduler/framework/v1alpha1"
	"github.com/divinerapier/learn-kubernetes/pkg/scheduler/nodeinfo"
)

// NodeAffinity is a plugin that checks if a pod node selector matches the node label.
type NodeAffinity struct {
	handle framework.FrameworkHandle
}

var _ framework.FilterPlugin = &NodeAffinity{}
var _ framework.ScorePlugin = &NodeAffinity{}

// Name is the name of the plugin used in the plugin registry and configurations.
const Name = "NodeAffinity"

// Name returns name of the plugin. It is used in logs, etc.
func (pl *NodeAffinity) Name() string {
	return Name
}

// Filter invoked at the filter extension point.
func (pl *NodeAffinity) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *nodeinfo.NodeInfo) *framework.Status {
	_, reasons, err := predicates.PodMatchNodeSelector(pod, nil, nodeInfo)
	return migration.PredicateResultToFrameworkStatus(reasons, err)
}

// Score invoked at the Score extension point.
func (pl *NodeAffinity) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	nodeInfo, exist := pl.handle.NodeInfoSnapshot().NodeInfoMap[nodeName]
	if !exist {
		return 0, framework.NewStatus(framework.Error, fmt.Sprintf("node %q does not exist in NodeInfoSnapshot", nodeName))
	}

	meta := migration.PriorityMetadata(state)
	s, err := priorities.CalculateNodeAffinityPriorityMap(pod, meta, nodeInfo)
	return s.Score, migration.ErrorToFrameworkStatus(err)
}

// NormalizeScore invoked after scoring all nodes.
func (pl *NodeAffinity) NormalizeScore(ctx context.Context, state *framework.CycleState, pod *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	// Note that CalculateNodeAffinityPriorityReduce doesn't use priority metadata, hence passing nil here.
	err := priorities.CalculateNodeAffinityPriorityReduce(pod, nil, pl.handle.NodeInfoSnapshot().NodeInfoMap, scores)
	return migration.ErrorToFrameworkStatus(err)
}

// ScoreExtensions of the Score plugin.
func (pl *NodeAffinity) ScoreExtensions() framework.ScoreExtensions {
	return pl
}

// New initializes a new plugin and returns it.
func New(_ *runtime.Unknown, h framework.FrameworkHandle) (framework.Plugin, error) {
	return &NodeAffinity{handle: h}, nil
}
