package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceIdScopeId{}

// IdentityGovernanceEntitlementManagementResourceIdScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Id Scope
type IdentityGovernanceEntitlementManagementResourceIdScopeId struct {
	AccessPackageResourceId      string
	AccessPackageResourceScopeId string
}

// NewIdentityGovernanceEntitlementManagementResourceIdScopeID returns a new IdentityGovernanceEntitlementManagementResourceIdScopeId struct
func NewIdentityGovernanceEntitlementManagementResourceIdScopeID(accessPackageResourceId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementResourceIdScopeId {
	return IdentityGovernanceEntitlementManagementResourceIdScopeId{
		AccessPackageResourceId:      accessPackageResourceId,
		AccessPackageResourceScopeId: accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceIdScopeID parses 'input' into a IdentityGovernanceEntitlementManagementResourceIdScopeId
func ParseIdentityGovernanceEntitlementManagementResourceIdScopeID(input string) (*IdentityGovernanceEntitlementManagementResourceIdScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceIdScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceIdScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceIdScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceIdScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceIdScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceIdScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceIdScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceIdScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceIdScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceIdScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Id Scope ID
func ValidateIdentityGovernanceEntitlementManagementResourceIdScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceIdScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Id Scope ID
func (id IdentityGovernanceEntitlementManagementResourceIdScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resources/%s/scopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Id Scope ID
func (id IdentityGovernanceEntitlementManagementResourceIdScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Id Scope ID
func (id IdentityGovernanceEntitlementManagementResourceIdScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Id Scope (%s)", strings.Join(components, "\n"))
}
