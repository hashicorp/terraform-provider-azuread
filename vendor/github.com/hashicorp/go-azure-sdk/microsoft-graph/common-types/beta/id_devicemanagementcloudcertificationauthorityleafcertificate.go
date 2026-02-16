package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCloudCertificationAuthorityLeafCertificateId{}

// DeviceManagementCloudCertificationAuthorityLeafCertificateId is a struct representing the Resource ID for a Device Management Cloud Certification Authority Leaf Certificate
type DeviceManagementCloudCertificationAuthorityLeafCertificateId struct {
	CloudCertificationAuthorityLeafCertificateId string
}

// NewDeviceManagementCloudCertificationAuthorityLeafCertificateID returns a new DeviceManagementCloudCertificationAuthorityLeafCertificateId struct
func NewDeviceManagementCloudCertificationAuthorityLeafCertificateID(cloudCertificationAuthorityLeafCertificateId string) DeviceManagementCloudCertificationAuthorityLeafCertificateId {
	return DeviceManagementCloudCertificationAuthorityLeafCertificateId{
		CloudCertificationAuthorityLeafCertificateId: cloudCertificationAuthorityLeafCertificateId,
	}
}

// ParseDeviceManagementCloudCertificationAuthorityLeafCertificateID parses 'input' into a DeviceManagementCloudCertificationAuthorityLeafCertificateId
func ParseDeviceManagementCloudCertificationAuthorityLeafCertificateID(input string) (*DeviceManagementCloudCertificationAuthorityLeafCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCloudCertificationAuthorityLeafCertificateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCloudCertificationAuthorityLeafCertificateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementCloudCertificationAuthorityLeafCertificateIDInsensitively parses 'input' case-insensitively into a DeviceManagementCloudCertificationAuthorityLeafCertificateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementCloudCertificationAuthorityLeafCertificateIDInsensitively(input string) (*DeviceManagementCloudCertificationAuthorityLeafCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCloudCertificationAuthorityLeafCertificateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCloudCertificationAuthorityLeafCertificateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementCloudCertificationAuthorityLeafCertificateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudCertificationAuthorityLeafCertificateId, ok = input.Parsed["cloudCertificationAuthorityLeafCertificateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudCertificationAuthorityLeafCertificateId", input)
	}

	return nil
}

// ValidateDeviceManagementCloudCertificationAuthorityLeafCertificateID checks that 'input' can be parsed as a Device Management Cloud Certification Authority Leaf Certificate ID
func ValidateDeviceManagementCloudCertificationAuthorityLeafCertificateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementCloudCertificationAuthorityLeafCertificateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Cloud Certification Authority Leaf Certificate ID
func (id DeviceManagementCloudCertificationAuthorityLeafCertificateId) ID() string {
	fmtString := "/deviceManagement/cloudCertificationAuthorityLeafCertificate/%s"
	return fmt.Sprintf(fmtString, id.CloudCertificationAuthorityLeafCertificateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Cloud Certification Authority Leaf Certificate ID
func (id DeviceManagementCloudCertificationAuthorityLeafCertificateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("cloudCertificationAuthorityLeafCertificate", "cloudCertificationAuthorityLeafCertificate", "cloudCertificationAuthorityLeafCertificate"),
		resourceids.UserSpecifiedSegment("cloudCertificationAuthorityLeafCertificateId", "cloudCertificationAuthorityLeafCertificateId"),
	}
}

// String returns a human-readable description of this Device Management Cloud Certification Authority Leaf Certificate ID
func (id DeviceManagementCloudCertificationAuthorityLeafCertificateId) String() string {
	components := []string{
		fmt.Sprintf("Cloud Certification Authority Leaf Certificate: %q", id.CloudCertificationAuthorityLeafCertificateId),
	}
	return fmt.Sprintf("Device Management Cloud Certification Authority Leaf Certificate (%s)", strings.Join(components, "\n"))
}
