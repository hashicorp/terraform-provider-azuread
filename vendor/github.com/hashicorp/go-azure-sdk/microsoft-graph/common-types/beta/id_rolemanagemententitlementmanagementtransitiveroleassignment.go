package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementTransitiveRoleAssignmentId{}

// RoleManagementEntitlementManagementTransitiveRoleAssignmentId is a struct representing the Resource ID for a Role Management Entitlement Management Transitive Role Assignment
type RoleManagementEntitlementManagementTransitiveRoleAssignmentId struct {
	UnifiedRoleAssignmentId string
}

// NewRoleManagementEntitlementManagementTransitiveRoleAssignmentID returns a new RoleManagementEntitlementManagementTransitiveRoleAssignmentId struct
func NewRoleManagementEntitlementManagementTransitiveRoleAssignmentID(unifiedRoleAssignmentId string) RoleManagementEntitlementManagementTransitiveRoleAssignmentId {
	return RoleManagementEntitlementManagementTransitiveRoleAssignmentId{
		UnifiedRoleAssignmentId: unifiedRoleAssignmentId,
	}
}

// ParseRoleManagementEntitlementManagementTransitiveRoleAssignmentID parses 'input' into a RoleManagementEntitlementManagementTransitiveRoleAssignmentId
func ParseRoleManagementEntitlementManagementTransitiveRoleAssignmentID(input string) (*RoleManagementEntitlementManagementTransitiveRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementTransitiveRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementTransitiveRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEntitlementManagementTransitiveRoleAssignmentIDInsensitively parses 'input' case-insensitively into a RoleManagementEntitlementManagementTransitiveRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEntitlementManagementTransitiveRoleAssignmentIDInsensitively(input string) (*RoleManagementEntitlementManagementTransitiveRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementTransitiveRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementTransitiveRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEntitlementManagementTransitiveRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentId, ok = input.Parsed["unifiedRoleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentId", input)
	}

	return nil
}

// ValidateRoleManagementEntitlementManagementTransitiveRoleAssignmentID checks that 'input' can be parsed as a Role Management Entitlement Management Transitive Role Assignment ID
func ValidateRoleManagementEntitlementManagementTransitiveRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEntitlementManagementTransitiveRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Entitlement Management Transitive Role Assignment ID
func (id RoleManagementEntitlementManagementTransitiveRoleAssignmentId) ID() string {
	fmtString := "/roleManagement/entitlementManagement/transitiveRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Entitlement Management Transitive Role Assignment ID
func (id RoleManagementEntitlementManagementTransitiveRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("transitiveRoleAssignments", "transitiveRoleAssignments", "transitiveRoleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentId", "unifiedRoleAssignmentId"),
	}
}

// String returns a human-readable description of this Role Management Entitlement Management Transitive Role Assignment ID
func (id RoleManagementEntitlementManagementTransitiveRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment: %q", id.UnifiedRoleAssignmentId),
	}
	return fmt.Sprintf("Role Management Entitlement Management Transitive Role Assignment (%s)", strings.Join(components, "\n"))
}
