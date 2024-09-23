package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileSkillId{}

// UserIdProfileSkillId is a struct representing the Resource ID for a User Id Profile Skill
type UserIdProfileSkillId struct {
	UserId             string
	SkillProficiencyId string
}

// NewUserIdProfileSkillID returns a new UserIdProfileSkillId struct
func NewUserIdProfileSkillID(userId string, skillProficiencyId string) UserIdProfileSkillId {
	return UserIdProfileSkillId{
		UserId:             userId,
		SkillProficiencyId: skillProficiencyId,
	}
}

// ParseUserIdProfileSkillID parses 'input' into a UserIdProfileSkillId
func ParseUserIdProfileSkillID(input string) (*UserIdProfileSkillId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileSkillId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileSkillId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfileSkillIDInsensitively parses 'input' case-insensitively into a UserIdProfileSkillId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfileSkillIDInsensitively(input string) (*UserIdProfileSkillId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileSkillId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileSkillId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfileSkillId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.SkillProficiencyId, ok = input.Parsed["skillProficiencyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "skillProficiencyId", input)
	}

	return nil
}

// ValidateUserIdProfileSkillID checks that 'input' can be parsed as a User Id Profile Skill ID
func ValidateUserIdProfileSkillID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfileSkillID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Skill ID
func (id UserIdProfileSkillId) ID() string {
	fmtString := "/users/%s/profile/skills/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SkillProficiencyId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Skill ID
func (id UserIdProfileSkillId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("skills", "skills", "skills"),
		resourceids.UserSpecifiedSegment("skillProficiencyId", "skillProficiencyId"),
	}
}

// String returns a human-readable description of this User Id Profile Skill ID
func (id UserIdProfileSkillId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Skill Proficiency: %q", id.SkillProficiencyId),
	}
	return fmt.Sprintf("User Id Profile Skill (%s)", strings.Join(components, "\n"))
}
