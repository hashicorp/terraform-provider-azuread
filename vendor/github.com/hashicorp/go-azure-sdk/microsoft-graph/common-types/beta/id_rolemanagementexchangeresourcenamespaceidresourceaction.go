package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementExchangeResourceNamespaceIdResourceActionId{}

// RoleManagementExchangeResourceNamespaceIdResourceActionId is a struct representing the Resource ID for a Role Management Exchange Resource Namespace Id Resource Action
type RoleManagementExchangeResourceNamespaceIdResourceActionId struct {
	UnifiedRbacResourceNamespaceId string
	UnifiedRbacResourceActionId    string
}

// NewRoleManagementExchangeResourceNamespaceIdResourceActionID returns a new RoleManagementExchangeResourceNamespaceIdResourceActionId struct
func NewRoleManagementExchangeResourceNamespaceIdResourceActionID(unifiedRbacResourceNamespaceId string, unifiedRbacResourceActionId string) RoleManagementExchangeResourceNamespaceIdResourceActionId {
	return RoleManagementExchangeResourceNamespaceIdResourceActionId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
		UnifiedRbacResourceActionId:    unifiedRbacResourceActionId,
	}
}

// ParseRoleManagementExchangeResourceNamespaceIdResourceActionID parses 'input' into a RoleManagementExchangeResourceNamespaceIdResourceActionId
func ParseRoleManagementExchangeResourceNamespaceIdResourceActionID(input string) (*RoleManagementExchangeResourceNamespaceIdResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementExchangeResourceNamespaceIdResourceActionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementExchangeResourceNamespaceIdResourceActionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementExchangeResourceNamespaceIdResourceActionIDInsensitively parses 'input' case-insensitively into a RoleManagementExchangeResourceNamespaceIdResourceActionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementExchangeResourceNamespaceIdResourceActionIDInsensitively(input string) (*RoleManagementExchangeResourceNamespaceIdResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementExchangeResourceNamespaceIdResourceActionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementExchangeResourceNamespaceIdResourceActionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementExchangeResourceNamespaceIdResourceActionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRbacResourceNamespaceId, ok = input.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", input)
	}

	if id.UnifiedRbacResourceActionId, ok = input.Parsed["unifiedRbacResourceActionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", input)
	}

	return nil
}

// ValidateRoleManagementExchangeResourceNamespaceIdResourceActionID checks that 'input' can be parsed as a Role Management Exchange Resource Namespace Id Resource Action ID
func ValidateRoleManagementExchangeResourceNamespaceIdResourceActionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementExchangeResourceNamespaceIdResourceActionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Exchange Resource Namespace Id Resource Action ID
func (id RoleManagementExchangeResourceNamespaceIdResourceActionId) ID() string {
	fmtString := "/roleManagement/exchange/resourceNamespaces/%s/resourceActions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId, id.UnifiedRbacResourceActionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Exchange Resource Namespace Id Resource Action ID
func (id RoleManagementExchangeResourceNamespaceIdResourceActionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("exchange", "exchange", "exchange"),
		resourceids.StaticSegment("resourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceId"),
		resourceids.StaticSegment("resourceActions", "resourceActions", "resourceActions"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceActionId", "unifiedRbacResourceActionId"),
	}
}

// String returns a human-readable description of this Role Management Exchange Resource Namespace Id Resource Action ID
func (id RoleManagementExchangeResourceNamespaceIdResourceActionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
		fmt.Sprintf("Unified Rbac Resource Action: %q", id.UnifiedRbacResourceActionId),
	}
	return fmt.Sprintf("Role Management Exchange Resource Namespace Id Resource Action (%s)", strings.Join(components, "\n"))
}
