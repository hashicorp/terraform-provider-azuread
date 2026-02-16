package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId{}

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Assignment Approval Id Step
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId struct {
	ApprovalId     string
	ApprovalStepId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepID returns a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepID(approvalId string, approvalStepId string) IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId {
	return IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId{
		ApprovalId:     approvalId,
		ApprovalStepId: approvalStepId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	if id.ApprovalStepId, ok = input.Parsed["approvalStepId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Assignment Approval Id Step ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Assignment Approval Id Step ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/%s/steps/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Assignment Approval Id Step ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageAssignmentApprovals", "accessPackageAssignmentApprovals", "accessPackageAssignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalId"),
		resourceids.StaticSegment("steps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Assignment Approval Id Step ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Assignment Approval Id Step (%s)", strings.Join(components, "\n"))
}
