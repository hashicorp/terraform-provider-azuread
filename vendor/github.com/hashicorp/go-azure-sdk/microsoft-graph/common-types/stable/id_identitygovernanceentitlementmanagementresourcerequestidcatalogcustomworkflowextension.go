package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId{}

// IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Request Id Catalog Custom Workflow Extension
type IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId struct {
	AccessPackageResourceRequestId string
	CustomCalloutExtensionId       string
}

// NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionID returns a new IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId struct
func NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionID(accessPackageResourceRequestId string, customCalloutExtensionId string) IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId {
	return IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
		CustomCalloutExtensionId:       customCalloutExtensionId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionID(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRequestId, ok = input.Parsed["accessPackageResourceRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRequestId", input)
	}

	if id.CustomCalloutExtensionId, ok = input.Parsed["customCalloutExtensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customCalloutExtensionId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Request Id Catalog Custom Workflow Extension ID
func ValidateIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Request Id Catalog Custom Workflow Extension ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRequests/%s/catalog/customWorkflowExtensions/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId, id.CustomCalloutExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Request Id Catalog Custom Workflow Extension ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRequests", "resourceRequests", "resourceRequests"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRequestId", "accessPackageResourceRequestId"),
		resourceids.StaticSegment("catalog", "catalog", "catalog"),
		resourceids.StaticSegment("customWorkflowExtensions", "customWorkflowExtensions", "customWorkflowExtensions"),
		resourceids.UserSpecifiedSegment("customCalloutExtensionId", "customCalloutExtensionId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Request Id Catalog Custom Workflow Extension ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
		fmt.Sprintf("Custom Callout Extension: %q", id.CustomCalloutExtensionId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Request Id Catalog Custom Workflow Extension (%s)", strings.Join(components, "\n"))
}
