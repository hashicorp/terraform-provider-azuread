package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileAccountId{}

// MeProfileAccountId is a struct representing the Resource ID for a Me Profile Account
type MeProfileAccountId struct {
	UserAccountInformationId string
}

// NewMeProfileAccountID returns a new MeProfileAccountId struct
func NewMeProfileAccountID(userAccountInformationId string) MeProfileAccountId {
	return MeProfileAccountId{
		UserAccountInformationId: userAccountInformationId,
	}
}

// ParseMeProfileAccountID parses 'input' into a MeProfileAccountId
func ParseMeProfileAccountID(input string) (*MeProfileAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileAccountId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileAccountId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfileAccountIDInsensitively parses 'input' case-insensitively into a MeProfileAccountId
// note: this method should only be used for API response data and not user input
func ParseMeProfileAccountIDInsensitively(input string) (*MeProfileAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileAccountId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileAccountId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfileAccountId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserAccountInformationId, ok = input.Parsed["userAccountInformationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userAccountInformationId", input)
	}

	return nil
}

// ValidateMeProfileAccountID checks that 'input' can be parsed as a Me Profile Account ID
func ValidateMeProfileAccountID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileAccountID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Account ID
func (id MeProfileAccountId) ID() string {
	fmtString := "/me/profile/account/%s"
	return fmt.Sprintf(fmtString, id.UserAccountInformationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Account ID
func (id MeProfileAccountId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("account", "account", "account"),
		resourceids.UserSpecifiedSegment("userAccountInformationId", "userAccountInformationId"),
	}
}

// String returns a human-readable description of this Me Profile Account ID
func (id MeProfileAccountId) String() string {
	components := []string{
		fmt.Sprintf("User Account Information: %q", id.UserAccountInformationId),
	}
	return fmt.Sprintf("Me Profile Account (%s)", strings.Join(components, "\n"))
}
