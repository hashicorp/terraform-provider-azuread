package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId{}

// RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId is a struct representing the Resource ID for a Role Management Entitlement Management Role Eligibility Schedule Instance
type RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId struct {
	UnifiedRoleEligibilityScheduleInstanceId string
}

// NewRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceID returns a new RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId struct
func NewRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceID(unifiedRoleEligibilityScheduleInstanceId string) RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId {
	return RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId{
		UnifiedRoleEligibilityScheduleInstanceId: unifiedRoleEligibilityScheduleInstanceId,
	}
}

// ParseRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceID parses 'input' into a RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId
func ParseRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceID(input string) (*RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceIDInsensitively parses 'input' case-insensitively into a RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceIDInsensitively(input string) (*RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleEligibilityScheduleInstanceId, ok = input.Parsed["unifiedRoleEligibilityScheduleInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleEligibilityScheduleInstanceId", input)
	}

	return nil
}

// ValidateRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceID checks that 'input' can be parsed as a Role Management Entitlement Management Role Eligibility Schedule Instance ID
func ValidateRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Entitlement Management Role Eligibility Schedule Instance ID
func (id RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId) ID() string {
	fmtString := "/roleManagement/entitlementManagement/roleEligibilityScheduleInstances/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleEligibilityScheduleInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Entitlement Management Role Eligibility Schedule Instance ID
func (id RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("roleEligibilityScheduleInstances", "roleEligibilityScheduleInstances", "roleEligibilityScheduleInstances"),
		resourceids.UserSpecifiedSegment("unifiedRoleEligibilityScheduleInstanceId", "unifiedRoleEligibilityScheduleInstanceId"),
	}
}

// String returns a human-readable description of this Role Management Entitlement Management Role Eligibility Schedule Instance ID
func (id RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Eligibility Schedule Instance: %q", id.UnifiedRoleEligibilityScheduleInstanceId),
	}
	return fmt.Sprintf("Role Management Entitlement Management Role Eligibility Schedule Instance (%s)", strings.Join(components, "\n"))
}
