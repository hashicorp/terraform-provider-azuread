package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListItemIdVersionId{}

// MeDriveIdListItemIdVersionId is a struct representing the Resource ID for a Me Drive Id List Item Id Version
type MeDriveIdListItemIdVersionId struct {
	DriveId           string
	ListItemId        string
	ListItemVersionId string
}

// NewMeDriveIdListItemIdVersionID returns a new MeDriveIdListItemIdVersionId struct
func NewMeDriveIdListItemIdVersionID(driveId string, listItemId string, listItemVersionId string) MeDriveIdListItemIdVersionId {
	return MeDriveIdListItemIdVersionId{
		DriveId:           driveId,
		ListItemId:        listItemId,
		ListItemVersionId: listItemVersionId,
	}
}

// ParseMeDriveIdListItemIdVersionID parses 'input' into a MeDriveIdListItemIdVersionId
func ParseMeDriveIdListItemIdVersionID(input string) (*MeDriveIdListItemIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListItemIdVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListItemIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdListItemIdVersionIDInsensitively parses 'input' case-insensitively into a MeDriveIdListItemIdVersionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdListItemIdVersionIDInsensitively(input string) (*MeDriveIdListItemIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListItemIdVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListItemIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdListItemIdVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ListItemId, ok = input.Parsed["listItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemId", input)
	}

	if id.ListItemVersionId, ok = input.Parsed["listItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemVersionId", input)
	}

	return nil
}

// ValidateMeDriveIdListItemIdVersionID checks that 'input' can be parsed as a Me Drive Id List Item Id Version ID
func ValidateMeDriveIdListItemIdVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdListItemIdVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id List Item Id Version ID
func (id MeDriveIdListItemIdVersionId) ID() string {
	fmtString := "/me/drives/%s/list/items/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ListItemId, id.ListItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id List Item Id Version ID
func (id MeDriveIdListItemIdVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("listItemVersionId", "listItemVersionId"),
	}
}

// String returns a human-readable description of this Me Drive Id List Item Id Version ID
func (id MeDriveIdListItemIdVersionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("List Item Version: %q", id.ListItemVersionId),
	}
	return fmt.Sprintf("Me Drive Id List Item Id Version (%s)", strings.Join(components, "\n"))
}
