package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCertificateConnectorDetailId{}

// DeviceManagementCertificateConnectorDetailId is a struct representing the Resource ID for a Device Management Certificate Connector Detail
type DeviceManagementCertificateConnectorDetailId struct {
	CertificateConnectorDetailsId string
}

// NewDeviceManagementCertificateConnectorDetailID returns a new DeviceManagementCertificateConnectorDetailId struct
func NewDeviceManagementCertificateConnectorDetailID(certificateConnectorDetailsId string) DeviceManagementCertificateConnectorDetailId {
	return DeviceManagementCertificateConnectorDetailId{
		CertificateConnectorDetailsId: certificateConnectorDetailsId,
	}
}

// ParseDeviceManagementCertificateConnectorDetailID parses 'input' into a DeviceManagementCertificateConnectorDetailId
func ParseDeviceManagementCertificateConnectorDetailID(input string) (*DeviceManagementCertificateConnectorDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCertificateConnectorDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCertificateConnectorDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementCertificateConnectorDetailIDInsensitively parses 'input' case-insensitively into a DeviceManagementCertificateConnectorDetailId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementCertificateConnectorDetailIDInsensitively(input string) (*DeviceManagementCertificateConnectorDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCertificateConnectorDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCertificateConnectorDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementCertificateConnectorDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CertificateConnectorDetailsId, ok = input.Parsed["certificateConnectorDetailsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "certificateConnectorDetailsId", input)
	}

	return nil
}

// ValidateDeviceManagementCertificateConnectorDetailID checks that 'input' can be parsed as a Device Management Certificate Connector Detail ID
func ValidateDeviceManagementCertificateConnectorDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementCertificateConnectorDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Certificate Connector Detail ID
func (id DeviceManagementCertificateConnectorDetailId) ID() string {
	fmtString := "/deviceManagement/certificateConnectorDetails/%s"
	return fmt.Sprintf(fmtString, id.CertificateConnectorDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Certificate Connector Detail ID
func (id DeviceManagementCertificateConnectorDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("certificateConnectorDetails", "certificateConnectorDetails", "certificateConnectorDetails"),
		resourceids.UserSpecifiedSegment("certificateConnectorDetailsId", "certificateConnectorDetailsId"),
	}
}

// String returns a human-readable description of this Device Management Certificate Connector Detail ID
func (id DeviceManagementCertificateConnectorDetailId) String() string {
	components := []string{
		fmt.Sprintf("Certificate Connector Details: %q", id.CertificateConnectorDetailsId),
	}
	return fmt.Sprintf("Device Management Certificate Connector Detail (%s)", strings.Join(components, "\n"))
}
