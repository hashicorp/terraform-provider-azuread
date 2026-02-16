package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId{}

// IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Environment Id Resource Id Scope
type IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId struct {
	AccessPackageResourceEnvironmentId string
	AccessPackageResourceId            string
	AccessPackageResourceScopeId       string
}

// NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeID returns a new IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId struct
func NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeID(accessPackageResourceEnvironmentId string, accessPackageResourceId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId {
	return IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId{
		AccessPackageResourceEnvironmentId: accessPackageResourceEnvironmentId,
		AccessPackageResourceId:            accessPackageResourceId,
		AccessPackageResourceScopeId:       accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeID parses 'input' into a IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId
func ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeID(input string) (*IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceEnvironmentId, ok = input.Parsed["accessPackageResourceEnvironmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceEnvironmentId", input)
	}

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Environment Id Resource Id Scope ID
func ValidateIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Environment Id Resource Id Scope ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceEnvironments/%s/resources/%s/scopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceEnvironmentId, id.AccessPackageResourceId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Environment Id Resource Id Scope ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceEnvironments", "resourceEnvironments", "resourceEnvironments"),
		resourceids.UserSpecifiedSegment("accessPackageResourceEnvironmentId", "accessPackageResourceEnvironmentId"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Environment Id Resource Id Scope ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Environment: %q", id.AccessPackageResourceEnvironmentId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Environment Id Resource Id Scope (%s)", strings.Join(components, "\n"))
}
