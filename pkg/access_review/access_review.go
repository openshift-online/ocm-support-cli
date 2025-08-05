package access_review

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1Auth "github.com/openshift-online/ocm-sdk-go/authorizations/v1"
)

type AccessReview struct {
	Allowed        bool     `json:"allowed"`
	Username       string   `json:"username"`
	Action         string   `json:"action"`
	ResourceType   string   `json:"resource_type"`
	OrganizationID string   `json:"organization_id,omitempty"`
	SubscriptionID string   `json:"subscription_id,omitempty"`
	ClusterID      string   `json:"cluster_id,omitempty"`
	SuggestedRoles []string `json:"suggested_roles,omitempty"`
}

// ComprehensiveRoleMapping represents the complete role mapping structure
type ComprehensiveRoleMapping struct {
	Meta            RoleMappingMeta     `json:"meta"`
	Roles           map[string]RoleData `json:"roles"`
	ResourceToRoles map[string][]string `json:"resource_to_roles"`
}

type RoleMappingMeta struct {
	TotalRoles         int      `json:"total_roles"`
	TotalResourceTypes int      `json:"total_resource_types"`
	TotalActionTypes   int      `json:"total_action_types"`
	ResourceTypes      []string `json:"resource_types"`
	ActionTypes        []string `json:"action_types"`
}

type RoleData struct {
	RoleID      string                     `json:"role_id"`
	Permissions map[string]map[string]bool `json:"permissions"`
}

func PostAccessReview(username, action, resourceType, organizationID, subscriptionID, clusterID string, conn *sdk.Connection) (*v1Auth.AccessReviewResponse, error) {
	// Build the access review request
	requestBuilder := v1Auth.NewAccessReviewRequest().
		AccountUsername(username).
		Action(action).
		ResourceType(resourceType)

	// Add optional context parameters if provided
	if organizationID != "" {
		requestBuilder = requestBuilder.OrganizationID(organizationID)
	}
	if subscriptionID != "" {
		requestBuilder = requestBuilder.SubscriptionID(subscriptionID)
	}
	if clusterID != "" {
		requestBuilder = requestBuilder.ClusterID(clusterID)
	}

	accessReviewRequest, err := requestBuilder.Build()
	if err != nil {
		return nil, fmt.Errorf("can't build access review request: %w", err)
	}

	accessResponse, err := conn.Authorizations().V1().AccessReview().Post().
		Request(accessReviewRequest).
		Send()
	if err != nil {
		return nil, fmt.Errorf("can't retrieve access review: %w", err)
	}

	responseBody := accessResponse.Response()

	return responseBody, nil
}

// SuggestRoles returns a list of roles that might grant access to the specified action and resource
// Based on comprehensive analysis of role definitions from:
// /Users/tithakka/uhc-project-workspace/src/gitlab.cee.redhat.com/service/uhc-account-manager/pkg/api/roles/
func SuggestRoles(action, resourceType string) []string {
	var suggestions []string

	// Build comprehensive role-to-resource mappings
	roleResourceMap := buildComprehensiveRoleResourceMap()

	// Normalize resource name - API uses names like "ReservedResource" but our mapping uses "ReservedResourceResource"
	normalizedResourceType := normalizeResourceName(resourceType)

	// Convert API action to internal action name: delete -> DeleteAction, create -> CreateAction, etc.
	internalAction := normalizeActionName(action)

	// Check each role to see if it has the requested action on the requested resource
	for roleName, resources := range roleResourceMap {
		// Try both the original name and the normalized name
		for _, resourceToCheck := range []string{resourceType, normalizedResourceType} {
			if actions, hasResource := resources[resourceToCheck]; hasResource {
				// Check if this role has the requested action in any format
				for _, roleAction := range actions {
					// Check for exact internal action match (DeleteAction)
					// Check for StarAction (grants all permissions)
					if roleAction == internalAction || roleAction == "StarAction" {
						suggestions = append(suggestions, roleName)
						break
					}
				}
				break // Found a match, no need to check the other name variant
			}
		}
	}

	// Remove duplicates
	seen := make(map[string]bool)
	uniqueSuggestions := []string{}
	for _, role := range suggestions {
		if !seen[role] {
			seen[role] = true
			uniqueSuggestions = append(uniqueSuggestions, role)
		}
	}

	// No fallback suggestions - only suggest roles that explicitly have the permission

	return uniqueSuggestions
}

func PresentAccessReview(accessReviewResponse *v1Auth.AccessReviewResponse, username, action, resourceType, organizationID, subscriptionID, clusterID string, suggestRoles bool) *AccessReview {
	var review *AccessReview = nil

	if accessReviewResponse != nil {
		review = &AccessReview{
			Allowed:      accessReviewResponse.Allowed(),
			Username:     username,
			Action:       action,
			ResourceType: resourceType,
		}

		// Only include context IDs if they were provided
		if organizationID != "" {
			review.OrganizationID = organizationID
		}
		if subscriptionID != "" {
			review.SubscriptionID = subscriptionID
		}
		if clusterID != "" {
			review.ClusterID = clusterID
		}

		// Add role suggestions if requested and access is denied
		if suggestRoles && !accessReviewResponse.Allowed() {
			review.SuggestedRoles = SuggestRoles(action, resourceType)
		}
	}

	return review
}

// normalizeResourceName converts API resource names to internal mapping names
// API: "ReservedResource" -> Internal: "ReservedResourceResource"
// API: "Cluster" -> Internal: "ClusterResource"
func normalizeResourceName(apiResourceName string) string {
	// Handle special cases where API name differs significantly from internal name
	specialCases := map[string]string{
		"Cluster":      "ClusterResource",
		"Account":      "AccountResource",
		"Subscription": "SubscriptionResource",
		"Organization": "OrganizationResource",
		"MachinePool":  "MachinePoolResource",
		"Idp":          "IdpResource",
	}

	if internalName, exists := specialCases[apiResourceName]; exists {
		return internalName
	}

	// For resources that already end with "Resource", add another "Resource" suffix
	// This handles the pattern: "ReservedResource" -> "ReservedResourceResource"
	if apiResourceName != "" && strings.HasSuffix(apiResourceName, "Resource") {
		return apiResourceName + "Resource"
	}

	// For other resources, just add "Resource" suffix
	if apiResourceName != "" {
		return apiResourceName + "Resource"
	}

	return apiResourceName
}

// normalizeActionName converts API action names to internal action names
// API: "delete" -> Internal: "DeleteAction"
// API: "create" -> Internal: "CreateAction"
func normalizeActionName(apiAction string) string {
	switch strings.ToLower(apiAction) {
	case "delete":
		return "DeleteAction"
	case "create":
		return "CreateAction"
	case "get":
		return "GetAction"
	case "list":
		return "ListAction"
	case "update":
		return "UpdateAction"
	case "impersonate":
		return "ImpersonateAction"
	default:
		// If already in internal format or unknown, return as-is
		return apiAction
	}
}

// loadComprehensiveRoleMapping loads the comprehensive role mapping from JSON
func loadComprehensiveRoleMapping() (*ComprehensiveRoleMapping, error) {
	// Get the current working directory and construct the path to the JSON file
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	jsonPath := filepath.Join(wd, "comprehensive_role_mapping.json")

	// Read the JSON file
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read comprehensive role mapping file %s: %w", jsonPath, err)
	}

	// Parse the JSON
	var mapping ComprehensiveRoleMapping
	if err := json.Unmarshal(data, &mapping); err != nil {
		return nil, fmt.Errorf("failed to parse comprehensive role mapping JSON: %w", err)
	}

	return &mapping, nil
}

func buildComprehensiveRoleResourceMap() map[string]map[string][]string {
	// Load the comprehensive mapping from JSON with corrections
	mapping, err := loadComprehensiveRoleMapping()
	if err != nil {
		// Fallback to basic mapping if JSON loading fails
		fmt.Printf("Warning: Failed to load comprehensive role mapping: %v\n", err)
		return buildFallbackRoleResourceMap()
	}

	// Convert from JSON format to the format expected by suggestion logic with corrections
	roleResourceMap := make(map[string]map[string][]string)

	for roleName, roleData := range mapping.Roles {
		resourceMap := make(map[string][]string)

		for resourceName, actionMap := range roleData.Permissions {
			var actions []string

			// Check each action type
			for actionName, hasPermission := range actionMap {
				if hasPermission {
					actions = append(actions, actionName)
				}
			}

			if len(actions) > 0 {
				resourceMap[resourceName] = actions
			}
		}

		if len(resourceMap) > 0 {
			roleResourceMap[roleName] = resourceMap
		}
	}

	// Apply corrections for known parsing errors
	roleResourceMap = applyRoleCorrections(roleResourceMap)

	return roleResourceMap
}

// applyRoleCorrections fixes known parsing errors in role definitions
func applyRoleCorrections(roleResourceMap map[string]map[string][]string) map[string]map[string][]string {
	// Fix UHCSupport - should only have Get/List except for StatusBoard/WebRCA
	if uhcSupport, exists := roleResourceMap["UHCSupport"]; exists {
		correctedUHC := make(map[string][]string)
		for resource, actions := range uhcSupport {
			// Only keep StarAction for StatusBoard and WebRCA, Get/List for others
			if resource == "StatusBoardResource" || resource == "WebRCAResource" {
				correctedUHC[resource] = []string{"StarAction"}
			} else {
				// Only keep Get/List actions
				var filteredActions []string
				for _, action := range actions {
					if action == "GetAction" || action == "ListAction" {
						filteredActions = append(filteredActions, action)
					}
				}
				if len(filteredActions) > 0 {
					correctedUHC[resource] = filteredActions
				}
			}
		}
		roleResourceMap["UHCSupport"] = correctedUHC
	}

	// Fix AMSQEAutomation - has special resource restrictions
	if amsQE, exists := roleResourceMap["AMSQEAutomation"]; exists {
		// Apply special restrictions from ams_qe_automation.go
		specialRestrictions := map[string][]string{
			"ClusterResource":                   {"GetAction", "ListAction", "UpdateAction", "ImpersonateAction"},
			"RegistryResource":                  {"GetAction", "ListAction"},
			"SubscriptionLabelInternalResource": {"GetAction", "ListAction"},
			"OrganizationLabelInternalResource": {"GetAction", "ListAction"},
			"AccountLabelInternalResource":      {"GetAction", "ListAction"},
			"LabelInternalResource":             {"GetAction", "ListAction"},
		}

		for resource, restrictedActions := range specialRestrictions {
			if _, hasResource := amsQE[resource]; hasResource {
				amsQE[resource] = restrictedActions
			}
		}
		roleResourceMap["AMSQEAutomation"] = amsQE
	}

	return roleResourceMap
}

func buildFallbackRoleResourceMap() map[string]map[string][]string {
	return map[string]map[string][]string{
		"AMSQEAutomation": {
			"FlavourResource":                   []string{"GetAction", "ListAction", "CreateAction", "DeleteAction", "UpdateAction"},
			"LabelInternalResource":             []string{"GetAction", "ListAction"},
			"RoleResource":                      []string{"GetAction", "ListAction"},
			"RegistryResource":                  []string{"GetAction", "ListAction"},
			"ClusterResource":                   []string{"GetAction", "ListAction", "UpdateAction", "ImpersonateAction"},
			"AccountLabelInternalResource":      []string{"GetAction", "ListAction"},
			"OrganizationLabelInternalResource": []string{"GetAction", "ListAction"},
			"SubscriptionLabelInternalResource": []string{"GetAction", "ListAction"},
		},
		"AuthenticatedUser": {
			"FlavourResource": []string{"GetAction", "ListAction"},
		},
		"CSMaintainer": {
			"FlavourResource": []string{"GetAction", "ListAction", "UpdateAction"},
		},
		"SuperAdmin": {
			"FlavourResource": []string{"CreateAction", "DeleteAction", "GetAction", "ListAction", "UpdateAction"},
		},
		"UHCSupport": {
			// UHCSupport has Get/List on ALL resources except skippable ones (security-sensitive)
			// Plus StarAction on StatusBoard and WebRCA resources
			"AccessReviewResource":           []string{"GetAction", "ListAction"},
			"AccessTokenResource":            []string{"GetAction", "ListAction"},
			"AccountResource":                []string{"GetAction", "ListAction"},
			"AccountGroupResource":           []string{"GetAction", "ListAction"},
			"AccountGroupAssignmentResource": []string{"GetAction", "ListAction"},
			"AccountPoolResource":            []string{"GetAction", "ListAction"},
			"ClusterResource":                []string{"GetAction", "ListAction"},
			"MachinePoolResource":            []string{"GetAction", "ListAction"},
			"IdpResource":                    []string{"GetAction", "ListAction"},
			"ClusterInternalResource":        []string{"GetAction", "ListAction"},
			"CapabilityReviewResource":       []string{"GetAction", "ListAction"},
			"CapabilityResource":             []string{"GetAction", "ListAction"},
			"FlavourResource":                []string{"GetAction", "ListAction"},
			"LabelInternalResource":          []string{"GetAction", "ListAction"},
			"LabelResource":                  []string{"GetAction", "ListAction"},
			"RegistryResource":               []string{"GetAction", "ListAction"},
			"RoleResource":                   []string{"GetAction", "ListAction"},
			"RoleBindingResource":            []string{"GetAction", "ListAction"},
			"OrganizationResource":           []string{"GetAction", "ListAction"},
			"SubscriptionResource":           []string{"GetAction", "ListAction"},
			"ReservedResourceResource":       []string{"GetAction", "ListAction"},
			"CloudResourceResource":          []string{"GetAction", "ListAction"},
			"AddOnResource":                  []string{"GetAction", "ListAction"},
			"ServiceLogResource":             []string{"GetAction", "ListAction"},
			"InternalServiceLogResource":     []string{"GetAction", "ListAction"},
			"StatusBoardResource":            []string{"StarAction"},
			"WebRCAResource":                 []string{"StarAction"},
			// Note: Skippable security-sensitive Backplane resources are excluded
		},
		"ClusterService": {
			"ReservedResourceResource": []string{"StarAction"},
		},
		"AccountDeleter": {
			"LabelInternalResource": []string{"GetAction", "ListAction", "DeleteAction"},
		},
		"HackdayLead": {
			"LabelInternalResource": []string{"StarAction"},
		},
		"BackplaneService": {
			"RoleResource": []string{"GetAction", "ListAction"},
		},
		"OCMResourcesService": {
			"RoleResource": []string{"GetAction", "ListAction"},
		},
	}
}
