package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId{}

// IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Id Access Packages Incompatible With
type IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId struct {
	AccessPackageId  string
	AccessPackageId1 string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithID returns a new IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithID(accessPackageId string, accessPackageId1 string) IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId {
	return IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId{
		AccessPackageId:  accessPackageId,
		AccessPackageId1: accessPackageId1,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	if id.AccessPackageId1, ok = input.Parsed["accessPackageId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId1", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Id Access Packages Incompatible With ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Id Access Packages Incompatible With ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackages/%s/accessPackagesIncompatibleWith/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageId, id.AccessPackageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Id Access Packages Incompatible With ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
		resourceids.StaticSegment("accessPackagesIncompatibleWith", "accessPackagesIncompatibleWith", "accessPackagesIncompatibleWith"),
		resourceids.UserSpecifiedSegment("accessPackageId1", "accessPackageId1"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Id Access Packages Incompatible With ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId) String() string {
	components := []string{
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
		fmt.Sprintf("Access Package Id 1: %q", id.AccessPackageId1),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Id Access Packages Incompatible With (%s)", strings.Join(components, "\n"))
}
