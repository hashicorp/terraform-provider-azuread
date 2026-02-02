package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdSubscriptionId{}

// MeDriveIdItemIdSubscriptionId is a struct representing the Resource ID for a Me Drive Id Item Id Subscription
type MeDriveIdItemIdSubscriptionId struct {
	DriveId        string
	DriveItemId    string
	SubscriptionId string
}

// NewMeDriveIdItemIdSubscriptionID returns a new MeDriveIdItemIdSubscriptionId struct
func NewMeDriveIdItemIdSubscriptionID(driveId string, driveItemId string, subscriptionId string) MeDriveIdItemIdSubscriptionId {
	return MeDriveIdItemIdSubscriptionId{
		DriveId:        driveId,
		DriveItemId:    driveItemId,
		SubscriptionId: subscriptionId,
	}
}

// ParseMeDriveIdItemIdSubscriptionID parses 'input' into a MeDriveIdItemIdSubscriptionId
func ParseMeDriveIdItemIdSubscriptionID(input string) (*MeDriveIdItemIdSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdItemIdSubscriptionIDInsensitively parses 'input' case-insensitively into a MeDriveIdItemIdSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdItemIdSubscriptionIDInsensitively(input string) (*MeDriveIdItemIdSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdItemIdSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	return nil
}

// ValidateMeDriveIdItemIdSubscriptionID checks that 'input' can be parsed as a Me Drive Id Item Id Subscription ID
func ValidateMeDriveIdItemIdSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdItemIdSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Item Id Subscription ID
func (id MeDriveIdItemIdSubscriptionId) ID() string {
	fmtString := "/me/drives/%s/items/%s/subscriptions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId, id.SubscriptionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Item Id Subscription ID
func (id MeDriveIdItemIdSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("subscriptions", "subscriptions", "subscriptions"),
		resourceids.UserSpecifiedSegment("subscriptionId", "subscriptionId"),
	}
}

// String returns a human-readable description of this Me Drive Id Item Id Subscription ID
func (id MeDriveIdItemIdSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
	}
	return fmt.Sprintf("Me Drive Id Item Id Subscription (%s)", strings.Join(components, "\n"))
}
