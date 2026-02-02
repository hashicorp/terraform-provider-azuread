package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceRoleManagementAlertAlertId{}

// IdentityGovernanceRoleManagementAlertAlertId is a struct representing the Resource ID for a Identity Governance Role Management Alert Alert
type IdentityGovernanceRoleManagementAlertAlertId struct {
	UnifiedRoleManagementAlertId string
}

// NewIdentityGovernanceRoleManagementAlertAlertID returns a new IdentityGovernanceRoleManagementAlertAlertId struct
func NewIdentityGovernanceRoleManagementAlertAlertID(unifiedRoleManagementAlertId string) IdentityGovernanceRoleManagementAlertAlertId {
	return IdentityGovernanceRoleManagementAlertAlertId{
		UnifiedRoleManagementAlertId: unifiedRoleManagementAlertId,
	}
}

// ParseIdentityGovernanceRoleManagementAlertAlertID parses 'input' into a IdentityGovernanceRoleManagementAlertAlertId
func ParseIdentityGovernanceRoleManagementAlertAlertID(input string) (*IdentityGovernanceRoleManagementAlertAlertId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceRoleManagementAlertAlertId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceRoleManagementAlertAlertId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceRoleManagementAlertAlertIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceRoleManagementAlertAlertId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceRoleManagementAlertAlertIDInsensitively(input string) (*IdentityGovernanceRoleManagementAlertAlertId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceRoleManagementAlertAlertId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceRoleManagementAlertAlertId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceRoleManagementAlertAlertId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleManagementAlertId, ok = input.Parsed["unifiedRoleManagementAlertId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementAlertId", input)
	}

	return nil
}

// ValidateIdentityGovernanceRoleManagementAlertAlertID checks that 'input' can be parsed as a Identity Governance Role Management Alert Alert ID
func ValidateIdentityGovernanceRoleManagementAlertAlertID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceRoleManagementAlertAlertID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Role Management Alert Alert ID
func (id IdentityGovernanceRoleManagementAlertAlertId) ID() string {
	fmtString := "/identityGovernance/roleManagementAlerts/alerts/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleManagementAlertId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Role Management Alert Alert ID
func (id IdentityGovernanceRoleManagementAlertAlertId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("roleManagementAlerts", "roleManagementAlerts", "roleManagementAlerts"),
		resourceids.StaticSegment("alerts", "alerts", "alerts"),
		resourceids.UserSpecifiedSegment("unifiedRoleManagementAlertId", "unifiedRoleManagementAlertId"),
	}
}

// String returns a human-readable description of this Identity Governance Role Management Alert Alert ID
func (id IdentityGovernanceRoleManagementAlertAlertId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Management Alert: %q", id.UnifiedRoleManagementAlertId),
	}
	return fmt.Sprintf("Identity Governance Role Management Alert Alert (%s)", strings.Join(components, "\n"))
}
