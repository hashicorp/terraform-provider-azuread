package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileAnniversaryId{}

// UserIdProfileAnniversaryId is a struct representing the Resource ID for a User Id Profile Anniversary
type UserIdProfileAnniversaryId struct {
	UserId              string
	PersonAnnualEventId string
}

// NewUserIdProfileAnniversaryID returns a new UserIdProfileAnniversaryId struct
func NewUserIdProfileAnniversaryID(userId string, personAnnualEventId string) UserIdProfileAnniversaryId {
	return UserIdProfileAnniversaryId{
		UserId:              userId,
		PersonAnnualEventId: personAnnualEventId,
	}
}

// ParseUserIdProfileAnniversaryID parses 'input' into a UserIdProfileAnniversaryId
func ParseUserIdProfileAnniversaryID(input string) (*UserIdProfileAnniversaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileAnniversaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileAnniversaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfileAnniversaryIDInsensitively parses 'input' case-insensitively into a UserIdProfileAnniversaryId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfileAnniversaryIDInsensitively(input string) (*UserIdProfileAnniversaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileAnniversaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileAnniversaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfileAnniversaryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PersonAnnualEventId, ok = input.Parsed["personAnnualEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personAnnualEventId", input)
	}

	return nil
}

// ValidateUserIdProfileAnniversaryID checks that 'input' can be parsed as a User Id Profile Anniversary ID
func ValidateUserIdProfileAnniversaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfileAnniversaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Anniversary ID
func (id UserIdProfileAnniversaryId) ID() string {
	fmtString := "/users/%s/profile/anniversaries/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonAnnualEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Anniversary ID
func (id UserIdProfileAnniversaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("anniversaries", "anniversaries", "anniversaries"),
		resourceids.UserSpecifiedSegment("personAnnualEventId", "personAnnualEventId"),
	}
}

// String returns a human-readable description of this User Id Profile Anniversary ID
func (id UserIdProfileAnniversaryId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person Annual Event: %q", id.PersonAnnualEventId),
	}
	return fmt.Sprintf("User Id Profile Anniversary (%s)", strings.Join(components, "\n"))
}
