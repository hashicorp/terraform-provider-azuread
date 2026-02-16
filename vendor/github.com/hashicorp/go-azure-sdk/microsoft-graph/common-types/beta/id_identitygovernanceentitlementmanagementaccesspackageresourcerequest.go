package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId{}

// IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Resource Request
type IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId struct {
	AccessPackageResourceRequestId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageResourceRequestID returns a new IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageResourceRequestID(accessPackageResourceRequestId string) IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId {
	return IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRequestID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId
func ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRequestID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRequestIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRequestIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRequestId, ok = input.Parsed["accessPackageResourceRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRequestId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageResourceRequestID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Resource Request ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageResourceRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Resource Request ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageResourceRequests/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Resource Request ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageResourceRequests", "accessPackageResourceRequests", "accessPackageResourceRequests"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRequestId", "accessPackageResourceRequestId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Resource Request ID
func (id IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Resource Request (%s)", strings.Join(components, "\n"))
}
