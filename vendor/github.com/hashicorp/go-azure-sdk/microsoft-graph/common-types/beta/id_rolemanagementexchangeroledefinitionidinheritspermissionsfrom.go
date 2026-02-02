package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId{}

// RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId is a struct representing the Resource ID for a Role Management Exchange Role Definition Id Inherits Permissions From
type RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId struct {
	UnifiedRoleDefinitionId  string
	UnifiedRoleDefinitionId1 string
}

// NewRoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromID returns a new RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId struct
func NewRoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromID(unifiedRoleDefinitionId string, unifiedRoleDefinitionId1 string) RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId {
	return RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId{
		UnifiedRoleDefinitionId:  unifiedRoleDefinitionId,
		UnifiedRoleDefinitionId1: unifiedRoleDefinitionId1,
	}
}

// ParseRoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromID parses 'input' into a RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId
func ParseRoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromID(input string) (*RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromIDInsensitively parses 'input' case-insensitively into a RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromIDInsensitively(input string) (*RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleDefinitionId, ok = input.Parsed["unifiedRoleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", input)
	}

	if id.UnifiedRoleDefinitionId1, ok = input.Parsed["unifiedRoleDefinitionId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId1", input)
	}

	return nil
}

// ValidateRoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromID checks that 'input' can be parsed as a Role Management Exchange Role Definition Id Inherits Permissions From ID
func ValidateRoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Exchange Role Definition Id Inherits Permissions From ID
func (id RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId) ID() string {
	fmtString := "/roleManagement/exchange/roleDefinitions/%s/inheritsPermissionsFrom/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleDefinitionId, id.UnifiedRoleDefinitionId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Exchange Role Definition Id Inherits Permissions From ID
func (id RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("exchange", "exchange", "exchange"),
		resourceids.StaticSegment("roleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionId"),
		resourceids.StaticSegment("inheritsPermissionsFrom", "inheritsPermissionsFrom", "inheritsPermissionsFrom"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId1", "unifiedRoleDefinitionId1"),
	}
}

// String returns a human-readable description of this Role Management Exchange Role Definition Id Inherits Permissions From ID
func (id RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
		fmt.Sprintf("Unified Role Definition Id 1: %q", id.UnifiedRoleDefinitionId1),
	}
	return fmt.Sprintf("Role Management Exchange Role Definition Id Inherits Permissions From (%s)", strings.Join(components, "\n"))
}
