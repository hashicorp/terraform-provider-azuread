package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyTokenLifetimePolicyId{}

// PolicyTokenLifetimePolicyId is a struct representing the Resource ID for a Policy Token Lifetime Policy
type PolicyTokenLifetimePolicyId struct {
	TokenLifetimePolicyId string
}

// NewPolicyTokenLifetimePolicyID returns a new PolicyTokenLifetimePolicyId struct
func NewPolicyTokenLifetimePolicyID(tokenLifetimePolicyId string) PolicyTokenLifetimePolicyId {
	return PolicyTokenLifetimePolicyId{
		TokenLifetimePolicyId: tokenLifetimePolicyId,
	}
}

// ParsePolicyTokenLifetimePolicyID parses 'input' into a PolicyTokenLifetimePolicyId
func ParsePolicyTokenLifetimePolicyID(input string) (*PolicyTokenLifetimePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyTokenLifetimePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyTokenLifetimePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyTokenLifetimePolicyIDInsensitively parses 'input' case-insensitively into a PolicyTokenLifetimePolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyTokenLifetimePolicyIDInsensitively(input string) (*PolicyTokenLifetimePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyTokenLifetimePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyTokenLifetimePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyTokenLifetimePolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TokenLifetimePolicyId, ok = input.Parsed["tokenLifetimePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "tokenLifetimePolicyId", input)
	}

	return nil
}

// ValidatePolicyTokenLifetimePolicyID checks that 'input' can be parsed as a Policy Token Lifetime Policy ID
func ValidatePolicyTokenLifetimePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyTokenLifetimePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Token Lifetime Policy ID
func (id PolicyTokenLifetimePolicyId) ID() string {
	fmtString := "/policies/tokenLifetimePolicies/%s"
	return fmt.Sprintf(fmtString, id.TokenLifetimePolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Token Lifetime Policy ID
func (id PolicyTokenLifetimePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("tokenLifetimePolicies", "tokenLifetimePolicies", "tokenLifetimePolicies"),
		resourceids.UserSpecifiedSegment("tokenLifetimePolicyId", "tokenLifetimePolicyId"),
	}
}

// String returns a human-readable description of this Policy Token Lifetime Policy ID
func (id PolicyTokenLifetimePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Token Lifetime Policy: %q", id.TokenLifetimePolicyId),
	}
	return fmt.Sprintf("Policy Token Lifetime Policy (%s)", strings.Join(components, "\n"))
}
