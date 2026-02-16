package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeFollowedSiteId{}

// MeFollowedSiteId is a struct representing the Resource ID for a Me Followed Site
type MeFollowedSiteId struct {
	SiteId string
}

// NewMeFollowedSiteID returns a new MeFollowedSiteId struct
func NewMeFollowedSiteID(siteId string) MeFollowedSiteId {
	return MeFollowedSiteId{
		SiteId: siteId,
	}
}

// ParseMeFollowedSiteID parses 'input' into a MeFollowedSiteId
func ParseMeFollowedSiteID(input string) (*MeFollowedSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeFollowedSiteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeFollowedSiteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeFollowedSiteIDInsensitively parses 'input' case-insensitively into a MeFollowedSiteId
// note: this method should only be used for API response data and not user input
func ParseMeFollowedSiteIDInsensitively(input string) (*MeFollowedSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeFollowedSiteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeFollowedSiteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeFollowedSiteId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	return nil
}

// ValidateMeFollowedSiteID checks that 'input' can be parsed as a Me Followed Site ID
func ValidateMeFollowedSiteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeFollowedSiteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Followed Site ID
func (id MeFollowedSiteId) ID() string {
	fmtString := "/me/followedSites/%s"
	return fmt.Sprintf(fmtString, id.SiteId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Followed Site ID
func (id MeFollowedSiteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("followedSites", "followedSites", "followedSites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
	}
}

// String returns a human-readable description of this Me Followed Site ID
func (id MeFollowedSiteId) String() string {
	components := []string{
		fmt.Sprintf("Site: %q", id.SiteId),
	}
	return fmt.Sprintf("Me Followed Site (%s)", strings.Join(components, "\n"))
}
