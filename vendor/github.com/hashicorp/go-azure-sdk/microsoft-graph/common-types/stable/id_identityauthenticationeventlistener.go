package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityAuthenticationEventListenerId{}

// IdentityAuthenticationEventListenerId is a struct representing the Resource ID for a Identity Authentication Event Listener
type IdentityAuthenticationEventListenerId struct {
	AuthenticationEventListenerId string
}

// NewIdentityAuthenticationEventListenerID returns a new IdentityAuthenticationEventListenerId struct
func NewIdentityAuthenticationEventListenerID(authenticationEventListenerId string) IdentityAuthenticationEventListenerId {
	return IdentityAuthenticationEventListenerId{
		AuthenticationEventListenerId: authenticationEventListenerId,
	}
}

// ParseIdentityAuthenticationEventListenerID parses 'input' into a IdentityAuthenticationEventListenerId
func ParseIdentityAuthenticationEventListenerID(input string) (*IdentityAuthenticationEventListenerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityAuthenticationEventListenerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityAuthenticationEventListenerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityAuthenticationEventListenerIDInsensitively parses 'input' case-insensitively into a IdentityAuthenticationEventListenerId
// note: this method should only be used for API response data and not user input
func ParseIdentityAuthenticationEventListenerIDInsensitively(input string) (*IdentityAuthenticationEventListenerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityAuthenticationEventListenerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityAuthenticationEventListenerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityAuthenticationEventListenerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationEventListenerId, ok = input.Parsed["authenticationEventListenerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationEventListenerId", input)
	}

	return nil
}

// ValidateIdentityAuthenticationEventListenerID checks that 'input' can be parsed as a Identity Authentication Event Listener ID
func ValidateIdentityAuthenticationEventListenerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityAuthenticationEventListenerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Authentication Event Listener ID
func (id IdentityAuthenticationEventListenerId) ID() string {
	fmtString := "/identity/authenticationEventListeners/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationEventListenerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Authentication Event Listener ID
func (id IdentityAuthenticationEventListenerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("authenticationEventListeners", "authenticationEventListeners", "authenticationEventListeners"),
		resourceids.UserSpecifiedSegment("authenticationEventListenerId", "authenticationEventListenerId"),
	}
}

// String returns a human-readable description of this Identity Authentication Event Listener ID
func (id IdentityAuthenticationEventListenerId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Event Listener: %q", id.AuthenticationEventListenerId),
	}
	return fmt.Sprintf("Identity Authentication Event Listener (%s)", strings.Join(components, "\n"))
}
