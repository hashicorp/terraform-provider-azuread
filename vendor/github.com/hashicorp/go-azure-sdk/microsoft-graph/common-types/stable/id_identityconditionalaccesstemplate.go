package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccessTemplateId{}

// IdentityConditionalAccessTemplateId is a struct representing the Resource ID for a Identity Conditional Access Template
type IdentityConditionalAccessTemplateId struct {
	ConditionalAccessTemplateId string
}

// NewIdentityConditionalAccessTemplateID returns a new IdentityConditionalAccessTemplateId struct
func NewIdentityConditionalAccessTemplateID(conditionalAccessTemplateId string) IdentityConditionalAccessTemplateId {
	return IdentityConditionalAccessTemplateId{
		ConditionalAccessTemplateId: conditionalAccessTemplateId,
	}
}

// ParseIdentityConditionalAccessTemplateID parses 'input' into a IdentityConditionalAccessTemplateId
func ParseIdentityConditionalAccessTemplateID(input string) (*IdentityConditionalAccessTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccessTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccessTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityConditionalAccessTemplateIDInsensitively parses 'input' case-insensitively into a IdentityConditionalAccessTemplateId
// note: this method should only be used for API response data and not user input
func ParseIdentityConditionalAccessTemplateIDInsensitively(input string) (*IdentityConditionalAccessTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccessTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccessTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityConditionalAccessTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ConditionalAccessTemplateId, ok = input.Parsed["conditionalAccessTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conditionalAccessTemplateId", input)
	}

	return nil
}

// ValidateIdentityConditionalAccessTemplateID checks that 'input' can be parsed as a Identity Conditional Access Template ID
func ValidateIdentityConditionalAccessTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityConditionalAccessTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Conditional Access Template ID
func (id IdentityConditionalAccessTemplateId) ID() string {
	fmtString := "/identity/conditionalAccess/templates/%s"
	return fmt.Sprintf(fmtString, id.ConditionalAccessTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Conditional Access Template ID
func (id IdentityConditionalAccessTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("conditionalAccess", "conditionalAccess", "conditionalAccess"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("conditionalAccessTemplateId", "conditionalAccessTemplateId"),
	}
}

// String returns a human-readable description of this Identity Conditional Access Template ID
func (id IdentityConditionalAccessTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Conditional Access Template: %q", id.ConditionalAccessTemplateId),
	}
	return fmt.Sprintf("Identity Conditional Access Template (%s)", strings.Join(components, "\n"))
}
