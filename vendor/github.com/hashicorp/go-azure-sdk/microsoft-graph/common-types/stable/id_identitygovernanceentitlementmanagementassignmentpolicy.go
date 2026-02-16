package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAssignmentPolicyId{}

// IdentityGovernanceEntitlementManagementAssignmentPolicyId is a struct representing the Resource ID for a Identity Governance Entitlement Management Assignment Policy
type IdentityGovernanceEntitlementManagementAssignmentPolicyId struct {
	AccessPackageAssignmentPolicyId string
}

// NewIdentityGovernanceEntitlementManagementAssignmentPolicyID returns a new IdentityGovernanceEntitlementManagementAssignmentPolicyId struct
func NewIdentityGovernanceEntitlementManagementAssignmentPolicyID(accessPackageAssignmentPolicyId string) IdentityGovernanceEntitlementManagementAssignmentPolicyId {
	return IdentityGovernanceEntitlementManagementAssignmentPolicyId{
		AccessPackageAssignmentPolicyId: accessPackageAssignmentPolicyId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAssignmentPolicyID parses 'input' into a IdentityGovernanceEntitlementManagementAssignmentPolicyId
func ParseIdentityGovernanceEntitlementManagementAssignmentPolicyID(input string) (*IdentityGovernanceEntitlementManagementAssignmentPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAssignmentPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAssignmentPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAssignmentPolicyIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAssignmentPolicyId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAssignmentPolicyIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAssignmentPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAssignmentPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAssignmentPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAssignmentPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentPolicyId, ok = input.Parsed["accessPackageAssignmentPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentPolicyId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAssignmentPolicyID checks that 'input' can be parsed as a Identity Governance Entitlement Management Assignment Policy ID
func ValidateIdentityGovernanceEntitlementManagementAssignmentPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAssignmentPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Assignment Policy ID
func (id IdentityGovernanceEntitlementManagementAssignmentPolicyId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/assignmentPolicies/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Assignment Policy ID
func (id IdentityGovernanceEntitlementManagementAssignmentPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("assignmentPolicies", "assignmentPolicies", "assignmentPolicies"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentPolicyId", "accessPackageAssignmentPolicyId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Assignment Policy ID
func (id IdentityGovernanceEntitlementManagementAssignmentPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment Policy: %q", id.AccessPackageAssignmentPolicyId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Assignment Policy (%s)", strings.Join(components, "\n"))
}
