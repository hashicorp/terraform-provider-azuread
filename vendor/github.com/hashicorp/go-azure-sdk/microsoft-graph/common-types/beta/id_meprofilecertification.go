package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileCertificationId{}

// MeProfileCertificationId is a struct representing the Resource ID for a Me Profile Certification
type MeProfileCertificationId struct {
	PersonCertificationId string
}

// NewMeProfileCertificationID returns a new MeProfileCertificationId struct
func NewMeProfileCertificationID(personCertificationId string) MeProfileCertificationId {
	return MeProfileCertificationId{
		PersonCertificationId: personCertificationId,
	}
}

// ParseMeProfileCertificationID parses 'input' into a MeProfileCertificationId
func ParseMeProfileCertificationID(input string) (*MeProfileCertificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileCertificationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileCertificationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfileCertificationIDInsensitively parses 'input' case-insensitively into a MeProfileCertificationId
// note: this method should only be used for API response data and not user input
func ParseMeProfileCertificationIDInsensitively(input string) (*MeProfileCertificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileCertificationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileCertificationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfileCertificationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PersonCertificationId, ok = input.Parsed["personCertificationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personCertificationId", input)
	}

	return nil
}

// ValidateMeProfileCertificationID checks that 'input' can be parsed as a Me Profile Certification ID
func ValidateMeProfileCertificationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileCertificationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Certification ID
func (id MeProfileCertificationId) ID() string {
	fmtString := "/me/profile/certifications/%s"
	return fmt.Sprintf(fmtString, id.PersonCertificationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Certification ID
func (id MeProfileCertificationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("certifications", "certifications", "certifications"),
		resourceids.UserSpecifiedSegment("personCertificationId", "personCertificationId"),
	}
}

// String returns a human-readable description of this Me Profile Certification ID
func (id MeProfileCertificationId) String() string {
	components := []string{
		fmt.Sprintf("Person Certification: %q", id.PersonCertificationId),
	}
	return fmt.Sprintf("Me Profile Certification (%s)", strings.Join(components, "\n"))
}
