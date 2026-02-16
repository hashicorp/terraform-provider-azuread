package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationTemplateId{}

// ApplicationTemplateId is a struct representing the Resource ID for a Application Template
type ApplicationTemplateId struct {
	ApplicationTemplateId string
}

// NewApplicationTemplateID returns a new ApplicationTemplateId struct
func NewApplicationTemplateID(applicationTemplateId string) ApplicationTemplateId {
	return ApplicationTemplateId{
		ApplicationTemplateId: applicationTemplateId,
	}
}

// ParseApplicationTemplateID parses 'input' into a ApplicationTemplateId
func ParseApplicationTemplateID(input string) (*ApplicationTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApplicationTemplateIDInsensitively parses 'input' case-insensitively into a ApplicationTemplateId
// note: this method should only be used for API response data and not user input
func ParseApplicationTemplateIDInsensitively(input string) (*ApplicationTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApplicationTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationTemplateId, ok = input.Parsed["applicationTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationTemplateId", input)
	}

	return nil
}

// ValidateApplicationTemplateID checks that 'input' can be parsed as a Application Template ID
func ValidateApplicationTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Template ID
func (id ApplicationTemplateId) ID() string {
	fmtString := "/applicationTemplates/%s"
	return fmt.Sprintf(fmtString, id.ApplicationTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Template ID
func (id ApplicationTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applicationTemplates", "applicationTemplates", "applicationTemplates"),
		resourceids.UserSpecifiedSegment("applicationTemplateId", "applicationTemplateId"),
	}
}

// String returns a human-readable description of this Application Template ID
func (id ApplicationTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Application Template: %q", id.ApplicationTemplateId),
	}
	return fmt.Sprintf("Application Template (%s)", strings.Join(components, "\n"))
}
