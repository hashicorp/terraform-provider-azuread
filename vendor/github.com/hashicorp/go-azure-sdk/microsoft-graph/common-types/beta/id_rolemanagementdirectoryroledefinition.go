package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryRoleDefinitionId{}

// RoleManagementDirectoryRoleDefinitionId is a struct representing the Resource ID for a Role Management Directory Role Definition
type RoleManagementDirectoryRoleDefinitionId struct {
	UnifiedRoleDefinitionId string
}

// NewRoleManagementDirectoryRoleDefinitionID returns a new RoleManagementDirectoryRoleDefinitionId struct
func NewRoleManagementDirectoryRoleDefinitionID(unifiedRoleDefinitionId string) RoleManagementDirectoryRoleDefinitionId {
	return RoleManagementDirectoryRoleDefinitionId{
		UnifiedRoleDefinitionId: unifiedRoleDefinitionId,
	}
}

// ParseRoleManagementDirectoryRoleDefinitionID parses 'input' into a RoleManagementDirectoryRoleDefinitionId
func ParseRoleManagementDirectoryRoleDefinitionID(input string) (*RoleManagementDirectoryRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDirectoryRoleDefinitionIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryRoleDefinitionIDInsensitively(input string) (*RoleManagementDirectoryRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDirectoryRoleDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleDefinitionId, ok = input.Parsed["unifiedRoleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", input)
	}

	return nil
}

// ValidateRoleManagementDirectoryRoleDefinitionID checks that 'input' can be parsed as a Role Management Directory Role Definition ID
func ValidateRoleManagementDirectoryRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Role Definition ID
func (id RoleManagementDirectoryRoleDefinitionId) ID() string {
	fmtString := "/roleManagement/directory/roleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Role Definition ID
func (id RoleManagementDirectoryRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("roleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionId"),
	}
}

// String returns a human-readable description of this Role Management Directory Role Definition ID
func (id RoleManagementDirectoryRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
	}
	return fmt.Sprintf("Role Management Directory Role Definition (%s)", strings.Join(components, "\n"))
}
