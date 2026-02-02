package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdOwnerId{}

// ServicePrincipalIdOwnerId is a struct representing the Resource ID for a Service Principal Id Owner
type ServicePrincipalIdOwnerId struct {
	ServicePrincipalId string
	DirectoryObjectId  string
}

// NewServicePrincipalIdOwnerID returns a new ServicePrincipalIdOwnerId struct
func NewServicePrincipalIdOwnerID(servicePrincipalId string, directoryObjectId string) ServicePrincipalIdOwnerId {
	return ServicePrincipalIdOwnerId{
		ServicePrincipalId: servicePrincipalId,
		DirectoryObjectId:  directoryObjectId,
	}
}

// ParseServicePrincipalIdOwnerID parses 'input' into a ServicePrincipalIdOwnerId
func ParseServicePrincipalIdOwnerID(input string) (*ServicePrincipalIdOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdOwnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdOwnerIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdOwnerId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdOwnerIDInsensitively(input string) (*ServicePrincipalIdOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdOwnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdOwnerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateServicePrincipalIdOwnerID checks that 'input' can be parsed as a Service Principal Id Owner ID
func ValidateServicePrincipalIdOwnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdOwnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Owner ID
func (id ServicePrincipalIdOwnerId) ID() string {
	fmtString := "/servicePrincipals/%s/owners/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Owner ID
func (id ServicePrincipalIdOwnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("owners", "owners", "owners"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Service Principal Id Owner ID
func (id ServicePrincipalIdOwnerId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Service Principal Id Owner (%s)", strings.Join(components, "\n"))
}
