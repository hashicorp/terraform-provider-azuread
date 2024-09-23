package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId{}

// DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId is a struct representing the Resource ID for a Device Management User Experience Analytics App Health Device Performance Detail
type DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId struct {
	UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId string
}

// NewDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailID returns a new DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId struct
func NewDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailID(userExperienceAnalyticsAppHealthDevicePerformanceDetailsId string) DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId {
	return DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId{
		UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId: userExperienceAnalyticsAppHealthDevicePerformanceDetailsId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailID parses 'input' into a DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId
func ParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailID(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId, ok = input.Parsed["userExperienceAnalyticsAppHealthDevicePerformanceDetailsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsAppHealthDevicePerformanceDetailsId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailID checks that 'input' can be parsed as a Device Management User Experience Analytics App Health Device Performance Detail ID
func ValidateDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics App Health Device Performance Detail ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformanceDetails/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics App Health Device Performance Detail ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsAppHealthDevicePerformanceDetails", "userExperienceAnalyticsAppHealthDevicePerformanceDetails", "userExperienceAnalyticsAppHealthDevicePerformanceDetails"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsAppHealthDevicePerformanceDetailsId", "userExperienceAnalyticsAppHealthDevicePerformanceDetailsId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics App Health Device Performance Detail ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics App Health Device Performance Details: %q", id.UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics App Health Device Performance Detail (%s)", strings.Join(components, "\n"))
}
