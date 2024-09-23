package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId{}

// IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId is a struct representing the Resource ID for a Identity Governance Privileged Access Group Assignment Approval
type IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId struct {
	ApprovalId string
}

// NewIdentityGovernancePrivilegedAccessGroupAssignmentApprovalID returns a new IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId struct
func NewIdentityGovernancePrivilegedAccessGroupAssignmentApprovalID(approvalId string) IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId {
	return IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId{
		ApprovalId: approvalId,
	}
}

// ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalID parses 'input' into a IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId
func ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalID(input string) (*IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIDInsensitively(input string) (*IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccessGroupAssignmentApprovalID checks that 'input' can be parsed as a Identity Governance Privileged Access Group Assignment Approval ID
func ValidateIdentityGovernancePrivilegedAccessGroupAssignmentApprovalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Access Group Assignment Approval ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/assignmentApprovals/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Access Group Assignment Approval ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("assignmentApprovals", "assignmentApprovals", "assignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalId"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Access Group Assignment Approval ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
	}
	return fmt.Sprintf("Identity Governance Privileged Access Group Assignment Approval (%s)", strings.Join(components, "\n"))
}
