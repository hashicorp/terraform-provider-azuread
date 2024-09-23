package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCompliancePolicyId{}

// DeviceManagementCompliancePolicyId is a struct representing the Resource ID for a Device Management Compliance Policy
type DeviceManagementCompliancePolicyId struct {
	DeviceManagementCompliancePolicyId string
}

// NewDeviceManagementCompliancePolicyID returns a new DeviceManagementCompliancePolicyId struct
func NewDeviceManagementCompliancePolicyID(deviceManagementCompliancePolicyId string) DeviceManagementCompliancePolicyId {
	return DeviceManagementCompliancePolicyId{
		DeviceManagementCompliancePolicyId: deviceManagementCompliancePolicyId,
	}
}

// ParseDeviceManagementCompliancePolicyID parses 'input' into a DeviceManagementCompliancePolicyId
func ParseDeviceManagementCompliancePolicyID(input string) (*DeviceManagementCompliancePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCompliancePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCompliancePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementCompliancePolicyIDInsensitively parses 'input' case-insensitively into a DeviceManagementCompliancePolicyId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementCompliancePolicyIDInsensitively(input string) (*DeviceManagementCompliancePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCompliancePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCompliancePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementCompliancePolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementCompliancePolicyId, ok = input.Parsed["deviceManagementCompliancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementCompliancePolicyId", input)
	}

	return nil
}

// ValidateDeviceManagementCompliancePolicyID checks that 'input' can be parsed as a Device Management Compliance Policy ID
func ValidateDeviceManagementCompliancePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementCompliancePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Compliance Policy ID
func (id DeviceManagementCompliancePolicyId) ID() string {
	fmtString := "/deviceManagement/compliancePolicies/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementCompliancePolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Compliance Policy ID
func (id DeviceManagementCompliancePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("compliancePolicies", "compliancePolicies", "compliancePolicies"),
		resourceids.UserSpecifiedSegment("deviceManagementCompliancePolicyId", "deviceManagementCompliancePolicyId"),
	}
}

// String returns a human-readable description of this Device Management Compliance Policy ID
func (id DeviceManagementCompliancePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Compliance Policy: %q", id.DeviceManagementCompliancePolicyId),
	}
	return fmt.Sprintf("Device Management Compliance Policy (%s)", strings.Join(components, "\n"))
}
