package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId{}

// IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Id Incompatible Access Package
type IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId struct {
	AccessPackageId  string
	AccessPackageId1 string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageID returns a new IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageID(accessPackageId string, accessPackageId1 string) IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId {
	return IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId{
		AccessPackageId:  accessPackageId,
		AccessPackageId1: accessPackageId1,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	if id.AccessPackageId1, ok = input.Parsed["accessPackageId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId1", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Id Incompatible Access Package ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Id Incompatible Access Package ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackages/%s/incompatibleAccessPackages/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageId, id.AccessPackageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Id Incompatible Access Package ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
		resourceids.StaticSegment("incompatibleAccessPackages", "incompatibleAccessPackages", "incompatibleAccessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId1", "accessPackageId1"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Id Incompatible Access Package ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId) String() string {
	components := []string{
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
		fmt.Sprintf("Access Package Id 1: %q", id.AccessPackageId1),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Id Incompatible Access Package (%s)", strings.Join(components, "\n"))
}
