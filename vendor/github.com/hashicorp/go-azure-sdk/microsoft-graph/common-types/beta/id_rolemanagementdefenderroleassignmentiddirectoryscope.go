package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDefenderRoleAssignmentIdDirectoryScopeId{}

// RoleManagementDefenderRoleAssignmentIdDirectoryScopeId is a struct representing the Resource ID for a Role Management Defender Role Assignment Id Directory Scope
type RoleManagementDefenderRoleAssignmentIdDirectoryScopeId struct {
	UnifiedRoleAssignmentMultipleId string
	DirectoryObjectId               string
}

// NewRoleManagementDefenderRoleAssignmentIdDirectoryScopeID returns a new RoleManagementDefenderRoleAssignmentIdDirectoryScopeId struct
func NewRoleManagementDefenderRoleAssignmentIdDirectoryScopeID(unifiedRoleAssignmentMultipleId string, directoryObjectId string) RoleManagementDefenderRoleAssignmentIdDirectoryScopeId {
	return RoleManagementDefenderRoleAssignmentIdDirectoryScopeId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
		DirectoryObjectId:               directoryObjectId,
	}
}

// ParseRoleManagementDefenderRoleAssignmentIdDirectoryScopeID parses 'input' into a RoleManagementDefenderRoleAssignmentIdDirectoryScopeId
func ParseRoleManagementDefenderRoleAssignmentIdDirectoryScopeID(input string) (*RoleManagementDefenderRoleAssignmentIdDirectoryScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderRoleAssignmentIdDirectoryScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderRoleAssignmentIdDirectoryScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDefenderRoleAssignmentIdDirectoryScopeIDInsensitively parses 'input' case-insensitively into a RoleManagementDefenderRoleAssignmentIdDirectoryScopeId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDefenderRoleAssignmentIdDirectoryScopeIDInsensitively(input string) (*RoleManagementDefenderRoleAssignmentIdDirectoryScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderRoleAssignmentIdDirectoryScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderRoleAssignmentIdDirectoryScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDefenderRoleAssignmentIdDirectoryScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentMultipleId, ok = input.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateRoleManagementDefenderRoleAssignmentIdDirectoryScopeID checks that 'input' can be parsed as a Role Management Defender Role Assignment Id Directory Scope ID
func ValidateRoleManagementDefenderRoleAssignmentIdDirectoryScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDefenderRoleAssignmentIdDirectoryScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Defender Role Assignment Id Directory Scope ID
func (id RoleManagementDefenderRoleAssignmentIdDirectoryScopeId) ID() string {
	fmtString := "/roleManagement/defender/roleAssignments/%s/directoryScopes/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Defender Role Assignment Id Directory Scope ID
func (id RoleManagementDefenderRoleAssignmentIdDirectoryScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("defender", "defender", "defender"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleId"),
		resourceids.StaticSegment("directoryScopes", "directoryScopes", "directoryScopes"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Role Management Defender Role Assignment Id Directory Scope ID
func (id RoleManagementDefenderRoleAssignmentIdDirectoryScopeId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Role Management Defender Role Assignment Id Directory Scope (%s)", strings.Join(components, "\n"))
}
