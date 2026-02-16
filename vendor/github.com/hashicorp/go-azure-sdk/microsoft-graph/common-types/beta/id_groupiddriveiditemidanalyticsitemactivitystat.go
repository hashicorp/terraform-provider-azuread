package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemIdAnalyticsItemActivityStatId{}

// GroupIdDriveIdItemIdAnalyticsItemActivityStatId is a struct representing the Resource ID for a Group Id Drive Id Item Id Analytics Item Activity Stat
type GroupIdDriveIdItemIdAnalyticsItemActivityStatId struct {
	GroupId            string
	DriveId            string
	DriveItemId        string
	ItemActivityStatId string
}

// NewGroupIdDriveIdItemIdAnalyticsItemActivityStatID returns a new GroupIdDriveIdItemIdAnalyticsItemActivityStatId struct
func NewGroupIdDriveIdItemIdAnalyticsItemActivityStatID(groupId string, driveId string, driveItemId string, itemActivityStatId string) GroupIdDriveIdItemIdAnalyticsItemActivityStatId {
	return GroupIdDriveIdItemIdAnalyticsItemActivityStatId{
		GroupId:            groupId,
		DriveId:            driveId,
		DriveItemId:        driveItemId,
		ItemActivityStatId: itemActivityStatId,
	}
}

// ParseGroupIdDriveIdItemIdAnalyticsItemActivityStatID parses 'input' into a GroupIdDriveIdItemIdAnalyticsItemActivityStatId
func ParseGroupIdDriveIdItemIdAnalyticsItemActivityStatID(input string) (*GroupIdDriveIdItemIdAnalyticsItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdAnalyticsItemActivityStatId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdAnalyticsItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdItemIdAnalyticsItemActivityStatIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdItemIdAnalyticsItemActivityStatId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdItemIdAnalyticsItemActivityStatIDInsensitively(input string) (*GroupIdDriveIdItemIdAnalyticsItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdAnalyticsItemActivityStatId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdAnalyticsItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdItemIdAnalyticsItemActivityStatId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	if id.ItemActivityStatId, ok = input.Parsed["itemActivityStatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityStatId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdItemIdAnalyticsItemActivityStatID checks that 'input' can be parsed as a Group Id Drive Id Item Id Analytics Item Activity Stat ID
func ValidateGroupIdDriveIdItemIdAnalyticsItemActivityStatID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdItemIdAnalyticsItemActivityStatID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Item Id Analytics Item Activity Stat ID
func (id GroupIdDriveIdItemIdAnalyticsItemActivityStatId) ID() string {
	fmtString := "/groups/%s/drives/%s/items/%s/analytics/itemActivityStats/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId, id.ItemActivityStatId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Item Id Analytics Item Activity Stat ID
func (id GroupIdDriveIdItemIdAnalyticsItemActivityStatId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("itemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Item Id Analytics Item Activity Stat ID
func (id GroupIdDriveIdItemIdAnalyticsItemActivityStatId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
	}
	return fmt.Sprintf("Group Id Drive Id Item Id Analytics Item Activity Stat (%s)", strings.Join(components, "\n"))
}
