package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId{}

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Assignment Resource Role
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId struct {
	AccessPackageAssignmentResourceRoleId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleID returns a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleID(accessPackageAssignmentResourceRoleId string) IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId {
	return IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId{
		AccessPackageAssignmentResourceRoleId: accessPackageAssignmentResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentResourceRoleId, ok = input.Parsed["accessPackageAssignmentResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentResourceRoleId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Assignment Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Assignment Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Assignment Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageAssignmentResourceRoles", "accessPackageAssignmentResourceRoles", "accessPackageAssignmentResourceRoles"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentResourceRoleId", "accessPackageAssignmentResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Assignment Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment Resource Role: %q", id.AccessPackageAssignmentResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Assignment Resource Role (%s)", strings.Join(components, "\n"))
}
