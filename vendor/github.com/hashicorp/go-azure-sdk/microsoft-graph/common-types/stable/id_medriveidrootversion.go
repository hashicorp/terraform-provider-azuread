package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootVersionId{}

// MeDriveIdRootVersionId is a struct representing the Resource ID for a Me Drive Id Root Version
type MeDriveIdRootVersionId struct {
	DriveId            string
	DriveItemVersionId string
}

// NewMeDriveIdRootVersionID returns a new MeDriveIdRootVersionId struct
func NewMeDriveIdRootVersionID(driveId string, driveItemVersionId string) MeDriveIdRootVersionId {
	return MeDriveIdRootVersionId{
		DriveId:            driveId,
		DriveItemVersionId: driveItemVersionId,
	}
}

// ParseMeDriveIdRootVersionID parses 'input' into a MeDriveIdRootVersionId
func ParseMeDriveIdRootVersionID(input string) (*MeDriveIdRootVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdRootVersionIDInsensitively parses 'input' case-insensitively into a MeDriveIdRootVersionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdRootVersionIDInsensitively(input string) (*MeDriveIdRootVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdRootVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemVersionId, ok = input.Parsed["driveItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemVersionId", input)
	}

	return nil
}

// ValidateMeDriveIdRootVersionID checks that 'input' can be parsed as a Me Drive Id Root Version ID
func ValidateMeDriveIdRootVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdRootVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Root Version ID
func (id MeDriveIdRootVersionId) ID() string {
	fmtString := "/me/drives/%s/root/versions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Root Version ID
func (id MeDriveIdRootVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("driveItemVersionId", "driveItemVersionId"),
	}
}

// String returns a human-readable description of this Me Drive Id Root Version ID
func (id MeDriveIdRootVersionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item Version: %q", id.DriveItemVersionId),
	}
	return fmt.Sprintf("Me Drive Id Root Version (%s)", strings.Join(components, "\n"))
}
