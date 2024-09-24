package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnlineMeetingId{}

// MeOnlineMeetingId is a struct representing the Resource ID for a Me Online Meeting
type MeOnlineMeetingId struct {
	OnlineMeetingId string
}

// NewMeOnlineMeetingID returns a new MeOnlineMeetingId struct
func NewMeOnlineMeetingID(onlineMeetingId string) MeOnlineMeetingId {
	return MeOnlineMeetingId{
		OnlineMeetingId: onlineMeetingId,
	}
}

// ParseMeOnlineMeetingID parses 'input' into a MeOnlineMeetingId
func ParseMeOnlineMeetingID(input string) (*MeOnlineMeetingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnlineMeetingIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingIDInsensitively(input string) (*MeOnlineMeetingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnlineMeetingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	return nil
}

// ValidateMeOnlineMeetingID checks that 'input' can be parsed as a Me Online Meeting ID
func ValidateMeOnlineMeetingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting ID
func (id MeOnlineMeetingId) ID() string {
	fmtString := "/me/onlineMeetings/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting ID
func (id MeOnlineMeetingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
	}
}

// String returns a human-readable description of this Me Online Meeting ID
func (id MeOnlineMeetingId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
	}
	return fmt.Sprintf("Me Online Meeting (%s)", strings.Join(components, "\n"))
}
