package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleEligibilityScheduleId{}

// RoleManagementEnterpriseAppIdRoleEligibilityScheduleId is a struct representing the Resource ID for a Role Management Enterprise App Id Role Eligibility Schedule
type RoleManagementEnterpriseAppIdRoleEligibilityScheduleId struct {
	RbacApplicationId                string
	UnifiedRoleEligibilityScheduleId string
}

// NewRoleManagementEnterpriseAppIdRoleEligibilityScheduleID returns a new RoleManagementEnterpriseAppIdRoleEligibilityScheduleId struct
func NewRoleManagementEnterpriseAppIdRoleEligibilityScheduleID(rbacApplicationId string, unifiedRoleEligibilityScheduleId string) RoleManagementEnterpriseAppIdRoleEligibilityScheduleId {
	return RoleManagementEnterpriseAppIdRoleEligibilityScheduleId{
		RbacApplicationId:                rbacApplicationId,
		UnifiedRoleEligibilityScheduleId: unifiedRoleEligibilityScheduleId,
	}
}

// ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleID parses 'input' into a RoleManagementEnterpriseAppIdRoleEligibilityScheduleId
func ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleID(input string) (*RoleManagementEnterpriseAppIdRoleEligibilityScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleEligibilityScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleEligibilityScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppIdRoleEligibilityScheduleId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleIDInsensitively(input string) (*RoleManagementEnterpriseAppIdRoleEligibilityScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleEligibilityScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleEligibilityScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEnterpriseAppIdRoleEligibilityScheduleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RbacApplicationId, ok = input.Parsed["rbacApplicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", input)
	}

	if id.UnifiedRoleEligibilityScheduleId, ok = input.Parsed["unifiedRoleEligibilityScheduleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleEligibilityScheduleId", input)
	}

	return nil
}

// ValidateRoleManagementEnterpriseAppIdRoleEligibilityScheduleID checks that 'input' can be parsed as a Role Management Enterprise App Id Role Eligibility Schedule ID
func ValidateRoleManagementEnterpriseAppIdRoleEligibilityScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Id Role Eligibility Schedule ID
func (id RoleManagementEnterpriseAppIdRoleEligibilityScheduleId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleEligibilitySchedules/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleEligibilityScheduleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Id Role Eligibility Schedule ID
func (id RoleManagementEnterpriseAppIdRoleEligibilityScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("enterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationId"),
		resourceids.StaticSegment("roleEligibilitySchedules", "roleEligibilitySchedules", "roleEligibilitySchedules"),
		resourceids.UserSpecifiedSegment("unifiedRoleEligibilityScheduleId", "unifiedRoleEligibilityScheduleId"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Id Role Eligibility Schedule ID
func (id RoleManagementEnterpriseAppIdRoleEligibilityScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Eligibility Schedule: %q", id.UnifiedRoleEligibilityScheduleId),
	}
	return fmt.Sprintf("Role Management Enterprise App Id Role Eligibility Schedule (%s)", strings.Join(components, "\n"))
}
