package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2xUserFlowIdLanguageIdOverridesPageId{}

// IdentityB2xUserFlowIdLanguageIdOverridesPageId is a struct representing the Resource ID for a Identity B 2x User Flow Id Language Id Overrides Page
type IdentityB2xUserFlowIdLanguageIdOverridesPageId struct {
	B2xIdentityUserFlowId           string
	UserFlowLanguageConfigurationId string
	UserFlowLanguagePageId          string
}

// NewIdentityB2xUserFlowIdLanguageIdOverridesPageID returns a new IdentityB2xUserFlowIdLanguageIdOverridesPageId struct
func NewIdentityB2xUserFlowIdLanguageIdOverridesPageID(b2xIdentityUserFlowId string, userFlowLanguageConfigurationId string, userFlowLanguagePageId string) IdentityB2xUserFlowIdLanguageIdOverridesPageId {
	return IdentityB2xUserFlowIdLanguageIdOverridesPageId{
		B2xIdentityUserFlowId:           b2xIdentityUserFlowId,
		UserFlowLanguageConfigurationId: userFlowLanguageConfigurationId,
		UserFlowLanguagePageId:          userFlowLanguagePageId,
	}
}

// ParseIdentityB2xUserFlowIdLanguageIdOverridesPageID parses 'input' into a IdentityB2xUserFlowIdLanguageIdOverridesPageId
func ParseIdentityB2xUserFlowIdLanguageIdOverridesPageID(input string) (*IdentityB2xUserFlowIdLanguageIdOverridesPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2xUserFlowIdLanguageIdOverridesPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2xUserFlowIdLanguageIdOverridesPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityB2xUserFlowIdLanguageIdOverridesPageIDInsensitively parses 'input' case-insensitively into a IdentityB2xUserFlowIdLanguageIdOverridesPageId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2xUserFlowIdLanguageIdOverridesPageIDInsensitively(input string) (*IdentityB2xUserFlowIdLanguageIdOverridesPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2xUserFlowIdLanguageIdOverridesPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2xUserFlowIdLanguageIdOverridesPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityB2xUserFlowIdLanguageIdOverridesPageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.B2xIdentityUserFlowId, ok = input.Parsed["b2xIdentityUserFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", input)
	}

	if id.UserFlowLanguageConfigurationId, ok = input.Parsed["userFlowLanguageConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userFlowLanguageConfigurationId", input)
	}

	if id.UserFlowLanguagePageId, ok = input.Parsed["userFlowLanguagePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userFlowLanguagePageId", input)
	}

	return nil
}

// ValidateIdentityB2xUserFlowIdLanguageIdOverridesPageID checks that 'input' can be parsed as a Identity B 2x User Flow Id Language Id Overrides Page ID
func ValidateIdentityB2xUserFlowIdLanguageIdOverridesPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2xUserFlowIdLanguageIdOverridesPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2x User Flow Id Language Id Overrides Page ID
func (id IdentityB2xUserFlowIdLanguageIdOverridesPageId) ID() string {
	fmtString := "/identity/b2xUserFlows/%s/languages/%s/overridesPages/%s"
	return fmt.Sprintf(fmtString, id.B2xIdentityUserFlowId, id.UserFlowLanguageConfigurationId, id.UserFlowLanguagePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2x User Flow Id Language Id Overrides Page ID
func (id IdentityB2xUserFlowIdLanguageIdOverridesPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("b2xUserFlows", "b2xUserFlows", "b2xUserFlows"),
		resourceids.UserSpecifiedSegment("b2xIdentityUserFlowId", "b2xIdentityUserFlowId"),
		resourceids.StaticSegment("languages", "languages", "languages"),
		resourceids.UserSpecifiedSegment("userFlowLanguageConfigurationId", "userFlowLanguageConfigurationId"),
		resourceids.StaticSegment("overridesPages", "overridesPages", "overridesPages"),
		resourceids.UserSpecifiedSegment("userFlowLanguagePageId", "userFlowLanguagePageId"),
	}
}

// String returns a human-readable description of this Identity B 2x User Flow Id Language Id Overrides Page ID
func (id IdentityB2xUserFlowIdLanguageIdOverridesPageId) String() string {
	components := []string{
		fmt.Sprintf("B 2x Identity User Flow: %q", id.B2xIdentityUserFlowId),
		fmt.Sprintf("User Flow Language Configuration: %q", id.UserFlowLanguageConfigurationId),
		fmt.Sprintf("User Flow Language Page: %q", id.UserFlowLanguagePageId),
	}
	return fmt.Sprintf("Identity B 2x User Flow Id Language Id Overrides Page (%s)", strings.Join(components, "\n"))
}
