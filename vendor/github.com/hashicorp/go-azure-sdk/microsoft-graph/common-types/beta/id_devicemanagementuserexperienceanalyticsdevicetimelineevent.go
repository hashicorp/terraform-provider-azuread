package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId{}

// DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId is a struct representing the Resource ID for a Device Management User Experience Analytics Device Timeline Event
type DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId struct {
	UserExperienceAnalyticsDeviceTimelineEventId string
}

// NewDeviceManagementUserExperienceAnalyticsDeviceTimelineEventID returns a new DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId struct
func NewDeviceManagementUserExperienceAnalyticsDeviceTimelineEventID(userExperienceAnalyticsDeviceTimelineEventId string) DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId {
	return DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId{
		UserExperienceAnalyticsDeviceTimelineEventId: userExperienceAnalyticsDeviceTimelineEventId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsDeviceTimelineEventID parses 'input' into a DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId
func ParseDeviceManagementUserExperienceAnalyticsDeviceTimelineEventID(input string) (*DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsDeviceTimelineEventIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsDeviceTimelineEventIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsDeviceTimelineEventId, ok = input.Parsed["userExperienceAnalyticsDeviceTimelineEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsDeviceTimelineEventId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsDeviceTimelineEventID checks that 'input' can be parsed as a Device Management User Experience Analytics Device Timeline Event ID
func ValidateDeviceManagementUserExperienceAnalyticsDeviceTimelineEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsDeviceTimelineEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Device Timeline Event ID
func (id DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsDeviceTimelineEvent/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsDeviceTimelineEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Device Timeline Event ID
func (id DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsDeviceTimelineEvent", "userExperienceAnalyticsDeviceTimelineEvent", "userExperienceAnalyticsDeviceTimelineEvent"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsDeviceTimelineEventId", "userExperienceAnalyticsDeviceTimelineEventId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Device Timeline Event ID
func (id DeviceManagementUserExperienceAnalyticsDeviceTimelineEventId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Device Timeline Event: %q", id.UserExperienceAnalyticsDeviceTimelineEventId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Device Timeline Event (%s)", strings.Join(components, "\n"))
}
