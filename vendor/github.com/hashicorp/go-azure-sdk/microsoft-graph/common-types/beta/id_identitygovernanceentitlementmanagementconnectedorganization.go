package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementConnectedOrganizationId{}

// IdentityGovernanceEntitlementManagementConnectedOrganizationId is a struct representing the Resource ID for a Identity Governance Entitlement Management Connected Organization
type IdentityGovernanceEntitlementManagementConnectedOrganizationId struct {
	ConnectedOrganizationId string
}

// NewIdentityGovernanceEntitlementManagementConnectedOrganizationID returns a new IdentityGovernanceEntitlementManagementConnectedOrganizationId struct
func NewIdentityGovernanceEntitlementManagementConnectedOrganizationID(connectedOrganizationId string) IdentityGovernanceEntitlementManagementConnectedOrganizationId {
	return IdentityGovernanceEntitlementManagementConnectedOrganizationId{
		ConnectedOrganizationId: connectedOrganizationId,
	}
}

// ParseIdentityGovernanceEntitlementManagementConnectedOrganizationID parses 'input' into a IdentityGovernanceEntitlementManagementConnectedOrganizationId
func ParseIdentityGovernanceEntitlementManagementConnectedOrganizationID(input string) (*IdentityGovernanceEntitlementManagementConnectedOrganizationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementConnectedOrganizationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementConnectedOrganizationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementConnectedOrganizationId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementConnectedOrganizationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementConnectedOrganizationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementConnectedOrganizationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementConnectedOrganizationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ConnectedOrganizationId, ok = input.Parsed["connectedOrganizationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "connectedOrganizationId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementConnectedOrganizationID checks that 'input' can be parsed as a Identity Governance Entitlement Management Connected Organization ID
func ValidateIdentityGovernanceEntitlementManagementConnectedOrganizationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementConnectedOrganizationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Connected Organization ID
func (id IdentityGovernanceEntitlementManagementConnectedOrganizationId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/connectedOrganizations/%s"
	return fmt.Sprintf(fmtString, id.ConnectedOrganizationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Connected Organization ID
func (id IdentityGovernanceEntitlementManagementConnectedOrganizationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("connectedOrganizations", "connectedOrganizations", "connectedOrganizations"),
		resourceids.UserSpecifiedSegment("connectedOrganizationId", "connectedOrganizationId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Connected Organization ID
func (id IdentityGovernanceEntitlementManagementConnectedOrganizationId) String() string {
	components := []string{
		fmt.Sprintf("Connected Organization: %q", id.ConnectedOrganizationId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Connected Organization (%s)", strings.Join(components, "\n"))
}
