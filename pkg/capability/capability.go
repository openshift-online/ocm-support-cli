package capability

import (
	"fmt"
	"strings"

	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
)

type Capability struct {
	Name      string
	Value     string
	Inherited bool
}

type CapabilityList []Capability

// A list of known capabilities and notes on enabling them can be found on
// https://docs.google.com/spreadsheets/d/1XMFkbG9DeXoBSOn5HPbzk-l2T_OjQdLjz452a6vaYCE

// Capabilities must contain 3 sections, separated by "."
// capability.{type}.{name}
const CapabilityAggressiveClusterSetup = "capability.account.aggressive_cluster_cleanup"
const CapabilityCreateMoaClusters = "capability.account.create_moa_clusters"
const CapabilityManageClusterAdmin = "capability.cluster.manage_cluster_admin"
const CapabilityOrganizationRegistrationsPerHour = "capability.organization.clusters_registrations_per_hour"
const CapabilityOrganizationPinClusterToShard = "capability.organization.pin_cluster_to_shard"
const CapabilityHibernateCluster = "capability.organization.hibernate_cluster"
const CapabilitySubscribedOcp = "capability.cluster.subscribed_ocp"
const CapabilitySubscribedOcpMarketplace = "capability.cluster.subscribed_ocp_marketplace"
const CapabilitySubscribedOsdMarketplace = "capability.cluster.subscribed_osd_marketplace"
const CapabilityEnableTermsEnforcement = "capability.account.enable_terms_enforcement"
const CapabilityBareMetalInstallerAdmin = "capability.account.bare_metal_installer_admin"
const CapabilityReleaseOcpClusters = "capability.cluster.release_ocp_clusters"
const CapabilityAutoscaleClustersDeprecated = "capability.organization.autoscale_clusters"
const CapabilityAutoscaleClusters = "capability.cluster.autoscale_clusters"
const CapabilityOrganizationInstallConfigOverride = "capability.organization.install_config_override"
const CapabilityOrganizationInstallConfigDefault = "capability.organization.install_config_default"
const CapabilityOrganizationOverrideOsdTrialLength = "capability.organization.override_osdtrial_length_days"
const CapabilityOrganizationCreateClusterProxy = "capability.organization.create_cluster_proxy"
const CapabilityAllowGCPNonCCSPrivateClusters = "capability.organization.create_gcp_non_ccs_cluster"
const CapabilityAllowInstallEOLVersions = "capability.organization.allow_install_eol_versions"
const CapabilityAddOnVersionSelect = "capability.organization.addon_version_select"
const CapabilityOrganizationFipsCluster = "capability.organization.fips_cluster"
const CapabilityOrganizationOvnCluster = "capability.organization.ovn_cluster"
const CapabilityOrganizationHyperShift = "capability.organization.hypershift"
const CapabilityBypassMaxExpiration = "capability.organization.bypass_max_expiration"

// This "use_rosa_paid_ami" capability will only be used short-term and will eventually be removed
// See https://issues.redhat.com/browse/SDB-1698
const CapabilityUseRosaPaidAMI = "capability.account.use_rosa_paid_ami"

var availableCapabilities map[string]string = map[string]string{
	"CapabilityAggressiveClusterSetup":             CapabilityAggressiveClusterSetup,
	"CapabilityCreateMoaClusters":                  CapabilityCreateMoaClusters,
	"CapabilityManageClusterAdmin":                 CapabilityManageClusterAdmin,
	"CapabilityOrganizationRegistrationsPerHour":   CapabilityOrganizationRegistrationsPerHour,
	"CapabilityOrganizationPinClusterToShard":      CapabilityOrganizationPinClusterToShard,
	"CapabilityHibernateCluster":                   CapabilityHibernateCluster,
	"CapabilitySubscribedOcp":                      CapabilitySubscribedOcp,
	"CapabilitySubscribedOcpMarketplace":           CapabilitySubscribedOcpMarketplace,
	"CapabilitySubscribedOsdMarketplace":           CapabilitySubscribedOsdMarketplace,
	"CapabilityEnableTermsEnforcement":             CapabilityEnableTermsEnforcement,
	"CapabilityBareMetalInstallerAdmin":            CapabilityBareMetalInstallerAdmin,
	"CapabilityReleaseOcpClusters":                 CapabilityReleaseOcpClusters,
	"CapabilityAutoscaleClustersDeprecated":        CapabilityAutoscaleClustersDeprecated,
	"CapabilityAutoscaleClusters":                  CapabilityAutoscaleClusters,
	"CapabilityOrganizationInstallConfigOverride":  CapabilityOrganizationInstallConfigOverride,
	"CapabilityOrganizationInstallConfigDefault":   CapabilityOrganizationInstallConfigDefault,
	"CapabilityOrganizationOverrideOsdTrialLength": CapabilityOrganizationOverrideOsdTrialLength,
	"CapabilityOrganizationCreateClusterProxy":     CapabilityOrganizationCreateClusterProxy,
	"CapabilityAllowGCPNonCCSPrivateClusters":      CapabilityAllowGCPNonCCSPrivateClusters,
	"CapabilityAddOnVersionSelect":                 CapabilityAddOnVersionSelect,
	"CapabilityOrganizationFipsCluster":            CapabilityOrganizationFipsCluster,
	"CapabilityOrganizationOvnCluster":             CapabilityOrganizationOvnCluster,
	"CapabilityOrganizationHyperShift":             CapabilityOrganizationHyperShift,
	"CapabilityBypassMaxExpiration":                CapabilityBypassMaxExpiration,
	"CapabilityUseRosaPaidAMI":                     CapabilityUseRosaPaidAMI,
}

func PresentCapabilities(capabilities []*v1.Capability) CapabilityList {
	var capabilitiesList []Capability
	for _, capability := range capabilities {
		cap := Capability{
			Name:      capability.Name(),
			Value:     capability.Value(),
			Inherited: capability.Inherited(),
		}
		capabilitiesList = append(capabilitiesList, cap)
	}
	return capabilitiesList
}

func GetAvailableCapability(capability string, resourceType string) (string, error) {
	val, ok := availableCapabilities[capability]
	if !ok {
		return "", fmt.Errorf("capability not found")
	}
	if strings.Split(val, ".")[1] != resourceType {
		return "", fmt.Errorf("capability not available for given resource type")
	}
	return val, nil
}
