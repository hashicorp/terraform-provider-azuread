package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId{}

// DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId is a struct representing the Resource ID for a Device Management User Experience Analytics Devices Without Cloud Identity
type DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId struct {
	UserExperienceAnalyticsDeviceWithoutCloudIdentityId string
}

// NewDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityID returns a new DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId struct
func NewDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityID(userExperienceAnalyticsDeviceWithoutCloudIdentityId string) DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId {
	return DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId{
		UserExperienceAnalyticsDeviceWithoutCloudIdentityId: userExperienceAnalyticsDeviceWithoutCloudIdentityId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityID parses 'input' into a DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId
func ParseDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityID(input string) (*DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsDeviceWithoutCloudIdentityId, ok = input.Parsed["userExperienceAnalyticsDeviceWithoutCloudIdentityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsDeviceWithoutCloudIdentityId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityID checks that 'input' can be parsed as a Device Management User Experience Analytics Devices Without Cloud Identity ID
func ValidateDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Devices Without Cloud Identity ID
func (id DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsDevicesWithoutCloudIdentity/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsDeviceWithoutCloudIdentityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Devices Without Cloud Identity ID
func (id DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsDevicesWithoutCloudIdentity", "userExperienceAnalyticsDevicesWithoutCloudIdentity", "userExperienceAnalyticsDevicesWithoutCloudIdentity"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsDeviceWithoutCloudIdentityId", "userExperienceAnalyticsDeviceWithoutCloudIdentityId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Devices Without Cloud Identity ID
func (id DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Device Without Cloud Identity: %q", id.UserExperienceAnalyticsDeviceWithoutCloudIdentityId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Devices Without Cloud Identity (%s)", strings.Join(components, "\n"))
}
