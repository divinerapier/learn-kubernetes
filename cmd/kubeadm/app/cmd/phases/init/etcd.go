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

package phases

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"k8s.io/klog"
	"github.com/divinerapier/learn-kubernetes/cmd/kubeadm/app/cmd/options"
	"github.com/divinerapier/learn-kubernetes/cmd/kubeadm/app/cmd/phases/workflow"
	cmdutil "github.com/divinerapier/learn-kubernetes/cmd/kubeadm/app/cmd/util"
	etcdphase "github.com/divinerapier/learn-kubernetes/cmd/kubeadm/app/phases/etcd"
)

var (
	etcdLocalExample = cmdutil.Examples(`
		# Generates the static Pod manifest file for etcd, functionally
		# equivalent to what is generated by kubeadm init.
		kubeadm init phase etcd local

		# Generates the static Pod manifest file for etcd using options
		# read from a configuration file.
		kubeadm init phase etcd local --config config.yaml
		`)
)

// NewEtcdPhase creates a kubeadm workflow phase that implements handling of etcd.
func NewEtcdPhase() workflow.Phase {
	phase := workflow.Phase{
		Name:  "etcd",
		Short: "Generate static Pod manifest file for local etcd",
		Long:  cmdutil.MacroCommandLongDescription,
		Phases: []workflow.Phase{
			newEtcdLocalSubPhase(),
		},
	}
	return phase
}

func newEtcdLocalSubPhase() workflow.Phase {
	phase := workflow.Phase{
		Name:         "local",
		Short:        "Generate the static Pod manifest file for a local, single-node local etcd instance",
		Example:      etcdLocalExample,
		Run:          runEtcdPhaseLocal(),
		InheritFlags: getEtcdPhaseFlags(),
	}
	return phase
}

func getEtcdPhaseFlags() []string {
	flags := []string{
		options.CertificatesDir,
		options.CfgPath,
		options.ImageRepository,
		options.Kustomize,
	}
	return flags
}

func runEtcdPhaseLocal() func(c workflow.RunData) error {
	return func(c workflow.RunData) error {
		data, ok := c.(InitData)
		if !ok {
			return errors.New("etcd phase invoked with an invalid data struct")
		}
		cfg := data.Cfg()

		// Add etcd static pod spec only if external etcd is not configured
		if cfg.Etcd.External == nil {
			// creates target folder if doesn't exist already
			if !data.DryRun() {
				if err := os.MkdirAll(cfg.Etcd.Local.DataDir, 0700); err != nil {
					return errors.Wrapf(err, "failed to create etcd directory %q", cfg.Etcd.Local.DataDir)
				}
			} else {
				fmt.Printf("[dryrun] Would ensure that %q directory is present\n", cfg.Etcd.Local.DataDir)
			}
			fmt.Printf("[etcd] Creating static Pod manifest for local etcd in %q\n", data.ManifestDir())
			if err := etcdphase.CreateLocalEtcdStaticPodManifestFile(data.ManifestDir(), data.KustomizeDir(), cfg.NodeRegistration.Name, &cfg.ClusterConfiguration, &cfg.LocalAPIEndpoint); err != nil {
				return errors.Wrap(err, "error creating local etcd static pod manifest file")
			}
		} else {
			klog.V(1).Infoln("[etcd] External etcd mode. Skipping the creation of a manifest for local etcd")
		}
		return nil
	}
}
