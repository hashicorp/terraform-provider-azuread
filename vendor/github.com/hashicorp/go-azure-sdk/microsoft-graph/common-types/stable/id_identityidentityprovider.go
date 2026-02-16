package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityIdentityProviderId{}

// IdentityIdentityProviderId is a struct representing the Resource ID for a Identity Identity Provider
type IdentityIdentityProviderId struct {
	IdentityProviderBaseId string
}

// NewIdentityIdentityProviderID returns a new IdentityIdentityProviderId struct
func NewIdentityIdentityProviderID(identityProviderBaseId string) IdentityIdentityProviderId {
	return IdentityIdentityProviderId{
		IdentityProviderBaseId: identityProviderBaseId,
	}
}

// ParseIdentityIdentityProviderID parses 'input' into a IdentityIdentityProviderId
func ParseIdentityIdentityProviderID(input string) (*IdentityIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityIdentityProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityIdentityProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityIdentityProviderIDInsensitively parses 'input' case-insensitively into a IdentityIdentityProviderId
// note: this method should only be used for API response data and not user input
func ParseIdentityIdentityProviderIDInsensitively(input string) (*IdentityIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityIdentityProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityIdentityProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityIdentityProviderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.IdentityProviderBaseId, ok = input.Parsed["identityProviderBaseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "identityProviderBaseId", input)
	}

	return nil
}

// ValidateIdentityIdentityProviderID checks that 'input' can be parsed as a Identity Identity Provider ID
func ValidateIdentityIdentityProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityIdentityProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Identity Provider ID
func (id IdentityIdentityProviderId) ID() string {
	fmtString := "/identity/identityProviders/%s"
	return fmt.Sprintf(fmtString, id.IdentityProviderBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Identity Provider ID
func (id IdentityIdentityProviderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("identityProviders", "identityProviders", "identityProviders"),
		resourceids.UserSpecifiedSegment("identityProviderBaseId", "identityProviderBaseId"),
	}
}

// String returns a human-readable description of this Identity Identity Provider ID
func (id IdentityIdentityProviderId) String() string {
	components := []string{
		fmt.Sprintf("Identity Provider Base: %q", id.IdentityProviderBaseId),
	}
	return fmt.Sprintf("Identity Identity Provider (%s)", strings.Join(components, "\n"))
}
