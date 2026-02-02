package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementExchangeRoleAssignmentId{}

// RoleManagementExchangeRoleAssignmentId is a struct representing the Resource ID for a Role Management Exchange Role Assignment
type RoleManagementExchangeRoleAssignmentId struct {
	UnifiedRoleAssignmentId string
}

// NewRoleManagementExchangeRoleAssignmentID returns a new RoleManagementExchangeRoleAssignmentId struct
func NewRoleManagementExchangeRoleAssignmentID(unifiedRoleAssignmentId string) RoleManagementExchangeRoleAssignmentId {
	return RoleManagementExchangeRoleAssignmentId{
		UnifiedRoleAssignmentId: unifiedRoleAssignmentId,
	}
}

// ParseRoleManagementExchangeRoleAssignmentID parses 'input' into a RoleManagementExchangeRoleAssignmentId
func ParseRoleManagementExchangeRoleAssignmentID(input string) (*RoleManagementExchangeRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementExchangeRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementExchangeRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementExchangeRoleAssignmentIDInsensitively parses 'input' case-insensitively into a RoleManagementExchangeRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementExchangeRoleAssignmentIDInsensitively(input string) (*RoleManagementExchangeRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementExchangeRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementExchangeRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementExchangeRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentId, ok = input.Parsed["unifiedRoleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentId", input)
	}

	return nil
}

// ValidateRoleManagementExchangeRoleAssignmentID checks that 'input' can be parsed as a Role Management Exchange Role Assignment ID
func ValidateRoleManagementExchangeRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementExchangeRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Exchange Role Assignment ID
func (id RoleManagementExchangeRoleAssignmentId) ID() string {
	fmtString := "/roleManagement/exchange/roleAssignments/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Exchange Role Assignment ID
func (id RoleManagementExchangeRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("exchange", "exchange", "exchange"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentId", "unifiedRoleAssignmentId"),
	}
}

// String returns a human-readable description of this Role Management Exchange Role Assignment ID
func (id RoleManagementExchangeRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment: %q", id.UnifiedRoleAssignmentId),
	}
	return fmt.Sprintf("Role Management Exchange Role Assignment (%s)", strings.Join(components, "\n"))
}
