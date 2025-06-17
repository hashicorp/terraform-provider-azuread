package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId{}

// ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId is a struct representing the Resource ID for a Service Principal Id Remote Desktop Security Configuration Approved Client App
type ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId struct {
	ServicePrincipalId  string
	ApprovedClientAppId string
}

// NewServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppID returns a new ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId struct
func NewServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppID(servicePrincipalId string, approvedClientAppId string) ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId {
	return ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId{
		ServicePrincipalId:  servicePrincipalId,
		ApprovedClientAppId: approvedClientAppId,
	}
}

// ParseServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppID parses 'input' into a ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId
func ParseServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppID(input string) (*ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppIDInsensitively(input string) (*ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.ApprovedClientAppId, ok = input.Parsed["approvedClientAppId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvedClientAppId", input)
	}

	return nil
}

// ValidateServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppID checks that 'input' can be parsed as a Service Principal Id Remote Desktop Security Configuration Approved Client App ID
func ValidateServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Remote Desktop Security Configuration Approved Client App ID
func (id ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId) ID() string {
	fmtString := "/servicePrincipals/%s/remoteDesktopSecurityConfiguration/approvedClientApps/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.ApprovedClientAppId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Remote Desktop Security Configuration Approved Client App ID
func (id ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("remoteDesktopSecurityConfiguration", "remoteDesktopSecurityConfiguration", "remoteDesktopSecurityConfiguration"),
		resourceids.StaticSegment("approvedClientApps", "approvedClientApps", "approvedClientApps"),
		resourceids.UserSpecifiedSegment("approvedClientAppId", "approvedClientAppId"),
	}
}

// String returns a human-readable description of this Service Principal Id Remote Desktop Security Configuration Approved Client App ID
func (id ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Approved Client App: %q", id.ApprovedClientAppId),
	}
	return fmt.Sprintf("Service Principal Id Remote Desktop Security Configuration Approved Client App (%s)", strings.Join(components, "\n"))
}
