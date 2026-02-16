package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileAddressId{}

// MeProfileAddressId is a struct representing the Resource ID for a Me Profile Address
type MeProfileAddressId struct {
	ItemAddressId string
}

// NewMeProfileAddressID returns a new MeProfileAddressId struct
func NewMeProfileAddressID(itemAddressId string) MeProfileAddressId {
	return MeProfileAddressId{
		ItemAddressId: itemAddressId,
	}
}

// ParseMeProfileAddressID parses 'input' into a MeProfileAddressId
func ParseMeProfileAddressID(input string) (*MeProfileAddressId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileAddressId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileAddressId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfileAddressIDInsensitively parses 'input' case-insensitively into a MeProfileAddressId
// note: this method should only be used for API response data and not user input
func ParseMeProfileAddressIDInsensitively(input string) (*MeProfileAddressId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileAddressId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileAddressId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfileAddressId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ItemAddressId, ok = input.Parsed["itemAddressId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemAddressId", input)
	}

	return nil
}

// ValidateMeProfileAddressID checks that 'input' can be parsed as a Me Profile Address ID
func ValidateMeProfileAddressID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileAddressID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Address ID
func (id MeProfileAddressId) ID() string {
	fmtString := "/me/profile/addresses/%s"
	return fmt.Sprintf(fmtString, id.ItemAddressId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Address ID
func (id MeProfileAddressId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("addresses", "addresses", "addresses"),
		resourceids.UserSpecifiedSegment("itemAddressId", "itemAddressId"),
	}
}

// String returns a human-readable description of this Me Profile Address ID
func (id MeProfileAddressId) String() string {
	components := []string{
		fmt.Sprintf("Item Address: %q", id.ItemAddressId),
	}
	return fmt.Sprintf("Me Profile Address (%s)", strings.Join(components, "\n"))
}
