package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyRoleManagementPolicyAssignmentId{}

// PolicyRoleManagementPolicyAssignmentId is a struct representing the Resource ID for a Policy Role Management Policy Assignment
type PolicyRoleManagementPolicyAssignmentId struct {
	UnifiedRoleManagementPolicyAssignmentId string
}

// NewPolicyRoleManagementPolicyAssignmentID returns a new PolicyRoleManagementPolicyAssignmentId struct
func NewPolicyRoleManagementPolicyAssignmentID(unifiedRoleManagementPolicyAssignmentId string) PolicyRoleManagementPolicyAssignmentId {
	return PolicyRoleManagementPolicyAssignmentId{
		UnifiedRoleManagementPolicyAssignmentId: unifiedRoleManagementPolicyAssignmentId,
	}
}

// ParsePolicyRoleManagementPolicyAssignmentID parses 'input' into a PolicyRoleManagementPolicyAssignmentId
func ParsePolicyRoleManagementPolicyAssignmentID(input string) (*PolicyRoleManagementPolicyAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyRoleManagementPolicyAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyRoleManagementPolicyAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyRoleManagementPolicyAssignmentIDInsensitively parses 'input' case-insensitively into a PolicyRoleManagementPolicyAssignmentId
// note: this method should only be used for API response data and not user input
func ParsePolicyRoleManagementPolicyAssignmentIDInsensitively(input string) (*PolicyRoleManagementPolicyAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyRoleManagementPolicyAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyRoleManagementPolicyAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyRoleManagementPolicyAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleManagementPolicyAssignmentId, ok = input.Parsed["unifiedRoleManagementPolicyAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementPolicyAssignmentId", input)
	}

	return nil
}

// ValidatePolicyRoleManagementPolicyAssignmentID checks that 'input' can be parsed as a Policy Role Management Policy Assignment ID
func ValidatePolicyRoleManagementPolicyAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyRoleManagementPolicyAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Role Management Policy Assignment ID
func (id PolicyRoleManagementPolicyAssignmentId) ID() string {
	fmtString := "/policies/roleManagementPolicyAssignments/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleManagementPolicyAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Role Management Policy Assignment ID
func (id PolicyRoleManagementPolicyAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("roleManagementPolicyAssignments", "roleManagementPolicyAssignments", "roleManagementPolicyAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleManagementPolicyAssignmentId", "unifiedRoleManagementPolicyAssignmentId"),
	}
}

// String returns a human-readable description of this Policy Role Management Policy Assignment ID
func (id PolicyRoleManagementPolicyAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Management Policy Assignment: %q", id.UnifiedRoleManagementPolicyAssignmentId),
	}
	return fmt.Sprintf("Policy Role Management Policy Assignment (%s)", strings.Join(components, "\n"))
}
