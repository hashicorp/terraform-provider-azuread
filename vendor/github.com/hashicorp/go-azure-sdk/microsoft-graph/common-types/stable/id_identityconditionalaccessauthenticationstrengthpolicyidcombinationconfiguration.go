package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId{}

// IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId is a struct representing the Resource ID for a Identity Conditional Access Authentication Strength Policy Id Combination Configuration
type IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId struct {
	AuthenticationStrengthPolicyId           string
	AuthenticationCombinationConfigurationId string
}

// NewIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationID returns a new IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId struct
func NewIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationID(authenticationStrengthPolicyId string, authenticationCombinationConfigurationId string) IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId {
	return IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId{
		AuthenticationStrengthPolicyId:           authenticationStrengthPolicyId,
		AuthenticationCombinationConfigurationId: authenticationCombinationConfigurationId,
	}
}

// ParseIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationID parses 'input' into a IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId
func ParseIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationID(input string) (*IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationIDInsensitively parses 'input' case-insensitively into a IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId
// note: this method should only be used for API response data and not user input
func ParseIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationIDInsensitively(input string) (*IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationStrengthPolicyId, ok = input.Parsed["authenticationStrengthPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationStrengthPolicyId", input)
	}

	if id.AuthenticationCombinationConfigurationId, ok = input.Parsed["authenticationCombinationConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationCombinationConfigurationId", input)
	}

	return nil
}

// ValidateIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationID checks that 'input' can be parsed as a Identity Conditional Access Authentication Strength Policy Id Combination Configuration ID
func ValidateIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Conditional Access Authentication Strength Policy Id Combination Configuration ID
func (id IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId) ID() string {
	fmtString := "/identity/conditionalAccess/authenticationStrength/policies/%s/combinationConfigurations/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationStrengthPolicyId, id.AuthenticationCombinationConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Conditional Access Authentication Strength Policy Id Combination Configuration ID
func (id IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("conditionalAccess", "conditionalAccess", "conditionalAccess"),
		resourceids.StaticSegment("authenticationStrength", "authenticationStrength", "authenticationStrength"),
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.UserSpecifiedSegment("authenticationStrengthPolicyId", "authenticationStrengthPolicyId"),
		resourceids.StaticSegment("combinationConfigurations", "combinationConfigurations", "combinationConfigurations"),
		resourceids.UserSpecifiedSegment("authenticationCombinationConfigurationId", "authenticationCombinationConfigurationId"),
	}
}

// String returns a human-readable description of this Identity Conditional Access Authentication Strength Policy Id Combination Configuration ID
func (id IdentityConditionalAccessAuthenticationStrengthPolicyIdCombinationConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Strength Policy: %q", id.AuthenticationStrengthPolicyId),
		fmt.Sprintf("Authentication Combination Configuration: %q", id.AuthenticationCombinationConfigurationId),
	}
	return fmt.Sprintf("Identity Conditional Access Authentication Strength Policy Id Combination Configuration (%s)", strings.Join(components, "\n"))
}
