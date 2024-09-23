package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageId{}

// IdentityGovernanceEntitlementManagementAccessPackageId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package
type IdentityGovernanceEntitlementManagementAccessPackageId struct {
	AccessPackageId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageID returns a new IdentityGovernanceEntitlementManagementAccessPackageId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageID(accessPackageId string) IdentityGovernanceEntitlementManagementAccessPackageId {
	return IdentityGovernanceEntitlementManagementAccessPackageId{
		AccessPackageId: accessPackageId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageId
func ParseIdentityGovernanceEntitlementManagementAccessPackageID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package ID
func (id IdentityGovernanceEntitlementManagementAccessPackageId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackages/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package ID
func (id IdentityGovernanceEntitlementManagementAccessPackageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package ID
func (id IdentityGovernanceEntitlementManagementAccessPackageId) String() string {
	components := []string{
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package (%s)", strings.Join(components, "\n"))
}
