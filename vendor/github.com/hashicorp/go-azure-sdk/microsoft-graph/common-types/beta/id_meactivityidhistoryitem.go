package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeActivityIdHistoryItemId{}

// MeActivityIdHistoryItemId is a struct representing the Resource ID for a Me Activity Id History Item
type MeActivityIdHistoryItemId struct {
	UserActivityId        string
	ActivityHistoryItemId string
}

// NewMeActivityIdHistoryItemID returns a new MeActivityIdHistoryItemId struct
func NewMeActivityIdHistoryItemID(userActivityId string, activityHistoryItemId string) MeActivityIdHistoryItemId {
	return MeActivityIdHistoryItemId{
		UserActivityId:        userActivityId,
		ActivityHistoryItemId: activityHistoryItemId,
	}
}

// ParseMeActivityIdHistoryItemID parses 'input' into a MeActivityIdHistoryItemId
func ParseMeActivityIdHistoryItemID(input string) (*MeActivityIdHistoryItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeActivityIdHistoryItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeActivityIdHistoryItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeActivityIdHistoryItemIDInsensitively parses 'input' case-insensitively into a MeActivityIdHistoryItemId
// note: this method should only be used for API response data and not user input
func ParseMeActivityIdHistoryItemIDInsensitively(input string) (*MeActivityIdHistoryItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeActivityIdHistoryItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeActivityIdHistoryItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeActivityIdHistoryItemId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserActivityId, ok = input.Parsed["userActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userActivityId", input)
	}

	if id.ActivityHistoryItemId, ok = input.Parsed["activityHistoryItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "activityHistoryItemId", input)
	}

	return nil
}

// ValidateMeActivityIdHistoryItemID checks that 'input' can be parsed as a Me Activity Id History Item ID
func ValidateMeActivityIdHistoryItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeActivityIdHistoryItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Activity Id History Item ID
func (id MeActivityIdHistoryItemId) ID() string {
	fmtString := "/me/activities/%s/historyItems/%s"
	return fmt.Sprintf(fmtString, id.UserActivityId, id.ActivityHistoryItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Activity Id History Item ID
func (id MeActivityIdHistoryItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("userActivityId", "userActivityId"),
		resourceids.StaticSegment("historyItems", "historyItems", "historyItems"),
		resourceids.UserSpecifiedSegment("activityHistoryItemId", "activityHistoryItemId"),
	}
}

// String returns a human-readable description of this Me Activity Id History Item ID
func (id MeActivityIdHistoryItemId) String() string {
	components := []string{
		fmt.Sprintf("User Activity: %q", id.UserActivityId),
		fmt.Sprintf("Activity History Item: %q", id.ActivityHistoryItemId),
	}
	return fmt.Sprintf("Me Activity Id History Item (%s)", strings.Join(components, "\n"))
}
