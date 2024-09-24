package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdSynchronizationJobId{}

// ApplicationIdSynchronizationJobId is a struct representing the Resource ID for a Application Id Synchronization Job
type ApplicationIdSynchronizationJobId struct {
	ApplicationId        string
	SynchronizationJobId string
}

// NewApplicationIdSynchronizationJobID returns a new ApplicationIdSynchronizationJobId struct
func NewApplicationIdSynchronizationJobID(applicationId string, synchronizationJobId string) ApplicationIdSynchronizationJobId {
	return ApplicationIdSynchronizationJobId{
		ApplicationId:        applicationId,
		SynchronizationJobId: synchronizationJobId,
	}
}

// ParseApplicationIdSynchronizationJobID parses 'input' into a ApplicationIdSynchronizationJobId
func ParseApplicationIdSynchronizationJobID(input string) (*ApplicationIdSynchronizationJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdSynchronizationJobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdSynchronizationJobId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApplicationIdSynchronizationJobIDInsensitively parses 'input' case-insensitively into a ApplicationIdSynchronizationJobId
// note: this method should only be used for API response data and not user input
func ParseApplicationIdSynchronizationJobIDInsensitively(input string) (*ApplicationIdSynchronizationJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdSynchronizationJobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdSynchronizationJobId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApplicationIdSynchronizationJobId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.SynchronizationJobId, ok = input.Parsed["synchronizationJobId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "synchronizationJobId", input)
	}

	return nil
}

// ValidateApplicationIdSynchronizationJobID checks that 'input' can be parsed as a Application Id Synchronization Job ID
func ValidateApplicationIdSynchronizationJobID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationIdSynchronizationJobID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Id Synchronization Job ID
func (id ApplicationIdSynchronizationJobId) ID() string {
	fmtString := "/applications/%s/synchronization/jobs/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.SynchronizationJobId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Id Synchronization Job ID
func (id ApplicationIdSynchronizationJobId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationId"),
		resourceids.StaticSegment("synchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("jobs", "jobs", "jobs"),
		resourceids.UserSpecifiedSegment("synchronizationJobId", "synchronizationJobId"),
	}
}

// String returns a human-readable description of this Application Id Synchronization Job ID
func (id ApplicationIdSynchronizationJobId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Synchronization Job: %q", id.SynchronizationJobId),
	}
	return fmt.Sprintf("Application Id Synchronization Job (%s)", strings.Join(components, "\n"))
}
