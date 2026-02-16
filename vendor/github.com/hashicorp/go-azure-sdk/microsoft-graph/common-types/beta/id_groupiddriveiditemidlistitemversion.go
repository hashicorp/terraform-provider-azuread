package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemIdListItemVersionId{}

// GroupIdDriveIdItemIdListItemVersionId is a struct representing the Resource ID for a Group Id Drive Id Item Id List Item Version
type GroupIdDriveIdItemIdListItemVersionId struct {
	GroupId           string
	DriveId           string
	DriveItemId       string
	ListItemVersionId string
}

// NewGroupIdDriveIdItemIdListItemVersionID returns a new GroupIdDriveIdItemIdListItemVersionId struct
func NewGroupIdDriveIdItemIdListItemVersionID(groupId string, driveId string, driveItemId string, listItemVersionId string) GroupIdDriveIdItemIdListItemVersionId {
	return GroupIdDriveIdItemIdListItemVersionId{
		GroupId:           groupId,
		DriveId:           driveId,
		DriveItemId:       driveItemId,
		ListItemVersionId: listItemVersionId,
	}
}

// ParseGroupIdDriveIdItemIdListItemVersionID parses 'input' into a GroupIdDriveIdItemIdListItemVersionId
func ParseGroupIdDriveIdItemIdListItemVersionID(input string) (*GroupIdDriveIdItemIdListItemVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdListItemVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdListItemVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdItemIdListItemVersionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdItemIdListItemVersionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdItemIdListItemVersionIDInsensitively(input string) (*GroupIdDriveIdItemIdListItemVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdListItemVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdListItemVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdItemIdListItemVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ListItemVersionId, ok = input.Parsed["listItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemVersionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdItemIdListItemVersionID checks that 'input' can be parsed as a Group Id Drive Id Item Id List Item Version ID
func ValidateGroupIdDriveIdItemIdListItemVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdItemIdListItemVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Item Id List Item Version ID
func (id GroupIdDriveIdItemIdListItemVersionId) ID() string {
	fmtString := "/groups/%s/drives/%s/items/%s/listItem/versions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId, id.ListItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Item Id List Item Version ID
func (id GroupIdDriveIdItemIdListItemVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("listItemVersionId", "listItemVersionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Item Id List Item Version ID
func (id GroupIdDriveIdItemIdListItemVersionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("List Item Version: %q", id.ListItemVersionId),
	}
	return fmt.Sprintf("Group Id Drive Id Item Id List Item Version (%s)", strings.Join(components, "\n"))
}
