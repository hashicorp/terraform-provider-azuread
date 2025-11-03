package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootActivityId{}

// GroupIdDriveIdRootActivityId is a struct representing the Resource ID for a Group Id Drive Id Root Activity
type GroupIdDriveIdRootActivityId struct {
	GroupId           string
	DriveId           string
	ItemActivityOLDId string
}

// NewGroupIdDriveIdRootActivityID returns a new GroupIdDriveIdRootActivityId struct
func NewGroupIdDriveIdRootActivityID(groupId string, driveId string, itemActivityOLDId string) GroupIdDriveIdRootActivityId {
	return GroupIdDriveIdRootActivityId{
		GroupId:           groupId,
		DriveId:           driveId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseGroupIdDriveIdRootActivityID parses 'input' into a GroupIdDriveIdRootActivityId
func ParseGroupIdDriveIdRootActivityID(input string) (*GroupIdDriveIdRootActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdRootActivityIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdRootActivityId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdRootActivityIDInsensitively(input string) (*GroupIdDriveIdRootActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdRootActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ItemActivityOLDId, ok = input.Parsed["itemActivityOLDId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityOLDId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdRootActivityID checks that 'input' can be parsed as a Group Id Drive Id Root Activity ID
func ValidateGroupIdDriveIdRootActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdRootActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Root Activity ID
func (id GroupIdDriveIdRootActivityId) ID() string {
	fmtString := "/groups/%s/drives/%s/root/activities/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Root Activity ID
func (id GroupIdDriveIdRootActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Root Activity ID
func (id GroupIdDriveIdRootActivityId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Item Activity OLD: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("Group Id Drive Id Root Activity (%s)", strings.Join(components, "\n"))
}
