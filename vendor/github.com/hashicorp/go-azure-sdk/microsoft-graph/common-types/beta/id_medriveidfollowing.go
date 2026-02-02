package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdFollowingId{}

// MeDriveIdFollowingId is a struct representing the Resource ID for a Me Drive Id Following
type MeDriveIdFollowingId struct {
	DriveId     string
	DriveItemId string
}

// NewMeDriveIdFollowingID returns a new MeDriveIdFollowingId struct
func NewMeDriveIdFollowingID(driveId string, driveItemId string) MeDriveIdFollowingId {
	return MeDriveIdFollowingId{
		DriveId:     driveId,
		DriveItemId: driveItemId,
	}
}

// ParseMeDriveIdFollowingID parses 'input' into a MeDriveIdFollowingId
func ParseMeDriveIdFollowingID(input string) (*MeDriveIdFollowingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdFollowingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdFollowingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdFollowingIDInsensitively parses 'input' case-insensitively into a MeDriveIdFollowingId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdFollowingIDInsensitively(input string) (*MeDriveIdFollowingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdFollowingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdFollowingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdFollowingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	return nil
}

// ValidateMeDriveIdFollowingID checks that 'input' can be parsed as a Me Drive Id Following ID
func ValidateMeDriveIdFollowingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdFollowingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Following ID
func (id MeDriveIdFollowingId) ID() string {
	fmtString := "/me/drives/%s/following/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Following ID
func (id MeDriveIdFollowingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("following", "following", "following"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
	}
}

// String returns a human-readable description of this Me Drive Id Following ID
func (id MeDriveIdFollowingId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
	}
	return fmt.Sprintf("Me Drive Id Following (%s)", strings.Join(components, "\n"))
}
