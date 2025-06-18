package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId{}

// RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId is a struct representing the Resource ID for a Role Management Defender Role Definition Id Inherits Permissions From
type RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId struct {
	UnifiedRoleDefinitionId  string
	UnifiedRoleDefinitionId1 string
}

// NewRoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromID returns a new RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId struct
func NewRoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromID(unifiedRoleDefinitionId string, unifiedRoleDefinitionId1 string) RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId {
	return RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId{
		UnifiedRoleDefinitionId:  unifiedRoleDefinitionId,
		UnifiedRoleDefinitionId1: unifiedRoleDefinitionId1,
	}
}

// ParseRoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromID parses 'input' into a RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId
func ParseRoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromID(input string) (*RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromIDInsensitively parses 'input' case-insensitively into a RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromIDInsensitively(input string) (*RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleDefinitionId, ok = input.Parsed["unifiedRoleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", input)
	}

	if id.UnifiedRoleDefinitionId1, ok = input.Parsed["unifiedRoleDefinitionId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId1", input)
	}

	return nil
}

// ValidateRoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromID checks that 'input' can be parsed as a Role Management Defender Role Definition Id Inherits Permissions From ID
func ValidateRoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Defender Role Definition Id Inherits Permissions From ID
func (id RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId) ID() string {
	fmtString := "/roleManagement/defender/roleDefinitions/%s/inheritsPermissionsFrom/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleDefinitionId, id.UnifiedRoleDefinitionId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Defender Role Definition Id Inherits Permissions From ID
func (id RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("defender", "defender", "defender"),
		resourceids.StaticSegment("roleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionId"),
		resourceids.StaticSegment("inheritsPermissionsFrom", "inheritsPermissionsFrom", "inheritsPermissionsFrom"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId1", "unifiedRoleDefinitionId1"),
	}
}

// String returns a human-readable description of this Role Management Defender Role Definition Id Inherits Permissions From ID
func (id RoleManagementDefenderRoleDefinitionIdInheritsPermissionsFromId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
		fmt.Sprintf("Unified Role Definition Id 1: %q", id.UnifiedRoleDefinitionId1),
	}
	return fmt.Sprintf("Role Management Defender Role Definition Id Inherits Permissions From (%s)", strings.Join(components, "\n"))
}
