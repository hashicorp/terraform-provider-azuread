package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdActivityId{}

// MeDriveIdActivityId is a struct representing the Resource ID for a Me Drive Id Activity
type MeDriveIdActivityId struct {
	DriveId           string
	ItemActivityOLDId string
}

// NewMeDriveIdActivityID returns a new MeDriveIdActivityId struct
func NewMeDriveIdActivityID(driveId string, itemActivityOLDId string) MeDriveIdActivityId {
	return MeDriveIdActivityId{
		DriveId:           driveId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseMeDriveIdActivityID parses 'input' into a MeDriveIdActivityId
func ParseMeDriveIdActivityID(input string) (*MeDriveIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdActivityIDInsensitively parses 'input' case-insensitively into a MeDriveIdActivityId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdActivityIDInsensitively(input string) (*MeDriveIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ItemActivityOLDId, ok = input.Parsed["itemActivityOLDId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityOLDId", input)
	}

	return nil
}

// ValidateMeDriveIdActivityID checks that 'input' can be parsed as a Me Drive Id Activity ID
func ValidateMeDriveIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Activity ID
func (id MeDriveIdActivityId) ID() string {
	fmtString := "/me/drives/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Activity ID
func (id MeDriveIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this Me Drive Id Activity ID
func (id MeDriveIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Item Activity OLD: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("Me Drive Id Activity (%s)", strings.Join(components, "\n"))
}
