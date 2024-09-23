package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementExchangeOnPremisesPolicyId{}

// DeviceManagementExchangeOnPremisesPolicyId is a struct representing the Resource ID for a Device Management Exchange On Premises Policy
type DeviceManagementExchangeOnPremisesPolicyId struct {
	DeviceManagementExchangeOnPremisesPolicyId string
}

// NewDeviceManagementExchangeOnPremisesPolicyID returns a new DeviceManagementExchangeOnPremisesPolicyId struct
func NewDeviceManagementExchangeOnPremisesPolicyID(deviceManagementExchangeOnPremisesPolicyId string) DeviceManagementExchangeOnPremisesPolicyId {
	return DeviceManagementExchangeOnPremisesPolicyId{
		DeviceManagementExchangeOnPremisesPolicyId: deviceManagementExchangeOnPremisesPolicyId,
	}
}

// ParseDeviceManagementExchangeOnPremisesPolicyID parses 'input' into a DeviceManagementExchangeOnPremisesPolicyId
func ParseDeviceManagementExchangeOnPremisesPolicyID(input string) (*DeviceManagementExchangeOnPremisesPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementExchangeOnPremisesPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementExchangeOnPremisesPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementExchangeOnPremisesPolicyIDInsensitively parses 'input' case-insensitively into a DeviceManagementExchangeOnPremisesPolicyId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementExchangeOnPremisesPolicyIDInsensitively(input string) (*DeviceManagementExchangeOnPremisesPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementExchangeOnPremisesPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementExchangeOnPremisesPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementExchangeOnPremisesPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementExchangeOnPremisesPolicyId, ok = input.Parsed["deviceManagementExchangeOnPremisesPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementExchangeOnPremisesPolicyId", input)
	}

	return nil
}

// ValidateDeviceManagementExchangeOnPremisesPolicyID checks that 'input' can be parsed as a Device Management Exchange On Premises Policy ID
func ValidateDeviceManagementExchangeOnPremisesPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementExchangeOnPremisesPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Exchange On Premises Policy ID
func (id DeviceManagementExchangeOnPremisesPolicyId) ID() string {
	fmtString := "/deviceManagement/exchangeOnPremisesPolicies/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementExchangeOnPremisesPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Exchange On Premises Policy ID
func (id DeviceManagementExchangeOnPremisesPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("exchangeOnPremisesPolicies", "exchangeOnPremisesPolicies", "exchangeOnPremisesPolicies"),
		resourceids.UserSpecifiedSegment("deviceManagementExchangeOnPremisesPolicyId", "deviceManagementExchangeOnPremisesPolicyId"),
	}
}

// String returns a human-readable description of this Device Management Exchange On Premises Policy ID
func (id DeviceManagementExchangeOnPremisesPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Exchange On Premises Policy: %q", id.DeviceManagementExchangeOnPremisesPolicyId),
	}
	return fmt.Sprintf("Device Management Exchange On Premises Policy (%s)", strings.Join(components, "\n"))
}
