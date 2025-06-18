package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootExtensionId{}

// MeDriveIdRootExtensionId is a struct representing the Resource ID for a Me Drive Id Root Extension
type MeDriveIdRootExtensionId struct {
	DriveId     string
	ExtensionId string
}

// NewMeDriveIdRootExtensionID returns a new MeDriveIdRootExtensionId struct
func NewMeDriveIdRootExtensionID(driveId string, extensionId string) MeDriveIdRootExtensionId {
	return MeDriveIdRootExtensionId{
		DriveId:     driveId,
		ExtensionId: extensionId,
	}
}

// ParseMeDriveIdRootExtensionID parses 'input' into a MeDriveIdRootExtensionId
func ParseMeDriveIdRootExtensionID(input string) (*MeDriveIdRootExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdRootExtensionIDInsensitively parses 'input' case-insensitively into a MeDriveIdRootExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdRootExtensionIDInsensitively(input string) (*MeDriveIdRootExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdRootExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeDriveIdRootExtensionID checks that 'input' can be parsed as a Me Drive Id Root Extension ID
func ValidateMeDriveIdRootExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdRootExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Root Extension ID
func (id MeDriveIdRootExtensionId) ID() string {
	fmtString := "/me/drives/%s/root/extensions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Root Extension ID
func (id MeDriveIdRootExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Me Drive Id Root Extension ID
func (id MeDriveIdRootExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Drive Id Root Extension (%s)", strings.Join(components, "\n"))
}
