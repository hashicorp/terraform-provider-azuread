package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityAuthenticationEventsFlowId{}

// IdentityAuthenticationEventsFlowId is a struct representing the Resource ID for a Identity Authentication Events Flow
type IdentityAuthenticationEventsFlowId struct {
	AuthenticationEventsFlowId string
}

// NewIdentityAuthenticationEventsFlowID returns a new IdentityAuthenticationEventsFlowId struct
func NewIdentityAuthenticationEventsFlowID(authenticationEventsFlowId string) IdentityAuthenticationEventsFlowId {
	return IdentityAuthenticationEventsFlowId{
		AuthenticationEventsFlowId: authenticationEventsFlowId,
	}
}

// ParseIdentityAuthenticationEventsFlowID parses 'input' into a IdentityAuthenticationEventsFlowId
func ParseIdentityAuthenticationEventsFlowID(input string) (*IdentityAuthenticationEventsFlowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityAuthenticationEventsFlowId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityAuthenticationEventsFlowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityAuthenticationEventsFlowIDInsensitively parses 'input' case-insensitively into a IdentityAuthenticationEventsFlowId
// note: this method should only be used for API response data and not user input
func ParseIdentityAuthenticationEventsFlowIDInsensitively(input string) (*IdentityAuthenticationEventsFlowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityAuthenticationEventsFlowId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityAuthenticationEventsFlowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityAuthenticationEventsFlowId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationEventsFlowId, ok = input.Parsed["authenticationEventsFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationEventsFlowId", input)
	}

	return nil
}

// ValidateIdentityAuthenticationEventsFlowID checks that 'input' can be parsed as a Identity Authentication Events Flow ID
func ValidateIdentityAuthenticationEventsFlowID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityAuthenticationEventsFlowID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Authentication Events Flow ID
func (id IdentityAuthenticationEventsFlowId) ID() string {
	fmtString := "/identity/authenticationEventsFlows/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationEventsFlowId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Authentication Events Flow ID
func (id IdentityAuthenticationEventsFlowId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("authenticationEventsFlows", "authenticationEventsFlows", "authenticationEventsFlows"),
		resourceids.UserSpecifiedSegment("authenticationEventsFlowId", "authenticationEventsFlowId"),
	}
}

// String returns a human-readable description of this Identity Authentication Events Flow ID
func (id IdentityAuthenticationEventsFlowId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Events Flow: %q", id.AuthenticationEventsFlowId),
	}
	return fmt.Sprintf("Identity Authentication Events Flow (%s)", strings.Join(components, "\n"))
}
