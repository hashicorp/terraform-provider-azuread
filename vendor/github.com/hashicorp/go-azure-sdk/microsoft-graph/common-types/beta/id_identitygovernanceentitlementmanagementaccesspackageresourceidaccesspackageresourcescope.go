package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId{}

// IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Resource Id Access Package Resource Scope
type IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId struct {
	AccessPackageResourceId      string
	AccessPackageResourceScopeId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeID returns a new IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeID(accessPackageResourceId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId {
	return IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId{
		AccessPackageResourceId:      accessPackageResourceId,
		AccessPackageResourceScopeId: accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId
func ParseIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Resource Id Access Package Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Resource Id Access Package Resource Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageResources/%s/accessPackageResourceScopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Resource Id Access Package Resource Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageResources", "accessPackageResources", "accessPackageResources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("accessPackageResourceScopes", "accessPackageResourceScopes", "accessPackageResourceScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Resource Id Access Package Resource Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Resource Id Access Package Resource Scope (%s)", strings.Join(components, "\n"))
}
