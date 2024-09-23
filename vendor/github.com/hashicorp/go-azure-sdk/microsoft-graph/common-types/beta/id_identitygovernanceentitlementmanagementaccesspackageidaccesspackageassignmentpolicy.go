package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId{}

// IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy
type IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId struct {
	AccessPackageId                 string
	AccessPackageAssignmentPolicyId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyID returns a new IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyID(accessPackageId string, accessPackageAssignmentPolicyId string) IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId {
	return IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId{
		AccessPackageId:                 accessPackageId,
		AccessPackageAssignmentPolicyId: accessPackageAssignmentPolicyId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	if id.AccessPackageAssignmentPolicyId, ok = input.Parsed["accessPackageAssignmentPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentPolicyId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackages/%s/accessPackageAssignmentPolicies/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageId, id.AccessPackageAssignmentPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
		resourceids.StaticSegment("accessPackageAssignmentPolicies", "accessPackageAssignmentPolicies", "accessPackageAssignmentPolicies"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentPolicyId", "accessPackageAssignmentPolicyId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
		fmt.Sprintf("Access Package Assignment Policy: %q", id.AccessPackageAssignmentPolicyId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy (%s)", strings.Join(components, "\n"))
}
