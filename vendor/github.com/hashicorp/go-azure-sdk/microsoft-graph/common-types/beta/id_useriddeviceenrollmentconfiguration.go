package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDeviceEnrollmentConfigurationId{}

// UserIdDeviceEnrollmentConfigurationId is a struct representing the Resource ID for a User Id Device Enrollment Configuration
type UserIdDeviceEnrollmentConfigurationId struct {
	UserId                          string
	DeviceEnrollmentConfigurationId string
}

// NewUserIdDeviceEnrollmentConfigurationID returns a new UserIdDeviceEnrollmentConfigurationId struct
func NewUserIdDeviceEnrollmentConfigurationID(userId string, deviceEnrollmentConfigurationId string) UserIdDeviceEnrollmentConfigurationId {
	return UserIdDeviceEnrollmentConfigurationId{
		UserId:                          userId,
		DeviceEnrollmentConfigurationId: deviceEnrollmentConfigurationId,
	}
}

// ParseUserIdDeviceEnrollmentConfigurationID parses 'input' into a UserIdDeviceEnrollmentConfigurationId
func ParseUserIdDeviceEnrollmentConfigurationID(input string) (*UserIdDeviceEnrollmentConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceEnrollmentConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceEnrollmentConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDeviceEnrollmentConfigurationIDInsensitively parses 'input' case-insensitively into a UserIdDeviceEnrollmentConfigurationId
// note: this method should only be used for API response data and not user input
func ParseUserIdDeviceEnrollmentConfigurationIDInsensitively(input string) (*UserIdDeviceEnrollmentConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceEnrollmentConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceEnrollmentConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDeviceEnrollmentConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DeviceEnrollmentConfigurationId, ok = input.Parsed["deviceEnrollmentConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceEnrollmentConfigurationId", input)
	}

	return nil
}

// ValidateUserIdDeviceEnrollmentConfigurationID checks that 'input' can be parsed as a User Id Device Enrollment Configuration ID
func ValidateUserIdDeviceEnrollmentConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDeviceEnrollmentConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Device Enrollment Configuration ID
func (id UserIdDeviceEnrollmentConfigurationId) ID() string {
	fmtString := "/users/%s/deviceEnrollmentConfigurations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceEnrollmentConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Device Enrollment Configuration ID
func (id UserIdDeviceEnrollmentConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("deviceEnrollmentConfigurations", "deviceEnrollmentConfigurations", "deviceEnrollmentConfigurations"),
		resourceids.UserSpecifiedSegment("deviceEnrollmentConfigurationId", "deviceEnrollmentConfigurationId"),
	}
}

// String returns a human-readable description of this User Id Device Enrollment Configuration ID
func (id UserIdDeviceEnrollmentConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device Enrollment Configuration: %q", id.DeviceEnrollmentConfigurationId),
	}
	return fmt.Sprintf("User Id Device Enrollment Configuration (%s)", strings.Join(components, "\n"))
}
