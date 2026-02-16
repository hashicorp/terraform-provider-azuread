package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdBundleId{}

// MeDriveIdBundleId is a struct representing the Resource ID for a Me Drive Id Bundle
type MeDriveIdBundleId struct {
	DriveId     string
	DriveItemId string
}

// NewMeDriveIdBundleID returns a new MeDriveIdBundleId struct
func NewMeDriveIdBundleID(driveId string, driveItemId string) MeDriveIdBundleId {
	return MeDriveIdBundleId{
		DriveId:     driveId,
		DriveItemId: driveItemId,
	}
}

// ParseMeDriveIdBundleID parses 'input' into a MeDriveIdBundleId
func ParseMeDriveIdBundleID(input string) (*MeDriveIdBundleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdBundleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdBundleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdBundleIDInsensitively parses 'input' case-insensitively into a MeDriveIdBundleId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdBundleIDInsensitively(input string) (*MeDriveIdBundleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdBundleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdBundleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdBundleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	return nil
}

// ValidateMeDriveIdBundleID checks that 'input' can be parsed as a Me Drive Id Bundle ID
func ValidateMeDriveIdBundleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdBundleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Bundle ID
func (id MeDriveIdBundleId) ID() string {
	fmtString := "/me/drives/%s/bundles/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Bundle ID
func (id MeDriveIdBundleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("bundles", "bundles", "bundles"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
	}
}

// String returns a human-readable description of this Me Drive Id Bundle ID
func (id MeDriveIdBundleId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
	}
	return fmt.Sprintf("Me Drive Id Bundle (%s)", strings.Join(components, "\n"))
}
