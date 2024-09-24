package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId{}

// RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId is a struct representing the Resource ID for a Role Management Entitlement Management Role Assignment Approval Id Step
type RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId struct {
	ApprovalId     string
	ApprovalStepId string
}

// NewRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepID returns a new RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId struct
func NewRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepID(approvalId string, approvalStepId string) RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId {
	return RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId{
		ApprovalId:     approvalId,
		ApprovalStepId: approvalStepId,
	}
}

// ParseRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepID parses 'input' into a RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId
func ParseRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepID(input string) (*RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepIDInsensitively parses 'input' case-insensitively into a RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepIDInsensitively(input string) (*RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	if id.ApprovalStepId, ok = input.Parsed["approvalStepId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", input)
	}

	return nil
}

// ValidateRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepID checks that 'input' can be parsed as a Role Management Entitlement Management Role Assignment Approval Id Step ID
func ValidateRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Entitlement Management Role Assignment Approval Id Step ID
func (id RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId) ID() string {
	fmtString := "/roleManagement/entitlementManagement/roleAssignmentApprovals/%s/steps/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Entitlement Management Role Assignment Approval Id Step ID
func (id RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("roleAssignmentApprovals", "roleAssignmentApprovals", "roleAssignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalId"),
		resourceids.StaticSegment("steps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepId"),
	}
}

// String returns a human-readable description of this Role Management Entitlement Management Role Assignment Approval Id Step ID
func (id RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("Role Management Entitlement Management Role Assignment Approval Id Step (%s)", strings.Join(components, "\n"))
}
