package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceEnvironmentId{}

// IdentityGovernanceEntitlementManagementResourceEnvironmentId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Environment
type IdentityGovernanceEntitlementManagementResourceEnvironmentId struct {
	AccessPackageResourceEnvironmentId string
}

// NewIdentityGovernanceEntitlementManagementResourceEnvironmentID returns a new IdentityGovernanceEntitlementManagementResourceEnvironmentId struct
func NewIdentityGovernanceEntitlementManagementResourceEnvironmentID(accessPackageResourceEnvironmentId string) IdentityGovernanceEntitlementManagementResourceEnvironmentId {
	return IdentityGovernanceEntitlementManagementResourceEnvironmentId{
		AccessPackageResourceEnvironmentId: accessPackageResourceEnvironmentId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceEnvironmentID parses 'input' into a IdentityGovernanceEntitlementManagementResourceEnvironmentId
func ParseIdentityGovernanceEntitlementManagementResourceEnvironmentID(input string) (*IdentityGovernanceEntitlementManagementResourceEnvironmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceEnvironmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceEnvironmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceEnvironmentId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceEnvironmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceEnvironmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceEnvironmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceEnvironmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceEnvironmentId, ok = input.Parsed["accessPackageResourceEnvironmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceEnvironmentId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceEnvironmentID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Environment ID
func ValidateIdentityGovernanceEntitlementManagementResourceEnvironmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceEnvironmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Environment ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceEnvironments/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceEnvironmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Environment ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceEnvironments", "resourceEnvironments", "resourceEnvironments"),
		resourceids.UserSpecifiedSegment("accessPackageResourceEnvironmentId", "accessPackageResourceEnvironmentId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Environment ID
func (id IdentityGovernanceEntitlementManagementResourceEnvironmentId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Environment: %q", id.AccessPackageResourceEnvironmentId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Environment (%s)", strings.Join(components, "\n"))
}
