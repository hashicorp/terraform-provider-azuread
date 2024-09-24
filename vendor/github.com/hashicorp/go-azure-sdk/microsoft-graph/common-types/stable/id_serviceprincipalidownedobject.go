package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdOwnedObjectId{}

// ServicePrincipalIdOwnedObjectId is a struct representing the Resource ID for a Service Principal Id Owned Object
type ServicePrincipalIdOwnedObjectId struct {
	ServicePrincipalId string
	DirectoryObjectId  string
}

// NewServicePrincipalIdOwnedObjectID returns a new ServicePrincipalIdOwnedObjectId struct
func NewServicePrincipalIdOwnedObjectID(servicePrincipalId string, directoryObjectId string) ServicePrincipalIdOwnedObjectId {
	return ServicePrincipalIdOwnedObjectId{
		ServicePrincipalId: servicePrincipalId,
		DirectoryObjectId:  directoryObjectId,
	}
}

// ParseServicePrincipalIdOwnedObjectID parses 'input' into a ServicePrincipalIdOwnedObjectId
func ParseServicePrincipalIdOwnedObjectID(input string) (*ServicePrincipalIdOwnedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdOwnedObjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdOwnedObjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdOwnedObjectIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdOwnedObjectId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdOwnedObjectIDInsensitively(input string) (*ServicePrincipalIdOwnedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdOwnedObjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdOwnedObjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdOwnedObjectId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateServicePrincipalIdOwnedObjectID checks that 'input' can be parsed as a Service Principal Id Owned Object ID
func ValidateServicePrincipalIdOwnedObjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdOwnedObjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Owned Object ID
func (id ServicePrincipalIdOwnedObjectId) ID() string {
	fmtString := "/servicePrincipals/%s/ownedObjects/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Owned Object ID
func (id ServicePrincipalIdOwnedObjectId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("ownedObjects", "ownedObjects", "ownedObjects"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Service Principal Id Owned Object ID
func (id ServicePrincipalIdOwnedObjectId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Service Principal Id Owned Object (%s)", strings.Join(components, "\n"))
}
