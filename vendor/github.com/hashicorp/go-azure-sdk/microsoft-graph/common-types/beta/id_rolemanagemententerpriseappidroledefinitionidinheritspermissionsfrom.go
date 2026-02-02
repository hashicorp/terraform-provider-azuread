package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId{}

// RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId is a struct representing the Resource ID for a Role Management Enterprise App Id Role Definition Id Inherits Permissions From
type RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId struct {
	RbacApplicationId        string
	UnifiedRoleDefinitionId  string
	UnifiedRoleDefinitionId1 string
}

// NewRoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromID returns a new RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId struct
func NewRoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromID(rbacApplicationId string, unifiedRoleDefinitionId string, unifiedRoleDefinitionId1 string) RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId {
	return RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId{
		RbacApplicationId:        rbacApplicationId,
		UnifiedRoleDefinitionId:  unifiedRoleDefinitionId,
		UnifiedRoleDefinitionId1: unifiedRoleDefinitionId1,
	}
}

// ParseRoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromID parses 'input' into a RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId
func ParseRoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromID(input string) (*RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromIDInsensitively(input string) (*RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RbacApplicationId, ok = input.Parsed["rbacApplicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", input)
	}

	if id.UnifiedRoleDefinitionId, ok = input.Parsed["unifiedRoleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", input)
	}

	if id.UnifiedRoleDefinitionId1, ok = input.Parsed["unifiedRoleDefinitionId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId1", input)
	}

	return nil
}

// ValidateRoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromID checks that 'input' can be parsed as a Role Management Enterprise App Id Role Definition Id Inherits Permissions From ID
func ValidateRoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Id Role Definition Id Inherits Permissions From ID
func (id RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleDefinitions/%s/inheritsPermissionsFrom/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleDefinitionId, id.UnifiedRoleDefinitionId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Id Role Definition Id Inherits Permissions From ID
func (id RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("enterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationId"),
		resourceids.StaticSegment("roleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionId"),
		resourceids.StaticSegment("inheritsPermissionsFrom", "inheritsPermissionsFrom", "inheritsPermissionsFrom"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId1", "unifiedRoleDefinitionId1"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Id Role Definition Id Inherits Permissions From ID
func (id RoleManagementEnterpriseAppIdRoleDefinitionIdInheritsPermissionsFromId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
		fmt.Sprintf("Unified Role Definition Id 1: %q", id.UnifiedRoleDefinitionId1),
	}
	return fmt.Sprintf("Role Management Enterprise App Id Role Definition Id Inherits Permissions From (%s)", strings.Join(components, "\n"))
}
