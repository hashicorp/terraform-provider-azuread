package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId{}

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Assignment Approval
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId struct {
	ApprovalId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalID returns a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalID(approvalId string) IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId {
	return IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId{
		ApprovalId: approvalId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Assignment Approval ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Assignment Approval ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Assignment Approval ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageAssignmentApprovals", "accessPackageAssignmentApprovals", "accessPackageAssignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Assignment Approval ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Assignment Approval (%s)", strings.Join(components, "\n"))
}
