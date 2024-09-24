package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsImpactingProcessId{}

// DeviceManagementUserExperienceAnalyticsImpactingProcessId is a struct representing the Resource ID for a Device Management User Experience Analytics Impacting Process
type DeviceManagementUserExperienceAnalyticsImpactingProcessId struct {
	UserExperienceAnalyticsImpactingProcessId string
}

// NewDeviceManagementUserExperienceAnalyticsImpactingProcessID returns a new DeviceManagementUserExperienceAnalyticsImpactingProcessId struct
func NewDeviceManagementUserExperienceAnalyticsImpactingProcessID(userExperienceAnalyticsImpactingProcessId string) DeviceManagementUserExperienceAnalyticsImpactingProcessId {
	return DeviceManagementUserExperienceAnalyticsImpactingProcessId{
		UserExperienceAnalyticsImpactingProcessId: userExperienceAnalyticsImpactingProcessId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsImpactingProcessID parses 'input' into a DeviceManagementUserExperienceAnalyticsImpactingProcessId
func ParseDeviceManagementUserExperienceAnalyticsImpactingProcessID(input string) (*DeviceManagementUserExperienceAnalyticsImpactingProcessId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsImpactingProcessId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsImpactingProcessId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsImpactingProcessIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsImpactingProcessId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsImpactingProcessIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsImpactingProcessId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsImpactingProcessId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsImpactingProcessId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsImpactingProcessId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsImpactingProcessId, ok = input.Parsed["userExperienceAnalyticsImpactingProcessId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsImpactingProcessId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsImpactingProcessID checks that 'input' can be parsed as a Device Management User Experience Analytics Impacting Process ID
func ValidateDeviceManagementUserExperienceAnalyticsImpactingProcessID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsImpactingProcessID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Impacting Process ID
func (id DeviceManagementUserExperienceAnalyticsImpactingProcessId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsImpactingProcess/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsImpactingProcessId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Impacting Process ID
func (id DeviceManagementUserExperienceAnalyticsImpactingProcessId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsImpactingProcess", "userExperienceAnalyticsImpactingProcess", "userExperienceAnalyticsImpactingProcess"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsImpactingProcessId", "userExperienceAnalyticsImpactingProcessId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Impacting Process ID
func (id DeviceManagementUserExperienceAnalyticsImpactingProcessId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Impacting Process: %q", id.UserExperienceAnalyticsImpactingProcessId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Impacting Process (%s)", strings.Join(components, "\n"))
}
