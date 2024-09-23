package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementZebraFotaArtifactId{}

// DeviceManagementZebraFotaArtifactId is a struct representing the Resource ID for a Device Management Zebra Fota Artifact
type DeviceManagementZebraFotaArtifactId struct {
	ZebraFotaArtifactId string
}

// NewDeviceManagementZebraFotaArtifactID returns a new DeviceManagementZebraFotaArtifactId struct
func NewDeviceManagementZebraFotaArtifactID(zebraFotaArtifactId string) DeviceManagementZebraFotaArtifactId {
	return DeviceManagementZebraFotaArtifactId{
		ZebraFotaArtifactId: zebraFotaArtifactId,
	}
}

// ParseDeviceManagementZebraFotaArtifactID parses 'input' into a DeviceManagementZebraFotaArtifactId
func ParseDeviceManagementZebraFotaArtifactID(input string) (*DeviceManagementZebraFotaArtifactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementZebraFotaArtifactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementZebraFotaArtifactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementZebraFotaArtifactIDInsensitively parses 'input' case-insensitively into a DeviceManagementZebraFotaArtifactId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementZebraFotaArtifactIDInsensitively(input string) (*DeviceManagementZebraFotaArtifactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementZebraFotaArtifactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementZebraFotaArtifactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementZebraFotaArtifactId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ZebraFotaArtifactId, ok = input.Parsed["zebraFotaArtifactId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "zebraFotaArtifactId", input)
	}

	return nil
}

// ValidateDeviceManagementZebraFotaArtifactID checks that 'input' can be parsed as a Device Management Zebra Fota Artifact ID
func ValidateDeviceManagementZebraFotaArtifactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementZebraFotaArtifactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Zebra Fota Artifact ID
func (id DeviceManagementZebraFotaArtifactId) ID() string {
	fmtString := "/deviceManagement/zebraFotaArtifacts/%s"
	return fmt.Sprintf(fmtString, id.ZebraFotaArtifactId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Zebra Fota Artifact ID
func (id DeviceManagementZebraFotaArtifactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("zebraFotaArtifacts", "zebraFotaArtifacts", "zebraFotaArtifacts"),
		resourceids.UserSpecifiedSegment("zebraFotaArtifactId", "zebraFotaArtifactId"),
	}
}

// String returns a human-readable description of this Device Management Zebra Fota Artifact ID
func (id DeviceManagementZebraFotaArtifactId) String() string {
	components := []string{
		fmt.Sprintf("Zebra Fota Artifact: %q", id.ZebraFotaArtifactId),
	}
	return fmt.Sprintf("Device Management Zebra Fota Artifact (%s)", strings.Join(components, "\n"))
}
