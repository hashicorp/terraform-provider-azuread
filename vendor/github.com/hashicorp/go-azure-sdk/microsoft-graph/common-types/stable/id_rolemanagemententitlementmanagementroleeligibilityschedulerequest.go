package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId{}

// RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId is a struct representing the Resource ID for a Role Management Entitlement Management Role Eligibility Schedule Request
type RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId struct {
	UnifiedRoleEligibilityScheduleRequestId string
}

// NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestID returns a new RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId struct
func NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestID(unifiedRoleEligibilityScheduleRequestId string) RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId {
	return RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId{
		UnifiedRoleEligibilityScheduleRequestId: unifiedRoleEligibilityScheduleRequestId,
	}
}

// ParseRoleManagementEntitlementManagementRoleEligibilityScheduleRequestID parses 'input' into a RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId
func ParseRoleManagementEntitlementManagementRoleEligibilityScheduleRequestID(input string) (*RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEntitlementManagementRoleEligibilityScheduleRequestIDInsensitively parses 'input' case-insensitively into a RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEntitlementManagementRoleEligibilityScheduleRequestIDInsensitively(input string) (*RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleEligibilityScheduleRequestId, ok = input.Parsed["unifiedRoleEligibilityScheduleRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleEligibilityScheduleRequestId", input)
	}

	return nil
}

// ValidateRoleManagementEntitlementManagementRoleEligibilityScheduleRequestID checks that 'input' can be parsed as a Role Management Entitlement Management Role Eligibility Schedule Request ID
func ValidateRoleManagementEntitlementManagementRoleEligibilityScheduleRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEntitlementManagementRoleEligibilityScheduleRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Entitlement Management Role Eligibility Schedule Request ID
func (id RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId) ID() string {
	fmtString := "/roleManagement/entitlementManagement/roleEligibilityScheduleRequests/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleEligibilityScheduleRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Entitlement Management Role Eligibility Schedule Request ID
func (id RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("roleEligibilityScheduleRequests", "roleEligibilityScheduleRequests", "roleEligibilityScheduleRequests"),
		resourceids.UserSpecifiedSegment("unifiedRoleEligibilityScheduleRequestId", "unifiedRoleEligibilityScheduleRequestId"),
	}
}

// String returns a human-readable description of this Role Management Entitlement Management Role Eligibility Schedule Request ID
func (id RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Eligibility Schedule Request: %q", id.UnifiedRoleEligibilityScheduleRequestId),
	}
	return fmt.Sprintf("Role Management Entitlement Management Role Eligibility Schedule Request (%s)", strings.Join(components, "\n"))
}
