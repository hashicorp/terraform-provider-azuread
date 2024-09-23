package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListContentTypeIdColumnLinkId{}

// GroupIdDriveIdListContentTypeIdColumnLinkId is a struct representing the Resource ID for a Group Id Drive Id List Content Type Id Column Link
type GroupIdDriveIdListContentTypeIdColumnLinkId struct {
	GroupId       string
	DriveId       string
	ContentTypeId string
	ColumnLinkId  string
}

// NewGroupIdDriveIdListContentTypeIdColumnLinkID returns a new GroupIdDriveIdListContentTypeIdColumnLinkId struct
func NewGroupIdDriveIdListContentTypeIdColumnLinkID(groupId string, driveId string, contentTypeId string, columnLinkId string) GroupIdDriveIdListContentTypeIdColumnLinkId {
	return GroupIdDriveIdListContentTypeIdColumnLinkId{
		GroupId:       groupId,
		DriveId:       driveId,
		ContentTypeId: contentTypeId,
		ColumnLinkId:  columnLinkId,
	}
}

// ParseGroupIdDriveIdListContentTypeIdColumnLinkID parses 'input' into a GroupIdDriveIdListContentTypeIdColumnLinkId
func ParseGroupIdDriveIdListContentTypeIdColumnLinkID(input string) (*GroupIdDriveIdListContentTypeIdColumnLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListContentTypeIdColumnLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListContentTypeIdColumnLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdListContentTypeIdColumnLinkIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdListContentTypeIdColumnLinkId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdListContentTypeIdColumnLinkIDInsensitively(input string) (*GroupIdDriveIdListContentTypeIdColumnLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListContentTypeIdColumnLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListContentTypeIdColumnLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdListContentTypeIdColumnLinkId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ColumnLinkId, ok = input.Parsed["columnLinkId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnLinkId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdListContentTypeIdColumnLinkID checks that 'input' can be parsed as a Group Id Drive Id List Content Type Id Column Link ID
func ValidateGroupIdDriveIdListContentTypeIdColumnLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdListContentTypeIdColumnLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id List Content Type Id Column Link ID
func (id GroupIdDriveIdListContentTypeIdColumnLinkId) ID() string {
	fmtString := "/groups/%s/drives/%s/list/contentTypes/%s/columnLinks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ContentTypeId, id.ColumnLinkId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id List Content Type Id Column Link ID
func (id GroupIdDriveIdListContentTypeIdColumnLinkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("columnLinks", "columnLinks", "columnLinks"),
		resourceids.UserSpecifiedSegment("columnLinkId", "columnLinkId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id List Content Type Id Column Link ID
func (id GroupIdDriveIdListContentTypeIdColumnLinkId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Link: %q", id.ColumnLinkId),
	}
	return fmt.Sprintf("Group Id Drive Id List Content Type Id Column Link (%s)", strings.Join(components, "\n"))
}
