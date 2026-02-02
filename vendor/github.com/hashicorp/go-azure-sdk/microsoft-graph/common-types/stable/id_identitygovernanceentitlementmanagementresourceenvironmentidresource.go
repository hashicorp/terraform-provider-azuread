package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId{}

// IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Environment Id Resource
type IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId struct {
	AccessPackageResourceEnvironmentId string
	AccessPackageResourceId            string
}

// NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceID returns a new IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId struct
func NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceID(accessPackageResourceEnvironmentId string, accessPackageResourceId string) IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId {
	return IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId{
		AccessPackageResourceEnvironmentId: accessPackageResourceEnvironmentId,
		AccessPackageResourceId:            accessPackageResourceId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceID parses 'input' into a IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId
func ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceID(input string) (*IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceEnvironmentId, ok = input.Parsed["accessPackageResourceEnvironmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceEnvironmentId", input)
	}

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Environment Id Resource ID
func ValidateIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Environment Id Resource ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceEnvironments/%s/resources/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceEnvironmentId, id.AccessPackageResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Environment Id Resource ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceEnvironments", "resourceEnvironments", "resourceEnvironments"),
		resourceids.UserSpecifiedSegment("accessPackageResourceEnvironmentId", "accessPackageResourceEnvironmentId"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Environment Id Resource ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Environment: %q", id.AccessPackageResourceEnvironmentId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Environment Id Resource (%s)", strings.Join(components, "\n"))
}
