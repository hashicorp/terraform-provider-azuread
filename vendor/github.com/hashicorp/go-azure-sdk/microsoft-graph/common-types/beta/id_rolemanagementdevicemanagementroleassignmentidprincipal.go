package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDeviceManagementRoleAssignmentIdPrincipalId{}

// RoleManagementDeviceManagementRoleAssignmentIdPrincipalId is a struct representing the Resource ID for a Role Management Device Management Role Assignment Id Principal
type RoleManagementDeviceManagementRoleAssignmentIdPrincipalId struct {
	UnifiedRoleAssignmentMultipleId string
	DirectoryObjectId               string
}

// NewRoleManagementDeviceManagementRoleAssignmentIdPrincipalID returns a new RoleManagementDeviceManagementRoleAssignmentIdPrincipalId struct
func NewRoleManagementDeviceManagementRoleAssignmentIdPrincipalID(unifiedRoleAssignmentMultipleId string, directoryObjectId string) RoleManagementDeviceManagementRoleAssignmentIdPrincipalId {
	return RoleManagementDeviceManagementRoleAssignmentIdPrincipalId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
		DirectoryObjectId:               directoryObjectId,
	}
}

// ParseRoleManagementDeviceManagementRoleAssignmentIdPrincipalID parses 'input' into a RoleManagementDeviceManagementRoleAssignmentIdPrincipalId
func ParseRoleManagementDeviceManagementRoleAssignmentIdPrincipalID(input string) (*RoleManagementDeviceManagementRoleAssignmentIdPrincipalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDeviceManagementRoleAssignmentIdPrincipalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDeviceManagementRoleAssignmentIdPrincipalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDeviceManagementRoleAssignmentIdPrincipalIDInsensitively parses 'input' case-insensitively into a RoleManagementDeviceManagementRoleAssignmentIdPrincipalId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDeviceManagementRoleAssignmentIdPrincipalIDInsensitively(input string) (*RoleManagementDeviceManagementRoleAssignmentIdPrincipalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDeviceManagementRoleAssignmentIdPrincipalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDeviceManagementRoleAssignmentIdPrincipalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDeviceManagementRoleAssignmentIdPrincipalId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentMultipleId, ok = input.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateRoleManagementDeviceManagementRoleAssignmentIdPrincipalID checks that 'input' can be parsed as a Role Management Device Management Role Assignment Id Principal ID
func ValidateRoleManagementDeviceManagementRoleAssignmentIdPrincipalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDeviceManagementRoleAssignmentIdPrincipalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Device Management Role Assignment Id Principal ID
func (id RoleManagementDeviceManagementRoleAssignmentIdPrincipalId) ID() string {
	fmtString := "/roleManagement/deviceManagement/roleAssignments/%s/principals/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Device Management Role Assignment Id Principal ID
func (id RoleManagementDeviceManagementRoleAssignmentIdPrincipalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleId"),
		resourceids.StaticSegment("principals", "principals", "principals"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Role Management Device Management Role Assignment Id Principal ID
func (id RoleManagementDeviceManagementRoleAssignmentIdPrincipalId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Role Management Device Management Role Assignment Id Principal (%s)", strings.Join(components, "\n"))
}
