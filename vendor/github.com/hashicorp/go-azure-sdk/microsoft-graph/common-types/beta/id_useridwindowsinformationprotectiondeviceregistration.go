package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdWindowsInformationProtectionDeviceRegistrationId{}

// UserIdWindowsInformationProtectionDeviceRegistrationId is a struct representing the Resource ID for a User Id Windows Information Protection Device Registration
type UserIdWindowsInformationProtectionDeviceRegistrationId struct {
	UserId                                           string
	WindowsInformationProtectionDeviceRegistrationId string
}

// NewUserIdWindowsInformationProtectionDeviceRegistrationID returns a new UserIdWindowsInformationProtectionDeviceRegistrationId struct
func NewUserIdWindowsInformationProtectionDeviceRegistrationID(userId string, windowsInformationProtectionDeviceRegistrationId string) UserIdWindowsInformationProtectionDeviceRegistrationId {
	return UserIdWindowsInformationProtectionDeviceRegistrationId{
		UserId: userId,
		WindowsInformationProtectionDeviceRegistrationId: windowsInformationProtectionDeviceRegistrationId,
	}
}

// ParseUserIdWindowsInformationProtectionDeviceRegistrationID parses 'input' into a UserIdWindowsInformationProtectionDeviceRegistrationId
func ParseUserIdWindowsInformationProtectionDeviceRegistrationID(input string) (*UserIdWindowsInformationProtectionDeviceRegistrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdWindowsInformationProtectionDeviceRegistrationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdWindowsInformationProtectionDeviceRegistrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdWindowsInformationProtectionDeviceRegistrationIDInsensitively parses 'input' case-insensitively into a UserIdWindowsInformationProtectionDeviceRegistrationId
// note: this method should only be used for API response data and not user input
func ParseUserIdWindowsInformationProtectionDeviceRegistrationIDInsensitively(input string) (*UserIdWindowsInformationProtectionDeviceRegistrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdWindowsInformationProtectionDeviceRegistrationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdWindowsInformationProtectionDeviceRegistrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdWindowsInformationProtectionDeviceRegistrationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.WindowsInformationProtectionDeviceRegistrationId, ok = input.Parsed["windowsInformationProtectionDeviceRegistrationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsInformationProtectionDeviceRegistrationId", input)
	}

	return nil
}

// ValidateUserIdWindowsInformationProtectionDeviceRegistrationID checks that 'input' can be parsed as a User Id Windows Information Protection Device Registration ID
func ValidateUserIdWindowsInformationProtectionDeviceRegistrationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdWindowsInformationProtectionDeviceRegistrationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Windows Information Protection Device Registration ID
func (id UserIdWindowsInformationProtectionDeviceRegistrationId) ID() string {
	fmtString := "/users/%s/windowsInformationProtectionDeviceRegistrations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.WindowsInformationProtectionDeviceRegistrationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Windows Information Protection Device Registration ID
func (id UserIdWindowsInformationProtectionDeviceRegistrationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("windowsInformationProtectionDeviceRegistrations", "windowsInformationProtectionDeviceRegistrations", "windowsInformationProtectionDeviceRegistrations"),
		resourceids.UserSpecifiedSegment("windowsInformationProtectionDeviceRegistrationId", "windowsInformationProtectionDeviceRegistrationId"),
	}
}

// String returns a human-readable description of this User Id Windows Information Protection Device Registration ID
func (id UserIdWindowsInformationProtectionDeviceRegistrationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Windows Information Protection Device Registration: %q", id.WindowsInformationProtectionDeviceRegistrationId),
	}
	return fmt.Sprintf("User Id Windows Information Protection Device Registration (%s)", strings.Join(components, "\n"))
}
