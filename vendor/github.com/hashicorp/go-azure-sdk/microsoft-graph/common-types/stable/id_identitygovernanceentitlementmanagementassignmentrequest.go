package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAssignmentRequestId{}

// IdentityGovernanceEntitlementManagementAssignmentRequestId is a struct representing the Resource ID for a Identity Governance Entitlement Management Assignment Request
type IdentityGovernanceEntitlementManagementAssignmentRequestId struct {
	AccessPackageAssignmentRequestId string
}

// NewIdentityGovernanceEntitlementManagementAssignmentRequestID returns a new IdentityGovernanceEntitlementManagementAssignmentRequestId struct
func NewIdentityGovernanceEntitlementManagementAssignmentRequestID(accessPackageAssignmentRequestId string) IdentityGovernanceEntitlementManagementAssignmentRequestId {
	return IdentityGovernanceEntitlementManagementAssignmentRequestId{
		AccessPackageAssignmentRequestId: accessPackageAssignmentRequestId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAssignmentRequestID parses 'input' into a IdentityGovernanceEntitlementManagementAssignmentRequestId
func ParseIdentityGovernanceEntitlementManagementAssignmentRequestID(input string) (*IdentityGovernanceEntitlementManagementAssignmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAssignmentRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAssignmentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAssignmentRequestIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAssignmentRequestId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAssignmentRequestIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAssignmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAssignmentRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAssignmentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAssignmentRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentRequestId, ok = input.Parsed["accessPackageAssignmentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentRequestId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAssignmentRequestID checks that 'input' can be parsed as a Identity Governance Entitlement Management Assignment Request ID
func ValidateIdentityGovernanceEntitlementManagementAssignmentRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAssignmentRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Assignment Request ID
func (id IdentityGovernanceEntitlementManagementAssignmentRequestId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/assignmentRequests/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Assignment Request ID
func (id IdentityGovernanceEntitlementManagementAssignmentRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("assignmentRequests", "assignmentRequests", "assignmentRequests"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentRequestId", "accessPackageAssignmentRequestId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Assignment Request ID
func (id IdentityGovernanceEntitlementManagementAssignmentRequestId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment Request: %q", id.AccessPackageAssignmentRequestId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Assignment Request (%s)", strings.Join(components, "\n"))
}
