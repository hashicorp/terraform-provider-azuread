package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId{}

// RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId is a struct representing the Resource ID for a Role Management Enterprise App Id Transitive Role Assignment
type RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId struct {
	RbacApplicationId       string
	UnifiedRoleAssignmentId string
}

// NewRoleManagementEnterpriseAppIdTransitiveRoleAssignmentID returns a new RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId struct
func NewRoleManagementEnterpriseAppIdTransitiveRoleAssignmentID(rbacApplicationId string, unifiedRoleAssignmentId string) RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId {
	return RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId{
		RbacApplicationId:       rbacApplicationId,
		UnifiedRoleAssignmentId: unifiedRoleAssignmentId,
	}
}

// ParseRoleManagementEnterpriseAppIdTransitiveRoleAssignmentID parses 'input' into a RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId
func ParseRoleManagementEnterpriseAppIdTransitiveRoleAssignmentID(input string) (*RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppIdTransitiveRoleAssignmentIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppIdTransitiveRoleAssignmentIDInsensitively(input string) (*RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RbacApplicationId, ok = input.Parsed["rbacApplicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", input)
	}

	if id.UnifiedRoleAssignmentId, ok = input.Parsed["unifiedRoleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentId", input)
	}

	return nil
}

// ValidateRoleManagementEnterpriseAppIdTransitiveRoleAssignmentID checks that 'input' can be parsed as a Role Management Enterprise App Id Transitive Role Assignment ID
func ValidateRoleManagementEnterpriseAppIdTransitiveRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppIdTransitiveRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Id Transitive Role Assignment ID
func (id RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/transitiveRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Id Transitive Role Assignment ID
func (id RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("enterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationId"),
		resourceids.StaticSegment("transitiveRoleAssignments", "transitiveRoleAssignments", "transitiveRoleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentId", "unifiedRoleAssignmentId"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Id Transitive Role Assignment ID
func (id RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Assignment: %q", id.UnifiedRoleAssignmentId),
	}
	return fmt.Sprintf("Role Management Enterprise App Id Transitive Role Assignment (%s)", strings.Join(components, "\n"))
}
