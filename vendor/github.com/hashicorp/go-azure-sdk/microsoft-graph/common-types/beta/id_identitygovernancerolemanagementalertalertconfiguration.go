package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceRoleManagementAlertAlertConfigurationId{}

// IdentityGovernanceRoleManagementAlertAlertConfigurationId is a struct representing the Resource ID for a Identity Governance Role Management Alert Alert Configuration
type IdentityGovernanceRoleManagementAlertAlertConfigurationId struct {
	UnifiedRoleManagementAlertConfigurationId string
}

// NewIdentityGovernanceRoleManagementAlertAlertConfigurationID returns a new IdentityGovernanceRoleManagementAlertAlertConfigurationId struct
func NewIdentityGovernanceRoleManagementAlertAlertConfigurationID(unifiedRoleManagementAlertConfigurationId string) IdentityGovernanceRoleManagementAlertAlertConfigurationId {
	return IdentityGovernanceRoleManagementAlertAlertConfigurationId{
		UnifiedRoleManagementAlertConfigurationId: unifiedRoleManagementAlertConfigurationId,
	}
}

// ParseIdentityGovernanceRoleManagementAlertAlertConfigurationID parses 'input' into a IdentityGovernanceRoleManagementAlertAlertConfigurationId
func ParseIdentityGovernanceRoleManagementAlertAlertConfigurationID(input string) (*IdentityGovernanceRoleManagementAlertAlertConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceRoleManagementAlertAlertConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceRoleManagementAlertAlertConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceRoleManagementAlertAlertConfigurationIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceRoleManagementAlertAlertConfigurationId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceRoleManagementAlertAlertConfigurationIDInsensitively(input string) (*IdentityGovernanceRoleManagementAlertAlertConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceRoleManagementAlertAlertConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceRoleManagementAlertAlertConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceRoleManagementAlertAlertConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleManagementAlertConfigurationId, ok = input.Parsed["unifiedRoleManagementAlertConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementAlertConfigurationId", input)
	}

	return nil
}

// ValidateIdentityGovernanceRoleManagementAlertAlertConfigurationID checks that 'input' can be parsed as a Identity Governance Role Management Alert Alert Configuration ID
func ValidateIdentityGovernanceRoleManagementAlertAlertConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceRoleManagementAlertAlertConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Role Management Alert Alert Configuration ID
func (id IdentityGovernanceRoleManagementAlertAlertConfigurationId) ID() string {
	fmtString := "/identityGovernance/roleManagementAlerts/alertConfigurations/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleManagementAlertConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Role Management Alert Alert Configuration ID
func (id IdentityGovernanceRoleManagementAlertAlertConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("roleManagementAlerts", "roleManagementAlerts", "roleManagementAlerts"),
		resourceids.StaticSegment("alertConfigurations", "alertConfigurations", "alertConfigurations"),
		resourceids.UserSpecifiedSegment("unifiedRoleManagementAlertConfigurationId", "unifiedRoleManagementAlertConfigurationId"),
	}
}

// String returns a human-readable description of this Identity Governance Role Management Alert Alert Configuration ID
func (id IdentityGovernanceRoleManagementAlertAlertConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Management Alert Configuration: %q", id.UnifiedRoleManagementAlertConfigurationId),
	}
	return fmt.Sprintf("Identity Governance Role Management Alert Alert Configuration (%s)", strings.Join(components, "\n"))
}
