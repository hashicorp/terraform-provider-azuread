package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCartToClassAssociationId{}

// DeviceManagementCartToClassAssociationId is a struct representing the Resource ID for a Device Management Cart To Class Association
type DeviceManagementCartToClassAssociationId struct {
	CartToClassAssociationId string
}

// NewDeviceManagementCartToClassAssociationID returns a new DeviceManagementCartToClassAssociationId struct
func NewDeviceManagementCartToClassAssociationID(cartToClassAssociationId string) DeviceManagementCartToClassAssociationId {
	return DeviceManagementCartToClassAssociationId{
		CartToClassAssociationId: cartToClassAssociationId,
	}
}

// ParseDeviceManagementCartToClassAssociationID parses 'input' into a DeviceManagementCartToClassAssociationId
func ParseDeviceManagementCartToClassAssociationID(input string) (*DeviceManagementCartToClassAssociationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCartToClassAssociationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCartToClassAssociationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementCartToClassAssociationIDInsensitively parses 'input' case-insensitively into a DeviceManagementCartToClassAssociationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementCartToClassAssociationIDInsensitively(input string) (*DeviceManagementCartToClassAssociationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCartToClassAssociationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCartToClassAssociationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementCartToClassAssociationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CartToClassAssociationId, ok = input.Parsed["cartToClassAssociationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cartToClassAssociationId", input)
	}

	return nil
}

// ValidateDeviceManagementCartToClassAssociationID checks that 'input' can be parsed as a Device Management Cart To Class Association ID
func ValidateDeviceManagementCartToClassAssociationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementCartToClassAssociationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Cart To Class Association ID
func (id DeviceManagementCartToClassAssociationId) ID() string {
	fmtString := "/deviceManagement/cartToClassAssociations/%s"
	return fmt.Sprintf(fmtString, id.CartToClassAssociationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Cart To Class Association ID
func (id DeviceManagementCartToClassAssociationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("cartToClassAssociations", "cartToClassAssociations", "cartToClassAssociations"),
		resourceids.UserSpecifiedSegment("cartToClassAssociationId", "cartToClassAssociationId"),
	}
}

// String returns a human-readable description of this Device Management Cart To Class Association ID
func (id DeviceManagementCartToClassAssociationId) String() string {
	components := []string{
		fmt.Sprintf("Cart To Class Association: %q", id.CartToClassAssociationId),
	}
	return fmt.Sprintf("Device Management Cart To Class Association (%s)", strings.Join(components, "\n"))
}
