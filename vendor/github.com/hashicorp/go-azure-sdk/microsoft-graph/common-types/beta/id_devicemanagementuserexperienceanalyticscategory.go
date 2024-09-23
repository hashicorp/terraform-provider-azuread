package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsCategoryId{}

// DeviceManagementUserExperienceAnalyticsCategoryId is a struct representing the Resource ID for a Device Management User Experience Analytics Category
type DeviceManagementUserExperienceAnalyticsCategoryId struct {
	UserExperienceAnalyticsCategoryId string
}

// NewDeviceManagementUserExperienceAnalyticsCategoryID returns a new DeviceManagementUserExperienceAnalyticsCategoryId struct
func NewDeviceManagementUserExperienceAnalyticsCategoryID(userExperienceAnalyticsCategoryId string) DeviceManagementUserExperienceAnalyticsCategoryId {
	return DeviceManagementUserExperienceAnalyticsCategoryId{
		UserExperienceAnalyticsCategoryId: userExperienceAnalyticsCategoryId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsCategoryID parses 'input' into a DeviceManagementUserExperienceAnalyticsCategoryId
func ParseDeviceManagementUserExperienceAnalyticsCategoryID(input string) (*DeviceManagementUserExperienceAnalyticsCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsCategoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsCategoryIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsCategoryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsCategoryIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsCategoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsCategoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsCategoryId, ok = input.Parsed["userExperienceAnalyticsCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsCategoryId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsCategoryID checks that 'input' can be parsed as a Device Management User Experience Analytics Category ID
func ValidateDeviceManagementUserExperienceAnalyticsCategoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsCategoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Category ID
func (id DeviceManagementUserExperienceAnalyticsCategoryId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsCategories/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsCategoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Category ID
func (id DeviceManagementUserExperienceAnalyticsCategoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsCategories", "userExperienceAnalyticsCategories", "userExperienceAnalyticsCategories"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsCategoryId", "userExperienceAnalyticsCategoryId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Category ID
func (id DeviceManagementUserExperienceAnalyticsCategoryId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Category: %q", id.UserExperienceAnalyticsCategoryId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Category (%s)", strings.Join(components, "\n"))
}
