package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeExtensionId{}

// MeExtensionId is a struct representing the Resource ID for a Me Extension
type MeExtensionId struct {
	ExtensionId string
}

// NewMeExtensionID returns a new MeExtensionId struct
func NewMeExtensionID(extensionId string) MeExtensionId {
	return MeExtensionId{
		ExtensionId: extensionId,
	}
}

// ParseMeExtensionID parses 'input' into a MeExtensionId
func ParseMeExtensionID(input string) (*MeExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeExtensionIDInsensitively parses 'input' case-insensitively into a MeExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeExtensionIDInsensitively(input string) (*MeExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeExtensionID checks that 'input' can be parsed as a Me Extension ID
func ValidateMeExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Extension ID
func (id MeExtensionId) ID() string {
	fmtString := "/me/extensions/%s"
	return fmt.Sprintf(fmtString, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Extension ID
func (id MeExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Me Extension ID
func (id MeExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Extension (%s)", strings.Join(components, "\n"))
}
