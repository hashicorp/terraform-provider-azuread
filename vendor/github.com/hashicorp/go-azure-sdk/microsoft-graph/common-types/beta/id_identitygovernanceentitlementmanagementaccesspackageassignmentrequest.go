package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId{}

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Assignment Request
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId struct {
	AccessPackageAssignmentRequestId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID returns a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID(accessPackageAssignmentRequestId string) IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId {
	return IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId{
		AccessPackageAssignmentRequestId: accessPackageAssignmentRequestId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentRequestId, ok = input.Parsed["accessPackageAssignmentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentRequestId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Assignment Request ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Assignment Request ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageAssignmentRequests/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Assignment Request ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageAssignmentRequests", "accessPackageAssignmentRequests", "accessPackageAssignmentRequests"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentRequestId", "accessPackageAssignmentRequestId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Assignment Request ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment Request: %q", id.AccessPackageAssignmentRequestId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Assignment Request (%s)", strings.Join(components, "\n"))
}
