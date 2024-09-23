package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryOutboundSharedUserProfileId{}

// DirectoryOutboundSharedUserProfileId is a struct representing the Resource ID for a Directory Outbound Shared User Profile
type DirectoryOutboundSharedUserProfileId struct {
	OutboundSharedUserProfileUserId string
}

// NewDirectoryOutboundSharedUserProfileID returns a new DirectoryOutboundSharedUserProfileId struct
func NewDirectoryOutboundSharedUserProfileID(outboundSharedUserProfileUserId string) DirectoryOutboundSharedUserProfileId {
	return DirectoryOutboundSharedUserProfileId{
		OutboundSharedUserProfileUserId: outboundSharedUserProfileUserId,
	}
}

// ParseDirectoryOutboundSharedUserProfileID parses 'input' into a DirectoryOutboundSharedUserProfileId
func ParseDirectoryOutboundSharedUserProfileID(input string) (*DirectoryOutboundSharedUserProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryOutboundSharedUserProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryOutboundSharedUserProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryOutboundSharedUserProfileIDInsensitively parses 'input' case-insensitively into a DirectoryOutboundSharedUserProfileId
// note: this method should only be used for API response data and not user input
func ParseDirectoryOutboundSharedUserProfileIDInsensitively(input string) (*DirectoryOutboundSharedUserProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryOutboundSharedUserProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryOutboundSharedUserProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryOutboundSharedUserProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OutboundSharedUserProfileUserId, ok = input.Parsed["outboundSharedUserProfileUserId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outboundSharedUserProfileUserId", input)
	}

	return nil
}

// ValidateDirectoryOutboundSharedUserProfileID checks that 'input' can be parsed as a Directory Outbound Shared User Profile ID
func ValidateDirectoryOutboundSharedUserProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryOutboundSharedUserProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Outbound Shared User Profile ID
func (id DirectoryOutboundSharedUserProfileId) ID() string {
	fmtString := "/directory/outboundSharedUserProfiles/%s"
	return fmt.Sprintf(fmtString, id.OutboundSharedUserProfileUserId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Outbound Shared User Profile ID
func (id DirectoryOutboundSharedUserProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("outboundSharedUserProfiles", "outboundSharedUserProfiles", "outboundSharedUserProfiles"),
		resourceids.UserSpecifiedSegment("outboundSharedUserProfileUserId", "outboundSharedUserProfileUserId"),
	}
}

// String returns a human-readable description of this Directory Outbound Shared User Profile ID
func (id DirectoryOutboundSharedUserProfileId) String() string {
	components := []string{
		fmt.Sprintf("Outbound Shared User Profile User: %q", id.OutboundSharedUserProfileUserId),
	}
	return fmt.Sprintf("Directory Outbound Shared User Profile (%s)", strings.Join(components, "\n"))
}
