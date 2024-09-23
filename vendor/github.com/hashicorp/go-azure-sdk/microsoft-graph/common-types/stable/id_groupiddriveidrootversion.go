package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootVersionId{}

// GroupIdDriveIdRootVersionId is a struct representing the Resource ID for a Group Id Drive Id Root Version
type GroupIdDriveIdRootVersionId struct {
	GroupId            string
	DriveId            string
	DriveItemVersionId string
}

// NewGroupIdDriveIdRootVersionID returns a new GroupIdDriveIdRootVersionId struct
func NewGroupIdDriveIdRootVersionID(groupId string, driveId string, driveItemVersionId string) GroupIdDriveIdRootVersionId {
	return GroupIdDriveIdRootVersionId{
		GroupId:            groupId,
		DriveId:            driveId,
		DriveItemVersionId: driveItemVersionId,
	}
}

// ParseGroupIdDriveIdRootVersionID parses 'input' into a GroupIdDriveIdRootVersionId
func ParseGroupIdDriveIdRootVersionID(input string) (*GroupIdDriveIdRootVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdRootVersionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdRootVersionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdRootVersionIDInsensitively(input string) (*GroupIdDriveIdRootVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdRootVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemVersionId, ok = input.Parsed["driveItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemVersionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdRootVersionID checks that 'input' can be parsed as a Group Id Drive Id Root Version ID
func ValidateGroupIdDriveIdRootVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdRootVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Root Version ID
func (id GroupIdDriveIdRootVersionId) ID() string {
	fmtString := "/groups/%s/drives/%s/root/versions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Root Version ID
func (id GroupIdDriveIdRootVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("driveItemVersionId", "driveItemVersionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Root Version ID
func (id GroupIdDriveIdRootVersionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item Version: %q", id.DriveItemVersionId),
	}
	return fmt.Sprintf("Group Id Drive Id Root Version (%s)", strings.Join(components, "\n"))
}
