package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageResourceId{}

// IdentityGovernanceEntitlementManagementAccessPackageResourceId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Resource
type IdentityGovernanceEntitlementManagementAccessPackageResourceId struct {
	AccessPackageResourceId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageResourceID returns a new IdentityGovernanceEntitlementManagementAccessPackageResourceId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageResourceID(accessPackageResourceId string) IdentityGovernanceEntitlementManagementAccessPackageResourceId {
	return IdentityGovernanceEntitlementManagementAccessPackageResourceId{
		AccessPackageResourceId: accessPackageResourceId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageResourceID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageResourceId
func ParseIdentityGovernanceEntitlementManagementAccessPackageResourceID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageResourceIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageResourceId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageResourceIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageResourceID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Resource ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Resource ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageResources/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Resource ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageResources", "accessPackageResources", "accessPackageResources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Resource ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Resource (%s)", strings.Join(components, "\n"))
}
