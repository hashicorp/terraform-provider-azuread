package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceId{}

// IdentityGovernanceEntitlementManagementResourceId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource
type IdentityGovernanceEntitlementManagementResourceId struct {
	AccessPackageResourceId string
}

// NewIdentityGovernanceEntitlementManagementResourceID returns a new IdentityGovernanceEntitlementManagementResourceId struct
func NewIdentityGovernanceEntitlementManagementResourceID(accessPackageResourceId string) IdentityGovernanceEntitlementManagementResourceId {
	return IdentityGovernanceEntitlementManagementResourceId{
		AccessPackageResourceId: accessPackageResourceId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceID parses 'input' into a IdentityGovernanceEntitlementManagementResourceId
func ParseIdentityGovernanceEntitlementManagementResourceID(input string) (*IdentityGovernanceEntitlementManagementResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource ID
func ValidateIdentityGovernanceEntitlementManagementResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource ID
func (id IdentityGovernanceEntitlementManagementResourceId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resources/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource ID
func (id IdentityGovernanceEntitlementManagementResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource ID
func (id IdentityGovernanceEntitlementManagementResourceId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource (%s)", strings.Join(components, "\n"))
}
