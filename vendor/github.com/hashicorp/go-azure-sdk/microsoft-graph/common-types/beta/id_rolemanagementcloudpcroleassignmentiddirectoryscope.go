package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId{}

// RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId is a struct representing the Resource ID for a Role Management Cloud P C Role Assignment Id Directory Scope
type RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId struct {
	UnifiedRoleAssignmentMultipleId string
	DirectoryObjectId               string
}

// NewRoleManagementCloudPCRoleAssignmentIdDirectoryScopeID returns a new RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId struct
func NewRoleManagementCloudPCRoleAssignmentIdDirectoryScopeID(unifiedRoleAssignmentMultipleId string, directoryObjectId string) RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId {
	return RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
		DirectoryObjectId:               directoryObjectId,
	}
}

// ParseRoleManagementCloudPCRoleAssignmentIdDirectoryScopeID parses 'input' into a RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId
func ParseRoleManagementCloudPCRoleAssignmentIdDirectoryScopeID(input string) (*RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementCloudPCRoleAssignmentIdDirectoryScopeIDInsensitively parses 'input' case-insensitively into a RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementCloudPCRoleAssignmentIdDirectoryScopeIDInsensitively(input string) (*RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentMultipleId, ok = input.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateRoleManagementCloudPCRoleAssignmentIdDirectoryScopeID checks that 'input' can be parsed as a Role Management Cloud P C Role Assignment Id Directory Scope ID
func ValidateRoleManagementCloudPCRoleAssignmentIdDirectoryScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementCloudPCRoleAssignmentIdDirectoryScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Cloud P C Role Assignment Id Directory Scope ID
func (id RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId) ID() string {
	fmtString := "/roleManagement/cloudPC/roleAssignments/%s/directoryScopes/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Cloud P C Role Assignment Id Directory Scope ID
func (id RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("cloudPC", "cloudPC", "cloudPC"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleId"),
		resourceids.StaticSegment("directoryScopes", "directoryScopes", "directoryScopes"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Role Management Cloud P C Role Assignment Id Directory Scope ID
func (id RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Role Management Cloud P C Role Assignment Id Directory Scope (%s)", strings.Join(components, "\n"))
}
