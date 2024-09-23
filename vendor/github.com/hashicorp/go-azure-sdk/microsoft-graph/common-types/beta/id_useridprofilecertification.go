package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileCertificationId{}

// UserIdProfileCertificationId is a struct representing the Resource ID for a User Id Profile Certification
type UserIdProfileCertificationId struct {
	UserId                string
	PersonCertificationId string
}

// NewUserIdProfileCertificationID returns a new UserIdProfileCertificationId struct
func NewUserIdProfileCertificationID(userId string, personCertificationId string) UserIdProfileCertificationId {
	return UserIdProfileCertificationId{
		UserId:                userId,
		PersonCertificationId: personCertificationId,
	}
}

// ParseUserIdProfileCertificationID parses 'input' into a UserIdProfileCertificationId
func ParseUserIdProfileCertificationID(input string) (*UserIdProfileCertificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileCertificationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileCertificationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfileCertificationIDInsensitively parses 'input' case-insensitively into a UserIdProfileCertificationId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfileCertificationIDInsensitively(input string) (*UserIdProfileCertificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileCertificationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileCertificationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfileCertificationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PersonCertificationId, ok = input.Parsed["personCertificationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personCertificationId", input)
	}

	return nil
}

// ValidateUserIdProfileCertificationID checks that 'input' can be parsed as a User Id Profile Certification ID
func ValidateUserIdProfileCertificationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfileCertificationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Certification ID
func (id UserIdProfileCertificationId) ID() string {
	fmtString := "/users/%s/profile/certifications/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonCertificationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Certification ID
func (id UserIdProfileCertificationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("certifications", "certifications", "certifications"),
		resourceids.UserSpecifiedSegment("personCertificationId", "personCertificationId"),
	}
}

// String returns a human-readable description of this User Id Profile Certification ID
func (id UserIdProfileCertificationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person Certification: %q", id.PersonCertificationId),
	}
	return fmt.Sprintf("User Id Profile Certification (%s)", strings.Join(components, "\n"))
}
