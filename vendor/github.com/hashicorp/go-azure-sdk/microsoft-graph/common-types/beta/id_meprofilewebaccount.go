package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileWebAccountId{}

// MeProfileWebAccountId is a struct representing the Resource ID for a Me Profile Web Account
type MeProfileWebAccountId struct {
	WebAccountId string
}

// NewMeProfileWebAccountID returns a new MeProfileWebAccountId struct
func NewMeProfileWebAccountID(webAccountId string) MeProfileWebAccountId {
	return MeProfileWebAccountId{
		WebAccountId: webAccountId,
	}
}

// ParseMeProfileWebAccountID parses 'input' into a MeProfileWebAccountId
func ParseMeProfileWebAccountID(input string) (*MeProfileWebAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileWebAccountId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileWebAccountId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfileWebAccountIDInsensitively parses 'input' case-insensitively into a MeProfileWebAccountId
// note: this method should only be used for API response data and not user input
func ParseMeProfileWebAccountIDInsensitively(input string) (*MeProfileWebAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileWebAccountId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileWebAccountId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfileWebAccountId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WebAccountId, ok = input.Parsed["webAccountId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "webAccountId", input)
	}

	return nil
}

// ValidateMeProfileWebAccountID checks that 'input' can be parsed as a Me Profile Web Account ID
func ValidateMeProfileWebAccountID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileWebAccountID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Web Account ID
func (id MeProfileWebAccountId) ID() string {
	fmtString := "/me/profile/webAccounts/%s"
	return fmt.Sprintf(fmtString, id.WebAccountId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Web Account ID
func (id MeProfileWebAccountId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("webAccounts", "webAccounts", "webAccounts"),
		resourceids.UserSpecifiedSegment("webAccountId", "webAccountId"),
	}
}

// String returns a human-readable description of this Me Profile Web Account ID
func (id MeProfileWebAccountId) String() string {
	components := []string{
		fmt.Sprintf("Web Account: %q", id.WebAccountId),
	}
	return fmt.Sprintf("Me Profile Web Account (%s)", strings.Join(components, "\n"))
}
