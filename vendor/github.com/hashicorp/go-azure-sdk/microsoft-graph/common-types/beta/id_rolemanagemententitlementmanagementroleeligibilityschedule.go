package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementRoleEligibilityScheduleId{}

// RoleManagementEntitlementManagementRoleEligibilityScheduleId is a struct representing the Resource ID for a Role Management Entitlement Management Role Eligibility Schedule
type RoleManagementEntitlementManagementRoleEligibilityScheduleId struct {
	UnifiedRoleEligibilityScheduleId string
}

// NewRoleManagementEntitlementManagementRoleEligibilityScheduleID returns a new RoleManagementEntitlementManagementRoleEligibilityScheduleId struct
func NewRoleManagementEntitlementManagementRoleEligibilityScheduleID(unifiedRoleEligibilityScheduleId string) RoleManagementEntitlementManagementRoleEligibilityScheduleId {
	return RoleManagementEntitlementManagementRoleEligibilityScheduleId{
		UnifiedRoleEligibilityScheduleId: unifiedRoleEligibilityScheduleId,
	}
}

// ParseRoleManagementEntitlementManagementRoleEligibilityScheduleID parses 'input' into a RoleManagementEntitlementManagementRoleEligibilityScheduleId
func ParseRoleManagementEntitlementManagementRoleEligibilityScheduleID(input string) (*RoleManagementEntitlementManagementRoleEligibilityScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleEligibilityScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleEligibilityScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEntitlementManagementRoleEligibilityScheduleIDInsensitively parses 'input' case-insensitively into a RoleManagementEntitlementManagementRoleEligibilityScheduleId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEntitlementManagementRoleEligibilityScheduleIDInsensitively(input string) (*RoleManagementEntitlementManagementRoleEligibilityScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleEligibilityScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleEligibilityScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEntitlementManagementRoleEligibilityScheduleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleEligibilityScheduleId, ok = input.Parsed["unifiedRoleEligibilityScheduleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleEligibilityScheduleId", input)
	}

	return nil
}

// ValidateRoleManagementEntitlementManagementRoleEligibilityScheduleID checks that 'input' can be parsed as a Role Management Entitlement Management Role Eligibility Schedule ID
func ValidateRoleManagementEntitlementManagementRoleEligibilityScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEntitlementManagementRoleEligibilityScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Entitlement Management Role Eligibility Schedule ID
func (id RoleManagementEntitlementManagementRoleEligibilityScheduleId) ID() string {
	fmtString := "/roleManagement/entitlementManagement/roleEligibilitySchedules/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleEligibilityScheduleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Entitlement Management Role Eligibility Schedule ID
func (id RoleManagementEntitlementManagementRoleEligibilityScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("roleEligibilitySchedules", "roleEligibilitySchedules", "roleEligibilitySchedules"),
		resourceids.UserSpecifiedSegment("unifiedRoleEligibilityScheduleId", "unifiedRoleEligibilityScheduleId"),
	}
}

// String returns a human-readable description of this Role Management Entitlement Management Role Eligibility Schedule ID
func (id RoleManagementEntitlementManagementRoleEligibilityScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Eligibility Schedule: %q", id.UnifiedRoleEligibilityScheduleId),
	}
	return fmt.Sprintf("Role Management Entitlement Management Role Eligibility Schedule (%s)", strings.Join(components, "\n"))
}
