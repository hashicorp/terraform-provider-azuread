package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SchemaExtensionId{}

// SchemaExtensionId is a struct representing the Resource ID for a Schema Extension
type SchemaExtensionId struct {
	SchemaExtensionId string
}

// NewSchemaExtensionID returns a new SchemaExtensionId struct
func NewSchemaExtensionID(schemaExtensionId string) SchemaExtensionId {
	return SchemaExtensionId{
		SchemaExtensionId: schemaExtensionId,
	}
}

// ParseSchemaExtensionID parses 'input' into a SchemaExtensionId
func ParseSchemaExtensionID(input string) (*SchemaExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SchemaExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SchemaExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSchemaExtensionIDInsensitively parses 'input' case-insensitively into a SchemaExtensionId
// note: this method should only be used for API response data and not user input
func ParseSchemaExtensionIDInsensitively(input string) (*SchemaExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SchemaExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SchemaExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SchemaExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SchemaExtensionId, ok = input.Parsed["schemaExtensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "schemaExtensionId", input)
	}

	return nil
}

// ValidateSchemaExtensionID checks that 'input' can be parsed as a Schema Extension ID
func ValidateSchemaExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSchemaExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Schema Extension ID
func (id SchemaExtensionId) ID() string {
	fmtString := "/schemaExtensions/%s"
	return fmt.Sprintf(fmtString, id.SchemaExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Schema Extension ID
func (id SchemaExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("schemaExtensions", "schemaExtensions", "schemaExtensions"),
		resourceids.UserSpecifiedSegment("schemaExtensionId", "schemaExtensionId"),
	}
}

// String returns a human-readable description of this Schema Extension ID
func (id SchemaExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Schema Extension: %q", id.SchemaExtensionId),
	}
	return fmt.Sprintf("Schema Extension (%s)", strings.Join(components, "\n"))
}
