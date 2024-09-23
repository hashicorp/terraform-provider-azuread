package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId{}

// PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId is a struct representing the Resource ID for a Policy Authentication Strength Policy Id Combination Configuration
type PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId struct {
	AuthenticationStrengthPolicyId           string
	AuthenticationCombinationConfigurationId string
}

// NewPolicyAuthenticationStrengthPolicyIdCombinationConfigurationID returns a new PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId struct
func NewPolicyAuthenticationStrengthPolicyIdCombinationConfigurationID(authenticationStrengthPolicyId string, authenticationCombinationConfigurationId string) PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId {
	return PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId{
		AuthenticationStrengthPolicyId:           authenticationStrengthPolicyId,
		AuthenticationCombinationConfigurationId: authenticationCombinationConfigurationId,
	}
}

// ParsePolicyAuthenticationStrengthPolicyIdCombinationConfigurationID parses 'input' into a PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId
func ParsePolicyAuthenticationStrengthPolicyIdCombinationConfigurationID(input string) (*PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyAuthenticationStrengthPolicyIdCombinationConfigurationIDInsensitively parses 'input' case-insensitively into a PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId
// note: this method should only be used for API response data and not user input
func ParsePolicyAuthenticationStrengthPolicyIdCombinationConfigurationIDInsensitively(input string) (*PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationStrengthPolicyId, ok = input.Parsed["authenticationStrengthPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationStrengthPolicyId", input)
	}

	if id.AuthenticationCombinationConfigurationId, ok = input.Parsed["authenticationCombinationConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationCombinationConfigurationId", input)
	}

	return nil
}

// ValidatePolicyAuthenticationStrengthPolicyIdCombinationConfigurationID checks that 'input' can be parsed as a Policy Authentication Strength Policy Id Combination Configuration ID
func ValidatePolicyAuthenticationStrengthPolicyIdCombinationConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyAuthenticationStrengthPolicyIdCombinationConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Authentication Strength Policy Id Combination Configuration ID
func (id PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId) ID() string {
	fmtString := "/policies/authenticationStrengthPolicies/%s/combinationConfigurations/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationStrengthPolicyId, id.AuthenticationCombinationConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Authentication Strength Policy Id Combination Configuration ID
func (id PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("authenticationStrengthPolicies", "authenticationStrengthPolicies", "authenticationStrengthPolicies"),
		resourceids.UserSpecifiedSegment("authenticationStrengthPolicyId", "authenticationStrengthPolicyId"),
		resourceids.StaticSegment("combinationConfigurations", "combinationConfigurations", "combinationConfigurations"),
		resourceids.UserSpecifiedSegment("authenticationCombinationConfigurationId", "authenticationCombinationConfigurationId"),
	}
}

// String returns a human-readable description of this Policy Authentication Strength Policy Id Combination Configuration ID
func (id PolicyAuthenticationStrengthPolicyIdCombinationConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Strength Policy: %q", id.AuthenticationStrengthPolicyId),
		fmt.Sprintf("Authentication Combination Configuration: %q", id.AuthenticationCombinationConfigurationId),
	}
	return fmt.Sprintf("Policy Authentication Strength Policy Id Combination Configuration (%s)", strings.Join(components, "\n"))
}
