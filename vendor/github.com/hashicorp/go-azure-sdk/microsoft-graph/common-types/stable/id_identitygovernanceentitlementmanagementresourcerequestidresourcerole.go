package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId{}

// IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Request Id Resource Role
type IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId struct {
	AccessPackageResourceRequestId string
	AccessPackageResourceRoleId    string
}

// NewIdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleID returns a new IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleID(accessPackageResourceRequestId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId {
	return IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
		AccessPackageResourceRoleId:    accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRequestId, ok = input.Parsed["accessPackageResourceRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRequestId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Request Id Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Request Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRequests/%s/resource/roles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Request Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRequests", "resourceRequests", "resourceRequests"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRequestId", "accessPackageResourceRequestId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Request Id Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Request Id Resource Role (%s)", strings.Join(components, "\n"))
}
