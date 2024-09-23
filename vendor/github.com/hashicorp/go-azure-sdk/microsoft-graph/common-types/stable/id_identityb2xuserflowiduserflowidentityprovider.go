package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2xUserFlowIdUserFlowIdentityProviderId{}

// IdentityB2xUserFlowIdUserFlowIdentityProviderId is a struct representing the Resource ID for a Identity B 2x User Flow Id User Flow Identity Provider
type IdentityB2xUserFlowIdUserFlowIdentityProviderId struct {
	B2xIdentityUserFlowId  string
	IdentityProviderBaseId string
}

// NewIdentityB2xUserFlowIdUserFlowIdentityProviderID returns a new IdentityB2xUserFlowIdUserFlowIdentityProviderId struct
func NewIdentityB2xUserFlowIdUserFlowIdentityProviderID(b2xIdentityUserFlowId string, identityProviderBaseId string) IdentityB2xUserFlowIdUserFlowIdentityProviderId {
	return IdentityB2xUserFlowIdUserFlowIdentityProviderId{
		B2xIdentityUserFlowId:  b2xIdentityUserFlowId,
		IdentityProviderBaseId: identityProviderBaseId,
	}
}

// ParseIdentityB2xUserFlowIdUserFlowIdentityProviderID parses 'input' into a IdentityB2xUserFlowIdUserFlowIdentityProviderId
func ParseIdentityB2xUserFlowIdUserFlowIdentityProviderID(input string) (*IdentityB2xUserFlowIdUserFlowIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2xUserFlowIdUserFlowIdentityProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2xUserFlowIdUserFlowIdentityProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityB2xUserFlowIdUserFlowIdentityProviderIDInsensitively parses 'input' case-insensitively into a IdentityB2xUserFlowIdUserFlowIdentityProviderId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2xUserFlowIdUserFlowIdentityProviderIDInsensitively(input string) (*IdentityB2xUserFlowIdUserFlowIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2xUserFlowIdUserFlowIdentityProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2xUserFlowIdUserFlowIdentityProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityB2xUserFlowIdUserFlowIdentityProviderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.B2xIdentityUserFlowId, ok = input.Parsed["b2xIdentityUserFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", input)
	}

	if id.IdentityProviderBaseId, ok = input.Parsed["identityProviderBaseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "identityProviderBaseId", input)
	}

	return nil
}

// ValidateIdentityB2xUserFlowIdUserFlowIdentityProviderID checks that 'input' can be parsed as a Identity B 2x User Flow Id User Flow Identity Provider ID
func ValidateIdentityB2xUserFlowIdUserFlowIdentityProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2xUserFlowIdUserFlowIdentityProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2x User Flow Id User Flow Identity Provider ID
func (id IdentityB2xUserFlowIdUserFlowIdentityProviderId) ID() string {
	fmtString := "/identity/b2xUserFlows/%s/userFlowIdentityProviders/%s"
	return fmt.Sprintf(fmtString, id.B2xIdentityUserFlowId, id.IdentityProviderBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2x User Flow Id User Flow Identity Provider ID
func (id IdentityB2xUserFlowIdUserFlowIdentityProviderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("b2xUserFlows", "b2xUserFlows", "b2xUserFlows"),
		resourceids.UserSpecifiedSegment("b2xIdentityUserFlowId", "b2xIdentityUserFlowId"),
		resourceids.StaticSegment("userFlowIdentityProviders", "userFlowIdentityProviders", "userFlowIdentityProviders"),
		resourceids.UserSpecifiedSegment("identityProviderBaseId", "identityProviderBaseId"),
	}
}

// String returns a human-readable description of this Identity B 2x User Flow Id User Flow Identity Provider ID
func (id IdentityB2xUserFlowIdUserFlowIdentityProviderId) String() string {
	components := []string{
		fmt.Sprintf("B 2x Identity User Flow: %q", id.B2xIdentityUserFlowId),
		fmt.Sprintf("Identity Provider Base: %q", id.IdentityProviderBaseId),
	}
	return fmt.Sprintf("Identity B 2x User Flow Id User Flow Identity Provider (%s)", strings.Join(components, "\n"))
}
