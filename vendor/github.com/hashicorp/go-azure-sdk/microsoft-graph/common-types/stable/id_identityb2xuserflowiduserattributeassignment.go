package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2xUserFlowIdUserAttributeAssignmentId{}

// IdentityB2xUserFlowIdUserAttributeAssignmentId is a struct representing the Resource ID for a Identity B 2 x User Flow Id User Attribute Assignment
type IdentityB2xUserFlowIdUserAttributeAssignmentId struct {
	B2xIdentityUserFlowId                 string
	IdentityUserFlowAttributeAssignmentId string
}

// NewIdentityB2xUserFlowIdUserAttributeAssignmentID returns a new IdentityB2xUserFlowIdUserAttributeAssignmentId struct
func NewIdentityB2xUserFlowIdUserAttributeAssignmentID(b2xIdentityUserFlowId string, identityUserFlowAttributeAssignmentId string) IdentityB2xUserFlowIdUserAttributeAssignmentId {
	return IdentityB2xUserFlowIdUserAttributeAssignmentId{
		B2xIdentityUserFlowId:                 b2xIdentityUserFlowId,
		IdentityUserFlowAttributeAssignmentId: identityUserFlowAttributeAssignmentId,
	}
}

// ParseIdentityB2xUserFlowIdUserAttributeAssignmentID parses 'input' into a IdentityB2xUserFlowIdUserAttributeAssignmentId
func ParseIdentityB2xUserFlowIdUserAttributeAssignmentID(input string) (*IdentityB2xUserFlowIdUserAttributeAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2xUserFlowIdUserAttributeAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2xUserFlowIdUserAttributeAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityB2xUserFlowIdUserAttributeAssignmentIDInsensitively parses 'input' case-insensitively into a IdentityB2xUserFlowIdUserAttributeAssignmentId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2xUserFlowIdUserAttributeAssignmentIDInsensitively(input string) (*IdentityB2xUserFlowIdUserAttributeAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2xUserFlowIdUserAttributeAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2xUserFlowIdUserAttributeAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityB2xUserFlowIdUserAttributeAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.B2xIdentityUserFlowId, ok = input.Parsed["b2xIdentityUserFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", input)
	}

	if id.IdentityUserFlowAttributeAssignmentId, ok = input.Parsed["identityUserFlowAttributeAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "identityUserFlowAttributeAssignmentId", input)
	}

	return nil
}

// ValidateIdentityB2xUserFlowIdUserAttributeAssignmentID checks that 'input' can be parsed as a Identity B 2 x User Flow Id User Attribute Assignment ID
func ValidateIdentityB2xUserFlowIdUserAttributeAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2xUserFlowIdUserAttributeAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2 x User Flow Id User Attribute Assignment ID
func (id IdentityB2xUserFlowIdUserAttributeAssignmentId) ID() string {
	fmtString := "/identity/b2xUserFlows/%s/userAttributeAssignments/%s"
	return fmt.Sprintf(fmtString, id.B2xIdentityUserFlowId, id.IdentityUserFlowAttributeAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2 x User Flow Id User Attribute Assignment ID
func (id IdentityB2xUserFlowIdUserAttributeAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("b2xUserFlows", "b2xUserFlows", "b2xUserFlows"),
		resourceids.UserSpecifiedSegment("b2xIdentityUserFlowId", "b2xIdentityUserFlowId"),
		resourceids.StaticSegment("userAttributeAssignments", "userAttributeAssignments", "userAttributeAssignments"),
		resourceids.UserSpecifiedSegment("identityUserFlowAttributeAssignmentId", "identityUserFlowAttributeAssignmentId"),
	}
}

// String returns a human-readable description of this Identity B 2 x User Flow Id User Attribute Assignment ID
func (id IdentityB2xUserFlowIdUserAttributeAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("B 2 x Identity User Flow: %q", id.B2xIdentityUserFlowId),
		fmt.Sprintf("Identity User Flow Attribute Assignment: %q", id.IdentityUserFlowAttributeAssignmentId),
	}
	return fmt.Sprintf("Identity B 2 x User Flow Id User Attribute Assignment (%s)", strings.Join(components, "\n"))
}
