package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId{}

// IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Id Assignment Policy Id Custom Extension Stage Setting
type IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId struct {
	AccessPackageId                 string
	AccessPackageAssignmentPolicyId string
	CustomExtensionStageSettingId   string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingID returns a new IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingID(accessPackageId string, accessPackageAssignmentPolicyId string, customExtensionStageSettingId string) IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId {
	return IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId{
		AccessPackageId:                 accessPackageId,
		AccessPackageAssignmentPolicyId: accessPackageAssignmentPolicyId,
		CustomExtensionStageSettingId:   customExtensionStageSettingId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Id Assignment Policy Id Custom Extension Stage Setting ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Id Assignment Policy Id Custom Extension Stage Setting ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackages/%s/assignmentPolicies/%s/customExtensionStageSettings/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageId, id.AccessPackageAssignmentPolicyId, id.CustomExtensionStageSettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Id Assignment Policy Id Custom Extension Stage Setting ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
		resourceids.StaticSegment("assignmentPolicies", "assignmentPolicies", "assignmentPolicies"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentPolicyId", "accessPackageAssignmentPolicyId"),
		resourceids.StaticSegment("customExtensionStageSettings", "customExtensionStageSettings", "customExtensionStageSettings"),
		resourceids.UserSpecifiedSegment("customExtensionStageSettingId", "customExtensionStageSettingId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Id Assignment Policy Id Custom Extension Stage Setting ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdCustomExtensionStageSettingId) String() string {
	components := []string{
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
		fmt.Sprintf("Access Package Assignment Policy: %q", id.AccessPackageAssignmentPolicyId),
		fmt.Sprintf("Custom Extension Stage Setting: %q", id.CustomExtensionStageSettingId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Id Assignment Policy Id Custom Extension Stage Setting (%s)", strings.Join(components, "\n"))
}
