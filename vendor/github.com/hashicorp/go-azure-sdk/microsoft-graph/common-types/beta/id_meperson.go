package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePersonId{}

// MePersonId is a struct representing the Resource ID for a Me Person
type MePersonId struct {
	PersonId string
}

// NewMePersonID returns a new MePersonId struct
func NewMePersonID(personId string) MePersonId {
	return MePersonId{
		PersonId: personId,
	}
}

// ParseMePersonID parses 'input' into a MePersonId
func ParseMePersonID(input string) (*MePersonId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePersonId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePersonId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePersonIDInsensitively parses 'input' case-insensitively into a MePersonId
// note: this method should only be used for API response data and not user input
func ParseMePersonIDInsensitively(input string) (*MePersonId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePersonId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePersonId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePersonId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PersonId, ok = input.Parsed["personId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personId", input)
	}

	return nil
}

// ValidateMePersonID checks that 'input' can be parsed as a Me Person ID
func ValidateMePersonID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePersonID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Person ID
func (id MePersonId) ID() string {
	fmtString := "/me/people/%s"
	return fmt.Sprintf(fmtString, id.PersonId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Person ID
func (id MePersonId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("people", "people", "people"),
		resourceids.UserSpecifiedSegment("personId", "personId"),
	}
}

// String returns a human-readable description of this Me Person ID
func (id MePersonId) String() string {
	components := []string{
		fmt.Sprintf("Person: %q", id.PersonId),
	}
	return fmt.Sprintf("Me Person (%s)", strings.Join(components, "\n"))
}
