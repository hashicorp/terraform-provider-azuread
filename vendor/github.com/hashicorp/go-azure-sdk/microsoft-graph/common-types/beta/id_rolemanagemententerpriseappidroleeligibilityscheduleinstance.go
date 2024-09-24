package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId{}

// RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId is a struct representing the Resource ID for a Role Management Enterprise App Id Role Eligibility Schedule Instance
type RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId struct {
	RbacApplicationId                        string
	UnifiedRoleEligibilityScheduleInstanceId string
}

// NewRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceID returns a new RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId struct
func NewRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceID(rbacApplicationId string, unifiedRoleEligibilityScheduleInstanceId string) RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId {
	return RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId{
		RbacApplicationId:                        rbacApplicationId,
		UnifiedRoleEligibilityScheduleInstanceId: unifiedRoleEligibilityScheduleInstanceId,
	}
}

// ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceID parses 'input' into a RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId
func ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceID(input string) (*RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceIDInsensitively(input string) (*RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RbacApplicationId, ok = input.Parsed["rbacApplicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", input)
	}

	if id.UnifiedRoleEligibilityScheduleInstanceId, ok = input.Parsed["unifiedRoleEligibilityScheduleInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleEligibilityScheduleInstanceId", input)
	}

	return nil
}

// ValidateRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceID checks that 'input' can be parsed as a Role Management Enterprise App Id Role Eligibility Schedule Instance ID
func ValidateRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Id Role Eligibility Schedule Instance ID
func (id RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleEligibilityScheduleInstances/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleEligibilityScheduleInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Id Role Eligibility Schedule Instance ID
func (id RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("enterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationId"),
		resourceids.StaticSegment("roleEligibilityScheduleInstances", "roleEligibilityScheduleInstances", "roleEligibilityScheduleInstances"),
		resourceids.UserSpecifiedSegment("unifiedRoleEligibilityScheduleInstanceId", "unifiedRoleEligibilityScheduleInstanceId"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Id Role Eligibility Schedule Instance ID
func (id RoleManagementEnterpriseAppIdRoleEligibilityScheduleInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Eligibility Schedule Instance: %q", id.UnifiedRoleEligibilityScheduleInstanceId),
	}
	return fmt.Sprintf("Role Management Enterprise App Id Role Eligibility Schedule Instance (%s)", strings.Join(components, "\n"))
}
