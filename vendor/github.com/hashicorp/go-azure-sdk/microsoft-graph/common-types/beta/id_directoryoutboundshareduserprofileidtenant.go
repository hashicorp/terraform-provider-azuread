package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryOutboundSharedUserProfileIdTenantId{}

// DirectoryOutboundSharedUserProfileIdTenantId is a struct representing the Resource ID for a Directory Outbound Shared User Profile Id Tenant
type DirectoryOutboundSharedUserProfileIdTenantId struct {
	OutboundSharedUserProfileUserId string
	TenantReferenceTenantId         string
}

// NewDirectoryOutboundSharedUserProfileIdTenantID returns a new DirectoryOutboundSharedUserProfileIdTenantId struct
func NewDirectoryOutboundSharedUserProfileIdTenantID(outboundSharedUserProfileUserId string, tenantReferenceTenantId string) DirectoryOutboundSharedUserProfileIdTenantId {
	return DirectoryOutboundSharedUserProfileIdTenantId{
		OutboundSharedUserProfileUserId: outboundSharedUserProfileUserId,
		TenantReferenceTenantId:         tenantReferenceTenantId,
	}
}

// ParseDirectoryOutboundSharedUserProfileIdTenantID parses 'input' into a DirectoryOutboundSharedUserProfileIdTenantId
func ParseDirectoryOutboundSharedUserProfileIdTenantID(input string) (*DirectoryOutboundSharedUserProfileIdTenantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryOutboundSharedUserProfileIdTenantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryOutboundSharedUserProfileIdTenantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryOutboundSharedUserProfileIdTenantIDInsensitively parses 'input' case-insensitively into a DirectoryOutboundSharedUserProfileIdTenantId
// note: this method should only be used for API response data and not user input
func ParseDirectoryOutboundSharedUserProfileIdTenantIDInsensitively(input string) (*DirectoryOutboundSharedUserProfileIdTenantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryOutboundSharedUserProfileIdTenantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryOutboundSharedUserProfileIdTenantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryOutboundSharedUserProfileIdTenantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OutboundSharedUserProfileUserId, ok = input.Parsed["outboundSharedUserProfileUserId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outboundSharedUserProfileUserId", input)
	}

	if id.TenantReferenceTenantId, ok = input.Parsed["tenantReferenceTenantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "tenantReferenceTenantId", input)
	}

	return nil
}

// ValidateDirectoryOutboundSharedUserProfileIdTenantID checks that 'input' can be parsed as a Directory Outbound Shared User Profile Id Tenant ID
func ValidateDirectoryOutboundSharedUserProfileIdTenantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryOutboundSharedUserProfileIdTenantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Outbound Shared User Profile Id Tenant ID
func (id DirectoryOutboundSharedUserProfileIdTenantId) ID() string {
	fmtString := "/directory/outboundSharedUserProfiles/%s/tenants/%s"
	return fmt.Sprintf(fmtString, id.OutboundSharedUserProfileUserId, id.TenantReferenceTenantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Outbound Shared User Profile Id Tenant ID
func (id DirectoryOutboundSharedUserProfileIdTenantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("outboundSharedUserProfiles", "outboundSharedUserProfiles", "outboundSharedUserProfiles"),
		resourceids.UserSpecifiedSegment("outboundSharedUserProfileUserId", "outboundSharedUserProfileUserId"),
		resourceids.StaticSegment("tenants", "tenants", "tenants"),
		resourceids.UserSpecifiedSegment("tenantReferenceTenantId", "tenantReferenceTenantId"),
	}
}

// String returns a human-readable description of this Directory Outbound Shared User Profile Id Tenant ID
func (id DirectoryOutboundSharedUserProfileIdTenantId) String() string {
	components := []string{
		fmt.Sprintf("Outbound Shared User Profile User: %q", id.OutboundSharedUserProfileUserId),
		fmt.Sprintf("Tenant Reference Tenant: %q", id.TenantReferenceTenantId),
	}
	return fmt.Sprintf("Directory Outbound Shared User Profile Id Tenant (%s)", strings.Join(components, "\n"))
}
