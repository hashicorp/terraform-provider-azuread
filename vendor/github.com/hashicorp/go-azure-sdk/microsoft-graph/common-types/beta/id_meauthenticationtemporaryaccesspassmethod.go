package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationTemporaryAccessPassMethodId{}

// MeAuthenticationTemporaryAccessPassMethodId is a struct representing the Resource ID for a Me Authentication Temporary Access Pass Method
type MeAuthenticationTemporaryAccessPassMethodId struct {
	TemporaryAccessPassAuthenticationMethodId string
}

// NewMeAuthenticationTemporaryAccessPassMethodID returns a new MeAuthenticationTemporaryAccessPassMethodId struct
func NewMeAuthenticationTemporaryAccessPassMethodID(temporaryAccessPassAuthenticationMethodId string) MeAuthenticationTemporaryAccessPassMethodId {
	return MeAuthenticationTemporaryAccessPassMethodId{
		TemporaryAccessPassAuthenticationMethodId: temporaryAccessPassAuthenticationMethodId,
	}
}

// ParseMeAuthenticationTemporaryAccessPassMethodID parses 'input' into a MeAuthenticationTemporaryAccessPassMethodId
func ParseMeAuthenticationTemporaryAccessPassMethodID(input string) (*MeAuthenticationTemporaryAccessPassMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationTemporaryAccessPassMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationTemporaryAccessPassMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAuthenticationTemporaryAccessPassMethodIDInsensitively parses 'input' case-insensitively into a MeAuthenticationTemporaryAccessPassMethodId
// note: this method should only be used for API response data and not user input
func ParseMeAuthenticationTemporaryAccessPassMethodIDInsensitively(input string) (*MeAuthenticationTemporaryAccessPassMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationTemporaryAccessPassMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationTemporaryAccessPassMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAuthenticationTemporaryAccessPassMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TemporaryAccessPassAuthenticationMethodId, ok = input.Parsed["temporaryAccessPassAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "temporaryAccessPassAuthenticationMethodId", input)
	}

	return nil
}

// ValidateMeAuthenticationTemporaryAccessPassMethodID checks that 'input' can be parsed as a Me Authentication Temporary Access Pass Method ID
func ValidateMeAuthenticationTemporaryAccessPassMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAuthenticationTemporaryAccessPassMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Authentication Temporary Access Pass Method ID
func (id MeAuthenticationTemporaryAccessPassMethodId) ID() string {
	fmtString := "/me/authentication/temporaryAccessPassMethods/%s"
	return fmt.Sprintf(fmtString, id.TemporaryAccessPassAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Authentication Temporary Access Pass Method ID
func (id MeAuthenticationTemporaryAccessPassMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("temporaryAccessPassMethods", "temporaryAccessPassMethods", "temporaryAccessPassMethods"),
		resourceids.UserSpecifiedSegment("temporaryAccessPassAuthenticationMethodId", "temporaryAccessPassAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this Me Authentication Temporary Access Pass Method ID
func (id MeAuthenticationTemporaryAccessPassMethodId) String() string {
	components := []string{
		fmt.Sprintf("Temporary Access Pass Authentication Method: %q", id.TemporaryAccessPassAuthenticationMethodId),
	}
	return fmt.Sprintf("Me Authentication Temporary Access Pass Method (%s)", strings.Join(components, "\n"))
}
