package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdExtensionPropertyId{}

// ApplicationIdExtensionPropertyId is a struct representing the Resource ID for a Application Id Extension Property
type ApplicationIdExtensionPropertyId struct {
	ApplicationId       string
	ExtensionPropertyId string
}

// NewApplicationIdExtensionPropertyID returns a new ApplicationIdExtensionPropertyId struct
func NewApplicationIdExtensionPropertyID(applicationId string, extensionPropertyId string) ApplicationIdExtensionPropertyId {
	return ApplicationIdExtensionPropertyId{
		ApplicationId:       applicationId,
		ExtensionPropertyId: extensionPropertyId,
	}
}

// ParseApplicationIdExtensionPropertyID parses 'input' into a ApplicationIdExtensionPropertyId
func ParseApplicationIdExtensionPropertyID(input string) (*ApplicationIdExtensionPropertyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdExtensionPropertyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdExtensionPropertyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApplicationIdExtensionPropertyIDInsensitively parses 'input' case-insensitively into a ApplicationIdExtensionPropertyId
// note: this method should only be used for API response data and not user input
func ParseApplicationIdExtensionPropertyIDInsensitively(input string) (*ApplicationIdExtensionPropertyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdExtensionPropertyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdExtensionPropertyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApplicationIdExtensionPropertyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.ExtensionPropertyId, ok = input.Parsed["extensionPropertyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionPropertyId", input)
	}

	return nil
}

// ValidateApplicationIdExtensionPropertyID checks that 'input' can be parsed as a Application Id Extension Property ID
func ValidateApplicationIdExtensionPropertyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationIdExtensionPropertyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Id Extension Property ID
func (id ApplicationIdExtensionPropertyId) ID() string {
	fmtString := "/applications/%s/extensionProperties/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.ExtensionPropertyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Id Extension Property ID
func (id ApplicationIdExtensionPropertyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationId"),
		resourceids.StaticSegment("extensionProperties", "extensionProperties", "extensionProperties"),
		resourceids.UserSpecifiedSegment("extensionPropertyId", "extensionPropertyId"),
	}
}

// String returns a human-readable description of this Application Id Extension Property ID
func (id ApplicationIdExtensionPropertyId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Extension Property: %q", id.ExtensionPropertyId),
	}
	return fmt.Sprintf("Application Id Extension Property (%s)", strings.Join(components, "\n"))
}
