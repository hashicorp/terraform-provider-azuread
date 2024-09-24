package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemIdVersionId{}

// GroupIdDriveIdItemIdVersionId is a struct representing the Resource ID for a Group Id Drive Id Item Id Version
type GroupIdDriveIdItemIdVersionId struct {
	GroupId            string
	DriveId            string
	DriveItemId        string
	DriveItemVersionId string
}

// NewGroupIdDriveIdItemIdVersionID returns a new GroupIdDriveIdItemIdVersionId struct
func NewGroupIdDriveIdItemIdVersionID(groupId string, driveId string, driveItemId string, driveItemVersionId string) GroupIdDriveIdItemIdVersionId {
	return GroupIdDriveIdItemIdVersionId{
		GroupId:            groupId,
		DriveId:            driveId,
		DriveItemId:        driveItemId,
		DriveItemVersionId: driveItemVersionId,
	}
}

// ParseGroupIdDriveIdItemIdVersionID parses 'input' into a GroupIdDriveIdItemIdVersionId
func ParseGroupIdDriveIdItemIdVersionID(input string) (*GroupIdDriveIdItemIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdItemIdVersionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdItemIdVersionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdItemIdVersionIDInsensitively(input string) (*GroupIdDriveIdItemIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdItemIdVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.DriveItemVersionId, ok = input.Parsed["driveItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemVersionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdItemIdVersionID checks that 'input' can be parsed as a Group Id Drive Id Item Id Version ID
func ValidateGroupIdDriveIdItemIdVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdItemIdVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Item Id Version ID
func (id GroupIdDriveIdItemIdVersionId) ID() string {
	fmtString := "/groups/%s/drives/%s/items/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId, id.DriveItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Item Id Version ID
func (id GroupIdDriveIdItemIdVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("driveItemVersionId", "driveItemVersionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Item Id Version ID
func (id GroupIdDriveIdItemIdVersionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Drive Item Version: %q", id.DriveItemVersionId),
	}
	return fmt.Sprintf("Group Id Drive Id Item Id Version (%s)", strings.Join(components, "\n"))
}
