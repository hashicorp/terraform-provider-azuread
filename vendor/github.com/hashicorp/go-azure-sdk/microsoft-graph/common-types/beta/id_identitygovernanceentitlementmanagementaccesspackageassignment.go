package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentId{}

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Assignment
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentId struct {
	AccessPackageAssignmentId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentID returns a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentID(accessPackageAssignmentId string) IdentityGovernanceEntitlementManagementAccessPackageAssignmentId {
	return IdentityGovernanceEntitlementManagementAccessPackageAssignmentId{
		AccessPackageAssignmentId: accessPackageAssignmentId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentId
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentId, ok = input.Parsed["accessPackageAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Assignment ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Assignment ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageAssignments/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Assignment ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageAssignments", "accessPackageAssignments", "accessPackageAssignments"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentId", "accessPackageAssignmentId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Assignment ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment: %q", id.AccessPackageAssignmentId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Assignment (%s)", strings.Join(components, "\n"))
}
