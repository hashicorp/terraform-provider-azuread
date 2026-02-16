package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootExtensionId{}

// GroupIdDriveIdRootExtensionId is a struct representing the Resource ID for a Group Id Drive Id Root Extension
type GroupIdDriveIdRootExtensionId struct {
	GroupId     string
	DriveId     string
	ExtensionId string
}

// NewGroupIdDriveIdRootExtensionID returns a new GroupIdDriveIdRootExtensionId struct
func NewGroupIdDriveIdRootExtensionID(groupId string, driveId string, extensionId string) GroupIdDriveIdRootExtensionId {
	return GroupIdDriveIdRootExtensionId{
		GroupId:     groupId,
		DriveId:     driveId,
		ExtensionId: extensionId,
	}
}

// ParseGroupIdDriveIdRootExtensionID parses 'input' into a GroupIdDriveIdRootExtensionId
func ParseGroupIdDriveIdRootExtensionID(input string) (*GroupIdDriveIdRootExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdRootExtensionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdRootExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdRootExtensionIDInsensitively(input string) (*GroupIdDriveIdRootExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdRootExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdRootExtensionID checks that 'input' can be parsed as a Group Id Drive Id Root Extension ID
func ValidateGroupIdDriveIdRootExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdRootExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Root Extension ID
func (id GroupIdDriveIdRootExtensionId) ID() string {
	fmtString := "/groups/%s/drives/%s/root/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Root Extension ID
func (id GroupIdDriveIdRootExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Root Extension ID
func (id GroupIdDriveIdRootExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Id Drive Id Root Extension (%s)", strings.Join(components, "\n"))
}
