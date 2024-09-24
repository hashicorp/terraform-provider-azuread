package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdSynchronizationTemplateId{}

// ApplicationIdSynchronizationTemplateId is a struct representing the Resource ID for a Application Id Synchronization Template
type ApplicationIdSynchronizationTemplateId struct {
	ApplicationId             string
	SynchronizationTemplateId string
}

// NewApplicationIdSynchronizationTemplateID returns a new ApplicationIdSynchronizationTemplateId struct
func NewApplicationIdSynchronizationTemplateID(applicationId string, synchronizationTemplateId string) ApplicationIdSynchronizationTemplateId {
	return ApplicationIdSynchronizationTemplateId{
		ApplicationId:             applicationId,
		SynchronizationTemplateId: synchronizationTemplateId,
	}
}

// ParseApplicationIdSynchronizationTemplateID parses 'input' into a ApplicationIdSynchronizationTemplateId
func ParseApplicationIdSynchronizationTemplateID(input string) (*ApplicationIdSynchronizationTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdSynchronizationTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdSynchronizationTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApplicationIdSynchronizationTemplateIDInsensitively parses 'input' case-insensitively into a ApplicationIdSynchronizationTemplateId
// note: this method should only be used for API response data and not user input
func ParseApplicationIdSynchronizationTemplateIDInsensitively(input string) (*ApplicationIdSynchronizationTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdSynchronizationTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdSynchronizationTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApplicationIdSynchronizationTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.SynchronizationTemplateId, ok = input.Parsed["synchronizationTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "synchronizationTemplateId", input)
	}

	return nil
}

// ValidateApplicationIdSynchronizationTemplateID checks that 'input' can be parsed as a Application Id Synchronization Template ID
func ValidateApplicationIdSynchronizationTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationIdSynchronizationTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Id Synchronization Template ID
func (id ApplicationIdSynchronizationTemplateId) ID() string {
	fmtString := "/applications/%s/synchronization/templates/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.SynchronizationTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Id Synchronization Template ID
func (id ApplicationIdSynchronizationTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationId"),
		resourceids.StaticSegment("synchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("synchronizationTemplateId", "synchronizationTemplateId"),
	}
}

// String returns a human-readable description of this Application Id Synchronization Template ID
func (id ApplicationIdSynchronizationTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Synchronization Template: %q", id.SynchronizationTemplateId),
	}
	return fmt.Sprintf("Application Id Synchronization Template (%s)", strings.Join(components, "\n"))
}
