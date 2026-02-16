package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemIdExtensionId{}

// GroupIdDriveIdItemIdExtensionId is a struct representing the Resource ID for a Group Id Drive Id Item Id Extension
type GroupIdDriveIdItemIdExtensionId struct {
	GroupId     string
	DriveId     string
	DriveItemId string
	ExtensionId string
}

// NewGroupIdDriveIdItemIdExtensionID returns a new GroupIdDriveIdItemIdExtensionId struct
func NewGroupIdDriveIdItemIdExtensionID(groupId string, driveId string, driveItemId string, extensionId string) GroupIdDriveIdItemIdExtensionId {
	return GroupIdDriveIdItemIdExtensionId{
		GroupId:     groupId,
		DriveId:     driveId,
		DriveItemId: driveItemId,
		ExtensionId: extensionId,
	}
}

// ParseGroupIdDriveIdItemIdExtensionID parses 'input' into a GroupIdDriveIdItemIdExtensionId
func ParseGroupIdDriveIdItemIdExtensionID(input string) (*GroupIdDriveIdItemIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdItemIdExtensionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdItemIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdItemIdExtensionIDInsensitively(input string) (*GroupIdDriveIdItemIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdItemIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdItemIdExtensionID checks that 'input' can be parsed as a Group Id Drive Id Item Id Extension ID
func ValidateGroupIdDriveIdItemIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdItemIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Item Id Extension ID
func (id GroupIdDriveIdItemIdExtensionId) ID() string {
	fmtString := "/groups/%s/drives/%s/items/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Item Id Extension ID
func (id GroupIdDriveIdItemIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Item Id Extension ID
func (id GroupIdDriveIdItemIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Id Drive Id Item Id Extension (%s)", strings.Join(components, "\n"))
}
