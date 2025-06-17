package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId{}

// DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId is a struct representing the Resource ID for a Device Management Cloud Certification Authority Id Cloud Certification Authority Leaf Certificate
type DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId struct {
	CloudCertificationAuthorityId                string
	CloudCertificationAuthorityLeafCertificateId string
}

// NewDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateID returns a new DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId struct
func NewDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateID(cloudCertificationAuthorityId string, cloudCertificationAuthorityLeafCertificateId string) DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId {
	return DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId{
		CloudCertificationAuthorityId:                cloudCertificationAuthorityId,
		CloudCertificationAuthorityLeafCertificateId: cloudCertificationAuthorityLeafCertificateId,
	}
}

// ParseDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateID parses 'input' into a DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId
func ParseDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateID(input string) (*DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateIDInsensitively parses 'input' case-insensitively into a DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateIDInsensitively(input string) (*DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudCertificationAuthorityId, ok = input.Parsed["cloudCertificationAuthorityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudCertificationAuthorityId", input)
	}

	if id.CloudCertificationAuthorityLeafCertificateId, ok = input.Parsed["cloudCertificationAuthorityLeafCertificateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudCertificationAuthorityLeafCertificateId", input)
	}

	return nil
}

// ValidateDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateID checks that 'input' can be parsed as a Device Management Cloud Certification Authority Id Cloud Certification Authority Leaf Certificate ID
func ValidateDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Cloud Certification Authority Id Cloud Certification Authority Leaf Certificate ID
func (id DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId) ID() string {
	fmtString := "/deviceManagement/cloudCertificationAuthority/%s/cloudCertificationAuthorityLeafCertificate/%s"
	return fmt.Sprintf(fmtString, id.CloudCertificationAuthorityId, id.CloudCertificationAuthorityLeafCertificateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Cloud Certification Authority Id Cloud Certification Authority Leaf Certificate ID
func (id DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("cloudCertificationAuthority", "cloudCertificationAuthority", "cloudCertificationAuthority"),
		resourceids.UserSpecifiedSegment("cloudCertificationAuthorityId", "cloudCertificationAuthorityId"),
		resourceids.StaticSegment("cloudCertificationAuthorityLeafCertificate", "cloudCertificationAuthorityLeafCertificate", "cloudCertificationAuthorityLeafCertificate"),
		resourceids.UserSpecifiedSegment("cloudCertificationAuthorityLeafCertificateId", "cloudCertificationAuthorityLeafCertificateId"),
	}
}

// String returns a human-readable description of this Device Management Cloud Certification Authority Id Cloud Certification Authority Leaf Certificate ID
func (id DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId) String() string {
	components := []string{
		fmt.Sprintf("Cloud Certification Authority: %q", id.CloudCertificationAuthorityId),
		fmt.Sprintf("Cloud Certification Authority Leaf Certificate: %q", id.CloudCertificationAuthorityLeafCertificateId),
	}
	return fmt.Sprintf("Device Management Cloud Certification Authority Id Cloud Certification Authority Leaf Certificate (%s)", strings.Join(components, "\n"))
}
