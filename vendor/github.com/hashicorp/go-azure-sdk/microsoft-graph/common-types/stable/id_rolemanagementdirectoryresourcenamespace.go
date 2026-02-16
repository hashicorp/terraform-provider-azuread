package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryResourceNamespaceId{}

// RoleManagementDirectoryResourceNamespaceId is a struct representing the Resource ID for a Role Management Directory Resource Namespace
type RoleManagementDirectoryResourceNamespaceId struct {
	UnifiedRbacResourceNamespaceId string
}

// NewRoleManagementDirectoryResourceNamespaceID returns a new RoleManagementDirectoryResourceNamespaceId struct
func NewRoleManagementDirectoryResourceNamespaceID(unifiedRbacResourceNamespaceId string) RoleManagementDirectoryResourceNamespaceId {
	return RoleManagementDirectoryResourceNamespaceId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
	}
}

// ParseRoleManagementDirectoryResourceNamespaceID parses 'input' into a RoleManagementDirectoryResourceNamespaceId
func ParseRoleManagementDirectoryResourceNamespaceID(input string) (*RoleManagementDirectoryResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryResourceNamespaceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryResourceNamespaceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDirectoryResourceNamespaceIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryResourceNamespaceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryResourceNamespaceIDInsensitively(input string) (*RoleManagementDirectoryResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryResourceNamespaceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryResourceNamespaceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDirectoryResourceNamespaceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRbacResourceNamespaceId, ok = input.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", input)
	}

	return nil
}

// ValidateRoleManagementDirectoryResourceNamespaceID checks that 'input' can be parsed as a Role Management Directory Resource Namespace ID
func ValidateRoleManagementDirectoryResourceNamespaceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryResourceNamespaceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Resource Namespace ID
func (id RoleManagementDirectoryResourceNamespaceId) ID() string {
	fmtString := "/roleManagement/directory/resourceNamespaces/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Resource Namespace ID
func (id RoleManagementDirectoryResourceNamespaceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("resourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceId"),
	}
}

// String returns a human-readable description of this Role Management Directory Resource Namespace ID
func (id RoleManagementDirectoryResourceNamespaceId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
	}
	return fmt.Sprintf("Role Management Directory Resource Namespace (%s)", strings.Join(components, "\n"))
}
