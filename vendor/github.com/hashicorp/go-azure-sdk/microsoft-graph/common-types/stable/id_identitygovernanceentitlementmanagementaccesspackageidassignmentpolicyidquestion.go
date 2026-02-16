package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId{}

// IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Id Assignment Policy Id Question
type IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId struct {
	AccessPackageId                 string
	AccessPackageAssignmentPolicyId string
	AccessPackageQuestionId         string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionID returns a new IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionID(accessPackageId string, accessPackageAssignmentPolicyId string, accessPackageQuestionId string) IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId {
	return IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId{
		AccessPackageId:                 accessPackageId,
		AccessPackageAssignmentPolicyId: accessPackageAssignmentPolicyId,
		AccessPackageQuestionId:         accessPackageQuestionId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	if id.AccessPackageAssignmentPolicyId, ok = input.Parsed["accessPackageAssignmentPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentPolicyId", input)
	}

	if id.AccessPackageQuestionId, ok = input.Parsed["accessPackageQuestionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageQuestionId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Id Assignment Policy Id Question ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Id Assignment Policy Id Question ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackages/%s/assignmentPolicies/%s/questions/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageId, id.AccessPackageAssignmentPolicyId, id.AccessPackageQuestionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Id Assignment Policy Id Question ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
		resourceids.StaticSegment("assignmentPolicies", "assignmentPolicies", "assignmentPolicies"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentPolicyId", "accessPackageAssignmentPolicyId"),
		resourceids.StaticSegment("questions", "questions", "questions"),
		resourceids.UserSpecifiedSegment("accessPackageQuestionId", "accessPackageQuestionId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Id Assignment Policy Id Question ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId) String() string {
	components := []string{
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
		fmt.Sprintf("Access Package Assignment Policy: %q", id.AccessPackageAssignmentPolicyId),
		fmt.Sprintf("Access Package Question: %q", id.AccessPackageQuestionId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Id Assignment Policy Id Question (%s)", strings.Join(components, "\n"))
}
