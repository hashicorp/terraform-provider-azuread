package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteResourceId{}

// MeOnenoteResourceId is a struct representing the Resource ID for a Me Onenote Resource
type MeOnenoteResourceId struct {
	OnenoteResourceId string
}

// NewMeOnenoteResourceID returns a new MeOnenoteResourceId struct
func NewMeOnenoteResourceID(onenoteResourceId string) MeOnenoteResourceId {
	return MeOnenoteResourceId{
		OnenoteResourceId: onenoteResourceId,
	}
}

// ParseMeOnenoteResourceID parses 'input' into a MeOnenoteResourceId
func ParseMeOnenoteResourceID(input string) (*MeOnenoteResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnenoteResourceIDInsensitively parses 'input' case-insensitively into a MeOnenoteResourceId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteResourceIDInsensitively(input string) (*MeOnenoteResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnenoteResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OnenoteResourceId, ok = input.Parsed["onenoteResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteResourceId", input)
	}

	return nil
}

// ValidateMeOnenoteResourceID checks that 'input' can be parsed as a Me Onenote Resource ID
func ValidateMeOnenoteResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Resource ID
func (id MeOnenoteResourceId) ID() string {
	fmtString := "/me/onenote/resources/%s"
	return fmt.Sprintf(fmtString, id.OnenoteResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Resource ID
func (id MeOnenoteResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("onenoteResourceId", "onenoteResourceId"),
	}
}

// String returns a human-readable description of this Me Onenote Resource ID
func (id MeOnenoteResourceId) String() string {
	components := []string{
		fmt.Sprintf("Onenote Resource: %q", id.OnenoteResourceId),
	}
	return fmt.Sprintf("Me Onenote Resource (%s)", strings.Join(components, "\n"))
}
