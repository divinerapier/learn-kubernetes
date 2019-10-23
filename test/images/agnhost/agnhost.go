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

package main

import (
	"flag"

	"github.com/spf13/cobra"

	"k8s.io/klog"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/audit-proxy"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/connect"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/crd-conversion-webhook"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/dns"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/entrypoint-tester"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/fakegitserver"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/inclusterclient"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/liveness"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/logs-generator"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/net"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/netexec"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/nettest"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/no-snat-test"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/no-snat-test-proxy"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/pause"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/port-forward-tester"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/porter"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/serve-hostname"
	"github.com/divinerapier/learn-kubernetes/test/images/agnhost/webhook"
)

func main() {
	rootCmd := &cobra.Command{Use: "app", Version: "2.6"}

	rootCmd.AddCommand(auditproxy.CmdAuditProxy)
	rootCmd.AddCommand(connect.CmdConnect)
	rootCmd.AddCommand(crdconvwebhook.CmdCrdConversionWebhook)
	rootCmd.AddCommand(dns.CmdDNSSuffix)
	rootCmd.AddCommand(dns.CmdDNSServerList)
	rootCmd.AddCommand(dns.CmdEtcHosts)
	rootCmd.AddCommand(entrypoint.CmdEntrypointTester)
	rootCmd.AddCommand(fakegitserver.CmdFakeGitServer)
	rootCmd.AddCommand(inclusterclient.CmdInClusterClient)
	rootCmd.AddCommand(liveness.CmdLiveness)
	rootCmd.AddCommand(logsgen.CmdLogsGenerator)
	rootCmd.AddCommand(net.CmdNet)
	rootCmd.AddCommand(netexec.CmdNetexec)
	rootCmd.AddCommand(nettest.CmdNettest)
	rootCmd.AddCommand(nosnat.CmdNoSnatTest)
	rootCmd.AddCommand(nosnatproxy.CmdNoSnatTestProxy)
	rootCmd.AddCommand(pause.CmdPause)
	rootCmd.AddCommand(porter.CmdPorter)
	rootCmd.AddCommand(portforwardtester.CmdPortForwardTester)
	rootCmd.AddCommand(servehostname.CmdServeHostname)
	rootCmd.AddCommand(webhook.CmdWebhook)

	// NOTE(claudiub): Some tests are passing logging related flags, so we need to be able to
	// accept them. This will also include them in the printed help.
	loggingFlags := &flag.FlagSet{}
	klog.InitFlags(loggingFlags)
	rootCmd.PersistentFlags().AddGoFlagSet(loggingFlags)
	rootCmd.Execute()
}
