package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveId{}

// MeDriveId is a struct representing the Resource ID for a Me Drive
type MeDriveId struct {
	DriveId string
}

// NewMeDriveID returns a new MeDriveId struct
func NewMeDriveID(driveId string) MeDriveId {
	return MeDriveId{
		DriveId: driveId,
	}
}

// ParseMeDriveID parses 'input' into a MeDriveId
func ParseMeDriveID(input string) (*MeDriveId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIDInsensitively parses 'input' case-insensitively into a MeDriveId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIDInsensitively(input string) (*MeDriveId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	return nil
}

// ValidateMeDriveID checks that 'input' can be parsed as a Me Drive ID
func ValidateMeDriveID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive ID
func (id MeDriveId) ID() string {
	fmtString := "/me/drives/%s"
	return fmt.Sprintf(fmtString, id.DriveId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive ID
func (id MeDriveId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
	}
}

// String returns a human-readable description of this Me Drive ID
func (id MeDriveId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
	}
	return fmt.Sprintf("Me Drive (%s)", strings.Join(components, "\n"))
}
