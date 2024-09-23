package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryRoleAssignmentApprovalIdStepId{}

// RoleManagementDirectoryRoleAssignmentApprovalIdStepId is a struct representing the Resource ID for a Role Management Directory Role Assignment Approval Id Step
type RoleManagementDirectoryRoleAssignmentApprovalIdStepId struct {
	ApprovalId     string
	ApprovalStepId string
}

// NewRoleManagementDirectoryRoleAssignmentApprovalIdStepID returns a new RoleManagementDirectoryRoleAssignmentApprovalIdStepId struct
func NewRoleManagementDirectoryRoleAssignmentApprovalIdStepID(approvalId string, approvalStepId string) RoleManagementDirectoryRoleAssignmentApprovalIdStepId {
	return RoleManagementDirectoryRoleAssignmentApprovalIdStepId{
		ApprovalId:     approvalId,
		ApprovalStepId: approvalStepId,
	}
}

// ParseRoleManagementDirectoryRoleAssignmentApprovalIdStepID parses 'input' into a RoleManagementDirectoryRoleAssignmentApprovalIdStepId
func ParseRoleManagementDirectoryRoleAssignmentApprovalIdStepID(input string) (*RoleManagementDirectoryRoleAssignmentApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleAssignmentApprovalIdStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleAssignmentApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDirectoryRoleAssignmentApprovalIdStepIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryRoleAssignmentApprovalIdStepId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryRoleAssignmentApprovalIdStepIDInsensitively(input string) (*RoleManagementDirectoryRoleAssignmentApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleAssignmentApprovalIdStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleAssignmentApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDirectoryRoleAssignmentApprovalIdStepId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	if id.ApprovalStepId, ok = input.Parsed["approvalStepId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", input)
	}

	return nil
}

// ValidateRoleManagementDirectoryRoleAssignmentApprovalIdStepID checks that 'input' can be parsed as a Role Management Directory Role Assignment Approval Id Step ID
func ValidateRoleManagementDirectoryRoleAssignmentApprovalIdStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryRoleAssignmentApprovalIdStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Role Assignment Approval Id Step ID
func (id RoleManagementDirectoryRoleAssignmentApprovalIdStepId) ID() string {
	fmtString := "/roleManagement/directory/roleAssignmentApprovals/%s/steps/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Role Assignment Approval Id Step ID
func (id RoleManagementDirectoryRoleAssignmentApprovalIdStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("roleAssignmentApprovals", "roleAssignmentApprovals", "roleAssignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalId"),
		resourceids.StaticSegment("steps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepId"),
	}
}

// String returns a human-readable description of this Role Management Directory Role Assignment Approval Id Step ID
func (id RoleManagementDirectoryRoleAssignmentApprovalIdStepId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("Role Management Directory Role Assignment Approval Id Step (%s)", strings.Join(components, "\n"))
}
