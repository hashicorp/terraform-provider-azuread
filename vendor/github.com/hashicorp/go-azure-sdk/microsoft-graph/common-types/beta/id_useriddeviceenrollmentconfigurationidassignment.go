package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDeviceEnrollmentConfigurationIdAssignmentId{}

// UserIdDeviceEnrollmentConfigurationIdAssignmentId is a struct representing the Resource ID for a User Id Device Enrollment Configuration Id Assignment
type UserIdDeviceEnrollmentConfigurationIdAssignmentId struct {
	UserId                              string
	DeviceEnrollmentConfigurationId     string
	EnrollmentConfigurationAssignmentId string
}

// NewUserIdDeviceEnrollmentConfigurationIdAssignmentID returns a new UserIdDeviceEnrollmentConfigurationIdAssignmentId struct
func NewUserIdDeviceEnrollmentConfigurationIdAssignmentID(userId string, deviceEnrollmentConfigurationId string, enrollmentConfigurationAssignmentId string) UserIdDeviceEnrollmentConfigurationIdAssignmentId {
	return UserIdDeviceEnrollmentConfigurationIdAssignmentId{
		UserId:                              userId,
		DeviceEnrollmentConfigurationId:     deviceEnrollmentConfigurationId,
		EnrollmentConfigurationAssignmentId: enrollmentConfigurationAssignmentId,
	}
}

// ParseUserIdDeviceEnrollmentConfigurationIdAssignmentID parses 'input' into a UserIdDeviceEnrollmentConfigurationIdAssignmentId
func ParseUserIdDeviceEnrollmentConfigurationIdAssignmentID(input string) (*UserIdDeviceEnrollmentConfigurationIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceEnrollmentConfigurationIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceEnrollmentConfigurationIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDeviceEnrollmentConfigurationIdAssignmentIDInsensitively parses 'input' case-insensitively into a UserIdDeviceEnrollmentConfigurationIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdDeviceEnrollmentConfigurationIdAssignmentIDInsensitively(input string) (*UserIdDeviceEnrollmentConfigurationIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceEnrollmentConfigurationIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceEnrollmentConfigurationIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDeviceEnrollmentConfigurationIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DeviceEnrollmentConfigurationId, ok = input.Parsed["deviceEnrollmentConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceEnrollmentConfigurationId", input)
	}

	if id.EnrollmentConfigurationAssignmentId, ok = input.Parsed["enrollmentConfigurationAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "enrollmentConfigurationAssignmentId", input)
	}

	return nil
}

// ValidateUserIdDeviceEnrollmentConfigurationIdAssignmentID checks that 'input' can be parsed as a User Id Device Enrollment Configuration Id Assignment ID
func ValidateUserIdDeviceEnrollmentConfigurationIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDeviceEnrollmentConfigurationIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Device Enrollment Configuration Id Assignment ID
func (id UserIdDeviceEnrollmentConfigurationIdAssignmentId) ID() string {
	fmtString := "/users/%s/deviceEnrollmentConfigurations/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceEnrollmentConfigurationId, id.EnrollmentConfigurationAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Device Enrollment Configuration Id Assignment ID
func (id UserIdDeviceEnrollmentConfigurationIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("deviceEnrollmentConfigurations", "deviceEnrollmentConfigurations", "deviceEnrollmentConfigurations"),
		resourceids.UserSpecifiedSegment("deviceEnrollmentConfigurationId", "deviceEnrollmentConfigurationId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("enrollmentConfigurationAssignmentId", "enrollmentConfigurationAssignmentId"),
	}
}

// String returns a human-readable description of this User Id Device Enrollment Configuration Id Assignment ID
func (id UserIdDeviceEnrollmentConfigurationIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device Enrollment Configuration: %q", id.DeviceEnrollmentConfigurationId),
		fmt.Sprintf("Enrollment Configuration Assignment: %q", id.EnrollmentConfigurationAssignmentId),
	}
	return fmt.Sprintf("User Id Device Enrollment Configuration Id Assignment (%s)", strings.Join(components, "\n"))
}
