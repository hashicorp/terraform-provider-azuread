package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdSynchronizationJobIdSchemaDirectoryId{}

// ApplicationIdSynchronizationJobIdSchemaDirectoryId is a struct representing the Resource ID for a Application Id Synchronization Job Id Schema Directory
type ApplicationIdSynchronizationJobIdSchemaDirectoryId struct {
	ApplicationId         string
	SynchronizationJobId  string
	DirectoryDefinitionId string
}

// NewApplicationIdSynchronizationJobIdSchemaDirectoryID returns a new ApplicationIdSynchronizationJobIdSchemaDirectoryId struct
func NewApplicationIdSynchronizationJobIdSchemaDirectoryID(applicationId string, synchronizationJobId string, directoryDefinitionId string) ApplicationIdSynchronizationJobIdSchemaDirectoryId {
	return ApplicationIdSynchronizationJobIdSchemaDirectoryId{
		ApplicationId:         applicationId,
		SynchronizationJobId:  synchronizationJobId,
		DirectoryDefinitionId: directoryDefinitionId,
	}
}

// ParseApplicationIdSynchronizationJobIdSchemaDirectoryID parses 'input' into a ApplicationIdSynchronizationJobIdSchemaDirectoryId
func ParseApplicationIdSynchronizationJobIdSchemaDirectoryID(input string) (*ApplicationIdSynchronizationJobIdSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdSynchronizationJobIdSchemaDirectoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdSynchronizationJobIdSchemaDirectoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApplicationIdSynchronizationJobIdSchemaDirectoryIDInsensitively parses 'input' case-insensitively into a ApplicationIdSynchronizationJobIdSchemaDirectoryId
// note: this method should only be used for API response data and not user input
func ParseApplicationIdSynchronizationJobIdSchemaDirectoryIDInsensitively(input string) (*ApplicationIdSynchronizationJobIdSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdSynchronizationJobIdSchemaDirectoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdSynchronizationJobIdSchemaDirectoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApplicationIdSynchronizationJobIdSchemaDirectoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.SynchronizationJobId, ok = input.Parsed["synchronizationJobId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "synchronizationJobId", input)
	}

	if id.DirectoryDefinitionId, ok = input.Parsed["directoryDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryDefinitionId", input)
	}

	return nil
}

// ValidateApplicationIdSynchronizationJobIdSchemaDirectoryID checks that 'input' can be parsed as a Application Id Synchronization Job Id Schema Directory ID
func ValidateApplicationIdSynchronizationJobIdSchemaDirectoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationIdSynchronizationJobIdSchemaDirectoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Id Synchronization Job Id Schema Directory ID
func (id ApplicationIdSynchronizationJobIdSchemaDirectoryId) ID() string {
	fmtString := "/applications/%s/synchronization/jobs/%s/schema/directories/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.SynchronizationJobId, id.DirectoryDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Id Synchronization Job Id Schema Directory ID
func (id ApplicationIdSynchronizationJobIdSchemaDirectoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationId"),
		resourceids.StaticSegment("synchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("jobs", "jobs", "jobs"),
		resourceids.UserSpecifiedSegment("synchronizationJobId", "synchronizationJobId"),
		resourceids.StaticSegment("schema", "schema", "schema"),
		resourceids.StaticSegment("directories", "directories", "directories"),
		resourceids.UserSpecifiedSegment("directoryDefinitionId", "directoryDefinitionId"),
	}
}

// String returns a human-readable description of this Application Id Synchronization Job Id Schema Directory ID
func (id ApplicationIdSynchronizationJobIdSchemaDirectoryId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Synchronization Job: %q", id.SynchronizationJobId),
		fmt.Sprintf("Directory Definition: %q", id.DirectoryDefinitionId),
	}
	return fmt.Sprintf("Application Id Synchronization Job Id Schema Directory (%s)", strings.Join(components, "\n"))
}
