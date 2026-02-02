package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemIdListItemActivityId{}

// GroupIdDriveIdItemIdListItemActivityId is a struct representing the Resource ID for a Group Id Drive Id Item Id List Item Activity
type GroupIdDriveIdItemIdListItemActivityId struct {
	GroupId           string
	DriveId           string
	DriveItemId       string
	ItemActivityOLDId string
}

// NewGroupIdDriveIdItemIdListItemActivityID returns a new GroupIdDriveIdItemIdListItemActivityId struct
func NewGroupIdDriveIdItemIdListItemActivityID(groupId string, driveId string, driveItemId string, itemActivityOLDId string) GroupIdDriveIdItemIdListItemActivityId {
	return GroupIdDriveIdItemIdListItemActivityId{
		GroupId:           groupId,
		DriveId:           driveId,
		DriveItemId:       driveItemId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseGroupIdDriveIdItemIdListItemActivityID parses 'input' into a GroupIdDriveIdItemIdListItemActivityId
func ParseGroupIdDriveIdItemIdListItemActivityID(input string) (*GroupIdDriveIdItemIdListItemActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdListItemActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdListItemActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdItemIdListItemActivityIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdItemIdListItemActivityId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdItemIdListItemActivityIDInsensitively(input string) (*GroupIdDriveIdItemIdListItemActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdListItemActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdListItemActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdItemIdListItemActivityId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ItemActivityOLDId, ok = input.Parsed["itemActivityOLDId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityOLDId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdItemIdListItemActivityID checks that 'input' can be parsed as a Group Id Drive Id Item Id List Item Activity ID
func ValidateGroupIdDriveIdItemIdListItemActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdItemIdListItemActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Item Id List Item Activity ID
func (id GroupIdDriveIdItemIdListItemActivityId) ID() string {
	fmtString := "/groups/%s/drives/%s/items/%s/listItem/activities/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Item Id List Item Activity ID
func (id GroupIdDriveIdItemIdListItemActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Item Id List Item Activity ID
func (id GroupIdDriveIdItemIdListItemActivityId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Item Activity OLD: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("Group Id Drive Id Item Id List Item Activity (%s)", strings.Join(components, "\n"))
}
