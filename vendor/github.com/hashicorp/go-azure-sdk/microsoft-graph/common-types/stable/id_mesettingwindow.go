package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeSettingWindowId{}

// MeSettingWindowId is a struct representing the Resource ID for a Me Setting Window
type MeSettingWindowId struct {
	WindowsSettingId string
}

// NewMeSettingWindowID returns a new MeSettingWindowId struct
func NewMeSettingWindowID(windowsSettingId string) MeSettingWindowId {
	return MeSettingWindowId{
		WindowsSettingId: windowsSettingId,
	}
}

// ParseMeSettingWindowID parses 'input' into a MeSettingWindowId
func ParseMeSettingWindowID(input string) (*MeSettingWindowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeSettingWindowId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeSettingWindowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeSettingWindowIDInsensitively parses 'input' case-insensitively into a MeSettingWindowId
// note: this method should only be used for API response data and not user input
func ParseMeSettingWindowIDInsensitively(input string) (*MeSettingWindowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeSettingWindowId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeSettingWindowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeSettingWindowId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsSettingId, ok = input.Parsed["windowsSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsSettingId", input)
	}

	return nil
}

// ValidateMeSettingWindowID checks that 'input' can be parsed as a Me Setting Window ID
func ValidateMeSettingWindowID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeSettingWindowID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Setting Window ID
func (id MeSettingWindowId) ID() string {
	fmtString := "/me/settings/windows/%s"
	return fmt.Sprintf(fmtString, id.WindowsSettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Setting Window ID
func (id MeSettingWindowId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.StaticSegment("windows", "windows", "windows"),
		resourceids.UserSpecifiedSegment("windowsSettingId", "windowsSettingId"),
	}
}

// String returns a human-readable description of this Me Setting Window ID
func (id MeSettingWindowId) String() string {
	components := []string{
		fmt.Sprintf("Windows Setting: %q", id.WindowsSettingId),
	}
	return fmt.Sprintf("Me Setting Window (%s)", strings.Join(components, "\n"))
}
