package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileAwardId{}

// MeProfileAwardId is a struct representing the Resource ID for a Me Profile Award
type MeProfileAwardId struct {
	PersonAwardId string
}

// NewMeProfileAwardID returns a new MeProfileAwardId struct
func NewMeProfileAwardID(personAwardId string) MeProfileAwardId {
	return MeProfileAwardId{
		PersonAwardId: personAwardId,
	}
}

// ParseMeProfileAwardID parses 'input' into a MeProfileAwardId
func ParseMeProfileAwardID(input string) (*MeProfileAwardId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileAwardId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileAwardId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfileAwardIDInsensitively parses 'input' case-insensitively into a MeProfileAwardId
// note: this method should only be used for API response data and not user input
func ParseMeProfileAwardIDInsensitively(input string) (*MeProfileAwardId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileAwardId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileAwardId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfileAwardId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PersonAwardId, ok = input.Parsed["personAwardId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personAwardId", input)
	}

	return nil
}

// ValidateMeProfileAwardID checks that 'input' can be parsed as a Me Profile Award ID
func ValidateMeProfileAwardID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileAwardID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Award ID
func (id MeProfileAwardId) ID() string {
	fmtString := "/me/profile/awards/%s"
	return fmt.Sprintf(fmtString, id.PersonAwardId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Award ID
func (id MeProfileAwardId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("awards", "awards", "awards"),
		resourceids.UserSpecifiedSegment("personAwardId", "personAwardId"),
	}
}

// String returns a human-readable description of this Me Profile Award ID
func (id MeProfileAwardId) String() string {
	components := []string{
		fmt.Sprintf("Person Award: %q", id.PersonAwardId),
	}
	return fmt.Sprintf("Me Profile Award (%s)", strings.Join(components, "\n"))
}
