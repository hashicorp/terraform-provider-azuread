package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2xUserFlowIdIdentityProviderId{}

// IdentityB2xUserFlowIdIdentityProviderId is a struct representing the Resource ID for a Identity B 2 x User Flow Id Identity Provider
type IdentityB2xUserFlowIdIdentityProviderId struct {
	B2xIdentityUserFlowId string
	IdentityProviderId    string
}

// NewIdentityB2xUserFlowIdIdentityProviderID returns a new IdentityB2xUserFlowIdIdentityProviderId struct
func NewIdentityB2xUserFlowIdIdentityProviderID(b2xIdentityUserFlowId string, identityProviderId string) IdentityB2xUserFlowIdIdentityProviderId {
	return IdentityB2xUserFlowIdIdentityProviderId{
		B2xIdentityUserFlowId: b2xIdentityUserFlowId,
		IdentityProviderId:    identityProviderId,
	}
}

// ParseIdentityB2xUserFlowIdIdentityProviderID parses 'input' into a IdentityB2xUserFlowIdIdentityProviderId
func ParseIdentityB2xUserFlowIdIdentityProviderID(input string) (*IdentityB2xUserFlowIdIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2xUserFlowIdIdentityProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2xUserFlowIdIdentityProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityB2xUserFlowIdIdentityProviderIDInsensitively parses 'input' case-insensitively into a IdentityB2xUserFlowIdIdentityProviderId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2xUserFlowIdIdentityProviderIDInsensitively(input string) (*IdentityB2xUserFlowIdIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2xUserFlowIdIdentityProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2xUserFlowIdIdentityProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityB2xUserFlowIdIdentityProviderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.B2xIdentityUserFlowId, ok = input.Parsed["b2xIdentityUserFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", input)
	}

	if id.IdentityProviderId, ok = input.Parsed["identityProviderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "identityProviderId", input)
	}

	return nil
}

// ValidateIdentityB2xUserFlowIdIdentityProviderID checks that 'input' can be parsed as a Identity B 2 x User Flow Id Identity Provider ID
func ValidateIdentityB2xUserFlowIdIdentityProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2xUserFlowIdIdentityProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2 x User Flow Id Identity Provider ID
func (id IdentityB2xUserFlowIdIdentityProviderId) ID() string {
	fmtString := "/identity/b2xUserFlows/%s/identityProviders/%s"
	return fmt.Sprintf(fmtString, id.B2xIdentityUserFlowId, id.IdentityProviderId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2 x User Flow Id Identity Provider ID
func (id IdentityB2xUserFlowIdIdentityProviderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("b2xUserFlows", "b2xUserFlows", "b2xUserFlows"),
		resourceids.UserSpecifiedSegment("b2xIdentityUserFlowId", "b2xIdentityUserFlowId"),
		resourceids.StaticSegment("identityProviders", "identityProviders", "identityProviders"),
		resourceids.UserSpecifiedSegment("identityProviderId", "identityProviderId"),
	}
}

// String returns a human-readable description of this Identity B 2 x User Flow Id Identity Provider ID
func (id IdentityB2xUserFlowIdIdentityProviderId) String() string {
	components := []string{
		fmt.Sprintf("B 2 x Identity User Flow: %q", id.B2xIdentityUserFlowId),
		fmt.Sprintf("Identity Provider: %q", id.IdentityProviderId),
	}
	return fmt.Sprintf("Identity B 2 x User Flow Id Identity Provider (%s)", strings.Join(components, "\n"))
}
