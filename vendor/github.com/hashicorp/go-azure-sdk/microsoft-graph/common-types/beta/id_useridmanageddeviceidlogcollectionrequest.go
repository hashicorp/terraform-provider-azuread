package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedDeviceIdLogCollectionRequestId{}

// UserIdManagedDeviceIdLogCollectionRequestId is a struct representing the Resource ID for a User Id Managed Device Id Log Collection Request
type UserIdManagedDeviceIdLogCollectionRequestId struct {
	UserId                        string
	ManagedDeviceId               string
	DeviceLogCollectionResponseId string
}

// NewUserIdManagedDeviceIdLogCollectionRequestID returns a new UserIdManagedDeviceIdLogCollectionRequestId struct
func NewUserIdManagedDeviceIdLogCollectionRequestID(userId string, managedDeviceId string, deviceLogCollectionResponseId string) UserIdManagedDeviceIdLogCollectionRequestId {
	return UserIdManagedDeviceIdLogCollectionRequestId{
		UserId:                        userId,
		ManagedDeviceId:               managedDeviceId,
		DeviceLogCollectionResponseId: deviceLogCollectionResponseId,
	}
}

// ParseUserIdManagedDeviceIdLogCollectionRequestID parses 'input' into a UserIdManagedDeviceIdLogCollectionRequestId
func ParseUserIdManagedDeviceIdLogCollectionRequestID(input string) (*UserIdManagedDeviceIdLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdLogCollectionRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdManagedDeviceIdLogCollectionRequestIDInsensitively parses 'input' case-insensitively into a UserIdManagedDeviceIdLogCollectionRequestId
// note: this method should only be used for API response data and not user input
func ParseUserIdManagedDeviceIdLogCollectionRequestIDInsensitively(input string) (*UserIdManagedDeviceIdLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdLogCollectionRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdManagedDeviceIdLogCollectionRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DeviceLogCollectionResponseId, ok = input.Parsed["deviceLogCollectionResponseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceLogCollectionResponseId", input)
	}

	return nil
}

// ValidateUserIdManagedDeviceIdLogCollectionRequestID checks that 'input' can be parsed as a User Id Managed Device Id Log Collection Request ID
func ValidateUserIdManagedDeviceIdLogCollectionRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdManagedDeviceIdLogCollectionRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Managed Device Id Log Collection Request ID
func (id UserIdManagedDeviceIdLogCollectionRequestId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/logCollectionRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.DeviceLogCollectionResponseId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Managed Device Id Log Collection Request ID
func (id UserIdManagedDeviceIdLogCollectionRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("logCollectionRequests", "logCollectionRequests", "logCollectionRequests"),
		resourceids.UserSpecifiedSegment("deviceLogCollectionResponseId", "deviceLogCollectionResponseId"),
	}
}

// String returns a human-readable description of this User Id Managed Device Id Log Collection Request ID
func (id UserIdManagedDeviceIdLogCollectionRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Log Collection Response: %q", id.DeviceLogCollectionResponseId),
	}
	return fmt.Sprintf("User Id Managed Device Id Log Collection Request (%s)", strings.Join(components, "\n"))
}
