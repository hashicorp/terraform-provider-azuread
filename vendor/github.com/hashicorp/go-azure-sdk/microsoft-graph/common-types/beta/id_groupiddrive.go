package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveId{}

// GroupIdDriveId is a struct representing the Resource ID for a Group Id Drive
type GroupIdDriveId struct {
	GroupId string
	DriveId string
}

// NewGroupIdDriveID returns a new GroupIdDriveId struct
func NewGroupIdDriveID(groupId string, driveId string) GroupIdDriveId {
	return GroupIdDriveId{
		GroupId: groupId,
		DriveId: driveId,
	}
}

// ParseGroupIdDriveID parses 'input' into a GroupIdDriveId
func ParseGroupIdDriveID(input string) (*GroupIdDriveId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIDInsensitively parses 'input' case-insensitively into a GroupIdDriveId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIDInsensitively(input string) (*GroupIdDriveId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	return nil
}

// ValidateGroupIdDriveID checks that 'input' can be parsed as a Group Id Drive ID
func ValidateGroupIdDriveID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive ID
func (id GroupIdDriveId) ID() string {
	fmtString := "/groups/%s/drives/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive ID
func (id GroupIdDriveId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
	}
}

// String returns a human-readable description of this Group Id Drive ID
func (id GroupIdDriveId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
	}
	return fmt.Sprintf("Group Id Drive (%s)", strings.Join(components, "\n"))
}
