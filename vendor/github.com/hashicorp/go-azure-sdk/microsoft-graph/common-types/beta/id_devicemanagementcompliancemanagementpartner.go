package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComplianceManagementPartnerId{}

// DeviceManagementComplianceManagementPartnerId is a struct representing the Resource ID for a Device Management Compliance Management Partner
type DeviceManagementComplianceManagementPartnerId struct {
	ComplianceManagementPartnerId string
}

// NewDeviceManagementComplianceManagementPartnerID returns a new DeviceManagementComplianceManagementPartnerId struct
func NewDeviceManagementComplianceManagementPartnerID(complianceManagementPartnerId string) DeviceManagementComplianceManagementPartnerId {
	return DeviceManagementComplianceManagementPartnerId{
		ComplianceManagementPartnerId: complianceManagementPartnerId,
	}
}

// ParseDeviceManagementComplianceManagementPartnerID parses 'input' into a DeviceManagementComplianceManagementPartnerId
func ParseDeviceManagementComplianceManagementPartnerID(input string) (*DeviceManagementComplianceManagementPartnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComplianceManagementPartnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComplianceManagementPartnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementComplianceManagementPartnerIDInsensitively parses 'input' case-insensitively into a DeviceManagementComplianceManagementPartnerId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementComplianceManagementPartnerIDInsensitively(input string) (*DeviceManagementComplianceManagementPartnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComplianceManagementPartnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComplianceManagementPartnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementComplianceManagementPartnerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ComplianceManagementPartnerId, ok = input.Parsed["complianceManagementPartnerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "complianceManagementPartnerId", input)
	}

	return nil
}

// ValidateDeviceManagementComplianceManagementPartnerID checks that 'input' can be parsed as a Device Management Compliance Management Partner ID
func ValidateDeviceManagementComplianceManagementPartnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementComplianceManagementPartnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Compliance Management Partner ID
func (id DeviceManagementComplianceManagementPartnerId) ID() string {
	fmtString := "/deviceManagement/complianceManagementPartners/%s"
	return fmt.Sprintf(fmtString, id.ComplianceManagementPartnerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Compliance Management Partner ID
func (id DeviceManagementComplianceManagementPartnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("complianceManagementPartners", "complianceManagementPartners", "complianceManagementPartners"),
		resourceids.UserSpecifiedSegment("complianceManagementPartnerId", "complianceManagementPartnerId"),
	}
}

// String returns a human-readable description of this Device Management Compliance Management Partner ID
func (id DeviceManagementComplianceManagementPartnerId) String() string {
	components := []string{
		fmt.Sprintf("Compliance Management Partner: %q", id.ComplianceManagementPartnerId),
	}
	return fmt.Sprintf("Device Management Compliance Management Partner (%s)", strings.Join(components, "\n"))
}
