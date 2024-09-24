package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootAnalyticsItemActivityStatIdActivityId{}

// MeDriveIdRootAnalyticsItemActivityStatIdActivityId is a struct representing the Resource ID for a Me Drive Id Root Analytics Item Activity Stat Id Activity
type MeDriveIdRootAnalyticsItemActivityStatIdActivityId struct {
	DriveId            string
	ItemActivityStatId string
	ItemActivityId     string
}

// NewMeDriveIdRootAnalyticsItemActivityStatIdActivityID returns a new MeDriveIdRootAnalyticsItemActivityStatIdActivityId struct
func NewMeDriveIdRootAnalyticsItemActivityStatIdActivityID(driveId string, itemActivityStatId string, itemActivityId string) MeDriveIdRootAnalyticsItemActivityStatIdActivityId {
	return MeDriveIdRootAnalyticsItemActivityStatIdActivityId{
		DriveId:            driveId,
		ItemActivityStatId: itemActivityStatId,
		ItemActivityId:     itemActivityId,
	}
}

// ParseMeDriveIdRootAnalyticsItemActivityStatIdActivityID parses 'input' into a MeDriveIdRootAnalyticsItemActivityStatIdActivityId
func ParseMeDriveIdRootAnalyticsItemActivityStatIdActivityID(input string) (*MeDriveIdRootAnalyticsItemActivityStatIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootAnalyticsItemActivityStatIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootAnalyticsItemActivityStatIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdRootAnalyticsItemActivityStatIdActivityIDInsensitively parses 'input' case-insensitively into a MeDriveIdRootAnalyticsItemActivityStatIdActivityId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdRootAnalyticsItemActivityStatIdActivityIDInsensitively(input string) (*MeDriveIdRootAnalyticsItemActivityStatIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootAnalyticsItemActivityStatIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootAnalyticsItemActivityStatIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdRootAnalyticsItemActivityStatIdActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ItemActivityStatId, ok = input.Parsed["itemActivityStatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityStatId", input)
	}

	if id.ItemActivityId, ok = input.Parsed["itemActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityId", input)
	}

	return nil
}

// ValidateMeDriveIdRootAnalyticsItemActivityStatIdActivityID checks that 'input' can be parsed as a Me Drive Id Root Analytics Item Activity Stat Id Activity ID
func ValidateMeDriveIdRootAnalyticsItemActivityStatIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdRootAnalyticsItemActivityStatIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Root Analytics Item Activity Stat Id Activity ID
func (id MeDriveIdRootAnalyticsItemActivityStatIdActivityId) ID() string {
	fmtString := "/me/drives/%s/root/analytics/itemActivityStats/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ItemActivityStatId, id.ItemActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Root Analytics Item Activity Stat Id Activity ID
func (id MeDriveIdRootAnalyticsItemActivityStatIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("itemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityId", "itemActivityId"),
	}
}

// String returns a human-readable description of this Me Drive Id Root Analytics Item Activity Stat Id Activity ID
func (id MeDriveIdRootAnalyticsItemActivityStatIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
		fmt.Sprintf("Item Activity: %q", id.ItemActivityId),
	}
	return fmt.Sprintf("Me Drive Id Root Analytics Item Activity Stat Id Activity (%s)", strings.Join(components, "\n"))
}
