package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnlineMeetingIdRecordingId{}

// MeOnlineMeetingIdRecordingId is a struct representing the Resource ID for a Me Online Meeting Id Recording
type MeOnlineMeetingIdRecordingId struct {
	OnlineMeetingId string
	CallRecordingId string
}

// NewMeOnlineMeetingIdRecordingID returns a new MeOnlineMeetingIdRecordingId struct
func NewMeOnlineMeetingIdRecordingID(onlineMeetingId string, callRecordingId string) MeOnlineMeetingIdRecordingId {
	return MeOnlineMeetingIdRecordingId{
		OnlineMeetingId: onlineMeetingId,
		CallRecordingId: callRecordingId,
	}
}

// ParseMeOnlineMeetingIdRecordingID parses 'input' into a MeOnlineMeetingIdRecordingId
func ParseMeOnlineMeetingIdRecordingID(input string) (*MeOnlineMeetingIdRecordingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdRecordingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdRecordingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnlineMeetingIdRecordingIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingIdRecordingId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingIdRecordingIDInsensitively(input string) (*MeOnlineMeetingIdRecordingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdRecordingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdRecordingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnlineMeetingIdRecordingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	if id.CallRecordingId, ok = input.Parsed["callRecordingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "callRecordingId", input)
	}

	return nil
}

// ValidateMeOnlineMeetingIdRecordingID checks that 'input' can be parsed as a Me Online Meeting Id Recording ID
func ValidateMeOnlineMeetingIdRecordingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingIdRecordingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting Id Recording ID
func (id MeOnlineMeetingIdRecordingId) ID() string {
	fmtString := "/me/onlineMeetings/%s/recordings/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId, id.CallRecordingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting Id Recording ID
func (id MeOnlineMeetingIdRecordingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("recordings", "recordings", "recordings"),
		resourceids.UserSpecifiedSegment("callRecordingId", "callRecordingId"),
	}
}

// String returns a human-readable description of this Me Online Meeting Id Recording ID
func (id MeOnlineMeetingIdRecordingId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Call Recording: %q", id.CallRecordingId),
	}
	return fmt.Sprintf("Me Online Meeting Id Recording (%s)", strings.Join(components, "\n"))
}
