package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId{}

// PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId is a struct representing the Resource ID for a Policy Authentication Methods Policy Authentication Method Configuration
type PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId struct {
	AuthenticationMethodConfigurationId string
}

// NewPolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationID returns a new PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId struct
func NewPolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationID(authenticationMethodConfigurationId string) PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId {
	return PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId{
		AuthenticationMethodConfigurationId: authenticationMethodConfigurationId,
	}
}

// ParsePolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationID parses 'input' into a PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId
func ParsePolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationID(input string) (*PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationIDInsensitively parses 'input' case-insensitively into a PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId
// note: this method should only be used for API response data and not user input
func ParsePolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationIDInsensitively(input string) (*PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationMethodConfigurationId, ok = input.Parsed["authenticationMethodConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationMethodConfigurationId", input)
	}

	return nil
}

// ValidatePolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationID checks that 'input' can be parsed as a Policy Authentication Methods Policy Authentication Method Configuration ID
func ValidatePolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Authentication Methods Policy Authentication Method Configuration ID
func (id PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId) ID() string {
	fmtString := "/policies/authenticationMethodsPolicy/authenticationMethodConfigurations/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationMethodConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Authentication Methods Policy Authentication Method Configuration ID
func (id PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.StaticSegment("authenticationMethodsPolicy", "authenticationMethodsPolicy", "authenticationMethodsPolicy"),
		resourceids.StaticSegment("authenticationMethodConfigurations", "authenticationMethodConfigurations", "authenticationMethodConfigurations"),
		resourceids.UserSpecifiedSegment("authenticationMethodConfigurationId", "authenticationMethodConfigurationId"),
	}
}

// String returns a human-readable description of this Policy Authentication Methods Policy Authentication Method Configuration ID
func (id PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Method Configuration: %q", id.AuthenticationMethodConfigurationId),
	}
	return fmt.Sprintf("Policy Authentication Methods Policy Authentication Method Configuration (%s)", strings.Join(components, "\n"))
}
