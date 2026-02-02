package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootAnalyticsItemActivityStatId{}

// GroupIdDriveIdRootAnalyticsItemActivityStatId is a struct representing the Resource ID for a Group Id Drive Id Root Analytics Item Activity Stat
type GroupIdDriveIdRootAnalyticsItemActivityStatId struct {
	GroupId            string
	DriveId            string
	ItemActivityStatId string
}

// NewGroupIdDriveIdRootAnalyticsItemActivityStatID returns a new GroupIdDriveIdRootAnalyticsItemActivityStatId struct
func NewGroupIdDriveIdRootAnalyticsItemActivityStatID(groupId string, driveId string, itemActivityStatId string) GroupIdDriveIdRootAnalyticsItemActivityStatId {
	return GroupIdDriveIdRootAnalyticsItemActivityStatId{
		GroupId:            groupId,
		DriveId:            driveId,
		ItemActivityStatId: itemActivityStatId,
	}
}

// ParseGroupIdDriveIdRootAnalyticsItemActivityStatID parses 'input' into a GroupIdDriveIdRootAnalyticsItemActivityStatId
func ParseGroupIdDriveIdRootAnalyticsItemActivityStatID(input string) (*GroupIdDriveIdRootAnalyticsItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootAnalyticsItemActivityStatId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootAnalyticsItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdRootAnalyticsItemActivityStatIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdRootAnalyticsItemActivityStatId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdRootAnalyticsItemActivityStatIDInsensitively(input string) (*GroupIdDriveIdRootAnalyticsItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootAnalyticsItemActivityStatId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootAnalyticsItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdRootAnalyticsItemActivityStatId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ItemActivityStatId, ok = input.Parsed["itemActivityStatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityStatId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdRootAnalyticsItemActivityStatID checks that 'input' can be parsed as a Group Id Drive Id Root Analytics Item Activity Stat ID
func ValidateGroupIdDriveIdRootAnalyticsItemActivityStatID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdRootAnalyticsItemActivityStatID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Root Analytics Item Activity Stat ID
func (id GroupIdDriveIdRootAnalyticsItemActivityStatId) ID() string {
	fmtString := "/groups/%s/drives/%s/root/analytics/itemActivityStats/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ItemActivityStatId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Root Analytics Item Activity Stat ID
func (id GroupIdDriveIdRootAnalyticsItemActivityStatId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("itemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Root Analytics Item Activity Stat ID
func (id GroupIdDriveIdRootAnalyticsItemActivityStatId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
	}
	return fmt.Sprintf("Group Id Drive Id Root Analytics Item Activity Stat (%s)", strings.Join(components, "\n"))
}
