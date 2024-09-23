package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCloudClipboardItemId{}

// MeCloudClipboardItemId is a struct representing the Resource ID for a Me Cloud Clipboard Item
type MeCloudClipboardItemId struct {
	CloudClipboardItemId string
}

// NewMeCloudClipboardItemID returns a new MeCloudClipboardItemId struct
func NewMeCloudClipboardItemID(cloudClipboardItemId string) MeCloudClipboardItemId {
	return MeCloudClipboardItemId{
		CloudClipboardItemId: cloudClipboardItemId,
	}
}

// ParseMeCloudClipboardItemID parses 'input' into a MeCloudClipboardItemId
func ParseMeCloudClipboardItemID(input string) (*MeCloudClipboardItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCloudClipboardItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCloudClipboardItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCloudClipboardItemIDInsensitively parses 'input' case-insensitively into a MeCloudClipboardItemId
// note: this method should only be used for API response data and not user input
func ParseMeCloudClipboardItemIDInsensitively(input string) (*MeCloudClipboardItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCloudClipboardItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCloudClipboardItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCloudClipboardItemId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudClipboardItemId, ok = input.Parsed["cloudClipboardItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudClipboardItemId", input)
	}

	return nil
}

// ValidateMeCloudClipboardItemID checks that 'input' can be parsed as a Me Cloud Clipboard Item ID
func ValidateMeCloudClipboardItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCloudClipboardItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Cloud Clipboard Item ID
func (id MeCloudClipboardItemId) ID() string {
	fmtString := "/me/cloudClipboard/items/%s"
	return fmt.Sprintf(fmtString, id.CloudClipboardItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Cloud Clipboard Item ID
func (id MeCloudClipboardItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("cloudClipboard", "cloudClipboard", "cloudClipboard"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("cloudClipboardItemId", "cloudClipboardItemId"),
	}
}

// String returns a human-readable description of this Me Cloud Clipboard Item ID
func (id MeCloudClipboardItemId) String() string {
	components := []string{
		fmt.Sprintf("Cloud Clipboard Item: %q", id.CloudClipboardItemId),
	}
	return fmt.Sprintf("Me Cloud Clipboard Item (%s)", strings.Join(components, "\n"))
}
