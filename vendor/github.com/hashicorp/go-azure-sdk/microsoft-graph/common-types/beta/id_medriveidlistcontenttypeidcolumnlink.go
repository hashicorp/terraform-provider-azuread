package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListContentTypeIdColumnLinkId{}

// MeDriveIdListContentTypeIdColumnLinkId is a struct representing the Resource ID for a Me Drive Id List Content Type Id Column Link
type MeDriveIdListContentTypeIdColumnLinkId struct {
	DriveId       string
	ContentTypeId string
	ColumnLinkId  string
}

// NewMeDriveIdListContentTypeIdColumnLinkID returns a new MeDriveIdListContentTypeIdColumnLinkId struct
func NewMeDriveIdListContentTypeIdColumnLinkID(driveId string, contentTypeId string, columnLinkId string) MeDriveIdListContentTypeIdColumnLinkId {
	return MeDriveIdListContentTypeIdColumnLinkId{
		DriveId:       driveId,
		ContentTypeId: contentTypeId,
		ColumnLinkId:  columnLinkId,
	}
}

// ParseMeDriveIdListContentTypeIdColumnLinkID parses 'input' into a MeDriveIdListContentTypeIdColumnLinkId
func ParseMeDriveIdListContentTypeIdColumnLinkID(input string) (*MeDriveIdListContentTypeIdColumnLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListContentTypeIdColumnLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListContentTypeIdColumnLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdListContentTypeIdColumnLinkIDInsensitively parses 'input' case-insensitively into a MeDriveIdListContentTypeIdColumnLinkId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdListContentTypeIdColumnLinkIDInsensitively(input string) (*MeDriveIdListContentTypeIdColumnLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListContentTypeIdColumnLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListContentTypeIdColumnLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdListContentTypeIdColumnLinkId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ContentTypeId, ok = input.Parsed["contentTypeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", input)
	}

	if id.ColumnLinkId, ok = input.Parsed["columnLinkId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnLinkId", input)
	}

	return nil
}

// ValidateMeDriveIdListContentTypeIdColumnLinkID checks that 'input' can be parsed as a Me Drive Id List Content Type Id Column Link ID
func ValidateMeDriveIdListContentTypeIdColumnLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdListContentTypeIdColumnLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id List Content Type Id Column Link ID
func (id MeDriveIdListContentTypeIdColumnLinkId) ID() string {
	fmtString := "/me/drives/%s/list/contentTypes/%s/columnLinks/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ContentTypeId, id.ColumnLinkId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id List Content Type Id Column Link ID
func (id MeDriveIdListContentTypeIdColumnLinkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("columnLinks", "columnLinks", "columnLinks"),
		resourceids.UserSpecifiedSegment("columnLinkId", "columnLinkId"),
	}
}

// String returns a human-readable description of this Me Drive Id List Content Type Id Column Link ID
func (id MeDriveIdListContentTypeIdColumnLinkId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Link: %q", id.ColumnLinkId),
	}
	return fmt.Sprintf("Me Drive Id List Content Type Id Column Link (%s)", strings.Join(components, "\n"))
}
