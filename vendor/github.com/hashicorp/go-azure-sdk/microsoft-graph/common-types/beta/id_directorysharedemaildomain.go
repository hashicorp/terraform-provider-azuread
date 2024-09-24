package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectorySharedEmailDomainId{}

// DirectorySharedEmailDomainId is a struct representing the Resource ID for a Directory Shared Email Domain
type DirectorySharedEmailDomainId struct {
	SharedEmailDomainId string
}

// NewDirectorySharedEmailDomainID returns a new DirectorySharedEmailDomainId struct
func NewDirectorySharedEmailDomainID(sharedEmailDomainId string) DirectorySharedEmailDomainId {
	return DirectorySharedEmailDomainId{
		SharedEmailDomainId: sharedEmailDomainId,
	}
}

// ParseDirectorySharedEmailDomainID parses 'input' into a DirectorySharedEmailDomainId
func ParseDirectorySharedEmailDomainID(input string) (*DirectorySharedEmailDomainId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectorySharedEmailDomainId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectorySharedEmailDomainId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectorySharedEmailDomainIDInsensitively parses 'input' case-insensitively into a DirectorySharedEmailDomainId
// note: this method should only be used for API response data and not user input
func ParseDirectorySharedEmailDomainIDInsensitively(input string) (*DirectorySharedEmailDomainId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectorySharedEmailDomainId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectorySharedEmailDomainId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectorySharedEmailDomainId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SharedEmailDomainId, ok = input.Parsed["sharedEmailDomainId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sharedEmailDomainId", input)
	}

	return nil
}

// ValidateDirectorySharedEmailDomainID checks that 'input' can be parsed as a Directory Shared Email Domain ID
func ValidateDirectorySharedEmailDomainID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectorySharedEmailDomainID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Shared Email Domain ID
func (id DirectorySharedEmailDomainId) ID() string {
	fmtString := "/directory/sharedEmailDomains/%s"
	return fmt.Sprintf(fmtString, id.SharedEmailDomainId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Shared Email Domain ID
func (id DirectorySharedEmailDomainId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("sharedEmailDomains", "sharedEmailDomains", "sharedEmailDomains"),
		resourceids.UserSpecifiedSegment("sharedEmailDomainId", "sharedEmailDomainId"),
	}
}

// String returns a human-readable description of this Directory Shared Email Domain ID
func (id DirectorySharedEmailDomainId) String() string {
	components := []string{
		fmt.Sprintf("Shared Email Domain: %q", id.SharedEmailDomainId),
	}
	return fmt.Sprintf("Directory Shared Email Domain (%s)", strings.Join(components, "\n"))
}
