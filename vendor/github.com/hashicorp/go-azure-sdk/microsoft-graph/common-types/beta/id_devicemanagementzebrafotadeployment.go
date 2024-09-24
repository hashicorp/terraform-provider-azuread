package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementZebraFotaDeploymentId{}

// DeviceManagementZebraFotaDeploymentId is a struct representing the Resource ID for a Device Management Zebra Fota Deployment
type DeviceManagementZebraFotaDeploymentId struct {
	ZebraFotaDeploymentId string
}

// NewDeviceManagementZebraFotaDeploymentID returns a new DeviceManagementZebraFotaDeploymentId struct
func NewDeviceManagementZebraFotaDeploymentID(zebraFotaDeploymentId string) DeviceManagementZebraFotaDeploymentId {
	return DeviceManagementZebraFotaDeploymentId{
		ZebraFotaDeploymentId: zebraFotaDeploymentId,
	}
}

// ParseDeviceManagementZebraFotaDeploymentID parses 'input' into a DeviceManagementZebraFotaDeploymentId
func ParseDeviceManagementZebraFotaDeploymentID(input string) (*DeviceManagementZebraFotaDeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementZebraFotaDeploymentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementZebraFotaDeploymentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementZebraFotaDeploymentIDInsensitively parses 'input' case-insensitively into a DeviceManagementZebraFotaDeploymentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementZebraFotaDeploymentIDInsensitively(input string) (*DeviceManagementZebraFotaDeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementZebraFotaDeploymentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementZebraFotaDeploymentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementZebraFotaDeploymentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ZebraFotaDeploymentId, ok = input.Parsed["zebraFotaDeploymentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "zebraFotaDeploymentId", input)
	}

	return nil
}

// ValidateDeviceManagementZebraFotaDeploymentID checks that 'input' can be parsed as a Device Management Zebra Fota Deployment ID
func ValidateDeviceManagementZebraFotaDeploymentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementZebraFotaDeploymentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Zebra Fota Deployment ID
func (id DeviceManagementZebraFotaDeploymentId) ID() string {
	fmtString := "/deviceManagement/zebraFotaDeployments/%s"
	return fmt.Sprintf(fmtString, id.ZebraFotaDeploymentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Zebra Fota Deployment ID
func (id DeviceManagementZebraFotaDeploymentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("zebraFotaDeployments", "zebraFotaDeployments", "zebraFotaDeployments"),
		resourceids.UserSpecifiedSegment("zebraFotaDeploymentId", "zebraFotaDeploymentId"),
	}
}

// String returns a human-readable description of this Device Management Zebra Fota Deployment ID
func (id DeviceManagementZebraFotaDeploymentId) String() string {
	components := []string{
		fmt.Sprintf("Zebra Fota Deployment: %q", id.ZebraFotaDeploymentId),
	}
	return fmt.Sprintf("Device Management Zebra Fota Deployment (%s)", strings.Join(components, "\n"))
}
