package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTelecomExpenseManagementPartnerId{}

// DeviceManagementTelecomExpenseManagementPartnerId is a struct representing the Resource ID for a Device Management Telecom Expense Management Partner
type DeviceManagementTelecomExpenseManagementPartnerId struct {
	TelecomExpenseManagementPartnerId string
}

// NewDeviceManagementTelecomExpenseManagementPartnerID returns a new DeviceManagementTelecomExpenseManagementPartnerId struct
func NewDeviceManagementTelecomExpenseManagementPartnerID(telecomExpenseManagementPartnerId string) DeviceManagementTelecomExpenseManagementPartnerId {
	return DeviceManagementTelecomExpenseManagementPartnerId{
		TelecomExpenseManagementPartnerId: telecomExpenseManagementPartnerId,
	}
}

// ParseDeviceManagementTelecomExpenseManagementPartnerID parses 'input' into a DeviceManagementTelecomExpenseManagementPartnerId
func ParseDeviceManagementTelecomExpenseManagementPartnerID(input string) (*DeviceManagementTelecomExpenseManagementPartnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTelecomExpenseManagementPartnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTelecomExpenseManagementPartnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTelecomExpenseManagementPartnerIDInsensitively parses 'input' case-insensitively into a DeviceManagementTelecomExpenseManagementPartnerId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTelecomExpenseManagementPartnerIDInsensitively(input string) (*DeviceManagementTelecomExpenseManagementPartnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTelecomExpenseManagementPartnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTelecomExpenseManagementPartnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTelecomExpenseManagementPartnerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TelecomExpenseManagementPartnerId, ok = input.Parsed["telecomExpenseManagementPartnerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "telecomExpenseManagementPartnerId", input)
	}

	return nil
}

// ValidateDeviceManagementTelecomExpenseManagementPartnerID checks that 'input' can be parsed as a Device Management Telecom Expense Management Partner ID
func ValidateDeviceManagementTelecomExpenseManagementPartnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTelecomExpenseManagementPartnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Telecom Expense Management Partner ID
func (id DeviceManagementTelecomExpenseManagementPartnerId) ID() string {
	fmtString := "/deviceManagement/telecomExpenseManagementPartners/%s"
	return fmt.Sprintf(fmtString, id.TelecomExpenseManagementPartnerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Telecom Expense Management Partner ID
func (id DeviceManagementTelecomExpenseManagementPartnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("telecomExpenseManagementPartners", "telecomExpenseManagementPartners", "telecomExpenseManagementPartners"),
		resourceids.UserSpecifiedSegment("telecomExpenseManagementPartnerId", "telecomExpenseManagementPartnerId"),
	}
}

// String returns a human-readable description of this Device Management Telecom Expense Management Partner ID
func (id DeviceManagementTelecomExpenseManagementPartnerId) String() string {
	components := []string{
		fmt.Sprintf("Telecom Expense Management Partner: %q", id.TelecomExpenseManagementPartnerId),
	}
	return fmt.Sprintf("Device Management Telecom Expense Management Partner (%s)", strings.Join(components, "\n"))
}
