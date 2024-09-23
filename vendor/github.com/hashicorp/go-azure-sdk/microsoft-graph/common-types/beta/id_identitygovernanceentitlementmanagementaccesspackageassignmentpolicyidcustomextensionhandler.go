package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{}

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Assignment Policy Id Custom Extension Handler
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId struct {
	AccessPackageAssignmentPolicyId string
	CustomExtensionHandlerId        string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerID returns a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerID(accessPackageAssignmentPolicyId string, customExtensionHandlerId string) IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId {
	return IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{
		AccessPackageAssignmentPolicyId: accessPackageAssignmentPolicyId,
		CustomExtensionHandlerId:        customExtensionHandlerId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentPolicyId, ok = input.Parsed["accessPackageAssignmentPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentPolicyId", input)
	}

	if id.CustomExtensionHandlerId, ok = input.Parsed["customExtensionHandlerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customExtensionHandlerId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Assignment Policy Id Custom Extension Handler ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Assignment Policy Id Custom Extension Handler ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageAssignmentPolicies/%s/customExtensionHandlers/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentPolicyId, id.CustomExtensionHandlerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Assignment Policy Id Custom Extension Handler ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageAssignmentPolicies", "accessPackageAssignmentPolicies", "accessPackageAssignmentPolicies"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentPolicyId", "accessPackageAssignmentPolicyId"),
		resourceids.StaticSegment("customExtensionHandlers", "customExtensionHandlers", "customExtensionHandlers"),
		resourceids.UserSpecifiedSegment("customExtensionHandlerId", "customExtensionHandlerId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Assignment Policy Id Custom Extension Handler ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIdCustomExtensionHandlerId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment Policy: %q", id.AccessPackageAssignmentPolicyId),
		fmt.Sprintf("Custom Extension Handler: %q", id.CustomExtensionHandlerId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Assignment Policy Id Custom Extension Handler (%s)", strings.Join(components, "\n"))
}
