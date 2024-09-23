package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId{}

// IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId is a struct representing the Resource ID for a Identity Conditional Access Authentication Strength Authentication Method Mode
type IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId struct {
	AuthenticationMethodModeDetailId string
}

// NewIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeID returns a new IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId struct
func NewIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeID(authenticationMethodModeDetailId string) IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId {
	return IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId{
		AuthenticationMethodModeDetailId: authenticationMethodModeDetailId,
	}
}

// ParseIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeID parses 'input' into a IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId
func ParseIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeID(input string) (*IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeIDInsensitively parses 'input' case-insensitively into a IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId
// note: this method should only be used for API response data and not user input
func ParseIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeIDInsensitively(input string) (*IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationMethodModeDetailId, ok = input.Parsed["authenticationMethodModeDetailId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationMethodModeDetailId", input)
	}

	return nil
}

// ValidateIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeID checks that 'input' can be parsed as a Identity Conditional Access Authentication Strength Authentication Method Mode ID
func ValidateIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Conditional Access Authentication Strength Authentication Method Mode ID
func (id IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId) ID() string {
	fmtString := "/identity/conditionalAccess/authenticationStrength/authenticationMethodModes/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationMethodModeDetailId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Conditional Access Authentication Strength Authentication Method Mode ID
func (id IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("conditionalAccess", "conditionalAccess", "conditionalAccess"),
		resourceids.StaticSegment("authenticationStrength", "authenticationStrength", "authenticationStrength"),
		resourceids.StaticSegment("authenticationMethodModes", "authenticationMethodModes", "authenticationMethodModes"),
		resourceids.UserSpecifiedSegment("authenticationMethodModeDetailId", "authenticationMethodModeDetailId"),
	}
}

// String returns a human-readable description of this Identity Conditional Access Authentication Strength Authentication Method Mode ID
func (id IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Method Mode Detail: %q", id.AuthenticationMethodModeDetailId),
	}
	return fmt.Sprintf("Identity Conditional Access Authentication Strength Authentication Method Mode (%s)", strings.Join(components, "\n"))
}
