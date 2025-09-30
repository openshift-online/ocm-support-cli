package get

import (
	accountgroupassignments "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/get/account_group_assignments"
	accountgroups "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/get/account_groups"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/get/accounts"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/get/clusters"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/get/organizations"
	registrycredentials "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/get/registry_credentials"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/get/subscriptions"
	"github.com/spf13/cobra"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "get [COMMAND]",
	Short: "Gets the given resource",
	Long:  "Gets the given resource",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(accounts.CmdGetAccounts)
	Cmd.AddCommand(accountgroups.CmdGetAccountGroups)
	Cmd.AddCommand(accountgroupassignments.CmdGetAccountGroupAssignments)
	Cmd.AddCommand(organizations.CmdGetOrganizations)
	Cmd.AddCommand(registrycredentials.CmdGetRegistryCredentials)
	Cmd.AddCommand(subscriptions.CmdGetSubscriptions)
	Cmd.AddCommand(clusters.CmdGetClusters)
}
