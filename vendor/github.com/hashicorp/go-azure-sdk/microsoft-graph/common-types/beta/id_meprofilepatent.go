package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfilePatentId{}

// MeProfilePatentId is a struct representing the Resource ID for a Me Profile Patent
type MeProfilePatentId struct {
	ItemPatentId string
}

// NewMeProfilePatentID returns a new MeProfilePatentId struct
func NewMeProfilePatentID(itemPatentId string) MeProfilePatentId {
	return MeProfilePatentId{
		ItemPatentId: itemPatentId,
	}
}

// ParseMeProfilePatentID parses 'input' into a MeProfilePatentId
func ParseMeProfilePatentID(input string) (*MeProfilePatentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfilePatentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfilePatentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfilePatentIDInsensitively parses 'input' case-insensitively into a MeProfilePatentId
// note: this method should only be used for API response data and not user input
func ParseMeProfilePatentIDInsensitively(input string) (*MeProfilePatentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfilePatentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfilePatentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfilePatentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ItemPatentId, ok = input.Parsed["itemPatentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemPatentId", input)
	}

	return nil
}

// ValidateMeProfilePatentID checks that 'input' can be parsed as a Me Profile Patent ID
func ValidateMeProfilePatentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfilePatentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Patent ID
func (id MeProfilePatentId) ID() string {
	fmtString := "/me/profile/patents/%s"
	return fmt.Sprintf(fmtString, id.ItemPatentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Patent ID
func (id MeProfilePatentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("patents", "patents", "patents"),
		resourceids.UserSpecifiedSegment("itemPatentId", "itemPatentId"),
	}
}

// String returns a human-readable description of this Me Profile Patent ID
func (id MeProfilePatentId) String() string {
	components := []string{
		fmt.Sprintf("Item Patent: %q", id.ItemPatentId),
	}
	return fmt.Sprintf("Me Profile Patent (%s)", strings.Join(components, "\n"))
}
