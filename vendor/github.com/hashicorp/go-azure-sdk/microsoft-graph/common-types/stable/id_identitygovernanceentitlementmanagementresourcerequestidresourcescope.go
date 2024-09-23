package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId{}

// IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Request Id Resource Scope
type IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId struct {
	AccessPackageResourceRequestId string
	AccessPackageResourceScopeId   string
}

// NewIdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeID returns a new IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeID(accessPackageResourceRequestId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId {
	return IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
		AccessPackageResourceScopeId:   accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRequestId, ok = input.Parsed["accessPackageResourceRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRequestId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Request Id Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Request Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRequests/%s/resource/scopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Request Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRequests", "resourceRequests", "resourceRequests"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRequestId", "accessPackageResourceRequestId"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("scopes", "scopes", "scopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Request Id Resource Scope ID
func (id IdentityGovernanceEntitlementManagementResourceRequestIdResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Request Id Resource Scope (%s)", strings.Join(components, "\n"))
}
