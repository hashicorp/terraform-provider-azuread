package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdOAuth2PermissionGrantId{}

// ServicePrincipalIdOAuth2PermissionGrantId is a struct representing the Resource ID for a Service Principal Id O Auth 2 Permission Grant
type ServicePrincipalIdOAuth2PermissionGrantId struct {
	ServicePrincipalId      string
	OAuth2PermissionGrantId string
}

// NewServicePrincipalIdOAuth2PermissionGrantID returns a new ServicePrincipalIdOAuth2PermissionGrantId struct
func NewServicePrincipalIdOAuth2PermissionGrantID(servicePrincipalId string, oAuth2PermissionGrantId string) ServicePrincipalIdOAuth2PermissionGrantId {
	return ServicePrincipalIdOAuth2PermissionGrantId{
		ServicePrincipalId:      servicePrincipalId,
		OAuth2PermissionGrantId: oAuth2PermissionGrantId,
	}
}

// ParseServicePrincipalIdOAuth2PermissionGrantID parses 'input' into a ServicePrincipalIdOAuth2PermissionGrantId
func ParseServicePrincipalIdOAuth2PermissionGrantID(input string) (*ServicePrincipalIdOAuth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdOAuth2PermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdOAuth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdOAuth2PermissionGrantIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdOAuth2PermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdOAuth2PermissionGrantIDInsensitively(input string) (*ServicePrincipalIdOAuth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdOAuth2PermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdOAuth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdOAuth2PermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.OAuth2PermissionGrantId, ok = input.Parsed["oAuth2PermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "oAuth2PermissionGrantId", input)
	}

	return nil
}

// ValidateServicePrincipalIdOAuth2PermissionGrantID checks that 'input' can be parsed as a Service Principal Id O Auth 2 Permission Grant ID
func ValidateServicePrincipalIdOAuth2PermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdOAuth2PermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id O Auth 2 Permission Grant ID
func (id ServicePrincipalIdOAuth2PermissionGrantId) ID() string {
	fmtString := "/servicePrincipals/%s/oauth2PermissionGrants/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.OAuth2PermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id O Auth 2 Permission Grant ID
func (id ServicePrincipalIdOAuth2PermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("oauth2PermissionGrants", "oauth2PermissionGrants", "oauth2PermissionGrants"),
		resourceids.UserSpecifiedSegment("oAuth2PermissionGrantId", "oAuth2PermissionGrantId"),
	}
}

// String returns a human-readable description of this Service Principal Id O Auth 2 Permission Grant ID
func (id ServicePrincipalIdOAuth2PermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("O Auth 2 Permission Grant: %q", id.OAuth2PermissionGrantId),
	}
	return fmt.Sprintf("Service Principal Id O Auth 2 Permission Grant (%s)", strings.Join(components, "\n"))
}
