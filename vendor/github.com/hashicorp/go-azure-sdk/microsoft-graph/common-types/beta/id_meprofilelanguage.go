package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileLanguageId{}

// MeProfileLanguageId is a struct representing the Resource ID for a Me Profile Language
type MeProfileLanguageId struct {
	LanguageProficiencyId string
}

// NewMeProfileLanguageID returns a new MeProfileLanguageId struct
func NewMeProfileLanguageID(languageProficiencyId string) MeProfileLanguageId {
	return MeProfileLanguageId{
		LanguageProficiencyId: languageProficiencyId,
	}
}

// ParseMeProfileLanguageID parses 'input' into a MeProfileLanguageId
func ParseMeProfileLanguageID(input string) (*MeProfileLanguageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileLanguageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileLanguageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfileLanguageIDInsensitively parses 'input' case-insensitively into a MeProfileLanguageId
// note: this method should only be used for API response data and not user input
func ParseMeProfileLanguageIDInsensitively(input string) (*MeProfileLanguageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileLanguageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileLanguageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfileLanguageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.LanguageProficiencyId, ok = input.Parsed["languageProficiencyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "languageProficiencyId", input)
	}

	return nil
}

// ValidateMeProfileLanguageID checks that 'input' can be parsed as a Me Profile Language ID
func ValidateMeProfileLanguageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileLanguageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Language ID
func (id MeProfileLanguageId) ID() string {
	fmtString := "/me/profile/languages/%s"
	return fmt.Sprintf(fmtString, id.LanguageProficiencyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Language ID
func (id MeProfileLanguageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("languages", "languages", "languages"),
		resourceids.UserSpecifiedSegment("languageProficiencyId", "languageProficiencyId"),
	}
}

// String returns a human-readable description of this Me Profile Language ID
func (id MeProfileLanguageId) String() string {
	components := []string{
		fmt.Sprintf("Language Proficiency: %q", id.LanguageProficiencyId),
	}
	return fmt.Sprintf("Me Profile Language (%s)", strings.Join(components, "\n"))
}
