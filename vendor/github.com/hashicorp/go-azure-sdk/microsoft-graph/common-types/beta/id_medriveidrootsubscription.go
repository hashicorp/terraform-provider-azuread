package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootSubscriptionId{}

// MeDriveIdRootSubscriptionId is a struct representing the Resource ID for a Me Drive Id Root Subscription
type MeDriveIdRootSubscriptionId struct {
	DriveId        string
	SubscriptionId string
}

// NewMeDriveIdRootSubscriptionID returns a new MeDriveIdRootSubscriptionId struct
func NewMeDriveIdRootSubscriptionID(driveId string, subscriptionId string) MeDriveIdRootSubscriptionId {
	return MeDriveIdRootSubscriptionId{
		DriveId:        driveId,
		SubscriptionId: subscriptionId,
	}
}

// ParseMeDriveIdRootSubscriptionID parses 'input' into a MeDriveIdRootSubscriptionId
func ParseMeDriveIdRootSubscriptionID(input string) (*MeDriveIdRootSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdRootSubscriptionIDInsensitively parses 'input' case-insensitively into a MeDriveIdRootSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdRootSubscriptionIDInsensitively(input string) (*MeDriveIdRootSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdRootSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	return nil
}

// ValidateMeDriveIdRootSubscriptionID checks that 'input' can be parsed as a Me Drive Id Root Subscription ID
func ValidateMeDriveIdRootSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdRootSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Root Subscription ID
func (id MeDriveIdRootSubscriptionId) ID() string {
	fmtString := "/me/drives/%s/root/subscriptions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.SubscriptionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Root Subscription ID
func (id MeDriveIdRootSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("subscriptions", "subscriptions", "subscriptions"),
		resourceids.UserSpecifiedSegment("subscriptionId", "subscriptionId"),
	}
}

// String returns a human-readable description of this Me Drive Id Root Subscription ID
func (id MeDriveIdRootSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
	}
	return fmt.Sprintf("Me Drive Id Root Subscription (%s)", strings.Join(components, "\n"))
}
