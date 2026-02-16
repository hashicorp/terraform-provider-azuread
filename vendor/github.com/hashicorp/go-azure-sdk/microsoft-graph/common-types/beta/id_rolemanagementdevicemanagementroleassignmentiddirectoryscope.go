package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId{}

// RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId is a struct representing the Resource ID for a Role Management Device Management Role Assignment Id Directory Scope
type RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId struct {
	UnifiedRoleAssignmentMultipleId string
	DirectoryObjectId               string
}

// NewRoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeID returns a new RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId struct
func NewRoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeID(unifiedRoleAssignmentMultipleId string, directoryObjectId string) RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId {
	return RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
		DirectoryObjectId:               directoryObjectId,
	}
}

// ParseRoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeID parses 'input' into a RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId
func ParseRoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeID(input string) (*RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeIDInsensitively parses 'input' case-insensitively into a RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeIDInsensitively(input string) (*RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentMultipleId, ok = input.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateRoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeID checks that 'input' can be parsed as a Role Management Device Management Role Assignment Id Directory Scope ID
func ValidateRoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Device Management Role Assignment Id Directory Scope ID
func (id RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId) ID() string {
	fmtString := "/roleManagement/deviceManagement/roleAssignments/%s/directoryScopes/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Device Management Role Assignment Id Directory Scope ID
func (id RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleId"),
		resourceids.StaticSegment("directoryScopes", "directoryScopes", "directoryScopes"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Role Management Device Management Role Assignment Id Directory Scope ID
func (id RoleManagementDeviceManagementRoleAssignmentIdDirectoryScopeId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Role Management Device Management Role Assignment Id Directory Scope (%s)", strings.Join(components, "\n"))
}
