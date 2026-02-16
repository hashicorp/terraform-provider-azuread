package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootChildId{}

// GroupIdDriveIdRootChildId is a struct representing the Resource ID for a Group Id Drive Id Root Child
type GroupIdDriveIdRootChildId struct {
	GroupId     string
	DriveId     string
	DriveItemId string
}

// NewGroupIdDriveIdRootChildID returns a new GroupIdDriveIdRootChildId struct
func NewGroupIdDriveIdRootChildID(groupId string, driveId string, driveItemId string) GroupIdDriveIdRootChildId {
	return GroupIdDriveIdRootChildId{
		GroupId:     groupId,
		DriveId:     driveId,
		DriveItemId: driveItemId,
	}
}

// ParseGroupIdDriveIdRootChildID parses 'input' into a GroupIdDriveIdRootChildId
func ParseGroupIdDriveIdRootChildID(input string) (*GroupIdDriveIdRootChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootChildId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdRootChildIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdRootChildId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdRootChildIDInsensitively(input string) (*GroupIdDriveIdRootChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootChildId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdRootChildId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdDriveIdRootChildID checks that 'input' can be parsed as a Group Id Drive Id Root Child ID
func ValidateGroupIdDriveIdRootChildID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdRootChildID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Root Child ID
func (id GroupIdDriveIdRootChildId) ID() string {
	fmtString := "/groups/%s/drives/%s/root/children/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Root Child ID
func (id GroupIdDriveIdRootChildId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("children", "children", "children"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Root Child ID
func (id GroupIdDriveIdRootChildId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
	}
	return fmt.Sprintf("Group Id Drive Id Root Child (%s)", strings.Join(components, "\n"))
}
