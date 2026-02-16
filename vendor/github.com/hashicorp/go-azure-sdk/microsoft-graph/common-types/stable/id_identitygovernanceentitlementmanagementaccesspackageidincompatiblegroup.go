package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId{}

// IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Id Incompatible Group
type IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId struct {
	AccessPackageId string
	GroupId         string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupID returns a new IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupID(accessPackageId string, groupId string) IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId {
	return IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId{
		AccessPackageId: accessPackageId,
		GroupId:         groupId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageId, ok = input.Parsed["accessPackageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageId", input)
	}

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Id Incompatible Group ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Id Incompatible Group ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackages/%s/incompatibleGroups/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageId, id.GroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Id Incompatible Group ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackages", "accessPackages", "accessPackages"),
		resourceids.UserSpecifiedSegment("accessPackageId", "accessPackageId"),
		resourceids.StaticSegment("incompatibleGroups", "incompatibleGroups", "incompatibleGroups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Id Incompatible Group ID
func (id IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleGroupId) String() string {
	components := []string{
		fmt.Sprintf("Access Package: %q", id.AccessPackageId),
		fmt.Sprintf("Group: %q", id.GroupId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Id Incompatible Group (%s)", strings.Join(components, "\n"))
}
