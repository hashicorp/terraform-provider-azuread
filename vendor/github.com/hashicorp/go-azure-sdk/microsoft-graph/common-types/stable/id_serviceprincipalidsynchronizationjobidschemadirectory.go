package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId{}

// ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId is a struct representing the Resource ID for a Service Principal Id Synchronization Job Id Schema Directory
type ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId struct {
	ServicePrincipalId    string
	SynchronizationJobId  string
	DirectoryDefinitionId string
}

// NewServicePrincipalIdSynchronizationJobIdSchemaDirectoryID returns a new ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId struct
func NewServicePrincipalIdSynchronizationJobIdSchemaDirectoryID(servicePrincipalId string, synchronizationJobId string, directoryDefinitionId string) ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId {
	return ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId{
		ServicePrincipalId:    servicePrincipalId,
		SynchronizationJobId:  synchronizationJobId,
		DirectoryDefinitionId: directoryDefinitionId,
	}
}

// ParseServicePrincipalIdSynchronizationJobIdSchemaDirectoryID parses 'input' into a ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId
func ParseServicePrincipalIdSynchronizationJobIdSchemaDirectoryID(input string) (*ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdSynchronizationJobIdSchemaDirectoryIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdSynchronizationJobIdSchemaDirectoryIDInsensitively(input string) (*ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.SynchronizationJobId, ok = input.Parsed["synchronizationJobId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "synchronizationJobId", input)
	}

	if id.DirectoryDefinitionId, ok = input.Parsed["directoryDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryDefinitionId", input)
	}

	return nil
}

// ValidateServicePrincipalIdSynchronizationJobIdSchemaDirectoryID checks that 'input' can be parsed as a Service Principal Id Synchronization Job Id Schema Directory ID
func ValidateServicePrincipalIdSynchronizationJobIdSchemaDirectoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdSynchronizationJobIdSchemaDirectoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Synchronization Job Id Schema Directory ID
func (id ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId) ID() string {
	fmtString := "/servicePrincipals/%s/synchronization/jobs/%s/schema/directories/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.SynchronizationJobId, id.DirectoryDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Synchronization Job Id Schema Directory ID
func (id ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("synchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("jobs", "jobs", "jobs"),
		resourceids.UserSpecifiedSegment("synchronizationJobId", "synchronizationJobId"),
		resourceids.StaticSegment("schema", "schema", "schema"),
		resourceids.StaticSegment("directories", "directories", "directories"),
		resourceids.UserSpecifiedSegment("directoryDefinitionId", "directoryDefinitionId"),
	}
}

// String returns a human-readable description of this Service Principal Id Synchronization Job Id Schema Directory ID
func (id ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Synchronization Job: %q", id.SynchronizationJobId),
		fmt.Sprintf("Directory Definition: %q", id.DirectoryDefinitionId),
	}
	return fmt.Sprintf("Service Principal Id Synchronization Job Id Schema Directory (%s)", strings.Join(components, "\n"))
}
