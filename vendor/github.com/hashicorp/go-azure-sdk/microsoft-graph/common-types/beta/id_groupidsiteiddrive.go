package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdDriveId{}

// GroupIdSiteIdDriveId is a struct representing the Resource ID for a Group Id Site Id Drive
type GroupIdSiteIdDriveId struct {
	GroupId string
	SiteId  string
	DriveId string
}

// NewGroupIdSiteIdDriveID returns a new GroupIdSiteIdDriveId struct
func NewGroupIdSiteIdDriveID(groupId string, siteId string, driveId string) GroupIdSiteIdDriveId {
	return GroupIdSiteIdDriveId{
		GroupId: groupId,
		SiteId:  siteId,
		DriveId: driveId,
	}
}

// ParseGroupIdSiteIdDriveID parses 'input' into a GroupIdSiteIdDriveId
func ParseGroupIdSiteIdDriveID(input string) (*GroupIdSiteIdDriveId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdDriveId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdDriveId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdDriveIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdDriveId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdDriveIDInsensitively(input string) (*GroupIdSiteIdDriveId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdDriveId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdDriveId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdDriveId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdDriveID checks that 'input' can be parsed as a Group Id Site Id Drive ID
func ValidateGroupIdSiteIdDriveID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdDriveID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Drive ID
func (id GroupIdSiteIdDriveId) ID() string {
	fmtString := "/groups/%s/sites/%s/drives/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.DriveId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Drive ID
func (id GroupIdSiteIdDriveId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Drive ID
func (id GroupIdSiteIdDriveId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Drive: %q", id.DriveId),
	}
	return fmt.Sprintf("Group Id Site Id Drive (%s)", strings.Join(components, "\n"))
}
