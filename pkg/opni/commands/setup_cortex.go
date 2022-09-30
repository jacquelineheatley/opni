//go:build !noplugins

package commands

import (
	"github.com/rancher/opni/plugins/metrics/pkg/apis/cortexadmin"
	"github.com/rancher/opni/plugins/metrics/pkg/apis/cortexops"
	"github.com/spf13/cobra"
)

var adminClient cortexadmin.CortexAdminClient
var opsClient cortexops.CortexOpsClient

func ConfigureCortexAdminCommand(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("address", "a", "", "Management API address (default: auto-detect)")
	if cmd.PersistentPreRunE == nil {
		cmd.PersistentPreRunE = cortexAdminPreRunE
	} else {
		cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
			if err := cmd.PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			return cortexAdminPreRunE(cmd, args)
		}
	}
}

func cortexAdminPreRunE(cmd *cobra.Command, args []string) error {
	if managementListenAddress == "" {
		panic("bug: managementListenAddress is empty")
	}
	ac, err := cortexadmin.NewClient(cmd.Context(),
		cortexadmin.WithListenAddress(managementListenAddress))
	if err != nil {
		return err
	}
	adminClient = ac

	oc, err := cortexops.NewClient(cmd.Context(),
		cortexops.WithListenAddress(managementListenAddress))
	if err != nil {
		return err
	}
	opsClient = oc
	return nil
}
