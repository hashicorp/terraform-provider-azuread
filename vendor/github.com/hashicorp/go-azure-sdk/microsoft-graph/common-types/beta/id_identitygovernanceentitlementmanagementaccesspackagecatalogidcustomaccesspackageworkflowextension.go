package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId{}

// IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Catalog Id Custom Access Package Workflow Extension
type IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId struct {
	AccessPackageCatalogId                 string
	CustomAccessPackageWorkflowExtensionId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID returns a new IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID(accessPackageCatalogId string, customAccessPackageWorkflowExtensionId string) IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId {
	return IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId{
		AccessPackageCatalogId:                 accessPackageCatalogId,
		CustomAccessPackageWorkflowExtensionId: customAccessPackageWorkflowExtensionId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId
func ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.CustomAccessPackageWorkflowExtensionId, ok = input.Parsed["customAccessPackageWorkflowExtensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customAccessPackageWorkflowExtensionId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Catalog Id Custom Access Package Workflow Extension ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Catalog Id Custom Access Package Workflow Extension ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageCatalogs/%s/customAccessPackageWorkflowExtensions/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.CustomAccessPackageWorkflowExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Catalog Id Custom Access Package Workflow Extension ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageCatalogs", "accessPackageCatalogs", "accessPackageCatalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("customAccessPackageWorkflowExtensions", "customAccessPackageWorkflowExtensions", "customAccessPackageWorkflowExtensions"),
		resourceids.UserSpecifiedSegment("customAccessPackageWorkflowExtensionId", "customAccessPackageWorkflowExtensionId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Catalog Id Custom Access Package Workflow Extension ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Custom Access Package Workflow Extension: %q", id.CustomAccessPackageWorkflowExtensionId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Catalog Id Custom Access Package Workflow Extension (%s)", strings.Join(components, "\n"))
}
