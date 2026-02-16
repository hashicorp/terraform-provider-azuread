package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationMethodId{}

// MeAuthenticationMethodId is a struct representing the Resource ID for a Me Authentication Method
type MeAuthenticationMethodId struct {
	AuthenticationMethodId string
}

// NewMeAuthenticationMethodID returns a new MeAuthenticationMethodId struct
func NewMeAuthenticationMethodID(authenticationMethodId string) MeAuthenticationMethodId {
	return MeAuthenticationMethodId{
		AuthenticationMethodId: authenticationMethodId,
	}
}

// ParseMeAuthenticationMethodID parses 'input' into a MeAuthenticationMethodId
func ParseMeAuthenticationMethodID(input string) (*MeAuthenticationMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAuthenticationMethodIDInsensitively parses 'input' case-insensitively into a MeAuthenticationMethodId
// note: this method should only be used for API response data and not user input
func ParseMeAuthenticationMethodIDInsensitively(input string) (*MeAuthenticationMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAuthenticationMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationMethodId, ok = input.Parsed["authenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationMethodId", input)
	}

	return nil
}

// ValidateMeAuthenticationMethodID checks that 'input' can be parsed as a Me Authentication Method ID
func ValidateMeAuthenticationMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAuthenticationMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Authentication Method ID
func (id MeAuthenticationMethodId) ID() string {
	fmtString := "/me/authentication/methods/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Authentication Method ID
func (id MeAuthenticationMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("methods", "methods", "methods"),
		resourceids.UserSpecifiedSegment("authenticationMethodId", "authenticationMethodId"),
	}
}

// String returns a human-readable description of this Me Authentication Method ID
func (id MeAuthenticationMethodId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Method: %q", id.AuthenticationMethodId),
	}
	return fmt.Sprintf("Me Authentication Method (%s)", strings.Join(components, "\n"))
}
