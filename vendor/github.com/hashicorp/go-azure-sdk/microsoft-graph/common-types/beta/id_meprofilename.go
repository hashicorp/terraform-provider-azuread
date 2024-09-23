package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileNameId{}

// MeProfileNameId is a struct representing the Resource ID for a Me Profile Name
type MeProfileNameId struct {
	PersonNameId string
}

// NewMeProfileNameID returns a new MeProfileNameId struct
func NewMeProfileNameID(personNameId string) MeProfileNameId {
	return MeProfileNameId{
		PersonNameId: personNameId,
	}
}

// ParseMeProfileNameID parses 'input' into a MeProfileNameId
func ParseMeProfileNameID(input string) (*MeProfileNameId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileNameId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileNameId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfileNameIDInsensitively parses 'input' case-insensitively into a MeProfileNameId
// note: this method should only be used for API response data and not user input
func ParseMeProfileNameIDInsensitively(input string) (*MeProfileNameId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileNameId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileNameId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfileNameId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PersonNameId, ok = input.Parsed["personNameId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personNameId", input)
	}

	return nil
}

// ValidateMeProfileNameID checks that 'input' can be parsed as a Me Profile Name ID
func ValidateMeProfileNameID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileNameID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Name ID
func (id MeProfileNameId) ID() string {
	fmtString := "/me/profile/names/%s"
	return fmt.Sprintf(fmtString, id.PersonNameId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Name ID
func (id MeProfileNameId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("names", "names", "names"),
		resourceids.UserSpecifiedSegment("personNameId", "personNameId"),
	}
}

// String returns a human-readable description of this Me Profile Name ID
func (id MeProfileNameId) String() string {
	components := []string{
		fmt.Sprintf("Person Name: %q", id.PersonNameId),
	}
	return fmt.Sprintf("Me Profile Name (%s)", strings.Join(components, "\n"))
}
