package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId{}

// IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Role Id Resource Scope
type IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId struct {
	AccessPackageResourceRequestId string
	AccessPackageResourceId        string
	AccessPackageResourceRoleId    string
	AccessPackageResourceScopeId   string
}

// NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeID returns a new IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeID(accessPackageResourceRequestId string, accessPackageResourceId string, accessPackageResourceRoleId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId {
	return IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
		AccessPackageResourceId:        accessPackageResourceId,
		AccessPackageResourceRoleId:    accessPackageResourceRoleId,
		AccessPackageResourceScopeId:   accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRequestId, ok = input.Parsed["accessPackageResourceRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRequestId", input)
	}

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Role Id Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Role Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRequests/%s/catalog/resources/%s/roles/%s/resource/scopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId, id.AccessPackageResourceId, id.AccessPackageResourceRoleId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Role Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRequests", "resourceRequests", "resourceRequests"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRequestId", "accessPackageResourceRequestId"),
		resourceids.StaticSegment("catalog", "catalog", "catalog"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Role Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Request Id Catalog Resource Id Role Id Resource Scope (%s)", strings.Join(components, "\n"))
}
