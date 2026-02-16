package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdExtensionId{}

// MeDriveIdItemIdExtensionId is a struct representing the Resource ID for a Me Drive Id Item Id Extension
type MeDriveIdItemIdExtensionId struct {
	DriveId     string
	DriveItemId string
	ExtensionId string
}

// NewMeDriveIdItemIdExtensionID returns a new MeDriveIdItemIdExtensionId struct
func NewMeDriveIdItemIdExtensionID(driveId string, driveItemId string, extensionId string) MeDriveIdItemIdExtensionId {
	return MeDriveIdItemIdExtensionId{
		DriveId:     driveId,
		DriveItemId: driveItemId,
		ExtensionId: extensionId,
	}
}

// ParseMeDriveIdItemIdExtensionID parses 'input' into a MeDriveIdItemIdExtensionId
func ParseMeDriveIdItemIdExtensionID(input string) (*MeDriveIdItemIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdItemIdExtensionIDInsensitively parses 'input' case-insensitively into a MeDriveIdItemIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdItemIdExtensionIDInsensitively(input string) (*MeDriveIdItemIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdItemIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeDriveIdItemIdExtensionID checks that 'input' can be parsed as a Me Drive Id Item Id Extension ID
func ValidateMeDriveIdItemIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdItemIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Item Id Extension ID
func (id MeDriveIdItemIdExtensionId) ID() string {
	fmtString := "/me/drives/%s/items/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Item Id Extension ID
func (id MeDriveIdItemIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Me Drive Id Item Id Extension ID
func (id MeDriveIdItemIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Drive Id Item Id Extension (%s)", strings.Join(components, "\n"))
}
