package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDeviceManagementRoleDefinitionId{}

// RoleManagementDeviceManagementRoleDefinitionId is a struct representing the Resource ID for a Role Management Device Management Role Definition
type RoleManagementDeviceManagementRoleDefinitionId struct {
	UnifiedRoleDefinitionId string
}

// NewRoleManagementDeviceManagementRoleDefinitionID returns a new RoleManagementDeviceManagementRoleDefinitionId struct
func NewRoleManagementDeviceManagementRoleDefinitionID(unifiedRoleDefinitionId string) RoleManagementDeviceManagementRoleDefinitionId {
	return RoleManagementDeviceManagementRoleDefinitionId{
		UnifiedRoleDefinitionId: unifiedRoleDefinitionId,
	}
}

// ParseRoleManagementDeviceManagementRoleDefinitionID parses 'input' into a RoleManagementDeviceManagementRoleDefinitionId
func ParseRoleManagementDeviceManagementRoleDefinitionID(input string) (*RoleManagementDeviceManagementRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDeviceManagementRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDeviceManagementRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDeviceManagementRoleDefinitionIDInsensitively parses 'input' case-insensitively into a RoleManagementDeviceManagementRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDeviceManagementRoleDefinitionIDInsensitively(input string) (*RoleManagementDeviceManagementRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDeviceManagementRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDeviceManagementRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDeviceManagementRoleDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleDefinitionId, ok = input.Parsed["unifiedRoleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", input)
	}

	return nil
}

// ValidateRoleManagementDeviceManagementRoleDefinitionID checks that 'input' can be parsed as a Role Management Device Management Role Definition ID
func ValidateRoleManagementDeviceManagementRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDeviceManagementRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Device Management Role Definition ID
func (id RoleManagementDeviceManagementRoleDefinitionId) ID() string {
	fmtString := "/roleManagement/deviceManagement/roleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Device Management Role Definition ID
func (id RoleManagementDeviceManagementRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("roleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionId"),
	}
}

// String returns a human-readable description of this Role Management Device Management Role Definition ID
func (id RoleManagementDeviceManagementRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
	}
	return fmt.Sprintf("Role Management Device Management Role Definition (%s)", strings.Join(components, "\n"))
}
