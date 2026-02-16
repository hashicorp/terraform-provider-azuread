package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileLanguageId{}

// UserIdProfileLanguageId is a struct representing the Resource ID for a User Id Profile Language
type UserIdProfileLanguageId struct {
	UserId                string
	LanguageProficiencyId string
}

// NewUserIdProfileLanguageID returns a new UserIdProfileLanguageId struct
func NewUserIdProfileLanguageID(userId string, languageProficiencyId string) UserIdProfileLanguageId {
	return UserIdProfileLanguageId{
		UserId:                userId,
		LanguageProficiencyId: languageProficiencyId,
	}
}

// ParseUserIdProfileLanguageID parses 'input' into a UserIdProfileLanguageId
func ParseUserIdProfileLanguageID(input string) (*UserIdProfileLanguageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileLanguageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileLanguageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfileLanguageIDInsensitively parses 'input' case-insensitively into a UserIdProfileLanguageId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfileLanguageIDInsensitively(input string) (*UserIdProfileLanguageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileLanguageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileLanguageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfileLanguageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.LanguageProficiencyId, ok = input.Parsed["languageProficiencyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "languageProficiencyId", input)
	}

	return nil
}

// ValidateUserIdProfileLanguageID checks that 'input' can be parsed as a User Id Profile Language ID
func ValidateUserIdProfileLanguageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfileLanguageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Language ID
func (id UserIdProfileLanguageId) ID() string {
	fmtString := "/users/%s/profile/languages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.LanguageProficiencyId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Language ID
func (id UserIdProfileLanguageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("languages", "languages", "languages"),
		resourceids.UserSpecifiedSegment("languageProficiencyId", "languageProficiencyId"),
	}
}

// String returns a human-readable description of this User Id Profile Language ID
func (id UserIdProfileLanguageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Language Proficiency: %q", id.LanguageProficiencyId),
	}
	return fmt.Sprintf("User Id Profile Language (%s)", strings.Join(components, "\n"))
}
