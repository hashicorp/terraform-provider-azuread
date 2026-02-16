package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointProvisioningPolicyId{}

// DeviceManagementVirtualEndpointProvisioningPolicyId is a struct representing the Resource ID for a Device Management Virtual Endpoint Provisioning Policy
type DeviceManagementVirtualEndpointProvisioningPolicyId struct {
	CloudPCProvisioningPolicyId string
}

// NewDeviceManagementVirtualEndpointProvisioningPolicyID returns a new DeviceManagementVirtualEndpointProvisioningPolicyId struct
func NewDeviceManagementVirtualEndpointProvisioningPolicyID(cloudPCProvisioningPolicyId string) DeviceManagementVirtualEndpointProvisioningPolicyId {
	return DeviceManagementVirtualEndpointProvisioningPolicyId{
		CloudPCProvisioningPolicyId: cloudPCProvisioningPolicyId,
	}
}

// ParseDeviceManagementVirtualEndpointProvisioningPolicyID parses 'input' into a DeviceManagementVirtualEndpointProvisioningPolicyId
func ParseDeviceManagementVirtualEndpointProvisioningPolicyID(input string) (*DeviceManagementVirtualEndpointProvisioningPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointProvisioningPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointProvisioningPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointProvisioningPolicyIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointProvisioningPolicyId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointProvisioningPolicyIDInsensitively(input string) (*DeviceManagementVirtualEndpointProvisioningPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointProvisioningPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointProvisioningPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointProvisioningPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCProvisioningPolicyId, ok = input.Parsed["cloudPCProvisioningPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCProvisioningPolicyId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointProvisioningPolicyID checks that 'input' can be parsed as a Device Management Virtual Endpoint Provisioning Policy ID
func ValidateDeviceManagementVirtualEndpointProvisioningPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointProvisioningPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint Provisioning Policy ID
func (id DeviceManagementVirtualEndpointProvisioningPolicyId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/provisioningPolicies/%s"
	return fmt.Sprintf(fmtString, id.CloudPCProvisioningPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint Provisioning Policy ID
func (id DeviceManagementVirtualEndpointProvisioningPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("provisioningPolicies", "provisioningPolicies", "provisioningPolicies"),
		resourceids.UserSpecifiedSegment("cloudPCProvisioningPolicyId", "cloudPCProvisioningPolicyId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint Provisioning Policy ID
func (id DeviceManagementVirtualEndpointProvisioningPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC Provisioning Policy: %q", id.CloudPCProvisioningPolicyId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint Provisioning Policy (%s)", strings.Join(components, "\n"))
}
