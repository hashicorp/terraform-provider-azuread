package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationHardwareOathMethodId{}

// MeAuthenticationHardwareOathMethodId is a struct representing the Resource ID for a Me Authentication Hardware Oath Method
type MeAuthenticationHardwareOathMethodId struct {
	HardwareOathAuthenticationMethodId string
}

// NewMeAuthenticationHardwareOathMethodID returns a new MeAuthenticationHardwareOathMethodId struct
func NewMeAuthenticationHardwareOathMethodID(hardwareOathAuthenticationMethodId string) MeAuthenticationHardwareOathMethodId {
	return MeAuthenticationHardwareOathMethodId{
		HardwareOathAuthenticationMethodId: hardwareOathAuthenticationMethodId,
	}
}

// ParseMeAuthenticationHardwareOathMethodID parses 'input' into a MeAuthenticationHardwareOathMethodId
func ParseMeAuthenticationHardwareOathMethodID(input string) (*MeAuthenticationHardwareOathMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationHardwareOathMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationHardwareOathMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAuthenticationHardwareOathMethodIDInsensitively parses 'input' case-insensitively into a MeAuthenticationHardwareOathMethodId
// note: this method should only be used for API response data and not user input
func ParseMeAuthenticationHardwareOathMethodIDInsensitively(input string) (*MeAuthenticationHardwareOathMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationHardwareOathMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationHardwareOathMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAuthenticationHardwareOathMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.HardwareOathAuthenticationMethodId, ok = input.Parsed["hardwareOathAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwareOathAuthenticationMethodId", input)
	}

	return nil
}

// ValidateMeAuthenticationHardwareOathMethodID checks that 'input' can be parsed as a Me Authentication Hardware Oath Method ID
func ValidateMeAuthenticationHardwareOathMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAuthenticationHardwareOathMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Authentication Hardware Oath Method ID
func (id MeAuthenticationHardwareOathMethodId) ID() string {
	fmtString := "/me/authentication/hardwareOathMethods/%s"
	return fmt.Sprintf(fmtString, id.HardwareOathAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Authentication Hardware Oath Method ID
func (id MeAuthenticationHardwareOathMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("hardwareOathMethods", "hardwareOathMethods", "hardwareOathMethods"),
		resourceids.UserSpecifiedSegment("hardwareOathAuthenticationMethodId", "hardwareOathAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this Me Authentication Hardware Oath Method ID
func (id MeAuthenticationHardwareOathMethodId) String() string {
	components := []string{
		fmt.Sprintf("Hardware Oath Authentication Method: %q", id.HardwareOathAuthenticationMethodId),
	}
	return fmt.Sprintf("Me Authentication Hardware Oath Method (%s)", strings.Join(components, "\n"))
}
