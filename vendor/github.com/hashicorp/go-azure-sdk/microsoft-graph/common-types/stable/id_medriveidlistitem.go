package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListItemId{}

// MeDriveIdListItemId is a struct representing the Resource ID for a Me Drive Id List Item
type MeDriveIdListItemId struct {
	DriveId    string
	ListItemId string
}

// NewMeDriveIdListItemID returns a new MeDriveIdListItemId struct
func NewMeDriveIdListItemID(driveId string, listItemId string) MeDriveIdListItemId {
	return MeDriveIdListItemId{
		DriveId:    driveId,
		ListItemId: listItemId,
	}
}

// ParseMeDriveIdListItemID parses 'input' into a MeDriveIdListItemId
func ParseMeDriveIdListItemID(input string) (*MeDriveIdListItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdListItemIDInsensitively parses 'input' case-insensitively into a MeDriveIdListItemId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdListItemIDInsensitively(input string) (*MeDriveIdListItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdListItemId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ListItemId, ok = input.Parsed["listItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemId", input)
	}

	return nil
}

// ValidateMeDriveIdListItemID checks that 'input' can be parsed as a Me Drive Id List Item ID
func ValidateMeDriveIdListItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdListItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id List Item ID
func (id MeDriveIdListItemId) ID() string {
	fmtString := "/me/drives/%s/list/items/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ListItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id List Item ID
func (id MeDriveIdListItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
	}
}

// String returns a human-readable description of this Me Drive Id List Item ID
func (id MeDriveIdListItemId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
	}
	return fmt.Sprintf("Me Drive Id List Item (%s)", strings.Join(components, "\n"))
}
