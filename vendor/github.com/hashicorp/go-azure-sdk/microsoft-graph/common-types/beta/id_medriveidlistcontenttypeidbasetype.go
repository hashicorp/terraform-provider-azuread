package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListContentTypeIdBaseTypeId{}

// MeDriveIdListContentTypeIdBaseTypeId is a struct representing the Resource ID for a Me Drive Id List Content Type Id Base Type
type MeDriveIdListContentTypeIdBaseTypeId struct {
	DriveId        string
	ContentTypeId  string
	ContentTypeId1 string
}

// NewMeDriveIdListContentTypeIdBaseTypeID returns a new MeDriveIdListContentTypeIdBaseTypeId struct
func NewMeDriveIdListContentTypeIdBaseTypeID(driveId string, contentTypeId string, contentTypeId1 string) MeDriveIdListContentTypeIdBaseTypeId {
	return MeDriveIdListContentTypeIdBaseTypeId{
		DriveId:        driveId,
		ContentTypeId:  contentTypeId,
		ContentTypeId1: contentTypeId1,
	}
}

// ParseMeDriveIdListContentTypeIdBaseTypeID parses 'input' into a MeDriveIdListContentTypeIdBaseTypeId
func ParseMeDriveIdListContentTypeIdBaseTypeID(input string) (*MeDriveIdListContentTypeIdBaseTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListContentTypeIdBaseTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListContentTypeIdBaseTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdListContentTypeIdBaseTypeIDInsensitively parses 'input' case-insensitively into a MeDriveIdListContentTypeIdBaseTypeId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdListContentTypeIdBaseTypeIDInsensitively(input string) (*MeDriveIdListContentTypeIdBaseTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListContentTypeIdBaseTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListContentTypeIdBaseTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdListContentTypeIdBaseTypeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ContentTypeId, ok = input.Parsed["contentTypeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", input)
	}

	if id.ContentTypeId1, ok = input.Parsed["contentTypeId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId1", input)
	}

	return nil
}

// ValidateMeDriveIdListContentTypeIdBaseTypeID checks that 'input' can be parsed as a Me Drive Id List Content Type Id Base Type ID
func ValidateMeDriveIdListContentTypeIdBaseTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdListContentTypeIdBaseTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id List Content Type Id Base Type ID
func (id MeDriveIdListContentTypeIdBaseTypeId) ID() string {
	fmtString := "/me/drives/%s/list/contentTypes/%s/baseTypes/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ContentTypeId, id.ContentTypeId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id List Content Type Id Base Type ID
func (id MeDriveIdListContentTypeIdBaseTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("baseTypes", "baseTypes", "baseTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId1", "contentTypeId1"),
	}
}

// String returns a human-readable description of this Me Drive Id List Content Type Id Base Type ID
func (id MeDriveIdListContentTypeIdBaseTypeId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Content Type Id 1: %q", id.ContentTypeId1),
	}
	return fmt.Sprintf("Me Drive Id List Content Type Id Base Type (%s)", strings.Join(components, "\n"))
}
