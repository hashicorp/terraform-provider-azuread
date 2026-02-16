package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleDefinitionId{}

// RoleManagementEnterpriseAppIdRoleDefinitionId is a struct representing the Resource ID for a Role Management Enterprise App Id Role Definition
type RoleManagementEnterpriseAppIdRoleDefinitionId struct {
	RbacApplicationId       string
	UnifiedRoleDefinitionId string
}

// NewRoleManagementEnterpriseAppIdRoleDefinitionID returns a new RoleManagementEnterpriseAppIdRoleDefinitionId struct
func NewRoleManagementEnterpriseAppIdRoleDefinitionID(rbacApplicationId string, unifiedRoleDefinitionId string) RoleManagementEnterpriseAppIdRoleDefinitionId {
	return RoleManagementEnterpriseAppIdRoleDefinitionId{
		RbacApplicationId:       rbacApplicationId,
		UnifiedRoleDefinitionId: unifiedRoleDefinitionId,
	}
}

// ParseRoleManagementEnterpriseAppIdRoleDefinitionID parses 'input' into a RoleManagementEnterpriseAppIdRoleDefinitionId
func ParseRoleManagementEnterpriseAppIdRoleDefinitionID(input string) (*RoleManagementEnterpriseAppIdRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppIdRoleDefinitionIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppIdRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppIdRoleDefinitionIDInsensitively(input string) (*RoleManagementEnterpriseAppIdRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEnterpriseAppIdRoleDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RbacApplicationId, ok = input.Parsed["rbacApplicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", input)
	}

	if id.UnifiedRoleDefinitionId, ok = input.Parsed["unifiedRoleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", input)
	}

	return nil
}

// ValidateRoleManagementEnterpriseAppIdRoleDefinitionID checks that 'input' can be parsed as a Role Management Enterprise App Id Role Definition ID
func ValidateRoleManagementEnterpriseAppIdRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppIdRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Id Role Definition ID
func (id RoleManagementEnterpriseAppIdRoleDefinitionId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Id Role Definition ID
func (id RoleManagementEnterpriseAppIdRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("enterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationId"),
		resourceids.StaticSegment("roleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionId"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Id Role Definition ID
func (id RoleManagementEnterpriseAppIdRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
	}
	return fmt.Sprintf("Role Management Enterprise App Id Role Definition (%s)", strings.Join(components, "\n"))
}
