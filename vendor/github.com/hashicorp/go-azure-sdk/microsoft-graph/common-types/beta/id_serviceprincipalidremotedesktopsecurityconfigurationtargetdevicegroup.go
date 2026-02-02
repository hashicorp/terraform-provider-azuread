package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId{}

// ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId is a struct representing the Resource ID for a Service Principal Id Remote Desktop Security Configuration Target Device Group
type ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId struct {
	ServicePrincipalId  string
	TargetDeviceGroupId string
}

// NewServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupID returns a new ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId struct
func NewServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupID(servicePrincipalId string, targetDeviceGroupId string) ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId {
	return ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId{
		ServicePrincipalId:  servicePrincipalId,
		TargetDeviceGroupId: targetDeviceGroupId,
	}
}

// ParseServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupID parses 'input' into a ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId
func ParseServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupID(input string) (*ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupIDInsensitively(input string) (*ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.TargetDeviceGroupId, ok = input.Parsed["targetDeviceGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "targetDeviceGroupId", input)
	}

	return nil
}

// ValidateServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupID checks that 'input' can be parsed as a Service Principal Id Remote Desktop Security Configuration Target Device Group ID
func ValidateServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Remote Desktop Security Configuration Target Device Group ID
func (id ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId) ID() string {
	fmtString := "/servicePrincipals/%s/remoteDesktopSecurityConfiguration/targetDeviceGroups/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.TargetDeviceGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Remote Desktop Security Configuration Target Device Group ID
func (id ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
		resourceids.StaticSegment("remoteDesktopSecurityConfiguration", "remoteDesktopSecurityConfiguration", "remoteDesktopSecurityConfiguration"),
		resourceids.StaticSegment("targetDeviceGroups", "targetDeviceGroups", "targetDeviceGroups"),
		resourceids.UserSpecifiedSegment("targetDeviceGroupId", "targetDeviceGroupId"),
	}
}

// String returns a human-readable description of this Service Principal Id Remote Desktop Security Configuration Target Device Group ID
func (id ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Target Device Group: %q", id.TargetDeviceGroupId),
	}
	return fmt.Sprintf("Service Principal Id Remote Desktop Security Configuration Target Device Group (%s)", strings.Join(components, "\n"))
}
