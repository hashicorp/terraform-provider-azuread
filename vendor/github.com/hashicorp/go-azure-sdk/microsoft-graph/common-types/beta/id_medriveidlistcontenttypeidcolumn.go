package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListContentTypeIdColumnId{}

// MeDriveIdListContentTypeIdColumnId is a struct representing the Resource ID for a Me Drive Id List Content Type Id Column
type MeDriveIdListContentTypeIdColumnId struct {
	DriveId            string
	ContentTypeId      string
	ColumnDefinitionId string
}

// NewMeDriveIdListContentTypeIdColumnID returns a new MeDriveIdListContentTypeIdColumnId struct
func NewMeDriveIdListContentTypeIdColumnID(driveId string, contentTypeId string, columnDefinitionId string) MeDriveIdListContentTypeIdColumnId {
	return MeDriveIdListContentTypeIdColumnId{
		DriveId:            driveId,
		ContentTypeId:      contentTypeId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseMeDriveIdListContentTypeIdColumnID parses 'input' into a MeDriveIdListContentTypeIdColumnId
func ParseMeDriveIdListContentTypeIdColumnID(input string) (*MeDriveIdListContentTypeIdColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListContentTypeIdColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListContentTypeIdColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdListContentTypeIdColumnIDInsensitively parses 'input' case-insensitively into a MeDriveIdListContentTypeIdColumnId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdListContentTypeIdColumnIDInsensitively(input string) (*MeDriveIdListContentTypeIdColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListContentTypeIdColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListContentTypeIdColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdListContentTypeIdColumnId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ContentTypeId, ok = input.Parsed["contentTypeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", input)
	}

	if id.ColumnDefinitionId, ok = input.Parsed["columnDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", input)
	}

	return nil
}

// ValidateMeDriveIdListContentTypeIdColumnID checks that 'input' can be parsed as a Me Drive Id List Content Type Id Column ID
func ValidateMeDriveIdListContentTypeIdColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdListContentTypeIdColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id List Content Type Id Column ID
func (id MeDriveIdListContentTypeIdColumnId) ID() string {
	fmtString := "/me/drives/%s/list/contentTypes/%s/columns/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ContentTypeId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id List Content Type Id Column ID
func (id MeDriveIdListContentTypeIdColumnId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("columns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionId"),
	}
}

// String returns a human-readable description of this Me Drive Id List Content Type Id Column ID
func (id MeDriveIdListContentTypeIdColumnId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("Me Drive Id List Content Type Id Column (%s)", strings.Join(components, "\n"))
}
