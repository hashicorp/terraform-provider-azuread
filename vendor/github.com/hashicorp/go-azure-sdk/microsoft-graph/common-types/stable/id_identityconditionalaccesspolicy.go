package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccessPolicyId{}

// IdentityConditionalAccessPolicyId is a struct representing the Resource ID for a Identity Conditional Access Policy
type IdentityConditionalAccessPolicyId struct {
	ConditionalAccessPolicyId string
}

// NewIdentityConditionalAccessPolicyID returns a new IdentityConditionalAccessPolicyId struct
func NewIdentityConditionalAccessPolicyID(conditionalAccessPolicyId string) IdentityConditionalAccessPolicyId {
	return IdentityConditionalAccessPolicyId{
		ConditionalAccessPolicyId: conditionalAccessPolicyId,
	}
}

// ParseIdentityConditionalAccessPolicyID parses 'input' into a IdentityConditionalAccessPolicyId
func ParseIdentityConditionalAccessPolicyID(input string) (*IdentityConditionalAccessPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccessPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccessPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityConditionalAccessPolicyIDInsensitively parses 'input' case-insensitively into a IdentityConditionalAccessPolicyId
// note: this method should only be used for API response data and not user input
func ParseIdentityConditionalAccessPolicyIDInsensitively(input string) (*IdentityConditionalAccessPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccessPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccessPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityConditionalAccessPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ConditionalAccessPolicyId, ok = input.Parsed["conditionalAccessPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conditionalAccessPolicyId", input)
	}

	return nil
}

// ValidateIdentityConditionalAccessPolicyID checks that 'input' can be parsed as a Identity Conditional Access Policy ID
func ValidateIdentityConditionalAccessPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityConditionalAccessPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Conditional Access Policy ID
func (id IdentityConditionalAccessPolicyId) ID() string {
	fmtString := "/identity/conditionalAccess/policies/%s"
	return fmt.Sprintf(fmtString, id.ConditionalAccessPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Conditional Access Policy ID
func (id IdentityConditionalAccessPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("conditionalAccess", "conditionalAccess", "conditionalAccess"),
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.UserSpecifiedSegment("conditionalAccessPolicyId", "conditionalAccessPolicyId"),
	}
}

// String returns a human-readable description of this Identity Conditional Access Policy ID
func (id IdentityConditionalAccessPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Conditional Access Policy: %q", id.ConditionalAccessPolicyId),
	}
	return fmt.Sprintf("Identity Conditional Access Policy (%s)", strings.Join(components, "\n"))
}
