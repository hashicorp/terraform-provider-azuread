package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationId{}

// ApplicationId is a struct representing the Resource ID for a Application
type ApplicationId struct {
	ApplicationId string
}

// NewApplicationID returns a new ApplicationId struct
func NewApplicationID(applicationId string) ApplicationId {
	return ApplicationId{
		ApplicationId: applicationId,
	}
}

// ParseApplicationID parses 'input' into a ApplicationId
func ParseApplicationID(input string) (*ApplicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApplicationIDInsensitively parses 'input' case-insensitively into a ApplicationId
// note: this method should only be used for API response data and not user input
func ParseApplicationIDInsensitively(input string) (*ApplicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApplicationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	return nil
}

// ValidateApplicationID checks that 'input' can be parsed as a Application ID
func ValidateApplicationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application ID
func (id ApplicationId) ID() string {
	fmtString := "/applications/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application ID
func (id ApplicationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationId"),
	}
}

// String returns a human-readable description of this Application ID
func (id ApplicationId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
	}
	return fmt.Sprintf("Application (%s)", strings.Join(components, "\n"))
}
