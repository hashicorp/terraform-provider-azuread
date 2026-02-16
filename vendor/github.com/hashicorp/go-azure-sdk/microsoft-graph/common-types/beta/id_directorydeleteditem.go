package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryDeletedItemId{}

// DirectoryDeletedItemId is a struct representing the Resource ID for a Directory Deleted Item
type DirectoryDeletedItemId struct {
	DirectoryObjectId string
}

// NewDirectoryDeletedItemID returns a new DirectoryDeletedItemId struct
func NewDirectoryDeletedItemID(directoryObjectId string) DirectoryDeletedItemId {
	return DirectoryDeletedItemId{
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseDirectoryDeletedItemID parses 'input' into a DirectoryDeletedItemId
func ParseDirectoryDeletedItemID(input string) (*DirectoryDeletedItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryDeletedItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryDeletedItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryDeletedItemIDInsensitively parses 'input' case-insensitively into a DirectoryDeletedItemId
// note: this method should only be used for API response data and not user input
func ParseDirectoryDeletedItemIDInsensitively(input string) (*DirectoryDeletedItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryDeletedItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryDeletedItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryDeletedItemId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateDirectoryDeletedItemID checks that 'input' can be parsed as a Directory Deleted Item ID
func ValidateDirectoryDeletedItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryDeletedItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Deleted Item ID
func (id DirectoryDeletedItemId) ID() string {
	fmtString := "/directory/deletedItems/%s"
	return fmt.Sprintf(fmtString, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Deleted Item ID
func (id DirectoryDeletedItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("deletedItems", "deletedItems", "deletedItems"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Directory Deleted Item ID
func (id DirectoryDeletedItemId) String() string {
	components := []string{
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Directory Deleted Item (%s)", strings.Join(components, "\n"))
}
