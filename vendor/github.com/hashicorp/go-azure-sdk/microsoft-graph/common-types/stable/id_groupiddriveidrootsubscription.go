package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootSubscriptionId{}

// GroupIdDriveIdRootSubscriptionId is a struct representing the Resource ID for a Group Id Drive Id Root Subscription
type GroupIdDriveIdRootSubscriptionId struct {
	GroupId        string
	DriveId        string
	SubscriptionId string
}

// NewGroupIdDriveIdRootSubscriptionID returns a new GroupIdDriveIdRootSubscriptionId struct
func NewGroupIdDriveIdRootSubscriptionID(groupId string, driveId string, subscriptionId string) GroupIdDriveIdRootSubscriptionId {
	return GroupIdDriveIdRootSubscriptionId{
		GroupId:        groupId,
		DriveId:        driveId,
		SubscriptionId: subscriptionId,
	}
}

// ParseGroupIdDriveIdRootSubscriptionID parses 'input' into a GroupIdDriveIdRootSubscriptionId
func ParseGroupIdDriveIdRootSubscriptionID(input string) (*GroupIdDriveIdRootSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdRootSubscriptionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdRootSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdRootSubscriptionIDInsensitively(input string) (*GroupIdDriveIdRootSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdRootSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdDriveIdRootSubscriptionID checks that 'input' can be parsed as a Group Id Drive Id Root Subscription ID
func ValidateGroupIdDriveIdRootSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdRootSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Root Subscription ID
func (id GroupIdDriveIdRootSubscriptionId) ID() string {
	fmtString := "/groups/%s/drives/%s/root/subscriptions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.SubscriptionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Root Subscription ID
func (id GroupIdDriveIdRootSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("subscriptions", "subscriptions", "subscriptions"),
		resourceids.UserSpecifiedSegment("subscriptionId", "subscriptionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Root Subscription ID
func (id GroupIdDriveIdRootSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
	}
	return fmt.Sprintf("Group Id Drive Id Root Subscription (%s)", strings.Join(components, "\n"))
}
