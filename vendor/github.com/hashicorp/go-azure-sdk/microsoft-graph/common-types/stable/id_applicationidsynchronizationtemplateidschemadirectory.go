package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdSynchronizationTemplateIdSchemaDirectoryId{}

// ApplicationIdSynchronizationTemplateIdSchemaDirectoryId is a struct representing the Resource ID for a Application Id Synchronization Template Id Schema Directory
type ApplicationIdSynchronizationTemplateIdSchemaDirectoryId struct {
	ApplicationId             string
	SynchronizationTemplateId string
	DirectoryDefinitionId     string
}

// NewApplicationIdSynchronizationTemplateIdSchemaDirectoryID returns a new ApplicationIdSynchronizationTemplateIdSchemaDirectoryId struct
func NewApplicationIdSynchronizationTemplateIdSchemaDirectoryID(applicationId string, synchronizationTemplateId string, directoryDefinitionId string) ApplicationIdSynchronizationTemplateIdSchemaDirectoryId {
	return ApplicationIdSynchronizationTemplateIdSchemaDirectoryId{
		ApplicationId:             applicationId,
		SynchronizationTemplateId: synchronizationTemplateId,
		DirectoryDefinitionId:     directoryDefinitionId,
	}
}

// ParseApplicationIdSynchronizationTemplateIdSchemaDirectoryID parses 'input' into a ApplicationIdSynchronizationTemplateIdSchemaDirectoryId
func ParseApplicationIdSynchronizationTemplateIdSchemaDirectoryID(input string) (*ApplicationIdSynchronizationTemplateIdSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdSynchronizationTemplateIdSchemaDirectoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdSynchronizationTemplateIdSchemaDirectoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApplicationIdSynchronizationTemplateIdSchemaDirectoryIDInsensitively parses 'input' case-insensitively into a ApplicationIdSynchronizationTemplateIdSchemaDirectoryId
// note: this method should only be used for API response data and not user input
func ParseApplicationIdSynchronizationTemplateIdSchemaDirectoryIDInsensitively(input string) (*ApplicationIdSynchronizationTemplateIdSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdSynchronizationTemplateIdSchemaDirectoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdSynchronizationTemplateIdSchemaDirectoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApplicationIdSynchronizationTemplateIdSchemaDirectoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.SynchronizationTemplateId, ok = input.Parsed["synchronizationTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "synchronizationTemplateId", input)
	}

	if id.DirectoryDefinitionId, ok = input.Parsed["directoryDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryDefinitionId", input)
	}

	return nil
}

// ValidateApplicationIdSynchronizationTemplateIdSchemaDirectoryID checks that 'input' can be parsed as a Application Id Synchronization Template Id Schema Directory ID
func ValidateApplicationIdSynchronizationTemplateIdSchemaDirectoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationIdSynchronizationTemplateIdSchemaDirectoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Id Synchronization Template Id Schema Directory ID
func (id ApplicationIdSynchronizationTemplateIdSchemaDirectoryId) ID() string {
	fmtString := "/applications/%s/synchronization/templates/%s/schema/directories/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.SynchronizationTemplateId, id.DirectoryDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Id Synchronization Template Id Schema Directory ID
func (id ApplicationIdSynchronizationTemplateIdSchemaDirectoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationId"),
		resourceids.StaticSegment("synchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("synchronizationTemplateId", "synchronizationTemplateId"),
		resourceids.StaticSegment("schema", "schema", "schema"),
		resourceids.StaticSegment("directories", "directories", "directories"),
		resourceids.UserSpecifiedSegment("directoryDefinitionId", "directoryDefinitionId"),
	}
}

// String returns a human-readable description of this Application Id Synchronization Template Id Schema Directory ID
func (id ApplicationIdSynchronizationTemplateIdSchemaDirectoryId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Synchronization Template: %q", id.SynchronizationTemplateId),
		fmt.Sprintf("Directory Definition: %q", id.DirectoryDefinitionId),
	}
	return fmt.Sprintf("Application Id Synchronization Template Id Schema Directory (%s)", strings.Join(components, "\n"))
}
