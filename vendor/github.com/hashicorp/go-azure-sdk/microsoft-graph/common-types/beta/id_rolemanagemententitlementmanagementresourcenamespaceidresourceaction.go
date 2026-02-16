package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId{}

// RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId is a struct representing the Resource ID for a Role Management Entitlement Management Resource Namespace Id Resource Action
type RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId struct {
	UnifiedRbacResourceNamespaceId string
	UnifiedRbacResourceActionId    string
}

// NewRoleManagementEntitlementManagementResourceNamespaceIdResourceActionID returns a new RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId struct
func NewRoleManagementEntitlementManagementResourceNamespaceIdResourceActionID(unifiedRbacResourceNamespaceId string, unifiedRbacResourceActionId string) RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId {
	return RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
		UnifiedRbacResourceActionId:    unifiedRbacResourceActionId,
	}
}

// ParseRoleManagementEntitlementManagementResourceNamespaceIdResourceActionID parses 'input' into a RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId
func ParseRoleManagementEntitlementManagementResourceNamespaceIdResourceActionID(input string) (*RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEntitlementManagementResourceNamespaceIdResourceActionIDInsensitively parses 'input' case-insensitively into a RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEntitlementManagementResourceNamespaceIdResourceActionIDInsensitively(input string) (*RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRbacResourceNamespaceId, ok = input.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", input)
	}

	if id.UnifiedRbacResourceActionId, ok = input.Parsed["unifiedRbacResourceActionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", input)
	}

	return nil
}

// ValidateRoleManagementEntitlementManagementResourceNamespaceIdResourceActionID checks that 'input' can be parsed as a Role Management Entitlement Management Resource Namespace Id Resource Action ID
func ValidateRoleManagementEntitlementManagementResourceNamespaceIdResourceActionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEntitlementManagementResourceNamespaceIdResourceActionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Entitlement Management Resource Namespace Id Resource Action ID
func (id RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId) ID() string {
	fmtString := "/roleManagement/entitlementManagement/resourceNamespaces/%s/resourceActions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId, id.UnifiedRbacResourceActionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Entitlement Management Resource Namespace Id Resource Action ID
func (id RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceId"),
		resourceids.StaticSegment("resourceActions", "resourceActions", "resourceActions"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceActionId", "unifiedRbacResourceActionId"),
	}
}

// String returns a human-readable description of this Role Management Entitlement Management Resource Namespace Id Resource Action ID
func (id RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
		fmt.Sprintf("Unified Rbac Resource Action: %q", id.UnifiedRbacResourceActionId),
	}
	return fmt.Sprintf("Role Management Entitlement Management Resource Namespace Id Resource Action (%s)", strings.Join(components, "\n"))
}
