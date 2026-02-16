package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationWindowsHelloForBusinessMethodId{}

// MeAuthenticationWindowsHelloForBusinessMethodId is a struct representing the Resource ID for a Me Authentication Windows Hello For Business Method
type MeAuthenticationWindowsHelloForBusinessMethodId struct {
	WindowsHelloForBusinessAuthenticationMethodId string
}

// NewMeAuthenticationWindowsHelloForBusinessMethodID returns a new MeAuthenticationWindowsHelloForBusinessMethodId struct
func NewMeAuthenticationWindowsHelloForBusinessMethodID(windowsHelloForBusinessAuthenticationMethodId string) MeAuthenticationWindowsHelloForBusinessMethodId {
	return MeAuthenticationWindowsHelloForBusinessMethodId{
		WindowsHelloForBusinessAuthenticationMethodId: windowsHelloForBusinessAuthenticationMethodId,
	}
}

// ParseMeAuthenticationWindowsHelloForBusinessMethodID parses 'input' into a MeAuthenticationWindowsHelloForBusinessMethodId
func ParseMeAuthenticationWindowsHelloForBusinessMethodID(input string) (*MeAuthenticationWindowsHelloForBusinessMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationWindowsHelloForBusinessMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationWindowsHelloForBusinessMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAuthenticationWindowsHelloForBusinessMethodIDInsensitively parses 'input' case-insensitively into a MeAuthenticationWindowsHelloForBusinessMethodId
// note: this method should only be used for API response data and not user input
func ParseMeAuthenticationWindowsHelloForBusinessMethodIDInsensitively(input string) (*MeAuthenticationWindowsHelloForBusinessMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationWindowsHelloForBusinessMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationWindowsHelloForBusinessMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAuthenticationWindowsHelloForBusinessMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsHelloForBusinessAuthenticationMethodId, ok = input.Parsed["windowsHelloForBusinessAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsHelloForBusinessAuthenticationMethodId", input)
	}

	return nil
}

// ValidateMeAuthenticationWindowsHelloForBusinessMethodID checks that 'input' can be parsed as a Me Authentication Windows Hello For Business Method ID
func ValidateMeAuthenticationWindowsHelloForBusinessMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAuthenticationWindowsHelloForBusinessMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Authentication Windows Hello For Business Method ID
func (id MeAuthenticationWindowsHelloForBusinessMethodId) ID() string {
	fmtString := "/me/authentication/windowsHelloForBusinessMethods/%s"
	return fmt.Sprintf(fmtString, id.WindowsHelloForBusinessAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Authentication Windows Hello For Business Method ID
func (id MeAuthenticationWindowsHelloForBusinessMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("windowsHelloForBusinessMethods", "windowsHelloForBusinessMethods", "windowsHelloForBusinessMethods"),
		resourceids.UserSpecifiedSegment("windowsHelloForBusinessAuthenticationMethodId", "windowsHelloForBusinessAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this Me Authentication Windows Hello For Business Method ID
func (id MeAuthenticationWindowsHelloForBusinessMethodId) String() string {
	components := []string{
		fmt.Sprintf("Windows Hello For Business Authentication Method: %q", id.WindowsHelloForBusinessAuthenticationMethodId),
	}
	return fmt.Sprintf("Me Authentication Windows Hello For Business Method (%s)", strings.Join(components, "\n"))
}
