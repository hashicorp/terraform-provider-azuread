package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootListItemVersionId{}

// GroupIdDriveIdRootListItemVersionId is a struct representing the Resource ID for a Group Id Drive Id Root List Item Version
type GroupIdDriveIdRootListItemVersionId struct {
	GroupId           string
	DriveId           string
	ListItemVersionId string
}

// NewGroupIdDriveIdRootListItemVersionID returns a new GroupIdDriveIdRootListItemVersionId struct
func NewGroupIdDriveIdRootListItemVersionID(groupId string, driveId string, listItemVersionId string) GroupIdDriveIdRootListItemVersionId {
	return GroupIdDriveIdRootListItemVersionId{
		GroupId:           groupId,
		DriveId:           driveId,
		ListItemVersionId: listItemVersionId,
	}
}

// ParseGroupIdDriveIdRootListItemVersionID parses 'input' into a GroupIdDriveIdRootListItemVersionId
func ParseGroupIdDriveIdRootListItemVersionID(input string) (*GroupIdDriveIdRootListItemVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootListItemVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootListItemVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdRootListItemVersionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdRootListItemVersionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdRootListItemVersionIDInsensitively(input string) (*GroupIdDriveIdRootListItemVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootListItemVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootListItemVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdRootListItemVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ListItemVersionId, ok = input.Parsed["listItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemVersionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdRootListItemVersionID checks that 'input' can be parsed as a Group Id Drive Id Root List Item Version ID
func ValidateGroupIdDriveIdRootListItemVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdRootListItemVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Root List Item Version ID
func (id GroupIdDriveIdRootListItemVersionId) ID() string {
	fmtString := "/groups/%s/drives/%s/root/listItem/versions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ListItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Root List Item Version ID
func (id GroupIdDriveIdRootListItemVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("listItemVersionId", "listItemVersionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Root List Item Version ID
func (id GroupIdDriveIdRootListItemVersionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item Version: %q", id.ListItemVersionId),
	}
	return fmt.Sprintf("Group Id Drive Id Root List Item Version (%s)", strings.Join(components, "\n"))
}
