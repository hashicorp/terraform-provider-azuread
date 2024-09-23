package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAppRoleAssignedResourceId{}

// MeAppRoleAssignedResourceId is a struct representing the Resource ID for a Me App Role Assigned Resource
type MeAppRoleAssignedResourceId struct {
	ServicePrincipalId string
}

// NewMeAppRoleAssignedResourceID returns a new MeAppRoleAssignedResourceId struct
func NewMeAppRoleAssignedResourceID(servicePrincipalId string) MeAppRoleAssignedResourceId {
	return MeAppRoleAssignedResourceId{
		ServicePrincipalId: servicePrincipalId,
	}
}

// ParseMeAppRoleAssignedResourceID parses 'input' into a MeAppRoleAssignedResourceId
func ParseMeAppRoleAssignedResourceID(input string) (*MeAppRoleAssignedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAppRoleAssignedResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAppRoleAssignedResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAppRoleAssignedResourceIDInsensitively parses 'input' case-insensitively into a MeAppRoleAssignedResourceId
// note: this method should only be used for API response data and not user input
func ParseMeAppRoleAssignedResourceIDInsensitively(input string) (*MeAppRoleAssignedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAppRoleAssignedResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAppRoleAssignedResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAppRoleAssignedResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	return nil
}

// ValidateMeAppRoleAssignedResourceID checks that 'input' can be parsed as a Me App Role Assigned Resource ID
func ValidateMeAppRoleAssignedResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAppRoleAssignedResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me App Role Assigned Resource ID
func (id MeAppRoleAssignedResourceId) ID() string {
	fmtString := "/me/appRoleAssignedResources/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me App Role Assigned Resource ID
func (id MeAppRoleAssignedResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("appRoleAssignedResources", "appRoleAssignedResources", "appRoleAssignedResources"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalId"),
	}
}

// String returns a human-readable description of this Me App Role Assigned Resource ID
func (id MeAppRoleAssignedResourceId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
	}
	return fmt.Sprintf("Me App Role Assigned Resource (%s)", strings.Join(components, "\n"))
}
