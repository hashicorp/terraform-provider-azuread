package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{}

// IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Id Access Package Resource Scope Id Access Package Resource Access Package Resource Role
type IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId struct {
	AccessPackageCatalogId       string
	AccessPackageResourceId      string
	AccessPackageResourceScopeId string
	AccessPackageResourceRoleId  string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID returns a new IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(accessPackageCatalogId string, accessPackageResourceId string, accessPackageResourceScopeId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId {
	return IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{
		AccessPackageCatalogId:       accessPackageCatalogId,
		AccessPackageResourceId:      accessPackageResourceId,
		AccessPackageResourceScopeId: accessPackageResourceScopeId,
		AccessPackageResourceRoleId:  accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId
func ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Id Access Package Resource Scope Id Access Package Resource Access Package Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Id Access Package Resource Scope Id Access Package Resource Access Package Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageCatalogs/%s/accessPackageResources/%s/accessPackageResourceScopes/%s/accessPackageResource/accessPackageResourceRoles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.AccessPackageResourceId, id.AccessPackageResourceScopeId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Id Access Package Resource Scope Id Access Package Resource Access Package Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageCatalogs", "accessPackageCatalogs", "accessPackageCatalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("accessPackageResources", "accessPackageResources", "accessPackageResources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("accessPackageResourceScopes", "accessPackageResourceScopes", "accessPackageResourceScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
		resourceids.StaticSegment("accessPackageResource", "accessPackageResource", "accessPackageResource"),
		resourceids.StaticSegment("accessPackageResourceRoles", "accessPackageResourceRoles", "accessPackageResourceRoles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Id Access Package Resource Scope Id Access Package Resource Access Package Resource Role ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Id Access Package Resource Scope Id Access Package Resource Access Package Resource Role (%s)", strings.Join(components, "\n"))
}
