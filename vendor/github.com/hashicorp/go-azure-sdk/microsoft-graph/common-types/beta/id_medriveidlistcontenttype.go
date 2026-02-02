package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListContentTypeId{}

// MeDriveIdListContentTypeId is a struct representing the Resource ID for a Me Drive Id List Content Type
type MeDriveIdListContentTypeId struct {
	DriveId       string
	ContentTypeId string
}

// NewMeDriveIdListContentTypeID returns a new MeDriveIdListContentTypeId struct
func NewMeDriveIdListContentTypeID(driveId string, contentTypeId string) MeDriveIdListContentTypeId {
	return MeDriveIdListContentTypeId{
		DriveId:       driveId,
		ContentTypeId: contentTypeId,
	}
}

// ParseMeDriveIdListContentTypeID parses 'input' into a MeDriveIdListContentTypeId
func ParseMeDriveIdListContentTypeID(input string) (*MeDriveIdListContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListContentTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListContentTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdListContentTypeIDInsensitively parses 'input' case-insensitively into a MeDriveIdListContentTypeId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdListContentTypeIDInsensitively(input string) (*MeDriveIdListContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListContentTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListContentTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdListContentTypeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ContentTypeId, ok = input.Parsed["contentTypeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", input)
	}

	return nil
}

// ValidateMeDriveIdListContentTypeID checks that 'input' can be parsed as a Me Drive Id List Content Type ID
func ValidateMeDriveIdListContentTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdListContentTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id List Content Type ID
func (id MeDriveIdListContentTypeId) ID() string {
	fmtString := "/me/drives/%s/list/contentTypes/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ContentTypeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id List Content Type ID
func (id MeDriveIdListContentTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
	}
}

// String returns a human-readable description of this Me Drive Id List Content Type ID
func (id MeDriveIdListContentTypeId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
	}
	return fmt.Sprintf("Me Drive Id List Content Type (%s)", strings.Join(components, "\n"))
}
