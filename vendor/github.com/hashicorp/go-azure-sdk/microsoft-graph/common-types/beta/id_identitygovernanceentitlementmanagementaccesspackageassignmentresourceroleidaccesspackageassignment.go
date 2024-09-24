package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId{}

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Assignment Resource Role Id Access Package Assignment
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId struct {
	AccessPackageAssignmentResourceRoleId string
	AccessPackageAssignmentId             string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentID returns a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentID(accessPackageAssignmentResourceRoleId string, accessPackageAssignmentId string) IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId {
	return IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId{
		AccessPackageAssignmentResourceRoleId: accessPackageAssignmentResourceRoleId,
		AccessPackageAssignmentId:             accessPackageAssignmentId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentResourceRoleId, ok = input.Parsed["accessPackageAssignmentResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentResourceRoleId", input)
	}

	if id.AccessPackageAssignmentId, ok = input.Parsed["accessPackageAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Assignment Resource Role Id Access Package Assignment ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Assignment Resource Role Id Access Package Assignment ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/%s/accessPackageAssignments/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentResourceRoleId, id.AccessPackageAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Assignment Resource Role Id Access Package Assignment ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageAssignmentResourceRoles", "accessPackageAssignmentResourceRoles", "accessPackageAssignmentResourceRoles"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentResourceRoleId", "accessPackageAssignmentResourceRoleId"),
		resourceids.StaticSegment("accessPackageAssignments", "accessPackageAssignments", "accessPackageAssignments"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentId", "accessPackageAssignmentId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Assignment Resource Role Id Access Package Assignment ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment Resource Role: %q", id.AccessPackageAssignmentResourceRoleId),
		fmt.Sprintf("Access Package Assignment: %q", id.AccessPackageAssignmentId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Assignment Resource Role Id Access Package Assignment (%s)", strings.Join(components, "\n"))
}
