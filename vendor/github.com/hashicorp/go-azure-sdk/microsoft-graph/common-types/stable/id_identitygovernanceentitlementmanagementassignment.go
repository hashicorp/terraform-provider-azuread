package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAssignmentId{}

// IdentityGovernanceEntitlementManagementAssignmentId is a struct representing the Resource ID for a Identity Governance Entitlement Management Assignment
type IdentityGovernanceEntitlementManagementAssignmentId struct {
	AccessPackageAssignmentId string
}

// NewIdentityGovernanceEntitlementManagementAssignmentID returns a new IdentityGovernanceEntitlementManagementAssignmentId struct
func NewIdentityGovernanceEntitlementManagementAssignmentID(accessPackageAssignmentId string) IdentityGovernanceEntitlementManagementAssignmentId {
	return IdentityGovernanceEntitlementManagementAssignmentId{
		AccessPackageAssignmentId: accessPackageAssignmentId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAssignmentID parses 'input' into a IdentityGovernanceEntitlementManagementAssignmentId
func ParseIdentityGovernanceEntitlementManagementAssignmentID(input string) (*IdentityGovernanceEntitlementManagementAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAssignmentIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAssignmentId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAssignmentIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentId, ok = input.Parsed["accessPackageAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAssignmentID checks that 'input' can be parsed as a Identity Governance Entitlement Management Assignment ID
func ValidateIdentityGovernanceEntitlementManagementAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Assignment ID
func (id IdentityGovernanceEntitlementManagementAssignmentId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/assignments/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Assignment ID
func (id IdentityGovernanceEntitlementManagementAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentId", "accessPackageAssignmentId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Assignment ID
func (id IdentityGovernanceEntitlementManagementAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment: %q", id.AccessPackageAssignmentId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Assignment (%s)", strings.Join(components, "\n"))
}
