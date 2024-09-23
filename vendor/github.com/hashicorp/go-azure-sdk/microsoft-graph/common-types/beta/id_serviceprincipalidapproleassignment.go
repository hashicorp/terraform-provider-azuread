package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdAppRoleAssignmentId{}

// ServicePrincipalIdAppRoleAssignmentId is a struct representing the Resource ID for a Service Principal Id App Role Assignment
type ServicePrincipalIdAppRoleAssignmentId struct {
	ServicePrincipalId  string
	AppRoleAssignmentId string
}

// NewServicePrincipalIdAppRoleAssignmentID returns a new ServicePrincipalIdAppRoleAssignmentId struct
func NewServicePrincipalIdAppRoleAssignmentID(servicePrincipalId string, appRoleAssignmentId string) ServicePrincipalIdAppRoleAssignmentId {
	return ServicePrincipalIdAppRoleAssignmentId{
		ServicePrincipalId:  servicePrincipalId,
		AppRoleAssignmentId: appRoleAssignmentId,
	}
}

// ParseServicePrincipalIdAppRoleAssignmentID parses 'input' into a ServicePrincipalIdAppRoleAssignmentId
func ParseServicePrincipalIdAppRoleAssignmentID(input string) (*ServicePrincipalIdAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdAppRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdAppRoleAssignmentIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdAppRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdAppRoleAssignmentIDInsensitively(input string) (*ServicePrincipalIdAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdAppRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdAppRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.AppRoleAssignmentId, ok = input.Parsed["appRoleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appRoleAssignmentId", input)
	}

	return nil
}

// ValidateServicePrincipalIdAppRoleAssignmentID checks that 'input' can be parsed as a Service Principal Id App Role Assignment ID
func ValidateServicePrincipalIdAppRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdAppRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id App Role Assignment ID
func (id ServicePrincipalIdAppRoleAssignmentId) ID() string {
	fmtString := "/servicePrincipals/%s/appRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.AppRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id App Role Assignment ID
func (id ServicePrincipalIdAppRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("appRoleAssignments", "appRoleAssignments", "appRoleAssignments"),
		resourceids.UserSpecifiedSegment("appRoleAssignmentId", "appRoleAssignmentId"),
	}
}

// String returns a human-readable description of this Service Principal Id App Role Assignment ID
func (id ServicePrincipalIdAppRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("App Role Assignment: %q", id.AppRoleAssignmentId),
	}
	return fmt.Sprintf("Service Principal Id App Role Assignment (%s)", strings.Join(components, "\n"))
}
