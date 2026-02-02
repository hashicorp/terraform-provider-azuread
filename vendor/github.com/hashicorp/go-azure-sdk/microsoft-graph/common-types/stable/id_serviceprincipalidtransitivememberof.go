package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdTransitiveMemberOfId{}

// ServicePrincipalIdTransitiveMemberOfId is a struct representing the Resource ID for a Service Principal Id Transitive Member Of
type ServicePrincipalIdTransitiveMemberOfId struct {
	ServicePrincipalId string
	DirectoryObjectId  string
}

// NewServicePrincipalIdTransitiveMemberOfID returns a new ServicePrincipalIdTransitiveMemberOfId struct
func NewServicePrincipalIdTransitiveMemberOfID(servicePrincipalId string, directoryObjectId string) ServicePrincipalIdTransitiveMemberOfId {
	return ServicePrincipalIdTransitiveMemberOfId{
		ServicePrincipalId: servicePrincipalId,
		DirectoryObjectId:  directoryObjectId,
	}
}

// ParseServicePrincipalIdTransitiveMemberOfID parses 'input' into a ServicePrincipalIdTransitiveMemberOfId
func ParseServicePrincipalIdTransitiveMemberOfID(input string) (*ServicePrincipalIdTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdTransitiveMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdTransitiveMemberOfIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdTransitiveMemberOfId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdTransitiveMemberOfIDInsensitively(input string) (*ServicePrincipalIdTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdTransitiveMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdTransitiveMemberOfId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateServicePrincipalIdTransitiveMemberOfID checks that 'input' can be parsed as a Service Principal Id Transitive Member Of ID
func ValidateServicePrincipalIdTransitiveMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdTransitiveMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Transitive Member Of ID
func (id ServicePrincipalIdTransitiveMemberOfId) ID() string {
	fmtString := "/servicePrincipals/%s/transitiveMemberOf/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Transitive Member Of ID
func (id ServicePrincipalIdTransitiveMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("transitiveMemberOf", "transitiveMemberOf", "transitiveMemberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Service Principal Id Transitive Member Of ID
func (id ServicePrincipalIdTransitiveMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Service Principal Id Transitive Member Of (%s)", strings.Join(components, "\n"))
}
