package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId{}

// RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId is a struct representing the Resource ID for a Role Management Enterprise App Id Role Assignment Approval Id Step
type RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId struct {
	RbacApplicationId string
	ApprovalId        string
	ApprovalStepId    string
}

// NewRoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepID returns a new RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId struct
func NewRoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepID(rbacApplicationId string, approvalId string, approvalStepId string) RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId {
	return RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId{
		RbacApplicationId: rbacApplicationId,
		ApprovalId:        approvalId,
		ApprovalStepId:    approvalStepId,
	}
}

// ParseRoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepID parses 'input' into a RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId
func ParseRoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepID(input string) (*RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepIDInsensitively(input string) (*RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RbacApplicationId, ok = input.Parsed["rbacApplicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", input)
	}

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	if id.ApprovalStepId, ok = input.Parsed["approvalStepId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", input)
	}

	return nil
}

// ValidateRoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepID checks that 'input' can be parsed as a Role Management Enterprise App Id Role Assignment Approval Id Step ID
func ValidateRoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Id Role Assignment Approval Id Step ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleAssignmentApprovals/%s/steps/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.ApprovalId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Id Role Assignment Approval Id Step ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("enterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationId"),
		resourceids.StaticSegment("roleAssignmentApprovals", "roleAssignmentApprovals", "roleAssignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalId"),
		resourceids.StaticSegment("steps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepId"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Id Role Assignment Approval Id Step ID
func (id RoleManagementEnterpriseAppIdRoleAssignmentApprovalIdStepId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("Role Management Enterprise App Id Role Assignment Approval Id Step (%s)", strings.Join(components, "\n"))
}
