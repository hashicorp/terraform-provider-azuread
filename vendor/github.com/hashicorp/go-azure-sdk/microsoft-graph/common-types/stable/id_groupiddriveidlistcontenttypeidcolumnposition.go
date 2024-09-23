package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListContentTypeIdColumnPositionId{}

// GroupIdDriveIdListContentTypeIdColumnPositionId is a struct representing the Resource ID for a Group Id Drive Id List Content Type Id Column Position
type GroupIdDriveIdListContentTypeIdColumnPositionId struct {
	GroupId            string
	DriveId            string
	ContentTypeId      string
	ColumnDefinitionId string
}

// NewGroupIdDriveIdListContentTypeIdColumnPositionID returns a new GroupIdDriveIdListContentTypeIdColumnPositionId struct
func NewGroupIdDriveIdListContentTypeIdColumnPositionID(groupId string, driveId string, contentTypeId string, columnDefinitionId string) GroupIdDriveIdListContentTypeIdColumnPositionId {
	return GroupIdDriveIdListContentTypeIdColumnPositionId{
		GroupId:            groupId,
		DriveId:            driveId,
		ContentTypeId:      contentTypeId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseGroupIdDriveIdListContentTypeIdColumnPositionID parses 'input' into a GroupIdDriveIdListContentTypeIdColumnPositionId
func ParseGroupIdDriveIdListContentTypeIdColumnPositionID(input string) (*GroupIdDriveIdListContentTypeIdColumnPositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListContentTypeIdColumnPositionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListContentTypeIdColumnPositionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdListContentTypeIdColumnPositionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdListContentTypeIdColumnPositionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdListContentTypeIdColumnPositionIDInsensitively(input string) (*GroupIdDriveIdListContentTypeIdColumnPositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListContentTypeIdColumnPositionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListContentTypeIdColumnPositionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdListContentTypeIdColumnPositionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ContentTypeId, ok = input.Parsed["contentTypeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", input)
	}

	if id.ColumnDefinitionId, ok = input.Parsed["columnDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdListContentTypeIdColumnPositionID checks that 'input' can be parsed as a Group Id Drive Id List Content Type Id Column Position ID
func ValidateGroupIdDriveIdListContentTypeIdColumnPositionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdListContentTypeIdColumnPositionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id List Content Type Id Column Position ID
func (id GroupIdDriveIdListContentTypeIdColumnPositionId) ID() string {
	fmtString := "/groups/%s/drives/%s/list/contentTypes/%s/columnPositions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ContentTypeId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id List Content Type Id Column Position ID
func (id GroupIdDriveIdListContentTypeIdColumnPositionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("columnPositions", "columnPositions", "columnPositions"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id List Content Type Id Column Position ID
func (id GroupIdDriveIdListContentTypeIdColumnPositionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("Group Id Drive Id List Content Type Id Column Position (%s)", strings.Join(components, "\n"))
}
