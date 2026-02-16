package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemId{}

// GroupIdDriveIdItemId is a struct representing the Resource ID for a Group Id Drive Id Item
type GroupIdDriveIdItemId struct {
	GroupId     string
	DriveId     string
	DriveItemId string
}

// NewGroupIdDriveIdItemID returns a new GroupIdDriveIdItemId struct
func NewGroupIdDriveIdItemID(groupId string, driveId string, driveItemId string) GroupIdDriveIdItemId {
	return GroupIdDriveIdItemId{
		GroupId:     groupId,
		DriveId:     driveId,
		DriveItemId: driveItemId,
	}
}

// ParseGroupIdDriveIdItemID parses 'input' into a GroupIdDriveIdItemId
func ParseGroupIdDriveIdItemID(input string) (*GroupIdDriveIdItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdItemIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdItemId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdItemIDInsensitively(input string) (*GroupIdDriveIdItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdItemId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdDriveIdItemID checks that 'input' can be parsed as a Group Id Drive Id Item ID
func ValidateGroupIdDriveIdItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Item ID
func (id GroupIdDriveIdItemId) ID() string {
	fmtString := "/groups/%s/drives/%s/items/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Item ID
func (id GroupIdDriveIdItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Item ID
func (id GroupIdDriveIdItemId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
	}
	return fmt.Sprintf("Group Id Drive Id Item (%s)", strings.Join(components, "\n"))
}
