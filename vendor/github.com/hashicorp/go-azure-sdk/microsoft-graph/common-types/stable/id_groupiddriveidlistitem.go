package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListItemId{}

// GroupIdDriveIdListItemId is a struct representing the Resource ID for a Group Id Drive Id List Item
type GroupIdDriveIdListItemId struct {
	GroupId    string
	DriveId    string
	ListItemId string
}

// NewGroupIdDriveIdListItemID returns a new GroupIdDriveIdListItemId struct
func NewGroupIdDriveIdListItemID(groupId string, driveId string, listItemId string) GroupIdDriveIdListItemId {
	return GroupIdDriveIdListItemId{
		GroupId:    groupId,
		DriveId:    driveId,
		ListItemId: listItemId,
	}
}

// ParseGroupIdDriveIdListItemID parses 'input' into a GroupIdDriveIdListItemId
func ParseGroupIdDriveIdListItemID(input string) (*GroupIdDriveIdListItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdListItemIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdListItemId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdListItemIDInsensitively(input string) (*GroupIdDriveIdListItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdListItemId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ListItemId, ok = input.Parsed["listItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdListItemID checks that 'input' can be parsed as a Group Id Drive Id List Item ID
func ValidateGroupIdDriveIdListItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdListItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id List Item ID
func (id GroupIdDriveIdListItemId) ID() string {
	fmtString := "/groups/%s/drives/%s/list/items/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ListItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id List Item ID
func (id GroupIdDriveIdListItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id List Item ID
func (id GroupIdDriveIdListItemId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
	}
	return fmt.Sprintf("Group Id Drive Id List Item (%s)", strings.Join(components, "\n"))
}
