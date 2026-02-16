package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsManagementPermissionsRequestChangeId{}

// IdentityGovernancePermissionsManagementPermissionsRequestChangeId is a struct representing the Resource ID for a Identity Governance Permissions Management Permissions Request Change
type IdentityGovernancePermissionsManagementPermissionsRequestChangeId struct {
	PermissionsRequestChangeId string
}

// NewIdentityGovernancePermissionsManagementPermissionsRequestChangeID returns a new IdentityGovernancePermissionsManagementPermissionsRequestChangeId struct
func NewIdentityGovernancePermissionsManagementPermissionsRequestChangeID(permissionsRequestChangeId string) IdentityGovernancePermissionsManagementPermissionsRequestChangeId {
	return IdentityGovernancePermissionsManagementPermissionsRequestChangeId{
		PermissionsRequestChangeId: permissionsRequestChangeId,
	}
}

// ParseIdentityGovernancePermissionsManagementPermissionsRequestChangeID parses 'input' into a IdentityGovernancePermissionsManagementPermissionsRequestChangeId
func ParseIdentityGovernancePermissionsManagementPermissionsRequestChangeID(input string) (*IdentityGovernancePermissionsManagementPermissionsRequestChangeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsManagementPermissionsRequestChangeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsManagementPermissionsRequestChangeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePermissionsManagementPermissionsRequestChangeIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePermissionsManagementPermissionsRequestChangeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePermissionsManagementPermissionsRequestChangeIDInsensitively(input string) (*IdentityGovernancePermissionsManagementPermissionsRequestChangeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePermissionsManagementPermissionsRequestChangeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePermissionsManagementPermissionsRequestChangeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePermissionsManagementPermissionsRequestChangeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PermissionsRequestChangeId, ok = input.Parsed["permissionsRequestChangeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionsRequestChangeId", input)
	}

	return nil
}

// ValidateIdentityGovernancePermissionsManagementPermissionsRequestChangeID checks that 'input' can be parsed as a Identity Governance Permissions Management Permissions Request Change ID
func ValidateIdentityGovernancePermissionsManagementPermissionsRequestChangeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePermissionsManagementPermissionsRequestChangeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Permissions Management Permissions Request Change ID
func (id IdentityGovernancePermissionsManagementPermissionsRequestChangeId) ID() string {
	fmtString := "/identityGovernance/permissionsManagement/permissionsRequestChanges/%s"
	return fmt.Sprintf(fmtString, id.PermissionsRequestChangeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Permissions Management Permissions Request Change ID
func (id IdentityGovernancePermissionsManagementPermissionsRequestChangeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("permissionsManagement", "permissionsManagement", "permissionsManagement"),
		resourceids.StaticSegment("permissionsRequestChanges", "permissionsRequestChanges", "permissionsRequestChanges"),
		resourceids.UserSpecifiedSegment("permissionsRequestChangeId", "permissionsRequestChangeId"),
	}
}

// String returns a human-readable description of this Identity Governance Permissions Management Permissions Request Change ID
func (id IdentityGovernancePermissionsManagementPermissionsRequestChangeId) String() string {
	components := []string{
		fmt.Sprintf("Permissions Request Change: %q", id.PermissionsRequestChangeId),
	}
	return fmt.Sprintf("Identity Governance Permissions Management Permissions Request Change (%s)", strings.Join(components, "\n"))
}
