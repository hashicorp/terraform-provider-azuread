package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdSynchronizationJobId{}

// ServicePrincipalIdSynchronizationJobId is a struct representing the Resource ID for a Service Principal Id Synchronization Job
type ServicePrincipalIdSynchronizationJobId struct {
	ServicePrincipalId   string
	SynchronizationJobId string
}

// NewServicePrincipalIdSynchronizationJobID returns a new ServicePrincipalIdSynchronizationJobId struct
func NewServicePrincipalIdSynchronizationJobID(servicePrincipalId string, synchronizationJobId string) ServicePrincipalIdSynchronizationJobId {
	return ServicePrincipalIdSynchronizationJobId{
		ServicePrincipalId:   servicePrincipalId,
		SynchronizationJobId: synchronizationJobId,
	}
}

// ParseServicePrincipalIdSynchronizationJobID parses 'input' into a ServicePrincipalIdSynchronizationJobId
func ParseServicePrincipalIdSynchronizationJobID(input string) (*ServicePrincipalIdSynchronizationJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdSynchronizationJobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdSynchronizationJobId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdSynchronizationJobIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdSynchronizationJobId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdSynchronizationJobIDInsensitively(input string) (*ServicePrincipalIdSynchronizationJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdSynchronizationJobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdSynchronizationJobId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdSynchronizationJobId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.SynchronizationJobId, ok = input.Parsed["synchronizationJobId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "synchronizationJobId", input)
	}

	return nil
}

// ValidateServicePrincipalIdSynchronizationJobID checks that 'input' can be parsed as a Service Principal Id Synchronization Job ID
func ValidateServicePrincipalIdSynchronizationJobID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdSynchronizationJobID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Synchronization Job ID
func (id ServicePrincipalIdSynchronizationJobId) ID() string {
	fmtString := "/servicePrincipals/%s/synchronization/jobs/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.SynchronizationJobId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Synchronization Job ID
func (id ServicePrincipalIdSynchronizationJobId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("synchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("jobs", "jobs", "jobs"),
		resourceids.UserSpecifiedSegment("synchronizationJobId", "synchronizationJobId"),
	}
}

// String returns a human-readable description of this Service Principal Id Synchronization Job ID
func (id ServicePrincipalIdSynchronizationJobId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Synchronization Job: %q", id.SynchronizationJobId),
	}
	return fmt.Sprintf("Service Principal Id Synchronization Job (%s)", strings.Join(components, "\n"))
}
