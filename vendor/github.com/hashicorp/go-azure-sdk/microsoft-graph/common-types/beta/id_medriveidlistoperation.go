package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListOperationId{}

// MeDriveIdListOperationId is a struct representing the Resource ID for a Me Drive Id List Operation
type MeDriveIdListOperationId struct {
	DriveId                    string
	RichLongRunningOperationId string
}

// NewMeDriveIdListOperationID returns a new MeDriveIdListOperationId struct
func NewMeDriveIdListOperationID(driveId string, richLongRunningOperationId string) MeDriveIdListOperationId {
	return MeDriveIdListOperationId{
		DriveId:                    driveId,
		RichLongRunningOperationId: richLongRunningOperationId,
	}
}

// ParseMeDriveIdListOperationID parses 'input' into a MeDriveIdListOperationId
func ParseMeDriveIdListOperationID(input string) (*MeDriveIdListOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdListOperationIDInsensitively parses 'input' case-insensitively into a MeDriveIdListOperationId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdListOperationIDInsensitively(input string) (*MeDriveIdListOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdListOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.RichLongRunningOperationId, ok = input.Parsed["richLongRunningOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "richLongRunningOperationId", input)
	}

	return nil
}

// ValidateMeDriveIdListOperationID checks that 'input' can be parsed as a Me Drive Id List Operation ID
func ValidateMeDriveIdListOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdListOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id List Operation ID
func (id MeDriveIdListOperationId) ID() string {
	fmtString := "/me/drives/%s/list/operations/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.RichLongRunningOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id List Operation ID
func (id MeDriveIdListOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("richLongRunningOperationId", "richLongRunningOperationId"),
	}
}

// String returns a human-readable description of this Me Drive Id List Operation ID
func (id MeDriveIdListOperationId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Rich Long Running Operation: %q", id.RichLongRunningOperationId),
	}
	return fmt.Sprintf("Me Drive Id List Operation (%s)", strings.Join(components, "\n"))
}
