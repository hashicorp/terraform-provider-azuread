package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListContentTypeId{}

// GroupIdDriveIdListContentTypeId is a struct representing the Resource ID for a Group Id Drive Id List Content Type
type GroupIdDriveIdListContentTypeId struct {
	GroupId       string
	DriveId       string
	ContentTypeId string
}

// NewGroupIdDriveIdListContentTypeID returns a new GroupIdDriveIdListContentTypeId struct
func NewGroupIdDriveIdListContentTypeID(groupId string, driveId string, contentTypeId string) GroupIdDriveIdListContentTypeId {
	return GroupIdDriveIdListContentTypeId{
		GroupId:       groupId,
		DriveId:       driveId,
		ContentTypeId: contentTypeId,
	}
}

// ParseGroupIdDriveIdListContentTypeID parses 'input' into a GroupIdDriveIdListContentTypeId
func ParseGroupIdDriveIdListContentTypeID(input string) (*GroupIdDriveIdListContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListContentTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListContentTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdListContentTypeIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdListContentTypeId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdListContentTypeIDInsensitively(input string) (*GroupIdDriveIdListContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListContentTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListContentTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdListContentTypeId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdDriveIdListContentTypeID checks that 'input' can be parsed as a Group Id Drive Id List Content Type ID
func ValidateGroupIdDriveIdListContentTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdListContentTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id List Content Type ID
func (id GroupIdDriveIdListContentTypeId) ID() string {
	fmtString := "/groups/%s/drives/%s/list/contentTypes/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ContentTypeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id List Content Type ID
func (id GroupIdDriveIdListContentTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id List Content Type ID
func (id GroupIdDriveIdListContentTypeId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
	}
	return fmt.Sprintf("Group Id Drive Id List Content Type (%s)", strings.Join(components, "\n"))
}
