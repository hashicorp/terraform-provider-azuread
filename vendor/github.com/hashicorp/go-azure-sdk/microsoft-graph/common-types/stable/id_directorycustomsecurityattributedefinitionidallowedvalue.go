package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId{}

// DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId is a struct representing the Resource ID for a Directory Custom Security Attribute Definition Id Allowed Value
type DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId struct {
	CustomSecurityAttributeDefinitionId string
	AllowedValueId                      string
}

// NewDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID returns a new DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId struct
func NewDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID(customSecurityAttributeDefinitionId string, allowedValueId string) DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId {
	return DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId{
		CustomSecurityAttributeDefinitionId: customSecurityAttributeDefinitionId,
		AllowedValueId:                      allowedValueId,
	}
}

// ParseDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID parses 'input' into a DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId
func ParseDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID(input string) (*DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryCustomSecurityAttributeDefinitionIdAllowedValueIDInsensitively parses 'input' case-insensitively into a DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId
// note: this method should only be used for API response data and not user input
func ParseDirectoryCustomSecurityAttributeDefinitionIdAllowedValueIDInsensitively(input string) (*DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CustomSecurityAttributeDefinitionId, ok = input.Parsed["customSecurityAttributeDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customSecurityAttributeDefinitionId", input)
	}

	if id.AllowedValueId, ok = input.Parsed["allowedValueId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "allowedValueId", input)
	}

	return nil
}

// ValidateDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID checks that 'input' can be parsed as a Directory Custom Security Attribute Definition Id Allowed Value ID
func ValidateDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Custom Security Attribute Definition Id Allowed Value ID
func (id DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId) ID() string {
	fmtString := "/directory/customSecurityAttributeDefinitions/%s/allowedValues/%s"
	return fmt.Sprintf(fmtString, id.CustomSecurityAttributeDefinitionId, id.AllowedValueId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Custom Security Attribute Definition Id Allowed Value ID
func (id DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("customSecurityAttributeDefinitions", "customSecurityAttributeDefinitions", "customSecurityAttributeDefinitions"),
		resourceids.UserSpecifiedSegment("customSecurityAttributeDefinitionId", "customSecurityAttributeDefinitionId"),
		resourceids.StaticSegment("allowedValues", "allowedValues", "allowedValues"),
		resourceids.UserSpecifiedSegment("allowedValueId", "allowedValueId"),
	}
}

// String returns a human-readable description of this Directory Custom Security Attribute Definition Id Allowed Value ID
func (id DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId) String() string {
	components := []string{
		fmt.Sprintf("Custom Security Attribute Definition: %q", id.CustomSecurityAttributeDefinitionId),
		fmt.Sprintf("Allowed Value: %q", id.AllowedValueId),
	}
	return fmt.Sprintf("Directory Custom Security Attribute Definition Id Allowed Value (%s)", strings.Join(components, "\n"))
}
