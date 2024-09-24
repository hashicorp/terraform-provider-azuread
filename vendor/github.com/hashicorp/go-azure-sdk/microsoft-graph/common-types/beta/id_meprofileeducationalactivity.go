package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileEducationalActivityId{}

// MeProfileEducationalActivityId is a struct representing the Resource ID for a Me Profile Educational Activity
type MeProfileEducationalActivityId struct {
	EducationalActivityId string
}

// NewMeProfileEducationalActivityID returns a new MeProfileEducationalActivityId struct
func NewMeProfileEducationalActivityID(educationalActivityId string) MeProfileEducationalActivityId {
	return MeProfileEducationalActivityId{
		EducationalActivityId: educationalActivityId,
	}
}

// ParseMeProfileEducationalActivityID parses 'input' into a MeProfileEducationalActivityId
func ParseMeProfileEducationalActivityID(input string) (*MeProfileEducationalActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileEducationalActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileEducationalActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfileEducationalActivityIDInsensitively parses 'input' case-insensitively into a MeProfileEducationalActivityId
// note: this method should only be used for API response data and not user input
func ParseMeProfileEducationalActivityIDInsensitively(input string) (*MeProfileEducationalActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileEducationalActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileEducationalActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfileEducationalActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EducationalActivityId, ok = input.Parsed["educationalActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "educationalActivityId", input)
	}

	return nil
}

// ValidateMeProfileEducationalActivityID checks that 'input' can be parsed as a Me Profile Educational Activity ID
func ValidateMeProfileEducationalActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileEducationalActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Educational Activity ID
func (id MeProfileEducationalActivityId) ID() string {
	fmtString := "/me/profile/educationalActivities/%s"
	return fmt.Sprintf(fmtString, id.EducationalActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Educational Activity ID
func (id MeProfileEducationalActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("educationalActivities", "educationalActivities", "educationalActivities"),
		resourceids.UserSpecifiedSegment("educationalActivityId", "educationalActivityId"),
	}
}

// String returns a human-readable description of this Me Profile Educational Activity ID
func (id MeProfileEducationalActivityId) String() string {
	components := []string{
		fmt.Sprintf("Educational Activity: %q", id.EducationalActivityId),
	}
	return fmt.Sprintf("Me Profile Educational Activity (%s)", strings.Join(components, "\n"))
}
