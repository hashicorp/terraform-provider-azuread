package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2xUserFlowId{}

// IdentityB2xUserFlowId is a struct representing the Resource ID for a Identity B 2 x User Flow
type IdentityB2xUserFlowId struct {
	B2xIdentityUserFlowId string
}

// NewIdentityB2xUserFlowID returns a new IdentityB2xUserFlowId struct
func NewIdentityB2xUserFlowID(b2xIdentityUserFlowId string) IdentityB2xUserFlowId {
	return IdentityB2xUserFlowId{
		B2xIdentityUserFlowId: b2xIdentityUserFlowId,
	}
}

// ParseIdentityB2xUserFlowID parses 'input' into a IdentityB2xUserFlowId
func ParseIdentityB2xUserFlowID(input string) (*IdentityB2xUserFlowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2xUserFlowId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2xUserFlowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityB2xUserFlowIDInsensitively parses 'input' case-insensitively into a IdentityB2xUserFlowId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2xUserFlowIDInsensitively(input string) (*IdentityB2xUserFlowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2xUserFlowId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2xUserFlowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityB2xUserFlowId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.B2xIdentityUserFlowId, ok = input.Parsed["b2xIdentityUserFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", input)
	}

	return nil
}

// ValidateIdentityB2xUserFlowID checks that 'input' can be parsed as a Identity B 2 x User Flow ID
func ValidateIdentityB2xUserFlowID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2xUserFlowID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2 x User Flow ID
func (id IdentityB2xUserFlowId) ID() string {
	fmtString := "/identity/b2xUserFlows/%s"
	return fmt.Sprintf(fmtString, id.B2xIdentityUserFlowId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2 x User Flow ID
func (id IdentityB2xUserFlowId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("b2xUserFlows", "b2xUserFlows", "b2xUserFlows"),
		resourceids.UserSpecifiedSegment("b2xIdentityUserFlowId", "b2xIdentityUserFlowId"),
	}
}

// String returns a human-readable description of this Identity B 2 x User Flow ID
func (id IdentityB2xUserFlowId) String() string {
	components := []string{
		fmt.Sprintf("B 2 x Identity User Flow: %q", id.B2xIdentityUserFlowId),
	}
	return fmt.Sprintf("Identity B 2 x User Flow (%s)", strings.Join(components, "\n"))
}
