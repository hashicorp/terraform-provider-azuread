package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryRoleTemplateId{}

// DirectoryRoleTemplateId is a struct representing the Resource ID for a Directory Role Template
type DirectoryRoleTemplateId struct {
	DirectoryRoleTemplateId string
}

// NewDirectoryRoleTemplateID returns a new DirectoryRoleTemplateId struct
func NewDirectoryRoleTemplateID(directoryRoleTemplateId string) DirectoryRoleTemplateId {
	return DirectoryRoleTemplateId{
		DirectoryRoleTemplateId: directoryRoleTemplateId,
	}
}

// ParseDirectoryRoleTemplateID parses 'input' into a DirectoryRoleTemplateId
func ParseDirectoryRoleTemplateID(input string) (*DirectoryRoleTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryRoleTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryRoleTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryRoleTemplateIDInsensitively parses 'input' case-insensitively into a DirectoryRoleTemplateId
// note: this method should only be used for API response data and not user input
func ParseDirectoryRoleTemplateIDInsensitively(input string) (*DirectoryRoleTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryRoleTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryRoleTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryRoleTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DirectoryRoleTemplateId, ok = input.Parsed["directoryRoleTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryRoleTemplateId", input)
	}

	return nil
}

// ValidateDirectoryRoleTemplateID checks that 'input' can be parsed as a Directory Role Template ID
func ValidateDirectoryRoleTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryRoleTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Role Template ID
func (id DirectoryRoleTemplateId) ID() string {
	fmtString := "/directoryRoleTemplates/%s"
	return fmt.Sprintf(fmtString, id.DirectoryRoleTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Role Template ID
func (id DirectoryRoleTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directoryRoleTemplates", "directoryRoleTemplates", "directoryRoleTemplates"),
		resourceids.UserSpecifiedSegment("directoryRoleTemplateId", "directoryRoleTemplateId"),
	}
}

// String returns a human-readable description of this Directory Role Template ID
func (id DirectoryRoleTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Directory Role Template: %q", id.DirectoryRoleTemplateId),
	}
	return fmt.Sprintf("Directory Role Template (%s)", strings.Join(components, "\n"))
}
