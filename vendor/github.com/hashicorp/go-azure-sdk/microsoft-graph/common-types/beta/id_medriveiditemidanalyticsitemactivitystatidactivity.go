package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId{}

// MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId is a struct representing the Resource ID for a Me Drive Id Item Id Analytics Item Activity Stat Id Activity
type MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId struct {
	DriveId            string
	DriveItemId        string
	ItemActivityStatId string
	ItemActivityId     string
}

// NewMeDriveIdItemIdAnalyticsItemActivityStatIdActivityID returns a new MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId struct
func NewMeDriveIdItemIdAnalyticsItemActivityStatIdActivityID(driveId string, driveItemId string, itemActivityStatId string, itemActivityId string) MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId {
	return MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId{
		DriveId:            driveId,
		DriveItemId:        driveItemId,
		ItemActivityStatId: itemActivityStatId,
		ItemActivityId:     itemActivityId,
	}
}

// ParseMeDriveIdItemIdAnalyticsItemActivityStatIdActivityID parses 'input' into a MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId
func ParseMeDriveIdItemIdAnalyticsItemActivityStatIdActivityID(input string) (*MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdItemIdAnalyticsItemActivityStatIdActivityIDInsensitively parses 'input' case-insensitively into a MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdItemIdAnalyticsItemActivityStatIdActivityIDInsensitively(input string) (*MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ItemActivityId, ok = input.Parsed["itemActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityId", input)
	}

	return nil
}

// ValidateMeDriveIdItemIdAnalyticsItemActivityStatIdActivityID checks that 'input' can be parsed as a Me Drive Id Item Id Analytics Item Activity Stat Id Activity ID
func ValidateMeDriveIdItemIdAnalyticsItemActivityStatIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdItemIdAnalyticsItemActivityStatIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Item Id Analytics Item Activity Stat Id Activity ID
func (id MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId) ID() string {
	fmtString := "/me/drives/%s/items/%s/analytics/itemActivityStats/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId, id.ItemActivityStatId, id.ItemActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Item Id Analytics Item Activity Stat Id Activity ID
func (id MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("itemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityId", "itemActivityId"),
	}
}

// String returns a human-readable description of this Me Drive Id Item Id Analytics Item Activity Stat Id Activity ID
func (id MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
		fmt.Sprintf("Item Activity: %q", id.ItemActivityId),
	}
	return fmt.Sprintf("Me Drive Id Item Id Analytics Item Activity Stat Id Activity (%s)", strings.Join(components, "\n"))
}
