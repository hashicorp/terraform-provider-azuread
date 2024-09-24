package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementExchangeTransitiveRoleAssignmentId{}

// RoleManagementExchangeTransitiveRoleAssignmentId is a struct representing the Resource ID for a Role Management Exchange Transitive Role Assignment
type RoleManagementExchangeTransitiveRoleAssignmentId struct {
	UnifiedRoleAssignmentId string
}

// NewRoleManagementExchangeTransitiveRoleAssignmentID returns a new RoleManagementExchangeTransitiveRoleAssignmentId struct
func NewRoleManagementExchangeTransitiveRoleAssignmentID(unifiedRoleAssignmentId string) RoleManagementExchangeTransitiveRoleAssignmentId {
	return RoleManagementExchangeTransitiveRoleAssignmentId{
		UnifiedRoleAssignmentId: unifiedRoleAssignmentId,
	}
}

// ParseRoleManagementExchangeTransitiveRoleAssignmentID parses 'input' into a RoleManagementExchangeTransitiveRoleAssignmentId
func ParseRoleManagementExchangeTransitiveRoleAssignmentID(input string) (*RoleManagementExchangeTransitiveRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementExchangeTransitiveRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementExchangeTransitiveRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementExchangeTransitiveRoleAssignmentIDInsensitively parses 'input' case-insensitively into a RoleManagementExchangeTransitiveRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementExchangeTransitiveRoleAssignmentIDInsensitively(input string) (*RoleManagementExchangeTransitiveRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementExchangeTransitiveRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementExchangeTransitiveRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementExchangeTransitiveRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentId, ok = input.Parsed["unifiedRoleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentId", input)
	}

	return nil
}

// ValidateRoleManagementExchangeTransitiveRoleAssignmentID checks that 'input' can be parsed as a Role Management Exchange Transitive Role Assignment ID
func ValidateRoleManagementExchangeTransitiveRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementExchangeTransitiveRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Exchange Transitive Role Assignment ID
func (id RoleManagementExchangeTransitiveRoleAssignmentId) ID() string {
	fmtString := "/roleManagement/exchange/transitiveRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Exchange Transitive Role Assignment ID
func (id RoleManagementExchangeTransitiveRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("exchange", "exchange", "exchange"),
		resourceids.StaticSegment("transitiveRoleAssignments", "transitiveRoleAssignments", "transitiveRoleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentId", "unifiedRoleAssignmentId"),
	}
}

// String returns a human-readable description of this Role Management Exchange Transitive Role Assignment ID
func (id RoleManagementExchangeTransitiveRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment: %q", id.UnifiedRoleAssignmentId),
	}
	return fmt.Sprintf("Role Management Exchange Transitive Role Assignment (%s)", strings.Join(components, "\n"))
}
