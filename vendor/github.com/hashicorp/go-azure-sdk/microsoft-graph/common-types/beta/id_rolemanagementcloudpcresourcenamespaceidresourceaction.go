package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementCloudPCResourceNamespaceIdResourceActionId{}

// RoleManagementCloudPCResourceNamespaceIdResourceActionId is a struct representing the Resource ID for a Role Management Cloud PC Resource Namespace Id Resource Action
type RoleManagementCloudPCResourceNamespaceIdResourceActionId struct {
	UnifiedRbacResourceNamespaceId string
	UnifiedRbacResourceActionId    string
}

// NewRoleManagementCloudPCResourceNamespaceIdResourceActionID returns a new RoleManagementCloudPCResourceNamespaceIdResourceActionId struct
func NewRoleManagementCloudPCResourceNamespaceIdResourceActionID(unifiedRbacResourceNamespaceId string, unifiedRbacResourceActionId string) RoleManagementCloudPCResourceNamespaceIdResourceActionId {
	return RoleManagementCloudPCResourceNamespaceIdResourceActionId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
		UnifiedRbacResourceActionId:    unifiedRbacResourceActionId,
	}
}

// ParseRoleManagementCloudPCResourceNamespaceIdResourceActionID parses 'input' into a RoleManagementCloudPCResourceNamespaceIdResourceActionId
func ParseRoleManagementCloudPCResourceNamespaceIdResourceActionID(input string) (*RoleManagementCloudPCResourceNamespaceIdResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementCloudPCResourceNamespaceIdResourceActionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementCloudPCResourceNamespaceIdResourceActionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementCloudPCResourceNamespaceIdResourceActionIDInsensitively parses 'input' case-insensitively into a RoleManagementCloudPCResourceNamespaceIdResourceActionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementCloudPCResourceNamespaceIdResourceActionIDInsensitively(input string) (*RoleManagementCloudPCResourceNamespaceIdResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementCloudPCResourceNamespaceIdResourceActionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementCloudPCResourceNamespaceIdResourceActionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementCloudPCResourceNamespaceIdResourceActionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRbacResourceNamespaceId, ok = input.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", input)
	}

	if id.UnifiedRbacResourceActionId, ok = input.Parsed["unifiedRbacResourceActionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", input)
	}

	return nil
}

// ValidateRoleManagementCloudPCResourceNamespaceIdResourceActionID checks that 'input' can be parsed as a Role Management Cloud PC Resource Namespace Id Resource Action ID
func ValidateRoleManagementCloudPCResourceNamespaceIdResourceActionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementCloudPCResourceNamespaceIdResourceActionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Cloud PC Resource Namespace Id Resource Action ID
func (id RoleManagementCloudPCResourceNamespaceIdResourceActionId) ID() string {
	fmtString := "/roleManagement/cloudPC/resourceNamespaces/%s/resourceActions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId, id.UnifiedRbacResourceActionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Cloud PC Resource Namespace Id Resource Action ID
func (id RoleManagementCloudPCResourceNamespaceIdResourceActionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("cloudPC", "cloudPC", "cloudPC"),
		resourceids.StaticSegment("resourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceId"),
		resourceids.StaticSegment("resourceActions", "resourceActions", "resourceActions"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceActionId", "unifiedRbacResourceActionId"),
	}
}

// String returns a human-readable description of this Role Management Cloud PC Resource Namespace Id Resource Action ID
func (id RoleManagementCloudPCResourceNamespaceIdResourceActionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
		fmt.Sprintf("Unified Rbac Resource Action: %q", id.UnifiedRbacResourceActionId),
	}
	return fmt.Sprintf("Role Management Cloud PC Resource Namespace Id Resource Action (%s)", strings.Join(components, "\n"))
}
