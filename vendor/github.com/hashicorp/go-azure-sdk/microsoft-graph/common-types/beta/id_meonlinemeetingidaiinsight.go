package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnlineMeetingIdAiInsightId{}

// MeOnlineMeetingIdAiInsightId is a struct representing the Resource ID for a Me Online Meeting Id Ai Insight
type MeOnlineMeetingIdAiInsightId struct {
	OnlineMeetingId string
	CallAiInsightId string
}

// NewMeOnlineMeetingIdAiInsightID returns a new MeOnlineMeetingIdAiInsightId struct
func NewMeOnlineMeetingIdAiInsightID(onlineMeetingId string, callAiInsightId string) MeOnlineMeetingIdAiInsightId {
	return MeOnlineMeetingIdAiInsightId{
		OnlineMeetingId: onlineMeetingId,
		CallAiInsightId: callAiInsightId,
	}
}

// ParseMeOnlineMeetingIdAiInsightID parses 'input' into a MeOnlineMeetingIdAiInsightId
func ParseMeOnlineMeetingIdAiInsightID(input string) (*MeOnlineMeetingIdAiInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdAiInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdAiInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnlineMeetingIdAiInsightIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingIdAiInsightId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingIdAiInsightIDInsensitively(input string) (*MeOnlineMeetingIdAiInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdAiInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdAiInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnlineMeetingIdAiInsightId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	if id.CallAiInsightId, ok = input.Parsed["callAiInsightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "callAiInsightId", input)
	}

	return nil
}

// ValidateMeOnlineMeetingIdAiInsightID checks that 'input' can be parsed as a Me Online Meeting Id Ai Insight ID
func ValidateMeOnlineMeetingIdAiInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingIdAiInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting Id Ai Insight ID
func (id MeOnlineMeetingIdAiInsightId) ID() string {
	fmtString := "/me/onlineMeetings/%s/aiInsights/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId, id.CallAiInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting Id Ai Insight ID
func (id MeOnlineMeetingIdAiInsightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("aiInsights", "aiInsights", "aiInsights"),
		resourceids.UserSpecifiedSegment("callAiInsightId", "callAiInsightId"),
	}
}

// String returns a human-readable description of this Me Online Meeting Id Ai Insight ID
func (id MeOnlineMeetingIdAiInsightId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Call Ai Insight: %q", id.CallAiInsightId),
	}
	return fmt.Sprintf("Me Online Meeting Id Ai Insight (%s)", strings.Join(components, "\n"))
}
