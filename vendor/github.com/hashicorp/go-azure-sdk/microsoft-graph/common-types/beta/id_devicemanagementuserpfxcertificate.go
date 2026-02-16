package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserPfxCertificateId{}

// DeviceManagementUserPfxCertificateId is a struct representing the Resource ID for a Device Management User Pfx Certificate
type DeviceManagementUserPfxCertificateId struct {
	UserPFXCertificateId string
}

// NewDeviceManagementUserPfxCertificateID returns a new DeviceManagementUserPfxCertificateId struct
func NewDeviceManagementUserPfxCertificateID(userPFXCertificateId string) DeviceManagementUserPfxCertificateId {
	return DeviceManagementUserPfxCertificateId{
		UserPFXCertificateId: userPFXCertificateId,
	}
}

// ParseDeviceManagementUserPfxCertificateID parses 'input' into a DeviceManagementUserPfxCertificateId
func ParseDeviceManagementUserPfxCertificateID(input string) (*DeviceManagementUserPfxCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserPfxCertificateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserPfxCertificateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserPfxCertificateIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserPfxCertificateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserPfxCertificateIDInsensitively(input string) (*DeviceManagementUserPfxCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserPfxCertificateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserPfxCertificateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserPfxCertificateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserPFXCertificateId, ok = input.Parsed["userPFXCertificateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userPFXCertificateId", input)
	}

	return nil
}

// ValidateDeviceManagementUserPfxCertificateID checks that 'input' can be parsed as a Device Management User Pfx Certificate ID
func ValidateDeviceManagementUserPfxCertificateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserPfxCertificateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Pfx Certificate ID
func (id DeviceManagementUserPfxCertificateId) ID() string {
	fmtString := "/deviceManagement/userPfxCertificates/%s"
	return fmt.Sprintf(fmtString, id.UserPFXCertificateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Pfx Certificate ID
func (id DeviceManagementUserPfxCertificateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userPfxCertificates", "userPfxCertificates", "userPfxCertificates"),
		resourceids.UserSpecifiedSegment("userPFXCertificateId", "userPFXCertificateId"),
	}
}

// String returns a human-readable description of this Device Management User Pfx Certificate ID
func (id DeviceManagementUserPfxCertificateId) String() string {
	components := []string{
		fmt.Sprintf("User PFX Certificate: %q", id.UserPFXCertificateId),
	}
	return fmt.Sprintf("Device Management User Pfx Certificate (%s)", strings.Join(components, "\n"))
}
