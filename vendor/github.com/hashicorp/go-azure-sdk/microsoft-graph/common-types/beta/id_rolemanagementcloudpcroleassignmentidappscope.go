package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementCloudPCRoleAssignmentIdAppScopeId{}

// RoleManagementCloudPCRoleAssignmentIdAppScopeId is a struct representing the Resource ID for a Role Management Cloud P C Role Assignment Id App Scope
type RoleManagementCloudPCRoleAssignmentIdAppScopeId struct {
	UnifiedRoleAssignmentMultipleId string
	AppScopeId                      string
}

// NewRoleManagementCloudPCRoleAssignmentIdAppScopeID returns a new RoleManagementCloudPCRoleAssignmentIdAppScopeId struct
func NewRoleManagementCloudPCRoleAssignmentIdAppScopeID(unifiedRoleAssignmentMultipleId string, appScopeId string) RoleManagementCloudPCRoleAssignmentIdAppScopeId {
	return RoleManagementCloudPCRoleAssignmentIdAppScopeId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
		AppScopeId:                      appScopeId,
	}
}

// ParseRoleManagementCloudPCRoleAssignmentIdAppScopeID parses 'input' into a RoleManagementCloudPCRoleAssignmentIdAppScopeId
func ParseRoleManagementCloudPCRoleAssignmentIdAppScopeID(input string) (*RoleManagementCloudPCRoleAssignmentIdAppScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementCloudPCRoleAssignmentIdAppScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementCloudPCRoleAssignmentIdAppScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementCloudPCRoleAssignmentIdAppScopeIDInsensitively parses 'input' case-insensitively into a RoleManagementCloudPCRoleAssignmentIdAppScopeId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementCloudPCRoleAssignmentIdAppScopeIDInsensitively(input string) (*RoleManagementCloudPCRoleAssignmentIdAppScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementCloudPCRoleAssignmentIdAppScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementCloudPCRoleAssignmentIdAppScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementCloudPCRoleAssignmentIdAppScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentMultipleId, ok = input.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", input)
	}

	if id.AppScopeId, ok = input.Parsed["appScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appScopeId", input)
	}

	return nil
}

// ValidateRoleManagementCloudPCRoleAssignmentIdAppScopeID checks that 'input' can be parsed as a Role Management Cloud P C Role Assignment Id App Scope ID
func ValidateRoleManagementCloudPCRoleAssignmentIdAppScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementCloudPCRoleAssignmentIdAppScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Cloud P C Role Assignment Id App Scope ID
func (id RoleManagementCloudPCRoleAssignmentIdAppScopeId) ID() string {
	fmtString := "/roleManagement/cloudPC/roleAssignments/%s/appScopes/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId, id.AppScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Cloud P C Role Assignment Id App Scope ID
func (id RoleManagementCloudPCRoleAssignmentIdAppScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("cloudPC", "cloudPC", "cloudPC"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleId"),
		resourceids.StaticSegment("appScopes", "appScopes", "appScopes"),
		resourceids.UserSpecifiedSegment("appScopeId", "appScopeId"),
	}
}

// String returns a human-readable description of this Role Management Cloud P C Role Assignment Id App Scope ID
func (id RoleManagementCloudPCRoleAssignmentIdAppScopeId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
		fmt.Sprintf("App Scope: %q", id.AppScopeId),
	}
	return fmt.Sprintf("Role Management Cloud P C Role Assignment Id App Scope (%s)", strings.Join(components, "\n"))
}
