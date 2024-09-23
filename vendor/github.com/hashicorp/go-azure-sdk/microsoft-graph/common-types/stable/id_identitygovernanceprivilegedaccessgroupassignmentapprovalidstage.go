package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId{}

// IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId is a struct representing the Resource ID for a Identity Governance Privileged Access Group Assignment Approval Id Stage
type IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId struct {
	ApprovalId      string
	ApprovalStageId string
}

// NewIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageID returns a new IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId struct
func NewIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageID(approvalId string, approvalStageId string) IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId {
	return IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId{
		ApprovalId:      approvalId,
		ApprovalStageId: approvalStageId,
	}
}

// ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageID parses 'input' into a IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId
func ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageID(input string) (*IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageIDInsensitively(input string) (*IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	if id.ApprovalStageId, ok = input.Parsed["approvalStageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalStageId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageID checks that 'input' can be parsed as a Identity Governance Privileged Access Group Assignment Approval Id Stage ID
func ValidateIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Access Group Assignment Approval Id Stage ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/assignmentApprovals/%s/stages/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId, id.ApprovalStageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Access Group Assignment Approval Id Stage ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("assignmentApprovals", "assignmentApprovals", "assignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalId"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("approvalStageId", "approvalStageId"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Access Group Assignment Approval Id Stage ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Stage: %q", id.ApprovalStageId),
	}
	return fmt.Sprintf("Identity Governance Privileged Access Group Assignment Approval Id Stage (%s)", strings.Join(components, "\n"))
}
