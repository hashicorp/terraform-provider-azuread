package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleAssignmentId{}

// RoleManagementEnterpriseAppIdRoleAssignmentId is a struct representing the Resource ID for a Role Management Enterprise App Id Role Assignment
type RoleManagementEnterpriseAppIdRoleAssignmentId struct {
	RbacApplicationId       string
	UnifiedRoleAssignmentId string
}

// NewRoleManagementEnterpriseAppIdRoleAssignmentID returns a new RoleManagementEnterpriseAppIdRoleAssignmentId struct
func NewRoleManagementEnterpriseAppIdRoleAssignmentID(rbacApplicationId string, unifiedRoleAssignmentId string) RoleManagementEnterpriseAppIdRoleAssignmentId {
	return RoleManagementEnterpriseAppIdRoleAssignmentId{
		RbacApplicationId:       rbacApplicationId,
		UnifiedRoleAssignmentId: unifiedRoleAssignmentId,
	}
}

// ParseRoleManagementEnterpriseAppIdRoleAssignmentID parses 'input' into a RoleManagementEnterpriseAppIdRoleAssignmentId
func ParseRoleManagementEnterpriseAppIdRoleAssignmentID(input string) (*RoleManagementEnterpriseAppIdRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppIdRoleAssignmentIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppIdRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppIdRoleAssignmentIDInsensitively(input string) (*RoleManagementEnterpriseAppIdRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEnterpriseAppIdRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RbacApplicationId, ok = input.Parsed["rbacApplicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", input)
	}

	if id.UnifiedRoleAssignmentId, ok = input.Parsed["unifiedRoleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentId", input)
	}

	return nil
}

// ValidateRoleManagementEnterpriseAppIdRoleAssignmentID checks that 'input' can be parsed as a Role Management Enterprise App Id Role Assignment ID
func ValidateRoleManagementEnterpriseAppIdRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppIdRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Id Role Assignment ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleAssignments/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Id Role Assignment ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("enterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationId"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentId", "unifiedRoleAssignmentId"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Id Role Assignment ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Assignment: %q", id.UnifiedRoleAssignmentId),
	}
	return fmt.Sprintf("Role Management Enterprise App Id Role Assignment (%s)", strings.Join(components, "\n"))
}
