package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDeviceIdDeviceTemplateId{}

// UserIdDeviceIdDeviceTemplateId is a struct representing the Resource ID for a User Id Device Id Device Template
type UserIdDeviceIdDeviceTemplateId struct {
	UserId           string
	DeviceId         string
	DeviceTemplateId string
}

// NewUserIdDeviceIdDeviceTemplateID returns a new UserIdDeviceIdDeviceTemplateId struct
func NewUserIdDeviceIdDeviceTemplateID(userId string, deviceId string, deviceTemplateId string) UserIdDeviceIdDeviceTemplateId {
	return UserIdDeviceIdDeviceTemplateId{
		UserId:           userId,
		DeviceId:         deviceId,
		DeviceTemplateId: deviceTemplateId,
	}
}

// ParseUserIdDeviceIdDeviceTemplateID parses 'input' into a UserIdDeviceIdDeviceTemplateId
func ParseUserIdDeviceIdDeviceTemplateID(input string) (*UserIdDeviceIdDeviceTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceIdDeviceTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceIdDeviceTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDeviceIdDeviceTemplateIDInsensitively parses 'input' case-insensitively into a UserIdDeviceIdDeviceTemplateId
// note: this method should only be used for API response data and not user input
func ParseUserIdDeviceIdDeviceTemplateIDInsensitively(input string) (*UserIdDeviceIdDeviceTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceIdDeviceTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceIdDeviceTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDeviceIdDeviceTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.DeviceTemplateId, ok = input.Parsed["deviceTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceTemplateId", input)
	}

	return nil
}

// ValidateUserIdDeviceIdDeviceTemplateID checks that 'input' can be parsed as a User Id Device Id Device Template ID
func ValidateUserIdDeviceIdDeviceTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDeviceIdDeviceTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Device Id Device Template ID
func (id UserIdDeviceIdDeviceTemplateId) ID() string {
	fmtString := "/users/%s/devices/%s/deviceTemplate/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceId, id.DeviceTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Device Id Device Template ID
func (id UserIdDeviceIdDeviceTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("deviceTemplate", "deviceTemplate", "deviceTemplate"),
		resourceids.UserSpecifiedSegment("deviceTemplateId", "deviceTemplateId"),
	}
}

// String returns a human-readable description of this User Id Device Id Device Template ID
func (id UserIdDeviceIdDeviceTemplateId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Device Template: %q", id.DeviceTemplateId),
	}
	return fmt.Sprintf("User Id Device Id Device Template (%s)", strings.Join(components, "\n"))
}
