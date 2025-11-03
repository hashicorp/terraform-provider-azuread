package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId{}

// DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId is a struct representing the Resource ID for a Device Management Virtual Endpoint Provisioning Policy Id Assignment
type DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId struct {
	CloudPCProvisioningPolicyId           string
	CloudPCProvisioningPolicyAssignmentId string
}

// NewDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentID returns a new DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId struct
func NewDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentID(cloudPCProvisioningPolicyId string, cloudPCProvisioningPolicyAssignmentId string) DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId {
	return DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId{
		CloudPCProvisioningPolicyId:           cloudPCProvisioningPolicyId,
		CloudPCProvisioningPolicyAssignmentId: cloudPCProvisioningPolicyAssignmentId,
	}
}

// ParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentID parses 'input' into a DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId
func ParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentID(input string) (*DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIDInsensitively(input string) (*DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCProvisioningPolicyId, ok = input.Parsed["cloudPCProvisioningPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCProvisioningPolicyId", input)
	}

	if id.CloudPCProvisioningPolicyAssignmentId, ok = input.Parsed["cloudPCProvisioningPolicyAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCProvisioningPolicyAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentID checks that 'input' can be parsed as a Device Management Virtual Endpoint Provisioning Policy Id Assignment ID
func ValidateDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint Provisioning Policy Id Assignment ID
func (id DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/provisioningPolicies/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.CloudPCProvisioningPolicyId, id.CloudPCProvisioningPolicyAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint Provisioning Policy Id Assignment ID
func (id DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("provisioningPolicies", "provisioningPolicies", "provisioningPolicies"),
		resourceids.UserSpecifiedSegment("cloudPCProvisioningPolicyId", "cloudPCProvisioningPolicyId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("cloudPCProvisioningPolicyAssignmentId", "cloudPCProvisioningPolicyAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint Provisioning Policy Id Assignment ID
func (id DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC Provisioning Policy: %q", id.CloudPCProvisioningPolicyId),
		fmt.Sprintf("Cloud PC Provisioning Policy Assignment: %q", id.CloudPCProvisioningPolicyAssignmentId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint Provisioning Policy Id Assignment (%s)", strings.Join(components, "\n"))
}
