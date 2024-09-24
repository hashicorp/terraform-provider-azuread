package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccessAuthenticationStrengthPolicyId{}

// IdentityConditionalAccessAuthenticationStrengthPolicyId is a struct representing the Resource ID for a Identity Conditional Access Authentication Strength Policy
type IdentityConditionalAccessAuthenticationStrengthPolicyId struct {
	AuthenticationStrengthPolicyId string
}

// NewIdentityConditionalAccessAuthenticationStrengthPolicyID returns a new IdentityConditionalAccessAuthenticationStrengthPolicyId struct
func NewIdentityConditionalAccessAuthenticationStrengthPolicyID(authenticationStrengthPolicyId string) IdentityConditionalAccessAuthenticationStrengthPolicyId {
	return IdentityConditionalAccessAuthenticationStrengthPolicyId{
		AuthenticationStrengthPolicyId: authenticationStrengthPolicyId,
	}
}

// ParseIdentityConditionalAccessAuthenticationStrengthPolicyID parses 'input' into a IdentityConditionalAccessAuthenticationStrengthPolicyId
func ParseIdentityConditionalAccessAuthenticationStrengthPolicyID(input string) (*IdentityConditionalAccessAuthenticationStrengthPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccessAuthenticationStrengthPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccessAuthenticationStrengthPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityConditionalAccessAuthenticationStrengthPolicyIDInsensitively parses 'input' case-insensitively into a IdentityConditionalAccessAuthenticationStrengthPolicyId
// note: this method should only be used for API response data and not user input
func ParseIdentityConditionalAccessAuthenticationStrengthPolicyIDInsensitively(input string) (*IdentityConditionalAccessAuthenticationStrengthPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccessAuthenticationStrengthPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccessAuthenticationStrengthPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityConditionalAccessAuthenticationStrengthPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationStrengthPolicyId, ok = input.Parsed["authenticationStrengthPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationStrengthPolicyId", input)
	}

	return nil
}

// ValidateIdentityConditionalAccessAuthenticationStrengthPolicyID checks that 'input' can be parsed as a Identity Conditional Access Authentication Strength Policy ID
func ValidateIdentityConditionalAccessAuthenticationStrengthPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityConditionalAccessAuthenticationStrengthPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Conditional Access Authentication Strength Policy ID
func (id IdentityConditionalAccessAuthenticationStrengthPolicyId) ID() string {
	fmtString := "/identity/conditionalAccess/authenticationStrength/policies/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationStrengthPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Conditional Access Authentication Strength Policy ID
func (id IdentityConditionalAccessAuthenticationStrengthPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("conditionalAccess", "conditionalAccess", "conditionalAccess"),
		resourceids.StaticSegment("authenticationStrength", "authenticationStrength", "authenticationStrength"),
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.UserSpecifiedSegment("authenticationStrengthPolicyId", "authenticationStrengthPolicyId"),
	}
}

// String returns a human-readable description of this Identity Conditional Access Authentication Strength Policy ID
func (id IdentityConditionalAccessAuthenticationStrengthPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Strength Policy: %q", id.AuthenticationStrengthPolicyId),
	}
	return fmt.Sprintf("Identity Conditional Access Authentication Strength Policy (%s)", strings.Join(components, "\n"))
}
