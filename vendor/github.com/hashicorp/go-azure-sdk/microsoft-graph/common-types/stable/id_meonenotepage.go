package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenotePageId{}

// MeOnenotePageId is a struct representing the Resource ID for a Me Onenote Page
type MeOnenotePageId struct {
	OnenotePageId string
}

// NewMeOnenotePageID returns a new MeOnenotePageId struct
func NewMeOnenotePageID(onenotePageId string) MeOnenotePageId {
	return MeOnenotePageId{
		OnenotePageId: onenotePageId,
	}
}

// ParseMeOnenotePageID parses 'input' into a MeOnenotePageId
func ParseMeOnenotePageID(input string) (*MeOnenotePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenotePageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenotePageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnenotePageIDInsensitively parses 'input' case-insensitively into a MeOnenotePageId
// note: this method should only be used for API response data and not user input
func ParseMeOnenotePageIDInsensitively(input string) (*MeOnenotePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenotePageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenotePageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnenotePageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateMeOnenotePageID checks that 'input' can be parsed as a Me Onenote Page ID
func ValidateMeOnenotePageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenotePageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Page ID
func (id MeOnenotePageId) ID() string {
	fmtString := "/me/onenote/pages/%s"
	return fmt.Sprintf(fmtString, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Page ID
func (id MeOnenotePageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this Me Onenote Page ID
func (id MeOnenotePageId) String() string {
	components := []string{
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Me Onenote Page (%s)", strings.Join(components, "\n"))
}
