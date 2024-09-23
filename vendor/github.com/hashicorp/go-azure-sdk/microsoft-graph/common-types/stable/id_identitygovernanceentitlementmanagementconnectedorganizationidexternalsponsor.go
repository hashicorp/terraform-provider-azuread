package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId{}

// IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId is a struct representing the Resource ID for a Identity Governance Entitlement Management Connected Organization Id External Sponsor
type IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId struct {
	ConnectedOrganizationId string
	DirectoryObjectId       string
}

// NewIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorID returns a new IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId struct
func NewIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorID(connectedOrganizationId string, directoryObjectId string) IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId {
	return IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId{
		ConnectedOrganizationId: connectedOrganizationId,
		DirectoryObjectId:       directoryObjectId,
	}
}

// ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorID parses 'input' into a IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId
func ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorID(input string) (*IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ConnectedOrganizationId, ok = input.Parsed["connectedOrganizationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "connectedOrganizationId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorID checks that 'input' can be parsed as a Identity Governance Entitlement Management Connected Organization Id External Sponsor ID
func ValidateIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Connected Organization Id External Sponsor ID
func (id IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/connectedOrganizations/%s/externalSponsors/%s"
	return fmt.Sprintf(fmtString, id.ConnectedOrganizationId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Connected Organization Id External Sponsor ID
func (id IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("connectedOrganizations", "connectedOrganizations", "connectedOrganizations"),
		resourceids.UserSpecifiedSegment("connectedOrganizationId", "connectedOrganizationId"),
		resourceids.StaticSegment("externalSponsors", "externalSponsors", "externalSponsors"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Connected Organization Id External Sponsor ID
func (id IdentityGovernanceEntitlementManagementConnectedOrganizationIdExternalSponsorId) String() string {
	components := []string{
		fmt.Sprintf("Connected Organization: %q", id.ConnectedOrganizationId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Connected Organization Id External Sponsor (%s)", strings.Join(components, "\n"))
}
