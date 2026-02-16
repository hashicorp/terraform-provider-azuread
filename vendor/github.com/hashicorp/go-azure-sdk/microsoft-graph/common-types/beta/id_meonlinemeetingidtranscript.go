package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnlineMeetingIdTranscriptId{}

// MeOnlineMeetingIdTranscriptId is a struct representing the Resource ID for a Me Online Meeting Id Transcript
type MeOnlineMeetingIdTranscriptId struct {
	OnlineMeetingId  string
	CallTranscriptId string
}

// NewMeOnlineMeetingIdTranscriptID returns a new MeOnlineMeetingIdTranscriptId struct
func NewMeOnlineMeetingIdTranscriptID(onlineMeetingId string, callTranscriptId string) MeOnlineMeetingIdTranscriptId {
	return MeOnlineMeetingIdTranscriptId{
		OnlineMeetingId:  onlineMeetingId,
		CallTranscriptId: callTranscriptId,
	}
}

// ParseMeOnlineMeetingIdTranscriptID parses 'input' into a MeOnlineMeetingIdTranscriptId
func ParseMeOnlineMeetingIdTranscriptID(input string) (*MeOnlineMeetingIdTranscriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdTranscriptId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdTranscriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnlineMeetingIdTranscriptIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingIdTranscriptId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingIdTranscriptIDInsensitively(input string) (*MeOnlineMeetingIdTranscriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdTranscriptId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdTranscriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnlineMeetingIdTranscriptId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	if id.CallTranscriptId, ok = input.Parsed["callTranscriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "callTranscriptId", input)
	}

	return nil
}

// ValidateMeOnlineMeetingIdTranscriptID checks that 'input' can be parsed as a Me Online Meeting Id Transcript ID
func ValidateMeOnlineMeetingIdTranscriptID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingIdTranscriptID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting Id Transcript ID
func (id MeOnlineMeetingIdTranscriptId) ID() string {
	fmtString := "/me/onlineMeetings/%s/transcripts/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId, id.CallTranscriptId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting Id Transcript ID
func (id MeOnlineMeetingIdTranscriptId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("transcripts", "transcripts", "transcripts"),
		resourceids.UserSpecifiedSegment("callTranscriptId", "callTranscriptId"),
	}
}

// String returns a human-readable description of this Me Online Meeting Id Transcript ID
func (id MeOnlineMeetingIdTranscriptId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Call Transcript: %q", id.CallTranscriptId),
	}
	return fmt.Sprintf("Me Online Meeting Id Transcript (%s)", strings.Join(components, "\n"))
}
