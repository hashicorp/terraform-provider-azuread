package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdContactIdExtensionId{}

// UserIdContactIdExtensionId is a struct representing the Resource ID for a User Id Contact Id Extension
type UserIdContactIdExtensionId struct {
	UserId      string
	ContactId   string
	ExtensionId string
}

// NewUserIdContactIdExtensionID returns a new UserIdContactIdExtensionId struct
func NewUserIdContactIdExtensionID(userId string, contactId string, extensionId string) UserIdContactIdExtensionId {
	return UserIdContactIdExtensionId{
		UserId:      userId,
		ContactId:   contactId,
		ExtensionId: extensionId,
	}
}

// ParseUserIdContactIdExtensionID parses 'input' into a UserIdContactIdExtensionId
func ParseUserIdContactIdExtensionID(input string) (*UserIdContactIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdContactIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdContactIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdContactIdExtensionIDInsensitively(input string) (*UserIdContactIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdContactIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ContactId, ok = input.Parsed["contactId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdContactIdExtensionID checks that 'input' can be parsed as a User Id Contact Id Extension ID
func ValidateUserIdContactIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdContactIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Contact Id Extension ID
func (id UserIdContactIdExtensionId) ID() string {
	fmtString := "/users/%s/contacts/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Contact Id Extension ID
func (id UserIdContactIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("contacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this User Id Contact Id Extension ID
func (id UserIdContactIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact: %q", id.ContactId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Contact Id Extension (%s)", strings.Join(components, "\n"))
}
