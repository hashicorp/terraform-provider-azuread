package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdContactId{}

// UserIdContactId is a struct representing the Resource ID for a User Id Contact
type UserIdContactId struct {
	UserId    string
	ContactId string
}

// NewUserIdContactID returns a new UserIdContactId struct
func NewUserIdContactID(userId string, contactId string) UserIdContactId {
	return UserIdContactId{
		UserId:    userId,
		ContactId: contactId,
	}
}

// ParseUserIdContactID parses 'input' into a UserIdContactId
func ParseUserIdContactID(input string) (*UserIdContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdContactIDInsensitively parses 'input' case-insensitively into a UserIdContactId
// note: this method should only be used for API response data and not user input
func ParseUserIdContactIDInsensitively(input string) (*UserIdContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdContactId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ContactId, ok = input.Parsed["contactId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactId", input)
	}

	return nil
}

// ValidateUserIdContactID checks that 'input' can be parsed as a User Id Contact ID
func ValidateUserIdContactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdContactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Contact ID
func (id UserIdContactId) ID() string {
	fmtString := "/users/%s/contacts/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Contact ID
func (id UserIdContactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("contacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactId"),
	}
}

// String returns a human-readable description of this User Id Contact ID
func (id UserIdContactId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact: %q", id.ContactId),
	}
	return fmt.Sprintf("User Id Contact (%s)", strings.Join(components, "\n"))
}
