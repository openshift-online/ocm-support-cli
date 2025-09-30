package delete

import (
	"github.com/spf13/cobra"

	account "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/account"
	accountgroupassignments "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/account_group_assignments"
	accountgroups "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/account_groups"
	capabilities "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/capabilities"
	accountCapability "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/capabilities/account_capability"
	organizationCapability "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/capabilities/organization_capability"
	subscriptionCapability "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/capabilities/subscription_capability"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/capability"
	accountLabel "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/labels/account_label"
	organizationLabel "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/labels/organization_label"
	subscriptionLabel "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/labels/subscription_label"
	registryCredentials "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/registry_credentials"
	applicationRoleBinding "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/role_bindings/application_role_binding"
	organizationRoleBinding "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/role_bindings/organization_role_binding"
	subscriptionRoleBinding "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/role_bindings/subscription_role_binding"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "delete [COMMAND]",
	Short: "Removes Labels, Capabilities from Accounts, Subscriptions, Organizations",
	Long:  "Removes Labels, Capabilities from Accounts, Subscriptions, Organizations",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(account.CmdDeleteAccount)
	Cmd.AddCommand(accountgroups.CmdDeleteAccountGroup)
	Cmd.AddCommand(accountgroupassignments.CmdDeleteAccountGroupAssignment)
	Cmd.AddCommand(registryCredentials.CmdDeleteRegistryCredentials)
	Cmd.AddCommand(accountCapability.CmdDeleteAccountCapability)
	Cmd.AddCommand(organizationCapability.CmdDeleteOrganizationCapability)
	Cmd.AddCommand(subscriptionCapability.CmdDeleteSubscriptionCapability)
	Cmd.AddCommand(accountLabel.CmdDeleteAccountLabel)
	Cmd.AddCommand(organizationLabel.CmdDeleteOrganizationLabel)
	Cmd.AddCommand(subscriptionLabel.CmdDeleteSubscriptionLabel)
	Cmd.AddCommand(applicationRoleBinding.CmdDeleteApplicationRoleBinding)
	Cmd.AddCommand(organizationRoleBinding.CmdDeleteOrganizationRoleBinding)
	Cmd.AddCommand(subscriptionRoleBinding.CmdDeleteSubscriptionRoleBinding)
	Cmd.AddCommand(capabilities.CmdDeleteCapabilities)
	Cmd.AddCommand(capability.CmdDeleteCapability)
}
