package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdActivityId{}

// MeDriveIdItemIdActivityId is a struct representing the Resource ID for a Me Drive Id Item Id Activity
type MeDriveIdItemIdActivityId struct {
	DriveId           string
	DriveItemId       string
	ItemActivityOLDId string
}

// NewMeDriveIdItemIdActivityID returns a new MeDriveIdItemIdActivityId struct
func NewMeDriveIdItemIdActivityID(driveId string, driveItemId string, itemActivityOLDId string) MeDriveIdItemIdActivityId {
	return MeDriveIdItemIdActivityId{
		DriveId:           driveId,
		DriveItemId:       driveItemId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseMeDriveIdItemIdActivityID parses 'input' into a MeDriveIdItemIdActivityId
func ParseMeDriveIdItemIdActivityID(input string) (*MeDriveIdItemIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdItemIdActivityIDInsensitively parses 'input' case-insensitively into a MeDriveIdItemIdActivityId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdItemIdActivityIDInsensitively(input string) (*MeDriveIdItemIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdItemIdActivityId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateMeDriveIdItemIdActivityID checks that 'input' can be parsed as a Me Drive Id Item Id Activity ID
func ValidateMeDriveIdItemIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdItemIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Item Id Activity ID
func (id MeDriveIdItemIdActivityId) ID() string {
	fmtString := "/me/drives/%s/items/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Item Id Activity ID
func (id MeDriveIdItemIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this Me Drive Id Item Id Activity ID
func (id MeDriveIdItemIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Item Activity OLD: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("Me Drive Id Item Id Activity (%s)", strings.Join(components, "\n"))
}
