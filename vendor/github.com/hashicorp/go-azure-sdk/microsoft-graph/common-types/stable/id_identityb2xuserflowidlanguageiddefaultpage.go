package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2xUserFlowIdLanguageIdDefaultPageId{}

// IdentityB2xUserFlowIdLanguageIdDefaultPageId is a struct representing the Resource ID for a Identity B 2 x User Flow Id Language Id Default Page
type IdentityB2xUserFlowIdLanguageIdDefaultPageId struct {
	B2xIdentityUserFlowId           string
	UserFlowLanguageConfigurationId string
	UserFlowLanguagePageId          string
}

// NewIdentityB2xUserFlowIdLanguageIdDefaultPageID returns a new IdentityB2xUserFlowIdLanguageIdDefaultPageId struct
func NewIdentityB2xUserFlowIdLanguageIdDefaultPageID(b2xIdentityUserFlowId string, userFlowLanguageConfigurationId string, userFlowLanguagePageId string) IdentityB2xUserFlowIdLanguageIdDefaultPageId {
	return IdentityB2xUserFlowIdLanguageIdDefaultPageId{
		B2xIdentityUserFlowId:           b2xIdentityUserFlowId,
		UserFlowLanguageConfigurationId: userFlowLanguageConfigurationId,
		UserFlowLanguagePageId:          userFlowLanguagePageId,
	}
}

// ParseIdentityB2xUserFlowIdLanguageIdDefaultPageID parses 'input' into a IdentityB2xUserFlowIdLanguageIdDefaultPageId
func ParseIdentityB2xUserFlowIdLanguageIdDefaultPageID(input string) (*IdentityB2xUserFlowIdLanguageIdDefaultPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2xUserFlowIdLanguageIdDefaultPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2xUserFlowIdLanguageIdDefaultPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityB2xUserFlowIdLanguageIdDefaultPageIDInsensitively parses 'input' case-insensitively into a IdentityB2xUserFlowIdLanguageIdDefaultPageId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2xUserFlowIdLanguageIdDefaultPageIDInsensitively(input string) (*IdentityB2xUserFlowIdLanguageIdDefaultPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2xUserFlowIdLanguageIdDefaultPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2xUserFlowIdLanguageIdDefaultPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityB2xUserFlowIdLanguageIdDefaultPageId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateIdentityB2xUserFlowIdLanguageIdDefaultPageID checks that 'input' can be parsed as a Identity B 2 x User Flow Id Language Id Default Page ID
func ValidateIdentityB2xUserFlowIdLanguageIdDefaultPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2xUserFlowIdLanguageIdDefaultPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2 x User Flow Id Language Id Default Page ID
func (id IdentityB2xUserFlowIdLanguageIdDefaultPageId) ID() string {
	fmtString := "/identity/b2xUserFlows/%s/languages/%s/defaultPages/%s"
	return fmt.Sprintf(fmtString, id.B2xIdentityUserFlowId, id.UserFlowLanguageConfigurationId, id.UserFlowLanguagePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2 x User Flow Id Language Id Default Page ID
func (id IdentityB2xUserFlowIdLanguageIdDefaultPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("b2xUserFlows", "b2xUserFlows", "b2xUserFlows"),
		resourceids.UserSpecifiedSegment("b2xIdentityUserFlowId", "b2xIdentityUserFlowId"),
		resourceids.StaticSegment("languages", "languages", "languages"),
		resourceids.UserSpecifiedSegment("userFlowLanguageConfigurationId", "userFlowLanguageConfigurationId"),
		resourceids.StaticSegment("defaultPages", "defaultPages", "defaultPages"),
		resourceids.UserSpecifiedSegment("userFlowLanguagePageId", "userFlowLanguagePageId"),
	}
}

// String returns a human-readable description of this Identity B 2 x User Flow Id Language Id Default Page ID
func (id IdentityB2xUserFlowIdLanguageIdDefaultPageId) String() string {
	components := []string{
		fmt.Sprintf("B 2 x Identity User Flow: %q", id.B2xIdentityUserFlowId),
		fmt.Sprintf("User Flow Language Configuration: %q", id.UserFlowLanguageConfigurationId),
		fmt.Sprintf("User Flow Language Page: %q", id.UserFlowLanguagePageId),
	}
	return fmt.Sprintf("Identity B 2 x User Flow Id Language Id Default Page (%s)", strings.Join(components, "\n"))
}
