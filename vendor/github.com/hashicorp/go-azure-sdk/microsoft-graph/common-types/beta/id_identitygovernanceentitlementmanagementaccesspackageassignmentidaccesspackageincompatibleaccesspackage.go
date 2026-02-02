package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId{}

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Assignment Id Access Package Incompatible Access Package
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId struct {
	AccessPackageAssignmentId string
	AccessPackageId           string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageID returns a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageID(accessPackageAssignmentId string, accessPackageId string) IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId {
	return IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId{
		AccessPackageAssignmentId: accessPackageAssignmentId,
		AccessPackageId:           accessPackageId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentId, ok = input.Parsed["accessPackageAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentId", input)
	}

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Assignment Id Access Package Incompatible Access Package ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Assignment Id Access Package Incompatible Access Package ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageAssignments/%s/accessPackage/incompatibleAccessPackages/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentId, id.AccessPackageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Assignment Id Access Package Incompatible Access Package ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageAssignments", "accessPackageAssignments", "accessPackageAssignments"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentId", "accessPackageAssignmentId"),
		resourceids.StaticSegment("accessPackage", "accessPackage", "accessPackage"),
		resourceids.StaticSegment("incompatibleAccessPackages", "incompatibleAccessPackages", "incompatibleAccessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Assignment Id Access Package Incompatible Access Package ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageIncompatibleAccessPackageId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment: %q", id.AccessPackageAssignmentId),
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Assignment Id Access Package Incompatible Access Package (%s)", strings.Join(components, "\n"))
}
