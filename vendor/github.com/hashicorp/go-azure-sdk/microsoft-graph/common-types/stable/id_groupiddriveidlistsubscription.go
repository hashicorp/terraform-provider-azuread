package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListSubscriptionId{}

// GroupIdDriveIdListSubscriptionId is a struct representing the Resource ID for a Group Id Drive Id List Subscription
type GroupIdDriveIdListSubscriptionId struct {
	GroupId        string
	DriveId        string
	SubscriptionId string
}

// NewGroupIdDriveIdListSubscriptionID returns a new GroupIdDriveIdListSubscriptionId struct
func NewGroupIdDriveIdListSubscriptionID(groupId string, driveId string, subscriptionId string) GroupIdDriveIdListSubscriptionId {
	return GroupIdDriveIdListSubscriptionId{
		GroupId:        groupId,
		DriveId:        driveId,
		SubscriptionId: subscriptionId,
	}
}

// ParseGroupIdDriveIdListSubscriptionID parses 'input' into a GroupIdDriveIdListSubscriptionId
func ParseGroupIdDriveIdListSubscriptionID(input string) (*GroupIdDriveIdListSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdListSubscriptionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdListSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdListSubscriptionIDInsensitively(input string) (*GroupIdDriveIdListSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdListSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdListSubscriptionID checks that 'input' can be parsed as a Group Id Drive Id List Subscription ID
func ValidateGroupIdDriveIdListSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdListSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id List Subscription ID
func (id GroupIdDriveIdListSubscriptionId) ID() string {
	fmtString := "/groups/%s/drives/%s/list/subscriptions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.SubscriptionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id List Subscription ID
func (id GroupIdDriveIdListSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("subscriptions", "subscriptions", "subscriptions"),
		resourceids.UserSpecifiedSegment("subscriptionId", "subscriptionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id List Subscription ID
func (id GroupIdDriveIdListSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
	}
	return fmt.Sprintf("Group Id Drive Id List Subscription (%s)", strings.Join(components, "\n"))
}
