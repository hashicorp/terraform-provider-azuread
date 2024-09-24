package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRoleScopeTagId{}

// DeviceManagementRoleScopeTagId is a struct representing the Resource ID for a Device Management Role Scope Tag
type DeviceManagementRoleScopeTagId struct {
	RoleScopeTagId string
}

// NewDeviceManagementRoleScopeTagID returns a new DeviceManagementRoleScopeTagId struct
func NewDeviceManagementRoleScopeTagID(roleScopeTagId string) DeviceManagementRoleScopeTagId {
	return DeviceManagementRoleScopeTagId{
		RoleScopeTagId: roleScopeTagId,
	}
}

// ParseDeviceManagementRoleScopeTagID parses 'input' into a DeviceManagementRoleScopeTagId
func ParseDeviceManagementRoleScopeTagID(input string) (*DeviceManagementRoleScopeTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRoleScopeTagId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRoleScopeTagId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementRoleScopeTagIDInsensitively parses 'input' case-insensitively into a DeviceManagementRoleScopeTagId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementRoleScopeTagIDInsensitively(input string) (*DeviceManagementRoleScopeTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRoleScopeTagId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRoleScopeTagId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementRoleScopeTagId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RoleScopeTagId, ok = input.Parsed["roleScopeTagId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "roleScopeTagId", input)
	}

	return nil
}

// ValidateDeviceManagementRoleScopeTagID checks that 'input' can be parsed as a Device Management Role Scope Tag ID
func ValidateDeviceManagementRoleScopeTagID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementRoleScopeTagID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Role Scope Tag ID
func (id DeviceManagementRoleScopeTagId) ID() string {
	fmtString := "/deviceManagement/roleScopeTags/%s"
	return fmt.Sprintf(fmtString, id.RoleScopeTagId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Role Scope Tag ID
func (id DeviceManagementRoleScopeTagId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("roleScopeTags", "roleScopeTags", "roleScopeTags"),
		resourceids.UserSpecifiedSegment("roleScopeTagId", "roleScopeTagId"),
	}
}

// String returns a human-readable description of this Device Management Role Scope Tag ID
func (id DeviceManagementRoleScopeTagId) String() string {
	components := []string{
		fmt.Sprintf("Role Scope Tag: %q", id.RoleScopeTagId),
	}
	return fmt.Sprintf("Device Management Role Scope Tag (%s)", strings.Join(components, "\n"))
}
