package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationPhoneMethodId{}

// MeAuthenticationPhoneMethodId is a struct representing the Resource ID for a Me Authentication Phone Method
type MeAuthenticationPhoneMethodId struct {
	PhoneAuthenticationMethodId string
}

// NewMeAuthenticationPhoneMethodID returns a new MeAuthenticationPhoneMethodId struct
func NewMeAuthenticationPhoneMethodID(phoneAuthenticationMethodId string) MeAuthenticationPhoneMethodId {
	return MeAuthenticationPhoneMethodId{
		PhoneAuthenticationMethodId: phoneAuthenticationMethodId,
	}
}

// ParseMeAuthenticationPhoneMethodID parses 'input' into a MeAuthenticationPhoneMethodId
func ParseMeAuthenticationPhoneMethodID(input string) (*MeAuthenticationPhoneMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationPhoneMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationPhoneMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAuthenticationPhoneMethodIDInsensitively parses 'input' case-insensitively into a MeAuthenticationPhoneMethodId
// note: this method should only be used for API response data and not user input
func ParseMeAuthenticationPhoneMethodIDInsensitively(input string) (*MeAuthenticationPhoneMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationPhoneMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationPhoneMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAuthenticationPhoneMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PhoneAuthenticationMethodId, ok = input.Parsed["phoneAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "phoneAuthenticationMethodId", input)
	}

	return nil
}

// ValidateMeAuthenticationPhoneMethodID checks that 'input' can be parsed as a Me Authentication Phone Method ID
func ValidateMeAuthenticationPhoneMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAuthenticationPhoneMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Authentication Phone Method ID
func (id MeAuthenticationPhoneMethodId) ID() string {
	fmtString := "/me/authentication/phoneMethods/%s"
	return fmt.Sprintf(fmtString, id.PhoneAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Authentication Phone Method ID
func (id MeAuthenticationPhoneMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("phoneMethods", "phoneMethods", "phoneMethods"),
		resourceids.UserSpecifiedSegment("phoneAuthenticationMethodId", "phoneAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this Me Authentication Phone Method ID
func (id MeAuthenticationPhoneMethodId) String() string {
	components := []string{
		fmt.Sprintf("Phone Authentication Method: %q", id.PhoneAuthenticationMethodId),
	}
	return fmt.Sprintf("Me Authentication Phone Method (%s)", strings.Join(components, "\n"))
}
