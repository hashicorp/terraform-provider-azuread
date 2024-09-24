package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId{}

// RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId is a struct representing the Resource ID for a Role Management Device Management Role Definition Id Inherits Permissions From
type RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId struct {
	UnifiedRoleDefinitionId  string
	UnifiedRoleDefinitionId1 string
}

// NewRoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromID returns a new RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId struct
func NewRoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromID(unifiedRoleDefinitionId string, unifiedRoleDefinitionId1 string) RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId {
	return RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId{
		UnifiedRoleDefinitionId:  unifiedRoleDefinitionId,
		UnifiedRoleDefinitionId1: unifiedRoleDefinitionId1,
	}
}

// ParseRoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromID parses 'input' into a RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId
func ParseRoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromID(input string) (*RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromIDInsensitively parses 'input' case-insensitively into a RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromIDInsensitively(input string) (*RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleDefinitionId, ok = input.Parsed["unifiedRoleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", input)
	}

	if id.UnifiedRoleDefinitionId1, ok = input.Parsed["unifiedRoleDefinitionId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId1", input)
	}

	return nil
}

// ValidateRoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromID checks that 'input' can be parsed as a Role Management Device Management Role Definition Id Inherits Permissions From ID
func ValidateRoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Device Management Role Definition Id Inherits Permissions From ID
func (id RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId) ID() string {
	fmtString := "/roleManagement/deviceManagement/roleDefinitions/%s/inheritsPermissionsFrom/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleDefinitionId, id.UnifiedRoleDefinitionId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Device Management Role Definition Id Inherits Permissions From ID
func (id RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("roleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionId"),
		resourceids.StaticSegment("inheritsPermissionsFrom", "inheritsPermissionsFrom", "inheritsPermissionsFrom"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId1", "unifiedRoleDefinitionId1"),
	}
}

// String returns a human-readable description of this Role Management Device Management Role Definition Id Inherits Permissions From ID
func (id RoleManagementDeviceManagementRoleDefinitionIdInheritsPermissionsFromId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
		fmt.Sprintf("Unified Role Definition Id 1: %q", id.UnifiedRoleDefinitionId1),
	}
	return fmt.Sprintf("Role Management Device Management Role Definition Id Inherits Permissions From (%s)", strings.Join(components, "\n"))
}
