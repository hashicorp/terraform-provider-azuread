package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMessageIdExtensionId{}

// MeMessageIdExtensionId is a struct representing the Resource ID for a Me Message Id Extension
type MeMessageIdExtensionId struct {
	MessageId   string
	ExtensionId string
}

// NewMeMessageIdExtensionID returns a new MeMessageIdExtensionId struct
func NewMeMessageIdExtensionID(messageId string, extensionId string) MeMessageIdExtensionId {
	return MeMessageIdExtensionId{
		MessageId:   messageId,
		ExtensionId: extensionId,
	}
}

// ParseMeMessageIdExtensionID parses 'input' into a MeMessageIdExtensionId
func ParseMeMessageIdExtensionID(input string) (*MeMessageIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMessageIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMessageIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMessageIdExtensionIDInsensitively parses 'input' case-insensitively into a MeMessageIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeMessageIdExtensionIDInsensitively(input string) (*MeMessageIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMessageIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMessageIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMessageIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MessageId, ok = input.Parsed["messageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeMessageIdExtensionID checks that 'input' can be parsed as a Me Message Id Extension ID
func ValidateMeMessageIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMessageIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Message Id Extension ID
func (id MeMessageIdExtensionId) ID() string {
	fmtString := "/me/messages/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.MessageId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Message Id Extension ID
func (id MeMessageIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Me Message Id Extension ID
func (id MeMessageIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Message Id Extension (%s)", strings.Join(components, "\n"))
}
