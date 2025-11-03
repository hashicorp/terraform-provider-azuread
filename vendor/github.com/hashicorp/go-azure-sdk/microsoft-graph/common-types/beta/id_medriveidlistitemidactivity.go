package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListItemIdActivityId{}

// MeDriveIdListItemIdActivityId is a struct representing the Resource ID for a Me Drive Id List Item Id Activity
type MeDriveIdListItemIdActivityId struct {
	DriveId           string
	ListItemId        string
	ItemActivityOLDId string
}

// NewMeDriveIdListItemIdActivityID returns a new MeDriveIdListItemIdActivityId struct
func NewMeDriveIdListItemIdActivityID(driveId string, listItemId string, itemActivityOLDId string) MeDriveIdListItemIdActivityId {
	return MeDriveIdListItemIdActivityId{
		DriveId:           driveId,
		ListItemId:        listItemId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseMeDriveIdListItemIdActivityID parses 'input' into a MeDriveIdListItemIdActivityId
func ParseMeDriveIdListItemIdActivityID(input string) (*MeDriveIdListItemIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListItemIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListItemIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdListItemIdActivityIDInsensitively parses 'input' case-insensitively into a MeDriveIdListItemIdActivityId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdListItemIdActivityIDInsensitively(input string) (*MeDriveIdListItemIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListItemIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListItemIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdListItemIdActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ListItemId, ok = input.Parsed["listItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemId", input)
	}

	if id.ItemActivityOLDId, ok = input.Parsed["itemActivityOLDId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityOLDId", input)
	}

	return nil
}

// ValidateMeDriveIdListItemIdActivityID checks that 'input' can be parsed as a Me Drive Id List Item Id Activity ID
func ValidateMeDriveIdListItemIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdListItemIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id List Item Id Activity ID
func (id MeDriveIdListItemIdActivityId) ID() string {
	fmtString := "/me/drives/%s/list/items/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ListItemId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id List Item Id Activity ID
func (id MeDriveIdListItemIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this Me Drive Id List Item Id Activity ID
func (id MeDriveIdListItemIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("Item Activity OLD: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("Me Drive Id List Item Id Activity (%s)", strings.Join(components, "\n"))
}
