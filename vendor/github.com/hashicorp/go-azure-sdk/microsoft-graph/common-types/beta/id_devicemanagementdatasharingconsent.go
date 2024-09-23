package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDataSharingConsentId{}

// DeviceManagementDataSharingConsentId is a struct representing the Resource ID for a Device Management Data Sharing Consent
type DeviceManagementDataSharingConsentId struct {
	DataSharingConsentId string
}

// NewDeviceManagementDataSharingConsentID returns a new DeviceManagementDataSharingConsentId struct
func NewDeviceManagementDataSharingConsentID(dataSharingConsentId string) DeviceManagementDataSharingConsentId {
	return DeviceManagementDataSharingConsentId{
		DataSharingConsentId: dataSharingConsentId,
	}
}

// ParseDeviceManagementDataSharingConsentID parses 'input' into a DeviceManagementDataSharingConsentId
func ParseDeviceManagementDataSharingConsentID(input string) (*DeviceManagementDataSharingConsentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDataSharingConsentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDataSharingConsentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDataSharingConsentIDInsensitively parses 'input' case-insensitively into a DeviceManagementDataSharingConsentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDataSharingConsentIDInsensitively(input string) (*DeviceManagementDataSharingConsentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDataSharingConsentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDataSharingConsentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDataSharingConsentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DataSharingConsentId, ok = input.Parsed["dataSharingConsentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataSharingConsentId", input)
	}

	return nil
}

// ValidateDeviceManagementDataSharingConsentID checks that 'input' can be parsed as a Device Management Data Sharing Consent ID
func ValidateDeviceManagementDataSharingConsentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDataSharingConsentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Data Sharing Consent ID
func (id DeviceManagementDataSharingConsentId) ID() string {
	fmtString := "/deviceManagement/dataSharingConsents/%s"
	return fmt.Sprintf(fmtString, id.DataSharingConsentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Data Sharing Consent ID
func (id DeviceManagementDataSharingConsentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("dataSharingConsents", "dataSharingConsents", "dataSharingConsents"),
		resourceids.UserSpecifiedSegment("dataSharingConsentId", "dataSharingConsentId"),
	}
}

// String returns a human-readable description of this Device Management Data Sharing Consent ID
func (id DeviceManagementDataSharingConsentId) String() string {
	components := []string{
		fmt.Sprintf("Data Sharing Consent: %q", id.DataSharingConsentId),
	}
	return fmt.Sprintf("Device Management Data Sharing Consent (%s)", strings.Join(components, "\n"))
}
