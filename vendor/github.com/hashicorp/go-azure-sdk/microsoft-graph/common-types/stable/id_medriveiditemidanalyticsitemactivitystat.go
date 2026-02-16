package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdAnalyticsItemActivityStatId{}

// MeDriveIdItemIdAnalyticsItemActivityStatId is a struct representing the Resource ID for a Me Drive Id Item Id Analytics Item Activity Stat
type MeDriveIdItemIdAnalyticsItemActivityStatId struct {
	DriveId            string
	DriveItemId        string
	ItemActivityStatId string
}

// NewMeDriveIdItemIdAnalyticsItemActivityStatID returns a new MeDriveIdItemIdAnalyticsItemActivityStatId struct
func NewMeDriveIdItemIdAnalyticsItemActivityStatID(driveId string, driveItemId string, itemActivityStatId string) MeDriveIdItemIdAnalyticsItemActivityStatId {
	return MeDriveIdItemIdAnalyticsItemActivityStatId{
		DriveId:            driveId,
		DriveItemId:        driveItemId,
		ItemActivityStatId: itemActivityStatId,
	}
}

// ParseMeDriveIdItemIdAnalyticsItemActivityStatID parses 'input' into a MeDriveIdItemIdAnalyticsItemActivityStatId
func ParseMeDriveIdItemIdAnalyticsItemActivityStatID(input string) (*MeDriveIdItemIdAnalyticsItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdAnalyticsItemActivityStatId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdAnalyticsItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdItemIdAnalyticsItemActivityStatIDInsensitively parses 'input' case-insensitively into a MeDriveIdItemIdAnalyticsItemActivityStatId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdItemIdAnalyticsItemActivityStatIDInsensitively(input string) (*MeDriveIdItemIdAnalyticsItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdAnalyticsItemActivityStatId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdAnalyticsItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdItemIdAnalyticsItemActivityStatId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	if id.ItemActivityStatId, ok = input.Parsed["itemActivityStatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityStatId", input)
	}

	return nil
}

// ValidateMeDriveIdItemIdAnalyticsItemActivityStatID checks that 'input' can be parsed as a Me Drive Id Item Id Analytics Item Activity Stat ID
func ValidateMeDriveIdItemIdAnalyticsItemActivityStatID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdItemIdAnalyticsItemActivityStatID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Item Id Analytics Item Activity Stat ID
func (id MeDriveIdItemIdAnalyticsItemActivityStatId) ID() string {
	fmtString := "/me/drives/%s/items/%s/analytics/itemActivityStats/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId, id.ItemActivityStatId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Item Id Analytics Item Activity Stat ID
func (id MeDriveIdItemIdAnalyticsItemActivityStatId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("itemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatId"),
	}
}

// String returns a human-readable description of this Me Drive Id Item Id Analytics Item Activity Stat ID
func (id MeDriveIdItemIdAnalyticsItemActivityStatId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
	}
	return fmt.Sprintf("Me Drive Id Item Id Analytics Item Activity Stat (%s)", strings.Join(components, "\n"))
}
