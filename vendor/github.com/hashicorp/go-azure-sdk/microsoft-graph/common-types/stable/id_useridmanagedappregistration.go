package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedAppRegistrationId{}

// UserIdManagedAppRegistrationId is a struct representing the Resource ID for a User Id Managed App Registration
type UserIdManagedAppRegistrationId struct {
	UserId                   string
	ManagedAppRegistrationId string
}

// NewUserIdManagedAppRegistrationID returns a new UserIdManagedAppRegistrationId struct
func NewUserIdManagedAppRegistrationID(userId string, managedAppRegistrationId string) UserIdManagedAppRegistrationId {
	return UserIdManagedAppRegistrationId{
		UserId:                   userId,
		ManagedAppRegistrationId: managedAppRegistrationId,
	}
}

// ParseUserIdManagedAppRegistrationID parses 'input' into a UserIdManagedAppRegistrationId
func ParseUserIdManagedAppRegistrationID(input string) (*UserIdManagedAppRegistrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedAppRegistrationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedAppRegistrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdManagedAppRegistrationIDInsensitively parses 'input' case-insensitively into a UserIdManagedAppRegistrationId
// note: this method should only be used for API response data and not user input
func ParseUserIdManagedAppRegistrationIDInsensitively(input string) (*UserIdManagedAppRegistrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedAppRegistrationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedAppRegistrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdManagedAppRegistrationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ManagedAppRegistrationId, ok = input.Parsed["managedAppRegistrationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedAppRegistrationId", input)
	}

	return nil
}

// ValidateUserIdManagedAppRegistrationID checks that 'input' can be parsed as a User Id Managed App Registration ID
func ValidateUserIdManagedAppRegistrationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdManagedAppRegistrationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Managed App Registration ID
func (id UserIdManagedAppRegistrationId) ID() string {
	fmtString := "/users/%s/managedAppRegistrations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedAppRegistrationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Managed App Registration ID
func (id UserIdManagedAppRegistrationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("managedAppRegistrations", "managedAppRegistrations", "managedAppRegistrations"),
		resourceids.UserSpecifiedSegment("managedAppRegistrationId", "managedAppRegistrationId"),
	}
}

// String returns a human-readable description of this User Id Managed App Registration ID
func (id UserIdManagedAppRegistrationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed App Registration: %q", id.ManagedAppRegistrationId),
	}
	return fmt.Sprintf("User Id Managed App Registration (%s)", strings.Join(components, "\n"))
}
