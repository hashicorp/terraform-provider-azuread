package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationFido2MethodId{}

// MeAuthenticationFido2MethodId is a struct representing the Resource ID for a Me Authentication Fido 2 Method
type MeAuthenticationFido2MethodId struct {
	Fido2AuthenticationMethodId string
}

// NewMeAuthenticationFido2MethodID returns a new MeAuthenticationFido2MethodId struct
func NewMeAuthenticationFido2MethodID(fido2AuthenticationMethodId string) MeAuthenticationFido2MethodId {
	return MeAuthenticationFido2MethodId{
		Fido2AuthenticationMethodId: fido2AuthenticationMethodId,
	}
}

// ParseMeAuthenticationFido2MethodID parses 'input' into a MeAuthenticationFido2MethodId
func ParseMeAuthenticationFido2MethodID(input string) (*MeAuthenticationFido2MethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationFido2MethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationFido2MethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAuthenticationFido2MethodIDInsensitively parses 'input' case-insensitively into a MeAuthenticationFido2MethodId
// note: this method should only be used for API response data and not user input
func ParseMeAuthenticationFido2MethodIDInsensitively(input string) (*MeAuthenticationFido2MethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationFido2MethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationFido2MethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAuthenticationFido2MethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Fido2AuthenticationMethodId, ok = input.Parsed["fido2AuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "fido2AuthenticationMethodId", input)
	}

	return nil
}

// ValidateMeAuthenticationFido2MethodID checks that 'input' can be parsed as a Me Authentication Fido 2 Method ID
func ValidateMeAuthenticationFido2MethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAuthenticationFido2MethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Authentication Fido 2 Method ID
func (id MeAuthenticationFido2MethodId) ID() string {
	fmtString := "/me/authentication/fido2Methods/%s"
	return fmt.Sprintf(fmtString, id.Fido2AuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Authentication Fido 2 Method ID
func (id MeAuthenticationFido2MethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("fido2Methods", "fido2Methods", "fido2Methods"),
		resourceids.UserSpecifiedSegment("fido2AuthenticationMethodId", "fido2AuthenticationMethodId"),
	}
}

// String returns a human-readable description of this Me Authentication Fido 2 Method ID
func (id MeAuthenticationFido2MethodId) String() string {
	components := []string{
		fmt.Sprintf("Fido 2 Authentication Method: %q", id.Fido2AuthenticationMethodId),
	}
	return fmt.Sprintf("Me Authentication Fido 2 Method (%s)", strings.Join(components, "\n"))
}
