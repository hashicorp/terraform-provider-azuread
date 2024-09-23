package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInsightUsedId{}

// MeInsightUsedId is a struct representing the Resource ID for a Me Insight Used
type MeInsightUsedId struct {
	UsedInsightId string
}

// NewMeInsightUsedID returns a new MeInsightUsedId struct
func NewMeInsightUsedID(usedInsightId string) MeInsightUsedId {
	return MeInsightUsedId{
		UsedInsightId: usedInsightId,
	}
}

// ParseMeInsightUsedID parses 'input' into a MeInsightUsedId
func ParseMeInsightUsedID(input string) (*MeInsightUsedId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInsightUsedId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInsightUsedId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeInsightUsedIDInsensitively parses 'input' case-insensitively into a MeInsightUsedId
// note: this method should only be used for API response data and not user input
func ParseMeInsightUsedIDInsensitively(input string) (*MeInsightUsedId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInsightUsedId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInsightUsedId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeInsightUsedId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UsedInsightId, ok = input.Parsed["usedInsightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "usedInsightId", input)
	}

	return nil
}

// ValidateMeInsightUsedID checks that 'input' can be parsed as a Me Insight Used ID
func ValidateMeInsightUsedID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeInsightUsedID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Insight Used ID
func (id MeInsightUsedId) ID() string {
	fmtString := "/me/insights/used/%s"
	return fmt.Sprintf(fmtString, id.UsedInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Insight Used ID
func (id MeInsightUsedId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("insights", "insights", "insights"),
		resourceids.StaticSegment("used", "used", "used"),
		resourceids.UserSpecifiedSegment("usedInsightId", "usedInsightId"),
	}
}

// String returns a human-readable description of this Me Insight Used ID
func (id MeInsightUsedId) String() string {
	components := []string{
		fmt.Sprintf("Used Insight: %q", id.UsedInsightId),
	}
	return fmt.Sprintf("Me Insight Used (%s)", strings.Join(components, "\n"))
}
