package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2xUserFlowIdLanguageId{}

// IdentityB2xUserFlowIdLanguageId is a struct representing the Resource ID for a Identity B 2 x User Flow Id Language
type IdentityB2xUserFlowIdLanguageId struct {
	B2xIdentityUserFlowId           string
	UserFlowLanguageConfigurationId string
}

// NewIdentityB2xUserFlowIdLanguageID returns a new IdentityB2xUserFlowIdLanguageId struct
func NewIdentityB2xUserFlowIdLanguageID(b2xIdentityUserFlowId string, userFlowLanguageConfigurationId string) IdentityB2xUserFlowIdLanguageId {
	return IdentityB2xUserFlowIdLanguageId{
		B2xIdentityUserFlowId:           b2xIdentityUserFlowId,
		UserFlowLanguageConfigurationId: userFlowLanguageConfigurationId,
	}
}

// ParseIdentityB2xUserFlowIdLanguageID parses 'input' into a IdentityB2xUserFlowIdLanguageId
func ParseIdentityB2xUserFlowIdLanguageID(input string) (*IdentityB2xUserFlowIdLanguageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2xUserFlowIdLanguageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2xUserFlowIdLanguageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityB2xUserFlowIdLanguageIDInsensitively parses 'input' case-insensitively into a IdentityB2xUserFlowIdLanguageId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2xUserFlowIdLanguageIDInsensitively(input string) (*IdentityB2xUserFlowIdLanguageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2xUserFlowIdLanguageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2xUserFlowIdLanguageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityB2xUserFlowIdLanguageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.B2xIdentityUserFlowId, ok = input.Parsed["b2xIdentityUserFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", input)
	}

	if id.UserFlowLanguageConfigurationId, ok = input.Parsed["userFlowLanguageConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userFlowLanguageConfigurationId", input)
	}

	return nil
}

// ValidateIdentityB2xUserFlowIdLanguageID checks that 'input' can be parsed as a Identity B 2 x User Flow Id Language ID
func ValidateIdentityB2xUserFlowIdLanguageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2xUserFlowIdLanguageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2 x User Flow Id Language ID
func (id IdentityB2xUserFlowIdLanguageId) ID() string {
	fmtString := "/identity/b2xUserFlows/%s/languages/%s"
	return fmt.Sprintf(fmtString, id.B2xIdentityUserFlowId, id.UserFlowLanguageConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2 x User Flow Id Language ID
func (id IdentityB2xUserFlowIdLanguageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("b2xUserFlows", "b2xUserFlows", "b2xUserFlows"),
		resourceids.UserSpecifiedSegment("b2xIdentityUserFlowId", "b2xIdentityUserFlowId"),
		resourceids.StaticSegment("languages", "languages", "languages"),
		resourceids.UserSpecifiedSegment("userFlowLanguageConfigurationId", "userFlowLanguageConfigurationId"),
	}
}

// String returns a human-readable description of this Identity B 2 x User Flow Id Language ID
func (id IdentityB2xUserFlowIdLanguageId) String() string {
	components := []string{
		fmt.Sprintf("B 2 x Identity User Flow: %q", id.B2xIdentityUserFlowId),
		fmt.Sprintf("User Flow Language Configuration: %q", id.UserFlowLanguageConfigurationId),
	}
	return fmt.Sprintf("Identity B 2 x User Flow Id Language (%s)", strings.Join(components, "\n"))
}
