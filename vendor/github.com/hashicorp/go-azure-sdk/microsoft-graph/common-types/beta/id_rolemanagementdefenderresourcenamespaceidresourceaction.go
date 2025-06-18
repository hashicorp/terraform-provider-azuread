package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDefenderResourceNamespaceIdResourceActionId{}

// RoleManagementDefenderResourceNamespaceIdResourceActionId is a struct representing the Resource ID for a Role Management Defender Resource Namespace Id Resource Action
type RoleManagementDefenderResourceNamespaceIdResourceActionId struct {
	UnifiedRbacResourceNamespaceId string
	UnifiedRbacResourceActionId    string
}

// NewRoleManagementDefenderResourceNamespaceIdResourceActionID returns a new RoleManagementDefenderResourceNamespaceIdResourceActionId struct
func NewRoleManagementDefenderResourceNamespaceIdResourceActionID(unifiedRbacResourceNamespaceId string, unifiedRbacResourceActionId string) RoleManagementDefenderResourceNamespaceIdResourceActionId {
	return RoleManagementDefenderResourceNamespaceIdResourceActionId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
		UnifiedRbacResourceActionId:    unifiedRbacResourceActionId,
	}
}

// ParseRoleManagementDefenderResourceNamespaceIdResourceActionID parses 'input' into a RoleManagementDefenderResourceNamespaceIdResourceActionId
func ParseRoleManagementDefenderResourceNamespaceIdResourceActionID(input string) (*RoleManagementDefenderResourceNamespaceIdResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderResourceNamespaceIdResourceActionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderResourceNamespaceIdResourceActionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDefenderResourceNamespaceIdResourceActionIDInsensitively parses 'input' case-insensitively into a RoleManagementDefenderResourceNamespaceIdResourceActionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDefenderResourceNamespaceIdResourceActionIDInsensitively(input string) (*RoleManagementDefenderResourceNamespaceIdResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderResourceNamespaceIdResourceActionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderResourceNamespaceIdResourceActionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDefenderResourceNamespaceIdResourceActionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRbacResourceNamespaceId, ok = input.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", input)
	}

	if id.UnifiedRbacResourceActionId, ok = input.Parsed["unifiedRbacResourceActionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", input)
	}

	return nil
}

// ValidateRoleManagementDefenderResourceNamespaceIdResourceActionID checks that 'input' can be parsed as a Role Management Defender Resource Namespace Id Resource Action ID
func ValidateRoleManagementDefenderResourceNamespaceIdResourceActionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDefenderResourceNamespaceIdResourceActionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Defender Resource Namespace Id Resource Action ID
func (id RoleManagementDefenderResourceNamespaceIdResourceActionId) ID() string {
	fmtString := "/roleManagement/defender/resourceNamespaces/%s/resourceActions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId, id.UnifiedRbacResourceActionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Defender Resource Namespace Id Resource Action ID
func (id RoleManagementDefenderResourceNamespaceIdResourceActionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("defender", "defender", "defender"),
		resourceids.StaticSegment("resourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceId"),
		resourceids.StaticSegment("resourceActions", "resourceActions", "resourceActions"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceActionId", "unifiedRbacResourceActionId"),
	}
}

// String returns a human-readable description of this Role Management Defender Resource Namespace Id Resource Action ID
func (id RoleManagementDefenderResourceNamespaceIdResourceActionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
		fmt.Sprintf("Unified Rbac Resource Action: %q", id.UnifiedRbacResourceActionId),
	}
	return fmt.Sprintf("Role Management Defender Resource Namespace Id Resource Action (%s)", strings.Join(components, "\n"))
}
