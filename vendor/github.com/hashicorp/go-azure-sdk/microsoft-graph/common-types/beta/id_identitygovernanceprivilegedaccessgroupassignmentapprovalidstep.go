package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId{}

// IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId is a struct representing the Resource ID for a Identity Governance Privileged Access Group Assignment Approval Id Step
type IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId struct {
	ApprovalId     string
	ApprovalStepId string
}

// NewIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepID returns a new IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId struct
func NewIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepID(approvalId string, approvalStepId string) IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId {
	return IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId{
		ApprovalId:     approvalId,
		ApprovalStepId: approvalStepId,
	}
}

// ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepID parses 'input' into a IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId
func ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepID(input string) (*IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepIDInsensitively(input string) (*IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	if id.ApprovalStepId, ok = input.Parsed["approvalStepId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepID checks that 'input' can be parsed as a Identity Governance Privileged Access Group Assignment Approval Id Step ID
func ValidateIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Access Group Assignment Approval Id Step ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/assignmentApprovals/%s/steps/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Access Group Assignment Approval Id Step ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("assignmentApprovals", "assignmentApprovals", "assignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalId"),
		resourceids.StaticSegment("steps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepId"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Access Group Assignment Approval Id Step ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("Identity Governance Privileged Access Group Assignment Approval Id Step (%s)", strings.Join(components, "\n"))
}
