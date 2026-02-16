package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId{}

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Assignment Id Access Package Assignment Policy Custom Extension Stage Setting
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId struct {
	AccessPackageAssignmentId     string
	CustomExtensionStageSettingId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingID returns a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingID(accessPackageAssignmentId string, customExtensionStageSettingId string) IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId {
	return IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId{
		AccessPackageAssignmentId:     accessPackageAssignmentId,
		CustomExtensionStageSettingId: customExtensionStageSettingId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentId, ok = input.Parsed["accessPackageAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentId", input)
	}

	if id.CustomExtensionStageSettingId, ok = input.Parsed["customExtensionStageSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customExtensionStageSettingId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Assignment Id Access Package Assignment Policy Custom Extension Stage Setting ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Assignment Id Access Package Assignment Policy Custom Extension Stage Setting ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageAssignments/%s/accessPackageAssignmentPolicy/customExtensionStageSettings/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentId, id.CustomExtensionStageSettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Assignment Id Access Package Assignment Policy Custom Extension Stage Setting ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageAssignments", "accessPackageAssignments", "accessPackageAssignments"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentId", "accessPackageAssignmentId"),
		resourceids.StaticSegment("accessPackageAssignmentPolicy", "accessPackageAssignmentPolicy", "accessPackageAssignmentPolicy"),
		resourceids.StaticSegment("customExtensionStageSettings", "customExtensionStageSettings", "customExtensionStageSettings"),
		resourceids.UserSpecifiedSegment("customExtensionStageSettingId", "customExtensionStageSettingId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Assignment Id Access Package Assignment Policy Custom Extension Stage Setting ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentPolicyCustomExtensionStageSettingId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment: %q", id.AccessPackageAssignmentId),
		fmt.Sprintf("Custom Extension Stage Setting: %q", id.CustomExtensionStageSettingId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Assignment Id Access Package Assignment Policy Custom Extension Stage Setting (%s)", strings.Join(components, "\n"))
}
