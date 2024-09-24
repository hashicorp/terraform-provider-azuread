package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdResourceNamespaceId{}

// RoleManagementEnterpriseAppIdResourceNamespaceId is a struct representing the Resource ID for a Role Management Enterprise App Id Resource Namespace
type RoleManagementEnterpriseAppIdResourceNamespaceId struct {
	RbacApplicationId              string
	UnifiedRbacResourceNamespaceId string
}

// NewRoleManagementEnterpriseAppIdResourceNamespaceID returns a new RoleManagementEnterpriseAppIdResourceNamespaceId struct
func NewRoleManagementEnterpriseAppIdResourceNamespaceID(rbacApplicationId string, unifiedRbacResourceNamespaceId string) RoleManagementEnterpriseAppIdResourceNamespaceId {
	return RoleManagementEnterpriseAppIdResourceNamespaceId{
		RbacApplicationId:              rbacApplicationId,
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
	}
}

// ParseRoleManagementEnterpriseAppIdResourceNamespaceID parses 'input' into a RoleManagementEnterpriseAppIdResourceNamespaceId
func ParseRoleManagementEnterpriseAppIdResourceNamespaceID(input string) (*RoleManagementEnterpriseAppIdResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdResourceNamespaceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdResourceNamespaceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppIdResourceNamespaceIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppIdResourceNamespaceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppIdResourceNamespaceIDInsensitively(input string) (*RoleManagementEnterpriseAppIdResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdResourceNamespaceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdResourceNamespaceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEnterpriseAppIdResourceNamespaceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RbacApplicationId, ok = input.Parsed["rbacApplicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", input)
	}

	if id.UnifiedRbacResourceNamespaceId, ok = input.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", input)
	}

	return nil
}

// ValidateRoleManagementEnterpriseAppIdResourceNamespaceID checks that 'input' can be parsed as a Role Management Enterprise App Id Resource Namespace ID
func ValidateRoleManagementEnterpriseAppIdResourceNamespaceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppIdResourceNamespaceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Id Resource Namespace ID
func (id RoleManagementEnterpriseAppIdResourceNamespaceId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/resourceNamespaces/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRbacResourceNamespaceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Id Resource Namespace ID
func (id RoleManagementEnterpriseAppIdResourceNamespaceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("enterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationId"),
		resourceids.StaticSegment("resourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceId"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Id Resource Namespace ID
func (id RoleManagementEnterpriseAppIdResourceNamespaceId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
	}
	return fmt.Sprintf("Role Management Enterprise App Id Resource Namespace (%s)", strings.Join(components, "\n"))
}
