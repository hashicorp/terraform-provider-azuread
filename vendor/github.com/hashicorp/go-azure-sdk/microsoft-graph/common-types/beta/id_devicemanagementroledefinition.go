package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRoleDefinitionId{}

// DeviceManagementRoleDefinitionId is a struct representing the Resource ID for a Device Management Role Definition
type DeviceManagementRoleDefinitionId struct {
	RoleDefinitionId string
}

// NewDeviceManagementRoleDefinitionID returns a new DeviceManagementRoleDefinitionId struct
func NewDeviceManagementRoleDefinitionID(roleDefinitionId string) DeviceManagementRoleDefinitionId {
	return DeviceManagementRoleDefinitionId{
		RoleDefinitionId: roleDefinitionId,
	}
}

// ParseDeviceManagementRoleDefinitionID parses 'input' into a DeviceManagementRoleDefinitionId
func ParseDeviceManagementRoleDefinitionID(input string) (*DeviceManagementRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementRoleDefinitionIDInsensitively parses 'input' case-insensitively into a DeviceManagementRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementRoleDefinitionIDInsensitively(input string) (*DeviceManagementRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementRoleDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RoleDefinitionId, ok = input.Parsed["roleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "roleDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementRoleDefinitionID checks that 'input' can be parsed as a Device Management Role Definition ID
func ValidateDeviceManagementRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Role Definition ID
func (id DeviceManagementRoleDefinitionId) ID() string {
	fmtString := "/deviceManagement/roleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.RoleDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Role Definition ID
func (id DeviceManagementRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("roleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("roleDefinitionId", "roleDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Role Definition ID
func (id DeviceManagementRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Role Definition: %q", id.RoleDefinitionId),
	}
	return fmt.Sprintf("Device Management Role Definition (%s)", strings.Join(components, "\n"))
}
