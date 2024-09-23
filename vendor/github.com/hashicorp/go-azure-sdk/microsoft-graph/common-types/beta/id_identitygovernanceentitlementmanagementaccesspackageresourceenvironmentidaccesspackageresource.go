package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId{}

// IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Resource Environment Id Access Package Resource
type IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId struct {
	AccessPackageResourceEnvironmentId string
	AccessPackageResourceId            string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceID returns a new IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceID(accessPackageResourceEnvironmentId string, accessPackageResourceId string) IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId {
	return IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId{
		AccessPackageResourceEnvironmentId: accessPackageResourceEnvironmentId,
		AccessPackageResourceId:            accessPackageResourceId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId
func ParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceEnvironmentId, ok = input.Parsed["accessPackageResourceEnvironmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceEnvironmentId", input)
	}

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Resource Environment Id Access Package Resource ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Resource Environment Id Access Package Resource ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/%s/accessPackageResources/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceEnvironmentId, id.AccessPackageResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Resource Environment Id Access Package Resource ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageResourceEnvironments", "accessPackageResourceEnvironments", "accessPackageResourceEnvironments"),
		resourceids.UserSpecifiedSegment("accessPackageResourceEnvironmentId", "accessPackageResourceEnvironmentId"),
		resourceids.StaticSegment("accessPackageResources", "accessPackageResources", "accessPackageResources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Resource Environment Id Access Package Resource ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Environment: %q", id.AccessPackageResourceEnvironmentId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Resource Environment Id Access Package Resource (%s)", strings.Join(components, "\n"))
}
