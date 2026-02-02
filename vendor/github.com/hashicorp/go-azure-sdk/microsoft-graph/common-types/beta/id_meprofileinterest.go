package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileInterestId{}

// MeProfileInterestId is a struct representing the Resource ID for a Me Profile Interest
type MeProfileInterestId struct {
	PersonInterestId string
}

// NewMeProfileInterestID returns a new MeProfileInterestId struct
func NewMeProfileInterestID(personInterestId string) MeProfileInterestId {
	return MeProfileInterestId{
		PersonInterestId: personInterestId,
	}
}

// ParseMeProfileInterestID parses 'input' into a MeProfileInterestId
func ParseMeProfileInterestID(input string) (*MeProfileInterestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileInterestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileInterestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfileInterestIDInsensitively parses 'input' case-insensitively into a MeProfileInterestId
// note: this method should only be used for API response data and not user input
func ParseMeProfileInterestIDInsensitively(input string) (*MeProfileInterestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileInterestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileInterestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfileInterestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PersonInterestId, ok = input.Parsed["personInterestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personInterestId", input)
	}

	return nil
}

// ValidateMeProfileInterestID checks that 'input' can be parsed as a Me Profile Interest ID
func ValidateMeProfileInterestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileInterestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Interest ID
func (id MeProfileInterestId) ID() string {
	fmtString := "/me/profile/interests/%s"
	return fmt.Sprintf(fmtString, id.PersonInterestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Interest ID
func (id MeProfileInterestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("interests", "interests", "interests"),
		resourceids.UserSpecifiedSegment("personInterestId", "personInterestId"),
	}
}

// String returns a human-readable description of this Me Profile Interest ID
func (id MeProfileInterestId) String() string {
	components := []string{
		fmt.Sprintf("Person Interest: %q", id.PersonInterestId),
	}
	return fmt.Sprintf("Me Profile Interest (%s)", strings.Join(components, "\n"))
}
