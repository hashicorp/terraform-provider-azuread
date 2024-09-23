package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityCustomAuthenticationExtensionId{}

// IdentityCustomAuthenticationExtensionId is a struct representing the Resource ID for a Identity Custom Authentication Extension
type IdentityCustomAuthenticationExtensionId struct {
	CustomAuthenticationExtensionId string
}

// NewIdentityCustomAuthenticationExtensionID returns a new IdentityCustomAuthenticationExtensionId struct
func NewIdentityCustomAuthenticationExtensionID(customAuthenticationExtensionId string) IdentityCustomAuthenticationExtensionId {
	return IdentityCustomAuthenticationExtensionId{
		CustomAuthenticationExtensionId: customAuthenticationExtensionId,
	}
}

// ParseIdentityCustomAuthenticationExtensionID parses 'input' into a IdentityCustomAuthenticationExtensionId
func ParseIdentityCustomAuthenticationExtensionID(input string) (*IdentityCustomAuthenticationExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityCustomAuthenticationExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityCustomAuthenticationExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityCustomAuthenticationExtensionIDInsensitively parses 'input' case-insensitively into a IdentityCustomAuthenticationExtensionId
// note: this method should only be used for API response data and not user input
func ParseIdentityCustomAuthenticationExtensionIDInsensitively(input string) (*IdentityCustomAuthenticationExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityCustomAuthenticationExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityCustomAuthenticationExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityCustomAuthenticationExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CustomAuthenticationExtensionId, ok = input.Parsed["customAuthenticationExtensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customAuthenticationExtensionId", input)
	}

	return nil
}

// ValidateIdentityCustomAuthenticationExtensionID checks that 'input' can be parsed as a Identity Custom Authentication Extension ID
func ValidateIdentityCustomAuthenticationExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityCustomAuthenticationExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Custom Authentication Extension ID
func (id IdentityCustomAuthenticationExtensionId) ID() string {
	fmtString := "/identity/customAuthenticationExtensions/%s"
	return fmt.Sprintf(fmtString, id.CustomAuthenticationExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Custom Authentication Extension ID
func (id IdentityCustomAuthenticationExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("customAuthenticationExtensions", "customAuthenticationExtensions", "customAuthenticationExtensions"),
		resourceids.UserSpecifiedSegment("customAuthenticationExtensionId", "customAuthenticationExtensionId"),
	}
}

// String returns a human-readable description of this Identity Custom Authentication Extension ID
func (id IdentityCustomAuthenticationExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Custom Authentication Extension: %q", id.CustomAuthenticationExtensionId),
	}
	return fmt.Sprintf("Identity Custom Authentication Extension (%s)", strings.Join(components, "\n"))
}
