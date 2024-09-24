package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId{}

// IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Catalog Id Access Package Custom Workflow Extension
type IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId struct {
	AccessPackageCatalogId   string
	CustomCalloutExtensionId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionID returns a new IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionID(accessPackageCatalogId string, customCalloutExtensionId string) IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId {
	return IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId{
		AccessPackageCatalogId:   accessPackageCatalogId,
		CustomCalloutExtensionId: customCalloutExtensionId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId
func ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.CustomCalloutExtensionId, ok = input.Parsed["customCalloutExtensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customCalloutExtensionId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Catalog Id Access Package Custom Workflow Extension ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Catalog Id Access Package Custom Workflow Extension ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageCatalogs/%s/accessPackageCustomWorkflowExtensions/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.CustomCalloutExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Catalog Id Access Package Custom Workflow Extension ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageCatalogs", "accessPackageCatalogs", "accessPackageCatalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("accessPackageCustomWorkflowExtensions", "accessPackageCustomWorkflowExtensions", "accessPackageCustomWorkflowExtensions"),
		resourceids.UserSpecifiedSegment("customCalloutExtensionId", "customCalloutExtensionId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Catalog Id Access Package Custom Workflow Extension ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Custom Callout Extension: %q", id.CustomCalloutExtensionId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Catalog Id Access Package Custom Workflow Extension (%s)", strings.Join(components, "\n"))
}
