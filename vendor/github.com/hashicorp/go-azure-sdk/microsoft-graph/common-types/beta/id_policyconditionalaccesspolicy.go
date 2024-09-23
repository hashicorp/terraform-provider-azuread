package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyConditionalAccessPolicyId{}

// PolicyConditionalAccessPolicyId is a struct representing the Resource ID for a Policy Conditional Access Policy
type PolicyConditionalAccessPolicyId struct {
	ConditionalAccessPolicyId string
}

// NewPolicyConditionalAccessPolicyID returns a new PolicyConditionalAccessPolicyId struct
func NewPolicyConditionalAccessPolicyID(conditionalAccessPolicyId string) PolicyConditionalAccessPolicyId {
	return PolicyConditionalAccessPolicyId{
		ConditionalAccessPolicyId: conditionalAccessPolicyId,
	}
}

// ParsePolicyConditionalAccessPolicyID parses 'input' into a PolicyConditionalAccessPolicyId
func ParsePolicyConditionalAccessPolicyID(input string) (*PolicyConditionalAccessPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyConditionalAccessPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyConditionalAccessPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyConditionalAccessPolicyIDInsensitively parses 'input' case-insensitively into a PolicyConditionalAccessPolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyConditionalAccessPolicyIDInsensitively(input string) (*PolicyConditionalAccessPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyConditionalAccessPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyConditionalAccessPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyConditionalAccessPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ConditionalAccessPolicyId, ok = input.Parsed["conditionalAccessPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conditionalAccessPolicyId", input)
	}

	return nil
}

// ValidatePolicyConditionalAccessPolicyID checks that 'input' can be parsed as a Policy Conditional Access Policy ID
func ValidatePolicyConditionalAccessPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyConditionalAccessPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Conditional Access Policy ID
func (id PolicyConditionalAccessPolicyId) ID() string {
	fmtString := "/policies/conditionalAccessPolicies/%s"
	return fmt.Sprintf(fmtString, id.ConditionalAccessPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Conditional Access Policy ID
func (id PolicyConditionalAccessPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("conditionalAccessPolicies", "conditionalAccessPolicies", "conditionalAccessPolicies"),
		resourceids.UserSpecifiedSegment("conditionalAccessPolicyId", "conditionalAccessPolicyId"),
	}
}

// String returns a human-readable description of this Policy Conditional Access Policy ID
func (id PolicyConditionalAccessPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Conditional Access Policy: %q", id.ConditionalAccessPolicyId),
	}
	return fmt.Sprintf("Policy Conditional Access Policy (%s)", strings.Join(components, "\n"))
}
