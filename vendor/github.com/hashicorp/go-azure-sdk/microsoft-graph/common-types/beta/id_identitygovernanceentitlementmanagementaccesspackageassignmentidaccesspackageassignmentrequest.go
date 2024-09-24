package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId{}

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Assignment Id Access Package Assignment Request
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId struct {
	AccessPackageAssignmentId        string
	AccessPackageAssignmentRequestId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestID returns a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestID(accessPackageAssignmentId string, accessPackageAssignmentRequestId string) IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId {
	return IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId{
		AccessPackageAssignmentId:        accessPackageAssignmentId,
		AccessPackageAssignmentRequestId: accessPackageAssignmentRequestId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentId, ok = input.Parsed["accessPackageAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentId", input)
	}

	if id.AccessPackageAssignmentRequestId, ok = input.Parsed["accessPackageAssignmentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentRequestId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Assignment Id Access Package Assignment Request ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Assignment Id Access Package Assignment Request ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageAssignments/%s/accessPackageAssignmentRequests/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentId, id.AccessPackageAssignmentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Assignment Id Access Package Assignment Request ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageAssignments", "accessPackageAssignments", "accessPackageAssignments"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentId", "accessPackageAssignmentId"),
		resourceids.StaticSegment("accessPackageAssignmentRequests", "accessPackageAssignmentRequests", "accessPackageAssignmentRequests"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentRequestId", "accessPackageAssignmentRequestId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Assignment Id Access Package Assignment Request ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment: %q", id.AccessPackageAssignmentId),
		fmt.Sprintf("Access Package Assignment Request: %q", id.AccessPackageAssignmentRequestId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Assignment Id Access Package Assignment Request (%s)", strings.Join(components, "\n"))
}
