package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDeviceManagementResourceNamespaceId{}

// RoleManagementDeviceManagementResourceNamespaceId is a struct representing the Resource ID for a Role Management Device Management Resource Namespace
type RoleManagementDeviceManagementResourceNamespaceId struct {
	UnifiedRbacResourceNamespaceId string
}

// NewRoleManagementDeviceManagementResourceNamespaceID returns a new RoleManagementDeviceManagementResourceNamespaceId struct
func NewRoleManagementDeviceManagementResourceNamespaceID(unifiedRbacResourceNamespaceId string) RoleManagementDeviceManagementResourceNamespaceId {
	return RoleManagementDeviceManagementResourceNamespaceId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
	}
}

// ParseRoleManagementDeviceManagementResourceNamespaceID parses 'input' into a RoleManagementDeviceManagementResourceNamespaceId
func ParseRoleManagementDeviceManagementResourceNamespaceID(input string) (*RoleManagementDeviceManagementResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDeviceManagementResourceNamespaceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDeviceManagementResourceNamespaceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDeviceManagementResourceNamespaceIDInsensitively parses 'input' case-insensitively into a RoleManagementDeviceManagementResourceNamespaceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDeviceManagementResourceNamespaceIDInsensitively(input string) (*RoleManagementDeviceManagementResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDeviceManagementResourceNamespaceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDeviceManagementResourceNamespaceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDeviceManagementResourceNamespaceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRbacResourceNamespaceId, ok = input.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", input)
	}

	return nil
}

// ValidateRoleManagementDeviceManagementResourceNamespaceID checks that 'input' can be parsed as a Role Management Device Management Resource Namespace ID
func ValidateRoleManagementDeviceManagementResourceNamespaceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDeviceManagementResourceNamespaceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Device Management Resource Namespace ID
func (id RoleManagementDeviceManagementResourceNamespaceId) ID() string {
	fmtString := "/roleManagement/deviceManagement/resourceNamespaces/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Device Management Resource Namespace ID
func (id RoleManagementDeviceManagementResourceNamespaceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("resourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceId"),
	}
}

// String returns a human-readable description of this Role Management Device Management Resource Namespace ID
func (id RoleManagementDeviceManagementResourceNamespaceId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
	}
	return fmt.Sprintf("Role Management Device Management Resource Namespace (%s)", strings.Join(components, "\n"))
}
