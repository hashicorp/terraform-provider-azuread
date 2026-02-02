package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeVirtualEventWebinarId{}

// MeVirtualEventWebinarId is a struct representing the Resource ID for a Me Virtual Event Webinar
type MeVirtualEventWebinarId struct {
	VirtualEventWebinarId string
}

// NewMeVirtualEventWebinarID returns a new MeVirtualEventWebinarId struct
func NewMeVirtualEventWebinarID(virtualEventWebinarId string) MeVirtualEventWebinarId {
	return MeVirtualEventWebinarId{
		VirtualEventWebinarId: virtualEventWebinarId,
	}
}

// ParseMeVirtualEventWebinarID parses 'input' into a MeVirtualEventWebinarId
func ParseMeVirtualEventWebinarID(input string) (*MeVirtualEventWebinarId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeVirtualEventWebinarId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeVirtualEventWebinarId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeVirtualEventWebinarIDInsensitively parses 'input' case-insensitively into a MeVirtualEventWebinarId
// note: this method should only be used for API response data and not user input
func ParseMeVirtualEventWebinarIDInsensitively(input string) (*MeVirtualEventWebinarId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeVirtualEventWebinarId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeVirtualEventWebinarId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeVirtualEventWebinarId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.VirtualEventWebinarId, ok = input.Parsed["virtualEventWebinarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "virtualEventWebinarId", input)
	}

	return nil
}

// ValidateMeVirtualEventWebinarID checks that 'input' can be parsed as a Me Virtual Event Webinar ID
func ValidateMeVirtualEventWebinarID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeVirtualEventWebinarID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Virtual Event Webinar ID
func (id MeVirtualEventWebinarId) ID() string {
	fmtString := "/me/virtualEvents/webinars/%s"
	return fmt.Sprintf(fmtString, id.VirtualEventWebinarId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Virtual Event Webinar ID
func (id MeVirtualEventWebinarId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("virtualEvents", "virtualEvents", "virtualEvents"),
		resourceids.StaticSegment("webinars", "webinars", "webinars"),
		resourceids.UserSpecifiedSegment("virtualEventWebinarId", "virtualEventWebinarId"),
	}
}

// String returns a human-readable description of this Me Virtual Event Webinar ID
func (id MeVirtualEventWebinarId) String() string {
	components := []string{
		fmt.Sprintf("Virtual Event Webinar: %q", id.VirtualEventWebinarId),
	}
	return fmt.Sprintf("Me Virtual Event Webinar (%s)", strings.Join(components, "\n"))
}
