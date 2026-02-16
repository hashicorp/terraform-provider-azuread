package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{}

// IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy Id Custom Extension Stage Setting
type IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId struct {
	AccessPackageId                 string
	AccessPackageAssignmentPolicyId string
	CustomExtensionStageSettingId   string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID returns a new IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID(accessPackageId string, accessPackageAssignmentPolicyId string, customExtensionStageSettingId string) IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId {
	return IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{
		AccessPackageId:                 accessPackageId,
		AccessPackageAssignmentPolicyId: accessPackageAssignmentPolicyId,
		CustomExtensionStageSettingId:   customExtensionStageSettingId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	if id.AccessPackageAssignmentPolicyId, ok = input.Parsed["accessPackageAssignmentPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentPolicyId", input)
	}

	if id.CustomExtensionStageSettingId, ok = input.Parsed["customExtensionStageSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customExtensionStageSettingId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy Id Custom Extension Stage Setting ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy Id Custom Extension Stage Setting ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackages/%s/accessPackageAssignmentPolicies/%s/customExtensionStageSettings/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageId, id.AccessPackageAssignmentPolicyId, id.CustomExtensionStageSettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy Id Custom Extension Stage Setting ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
		resourceids.StaticSegment("accessPackageAssignmentPolicies", "accessPackageAssignmentPolicies", "accessPackageAssignmentPolicies"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentPolicyId", "accessPackageAssignmentPolicyId"),
		resourceids.StaticSegment("customExtensionStageSettings", "customExtensionStageSettings", "customExtensionStageSettings"),
		resourceids.UserSpecifiedSegment("customExtensionStageSettingId", "customExtensionStageSettingId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy Id Custom Extension Stage Setting ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionStageSettingId) String() string {
	components := []string{
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
		fmt.Sprintf("Access Package Assignment Policy: %q", id.AccessPackageAssignmentPolicyId),
		fmt.Sprintf("Custom Extension Stage Setting: %q", id.CustomExtensionStageSettingId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy Id Custom Extension Stage Setting (%s)", strings.Join(components, "\n"))
}
