package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdEndpointId{}

// ServicePrincipalIdEndpointId is a struct representing the Resource ID for a Service Principal Id Endpoint
type ServicePrincipalIdEndpointId struct {
	ServicePrincipalId string
	EndpointId         string
}

// NewServicePrincipalIdEndpointID returns a new ServicePrincipalIdEndpointId struct
func NewServicePrincipalIdEndpointID(servicePrincipalId string, endpointId string) ServicePrincipalIdEndpointId {
	return ServicePrincipalIdEndpointId{
		ServicePrincipalId: servicePrincipalId,
		EndpointId:         endpointId,
	}
}

// ParseServicePrincipalIdEndpointID parses 'input' into a ServicePrincipalIdEndpointId
func ParseServicePrincipalIdEndpointID(input string) (*ServicePrincipalIdEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdEndpointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdEndpointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdEndpointIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdEndpointId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdEndpointIDInsensitively(input string) (*ServicePrincipalIdEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdEndpointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdEndpointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdEndpointId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.EndpointId, ok = input.Parsed["endpointId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "endpointId", input)
	}

	return nil
}

// ValidateServicePrincipalIdEndpointID checks that 'input' can be parsed as a Service Principal Id Endpoint ID
func ValidateServicePrincipalIdEndpointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdEndpointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Endpoint ID
func (id ServicePrincipalIdEndpointId) ID() string {
	fmtString := "/servicePrincipals/%s/endpoints/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.EndpointId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Endpoint ID
func (id ServicePrincipalIdEndpointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("endpoints", "endpoints", "endpoints"),
		resourceids.UserSpecifiedSegment("endpointId", "endpointId"),
	}
}

// String returns a human-readable description of this Service Principal Id Endpoint ID
func (id ServicePrincipalIdEndpointId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Endpoint: %q", id.EndpointId),
	}
	return fmt.Sprintf("Service Principal Id Endpoint (%s)", strings.Join(components, "\n"))
}
