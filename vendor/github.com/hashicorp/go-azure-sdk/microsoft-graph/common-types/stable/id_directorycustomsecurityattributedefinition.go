package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryCustomSecurityAttributeDefinitionId{}

// DirectoryCustomSecurityAttributeDefinitionId is a struct representing the Resource ID for a Directory Custom Security Attribute Definition
type DirectoryCustomSecurityAttributeDefinitionId struct {
	CustomSecurityAttributeDefinitionId string
}

// NewDirectoryCustomSecurityAttributeDefinitionID returns a new DirectoryCustomSecurityAttributeDefinitionId struct
func NewDirectoryCustomSecurityAttributeDefinitionID(customSecurityAttributeDefinitionId string) DirectoryCustomSecurityAttributeDefinitionId {
	return DirectoryCustomSecurityAttributeDefinitionId{
		CustomSecurityAttributeDefinitionId: customSecurityAttributeDefinitionId,
	}
}

// ParseDirectoryCustomSecurityAttributeDefinitionID parses 'input' into a DirectoryCustomSecurityAttributeDefinitionId
func ParseDirectoryCustomSecurityAttributeDefinitionID(input string) (*DirectoryCustomSecurityAttributeDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryCustomSecurityAttributeDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryCustomSecurityAttributeDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryCustomSecurityAttributeDefinitionIDInsensitively parses 'input' case-insensitively into a DirectoryCustomSecurityAttributeDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDirectoryCustomSecurityAttributeDefinitionIDInsensitively(input string) (*DirectoryCustomSecurityAttributeDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryCustomSecurityAttributeDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryCustomSecurityAttributeDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryCustomSecurityAttributeDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CustomSecurityAttributeDefinitionId, ok = input.Parsed["customSecurityAttributeDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customSecurityAttributeDefinitionId", input)
	}

	return nil
}

// ValidateDirectoryCustomSecurityAttributeDefinitionID checks that 'input' can be parsed as a Directory Custom Security Attribute Definition ID
func ValidateDirectoryCustomSecurityAttributeDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryCustomSecurityAttributeDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Custom Security Attribute Definition ID
func (id DirectoryCustomSecurityAttributeDefinitionId) ID() string {
	fmtString := "/directory/customSecurityAttributeDefinitions/%s"
	return fmt.Sprintf(fmtString, id.CustomSecurityAttributeDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Custom Security Attribute Definition ID
func (id DirectoryCustomSecurityAttributeDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("customSecurityAttributeDefinitions", "customSecurityAttributeDefinitions", "customSecurityAttributeDefinitions"),
		resourceids.UserSpecifiedSegment("customSecurityAttributeDefinitionId", "customSecurityAttributeDefinitionId"),
	}
}

// String returns a human-readable description of this Directory Custom Security Attribute Definition ID
func (id DirectoryCustomSecurityAttributeDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Custom Security Attribute Definition: %q", id.CustomSecurityAttributeDefinitionId),
	}
	return fmt.Sprintf("Directory Custom Security Attribute Definition (%s)", strings.Join(components, "\n"))
}
