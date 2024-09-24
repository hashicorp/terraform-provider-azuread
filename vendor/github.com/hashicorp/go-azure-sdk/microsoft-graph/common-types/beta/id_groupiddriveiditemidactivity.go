package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemIdActivityId{}

// GroupIdDriveIdItemIdActivityId is a struct representing the Resource ID for a Group Id Drive Id Item Id Activity
type GroupIdDriveIdItemIdActivityId struct {
	GroupId           string
	DriveId           string
	DriveItemId       string
	ItemActivityOLDId string
}

// NewGroupIdDriveIdItemIdActivityID returns a new GroupIdDriveIdItemIdActivityId struct
func NewGroupIdDriveIdItemIdActivityID(groupId string, driveId string, driveItemId string, itemActivityOLDId string) GroupIdDriveIdItemIdActivityId {
	return GroupIdDriveIdItemIdActivityId{
		GroupId:           groupId,
		DriveId:           driveId,
		DriveItemId:       driveItemId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseGroupIdDriveIdItemIdActivityID parses 'input' into a GroupIdDriveIdItemIdActivityId
func ParseGroupIdDriveIdItemIdActivityID(input string) (*GroupIdDriveIdItemIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdItemIdActivityIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdItemIdActivityId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdItemIdActivityIDInsensitively(input string) (*GroupIdDriveIdItemIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdItemIdActivityId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdDriveIdItemIdActivityID checks that 'input' can be parsed as a Group Id Drive Id Item Id Activity ID
func ValidateGroupIdDriveIdItemIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdItemIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Item Id Activity ID
func (id GroupIdDriveIdItemIdActivityId) ID() string {
	fmtString := "/groups/%s/drives/%s/items/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Item Id Activity ID
func (id GroupIdDriveIdItemIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Item Id Activity ID
func (id GroupIdDriveIdItemIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Item Activity O L D: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("Group Id Drive Id Item Id Activity (%s)", strings.Join(components, "\n"))
}
