package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementRoleDefinitionId{}

// RoleManagementEntitlementManagementRoleDefinitionId is a struct representing the Resource ID for a Role Management Entitlement Management Role Definition
type RoleManagementEntitlementManagementRoleDefinitionId struct {
	UnifiedRoleDefinitionId string
}

// NewRoleManagementEntitlementManagementRoleDefinitionID returns a new RoleManagementEntitlementManagementRoleDefinitionId struct
func NewRoleManagementEntitlementManagementRoleDefinitionID(unifiedRoleDefinitionId string) RoleManagementEntitlementManagementRoleDefinitionId {
	return RoleManagementEntitlementManagementRoleDefinitionId{
		UnifiedRoleDefinitionId: unifiedRoleDefinitionId,
	}
}

// ParseRoleManagementEntitlementManagementRoleDefinitionID parses 'input' into a RoleManagementEntitlementManagementRoleDefinitionId
func ParseRoleManagementEntitlementManagementRoleDefinitionID(input string) (*RoleManagementEntitlementManagementRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEntitlementManagementRoleDefinitionIDInsensitively parses 'input' case-insensitively into a RoleManagementEntitlementManagementRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEntitlementManagementRoleDefinitionIDInsensitively(input string) (*RoleManagementEntitlementManagementRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEntitlementManagementRoleDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleDefinitionId, ok = input.Parsed["unifiedRoleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", input)
	}

	return nil
}

// ValidateRoleManagementEntitlementManagementRoleDefinitionID checks that 'input' can be parsed as a Role Management Entitlement Management Role Definition ID
func ValidateRoleManagementEntitlementManagementRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEntitlementManagementRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Entitlement Management Role Definition ID
func (id RoleManagementEntitlementManagementRoleDefinitionId) ID() string {
	fmtString := "/roleManagement/entitlementManagement/roleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Entitlement Management Role Definition ID
func (id RoleManagementEntitlementManagementRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("roleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionId"),
	}
}

// String returns a human-readable description of this Role Management Entitlement Management Role Definition ID
func (id RoleManagementEntitlementManagementRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
	}
	return fmt.Sprintf("Role Management Entitlement Management Role Definition (%s)", strings.Join(components, "\n"))
}
