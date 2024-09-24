package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeActivityId{}

// MeActivityId is a struct representing the Resource ID for a Me Activity
type MeActivityId struct {
	UserActivityId string
}

// NewMeActivityID returns a new MeActivityId struct
func NewMeActivityID(userActivityId string) MeActivityId {
	return MeActivityId{
		UserActivityId: userActivityId,
	}
}

// ParseMeActivityID parses 'input' into a MeActivityId
func ParseMeActivityID(input string) (*MeActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeActivityIDInsensitively parses 'input' case-insensitively into a MeActivityId
// note: this method should only be used for API response data and not user input
func ParseMeActivityIDInsensitively(input string) (*MeActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserActivityId, ok = input.Parsed["userActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userActivityId", input)
	}

	return nil
}

// ValidateMeActivityID checks that 'input' can be parsed as a Me Activity ID
func ValidateMeActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Activity ID
func (id MeActivityId) ID() string {
	fmtString := "/me/activities/%s"
	return fmt.Sprintf(fmtString, id.UserActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Activity ID
func (id MeActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("userActivityId", "userActivityId"),
	}
}

// String returns a human-readable description of this Me Activity ID
func (id MeActivityId) String() string {
	components := []string{
		fmt.Sprintf("User Activity: %q", id.UserActivityId),
	}
	return fmt.Sprintf("Me Activity (%s)", strings.Join(components, "\n"))
}
