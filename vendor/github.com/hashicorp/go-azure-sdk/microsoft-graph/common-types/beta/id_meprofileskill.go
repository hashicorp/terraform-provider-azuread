package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileSkillId{}

// MeProfileSkillId is a struct representing the Resource ID for a Me Profile Skill
type MeProfileSkillId struct {
	SkillProficiencyId string
}

// NewMeProfileSkillID returns a new MeProfileSkillId struct
func NewMeProfileSkillID(skillProficiencyId string) MeProfileSkillId {
	return MeProfileSkillId{
		SkillProficiencyId: skillProficiencyId,
	}
}

// ParseMeProfileSkillID parses 'input' into a MeProfileSkillId
func ParseMeProfileSkillID(input string) (*MeProfileSkillId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileSkillId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileSkillId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfileSkillIDInsensitively parses 'input' case-insensitively into a MeProfileSkillId
// note: this method should only be used for API response data and not user input
func ParseMeProfileSkillIDInsensitively(input string) (*MeProfileSkillId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileSkillId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileSkillId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfileSkillId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SkillProficiencyId, ok = input.Parsed["skillProficiencyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "skillProficiencyId", input)
	}

	return nil
}

// ValidateMeProfileSkillID checks that 'input' can be parsed as a Me Profile Skill ID
func ValidateMeProfileSkillID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileSkillID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Skill ID
func (id MeProfileSkillId) ID() string {
	fmtString := "/me/profile/skills/%s"
	return fmt.Sprintf(fmtString, id.SkillProficiencyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Skill ID
func (id MeProfileSkillId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("skills", "skills", "skills"),
		resourceids.UserSpecifiedSegment("skillProficiencyId", "skillProficiencyId"),
	}
}

// String returns a human-readable description of this Me Profile Skill ID
func (id MeProfileSkillId) String() string {
	components := []string{
		fmt.Sprintf("Skill Proficiency: %q", id.SkillProficiencyId),
	}
	return fmt.Sprintf("Me Profile Skill (%s)", strings.Join(components, "\n"))
}
