package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdSynchronizationTemplateId{}

// ServicePrincipalIdSynchronizationTemplateId is a struct representing the Resource ID for a Service Principal Id Synchronization Template
type ServicePrincipalIdSynchronizationTemplateId struct {
	ServicePrincipalId        string
	SynchronizationTemplateId string
}

// NewServicePrincipalIdSynchronizationTemplateID returns a new ServicePrincipalIdSynchronizationTemplateId struct
func NewServicePrincipalIdSynchronizationTemplateID(servicePrincipalId string, synchronizationTemplateId string) ServicePrincipalIdSynchronizationTemplateId {
	return ServicePrincipalIdSynchronizationTemplateId{
		ServicePrincipalId:        servicePrincipalId,
		SynchronizationTemplateId: synchronizationTemplateId,
	}
}

// ParseServicePrincipalIdSynchronizationTemplateID parses 'input' into a ServicePrincipalIdSynchronizationTemplateId
func ParseServicePrincipalIdSynchronizationTemplateID(input string) (*ServicePrincipalIdSynchronizationTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdSynchronizationTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdSynchronizationTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdSynchronizationTemplateIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdSynchronizationTemplateId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdSynchronizationTemplateIDInsensitively(input string) (*ServicePrincipalIdSynchronizationTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdSynchronizationTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdSynchronizationTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdSynchronizationTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.SynchronizationTemplateId, ok = input.Parsed["synchronizationTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "synchronizationTemplateId", input)
	}

	return nil
}

// ValidateServicePrincipalIdSynchronizationTemplateID checks that 'input' can be parsed as a Service Principal Id Synchronization Template ID
func ValidateServicePrincipalIdSynchronizationTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdSynchronizationTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Synchronization Template ID
func (id ServicePrincipalIdSynchronizationTemplateId) ID() string {
	fmtString := "/servicePrincipals/%s/synchronization/templates/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.SynchronizationTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Synchronization Template ID
func (id ServicePrincipalIdSynchronizationTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("synchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("synchronizationTemplateId", "synchronizationTemplateId"),
	}
}

// String returns a human-readable description of this Service Principal Id Synchronization Template ID
func (id ServicePrincipalIdSynchronizationTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Synchronization Template: %q", id.SynchronizationTemplateId),
	}
	return fmt.Sprintf("Service Principal Id Synchronization Template (%s)", strings.Join(components, "\n"))
}
