package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationPlatformCredentialMethodId{}

// MeAuthenticationPlatformCredentialMethodId is a struct representing the Resource ID for a Me Authentication Platform Credential Method
type MeAuthenticationPlatformCredentialMethodId struct {
	PlatformCredentialAuthenticationMethodId string
}

// NewMeAuthenticationPlatformCredentialMethodID returns a new MeAuthenticationPlatformCredentialMethodId struct
func NewMeAuthenticationPlatformCredentialMethodID(platformCredentialAuthenticationMethodId string) MeAuthenticationPlatformCredentialMethodId {
	return MeAuthenticationPlatformCredentialMethodId{
		PlatformCredentialAuthenticationMethodId: platformCredentialAuthenticationMethodId,
	}
}

// ParseMeAuthenticationPlatformCredentialMethodID parses 'input' into a MeAuthenticationPlatformCredentialMethodId
func ParseMeAuthenticationPlatformCredentialMethodID(input string) (*MeAuthenticationPlatformCredentialMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationPlatformCredentialMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationPlatformCredentialMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAuthenticationPlatformCredentialMethodIDInsensitively parses 'input' case-insensitively into a MeAuthenticationPlatformCredentialMethodId
// note: this method should only be used for API response data and not user input
func ParseMeAuthenticationPlatformCredentialMethodIDInsensitively(input string) (*MeAuthenticationPlatformCredentialMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationPlatformCredentialMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationPlatformCredentialMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAuthenticationPlatformCredentialMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PlatformCredentialAuthenticationMethodId, ok = input.Parsed["platformCredentialAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "platformCredentialAuthenticationMethodId", input)
	}

	return nil
}

// ValidateMeAuthenticationPlatformCredentialMethodID checks that 'input' can be parsed as a Me Authentication Platform Credential Method ID
func ValidateMeAuthenticationPlatformCredentialMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAuthenticationPlatformCredentialMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Authentication Platform Credential Method ID
func (id MeAuthenticationPlatformCredentialMethodId) ID() string {
	fmtString := "/me/authentication/platformCredentialMethods/%s"
	return fmt.Sprintf(fmtString, id.PlatformCredentialAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Authentication Platform Credential Method ID
func (id MeAuthenticationPlatformCredentialMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("platformCredentialMethods", "platformCredentialMethods", "platformCredentialMethods"),
		resourceids.UserSpecifiedSegment("platformCredentialAuthenticationMethodId", "platformCredentialAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this Me Authentication Platform Credential Method ID
func (id MeAuthenticationPlatformCredentialMethodId) String() string {
	components := []string{
		fmt.Sprintf("Platform Credential Authentication Method: %q", id.PlatformCredentialAuthenticationMethodId),
	}
	return fmt.Sprintf("Me Authentication Platform Credential Method (%s)", strings.Join(components, "\n"))
}
