package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCloudCertificationAuthorityId{}

// DeviceManagementCloudCertificationAuthorityId is a struct representing the Resource ID for a Device Management Cloud Certification Authority
type DeviceManagementCloudCertificationAuthorityId struct {
	CloudCertificationAuthorityId string
}

// NewDeviceManagementCloudCertificationAuthorityID returns a new DeviceManagementCloudCertificationAuthorityId struct
func NewDeviceManagementCloudCertificationAuthorityID(cloudCertificationAuthorityId string) DeviceManagementCloudCertificationAuthorityId {
	return DeviceManagementCloudCertificationAuthorityId{
		CloudCertificationAuthorityId: cloudCertificationAuthorityId,
	}
}

// ParseDeviceManagementCloudCertificationAuthorityID parses 'input' into a DeviceManagementCloudCertificationAuthorityId
func ParseDeviceManagementCloudCertificationAuthorityID(input string) (*DeviceManagementCloudCertificationAuthorityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCloudCertificationAuthorityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCloudCertificationAuthorityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementCloudCertificationAuthorityIDInsensitively parses 'input' case-insensitively into a DeviceManagementCloudCertificationAuthorityId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementCloudCertificationAuthorityIDInsensitively(input string) (*DeviceManagementCloudCertificationAuthorityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCloudCertificationAuthorityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCloudCertificationAuthorityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementCloudCertificationAuthorityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudCertificationAuthorityId, ok = input.Parsed["cloudCertificationAuthorityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudCertificationAuthorityId", input)
	}

	return nil
}

// ValidateDeviceManagementCloudCertificationAuthorityID checks that 'input' can be parsed as a Device Management Cloud Certification Authority ID
func ValidateDeviceManagementCloudCertificationAuthorityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementCloudCertificationAuthorityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Cloud Certification Authority ID
func (id DeviceManagementCloudCertificationAuthorityId) ID() string {
	fmtString := "/deviceManagement/cloudCertificationAuthority/%s"
	return fmt.Sprintf(fmtString, id.CloudCertificationAuthorityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Cloud Certification Authority ID
func (id DeviceManagementCloudCertificationAuthorityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("cloudCertificationAuthority", "cloudCertificationAuthority", "cloudCertificationAuthority"),
		resourceids.UserSpecifiedSegment("cloudCertificationAuthorityId", "cloudCertificationAuthorityId"),
	}
}

// String returns a human-readable description of this Device Management Cloud Certification Authority ID
func (id DeviceManagementCloudCertificationAuthorityId) String() string {
	components := []string{
		fmt.Sprintf("Cloud Certification Authority: %q", id.CloudCertificationAuthorityId),
	}
	return fmt.Sprintf("Device Management Cloud Certification Authority (%s)", strings.Join(components, "\n"))
}
