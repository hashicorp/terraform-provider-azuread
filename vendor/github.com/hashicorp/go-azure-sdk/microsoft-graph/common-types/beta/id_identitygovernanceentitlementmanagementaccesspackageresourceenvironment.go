package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId{}

// IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Resource Environment
type IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId struct {
	AccessPackageResourceEnvironmentId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentID returns a new IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentID(accessPackageResourceEnvironmentId string) IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId {
	return IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId{
		AccessPackageResourceEnvironmentId: accessPackageResourceEnvironmentId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId
func ParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceEnvironmentId, ok = input.Parsed["accessPackageResourceEnvironmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceEnvironmentId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Resource Environment ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Resource Environment ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceEnvironmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Resource Environment ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageResourceEnvironments", "accessPackageResourceEnvironments", "accessPackageResourceEnvironments"),
		resourceids.UserSpecifiedSegment("accessPackageResourceEnvironmentId", "accessPackageResourceEnvironmentId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Resource Environment ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Environment: %q", id.AccessPackageResourceEnvironmentId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Resource Environment (%s)", strings.Join(components, "\n"))
}
