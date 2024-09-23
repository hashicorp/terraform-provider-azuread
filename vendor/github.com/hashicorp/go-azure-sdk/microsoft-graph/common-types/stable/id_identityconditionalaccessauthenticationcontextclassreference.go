package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccessAuthenticationContextClassReferenceId{}

// IdentityConditionalAccessAuthenticationContextClassReferenceId is a struct representing the Resource ID for a Identity Conditional Access Authentication Context Class Reference
type IdentityConditionalAccessAuthenticationContextClassReferenceId struct {
	AuthenticationContextClassReferenceId string
}

// NewIdentityConditionalAccessAuthenticationContextClassReferenceID returns a new IdentityConditionalAccessAuthenticationContextClassReferenceId struct
func NewIdentityConditionalAccessAuthenticationContextClassReferenceID(authenticationContextClassReferenceId string) IdentityConditionalAccessAuthenticationContextClassReferenceId {
	return IdentityConditionalAccessAuthenticationContextClassReferenceId{
		AuthenticationContextClassReferenceId: authenticationContextClassReferenceId,
	}
}

// ParseIdentityConditionalAccessAuthenticationContextClassReferenceID parses 'input' into a IdentityConditionalAccessAuthenticationContextClassReferenceId
func ParseIdentityConditionalAccessAuthenticationContextClassReferenceID(input string) (*IdentityConditionalAccessAuthenticationContextClassReferenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccessAuthenticationContextClassReferenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccessAuthenticationContextClassReferenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityConditionalAccessAuthenticationContextClassReferenceIDInsensitively parses 'input' case-insensitively into a IdentityConditionalAccessAuthenticationContextClassReferenceId
// note: this method should only be used for API response data and not user input
func ParseIdentityConditionalAccessAuthenticationContextClassReferenceIDInsensitively(input string) (*IdentityConditionalAccessAuthenticationContextClassReferenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccessAuthenticationContextClassReferenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccessAuthenticationContextClassReferenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityConditionalAccessAuthenticationContextClassReferenceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationContextClassReferenceId, ok = input.Parsed["authenticationContextClassReferenceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationContextClassReferenceId", input)
	}

	return nil
}

// ValidateIdentityConditionalAccessAuthenticationContextClassReferenceID checks that 'input' can be parsed as a Identity Conditional Access Authentication Context Class Reference ID
func ValidateIdentityConditionalAccessAuthenticationContextClassReferenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityConditionalAccessAuthenticationContextClassReferenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Conditional Access Authentication Context Class Reference ID
func (id IdentityConditionalAccessAuthenticationContextClassReferenceId) ID() string {
	fmtString := "/identity/conditionalAccess/authenticationContextClassReferences/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationContextClassReferenceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Conditional Access Authentication Context Class Reference ID
func (id IdentityConditionalAccessAuthenticationContextClassReferenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("conditionalAccess", "conditionalAccess", "conditionalAccess"),
		resourceids.StaticSegment("authenticationContextClassReferences", "authenticationContextClassReferences", "authenticationContextClassReferences"),
		resourceids.UserSpecifiedSegment("authenticationContextClassReferenceId", "authenticationContextClassReferenceId"),
	}
}

// String returns a human-readable description of this Identity Conditional Access Authentication Context Class Reference ID
func (id IdentityConditionalAccessAuthenticationContextClassReferenceId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Context Class Reference: %q", id.AuthenticationContextClassReferenceId),
	}
	return fmt.Sprintf("Identity Conditional Access Authentication Context Class Reference (%s)", strings.Join(components, "\n"))
}
