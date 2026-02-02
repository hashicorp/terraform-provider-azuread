package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointGalleryImageId{}

// DeviceManagementVirtualEndpointGalleryImageId is a struct representing the Resource ID for a Device Management Virtual Endpoint Gallery Image
type DeviceManagementVirtualEndpointGalleryImageId struct {
	CloudPCGalleryImageId string
}

// NewDeviceManagementVirtualEndpointGalleryImageID returns a new DeviceManagementVirtualEndpointGalleryImageId struct
func NewDeviceManagementVirtualEndpointGalleryImageID(cloudPCGalleryImageId string) DeviceManagementVirtualEndpointGalleryImageId {
	return DeviceManagementVirtualEndpointGalleryImageId{
		CloudPCGalleryImageId: cloudPCGalleryImageId,
	}
}

// ParseDeviceManagementVirtualEndpointGalleryImageID parses 'input' into a DeviceManagementVirtualEndpointGalleryImageId
func ParseDeviceManagementVirtualEndpointGalleryImageID(input string) (*DeviceManagementVirtualEndpointGalleryImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointGalleryImageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointGalleryImageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointGalleryImageIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointGalleryImageId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointGalleryImageIDInsensitively(input string) (*DeviceManagementVirtualEndpointGalleryImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointGalleryImageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointGalleryImageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointGalleryImageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCGalleryImageId, ok = input.Parsed["cloudPCGalleryImageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCGalleryImageId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointGalleryImageID checks that 'input' can be parsed as a Device Management Virtual Endpoint Gallery Image ID
func ValidateDeviceManagementVirtualEndpointGalleryImageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointGalleryImageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint Gallery Image ID
func (id DeviceManagementVirtualEndpointGalleryImageId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/galleryImages/%s"
	return fmt.Sprintf(fmtString, id.CloudPCGalleryImageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint Gallery Image ID
func (id DeviceManagementVirtualEndpointGalleryImageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("galleryImages", "galleryImages", "galleryImages"),
		resourceids.UserSpecifiedSegment("cloudPCGalleryImageId", "cloudPCGalleryImageId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint Gallery Image ID
func (id DeviceManagementVirtualEndpointGalleryImageId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC Gallery Image: %q", id.CloudPCGalleryImageId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint Gallery Image (%s)", strings.Join(components, "\n"))
}
