package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemIdChildId{}

// GroupIdDriveIdItemIdChildId is a struct representing the Resource ID for a Group Id Drive Id Item Id Child
type GroupIdDriveIdItemIdChildId struct {
	GroupId      string
	DriveId      string
	DriveItemId  string
	DriveItemId1 string
}

// NewGroupIdDriveIdItemIdChildID returns a new GroupIdDriveIdItemIdChildId struct
func NewGroupIdDriveIdItemIdChildID(groupId string, driveId string, driveItemId string, driveItemId1 string) GroupIdDriveIdItemIdChildId {
	return GroupIdDriveIdItemIdChildId{
		GroupId:      groupId,
		DriveId:      driveId,
		DriveItemId:  driveItemId,
		DriveItemId1: driveItemId1,
	}
}

// ParseGroupIdDriveIdItemIdChildID parses 'input' into a GroupIdDriveIdItemIdChildId
func ParseGroupIdDriveIdItemIdChildID(input string) (*GroupIdDriveIdItemIdChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdChildId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdItemIdChildIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdItemIdChildId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdItemIdChildIDInsensitively(input string) (*GroupIdDriveIdItemIdChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdChildId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdItemIdChildId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.DriveItemId1, ok = input.Parsed["driveItemId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId1", input)
	}

	return nil
}

// ValidateGroupIdDriveIdItemIdChildID checks that 'input' can be parsed as a Group Id Drive Id Item Id Child ID
func ValidateGroupIdDriveIdItemIdChildID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdItemIdChildID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Item Id Child ID
func (id GroupIdDriveIdItemIdChildId) ID() string {
	fmtString := "/groups/%s/drives/%s/items/%s/children/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId, id.DriveItemId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Item Id Child ID
func (id GroupIdDriveIdItemIdChildId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("children", "children", "children"),
		resourceids.UserSpecifiedSegment("driveItemId1", "driveItemId1"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Item Id Child ID
func (id GroupIdDriveIdItemIdChildId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Drive Item Id 1: %q", id.DriveItemId1),
	}
	return fmt.Sprintf("Group Id Drive Id Item Id Child (%s)", strings.Join(components, "\n"))
}
