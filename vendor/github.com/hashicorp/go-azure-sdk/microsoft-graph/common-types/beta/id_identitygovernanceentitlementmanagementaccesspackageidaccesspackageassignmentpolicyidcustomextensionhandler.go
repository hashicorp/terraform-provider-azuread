package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{}

// IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy Id Custom Extension Handler
type IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId struct {
	AccessPackageId                 string
	AccessPackageAssignmentPolicyId string
	CustomExtensionHandlerId        string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerID returns a new IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerID(accessPackageId string, accessPackageAssignmentPolicyId string, customExtensionHandlerId string) IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId {
	return IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{
		AccessPackageId:                 accessPackageId,
		AccessPackageAssignmentPolicyId: accessPackageAssignmentPolicyId,
		CustomExtensionHandlerId:        customExtensionHandlerId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	if id.AccessPackageAssignmentPolicyId, ok = input.Parsed["accessPackageAssignmentPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentPolicyId", input)
	}

	if id.CustomExtensionHandlerId, ok = input.Parsed["customExtensionHandlerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customExtensionHandlerId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy Id Custom Extension Handler ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy Id Custom Extension Handler ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackages/%s/accessPackageAssignmentPolicies/%s/customExtensionHandlers/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageId, id.AccessPackageAssignmentPolicyId, id.CustomExtensionHandlerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy Id Custom Extension Handler ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
		resourceids.StaticSegment("accessPackageAssignmentPolicies", "accessPackageAssignmentPolicies", "accessPackageAssignmentPolicies"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentPolicyId", "accessPackageAssignmentPolicyId"),
		resourceids.StaticSegment("customExtensionHandlers", "customExtensionHandlers", "customExtensionHandlers"),
		resourceids.UserSpecifiedSegment("customExtensionHandlerId", "customExtensionHandlerId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy Id Custom Extension Handler ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyIdCustomExtensionHandlerId) String() string {
	components := []string{
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
		fmt.Sprintf("Access Package Assignment Policy: %q", id.AccessPackageAssignmentPolicyId),
		fmt.Sprintf("Custom Extension Handler: %q", id.CustomExtensionHandlerId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Id Access Package Assignment Policy Id Custom Extension Handler (%s)", strings.Join(components, "\n"))
}
