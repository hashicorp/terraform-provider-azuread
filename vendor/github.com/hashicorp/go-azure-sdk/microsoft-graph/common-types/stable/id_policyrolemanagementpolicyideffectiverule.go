package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyRoleManagementPolicyIdEffectiveRuleId{}

// PolicyRoleManagementPolicyIdEffectiveRuleId is a struct representing the Resource ID for a Policy Role Management Policy Id Effective Rule
type PolicyRoleManagementPolicyIdEffectiveRuleId struct {
	UnifiedRoleManagementPolicyId     string
	UnifiedRoleManagementPolicyRuleId string
}

// NewPolicyRoleManagementPolicyIdEffectiveRuleID returns a new PolicyRoleManagementPolicyIdEffectiveRuleId struct
func NewPolicyRoleManagementPolicyIdEffectiveRuleID(unifiedRoleManagementPolicyId string, unifiedRoleManagementPolicyRuleId string) PolicyRoleManagementPolicyIdEffectiveRuleId {
	return PolicyRoleManagementPolicyIdEffectiveRuleId{
		UnifiedRoleManagementPolicyId:     unifiedRoleManagementPolicyId,
		UnifiedRoleManagementPolicyRuleId: unifiedRoleManagementPolicyRuleId,
	}
}

// ParsePolicyRoleManagementPolicyIdEffectiveRuleID parses 'input' into a PolicyRoleManagementPolicyIdEffectiveRuleId
func ParsePolicyRoleManagementPolicyIdEffectiveRuleID(input string) (*PolicyRoleManagementPolicyIdEffectiveRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyRoleManagementPolicyIdEffectiveRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyRoleManagementPolicyIdEffectiveRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyRoleManagementPolicyIdEffectiveRuleIDInsensitively parses 'input' case-insensitively into a PolicyRoleManagementPolicyIdEffectiveRuleId
// note: this method should only be used for API response data and not user input
func ParsePolicyRoleManagementPolicyIdEffectiveRuleIDInsensitively(input string) (*PolicyRoleManagementPolicyIdEffectiveRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyRoleManagementPolicyIdEffectiveRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyRoleManagementPolicyIdEffectiveRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyRoleManagementPolicyIdEffectiveRuleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleManagementPolicyId, ok = input.Parsed["unifiedRoleManagementPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementPolicyId", input)
	}

	if id.UnifiedRoleManagementPolicyRuleId, ok = input.Parsed["unifiedRoleManagementPolicyRuleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementPolicyRuleId", input)
	}

	return nil
}

// ValidatePolicyRoleManagementPolicyIdEffectiveRuleID checks that 'input' can be parsed as a Policy Role Management Policy Id Effective Rule ID
func ValidatePolicyRoleManagementPolicyIdEffectiveRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyRoleManagementPolicyIdEffectiveRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Role Management Policy Id Effective Rule ID
func (id PolicyRoleManagementPolicyIdEffectiveRuleId) ID() string {
	fmtString := "/policies/roleManagementPolicies/%s/effectiveRules/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleManagementPolicyId, id.UnifiedRoleManagementPolicyRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Role Management Policy Id Effective Rule ID
func (id PolicyRoleManagementPolicyIdEffectiveRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("roleManagementPolicies", "roleManagementPolicies", "roleManagementPolicies"),
		resourceids.UserSpecifiedSegment("unifiedRoleManagementPolicyId", "unifiedRoleManagementPolicyId"),
		resourceids.StaticSegment("effectiveRules", "effectiveRules", "effectiveRules"),
		resourceids.UserSpecifiedSegment("unifiedRoleManagementPolicyRuleId", "unifiedRoleManagementPolicyRuleId"),
	}
}

// String returns a human-readable description of this Policy Role Management Policy Id Effective Rule ID
func (id PolicyRoleManagementPolicyIdEffectiveRuleId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Management Policy: %q", id.UnifiedRoleManagementPolicyId),
		fmt.Sprintf("Unified Role Management Policy Rule: %q", id.UnifiedRoleManagementPolicyRuleId),
	}
	return fmt.Sprintf("Policy Role Management Policy Id Effective Rule (%s)", strings.Join(components, "\n"))
}
