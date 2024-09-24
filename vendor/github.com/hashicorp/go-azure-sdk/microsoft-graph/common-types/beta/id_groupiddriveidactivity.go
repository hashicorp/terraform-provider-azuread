package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdActivityId{}

// GroupIdDriveIdActivityId is a struct representing the Resource ID for a Group Id Drive Id Activity
type GroupIdDriveIdActivityId struct {
	GroupId           string
	DriveId           string
	ItemActivityOLDId string
}

// NewGroupIdDriveIdActivityID returns a new GroupIdDriveIdActivityId struct
func NewGroupIdDriveIdActivityID(groupId string, driveId string, itemActivityOLDId string) GroupIdDriveIdActivityId {
	return GroupIdDriveIdActivityId{
		GroupId:           groupId,
		DriveId:           driveId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseGroupIdDriveIdActivityID parses 'input' into a GroupIdDriveIdActivityId
func ParseGroupIdDriveIdActivityID(input string) (*GroupIdDriveIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdActivityIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdActivityId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdActivityIDInsensitively(input string) (*GroupIdDriveIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdActivityId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdDriveIdActivityID checks that 'input' can be parsed as a Group Id Drive Id Activity ID
func ValidateGroupIdDriveIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Activity ID
func (id GroupIdDriveIdActivityId) ID() string {
	fmtString := "/groups/%s/drives/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Activity ID
func (id GroupIdDriveIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Activity ID
func (id GroupIdDriveIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Item Activity O L D: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("Group Id Drive Id Activity (%s)", strings.Join(components, "\n"))
}
