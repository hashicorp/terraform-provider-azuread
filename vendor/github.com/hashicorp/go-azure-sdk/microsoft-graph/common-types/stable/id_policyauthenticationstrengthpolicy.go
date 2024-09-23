package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyAuthenticationStrengthPolicyId{}

// PolicyAuthenticationStrengthPolicyId is a struct representing the Resource ID for a Policy Authentication Strength Policy
type PolicyAuthenticationStrengthPolicyId struct {
	AuthenticationStrengthPolicyId string
}

// NewPolicyAuthenticationStrengthPolicyID returns a new PolicyAuthenticationStrengthPolicyId struct
func NewPolicyAuthenticationStrengthPolicyID(authenticationStrengthPolicyId string) PolicyAuthenticationStrengthPolicyId {
	return PolicyAuthenticationStrengthPolicyId{
		AuthenticationStrengthPolicyId: authenticationStrengthPolicyId,
	}
}

// ParsePolicyAuthenticationStrengthPolicyID parses 'input' into a PolicyAuthenticationStrengthPolicyId
func ParsePolicyAuthenticationStrengthPolicyID(input string) (*PolicyAuthenticationStrengthPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAuthenticationStrengthPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAuthenticationStrengthPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyAuthenticationStrengthPolicyIDInsensitively parses 'input' case-insensitively into a PolicyAuthenticationStrengthPolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyAuthenticationStrengthPolicyIDInsensitively(input string) (*PolicyAuthenticationStrengthPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAuthenticationStrengthPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAuthenticationStrengthPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyAuthenticationStrengthPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationStrengthPolicyId, ok = input.Parsed["authenticationStrengthPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationStrengthPolicyId", input)
	}

	return nil
}

// ValidatePolicyAuthenticationStrengthPolicyID checks that 'input' can be parsed as a Policy Authentication Strength Policy ID
func ValidatePolicyAuthenticationStrengthPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyAuthenticationStrengthPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Authentication Strength Policy ID
func (id PolicyAuthenticationStrengthPolicyId) ID() string {
	fmtString := "/policies/authenticationStrengthPolicies/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationStrengthPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Authentication Strength Policy ID
func (id PolicyAuthenticationStrengthPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("authenticationStrengthPolicies", "authenticationStrengthPolicies", "authenticationStrengthPolicies"),
		resourceids.UserSpecifiedSegment("authenticationStrengthPolicyId", "authenticationStrengthPolicyId"),
	}
}

// String returns a human-readable description of this Policy Authentication Strength Policy ID
func (id PolicyAuthenticationStrengthPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Strength Policy: %q", id.AuthenticationStrengthPolicyId),
	}
	return fmt.Sprintf("Policy Authentication Strength Policy (%s)", strings.Join(components, "\n"))
}
