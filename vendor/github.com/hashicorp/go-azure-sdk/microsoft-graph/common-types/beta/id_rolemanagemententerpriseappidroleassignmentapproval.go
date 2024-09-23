package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleAssignmentApprovalId{}

// RoleManagementEnterpriseAppIdRoleAssignmentApprovalId is a struct representing the Resource ID for a Role Management Enterprise App Id Role Assignment Approval
type RoleManagementEnterpriseAppIdRoleAssignmentApprovalId struct {
	RbacApplicationId string
	ApprovalId        string
}

// NewRoleManagementEnterpriseAppIdRoleAssignmentApprovalID returns a new RoleManagementEnterpriseAppIdRoleAssignmentApprovalId struct
func NewRoleManagementEnterpriseAppIdRoleAssignmentApprovalID(rbacApplicationId string, approvalId string) RoleManagementEnterpriseAppIdRoleAssignmentApprovalId {
	return RoleManagementEnterpriseAppIdRoleAssignmentApprovalId{
		RbacApplicationId: rbacApplicationId,
		ApprovalId:        approvalId,
	}
}

// ParseRoleManagementEnterpriseAppIdRoleAssignmentApprovalID parses 'input' into a RoleManagementEnterpriseAppIdRoleAssignmentApprovalId
func ParseRoleManagementEnterpriseAppIdRoleAssignmentApprovalID(input string) (*RoleManagementEnterpriseAppIdRoleAssignmentApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleAssignmentApprovalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleAssignmentApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppIdRoleAssignmentApprovalIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppIdRoleAssignmentApprovalId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppIdRoleAssignmentApprovalIDInsensitively(input string) (*RoleManagementEnterpriseAppIdRoleAssignmentApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleAssignmentApprovalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleAssignmentApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEnterpriseAppIdRoleAssignmentApprovalId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RbacApplicationId, ok = input.Parsed["rbacApplicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", input)
	}

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	return nil
}

// ValidateRoleManagementEnterpriseAppIdRoleAssignmentApprovalID checks that 'input' can be parsed as a Role Management Enterprise App Id Role Assignment Approval ID
func ValidateRoleManagementEnterpriseAppIdRoleAssignmentApprovalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppIdRoleAssignmentApprovalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Id Role Assignment Approval ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentApprovalId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleAssignmentApprovals/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.ApprovalId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Id Role Assignment Approval ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentApprovalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("enterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationId"),
		resourceids.StaticSegment("roleAssignmentApprovals", "roleAssignmentApprovals", "roleAssignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalId"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Id Role Assignment Approval ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentApprovalId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Approval: %q", id.ApprovalId),
	}
	return fmt.Sprintf("Role Management Enterprise App Id Role Assignment Approval (%s)", strings.Join(components, "\n"))
}
