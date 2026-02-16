package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId{}

// RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId is a struct representing the Resource ID for a Role Management Enterprise App Id Role Eligibility Schedule Request
type RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId struct {
	RbacApplicationId                       string
	UnifiedRoleEligibilityScheduleRequestId string
}

// NewRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID returns a new RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId struct
func NewRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID(rbacApplicationId string, unifiedRoleEligibilityScheduleRequestId string) RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId {
	return RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId{
		RbacApplicationId:                       rbacApplicationId,
		UnifiedRoleEligibilityScheduleRequestId: unifiedRoleEligibilityScheduleRequestId,
	}
}

// ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID parses 'input' into a RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId
func ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID(input string) (*RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestIDInsensitively(input string) (*RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RbacApplicationId, ok = input.Parsed["rbacApplicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", input)
	}

	if id.UnifiedRoleEligibilityScheduleRequestId, ok = input.Parsed["unifiedRoleEligibilityScheduleRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleEligibilityScheduleRequestId", input)
	}

	return nil
}

// ValidateRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID checks that 'input' can be parsed as a Role Management Enterprise App Id Role Eligibility Schedule Request ID
func ValidateRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Id Role Eligibility Schedule Request ID
func (id RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleEligibilityScheduleRequests/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleEligibilityScheduleRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Id Role Eligibility Schedule Request ID
func (id RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("enterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationId"),
		resourceids.StaticSegment("roleEligibilityScheduleRequests", "roleEligibilityScheduleRequests", "roleEligibilityScheduleRequests"),
		resourceids.UserSpecifiedSegment("unifiedRoleEligibilityScheduleRequestId", "unifiedRoleEligibilityScheduleRequestId"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Id Role Eligibility Schedule Request ID
func (id RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Eligibility Schedule Request: %q", id.UnifiedRoleEligibilityScheduleRequestId),
	}
	return fmt.Sprintf("Role Management Enterprise App Id Role Eligibility Schedule Request (%s)", strings.Join(components, "\n"))
}
