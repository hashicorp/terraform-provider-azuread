package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId{}

// DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId is a struct representing the Resource ID for a Device Management Mobile App Troubleshooting Event Id App Log Collection Request
type DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId struct {
	MobileAppTroubleshootingEventId string
	AppLogCollectionRequestId       string
}

// NewDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestID returns a new DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId struct
func NewDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestID(mobileAppTroubleshootingEventId string, appLogCollectionRequestId string) DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId {
	return DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId{
		MobileAppTroubleshootingEventId: mobileAppTroubleshootingEventId,
		AppLogCollectionRequestId:       appLogCollectionRequestId,
	}
}

// ParseDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestID parses 'input' into a DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId
func ParseDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestID(input string) (*DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestIDInsensitively parses 'input' case-insensitively into a DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestIDInsensitively(input string) (*DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MobileAppTroubleshootingEventId, ok = input.Parsed["mobileAppTroubleshootingEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mobileAppTroubleshootingEventId", input)
	}

	if id.AppLogCollectionRequestId, ok = input.Parsed["appLogCollectionRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appLogCollectionRequestId", input)
	}

	return nil
}

// ValidateDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestID checks that 'input' can be parsed as a Device Management Mobile App Troubleshooting Event Id App Log Collection Request ID
func ValidateDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Mobile App Troubleshooting Event Id App Log Collection Request ID
func (id DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId) ID() string {
	fmtString := "/deviceManagement/mobileAppTroubleshootingEvents/%s/appLogCollectionRequests/%s"
	return fmt.Sprintf(fmtString, id.MobileAppTroubleshootingEventId, id.AppLogCollectionRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Mobile App Troubleshooting Event Id App Log Collection Request ID
func (id DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("mobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents"),
		resourceids.UserSpecifiedSegment("mobileAppTroubleshootingEventId", "mobileAppTroubleshootingEventId"),
		resourceids.StaticSegment("appLogCollectionRequests", "appLogCollectionRequests", "appLogCollectionRequests"),
		resourceids.UserSpecifiedSegment("appLogCollectionRequestId", "appLogCollectionRequestId"),
	}
}

// String returns a human-readable description of this Device Management Mobile App Troubleshooting Event Id App Log Collection Request ID
func (id DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId) String() string {
	components := []string{
		fmt.Sprintf("Mobile App Troubleshooting Event: %q", id.MobileAppTroubleshootingEventId),
		fmt.Sprintf("App Log Collection Request: %q", id.AppLogCollectionRequestId),
	}
	return fmt.Sprintf("Device Management Mobile App Troubleshooting Event Id App Log Collection Request (%s)", strings.Join(components, "\n"))
}
