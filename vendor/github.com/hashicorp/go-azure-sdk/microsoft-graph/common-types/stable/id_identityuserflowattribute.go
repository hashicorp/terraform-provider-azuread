package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityUserFlowAttributeId{}

// IdentityUserFlowAttributeId is a struct representing the Resource ID for a Identity User Flow Attribute
type IdentityUserFlowAttributeId struct {
	IdentityUserFlowAttributeId string
}

// NewIdentityUserFlowAttributeID returns a new IdentityUserFlowAttributeId struct
func NewIdentityUserFlowAttributeID(identityUserFlowAttributeId string) IdentityUserFlowAttributeId {
	return IdentityUserFlowAttributeId{
		IdentityUserFlowAttributeId: identityUserFlowAttributeId,
	}
}

// ParseIdentityUserFlowAttributeID parses 'input' into a IdentityUserFlowAttributeId
func ParseIdentityUserFlowAttributeID(input string) (*IdentityUserFlowAttributeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityUserFlowAttributeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityUserFlowAttributeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityUserFlowAttributeIDInsensitively parses 'input' case-insensitively into a IdentityUserFlowAttributeId
// note: this method should only be used for API response data and not user input
func ParseIdentityUserFlowAttributeIDInsensitively(input string) (*IdentityUserFlowAttributeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityUserFlowAttributeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityUserFlowAttributeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityUserFlowAttributeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.IdentityUserFlowAttributeId, ok = input.Parsed["identityUserFlowAttributeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "identityUserFlowAttributeId", input)
	}

	return nil
}

// ValidateIdentityUserFlowAttributeID checks that 'input' can be parsed as a Identity User Flow Attribute ID
func ValidateIdentityUserFlowAttributeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityUserFlowAttributeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity User Flow Attribute ID
func (id IdentityUserFlowAttributeId) ID() string {
	fmtString := "/identity/userFlowAttributes/%s"
	return fmt.Sprintf(fmtString, id.IdentityUserFlowAttributeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity User Flow Attribute ID
func (id IdentityUserFlowAttributeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("userFlowAttributes", "userFlowAttributes", "userFlowAttributes"),
		resourceids.UserSpecifiedSegment("identityUserFlowAttributeId", "identityUserFlowAttributeId"),
	}
}

// String returns a human-readable description of this Identity User Flow Attribute ID
func (id IdentityUserFlowAttributeId) String() string {
	components := []string{
		fmt.Sprintf("Identity User Flow Attribute: %q", id.IdentityUserFlowAttributeId),
	}
	return fmt.Sprintf("Identity User Flow Attribute (%s)", strings.Join(components, "\n"))
}
