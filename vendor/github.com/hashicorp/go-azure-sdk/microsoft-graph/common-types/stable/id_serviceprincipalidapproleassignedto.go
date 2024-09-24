package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdAppRoleAssignedToId{}

// ServicePrincipalIdAppRoleAssignedToId is a struct representing the Resource ID for a Service Principal Id App Role Assigned To
type ServicePrincipalIdAppRoleAssignedToId struct {
	ServicePrincipalId  string
	AppRoleAssignmentId string
}

// NewServicePrincipalIdAppRoleAssignedToID returns a new ServicePrincipalIdAppRoleAssignedToId struct
func NewServicePrincipalIdAppRoleAssignedToID(servicePrincipalId string, appRoleAssignmentId string) ServicePrincipalIdAppRoleAssignedToId {
	return ServicePrincipalIdAppRoleAssignedToId{
		ServicePrincipalId:  servicePrincipalId,
		AppRoleAssignmentId: appRoleAssignmentId,
	}
}

// ParseServicePrincipalIdAppRoleAssignedToID parses 'input' into a ServicePrincipalIdAppRoleAssignedToId
func ParseServicePrincipalIdAppRoleAssignedToID(input string) (*ServicePrincipalIdAppRoleAssignedToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdAppRoleAssignedToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdAppRoleAssignedToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdAppRoleAssignedToIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdAppRoleAssignedToId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdAppRoleAssignedToIDInsensitively(input string) (*ServicePrincipalIdAppRoleAssignedToId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdAppRoleAssignedToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdAppRoleAssignedToId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdAppRoleAssignedToId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.AppRoleAssignmentId, ok = input.Parsed["appRoleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appRoleAssignmentId", input)
	}

	return nil
}

// ValidateServicePrincipalIdAppRoleAssignedToID checks that 'input' can be parsed as a Service Principal Id App Role Assigned To ID
func ValidateServicePrincipalIdAppRoleAssignedToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdAppRoleAssignedToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id App Role Assigned To ID
func (id ServicePrincipalIdAppRoleAssignedToId) ID() string {
	fmtString := "/servicePrincipals/%s/appRoleAssignedTo/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.AppRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id App Role Assigned To ID
func (id ServicePrincipalIdAppRoleAssignedToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("appRoleAssignedTo", "appRoleAssignedTo", "appRoleAssignedTo"),
		resourceids.UserSpecifiedSegment("appRoleAssignmentId", "appRoleAssignmentId"),
	}
}

// String returns a human-readable description of this Service Principal Id App Role Assigned To ID
func (id ServicePrincipalIdAppRoleAssignedToId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("App Role Assignment: %q", id.AppRoleAssignmentId),
	}
	return fmt.Sprintf("Service Principal Id App Role Assigned To (%s)", strings.Join(components, "\n"))
}
