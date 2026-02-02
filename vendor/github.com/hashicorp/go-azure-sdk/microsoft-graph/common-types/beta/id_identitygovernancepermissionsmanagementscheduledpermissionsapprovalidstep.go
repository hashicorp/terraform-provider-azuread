package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId{}

// IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId is a struct representing the Resource ID for a Identity Governance Permissions Management Scheduled Permissions Approval Id Step
type IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId struct {
	ApprovalId     string
	ApprovalStepId string
}

// NewIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID returns a new IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId struct
func NewIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID(approvalId string, approvalStepId string) IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId {
	return IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId{
		ApprovalId:     approvalId,
		ApprovalStepId: approvalStepId,
	}
}

// ParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID parses 'input' into a IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId
func ParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID(input string) (*IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepIDInsensitively(input string) (*IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	if id.ApprovalStepId, ok = input.Parsed["approvalStepId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", input)
	}

	return nil
}

// ValidateIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID checks that 'input' can be parsed as a Identity Governance Permissions Management Scheduled Permissions Approval Id Step ID
func ValidateIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Permissions Management Scheduled Permissions Approval Id Step ID
func (id IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId) ID() string {
	fmtString := "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/%s/steps/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Permissions Management Scheduled Permissions Approval Id Step ID
func (id IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("permissionsManagement", "permissionsManagement", "permissionsManagement"),
		resourceids.StaticSegment("scheduledPermissionsApprovals", "scheduledPermissionsApprovals", "scheduledPermissionsApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalId"),
		resourceids.StaticSegment("steps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepId"),
	}
}

// String returns a human-readable description of this Identity Governance Permissions Management Scheduled Permissions Approval Id Step ID
func (id IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("Identity Governance Permissions Management Scheduled Permissions Approval Id Step (%s)", strings.Join(components, "\n"))
}
