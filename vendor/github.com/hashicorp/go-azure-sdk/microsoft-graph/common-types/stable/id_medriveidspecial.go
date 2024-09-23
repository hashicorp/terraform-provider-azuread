package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdSpecialId{}

// MeDriveIdSpecialId is a struct representing the Resource ID for a Me Drive Id Special
type MeDriveIdSpecialId struct {
	DriveId     string
	DriveItemId string
}

// NewMeDriveIdSpecialID returns a new MeDriveIdSpecialId struct
func NewMeDriveIdSpecialID(driveId string, driveItemId string) MeDriveIdSpecialId {
	return MeDriveIdSpecialId{
		DriveId:     driveId,
		DriveItemId: driveItemId,
	}
}

// ParseMeDriveIdSpecialID parses 'input' into a MeDriveIdSpecialId
func ParseMeDriveIdSpecialID(input string) (*MeDriveIdSpecialId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdSpecialId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdSpecialId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdSpecialIDInsensitively parses 'input' case-insensitively into a MeDriveIdSpecialId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdSpecialIDInsensitively(input string) (*MeDriveIdSpecialId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdSpecialId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdSpecialId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdSpecialId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	return nil
}

// ValidateMeDriveIdSpecialID checks that 'input' can be parsed as a Me Drive Id Special ID
func ValidateMeDriveIdSpecialID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdSpecialID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Special ID
func (id MeDriveIdSpecialId) ID() string {
	fmtString := "/me/drives/%s/special/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Special ID
func (id MeDriveIdSpecialId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("special", "special", "special"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
	}
}

// String returns a human-readable description of this Me Drive Id Special ID
func (id MeDriveIdSpecialId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
	}
	return fmt.Sprintf("Me Drive Id Special (%s)", strings.Join(components, "\n"))
}
