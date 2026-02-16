package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyRoleManagementPolicyIdRuleId{}

// PolicyRoleManagementPolicyIdRuleId is a struct representing the Resource ID for a Policy Role Management Policy Id Rule
type PolicyRoleManagementPolicyIdRuleId struct {
	UnifiedRoleManagementPolicyId     string
	UnifiedRoleManagementPolicyRuleId string
}

// NewPolicyRoleManagementPolicyIdRuleID returns a new PolicyRoleManagementPolicyIdRuleId struct
func NewPolicyRoleManagementPolicyIdRuleID(unifiedRoleManagementPolicyId string, unifiedRoleManagementPolicyRuleId string) PolicyRoleManagementPolicyIdRuleId {
	return PolicyRoleManagementPolicyIdRuleId{
		UnifiedRoleManagementPolicyId:     unifiedRoleManagementPolicyId,
		UnifiedRoleManagementPolicyRuleId: unifiedRoleManagementPolicyRuleId,
	}
}

// ParsePolicyRoleManagementPolicyIdRuleID parses 'input' into a PolicyRoleManagementPolicyIdRuleId
func ParsePolicyRoleManagementPolicyIdRuleID(input string) (*PolicyRoleManagementPolicyIdRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyRoleManagementPolicyIdRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyRoleManagementPolicyIdRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyRoleManagementPolicyIdRuleIDInsensitively parses 'input' case-insensitively into a PolicyRoleManagementPolicyIdRuleId
// note: this method should only be used for API response data and not user input
func ParsePolicyRoleManagementPolicyIdRuleIDInsensitively(input string) (*PolicyRoleManagementPolicyIdRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyRoleManagementPolicyIdRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyRoleManagementPolicyIdRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyRoleManagementPolicyIdRuleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleManagementPolicyId, ok = input.Parsed["unifiedRoleManagementPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementPolicyId", input)
	}

	if id.UnifiedRoleManagementPolicyRuleId, ok = input.Parsed["unifiedRoleManagementPolicyRuleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementPolicyRuleId", input)
	}

	return nil
}

// ValidatePolicyRoleManagementPolicyIdRuleID checks that 'input' can be parsed as a Policy Role Management Policy Id Rule ID
func ValidatePolicyRoleManagementPolicyIdRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyRoleManagementPolicyIdRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Role Management Policy Id Rule ID
func (id PolicyRoleManagementPolicyIdRuleId) ID() string {
	fmtString := "/policies/roleManagementPolicies/%s/rules/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleManagementPolicyId, id.UnifiedRoleManagementPolicyRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Role Management Policy Id Rule ID
func (id PolicyRoleManagementPolicyIdRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("roleManagementPolicies", "roleManagementPolicies", "roleManagementPolicies"),
		resourceids.UserSpecifiedSegment("unifiedRoleManagementPolicyId", "unifiedRoleManagementPolicyId"),
		resourceids.StaticSegment("rules", "rules", "rules"),
		resourceids.UserSpecifiedSegment("unifiedRoleManagementPolicyRuleId", "unifiedRoleManagementPolicyRuleId"),
	}
}

// String returns a human-readable description of this Policy Role Management Policy Id Rule ID
func (id PolicyRoleManagementPolicyIdRuleId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Management Policy: %q", id.UnifiedRoleManagementPolicyId),
		fmt.Sprintf("Unified Role Management Policy Rule: %q", id.UnifiedRoleManagementPolicyRuleId),
	}
	return fmt.Sprintf("Policy Role Management Policy Id Rule (%s)", strings.Join(components, "\n"))
}
