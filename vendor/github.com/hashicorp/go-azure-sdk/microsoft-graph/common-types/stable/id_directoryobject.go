package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryObjectId{}

// DirectoryObjectId is a struct representing the Resource ID for a Directory Object
type DirectoryObjectId struct {
	DirectoryObjectId string
}

// NewDirectoryObjectID returns a new DirectoryObjectId struct
func NewDirectoryObjectID(directoryObjectId string) DirectoryObjectId {
	return DirectoryObjectId{
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseDirectoryObjectID parses 'input' into a DirectoryObjectId
func ParseDirectoryObjectID(input string) (*DirectoryObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryObjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryObjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryObjectIDInsensitively parses 'input' case-insensitively into a DirectoryObjectId
// note: this method should only be used for API response data and not user input
func ParseDirectoryObjectIDInsensitively(input string) (*DirectoryObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryObjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryObjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryObjectId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateDirectoryObjectID checks that 'input' can be parsed as a Directory Object ID
func ValidateDirectoryObjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryObjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Object ID
func (id DirectoryObjectId) ID() string {
	fmtString := "/directoryObjects/%s"
	return fmt.Sprintf(fmtString, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Object ID
func (id DirectoryObjectId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directoryObjects", "directoryObjects", "directoryObjects"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Directory Object ID
func (id DirectoryObjectId) String() string {
	components := []string{
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Directory Object (%s)", strings.Join(components, "\n"))
}
