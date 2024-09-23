package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdListItemVersionId{}

// MeDriveIdItemIdListItemVersionId is a struct representing the Resource ID for a Me Drive Id Item Id List Item Version
type MeDriveIdItemIdListItemVersionId struct {
	DriveId           string
	DriveItemId       string
	ListItemVersionId string
}

// NewMeDriveIdItemIdListItemVersionID returns a new MeDriveIdItemIdListItemVersionId struct
func NewMeDriveIdItemIdListItemVersionID(driveId string, driveItemId string, listItemVersionId string) MeDriveIdItemIdListItemVersionId {
	return MeDriveIdItemIdListItemVersionId{
		DriveId:           driveId,
		DriveItemId:       driveItemId,
		ListItemVersionId: listItemVersionId,
	}
}

// ParseMeDriveIdItemIdListItemVersionID parses 'input' into a MeDriveIdItemIdListItemVersionId
func ParseMeDriveIdItemIdListItemVersionID(input string) (*MeDriveIdItemIdListItemVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdListItemVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdListItemVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdItemIdListItemVersionIDInsensitively parses 'input' case-insensitively into a MeDriveIdItemIdListItemVersionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdItemIdListItemVersionIDInsensitively(input string) (*MeDriveIdItemIdListItemVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdListItemVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdListItemVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdItemIdListItemVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	if id.ListItemVersionId, ok = input.Parsed["listItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemVersionId", input)
	}

	return nil
}

// ValidateMeDriveIdItemIdListItemVersionID checks that 'input' can be parsed as a Me Drive Id Item Id List Item Version ID
func ValidateMeDriveIdItemIdListItemVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdItemIdListItemVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Item Id List Item Version ID
func (id MeDriveIdItemIdListItemVersionId) ID() string {
	fmtString := "/me/drives/%s/items/%s/listItem/versions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId, id.ListItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Item Id List Item Version ID
func (id MeDriveIdItemIdListItemVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("listItemVersionId", "listItemVersionId"),
	}
}

// String returns a human-readable description of this Me Drive Id Item Id List Item Version ID
func (id MeDriveIdItemIdListItemVersionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("List Item Version: %q", id.ListItemVersionId),
	}
	return fmt.Sprintf("Me Drive Id Item Id List Item Version (%s)", strings.Join(components, "\n"))
}
