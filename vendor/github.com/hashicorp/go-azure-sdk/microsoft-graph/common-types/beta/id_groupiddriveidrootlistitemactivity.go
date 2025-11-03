package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootListItemActivityId{}

// GroupIdDriveIdRootListItemActivityId is a struct representing the Resource ID for a Group Id Drive Id Root List Item Activity
type GroupIdDriveIdRootListItemActivityId struct {
	GroupId           string
	DriveId           string
	ItemActivityOLDId string
}

// NewGroupIdDriveIdRootListItemActivityID returns a new GroupIdDriveIdRootListItemActivityId struct
func NewGroupIdDriveIdRootListItemActivityID(groupId string, driveId string, itemActivityOLDId string) GroupIdDriveIdRootListItemActivityId {
	return GroupIdDriveIdRootListItemActivityId{
		GroupId:           groupId,
		DriveId:           driveId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseGroupIdDriveIdRootListItemActivityID parses 'input' into a GroupIdDriveIdRootListItemActivityId
func ParseGroupIdDriveIdRootListItemActivityID(input string) (*GroupIdDriveIdRootListItemActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootListItemActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootListItemActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdRootListItemActivityIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdRootListItemActivityId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdRootListItemActivityIDInsensitively(input string) (*GroupIdDriveIdRootListItemActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootListItemActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootListItemActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdRootListItemActivityId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdDriveIdRootListItemActivityID checks that 'input' can be parsed as a Group Id Drive Id Root List Item Activity ID
func ValidateGroupIdDriveIdRootListItemActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdRootListItemActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Root List Item Activity ID
func (id GroupIdDriveIdRootListItemActivityId) ID() string {
	fmtString := "/groups/%s/drives/%s/root/listItem/activities/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Root List Item Activity ID
func (id GroupIdDriveIdRootListItemActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Root List Item Activity ID
func (id GroupIdDriveIdRootListItemActivityId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Item Activity OLD: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("Group Id Drive Id Root List Item Activity (%s)", strings.Join(components, "\n"))
}
