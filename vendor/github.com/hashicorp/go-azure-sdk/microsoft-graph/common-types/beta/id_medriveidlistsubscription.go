package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListSubscriptionId{}

// MeDriveIdListSubscriptionId is a struct representing the Resource ID for a Me Drive Id List Subscription
type MeDriveIdListSubscriptionId struct {
	DriveId        string
	SubscriptionId string
}

// NewMeDriveIdListSubscriptionID returns a new MeDriveIdListSubscriptionId struct
func NewMeDriveIdListSubscriptionID(driveId string, subscriptionId string) MeDriveIdListSubscriptionId {
	return MeDriveIdListSubscriptionId{
		DriveId:        driveId,
		SubscriptionId: subscriptionId,
	}
}

// ParseMeDriveIdListSubscriptionID parses 'input' into a MeDriveIdListSubscriptionId
func ParseMeDriveIdListSubscriptionID(input string) (*MeDriveIdListSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdListSubscriptionIDInsensitively parses 'input' case-insensitively into a MeDriveIdListSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdListSubscriptionIDInsensitively(input string) (*MeDriveIdListSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdListSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	return nil
}

// ValidateMeDriveIdListSubscriptionID checks that 'input' can be parsed as a Me Drive Id List Subscription ID
func ValidateMeDriveIdListSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdListSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id List Subscription ID
func (id MeDriveIdListSubscriptionId) ID() string {
	fmtString := "/me/drives/%s/list/subscriptions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.SubscriptionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id List Subscription ID
func (id MeDriveIdListSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("subscriptions", "subscriptions", "subscriptions"),
		resourceids.UserSpecifiedSegment("subscriptionId", "subscriptionId"),
	}
}

// String returns a human-readable description of this Me Drive Id List Subscription ID
func (id MeDriveIdListSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
	}
	return fmt.Sprintf("Me Drive Id List Subscription (%s)", strings.Join(components, "\n"))
}
