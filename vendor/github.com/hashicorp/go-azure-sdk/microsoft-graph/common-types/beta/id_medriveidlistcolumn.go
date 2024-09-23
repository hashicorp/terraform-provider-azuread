package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListColumnId{}

// MeDriveIdListColumnId is a struct representing the Resource ID for a Me Drive Id List Column
type MeDriveIdListColumnId struct {
	DriveId            string
	ColumnDefinitionId string
}

// NewMeDriveIdListColumnID returns a new MeDriveIdListColumnId struct
func NewMeDriveIdListColumnID(driveId string, columnDefinitionId string) MeDriveIdListColumnId {
	return MeDriveIdListColumnId{
		DriveId:            driveId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseMeDriveIdListColumnID parses 'input' into a MeDriveIdListColumnId
func ParseMeDriveIdListColumnID(input string) (*MeDriveIdListColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdListColumnIDInsensitively parses 'input' case-insensitively into a MeDriveIdListColumnId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdListColumnIDInsensitively(input string) (*MeDriveIdListColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdListColumnId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ColumnDefinitionId, ok = input.Parsed["columnDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", input)
	}

	return nil
}

// ValidateMeDriveIdListColumnID checks that 'input' can be parsed as a Me Drive Id List Column ID
func ValidateMeDriveIdListColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdListColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id List Column ID
func (id MeDriveIdListColumnId) ID() string {
	fmtString := "/me/drives/%s/list/columns/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id List Column ID
func (id MeDriveIdListColumnId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("columns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionId"),
	}
}

// String returns a human-readable description of this Me Drive Id List Column ID
func (id MeDriveIdListColumnId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("Me Drive Id List Column (%s)", strings.Join(components, "\n"))
}
