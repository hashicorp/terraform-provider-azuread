package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId{}

// ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId is a struct representing the Resource ID for a Service Principal Id Synchronization Template Id Schema Directory
type ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId struct {
	ServicePrincipalId        string
	SynchronizationTemplateId string
	DirectoryDefinitionId     string
}

// NewServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID returns a new ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId struct
func NewServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID(servicePrincipalId string, synchronizationTemplateId string, directoryDefinitionId string) ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId {
	return ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId{
		ServicePrincipalId:        servicePrincipalId,
		SynchronizationTemplateId: synchronizationTemplateId,
		DirectoryDefinitionId:     directoryDefinitionId,
	}
}

// ParseServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID parses 'input' into a ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId
func ParseServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID(input string) (*ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryIDInsensitively(input string) (*ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.SynchronizationTemplateId, ok = input.Parsed["synchronizationTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "synchronizationTemplateId", input)
	}

	if id.DirectoryDefinitionId, ok = input.Parsed["directoryDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryDefinitionId", input)
	}

	return nil
}

// ValidateServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID checks that 'input' can be parsed as a Service Principal Id Synchronization Template Id Schema Directory ID
func ValidateServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Synchronization Template Id Schema Directory ID
func (id ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId) ID() string {
	fmtString := "/servicePrincipals/%s/synchronization/templates/%s/schema/directories/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.SynchronizationTemplateId, id.DirectoryDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Synchronization Template Id Schema Directory ID
func (id ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("synchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("synchronizationTemplateId", "synchronizationTemplateId"),
		resourceids.StaticSegment("schema", "schema", "schema"),
		resourceids.StaticSegment("directories", "directories", "directories"),
		resourceids.UserSpecifiedSegment("directoryDefinitionId", "directoryDefinitionId"),
	}
}

// String returns a human-readable description of this Service Principal Id Synchronization Template Id Schema Directory ID
func (id ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Synchronization Template: %q", id.SynchronizationTemplateId),
		fmt.Sprintf("Directory Definition: %q", id.DirectoryDefinitionId),
	}
	return fmt.Sprintf("Service Principal Id Synchronization Template Id Schema Directory (%s)", strings.Join(components, "\n"))
}
