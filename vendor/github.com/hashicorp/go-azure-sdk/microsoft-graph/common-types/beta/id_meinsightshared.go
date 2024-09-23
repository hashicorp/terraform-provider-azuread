package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInsightSharedId{}

// MeInsightSharedId is a struct representing the Resource ID for a Me Insight Shared
type MeInsightSharedId struct {
	SharedInsightId string
}

// NewMeInsightSharedID returns a new MeInsightSharedId struct
func NewMeInsightSharedID(sharedInsightId string) MeInsightSharedId {
	return MeInsightSharedId{
		SharedInsightId: sharedInsightId,
	}
}

// ParseMeInsightSharedID parses 'input' into a MeInsightSharedId
func ParseMeInsightSharedID(input string) (*MeInsightSharedId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInsightSharedId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInsightSharedId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeInsightSharedIDInsensitively parses 'input' case-insensitively into a MeInsightSharedId
// note: this method should only be used for API response data and not user input
func ParseMeInsightSharedIDInsensitively(input string) (*MeInsightSharedId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInsightSharedId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInsightSharedId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeInsightSharedId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SharedInsightId, ok = input.Parsed["sharedInsightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sharedInsightId", input)
	}

	return nil
}

// ValidateMeInsightSharedID checks that 'input' can be parsed as a Me Insight Shared ID
func ValidateMeInsightSharedID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeInsightSharedID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Insight Shared ID
func (id MeInsightSharedId) ID() string {
	fmtString := "/me/insights/shared/%s"
	return fmt.Sprintf(fmtString, id.SharedInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Insight Shared ID
func (id MeInsightSharedId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("insights", "insights", "insights"),
		resourceids.StaticSegment("shared", "shared", "shared"),
		resourceids.UserSpecifiedSegment("sharedInsightId", "sharedInsightId"),
	}
}

// String returns a human-readable description of this Me Insight Shared ID
func (id MeInsightSharedId) String() string {
	components := []string{
		fmt.Sprintf("Shared Insight: %q", id.SharedInsightId),
	}
	return fmt.Sprintf("Me Insight Shared (%s)", strings.Join(components, "\n"))
}
