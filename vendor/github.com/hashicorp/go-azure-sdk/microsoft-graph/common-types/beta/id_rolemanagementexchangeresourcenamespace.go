package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementExchangeResourceNamespaceId{}

// RoleManagementExchangeResourceNamespaceId is a struct representing the Resource ID for a Role Management Exchange Resource Namespace
type RoleManagementExchangeResourceNamespaceId struct {
	UnifiedRbacResourceNamespaceId string
}

// NewRoleManagementExchangeResourceNamespaceID returns a new RoleManagementExchangeResourceNamespaceId struct
func NewRoleManagementExchangeResourceNamespaceID(unifiedRbacResourceNamespaceId string) RoleManagementExchangeResourceNamespaceId {
	return RoleManagementExchangeResourceNamespaceId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
	}
}

// ParseRoleManagementExchangeResourceNamespaceID parses 'input' into a RoleManagementExchangeResourceNamespaceId
func ParseRoleManagementExchangeResourceNamespaceID(input string) (*RoleManagementExchangeResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementExchangeResourceNamespaceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementExchangeResourceNamespaceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementExchangeResourceNamespaceIDInsensitively parses 'input' case-insensitively into a RoleManagementExchangeResourceNamespaceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementExchangeResourceNamespaceIDInsensitively(input string) (*RoleManagementExchangeResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementExchangeResourceNamespaceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementExchangeResourceNamespaceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementExchangeResourceNamespaceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRbacResourceNamespaceId, ok = input.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", input)
	}

	return nil
}

// ValidateRoleManagementExchangeResourceNamespaceID checks that 'input' can be parsed as a Role Management Exchange Resource Namespace ID
func ValidateRoleManagementExchangeResourceNamespaceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementExchangeResourceNamespaceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Exchange Resource Namespace ID
func (id RoleManagementExchangeResourceNamespaceId) ID() string {
	fmtString := "/roleManagement/exchange/resourceNamespaces/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Exchange Resource Namespace ID
func (id RoleManagementExchangeResourceNamespaceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("exchange", "exchange", "exchange"),
		resourceids.StaticSegment("resourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceId"),
	}
}

// String returns a human-readable description of this Role Management Exchange Resource Namespace ID
func (id RoleManagementExchangeResourceNamespaceId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
	}
	return fmt.Sprintf("Role Management Exchange Resource Namespace (%s)", strings.Join(components, "\n"))
}
