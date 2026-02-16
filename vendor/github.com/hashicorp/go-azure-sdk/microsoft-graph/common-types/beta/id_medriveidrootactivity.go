package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootActivityId{}

// MeDriveIdRootActivityId is a struct representing the Resource ID for a Me Drive Id Root Activity
type MeDriveIdRootActivityId struct {
	DriveId           string
	ItemActivityOLDId string
}

// NewMeDriveIdRootActivityID returns a new MeDriveIdRootActivityId struct
func NewMeDriveIdRootActivityID(driveId string, itemActivityOLDId string) MeDriveIdRootActivityId {
	return MeDriveIdRootActivityId{
		DriveId:           driveId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseMeDriveIdRootActivityID parses 'input' into a MeDriveIdRootActivityId
func ParseMeDriveIdRootActivityID(input string) (*MeDriveIdRootActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdRootActivityIDInsensitively parses 'input' case-insensitively into a MeDriveIdRootActivityId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdRootActivityIDInsensitively(input string) (*MeDriveIdRootActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdRootActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ItemActivityOLDId, ok = input.Parsed["itemActivityOLDId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityOLDId", input)
	}

	return nil
}

// ValidateMeDriveIdRootActivityID checks that 'input' can be parsed as a Me Drive Id Root Activity ID
func ValidateMeDriveIdRootActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdRootActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Root Activity ID
func (id MeDriveIdRootActivityId) ID() string {
	fmtString := "/me/drives/%s/root/activities/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Root Activity ID
func (id MeDriveIdRootActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this Me Drive Id Root Activity ID
func (id MeDriveIdRootActivityId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Item Activity OLD: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("Me Drive Id Root Activity (%s)", strings.Join(components, "\n"))
}
