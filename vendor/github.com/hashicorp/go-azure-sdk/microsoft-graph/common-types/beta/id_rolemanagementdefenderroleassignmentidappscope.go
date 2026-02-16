package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDefenderRoleAssignmentIdAppScopeId{}

// RoleManagementDefenderRoleAssignmentIdAppScopeId is a struct representing the Resource ID for a Role Management Defender Role Assignment Id App Scope
type RoleManagementDefenderRoleAssignmentIdAppScopeId struct {
	UnifiedRoleAssignmentMultipleId string
	AppScopeId                      string
}

// NewRoleManagementDefenderRoleAssignmentIdAppScopeID returns a new RoleManagementDefenderRoleAssignmentIdAppScopeId struct
func NewRoleManagementDefenderRoleAssignmentIdAppScopeID(unifiedRoleAssignmentMultipleId string, appScopeId string) RoleManagementDefenderRoleAssignmentIdAppScopeId {
	return RoleManagementDefenderRoleAssignmentIdAppScopeId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
		AppScopeId:                      appScopeId,
	}
}

// ParseRoleManagementDefenderRoleAssignmentIdAppScopeID parses 'input' into a RoleManagementDefenderRoleAssignmentIdAppScopeId
func ParseRoleManagementDefenderRoleAssignmentIdAppScopeID(input string) (*RoleManagementDefenderRoleAssignmentIdAppScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderRoleAssignmentIdAppScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderRoleAssignmentIdAppScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDefenderRoleAssignmentIdAppScopeIDInsensitively parses 'input' case-insensitively into a RoleManagementDefenderRoleAssignmentIdAppScopeId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDefenderRoleAssignmentIdAppScopeIDInsensitively(input string) (*RoleManagementDefenderRoleAssignmentIdAppScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderRoleAssignmentIdAppScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderRoleAssignmentIdAppScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDefenderRoleAssignmentIdAppScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentMultipleId, ok = input.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", input)
	}

	if id.AppScopeId, ok = input.Parsed["appScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appScopeId", input)
	}

	return nil
}

// ValidateRoleManagementDefenderRoleAssignmentIdAppScopeID checks that 'input' can be parsed as a Role Management Defender Role Assignment Id App Scope ID
func ValidateRoleManagementDefenderRoleAssignmentIdAppScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDefenderRoleAssignmentIdAppScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Defender Role Assignment Id App Scope ID
func (id RoleManagementDefenderRoleAssignmentIdAppScopeId) ID() string {
	fmtString := "/roleManagement/defender/roleAssignments/%s/appScopes/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId, id.AppScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Defender Role Assignment Id App Scope ID
func (id RoleManagementDefenderRoleAssignmentIdAppScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("defender", "defender", "defender"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleId"),
		resourceids.StaticSegment("appScopes", "appScopes", "appScopes"),
		resourceids.UserSpecifiedSegment("appScopeId", "appScopeId"),
	}
}

// String returns a human-readable description of this Role Management Defender Role Assignment Id App Scope ID
func (id RoleManagementDefenderRoleAssignmentIdAppScopeId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
		fmt.Sprintf("App Scope: %q", id.AppScopeId),
	}
	return fmt.Sprintf("Role Management Defender Role Assignment Id App Scope (%s)", strings.Join(components, "\n"))
}
