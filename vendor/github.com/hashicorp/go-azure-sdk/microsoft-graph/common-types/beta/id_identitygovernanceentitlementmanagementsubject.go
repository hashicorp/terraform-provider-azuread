package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementSubjectId{}

// IdentityGovernanceEntitlementManagementSubjectId is a struct representing the Resource ID for a Identity Governance Entitlement Management Subject
type IdentityGovernanceEntitlementManagementSubjectId struct {
	AccessPackageSubjectId string
}

// NewIdentityGovernanceEntitlementManagementSubjectID returns a new IdentityGovernanceEntitlementManagementSubjectId struct
func NewIdentityGovernanceEntitlementManagementSubjectID(accessPackageSubjectId string) IdentityGovernanceEntitlementManagementSubjectId {
	return IdentityGovernanceEntitlementManagementSubjectId{
		AccessPackageSubjectId: accessPackageSubjectId,
	}
}

// ParseIdentityGovernanceEntitlementManagementSubjectID parses 'input' into a IdentityGovernanceEntitlementManagementSubjectId
func ParseIdentityGovernanceEntitlementManagementSubjectID(input string) (*IdentityGovernanceEntitlementManagementSubjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementSubjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementSubjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementSubjectIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementSubjectId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementSubjectIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementSubjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementSubjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementSubjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementSubjectId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageSubjectId, ok = input.Parsed["accessPackageSubjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageSubjectId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementSubjectID checks that 'input' can be parsed as a Identity Governance Entitlement Management Subject ID
func ValidateIdentityGovernanceEntitlementManagementSubjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementSubjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Subject ID
func (id IdentityGovernanceEntitlementManagementSubjectId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/subjects/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageSubjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Subject ID
func (id IdentityGovernanceEntitlementManagementSubjectId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("subjects", "subjects", "subjects"),
		resourceids.UserSpecifiedSegment("accessPackageSubjectId", "accessPackageSubjectId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Subject ID
func (id IdentityGovernanceEntitlementManagementSubjectId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Subject: %q", id.AccessPackageSubjectId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Subject (%s)", strings.Join(components, "\n"))
}
