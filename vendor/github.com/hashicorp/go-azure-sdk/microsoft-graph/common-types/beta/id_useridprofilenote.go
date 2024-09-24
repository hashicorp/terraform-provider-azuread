package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileNoteId{}

// UserIdProfileNoteId is a struct representing the Resource ID for a User Id Profile Note
type UserIdProfileNoteId struct {
	UserId             string
	PersonAnnotationId string
}

// NewUserIdProfileNoteID returns a new UserIdProfileNoteId struct
func NewUserIdProfileNoteID(userId string, personAnnotationId string) UserIdProfileNoteId {
	return UserIdProfileNoteId{
		UserId:             userId,
		PersonAnnotationId: personAnnotationId,
	}
}

// ParseUserIdProfileNoteID parses 'input' into a UserIdProfileNoteId
func ParseUserIdProfileNoteID(input string) (*UserIdProfileNoteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileNoteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileNoteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfileNoteIDInsensitively parses 'input' case-insensitively into a UserIdProfileNoteId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfileNoteIDInsensitively(input string) (*UserIdProfileNoteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileNoteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileNoteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfileNoteId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PersonAnnotationId, ok = input.Parsed["personAnnotationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personAnnotationId", input)
	}

	return nil
}

// ValidateUserIdProfileNoteID checks that 'input' can be parsed as a User Id Profile Note ID
func ValidateUserIdProfileNoteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfileNoteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Note ID
func (id UserIdProfileNoteId) ID() string {
	fmtString := "/users/%s/profile/notes/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonAnnotationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Note ID
func (id UserIdProfileNoteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("notes", "notes", "notes"),
		resourceids.UserSpecifiedSegment("personAnnotationId", "personAnnotationId"),
	}
}

// String returns a human-readable description of this User Id Profile Note ID
func (id UserIdProfileNoteId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person Annotation: %q", id.PersonAnnotationId),
	}
	return fmt.Sprintf("User Id Profile Note (%s)", strings.Join(components, "\n"))
}
