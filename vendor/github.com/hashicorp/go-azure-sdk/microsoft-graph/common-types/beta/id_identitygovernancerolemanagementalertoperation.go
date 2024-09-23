package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceRoleManagementAlertOperationId{}

// IdentityGovernanceRoleManagementAlertOperationId is a struct representing the Resource ID for a Identity Governance Role Management Alert Operation
type IdentityGovernanceRoleManagementAlertOperationId struct {
	LongRunningOperationId string
}

// NewIdentityGovernanceRoleManagementAlertOperationID returns a new IdentityGovernanceRoleManagementAlertOperationId struct
func NewIdentityGovernanceRoleManagementAlertOperationID(longRunningOperationId string) IdentityGovernanceRoleManagementAlertOperationId {
	return IdentityGovernanceRoleManagementAlertOperationId{
		LongRunningOperationId: longRunningOperationId,
	}
}

// ParseIdentityGovernanceRoleManagementAlertOperationID parses 'input' into a IdentityGovernanceRoleManagementAlertOperationId
func ParseIdentityGovernanceRoleManagementAlertOperationID(input string) (*IdentityGovernanceRoleManagementAlertOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceRoleManagementAlertOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceRoleManagementAlertOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceRoleManagementAlertOperationIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceRoleManagementAlertOperationId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceRoleManagementAlertOperationIDInsensitively(input string) (*IdentityGovernanceRoleManagementAlertOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceRoleManagementAlertOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceRoleManagementAlertOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceRoleManagementAlertOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.LongRunningOperationId, ok = input.Parsed["longRunningOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "longRunningOperationId", input)
	}

	return nil
}

// ValidateIdentityGovernanceRoleManagementAlertOperationID checks that 'input' can be parsed as a Identity Governance Role Management Alert Operation ID
func ValidateIdentityGovernanceRoleManagementAlertOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceRoleManagementAlertOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Role Management Alert Operation ID
func (id IdentityGovernanceRoleManagementAlertOperationId) ID() string {
	fmtString := "/identityGovernance/roleManagementAlerts/operations/%s"
	return fmt.Sprintf(fmtString, id.LongRunningOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Role Management Alert Operation ID
func (id IdentityGovernanceRoleManagementAlertOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("roleManagementAlerts", "roleManagementAlerts", "roleManagementAlerts"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("longRunningOperationId", "longRunningOperationId"),
	}
}

// String returns a human-readable description of this Identity Governance Role Management Alert Operation ID
func (id IdentityGovernanceRoleManagementAlertOperationId) String() string {
	components := []string{
		fmt.Sprintf("Long Running Operation: %q", id.LongRunningOperationId),
	}
	return fmt.Sprintf("Identity Governance Role Management Alert Operation (%s)", strings.Join(components, "\n"))
}
