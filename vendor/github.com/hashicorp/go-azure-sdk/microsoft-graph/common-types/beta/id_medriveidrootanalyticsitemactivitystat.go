package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootAnalyticsItemActivityStatId{}

// MeDriveIdRootAnalyticsItemActivityStatId is a struct representing the Resource ID for a Me Drive Id Root Analytics Item Activity Stat
type MeDriveIdRootAnalyticsItemActivityStatId struct {
	DriveId            string
	ItemActivityStatId string
}

// NewMeDriveIdRootAnalyticsItemActivityStatID returns a new MeDriveIdRootAnalyticsItemActivityStatId struct
func NewMeDriveIdRootAnalyticsItemActivityStatID(driveId string, itemActivityStatId string) MeDriveIdRootAnalyticsItemActivityStatId {
	return MeDriveIdRootAnalyticsItemActivityStatId{
		DriveId:            driveId,
		ItemActivityStatId: itemActivityStatId,
	}
}

// ParseMeDriveIdRootAnalyticsItemActivityStatID parses 'input' into a MeDriveIdRootAnalyticsItemActivityStatId
func ParseMeDriveIdRootAnalyticsItemActivityStatID(input string) (*MeDriveIdRootAnalyticsItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootAnalyticsItemActivityStatId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootAnalyticsItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdRootAnalyticsItemActivityStatIDInsensitively parses 'input' case-insensitively into a MeDriveIdRootAnalyticsItemActivityStatId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdRootAnalyticsItemActivityStatIDInsensitively(input string) (*MeDriveIdRootAnalyticsItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootAnalyticsItemActivityStatId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootAnalyticsItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdRootAnalyticsItemActivityStatId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ItemActivityStatId, ok = input.Parsed["itemActivityStatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityStatId", input)
	}

	return nil
}

// ValidateMeDriveIdRootAnalyticsItemActivityStatID checks that 'input' can be parsed as a Me Drive Id Root Analytics Item Activity Stat ID
func ValidateMeDriveIdRootAnalyticsItemActivityStatID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdRootAnalyticsItemActivityStatID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Root Analytics Item Activity Stat ID
func (id MeDriveIdRootAnalyticsItemActivityStatId) ID() string {
	fmtString := "/me/drives/%s/root/analytics/itemActivityStats/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ItemActivityStatId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Root Analytics Item Activity Stat ID
func (id MeDriveIdRootAnalyticsItemActivityStatId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("itemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatId"),
	}
}

// String returns a human-readable description of this Me Drive Id Root Analytics Item Activity Stat ID
func (id MeDriveIdRootAnalyticsItemActivityStatId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
	}
	return fmt.Sprintf("Me Drive Id Root Analytics Item Activity Stat (%s)", strings.Join(components, "\n"))
}
