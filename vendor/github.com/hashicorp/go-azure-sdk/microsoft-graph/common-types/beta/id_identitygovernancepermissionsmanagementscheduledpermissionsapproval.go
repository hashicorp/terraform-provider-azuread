package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId{}

// IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId is a struct representing the Resource ID for a Identity Governance Permissions Management Scheduled Permissions Approval
type IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId struct {
	ApprovalId string
}

// NewIdentityGovernancePermissionsManagementScheduledPermissionsApprovalID returns a new IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId struct
func NewIdentityGovernancePermissionsManagementScheduledPermissionsApprovalID(approvalId string) IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId {
	return IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId{
		ApprovalId: approvalId,
	}
}

// ParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalID parses 'input' into a IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId
func ParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalID(input string) (*IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIDInsensitively(input string) (*IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	return nil
}

// ValidateIdentityGovernancePermissionsManagementScheduledPermissionsApprovalID checks that 'input' can be parsed as a Identity Governance Permissions Management Scheduled Permissions Approval ID
func ValidateIdentityGovernancePermissionsManagementScheduledPermissionsApprovalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Permissions Management Scheduled Permissions Approval ID
func (id IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId) ID() string {
	fmtString := "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Permissions Management Scheduled Permissions Approval ID
func (id IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("permissionsManagement", "permissionsManagement", "permissionsManagement"),
		resourceids.StaticSegment("scheduledPermissionsApprovals", "scheduledPermissionsApprovals", "scheduledPermissionsApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalId"),
	}
}

// String returns a human-readable description of this Identity Governance Permissions Management Scheduled Permissions Approval ID
func (id IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
	}
	return fmt.Sprintf("Identity Governance Permissions Management Scheduled Permissions Approval (%s)", strings.Join(components, "\n"))
}
