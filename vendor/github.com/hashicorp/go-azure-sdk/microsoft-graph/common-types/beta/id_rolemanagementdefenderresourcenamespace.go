package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDefenderResourceNamespaceId{}

// RoleManagementDefenderResourceNamespaceId is a struct representing the Resource ID for a Role Management Defender Resource Namespace
type RoleManagementDefenderResourceNamespaceId struct {
	UnifiedRbacResourceNamespaceId string
}

// NewRoleManagementDefenderResourceNamespaceID returns a new RoleManagementDefenderResourceNamespaceId struct
func NewRoleManagementDefenderResourceNamespaceID(unifiedRbacResourceNamespaceId string) RoleManagementDefenderResourceNamespaceId {
	return RoleManagementDefenderResourceNamespaceId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
	}
}

// ParseRoleManagementDefenderResourceNamespaceID parses 'input' into a RoleManagementDefenderResourceNamespaceId
func ParseRoleManagementDefenderResourceNamespaceID(input string) (*RoleManagementDefenderResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderResourceNamespaceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderResourceNamespaceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDefenderResourceNamespaceIDInsensitively parses 'input' case-insensitively into a RoleManagementDefenderResourceNamespaceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDefenderResourceNamespaceIDInsensitively(input string) (*RoleManagementDefenderResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderResourceNamespaceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderResourceNamespaceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDefenderResourceNamespaceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRbacResourceNamespaceId, ok = input.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", input)
	}

	return nil
}

// ValidateRoleManagementDefenderResourceNamespaceID checks that 'input' can be parsed as a Role Management Defender Resource Namespace ID
func ValidateRoleManagementDefenderResourceNamespaceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDefenderResourceNamespaceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Defender Resource Namespace ID
func (id RoleManagementDefenderResourceNamespaceId) ID() string {
	fmtString := "/roleManagement/defender/resourceNamespaces/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Defender Resource Namespace ID
func (id RoleManagementDefenderResourceNamespaceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("defender", "defender", "defender"),
		resourceids.StaticSegment("resourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceId"),
	}
}

// String returns a human-readable description of this Role Management Defender Resource Namespace ID
func (id RoleManagementDefenderResourceNamespaceId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
	}
	return fmt.Sprintf("Role Management Defender Resource Namespace (%s)", strings.Join(components, "\n"))
}
