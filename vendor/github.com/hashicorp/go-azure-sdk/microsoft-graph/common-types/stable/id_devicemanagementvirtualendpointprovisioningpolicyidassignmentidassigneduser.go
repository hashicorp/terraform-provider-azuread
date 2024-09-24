package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId{}

// DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId is a struct representing the Resource ID for a Device Management Virtual Endpoint Provisioning Policy Id Assignment Id Assigned User
type DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId struct {
	CloudPCProvisioningPolicyId           string
	CloudPCProvisioningPolicyAssignmentId string
	UserId                                string
}

// NewDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserID returns a new DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId struct
func NewDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserID(cloudPCProvisioningPolicyId string, cloudPCProvisioningPolicyAssignmentId string, userId string) DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId {
	return DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId{
		CloudPCProvisioningPolicyId:           cloudPCProvisioningPolicyId,
		CloudPCProvisioningPolicyAssignmentId: cloudPCProvisioningPolicyAssignmentId,
		UserId:                                userId,
	}
}

// ParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserID parses 'input' into a DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId
func ParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserID(input string) (*DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserIDInsensitively(input string) (*DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCProvisioningPolicyId, ok = input.Parsed["cloudPCProvisioningPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCProvisioningPolicyId", input)
	}

	if id.CloudPCProvisioningPolicyAssignmentId, ok = input.Parsed["cloudPCProvisioningPolicyAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCProvisioningPolicyAssignmentId", input)
	}

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserID checks that 'input' can be parsed as a Device Management Virtual Endpoint Provisioning Policy Id Assignment Id Assigned User ID
func ValidateDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint Provisioning Policy Id Assignment Id Assigned User ID
func (id DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/provisioningPolicies/%s/assignments/%s/assignedUsers/%s"
	return fmt.Sprintf(fmtString, id.CloudPCProvisioningPolicyId, id.CloudPCProvisioningPolicyAssignmentId, id.UserId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint Provisioning Policy Id Assignment Id Assigned User ID
func (id DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("provisioningPolicies", "provisioningPolicies", "provisioningPolicies"),
		resourceids.UserSpecifiedSegment("cloudPCProvisioningPolicyId", "cloudPCProvisioningPolicyId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("cloudPCProvisioningPolicyAssignmentId", "cloudPCProvisioningPolicyAssignmentId"),
		resourceids.StaticSegment("assignedUsers", "assignedUsers", "assignedUsers"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint Provisioning Policy Id Assignment Id Assigned User ID
func (id DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId) String() string {
	components := []string{
		fmt.Sprintf("Cloud P C Provisioning Policy: %q", id.CloudPCProvisioningPolicyId),
		fmt.Sprintf("Cloud P C Provisioning Policy Assignment: %q", id.CloudPCProvisioningPolicyAssignmentId),
		fmt.Sprintf("User: %q", id.UserId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint Provisioning Policy Id Assignment Id Assigned User (%s)", strings.Join(components, "\n"))
}
