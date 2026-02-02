package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteId{}

// GroupIdSiteId is a struct representing the Resource ID for a Group Id Site
type GroupIdSiteId struct {
	GroupId string
	SiteId  string
}

// NewGroupIdSiteID returns a new GroupIdSiteId struct
func NewGroupIdSiteID(groupId string, siteId string) GroupIdSiteId {
	return GroupIdSiteId{
		GroupId: groupId,
		SiteId:  siteId,
	}
}

// ParseGroupIdSiteID parses 'input' into a GroupIdSiteId
func ParseGroupIdSiteID(input string) (*GroupIdSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIDInsensitively parses 'input' case-insensitively into a GroupIdSiteId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIDInsensitively(input string) (*GroupIdSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	return nil
}

// ValidateGroupIdSiteID checks that 'input' can be parsed as a Group Id Site ID
func ValidateGroupIdSiteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site ID
func (id GroupIdSiteId) ID() string {
	fmtString := "/groups/%s/sites/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site ID
func (id GroupIdSiteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
	}
}

// String returns a human-readable description of this Group Id Site ID
func (id GroupIdSiteId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
	}
	return fmt.Sprintf("Group Id Site (%s)", strings.Join(components, "\n"))
}
