package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfilePhoneId{}

// MeProfilePhoneId is a struct representing the Resource ID for a Me Profile Phone
type MeProfilePhoneId struct {
	ItemPhoneId string
}

// NewMeProfilePhoneID returns a new MeProfilePhoneId struct
func NewMeProfilePhoneID(itemPhoneId string) MeProfilePhoneId {
	return MeProfilePhoneId{
		ItemPhoneId: itemPhoneId,
	}
}

// ParseMeProfilePhoneID parses 'input' into a MeProfilePhoneId
func ParseMeProfilePhoneID(input string) (*MeProfilePhoneId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfilePhoneId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfilePhoneId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfilePhoneIDInsensitively parses 'input' case-insensitively into a MeProfilePhoneId
// note: this method should only be used for API response data and not user input
func ParseMeProfilePhoneIDInsensitively(input string) (*MeProfilePhoneId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfilePhoneId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfilePhoneId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfilePhoneId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ItemPhoneId, ok = input.Parsed["itemPhoneId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemPhoneId", input)
	}

	return nil
}

// ValidateMeProfilePhoneID checks that 'input' can be parsed as a Me Profile Phone ID
func ValidateMeProfilePhoneID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfilePhoneID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Phone ID
func (id MeProfilePhoneId) ID() string {
	fmtString := "/me/profile/phones/%s"
	return fmt.Sprintf(fmtString, id.ItemPhoneId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Phone ID
func (id MeProfilePhoneId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("phones", "phones", "phones"),
		resourceids.UserSpecifiedSegment("itemPhoneId", "itemPhoneId"),
	}
}

// String returns a human-readable description of this Me Profile Phone ID
func (id MeProfilePhoneId) String() string {
	components := []string{
		fmt.Sprintf("Item Phone: %q", id.ItemPhoneId),
	}
	return fmt.Sprintf("Me Profile Phone (%s)", strings.Join(components, "\n"))
}
