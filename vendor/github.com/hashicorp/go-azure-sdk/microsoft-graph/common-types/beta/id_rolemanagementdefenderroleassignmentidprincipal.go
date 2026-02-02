package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDefenderRoleAssignmentIdPrincipalId{}

// RoleManagementDefenderRoleAssignmentIdPrincipalId is a struct representing the Resource ID for a Role Management Defender Role Assignment Id Principal
type RoleManagementDefenderRoleAssignmentIdPrincipalId struct {
	UnifiedRoleAssignmentMultipleId string
	DirectoryObjectId               string
}

// NewRoleManagementDefenderRoleAssignmentIdPrincipalID returns a new RoleManagementDefenderRoleAssignmentIdPrincipalId struct
func NewRoleManagementDefenderRoleAssignmentIdPrincipalID(unifiedRoleAssignmentMultipleId string, directoryObjectId string) RoleManagementDefenderRoleAssignmentIdPrincipalId {
	return RoleManagementDefenderRoleAssignmentIdPrincipalId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
		DirectoryObjectId:               directoryObjectId,
	}
}

// ParseRoleManagementDefenderRoleAssignmentIdPrincipalID parses 'input' into a RoleManagementDefenderRoleAssignmentIdPrincipalId
func ParseRoleManagementDefenderRoleAssignmentIdPrincipalID(input string) (*RoleManagementDefenderRoleAssignmentIdPrincipalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderRoleAssignmentIdPrincipalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderRoleAssignmentIdPrincipalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDefenderRoleAssignmentIdPrincipalIDInsensitively parses 'input' case-insensitively into a RoleManagementDefenderRoleAssignmentIdPrincipalId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDefenderRoleAssignmentIdPrincipalIDInsensitively(input string) (*RoleManagementDefenderRoleAssignmentIdPrincipalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderRoleAssignmentIdPrincipalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderRoleAssignmentIdPrincipalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDefenderRoleAssignmentIdPrincipalId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentMultipleId, ok = input.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateRoleManagementDefenderRoleAssignmentIdPrincipalID checks that 'input' can be parsed as a Role Management Defender Role Assignment Id Principal ID
func ValidateRoleManagementDefenderRoleAssignmentIdPrincipalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDefenderRoleAssignmentIdPrincipalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Defender Role Assignment Id Principal ID
func (id RoleManagementDefenderRoleAssignmentIdPrincipalId) ID() string {
	fmtString := "/roleManagement/defender/roleAssignments/%s/principals/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Defender Role Assignment Id Principal ID
func (id RoleManagementDefenderRoleAssignmentIdPrincipalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("defender", "defender", "defender"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleId"),
		resourceids.StaticSegment("principals", "principals", "principals"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Role Management Defender Role Assignment Id Principal ID
func (id RoleManagementDefenderRoleAssignmentIdPrincipalId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Role Management Defender Role Assignment Id Principal (%s)", strings.Join(components, "\n"))
}
