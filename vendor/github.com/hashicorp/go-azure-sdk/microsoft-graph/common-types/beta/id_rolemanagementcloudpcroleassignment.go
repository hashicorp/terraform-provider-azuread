package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementCloudPCRoleAssignmentId{}

// RoleManagementCloudPCRoleAssignmentId is a struct representing the Resource ID for a Role Management Cloud PC Role Assignment
type RoleManagementCloudPCRoleAssignmentId struct {
	UnifiedRoleAssignmentMultipleId string
}

// NewRoleManagementCloudPCRoleAssignmentID returns a new RoleManagementCloudPCRoleAssignmentId struct
func NewRoleManagementCloudPCRoleAssignmentID(unifiedRoleAssignmentMultipleId string) RoleManagementCloudPCRoleAssignmentId {
	return RoleManagementCloudPCRoleAssignmentId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
	}
}

// ParseRoleManagementCloudPCRoleAssignmentID parses 'input' into a RoleManagementCloudPCRoleAssignmentId
func ParseRoleManagementCloudPCRoleAssignmentID(input string) (*RoleManagementCloudPCRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementCloudPCRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementCloudPCRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementCloudPCRoleAssignmentIDInsensitively parses 'input' case-insensitively into a RoleManagementCloudPCRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementCloudPCRoleAssignmentIDInsensitively(input string) (*RoleManagementCloudPCRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementCloudPCRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementCloudPCRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementCloudPCRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentMultipleId, ok = input.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", input)
	}

	return nil
}

// ValidateRoleManagementCloudPCRoleAssignmentID checks that 'input' can be parsed as a Role Management Cloud PC Role Assignment ID
func ValidateRoleManagementCloudPCRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementCloudPCRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Cloud PC Role Assignment ID
func (id RoleManagementCloudPCRoleAssignmentId) ID() string {
	fmtString := "/roleManagement/cloudPC/roleAssignments/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Cloud PC Role Assignment ID
func (id RoleManagementCloudPCRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("cloudPC", "cloudPC", "cloudPC"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleId"),
	}
}

// String returns a human-readable description of this Role Management Cloud PC Role Assignment ID
func (id RoleManagementCloudPCRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
	}
	return fmt.Sprintf("Role Management Cloud PC Role Assignment (%s)", strings.Join(components, "\n"))
}
