package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestId{}

// IdentityGovernanceEntitlementManagementResourceRequestId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Request
type IdentityGovernanceEntitlementManagementResourceRequestId struct {
	AccessPackageResourceRequestId string
}

// NewIdentityGovernanceEntitlementManagementResourceRequestID returns a new IdentityGovernanceEntitlementManagementResourceRequestId struct
func NewIdentityGovernanceEntitlementManagementResourceRequestID(accessPackageResourceRequestId string) IdentityGovernanceEntitlementManagementResourceRequestId {
	return IdentityGovernanceEntitlementManagementResourceRequestId{
		AccessPackageResourceRequestId: accessPackageResourceRequestId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRequestId
func ParseIdentityGovernanceEntitlementManagementResourceRequestID(input string) (*IdentityGovernanceEntitlementManagementResourceRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRequestIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRequestId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRequestIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRequestId, ok = input.Parsed["accessPackageResourceRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRequestId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRequestID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Request ID
func ValidateIdentityGovernanceEntitlementManagementResourceRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Request ID
func (id IdentityGovernanceEntitlementManagementResourceRequestId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRequests/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Request ID
func (id IdentityGovernanceEntitlementManagementResourceRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRequests", "resourceRequests", "resourceRequests"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRequestId", "accessPackageResourceRequestId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Request ID
func (id IdentityGovernanceEntitlementManagementResourceRequestId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Request: %q", id.AccessPackageResourceRequestId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Request (%s)", strings.Join(components, "\n"))
}
