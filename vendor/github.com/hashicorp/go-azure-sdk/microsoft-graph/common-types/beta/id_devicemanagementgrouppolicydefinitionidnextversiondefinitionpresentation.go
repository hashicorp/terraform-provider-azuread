package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId{}

// DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId is a struct representing the Resource ID for a Device Management Group Policy Definition Id Next Version Definition Presentation
type DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId struct {
	GroupPolicyDefinitionId   string
	GroupPolicyPresentationId string
}

// NewDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationID returns a new DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId struct
func NewDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationID(groupPolicyDefinitionId string, groupPolicyPresentationId string) DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId {
	return DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId{
		GroupPolicyDefinitionId:   groupPolicyDefinitionId,
		GroupPolicyPresentationId: groupPolicyPresentationId,
	}
}

// ParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationID parses 'input' into a DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId
func ParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationID(input string) (*DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationIDInsensitively(input string) (*DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyDefinitionId, ok = input.Parsed["groupPolicyDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyDefinitionId", input)
	}

	if id.GroupPolicyPresentationId, ok = input.Parsed["groupPolicyPresentationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyPresentationId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationID checks that 'input' can be parsed as a Device Management Group Policy Definition Id Next Version Definition Presentation ID
func ValidateDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Definition Id Next Version Definition Presentation ID
func (id DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId) ID() string {
	fmtString := "/deviceManagement/groupPolicyDefinitions/%s/nextVersionDefinition/presentations/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyDefinitionId, id.GroupPolicyPresentationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Definition Id Next Version Definition Presentation ID
func (id DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyDefinitions", "groupPolicyDefinitions", "groupPolicyDefinitions"),
		resourceids.UserSpecifiedSegment("groupPolicyDefinitionId", "groupPolicyDefinitionId"),
		resourceids.StaticSegment("nextVersionDefinition", "nextVersionDefinition", "nextVersionDefinition"),
		resourceids.StaticSegment("presentations", "presentations", "presentations"),
		resourceids.UserSpecifiedSegment("groupPolicyPresentationId", "groupPolicyPresentationId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Definition Id Next Version Definition Presentation ID
func (id DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPresentationId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Definition: %q", id.GroupPolicyDefinitionId),
		fmt.Sprintf("Group Policy Presentation: %q", id.GroupPolicyPresentationId),
	}
	return fmt.Sprintf("Device Management Group Policy Definition Id Next Version Definition Presentation (%s)", strings.Join(components, "\n"))
}
