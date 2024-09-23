package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileAnniversaryId{}

// MeProfileAnniversaryId is a struct representing the Resource ID for a Me Profile Anniversary
type MeProfileAnniversaryId struct {
	PersonAnnualEventId string
}

// NewMeProfileAnniversaryID returns a new MeProfileAnniversaryId struct
func NewMeProfileAnniversaryID(personAnnualEventId string) MeProfileAnniversaryId {
	return MeProfileAnniversaryId{
		PersonAnnualEventId: personAnnualEventId,
	}
}

// ParseMeProfileAnniversaryID parses 'input' into a MeProfileAnniversaryId
func ParseMeProfileAnniversaryID(input string) (*MeProfileAnniversaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileAnniversaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileAnniversaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfileAnniversaryIDInsensitively parses 'input' case-insensitively into a MeProfileAnniversaryId
// note: this method should only be used for API response data and not user input
func ParseMeProfileAnniversaryIDInsensitively(input string) (*MeProfileAnniversaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileAnniversaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileAnniversaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfileAnniversaryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PersonAnnualEventId, ok = input.Parsed["personAnnualEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personAnnualEventId", input)
	}

	return nil
}

// ValidateMeProfileAnniversaryID checks that 'input' can be parsed as a Me Profile Anniversary ID
func ValidateMeProfileAnniversaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileAnniversaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Anniversary ID
func (id MeProfileAnniversaryId) ID() string {
	fmtString := "/me/profile/anniversaries/%s"
	return fmt.Sprintf(fmtString, id.PersonAnnualEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Anniversary ID
func (id MeProfileAnniversaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("anniversaries", "anniversaries", "anniversaries"),
		resourceids.UserSpecifiedSegment("personAnnualEventId", "personAnnualEventId"),
	}
}

// String returns a human-readable description of this Me Profile Anniversary ID
func (id MeProfileAnniversaryId) String() string {
	components := []string{
		fmt.Sprintf("Person Annual Event: %q", id.PersonAnnualEventId),
	}
	return fmt.Sprintf("Me Profile Anniversary (%s)", strings.Join(components, "\n"))
}
