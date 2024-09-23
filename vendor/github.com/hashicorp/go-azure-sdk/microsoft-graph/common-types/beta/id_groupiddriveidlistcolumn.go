package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListColumnId{}

// GroupIdDriveIdListColumnId is a struct representing the Resource ID for a Group Id Drive Id List Column
type GroupIdDriveIdListColumnId struct {
	GroupId            string
	DriveId            string
	ColumnDefinitionId string
}

// NewGroupIdDriveIdListColumnID returns a new GroupIdDriveIdListColumnId struct
func NewGroupIdDriveIdListColumnID(groupId string, driveId string, columnDefinitionId string) GroupIdDriveIdListColumnId {
	return GroupIdDriveIdListColumnId{
		GroupId:            groupId,
		DriveId:            driveId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseGroupIdDriveIdListColumnID parses 'input' into a GroupIdDriveIdListColumnId
func ParseGroupIdDriveIdListColumnID(input string) (*GroupIdDriveIdListColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdListColumnIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdListColumnId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdListColumnIDInsensitively(input string) (*GroupIdDriveIdListColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdListColumnId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ColumnDefinitionId, ok = input.Parsed["columnDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdListColumnID checks that 'input' can be parsed as a Group Id Drive Id List Column ID
func ValidateGroupIdDriveIdListColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdListColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id List Column ID
func (id GroupIdDriveIdListColumnId) ID() string {
	fmtString := "/groups/%s/drives/%s/list/columns/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id List Column ID
func (id GroupIdDriveIdListColumnId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("columns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id List Column ID
func (id GroupIdDriveIdListColumnId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("Group Id Drive Id List Column (%s)", strings.Join(components, "\n"))
}
