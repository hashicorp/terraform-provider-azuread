package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId{}

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Assignment Approval Id Stage
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId struct {
	ApprovalId      string
	ApprovalStageId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageID returns a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageID(approvalId string, approvalStageId string) IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId {
	return IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId{
		ApprovalId:      approvalId,
		ApprovalStageId: approvalStageId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	if id.ApprovalStageId, ok = input.Parsed["approvalStageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalStageId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Assignment Approval Id Stage ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Assignment Approval Id Stage ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/%s/stages/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId, id.ApprovalStageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Assignment Approval Id Stage ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageAssignmentApprovals", "accessPackageAssignmentApprovals", "accessPackageAssignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalId"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("approvalStageId", "approvalStageId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Assignment Approval Id Stage ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Stage: %q", id.ApprovalStageId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Assignment Approval Id Stage (%s)", strings.Join(components, "\n"))
}
