package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationMicrosoftAuthenticatorMethodId{}

// MeAuthenticationMicrosoftAuthenticatorMethodId is a struct representing the Resource ID for a Me Authentication Microsoft Authenticator Method
type MeAuthenticationMicrosoftAuthenticatorMethodId struct {
	MicrosoftAuthenticatorAuthenticationMethodId string
}

// NewMeAuthenticationMicrosoftAuthenticatorMethodID returns a new MeAuthenticationMicrosoftAuthenticatorMethodId struct
func NewMeAuthenticationMicrosoftAuthenticatorMethodID(microsoftAuthenticatorAuthenticationMethodId string) MeAuthenticationMicrosoftAuthenticatorMethodId {
	return MeAuthenticationMicrosoftAuthenticatorMethodId{
		MicrosoftAuthenticatorAuthenticationMethodId: microsoftAuthenticatorAuthenticationMethodId,
	}
}

// ParseMeAuthenticationMicrosoftAuthenticatorMethodID parses 'input' into a MeAuthenticationMicrosoftAuthenticatorMethodId
func ParseMeAuthenticationMicrosoftAuthenticatorMethodID(input string) (*MeAuthenticationMicrosoftAuthenticatorMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationMicrosoftAuthenticatorMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationMicrosoftAuthenticatorMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAuthenticationMicrosoftAuthenticatorMethodIDInsensitively parses 'input' case-insensitively into a MeAuthenticationMicrosoftAuthenticatorMethodId
// note: this method should only be used for API response data and not user input
func ParseMeAuthenticationMicrosoftAuthenticatorMethodIDInsensitively(input string) (*MeAuthenticationMicrosoftAuthenticatorMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationMicrosoftAuthenticatorMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationMicrosoftAuthenticatorMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAuthenticationMicrosoftAuthenticatorMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MicrosoftAuthenticatorAuthenticationMethodId, ok = input.Parsed["microsoftAuthenticatorAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "microsoftAuthenticatorAuthenticationMethodId", input)
	}

	return nil
}

// ValidateMeAuthenticationMicrosoftAuthenticatorMethodID checks that 'input' can be parsed as a Me Authentication Microsoft Authenticator Method ID
func ValidateMeAuthenticationMicrosoftAuthenticatorMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAuthenticationMicrosoftAuthenticatorMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Authentication Microsoft Authenticator Method ID
func (id MeAuthenticationMicrosoftAuthenticatorMethodId) ID() string {
	fmtString := "/me/authentication/microsoftAuthenticatorMethods/%s"
	return fmt.Sprintf(fmtString, id.MicrosoftAuthenticatorAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Authentication Microsoft Authenticator Method ID
func (id MeAuthenticationMicrosoftAuthenticatorMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("microsoftAuthenticatorMethods", "microsoftAuthenticatorMethods", "microsoftAuthenticatorMethods"),
		resourceids.UserSpecifiedSegment("microsoftAuthenticatorAuthenticationMethodId", "microsoftAuthenticatorAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this Me Authentication Microsoft Authenticator Method ID
func (id MeAuthenticationMicrosoftAuthenticatorMethodId) String() string {
	components := []string{
		fmt.Sprintf("Microsoft Authenticator Authentication Method: %q", id.MicrosoftAuthenticatorAuthenticationMethodId),
	}
	return fmt.Sprintf("Me Authentication Microsoft Authenticator Method (%s)", strings.Join(components, "\n"))
}
