package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdCreatedObjectId{}

// ServicePrincipalIdCreatedObjectId is a struct representing the Resource ID for a Service Principal Id Created Object
type ServicePrincipalIdCreatedObjectId struct {
	ServicePrincipalId string
	DirectoryObjectId  string
}

// NewServicePrincipalIdCreatedObjectID returns a new ServicePrincipalIdCreatedObjectId struct
func NewServicePrincipalIdCreatedObjectID(servicePrincipalId string, directoryObjectId string) ServicePrincipalIdCreatedObjectId {
	return ServicePrincipalIdCreatedObjectId{
		ServicePrincipalId: servicePrincipalId,
		DirectoryObjectId:  directoryObjectId,
	}
}

// ParseServicePrincipalIdCreatedObjectID parses 'input' into a ServicePrincipalIdCreatedObjectId
func ParseServicePrincipalIdCreatedObjectID(input string) (*ServicePrincipalIdCreatedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdCreatedObjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdCreatedObjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdCreatedObjectIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdCreatedObjectId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdCreatedObjectIDInsensitively(input string) (*ServicePrincipalIdCreatedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdCreatedObjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdCreatedObjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdCreatedObjectId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateServicePrincipalIdCreatedObjectID checks that 'input' can be parsed as a Service Principal Id Created Object ID
func ValidateServicePrincipalIdCreatedObjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdCreatedObjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Created Object ID
func (id ServicePrincipalIdCreatedObjectId) ID() string {
	fmtString := "/servicePrincipals/%s/createdObjects/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Created Object ID
func (id ServicePrincipalIdCreatedObjectId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("createdObjects", "createdObjects", "createdObjects"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Service Principal Id Created Object ID
func (id ServicePrincipalIdCreatedObjectId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Service Principal Id Created Object (%s)", strings.Join(components, "\n"))
}
