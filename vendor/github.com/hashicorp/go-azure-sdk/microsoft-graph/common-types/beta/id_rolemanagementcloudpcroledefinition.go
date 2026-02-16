package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementCloudPCRoleDefinitionId{}

// RoleManagementCloudPCRoleDefinitionId is a struct representing the Resource ID for a Role Management Cloud PC Role Definition
type RoleManagementCloudPCRoleDefinitionId struct {
	UnifiedRoleDefinitionId string
}

// NewRoleManagementCloudPCRoleDefinitionID returns a new RoleManagementCloudPCRoleDefinitionId struct
func NewRoleManagementCloudPCRoleDefinitionID(unifiedRoleDefinitionId string) RoleManagementCloudPCRoleDefinitionId {
	return RoleManagementCloudPCRoleDefinitionId{
		UnifiedRoleDefinitionId: unifiedRoleDefinitionId,
	}
}

// ParseRoleManagementCloudPCRoleDefinitionID parses 'input' into a RoleManagementCloudPCRoleDefinitionId
func ParseRoleManagementCloudPCRoleDefinitionID(input string) (*RoleManagementCloudPCRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementCloudPCRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementCloudPCRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementCloudPCRoleDefinitionIDInsensitively parses 'input' case-insensitively into a RoleManagementCloudPCRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementCloudPCRoleDefinitionIDInsensitively(input string) (*RoleManagementCloudPCRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementCloudPCRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementCloudPCRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementCloudPCRoleDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleDefinitionId, ok = input.Parsed["unifiedRoleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", input)
	}

	return nil
}

// ValidateRoleManagementCloudPCRoleDefinitionID checks that 'input' can be parsed as a Role Management Cloud PC Role Definition ID
func ValidateRoleManagementCloudPCRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementCloudPCRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Cloud PC Role Definition ID
func (id RoleManagementCloudPCRoleDefinitionId) ID() string {
	fmtString := "/roleManagement/cloudPC/roleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Cloud PC Role Definition ID
func (id RoleManagementCloudPCRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("cloudPC", "cloudPC", "cloudPC"),
		resourceids.StaticSegment("roleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionId"),
	}
}

// String returns a human-readable description of this Role Management Cloud PC Role Definition ID
func (id RoleManagementCloudPCRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
	}
	return fmt.Sprintf("Role Management Cloud PC Role Definition (%s)", strings.Join(components, "\n"))
}
