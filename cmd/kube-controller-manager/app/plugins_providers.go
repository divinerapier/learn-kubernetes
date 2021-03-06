// +build !providerless

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

package app

import (
	"github.com/divinerapier/learn-kubernetes/pkg/volume"
	"github.com/divinerapier/learn-kubernetes/pkg/volume/awsebs"
	"github.com/divinerapier/learn-kubernetes/pkg/volume/azure_dd"
	"github.com/divinerapier/learn-kubernetes/pkg/volume/azure_file"
	"github.com/divinerapier/learn-kubernetes/pkg/volume/cinder"
	"github.com/divinerapier/learn-kubernetes/pkg/volume/gcepd"
	"github.com/divinerapier/learn-kubernetes/pkg/volume/vsphere_volume"
)

func appendAttachableLegacyProviderVolumes(allPlugins []volume.VolumePlugin) []volume.VolumePlugin {
	allPlugins = append(allPlugins, awsebs.ProbeVolumePlugins()...)
	allPlugins = append(allPlugins, azure_dd.ProbeVolumePlugins()...)
	allPlugins = append(allPlugins, cinder.ProbeVolumePlugins()...)
	allPlugins = append(allPlugins, gcepd.ProbeVolumePlugins()...)
	allPlugins = append(allPlugins, vsphere_volume.ProbeVolumePlugins()...)
	return allPlugins
}

func appendExpandableLegacyProviderVolumes(allPlugins []volume.VolumePlugin) []volume.VolumePlugin {
	return appendLegacyProviderVolumes(allPlugins)
}

func appendLegacyProviderVolumes(allPlugins []volume.VolumePlugin) []volume.VolumePlugin {
	allPlugins = append(allPlugins, awsebs.ProbeVolumePlugins()...)
	allPlugins = append(allPlugins, azure_dd.ProbeVolumePlugins()...)
	allPlugins = append(allPlugins, azure_file.ProbeVolumePlugins()...)
	allPlugins = append(allPlugins, cinder.ProbeVolumePlugins()...)
	allPlugins = append(allPlugins, gcepd.ProbeVolumePlugins()...)
	allPlugins = append(allPlugins, vsphere_volume.ProbeVolumePlugins()...)
	return allPlugins
}
