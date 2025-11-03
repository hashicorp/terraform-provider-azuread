package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementCloudPCRoleAssignmentIdPrincipalId{}

// RoleManagementCloudPCRoleAssignmentIdPrincipalId is a struct representing the Resource ID for a Role Management Cloud PC Role Assignment Id Principal
type RoleManagementCloudPCRoleAssignmentIdPrincipalId struct {
	UnifiedRoleAssignmentMultipleId string
	DirectoryObjectId               string
}

// NewRoleManagementCloudPCRoleAssignmentIdPrincipalID returns a new RoleManagementCloudPCRoleAssignmentIdPrincipalId struct
func NewRoleManagementCloudPCRoleAssignmentIdPrincipalID(unifiedRoleAssignmentMultipleId string, directoryObjectId string) RoleManagementCloudPCRoleAssignmentIdPrincipalId {
	return RoleManagementCloudPCRoleAssignmentIdPrincipalId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
		DirectoryObjectId:               directoryObjectId,
	}
}

// ParseRoleManagementCloudPCRoleAssignmentIdPrincipalID parses 'input' into a RoleManagementCloudPCRoleAssignmentIdPrincipalId
func ParseRoleManagementCloudPCRoleAssignmentIdPrincipalID(input string) (*RoleManagementCloudPCRoleAssignmentIdPrincipalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementCloudPCRoleAssignmentIdPrincipalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementCloudPCRoleAssignmentIdPrincipalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementCloudPCRoleAssignmentIdPrincipalIDInsensitively parses 'input' case-insensitively into a RoleManagementCloudPCRoleAssignmentIdPrincipalId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementCloudPCRoleAssignmentIdPrincipalIDInsensitively(input string) (*RoleManagementCloudPCRoleAssignmentIdPrincipalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementCloudPCRoleAssignmentIdPrincipalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementCloudPCRoleAssignmentIdPrincipalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementCloudPCRoleAssignmentIdPrincipalId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentMultipleId, ok = input.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateRoleManagementCloudPCRoleAssignmentIdPrincipalID checks that 'input' can be parsed as a Role Management Cloud PC Role Assignment Id Principal ID
func ValidateRoleManagementCloudPCRoleAssignmentIdPrincipalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementCloudPCRoleAssignmentIdPrincipalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Cloud PC Role Assignment Id Principal ID
func (id RoleManagementCloudPCRoleAssignmentIdPrincipalId) ID() string {
	fmtString := "/roleManagement/cloudPC/roleAssignments/%s/principals/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Cloud PC Role Assignment Id Principal ID
func (id RoleManagementCloudPCRoleAssignmentIdPrincipalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("cloudPC", "cloudPC", "cloudPC"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleId"),
		resourceids.StaticSegment("principals", "principals", "principals"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Role Management Cloud PC Role Assignment Id Principal ID
func (id RoleManagementCloudPCRoleAssignmentIdPrincipalId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Role Management Cloud PC Role Assignment Id Principal (%s)", strings.Join(components, "\n"))
}
