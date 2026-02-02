package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeSettingWindowIdInstanceId{}

// MeSettingWindowIdInstanceId is a struct representing the Resource ID for a Me Setting Window Id Instance
type MeSettingWindowIdInstanceId struct {
	WindowsSettingId         string
	WindowsSettingInstanceId string
}

// NewMeSettingWindowIdInstanceID returns a new MeSettingWindowIdInstanceId struct
func NewMeSettingWindowIdInstanceID(windowsSettingId string, windowsSettingInstanceId string) MeSettingWindowIdInstanceId {
	return MeSettingWindowIdInstanceId{
		WindowsSettingId:         windowsSettingId,
		WindowsSettingInstanceId: windowsSettingInstanceId,
	}
}

// ParseMeSettingWindowIdInstanceID parses 'input' into a MeSettingWindowIdInstanceId
func ParseMeSettingWindowIdInstanceID(input string) (*MeSettingWindowIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeSettingWindowIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeSettingWindowIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeSettingWindowIdInstanceIDInsensitively parses 'input' case-insensitively into a MeSettingWindowIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeSettingWindowIdInstanceIDInsensitively(input string) (*MeSettingWindowIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeSettingWindowIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeSettingWindowIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeSettingWindowIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsSettingId, ok = input.Parsed["windowsSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsSettingId", input)
	}

	if id.WindowsSettingInstanceId, ok = input.Parsed["windowsSettingInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsSettingInstanceId", input)
	}

	return nil
}

// ValidateMeSettingWindowIdInstanceID checks that 'input' can be parsed as a Me Setting Window Id Instance ID
func ValidateMeSettingWindowIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeSettingWindowIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Setting Window Id Instance ID
func (id MeSettingWindowIdInstanceId) ID() string {
	fmtString := "/me/settings/windows/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.WindowsSettingId, id.WindowsSettingInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Setting Window Id Instance ID
func (id MeSettingWindowIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.StaticSegment("windows", "windows", "windows"),
		resourceids.UserSpecifiedSegment("windowsSettingId", "windowsSettingId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("windowsSettingInstanceId", "windowsSettingInstanceId"),
	}
}

// String returns a human-readable description of this Me Setting Window Id Instance ID
func (id MeSettingWindowIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Windows Setting: %q", id.WindowsSettingId),
		fmt.Sprintf("Windows Setting Instance: %q", id.WindowsSettingInstanceId),
	}
	return fmt.Sprintf("Me Setting Window Id Instance (%s)", strings.Join(components, "\n"))
}
