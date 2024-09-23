package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId{}

// IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId is a struct representing the Resource ID for a Identity Governance Entitlement Management Catalog Id Custom Workflow Extension
type IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId struct {
	AccessPackageCatalogId   string
	CustomCalloutExtensionId string
}

// NewIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID returns a new IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId struct
func NewIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID(accessPackageCatalogId string, customCalloutExtensionId string) IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId {
	return IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId{
		AccessPackageCatalogId:   accessPackageCatalogId,
		CustomCalloutExtensionId: customCalloutExtensionId,
	}
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID parses 'input' into a IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId
func ParseIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID(input string) (*IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.CustomCalloutExtensionId, ok = input.Parsed["customCalloutExtensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customCalloutExtensionId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID checks that 'input' can be parsed as a Identity Governance Entitlement Management Catalog Id Custom Workflow Extension ID
func ValidateIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Catalog Id Custom Workflow Extension ID
func (id IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/catalogs/%s/customWorkflowExtensions/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.CustomCalloutExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Catalog Id Custom Workflow Extension ID
func (id IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("catalogs", "catalogs", "catalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("customWorkflowExtensions", "customWorkflowExtensions", "customWorkflowExtensions"),
		resourceids.UserSpecifiedSegment("customCalloutExtensionId", "customCalloutExtensionId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Catalog Id Custom Workflow Extension ID
func (id IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Custom Callout Extension: %q", id.CustomCalloutExtensionId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Catalog Id Custom Workflow Extension (%s)", strings.Join(components, "\n"))
}
