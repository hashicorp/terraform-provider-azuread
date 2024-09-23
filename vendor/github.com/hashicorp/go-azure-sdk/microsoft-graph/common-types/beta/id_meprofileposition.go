package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfilePositionId{}

// MeProfilePositionId is a struct representing the Resource ID for a Me Profile Position
type MeProfilePositionId struct {
	WorkPositionId string
}

// NewMeProfilePositionID returns a new MeProfilePositionId struct
func NewMeProfilePositionID(workPositionId string) MeProfilePositionId {
	return MeProfilePositionId{
		WorkPositionId: workPositionId,
	}
}

// ParseMeProfilePositionID parses 'input' into a MeProfilePositionId
func ParseMeProfilePositionID(input string) (*MeProfilePositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfilePositionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfilePositionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfilePositionIDInsensitively parses 'input' case-insensitively into a MeProfilePositionId
// note: this method should only be used for API response data and not user input
func ParseMeProfilePositionIDInsensitively(input string) (*MeProfilePositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfilePositionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfilePositionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfilePositionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkPositionId, ok = input.Parsed["workPositionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workPositionId", input)
	}

	return nil
}

// ValidateMeProfilePositionID checks that 'input' can be parsed as a Me Profile Position ID
func ValidateMeProfilePositionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfilePositionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Position ID
func (id MeProfilePositionId) ID() string {
	fmtString := "/me/profile/positions/%s"
	return fmt.Sprintf(fmtString, id.WorkPositionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Position ID
func (id MeProfilePositionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("positions", "positions", "positions"),
		resourceids.UserSpecifiedSegment("workPositionId", "workPositionId"),
	}
}

// String returns a human-readable description of this Me Profile Position ID
func (id MeProfilePositionId) String() string {
	components := []string{
		fmt.Sprintf("Work Position: %q", id.WorkPositionId),
	}
	return fmt.Sprintf("Me Profile Position (%s)", strings.Join(components, "\n"))
}
