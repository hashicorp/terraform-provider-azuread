package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId{}

// RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId is a struct representing the Resource ID for a Role Management Enterprise App Id Resource Namespace Id Resource Action
type RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId struct {
	RbacApplicationId              string
	UnifiedRbacResourceNamespaceId string
	UnifiedRbacResourceActionId    string
}

// NewRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionID returns a new RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId struct
func NewRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionID(rbacApplicationId string, unifiedRbacResourceNamespaceId string, unifiedRbacResourceActionId string) RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId {
	return RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId{
		RbacApplicationId:              rbacApplicationId,
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
		UnifiedRbacResourceActionId:    unifiedRbacResourceActionId,
	}
}

// ParseRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionID parses 'input' into a RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId
func ParseRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionID(input string) (*RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionIDInsensitively(input string) (*RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RbacApplicationId, ok = input.Parsed["rbacApplicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", input)
	}

	if id.UnifiedRbacResourceNamespaceId, ok = input.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", input)
	}

	if id.UnifiedRbacResourceActionId, ok = input.Parsed["unifiedRbacResourceActionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", input)
	}

	return nil
}

// ValidateRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionID checks that 'input' can be parsed as a Role Management Enterprise App Id Resource Namespace Id Resource Action ID
func ValidateRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Id Resource Namespace Id Resource Action ID
func (id RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/resourceNamespaces/%s/resourceActions/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRbacResourceNamespaceId, id.UnifiedRbacResourceActionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Id Resource Namespace Id Resource Action ID
func (id RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("enterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationId"),
		resourceids.StaticSegment("resourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceId"),
		resourceids.StaticSegment("resourceActions", "resourceActions", "resourceActions"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceActionId", "unifiedRbacResourceActionId"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Id Resource Namespace Id Resource Action ID
func (id RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
		fmt.Sprintf("Unified Rbac Resource Action: %q", id.UnifiedRbacResourceActionId),
	}
	return fmt.Sprintf("Role Management Enterprise App Id Resource Namespace Id Resource Action (%s)", strings.Join(components, "\n"))
}
