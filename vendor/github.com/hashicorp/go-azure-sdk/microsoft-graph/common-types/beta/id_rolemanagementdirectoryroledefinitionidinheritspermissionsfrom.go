package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId{}

// RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId is a struct representing the Resource ID for a Role Management Directory Role Definition Id Inherits Permissions From
type RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId struct {
	UnifiedRoleDefinitionId  string
	UnifiedRoleDefinitionId1 string
}

// NewRoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromID returns a new RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId struct
func NewRoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromID(unifiedRoleDefinitionId string, unifiedRoleDefinitionId1 string) RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId {
	return RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId{
		UnifiedRoleDefinitionId:  unifiedRoleDefinitionId,
		UnifiedRoleDefinitionId1: unifiedRoleDefinitionId1,
	}
}

// ParseRoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromID parses 'input' into a RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId
func ParseRoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromID(input string) (*RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromIDInsensitively(input string) (*RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleDefinitionId, ok = input.Parsed["unifiedRoleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", input)
	}

	if id.UnifiedRoleDefinitionId1, ok = input.Parsed["unifiedRoleDefinitionId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId1", input)
	}

	return nil
}

// ValidateRoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromID checks that 'input' can be parsed as a Role Management Directory Role Definition Id Inherits Permissions From ID
func ValidateRoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Role Definition Id Inherits Permissions From ID
func (id RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId) ID() string {
	fmtString := "/roleManagement/directory/roleDefinitions/%s/inheritsPermissionsFrom/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleDefinitionId, id.UnifiedRoleDefinitionId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Role Definition Id Inherits Permissions From ID
func (id RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("roleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionId"),
		resourceids.StaticSegment("inheritsPermissionsFrom", "inheritsPermissionsFrom", "inheritsPermissionsFrom"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId1", "unifiedRoleDefinitionId1"),
	}
}

// String returns a human-readable description of this Role Management Directory Role Definition Id Inherits Permissions From ID
func (id RoleManagementDirectoryRoleDefinitionIdInheritsPermissionsFromId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
		fmt.Sprintf("Unified Role Definition Id 1: %q", id.UnifiedRoleDefinitionId1),
	}
	return fmt.Sprintf("Role Management Directory Role Definition Id Inherits Permissions From (%s)", strings.Join(components, "\n"))
}
