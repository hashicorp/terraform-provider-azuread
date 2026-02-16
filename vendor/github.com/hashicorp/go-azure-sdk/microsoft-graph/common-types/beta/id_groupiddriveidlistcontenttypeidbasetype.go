package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListContentTypeIdBaseTypeId{}

// GroupIdDriveIdListContentTypeIdBaseTypeId is a struct representing the Resource ID for a Group Id Drive Id List Content Type Id Base Type
type GroupIdDriveIdListContentTypeIdBaseTypeId struct {
	GroupId        string
	DriveId        string
	ContentTypeId  string
	ContentTypeId1 string
}

// NewGroupIdDriveIdListContentTypeIdBaseTypeID returns a new GroupIdDriveIdListContentTypeIdBaseTypeId struct
func NewGroupIdDriveIdListContentTypeIdBaseTypeID(groupId string, driveId string, contentTypeId string, contentTypeId1 string) GroupIdDriveIdListContentTypeIdBaseTypeId {
	return GroupIdDriveIdListContentTypeIdBaseTypeId{
		GroupId:        groupId,
		DriveId:        driveId,
		ContentTypeId:  contentTypeId,
		ContentTypeId1: contentTypeId1,
	}
}

// ParseGroupIdDriveIdListContentTypeIdBaseTypeID parses 'input' into a GroupIdDriveIdListContentTypeIdBaseTypeId
func ParseGroupIdDriveIdListContentTypeIdBaseTypeID(input string) (*GroupIdDriveIdListContentTypeIdBaseTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListContentTypeIdBaseTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListContentTypeIdBaseTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdListContentTypeIdBaseTypeIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdListContentTypeIdBaseTypeId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdListContentTypeIdBaseTypeIDInsensitively(input string) (*GroupIdDriveIdListContentTypeIdBaseTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListContentTypeIdBaseTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListContentTypeIdBaseTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdListContentTypeIdBaseTypeId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ContentTypeId1, ok = input.Parsed["contentTypeId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId1", input)
	}

	return nil
}

// ValidateGroupIdDriveIdListContentTypeIdBaseTypeID checks that 'input' can be parsed as a Group Id Drive Id List Content Type Id Base Type ID
func ValidateGroupIdDriveIdListContentTypeIdBaseTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdListContentTypeIdBaseTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id List Content Type Id Base Type ID
func (id GroupIdDriveIdListContentTypeIdBaseTypeId) ID() string {
	fmtString := "/groups/%s/drives/%s/list/contentTypes/%s/baseTypes/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ContentTypeId, id.ContentTypeId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id List Content Type Id Base Type ID
func (id GroupIdDriveIdListContentTypeIdBaseTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("baseTypes", "baseTypes", "baseTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId1", "contentTypeId1"),
	}
}

// String returns a human-readable description of this Group Id Drive Id List Content Type Id Base Type ID
func (id GroupIdDriveIdListContentTypeIdBaseTypeId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Content Type Id 1: %q", id.ContentTypeId1),
	}
	return fmt.Sprintf("Group Id Drive Id List Content Type Id Base Type (%s)", strings.Join(components, "\n"))
}
