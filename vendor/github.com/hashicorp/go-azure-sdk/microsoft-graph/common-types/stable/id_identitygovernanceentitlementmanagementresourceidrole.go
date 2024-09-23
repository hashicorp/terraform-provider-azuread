package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceIdRoleId{}

// IdentityGovernanceEntitlementManagementResourceIdRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Id Role
type IdentityGovernanceEntitlementManagementResourceIdRoleId struct {
	AccessPackageResourceId     string
	AccessPackageResourceRoleId string
}

// NewIdentityGovernanceEntitlementManagementResourceIdRoleID returns a new IdentityGovernanceEntitlementManagementResourceIdRoleId struct
func NewIdentityGovernanceEntitlementManagementResourceIdRoleID(accessPackageResourceId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementResourceIdRoleId {
	return IdentityGovernanceEntitlementManagementResourceIdRoleId{
		AccessPackageResourceId:     accessPackageResourceId,
		AccessPackageResourceRoleId: accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceIdRoleID parses 'input' into a IdentityGovernanceEntitlementManagementResourceIdRoleId
func ParseIdentityGovernanceEntitlementManagementResourceIdRoleID(input string) (*IdentityGovernanceEntitlementManagementResourceIdRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceIdRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceIdRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceIdRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceIdRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceIdRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceIdRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceIdRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceIdRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceIdRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceIdRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Id Role ID
func ValidateIdentityGovernanceEntitlementManagementResourceIdRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceIdRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Id Role ID
func (id IdentityGovernanceEntitlementManagementResourceIdRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resources/%s/roles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Id Role ID
func (id IdentityGovernanceEntitlementManagementResourceIdRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Id Role ID
func (id IdentityGovernanceEntitlementManagementResourceIdRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Id Role (%s)", strings.Join(components, "\n"))
}
