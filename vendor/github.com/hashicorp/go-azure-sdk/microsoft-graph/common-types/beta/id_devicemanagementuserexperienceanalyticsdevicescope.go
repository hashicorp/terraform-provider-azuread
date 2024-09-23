package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDeviceScopeId{}

// DeviceManagementUserExperienceAnalyticsDeviceScopeId is a struct representing the Resource ID for a Device Management User Experience Analytics Device Scope
type DeviceManagementUserExperienceAnalyticsDeviceScopeId struct {
	UserExperienceAnalyticsDeviceScopeId string
}

// NewDeviceManagementUserExperienceAnalyticsDeviceScopeID returns a new DeviceManagementUserExperienceAnalyticsDeviceScopeId struct
func NewDeviceManagementUserExperienceAnalyticsDeviceScopeID(userExperienceAnalyticsDeviceScopeId string) DeviceManagementUserExperienceAnalyticsDeviceScopeId {
	return DeviceManagementUserExperienceAnalyticsDeviceScopeId{
		UserExperienceAnalyticsDeviceScopeId: userExperienceAnalyticsDeviceScopeId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsDeviceScopeID parses 'input' into a DeviceManagementUserExperienceAnalyticsDeviceScopeId
func ParseDeviceManagementUserExperienceAnalyticsDeviceScopeID(input string) (*DeviceManagementUserExperienceAnalyticsDeviceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDeviceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDeviceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsDeviceScopeIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsDeviceScopeId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsDeviceScopeIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsDeviceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDeviceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDeviceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsDeviceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsDeviceScopeId, ok = input.Parsed["userExperienceAnalyticsDeviceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsDeviceScopeId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsDeviceScopeID checks that 'input' can be parsed as a Device Management User Experience Analytics Device Scope ID
func ValidateDeviceManagementUserExperienceAnalyticsDeviceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsDeviceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Device Scope ID
func (id DeviceManagementUserExperienceAnalyticsDeviceScopeId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsDeviceScopes/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsDeviceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Device Scope ID
func (id DeviceManagementUserExperienceAnalyticsDeviceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsDeviceScopes", "userExperienceAnalyticsDeviceScopes", "userExperienceAnalyticsDeviceScopes"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsDeviceScopeId", "userExperienceAnalyticsDeviceScopeId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Device Scope ID
func (id DeviceManagementUserExperienceAnalyticsDeviceScopeId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Device Scope: %q", id.UserExperienceAnalyticsDeviceScopeId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Device Scope (%s)", strings.Join(components, "\n"))
}
