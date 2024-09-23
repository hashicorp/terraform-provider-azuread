package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdMemberOfId{}

// ServicePrincipalIdMemberOfId is a struct representing the Resource ID for a Service Principal Id Member Of
type ServicePrincipalIdMemberOfId struct {
	ServicePrincipalId string
	DirectoryObjectId  string
}

// NewServicePrincipalIdMemberOfID returns a new ServicePrincipalIdMemberOfId struct
func NewServicePrincipalIdMemberOfID(servicePrincipalId string, directoryObjectId string) ServicePrincipalIdMemberOfId {
	return ServicePrincipalIdMemberOfId{
		ServicePrincipalId: servicePrincipalId,
		DirectoryObjectId:  directoryObjectId,
	}
}

// ParseServicePrincipalIdMemberOfID parses 'input' into a ServicePrincipalIdMemberOfId
func ParseServicePrincipalIdMemberOfID(input string) (*ServicePrincipalIdMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdMemberOfIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdMemberOfId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdMemberOfIDInsensitively(input string) (*ServicePrincipalIdMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdMemberOfId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateServicePrincipalIdMemberOfID checks that 'input' can be parsed as a Service Principal Id Member Of ID
func ValidateServicePrincipalIdMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Member Of ID
func (id ServicePrincipalIdMemberOfId) ID() string {
	fmtString := "/servicePrincipals/%s/memberOf/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Member Of ID
func (id ServicePrincipalIdMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("memberOf", "memberOf", "memberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Service Principal Id Member Of ID
func (id ServicePrincipalIdMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Service Principal Id Member Of (%s)", strings.Join(components, "\n"))
}
