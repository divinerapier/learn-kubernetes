/*
Copyright 2018 The Kubernetes Authors.

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

package defaults

import (
	"github.com/divinerapier/learn-kubernetes/pkg/scheduler"
	"github.com/divinerapier/learn-kubernetes/pkg/scheduler/algorithm"
	"github.com/divinerapier/learn-kubernetes/pkg/scheduler/algorithm/priorities"
	"github.com/divinerapier/learn-kubernetes/pkg/scheduler/core"
)

func init() {
	// Register functions that extract metadata used by priorities computations.
	scheduler.RegisterPriorityMetadataProducerFactory(
		func(args scheduler.PluginFactoryArgs) priorities.PriorityMetadataProducer {
			return priorities.NewPriorityMetadataFactory(args.ServiceLister, args.ControllerLister, args.ReplicaSetLister, args.StatefulSetLister)
		})

	// ServiceSpreadingPriority is a priority config factory that spreads pods by minimizing
	// the number of pods (belonging to the same service) on the same node.
	// Register the factory so that it's available, but do not include it as part of the default priorities
	// Largely replaced by "SelectorSpreadPriority", but registered for backward compatibility with 1.0
	scheduler.RegisterPriorityConfigFactory(
		priorities.ServiceSpreadingPriority,
		scheduler.PriorityConfigFactory{
			MapReduceFunction: func(args scheduler.PluginFactoryArgs) (priorities.PriorityMapFunction, priorities.PriorityReduceFunction) {
				return priorities.NewSelectorSpreadPriority(args.ServiceLister, algorithm.EmptyControllerLister{}, algorithm.EmptyReplicaSetLister{}, algorithm.EmptyStatefulSetLister{})
			},
			Weight: 1,
		},
	)
	// EqualPriority is a prioritizer function that gives an equal weight of one to all nodes
	// Register the priority function so that its available
	// but do not include it as part of the default priorities
	scheduler.RegisterPriorityMapReduceFunction(priorities.EqualPriority, core.EqualPriorityMap, nil, 1)
	// Optional, cluster-autoscaler friendly priority function - give used nodes higher priority.
	scheduler.RegisterPriorityMapReduceFunction(priorities.MostRequestedPriority, priorities.MostRequestedPriorityMap, nil, 1)
	scheduler.RegisterPriorityMapReduceFunction(
		priorities.RequestedToCapacityRatioPriority,
		priorities.RequestedToCapacityRatioResourceAllocationPriorityDefault().PriorityMap,
		nil,
		1)
	// spreads pods by minimizing the number of pods (belonging to the same service or replication controller) on the same node.
	scheduler.RegisterPriorityConfigFactory(
		priorities.SelectorSpreadPriority,
		scheduler.PriorityConfigFactory{
			MapReduceFunction: func(args scheduler.PluginFactoryArgs) (priorities.PriorityMapFunction, priorities.PriorityReduceFunction) {
				return priorities.NewSelectorSpreadPriority(args.ServiceLister, args.ControllerLister, args.ReplicaSetLister, args.StatefulSetLister)
			},
			Weight: 1,
		},
	)
	// pods should be placed in the same topological domain (e.g. same node, same rack, same zone, same power domain, etc.)
	// as some other pods, or, conversely, should not be placed in the same topological domain as some other pods.
	scheduler.RegisterPriorityConfigFactory(
		priorities.InterPodAffinityPriority,
		scheduler.PriorityConfigFactory{
			Function: func(args scheduler.PluginFactoryArgs) priorities.PriorityFunction {
				return priorities.NewInterPodAffinityPriority(args.NodeInfo, args.HardPodAffinitySymmetricWeight)
			},
			Weight: 1,
		},
	)

	// Prioritize nodes by least requested utilization.
	scheduler.RegisterPriorityMapReduceFunction(priorities.LeastRequestedPriority, priorities.LeastRequestedPriorityMap, nil, 1)

	// Prioritizes nodes to help achieve balanced resource usage
	scheduler.RegisterPriorityMapReduceFunction(priorities.BalancedResourceAllocation, priorities.BalancedResourceAllocationMap, nil, 1)

	// Set this weight large enough to override all other priority functions.
	// TODO: Figure out a better way to do this, maybe at same time as fixing #24720.
	scheduler.RegisterPriorityMapReduceFunction(priorities.NodePreferAvoidPodsPriority, priorities.CalculateNodePreferAvoidPodsPriorityMap, nil, 10000)

	// Prioritizes nodes that have labels matching NodeAffinity
	scheduler.RegisterPriorityMapReduceFunction(priorities.NodeAffinityPriority, priorities.CalculateNodeAffinityPriorityMap, priorities.CalculateNodeAffinityPriorityReduce, 1)

	// Prioritizes nodes that marked with taint which pod can tolerate.
	scheduler.RegisterPriorityMapReduceFunction(priorities.TaintTolerationPriority, priorities.ComputeTaintTolerationPriorityMap, priorities.ComputeTaintTolerationPriorityReduce, 1)

	// ImageLocalityPriority prioritizes nodes that have images requested by the pod present.
	scheduler.RegisterPriorityMapReduceFunction(priorities.ImageLocalityPriority, priorities.ImageLocalityPriorityMap, nil, 1)
}
