package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListItemIdVersionId{}

// GroupIdDriveIdListItemIdVersionId is a struct representing the Resource ID for a Group Id Drive Id List Item Id Version
type GroupIdDriveIdListItemIdVersionId struct {
	GroupId           string
	DriveId           string
	ListItemId        string
	ListItemVersionId string
}

// NewGroupIdDriveIdListItemIdVersionID returns a new GroupIdDriveIdListItemIdVersionId struct
func NewGroupIdDriveIdListItemIdVersionID(groupId string, driveId string, listItemId string, listItemVersionId string) GroupIdDriveIdListItemIdVersionId {
	return GroupIdDriveIdListItemIdVersionId{
		GroupId:           groupId,
		DriveId:           driveId,
		ListItemId:        listItemId,
		ListItemVersionId: listItemVersionId,
	}
}

// ParseGroupIdDriveIdListItemIdVersionID parses 'input' into a GroupIdDriveIdListItemIdVersionId
func ParseGroupIdDriveIdListItemIdVersionID(input string) (*GroupIdDriveIdListItemIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListItemIdVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListItemIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdListItemIdVersionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdListItemIdVersionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdListItemIdVersionIDInsensitively(input string) (*GroupIdDriveIdListItemIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListItemIdVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListItemIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdListItemIdVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ListItemVersionId, ok = input.Parsed["listItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemVersionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdListItemIdVersionID checks that 'input' can be parsed as a Group Id Drive Id List Item Id Version ID
func ValidateGroupIdDriveIdListItemIdVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdListItemIdVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id List Item Id Version ID
func (id GroupIdDriveIdListItemIdVersionId) ID() string {
	fmtString := "/groups/%s/drives/%s/list/items/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ListItemId, id.ListItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id List Item Id Version ID
func (id GroupIdDriveIdListItemIdVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("listItemVersionId", "listItemVersionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id List Item Id Version ID
func (id GroupIdDriveIdListItemIdVersionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("List Item Version: %q", id.ListItemVersionId),
	}
	return fmt.Sprintf("Group Id Drive Id List Item Id Version (%s)", strings.Join(components, "\n"))
}
