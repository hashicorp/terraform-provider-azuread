package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId{}

// IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId is a struct representing the Resource ID for a Identity Governance Entitlement Management Assignment Policy Id Question
type IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId struct {
	AccessPackageAssignmentPolicyId string
	AccessPackageQuestionId         string
}

// NewIdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionID returns a new IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId struct
func NewIdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionID(accessPackageAssignmentPolicyId string, accessPackageQuestionId string) IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId {
	return IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId{
		AccessPackageAssignmentPolicyId: accessPackageAssignmentPolicyId,
		AccessPackageQuestionId:         accessPackageQuestionId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionID parses 'input' into a IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId
func ParseIdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionID(input string) (*IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentPolicyId, ok = input.Parsed["accessPackageAssignmentPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentPolicyId", input)
	}

	if id.AccessPackageQuestionId, ok = input.Parsed["accessPackageQuestionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageQuestionId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionID checks that 'input' can be parsed as a Identity Governance Entitlement Management Assignment Policy Id Question ID
func ValidateIdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Assignment Policy Id Question ID
func (id IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/assignmentPolicies/%s/questions/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentPolicyId, id.AccessPackageQuestionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Assignment Policy Id Question ID
func (id IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("assignmentPolicies", "assignmentPolicies", "assignmentPolicies"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentPolicyId", "accessPackageAssignmentPolicyId"),
		resourceids.StaticSegment("questions", "questions", "questions"),
		resourceids.UserSpecifiedSegment("accessPackageQuestionId", "accessPackageQuestionId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Assignment Policy Id Question ID
func (id IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment Policy: %q", id.AccessPackageAssignmentPolicyId),
		fmt.Sprintf("Access Package Question: %q", id.AccessPackageQuestionId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Assignment Policy Id Question (%s)", strings.Join(components, "\n"))
}
