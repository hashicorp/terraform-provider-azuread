package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemIdSubscriptionId{}

// GroupIdDriveIdItemIdSubscriptionId is a struct representing the Resource ID for a Group Id Drive Id Item Id Subscription
type GroupIdDriveIdItemIdSubscriptionId struct {
	GroupId        string
	DriveId        string
	DriveItemId    string
	SubscriptionId string
}

// NewGroupIdDriveIdItemIdSubscriptionID returns a new GroupIdDriveIdItemIdSubscriptionId struct
func NewGroupIdDriveIdItemIdSubscriptionID(groupId string, driveId string, driveItemId string, subscriptionId string) GroupIdDriveIdItemIdSubscriptionId {
	return GroupIdDriveIdItemIdSubscriptionId{
		GroupId:        groupId,
		DriveId:        driveId,
		DriveItemId:    driveItemId,
		SubscriptionId: subscriptionId,
	}
}

// ParseGroupIdDriveIdItemIdSubscriptionID parses 'input' into a GroupIdDriveIdItemIdSubscriptionId
func ParseGroupIdDriveIdItemIdSubscriptionID(input string) (*GroupIdDriveIdItemIdSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdItemIdSubscriptionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdItemIdSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdItemIdSubscriptionIDInsensitively(input string) (*GroupIdDriveIdItemIdSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdItemIdSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

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

// ValidateGroupIdDriveIdItemIdSubscriptionID checks that 'input' can be parsed as a Group Id Drive Id Item Id Subscription ID
func ValidateGroupIdDriveIdItemIdSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdItemIdSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Item Id Subscription ID
func (id GroupIdDriveIdItemIdSubscriptionId) ID() string {
	fmtString := "/groups/%s/drives/%s/items/%s/subscriptions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId, id.SubscriptionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Item Id Subscription ID
func (id GroupIdDriveIdItemIdSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("subscriptions", "subscriptions", "subscriptions"),
		resourceids.UserSpecifiedSegment("subscriptionId", "subscriptionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Item Id Subscription ID
func (id GroupIdDriveIdItemIdSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
	}
	return fmt.Sprintf("Group Id Drive Id Item Id Subscription (%s)", strings.Join(components, "\n"))
}
