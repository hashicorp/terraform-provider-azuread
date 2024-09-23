package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId{}

// DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId is a struct representing the Resource ID for a Device Management Apple User Initiated Enrollment Profile Id Assignment
type DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId struct {
	AppleUserInitiatedEnrollmentProfileId string
	AppleEnrollmentProfileAssignmentId    string
}

// NewDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentID returns a new DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId struct
func NewDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentID(appleUserInitiatedEnrollmentProfileId string, appleEnrollmentProfileAssignmentId string) DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId {
	return DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId{
		AppleUserInitiatedEnrollmentProfileId: appleUserInitiatedEnrollmentProfileId,
		AppleEnrollmentProfileAssignmentId:    appleEnrollmentProfileAssignmentId,
	}
}

// ParseDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentID parses 'input' into a DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId
func ParseDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentID(input string) (*DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentIDInsensitively(input string) (*DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AppleUserInitiatedEnrollmentProfileId, ok = input.Parsed["appleUserInitiatedEnrollmentProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appleUserInitiatedEnrollmentProfileId", input)
	}

	if id.AppleEnrollmentProfileAssignmentId, ok = input.Parsed["appleEnrollmentProfileAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appleEnrollmentProfileAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentID checks that 'input' can be parsed as a Device Management Apple User Initiated Enrollment Profile Id Assignment ID
func ValidateDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Apple User Initiated Enrollment Profile Id Assignment ID
func (id DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/appleUserInitiatedEnrollmentProfiles/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.AppleUserInitiatedEnrollmentProfileId, id.AppleEnrollmentProfileAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Apple User Initiated Enrollment Profile Id Assignment ID
func (id DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("appleUserInitiatedEnrollmentProfiles", "appleUserInitiatedEnrollmentProfiles", "appleUserInitiatedEnrollmentProfiles"),
		resourceids.UserSpecifiedSegment("appleUserInitiatedEnrollmentProfileId", "appleUserInitiatedEnrollmentProfileId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("appleEnrollmentProfileAssignmentId", "appleEnrollmentProfileAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Apple User Initiated Enrollment Profile Id Assignment ID
func (id DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Apple User Initiated Enrollment Profile: %q", id.AppleUserInitiatedEnrollmentProfileId),
		fmt.Sprintf("Apple Enrollment Profile Assignment: %q", id.AppleEnrollmentProfileAssignmentId),
	}
	return fmt.Sprintf("Device Management Apple User Initiated Enrollment Profile Id Assignment (%s)", strings.Join(components, "\n"))
}
