package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId{}

// IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId is a struct representing the Resource ID for a Identity Governance Entitlement Management Connected Organization Id Internal Sponsor
type IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId struct {
	ConnectedOrganizationId string
	DirectoryObjectId       string
}

// NewIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID returns a new IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId struct
func NewIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID(connectedOrganizationId string, directoryObjectId string) IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId {
	return IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId{
		ConnectedOrganizationId: connectedOrganizationId,
		DirectoryObjectId:       directoryObjectId,
	}
}

// ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID parses 'input' into a IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId
func ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID(input string) (*IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ConnectedOrganizationId, ok = input.Parsed["connectedOrganizationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "connectedOrganizationId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID checks that 'input' can be parsed as a Identity Governance Entitlement Management Connected Organization Id Internal Sponsor ID
func ValidateIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Connected Organization Id Internal Sponsor ID
func (id IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/connectedOrganizations/%s/internalSponsors/%s"
	return fmt.Sprintf(fmtString, id.ConnectedOrganizationId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Connected Organization Id Internal Sponsor ID
func (id IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("connectedOrganizations", "connectedOrganizations", "connectedOrganizations"),
		resourceids.UserSpecifiedSegment("connectedOrganizationId", "connectedOrganizationId"),
		resourceids.StaticSegment("internalSponsors", "internalSponsors", "internalSponsors"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Connected Organization Id Internal Sponsor ID
func (id IdentityGovernanceEntitlementManagementConnectedOrganizationIdInternalSponsorId) String() string {
	components := []string{
		fmt.Sprintf("Connected Organization: %q", id.ConnectedOrganizationId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Connected Organization Id Internal Sponsor (%s)", strings.Join(components, "\n"))
}
