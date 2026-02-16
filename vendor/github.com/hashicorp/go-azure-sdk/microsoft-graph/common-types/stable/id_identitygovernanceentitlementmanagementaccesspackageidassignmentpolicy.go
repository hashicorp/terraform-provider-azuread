package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId{}

// IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Id Assignment Policy
type IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId struct {
	AccessPackageId                 string
	AccessPackageAssignmentPolicyId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyID returns a new IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyID(accessPackageId string, accessPackageAssignmentPolicyId string) IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId {
	return IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId{
		AccessPackageId:                 accessPackageId,
		AccessPackageAssignmentPolicyId: accessPackageAssignmentPolicyId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	if id.AccessPackageAssignmentPolicyId, ok = input.Parsed["accessPackageAssignmentPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentPolicyId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Id Assignment Policy ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Id Assignment Policy ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackages/%s/assignmentPolicies/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageId, id.AccessPackageAssignmentPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Id Assignment Policy ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
		resourceids.StaticSegment("assignmentPolicies", "assignmentPolicies", "assignmentPolicies"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentPolicyId", "accessPackageAssignmentPolicyId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Id Assignment Policy ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
		fmt.Sprintf("Access Package Assignment Policy: %q", id.AccessPackageAssignmentPolicyId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Id Assignment Policy (%s)", strings.Join(components, "\n"))
}
