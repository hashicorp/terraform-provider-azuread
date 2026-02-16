package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileNoteId{}

// MeProfileNoteId is a struct representing the Resource ID for a Me Profile Note
type MeProfileNoteId struct {
	PersonAnnotationId string
}

// NewMeProfileNoteID returns a new MeProfileNoteId struct
func NewMeProfileNoteID(personAnnotationId string) MeProfileNoteId {
	return MeProfileNoteId{
		PersonAnnotationId: personAnnotationId,
	}
}

// ParseMeProfileNoteID parses 'input' into a MeProfileNoteId
func ParseMeProfileNoteID(input string) (*MeProfileNoteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileNoteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileNoteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfileNoteIDInsensitively parses 'input' case-insensitively into a MeProfileNoteId
// note: this method should only be used for API response data and not user input
func ParseMeProfileNoteIDInsensitively(input string) (*MeProfileNoteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileNoteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileNoteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfileNoteId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PersonAnnotationId, ok = input.Parsed["personAnnotationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personAnnotationId", input)
	}

	return nil
}

// ValidateMeProfileNoteID checks that 'input' can be parsed as a Me Profile Note ID
func ValidateMeProfileNoteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileNoteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Note ID
func (id MeProfileNoteId) ID() string {
	fmtString := "/me/profile/notes/%s"
	return fmt.Sprintf(fmtString, id.PersonAnnotationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Note ID
func (id MeProfileNoteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("notes", "notes", "notes"),
		resourceids.UserSpecifiedSegment("personAnnotationId", "personAnnotationId"),
	}
}

// String returns a human-readable description of this Me Profile Note ID
func (id MeProfileNoteId) String() string {
	components := []string{
		fmt.Sprintf("Person Annotation: %q", id.PersonAnnotationId),
	}
	return fmt.Sprintf("Me Profile Note (%s)", strings.Join(components, "\n"))
}
