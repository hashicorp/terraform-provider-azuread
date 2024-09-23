package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{}

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Assignment Policy Id Custom Extension Stage Setting
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId struct {
	AccessPackageAssignmentPolicyId string
	CustomExtensionStageSettingId   string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID returns a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID(accessPackageAssignmentPolicyId string, customExtensionStageSettingId string) IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId {
	return IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{
		AccessPackageAssignmentPolicyId: accessPackageAssignmentPolicyId,
		CustomExtensionStageSettingId:   customExtensionStageSettingId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentPolicyId, ok = input.Parsed["accessPackageAssignmentPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentPolicyId", input)
	}

	if id.CustomExtensionStageSettingId, ok = input.Parsed["customExtensionStageSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customExtensionStageSettingId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Assignment Policy Id Custom Extension Stage Setting ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Assignment Policy Id Custom Extension Stage Setting ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageAssignmentPolicies/%s/customExtensionStageSettings/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentPolicyId, id.CustomExtensionStageSettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Assignment Policy Id Custom Extension Stage Setting ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageAssignmentPolicies", "accessPackageAssignmentPolicies", "accessPackageAssignmentPolicies"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentPolicyId", "accessPackageAssignmentPolicyId"),
		resourceids.StaticSegment("customExtensionStageSettings", "customExtensionStageSettings", "customExtensionStageSettings"),
		resourceids.UserSpecifiedSegment("customExtensionStageSettingId", "customExtensionStageSettingId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Assignment Policy Id Custom Extension Stage Setting ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment Policy: %q", id.AccessPackageAssignmentPolicyId),
		fmt.Sprintf("Custom Extension Stage Setting: %q", id.CustomExtensionStageSettingId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Assignment Policy Id Custom Extension Stage Setting (%s)", strings.Join(components, "\n"))
}
