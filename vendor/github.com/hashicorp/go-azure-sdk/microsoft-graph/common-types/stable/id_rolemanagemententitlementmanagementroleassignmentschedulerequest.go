package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId{}

// RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId is a struct representing the Resource ID for a Role Management Entitlement Management Role Assignment Schedule Request
type RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId struct {
	UnifiedRoleAssignmentScheduleRequestId string
}

// NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestID returns a new RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId struct
func NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestID(unifiedRoleAssignmentScheduleRequestId string) RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId {
	return RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId{
		UnifiedRoleAssignmentScheduleRequestId: unifiedRoleAssignmentScheduleRequestId,
	}
}

// ParseRoleManagementEntitlementManagementRoleAssignmentScheduleRequestID parses 'input' into a RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId
func ParseRoleManagementEntitlementManagementRoleAssignmentScheduleRequestID(input string) (*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEntitlementManagementRoleAssignmentScheduleRequestIDInsensitively parses 'input' case-insensitively into a RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEntitlementManagementRoleAssignmentScheduleRequestIDInsensitively(input string) (*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentScheduleRequestId, ok = input.Parsed["unifiedRoleAssignmentScheduleRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleRequestId", input)
	}

	return nil
}

// ValidateRoleManagementEntitlementManagementRoleAssignmentScheduleRequestID checks that 'input' can be parsed as a Role Management Entitlement Management Role Assignment Schedule Request ID
func ValidateRoleManagementEntitlementManagementRoleAssignmentScheduleRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEntitlementManagementRoleAssignmentScheduleRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Entitlement Management Role Assignment Schedule Request ID
func (id RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId) ID() string {
	fmtString := "/roleManagement/entitlementManagement/roleAssignmentScheduleRequests/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentScheduleRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Entitlement Management Role Assignment Schedule Request ID
func (id RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("roleAssignmentScheduleRequests", "roleAssignmentScheduleRequests", "roleAssignmentScheduleRequests"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentScheduleRequestId", "unifiedRoleAssignmentScheduleRequestId"),
	}
}

// String returns a human-readable description of this Role Management Entitlement Management Role Assignment Schedule Request ID
func (id RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Schedule Request: %q", id.UnifiedRoleAssignmentScheduleRequestId),
	}
	return fmt.Sprintf("Role Management Entitlement Management Role Assignment Schedule Request (%s)", strings.Join(components, "\n"))
}
