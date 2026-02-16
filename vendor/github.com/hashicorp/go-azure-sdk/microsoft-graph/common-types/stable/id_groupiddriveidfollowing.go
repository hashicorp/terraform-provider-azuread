package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdFollowingId{}

// GroupIdDriveIdFollowingId is a struct representing the Resource ID for a Group Id Drive Id Following
type GroupIdDriveIdFollowingId struct {
	GroupId     string
	DriveId     string
	DriveItemId string
}

// NewGroupIdDriveIdFollowingID returns a new GroupIdDriveIdFollowingId struct
func NewGroupIdDriveIdFollowingID(groupId string, driveId string, driveItemId string) GroupIdDriveIdFollowingId {
	return GroupIdDriveIdFollowingId{
		GroupId:     groupId,
		DriveId:     driveId,
		DriveItemId: driveItemId,
	}
}

// ParseGroupIdDriveIdFollowingID parses 'input' into a GroupIdDriveIdFollowingId
func ParseGroupIdDriveIdFollowingID(input string) (*GroupIdDriveIdFollowingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdFollowingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdFollowingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdFollowingIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdFollowingId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdFollowingIDInsensitively(input string) (*GroupIdDriveIdFollowingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdFollowingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdFollowingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdFollowingId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdDriveIdFollowingID checks that 'input' can be parsed as a Group Id Drive Id Following ID
func ValidateGroupIdDriveIdFollowingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdFollowingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Following ID
func (id GroupIdDriveIdFollowingId) ID() string {
	fmtString := "/groups/%s/drives/%s/following/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Following ID
func (id GroupIdDriveIdFollowingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("following", "following", "following"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Following ID
func (id GroupIdDriveIdFollowingId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
	}
	return fmt.Sprintf("Group Id Drive Id Following (%s)", strings.Join(components, "\n"))
}
