package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccessNamedLocationId{}

// IdentityConditionalAccessNamedLocationId is a struct representing the Resource ID for a Identity Conditional Access Named Location
type IdentityConditionalAccessNamedLocationId struct {
	NamedLocationId string
}

// NewIdentityConditionalAccessNamedLocationID returns a new IdentityConditionalAccessNamedLocationId struct
func NewIdentityConditionalAccessNamedLocationID(namedLocationId string) IdentityConditionalAccessNamedLocationId {
	return IdentityConditionalAccessNamedLocationId{
		NamedLocationId: namedLocationId,
	}
}

// ParseIdentityConditionalAccessNamedLocationID parses 'input' into a IdentityConditionalAccessNamedLocationId
func ParseIdentityConditionalAccessNamedLocationID(input string) (*IdentityConditionalAccessNamedLocationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccessNamedLocationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccessNamedLocationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityConditionalAccessNamedLocationIDInsensitively parses 'input' case-insensitively into a IdentityConditionalAccessNamedLocationId
// note: this method should only be used for API response data and not user input
func ParseIdentityConditionalAccessNamedLocationIDInsensitively(input string) (*IdentityConditionalAccessNamedLocationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccessNamedLocationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccessNamedLocationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityConditionalAccessNamedLocationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.NamedLocationId, ok = input.Parsed["namedLocationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "namedLocationId", input)
	}

	return nil
}

// ValidateIdentityConditionalAccessNamedLocationID checks that 'input' can be parsed as a Identity Conditional Access Named Location ID
func ValidateIdentityConditionalAccessNamedLocationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityConditionalAccessNamedLocationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Conditional Access Named Location ID
func (id IdentityConditionalAccessNamedLocationId) ID() string {
	fmtString := "/identity/conditionalAccess/namedLocations/%s"
	return fmt.Sprintf(fmtString, id.NamedLocationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Conditional Access Named Location ID
func (id IdentityConditionalAccessNamedLocationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("conditionalAccess", "conditionalAccess", "conditionalAccess"),
		resourceids.StaticSegment("namedLocations", "namedLocations", "namedLocations"),
		resourceids.UserSpecifiedSegment("namedLocationId", "namedLocationId"),
	}
}

// String returns a human-readable description of this Identity Conditional Access Named Location ID
func (id IdentityConditionalAccessNamedLocationId) String() string {
	components := []string{
		fmt.Sprintf("Named Location: %q", id.NamedLocationId),
	}
	return fmt.Sprintf("Identity Conditional Access Named Location (%s)", strings.Join(components, "\n"))
}
