package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementCloudPCResourceNamespaceId{}

// RoleManagementCloudPCResourceNamespaceId is a struct representing the Resource ID for a Role Management Cloud PC Resource Namespace
type RoleManagementCloudPCResourceNamespaceId struct {
	UnifiedRbacResourceNamespaceId string
}

// NewRoleManagementCloudPCResourceNamespaceID returns a new RoleManagementCloudPCResourceNamespaceId struct
func NewRoleManagementCloudPCResourceNamespaceID(unifiedRbacResourceNamespaceId string) RoleManagementCloudPCResourceNamespaceId {
	return RoleManagementCloudPCResourceNamespaceId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
	}
}

// ParseRoleManagementCloudPCResourceNamespaceID parses 'input' into a RoleManagementCloudPCResourceNamespaceId
func ParseRoleManagementCloudPCResourceNamespaceID(input string) (*RoleManagementCloudPCResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementCloudPCResourceNamespaceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementCloudPCResourceNamespaceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementCloudPCResourceNamespaceIDInsensitively parses 'input' case-insensitively into a RoleManagementCloudPCResourceNamespaceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementCloudPCResourceNamespaceIDInsensitively(input string) (*RoleManagementCloudPCResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementCloudPCResourceNamespaceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementCloudPCResourceNamespaceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementCloudPCResourceNamespaceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRbacResourceNamespaceId, ok = input.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", input)
	}

	return nil
}

// ValidateRoleManagementCloudPCResourceNamespaceID checks that 'input' can be parsed as a Role Management Cloud PC Resource Namespace ID
func ValidateRoleManagementCloudPCResourceNamespaceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementCloudPCResourceNamespaceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Cloud PC Resource Namespace ID
func (id RoleManagementCloudPCResourceNamespaceId) ID() string {
	fmtString := "/roleManagement/cloudPC/resourceNamespaces/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Cloud PC Resource Namespace ID
func (id RoleManagementCloudPCResourceNamespaceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("cloudPC", "cloudPC", "cloudPC"),
		resourceids.StaticSegment("resourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceId"),
	}
}

// String returns a human-readable description of this Role Management Cloud PC Resource Namespace ID
func (id RoleManagementCloudPCResourceNamespaceId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
	}
	return fmt.Sprintf("Role Management Cloud PC Resource Namespace (%s)", strings.Join(components, "\n"))
}
