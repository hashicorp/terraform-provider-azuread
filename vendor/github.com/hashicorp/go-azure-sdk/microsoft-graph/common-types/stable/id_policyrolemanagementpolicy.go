package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyRoleManagementPolicyId{}

// PolicyRoleManagementPolicyId is a struct representing the Resource ID for a Policy Role Management Policy
type PolicyRoleManagementPolicyId struct {
	UnifiedRoleManagementPolicyId string
}

// NewPolicyRoleManagementPolicyID returns a new PolicyRoleManagementPolicyId struct
func NewPolicyRoleManagementPolicyID(unifiedRoleManagementPolicyId string) PolicyRoleManagementPolicyId {
	return PolicyRoleManagementPolicyId{
		UnifiedRoleManagementPolicyId: unifiedRoleManagementPolicyId,
	}
}

// ParsePolicyRoleManagementPolicyID parses 'input' into a PolicyRoleManagementPolicyId
func ParsePolicyRoleManagementPolicyID(input string) (*PolicyRoleManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyRoleManagementPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyRoleManagementPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyRoleManagementPolicyIDInsensitively parses 'input' case-insensitively into a PolicyRoleManagementPolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyRoleManagementPolicyIDInsensitively(input string) (*PolicyRoleManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyRoleManagementPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyRoleManagementPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyRoleManagementPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleManagementPolicyId, ok = input.Parsed["unifiedRoleManagementPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementPolicyId", input)
	}

	return nil
}

// ValidatePolicyRoleManagementPolicyID checks that 'input' can be parsed as a Policy Role Management Policy ID
func ValidatePolicyRoleManagementPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyRoleManagementPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Role Management Policy ID
func (id PolicyRoleManagementPolicyId) ID() string {
	fmtString := "/policies/roleManagementPolicies/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleManagementPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Role Management Policy ID
func (id PolicyRoleManagementPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("roleManagementPolicies", "roleManagementPolicies", "roleManagementPolicies"),
		resourceids.UserSpecifiedSegment("unifiedRoleManagementPolicyId", "unifiedRoleManagementPolicyId"),
	}
}

// String returns a human-readable description of this Policy Role Management Policy ID
func (id PolicyRoleManagementPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Management Policy: %q", id.UnifiedRoleManagementPolicyId),
	}
	return fmt.Sprintf("Policy Role Management Policy (%s)", strings.Join(components, "\n"))
}
