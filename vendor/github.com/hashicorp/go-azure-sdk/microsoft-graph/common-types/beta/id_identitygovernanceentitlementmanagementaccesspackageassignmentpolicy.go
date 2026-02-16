package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId{}

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Assignment Policy
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId struct {
	AccessPackageAssignmentPolicyId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID returns a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID(accessPackageAssignmentPolicyId string) IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId {
	return IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId{
		AccessPackageAssignmentPolicyId: accessPackageAssignmentPolicyId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentPolicyId, ok = input.Parsed["accessPackageAssignmentPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentPolicyId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Assignment Policy ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Assignment Policy ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageAssignmentPolicies/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Assignment Policy ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageAssignmentPolicies", "accessPackageAssignmentPolicies", "accessPackageAssignmentPolicies"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentPolicyId", "accessPackageAssignmentPolicyId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Assignment Policy ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment Policy: %q", id.AccessPackageAssignmentPolicyId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Assignment Policy (%s)", strings.Join(components, "\n"))
}
