package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryRoleId{}

// DirectoryRoleId is a struct representing the Resource ID for a Directory Role
type DirectoryRoleId struct {
	DirectoryRoleId string
}

// NewDirectoryRoleID returns a new DirectoryRoleId struct
func NewDirectoryRoleID(directoryRoleId string) DirectoryRoleId {
	return DirectoryRoleId{
		DirectoryRoleId: directoryRoleId,
	}
}

// ParseDirectoryRoleID parses 'input' into a DirectoryRoleId
func ParseDirectoryRoleID(input string) (*DirectoryRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryRoleIDInsensitively parses 'input' case-insensitively into a DirectoryRoleId
// note: this method should only be used for API response data and not user input
func ParseDirectoryRoleIDInsensitively(input string) (*DirectoryRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DirectoryRoleId, ok = input.Parsed["directoryRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryRoleId", input)
	}

	return nil
}

// ValidateDirectoryRoleID checks that 'input' can be parsed as a Directory Role ID
func ValidateDirectoryRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Role ID
func (id DirectoryRoleId) ID() string {
	fmtString := "/directoryRoles/%s"
	return fmt.Sprintf(fmtString, id.DirectoryRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Role ID
func (id DirectoryRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directoryRoles", "directoryRoles", "directoryRoles"),
		resourceids.UserSpecifiedSegment("directoryRoleId", "directoryRoleId"),
	}
}

// String returns a human-readable description of this Directory Role ID
func (id DirectoryRoleId) String() string {
	components := []string{
		fmt.Sprintf("Directory Role: %q", id.DirectoryRoleId),
	}
	return fmt.Sprintf("Directory Role (%s)", strings.Join(components, "\n"))
}
