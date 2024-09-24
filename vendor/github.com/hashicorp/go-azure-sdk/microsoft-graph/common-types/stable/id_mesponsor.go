package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeSponsorId{}

// MeSponsorId is a struct representing the Resource ID for a Me Sponsor
type MeSponsorId struct {
	DirectoryObjectId string
}

// NewMeSponsorID returns a new MeSponsorId struct
func NewMeSponsorID(directoryObjectId string) MeSponsorId {
	return MeSponsorId{
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeSponsorID parses 'input' into a MeSponsorId
func ParseMeSponsorID(input string) (*MeSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeSponsorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeSponsorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeSponsorIDInsensitively parses 'input' case-insensitively into a MeSponsorId
// note: this method should only be used for API response data and not user input
func ParseMeSponsorIDInsensitively(input string) (*MeSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeSponsorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeSponsorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeSponsorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateMeSponsorID checks that 'input' can be parsed as a Me Sponsor ID
func ValidateMeSponsorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeSponsorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Sponsor ID
func (id MeSponsorId) ID() string {
	fmtString := "/me/sponsors/%s"
	return fmt.Sprintf(fmtString, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Sponsor ID
func (id MeSponsorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("sponsors", "sponsors", "sponsors"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Me Sponsor ID
func (id MeSponsorId) String() string {
	components := []string{
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Sponsor (%s)", strings.Join(components, "\n"))
}
