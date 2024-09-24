package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId{}

// IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId is a struct representing the Resource ID for a Identity Governance Entitlement Management Assignment Policy Id Custom Extension Stage Setting
type IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId struct {
	AccessPackageAssignmentPolicyId string
	CustomExtensionStageSettingId   string
}

// NewIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingID returns a new IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId struct
func NewIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingID(accessPackageAssignmentPolicyId string, customExtensionStageSettingId string) IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId {
	return IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId{
		AccessPackageAssignmentPolicyId: accessPackageAssignmentPolicyId,
		CustomExtensionStageSettingId:   customExtensionStageSettingId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingID parses 'input' into a IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId
func ParseIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingID(input string) (*IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentPolicyId, ok = input.Parsed["accessPackageAssignmentPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentPolicyId", input)
	}

	if id.CustomExtensionStageSettingId, ok = input.Parsed["customExtensionStageSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customExtensionStageSettingId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingID checks that 'input' can be parsed as a Identity Governance Entitlement Management Assignment Policy Id Custom Extension Stage Setting ID
func ValidateIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Assignment Policy Id Custom Extension Stage Setting ID
func (id IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/assignmentPolicies/%s/customExtensionStageSettings/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentPolicyId, id.CustomExtensionStageSettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Assignment Policy Id Custom Extension Stage Setting ID
func (id IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("assignmentPolicies", "assignmentPolicies", "assignmentPolicies"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentPolicyId", "accessPackageAssignmentPolicyId"),
		resourceids.StaticSegment("customExtensionStageSettings", "customExtensionStageSettings", "customExtensionStageSettings"),
		resourceids.UserSpecifiedSegment("customExtensionStageSettingId", "customExtensionStageSettingId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Assignment Policy Id Custom Extension Stage Setting ID
func (id IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment Policy: %q", id.AccessPackageAssignmentPolicyId),
		fmt.Sprintf("Custom Extension Stage Setting: %q", id.CustomExtensionStageSettingId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Assignment Policy Id Custom Extension Stage Setting (%s)", strings.Join(components, "\n"))
}
