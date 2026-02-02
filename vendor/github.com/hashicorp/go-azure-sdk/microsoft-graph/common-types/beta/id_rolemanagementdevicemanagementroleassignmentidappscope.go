package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDeviceManagementRoleAssignmentIdAppScopeId{}

// RoleManagementDeviceManagementRoleAssignmentIdAppScopeId is a struct representing the Resource ID for a Role Management Device Management Role Assignment Id App Scope
type RoleManagementDeviceManagementRoleAssignmentIdAppScopeId struct {
	UnifiedRoleAssignmentMultipleId string
	AppScopeId                      string
}

// NewRoleManagementDeviceManagementRoleAssignmentIdAppScopeID returns a new RoleManagementDeviceManagementRoleAssignmentIdAppScopeId struct
func NewRoleManagementDeviceManagementRoleAssignmentIdAppScopeID(unifiedRoleAssignmentMultipleId string, appScopeId string) RoleManagementDeviceManagementRoleAssignmentIdAppScopeId {
	return RoleManagementDeviceManagementRoleAssignmentIdAppScopeId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
		AppScopeId:                      appScopeId,
	}
}

// ParseRoleManagementDeviceManagementRoleAssignmentIdAppScopeID parses 'input' into a RoleManagementDeviceManagementRoleAssignmentIdAppScopeId
func ParseRoleManagementDeviceManagementRoleAssignmentIdAppScopeID(input string) (*RoleManagementDeviceManagementRoleAssignmentIdAppScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDeviceManagementRoleAssignmentIdAppScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDeviceManagementRoleAssignmentIdAppScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDeviceManagementRoleAssignmentIdAppScopeIDInsensitively parses 'input' case-insensitively into a RoleManagementDeviceManagementRoleAssignmentIdAppScopeId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDeviceManagementRoleAssignmentIdAppScopeIDInsensitively(input string) (*RoleManagementDeviceManagementRoleAssignmentIdAppScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDeviceManagementRoleAssignmentIdAppScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDeviceManagementRoleAssignmentIdAppScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDeviceManagementRoleAssignmentIdAppScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentMultipleId, ok = input.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", input)
	}

	if id.AppScopeId, ok = input.Parsed["appScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appScopeId", input)
	}

	return nil
}

// ValidateRoleManagementDeviceManagementRoleAssignmentIdAppScopeID checks that 'input' can be parsed as a Role Management Device Management Role Assignment Id App Scope ID
func ValidateRoleManagementDeviceManagementRoleAssignmentIdAppScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDeviceManagementRoleAssignmentIdAppScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Device Management Role Assignment Id App Scope ID
func (id RoleManagementDeviceManagementRoleAssignmentIdAppScopeId) ID() string {
	fmtString := "/roleManagement/deviceManagement/roleAssignments/%s/appScopes/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId, id.AppScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Device Management Role Assignment Id App Scope ID
func (id RoleManagementDeviceManagementRoleAssignmentIdAppScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleId"),
		resourceids.StaticSegment("appScopes", "appScopes", "appScopes"),
		resourceids.UserSpecifiedSegment("appScopeId", "appScopeId"),
	}
}

// String returns a human-readable description of this Role Management Device Management Role Assignment Id App Scope ID
func (id RoleManagementDeviceManagementRoleAssignmentIdAppScopeId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
		fmt.Sprintf("App Scope: %q", id.AppScopeId),
	}
	return fmt.Sprintf("Role Management Device Management Role Assignment Id App Scope (%s)", strings.Join(components, "\n"))
}
