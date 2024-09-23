package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdListItemActivityId{}

// MeDriveIdItemIdListItemActivityId is a struct representing the Resource ID for a Me Drive Id Item Id List Item Activity
type MeDriveIdItemIdListItemActivityId struct {
	DriveId           string
	DriveItemId       string
	ItemActivityOLDId string
}

// NewMeDriveIdItemIdListItemActivityID returns a new MeDriveIdItemIdListItemActivityId struct
func NewMeDriveIdItemIdListItemActivityID(driveId string, driveItemId string, itemActivityOLDId string) MeDriveIdItemIdListItemActivityId {
	return MeDriveIdItemIdListItemActivityId{
		DriveId:           driveId,
		DriveItemId:       driveItemId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseMeDriveIdItemIdListItemActivityID parses 'input' into a MeDriveIdItemIdListItemActivityId
func ParseMeDriveIdItemIdListItemActivityID(input string) (*MeDriveIdItemIdListItemActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdListItemActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdListItemActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdItemIdListItemActivityIDInsensitively parses 'input' case-insensitively into a MeDriveIdItemIdListItemActivityId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdItemIdListItemActivityIDInsensitively(input string) (*MeDriveIdItemIdListItemActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdListItemActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdListItemActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdItemIdListItemActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	if id.ItemActivityOLDId, ok = input.Parsed["itemActivityOLDId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityOLDId", input)
	}

	return nil
}

// ValidateMeDriveIdItemIdListItemActivityID checks that 'input' can be parsed as a Me Drive Id Item Id List Item Activity ID
func ValidateMeDriveIdItemIdListItemActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdItemIdListItemActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Item Id List Item Activity ID
func (id MeDriveIdItemIdListItemActivityId) ID() string {
	fmtString := "/me/drives/%s/items/%s/listItem/activities/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Item Id List Item Activity ID
func (id MeDriveIdItemIdListItemActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this Me Drive Id Item Id List Item Activity ID
func (id MeDriveIdItemIdListItemActivityId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Item Activity O L D: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("Me Drive Id Item Id List Item Activity (%s)", strings.Join(components, "\n"))
}
