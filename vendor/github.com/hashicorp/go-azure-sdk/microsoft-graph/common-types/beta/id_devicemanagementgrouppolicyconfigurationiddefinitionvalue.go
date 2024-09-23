package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyConfigurationIdDefinitionValueId{}

// DeviceManagementGroupPolicyConfigurationIdDefinitionValueId is a struct representing the Resource ID for a Device Management Group Policy Configuration Id Definition Value
type DeviceManagementGroupPolicyConfigurationIdDefinitionValueId struct {
	GroupPolicyConfigurationId   string
	GroupPolicyDefinitionValueId string
}

// NewDeviceManagementGroupPolicyConfigurationIdDefinitionValueID returns a new DeviceManagementGroupPolicyConfigurationIdDefinitionValueId struct
func NewDeviceManagementGroupPolicyConfigurationIdDefinitionValueID(groupPolicyConfigurationId string, groupPolicyDefinitionValueId string) DeviceManagementGroupPolicyConfigurationIdDefinitionValueId {
	return DeviceManagementGroupPolicyConfigurationIdDefinitionValueId{
		GroupPolicyConfigurationId:   groupPolicyConfigurationId,
		GroupPolicyDefinitionValueId: groupPolicyDefinitionValueId,
	}
}

// ParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueID parses 'input' into a DeviceManagementGroupPolicyConfigurationIdDefinitionValueId
func ParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueID(input string) (*DeviceManagementGroupPolicyConfigurationIdDefinitionValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyConfigurationIdDefinitionValueId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyConfigurationIdDefinitionValueId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyConfigurationIdDefinitionValueId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueIDInsensitively(input string) (*DeviceManagementGroupPolicyConfigurationIdDefinitionValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyConfigurationIdDefinitionValueId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyConfigurationIdDefinitionValueId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyConfigurationIdDefinitionValueId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyConfigurationId, ok = input.Parsed["groupPolicyConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyConfigurationId", input)
	}

	if id.GroupPolicyDefinitionValueId, ok = input.Parsed["groupPolicyDefinitionValueId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyDefinitionValueId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyConfigurationIdDefinitionValueID checks that 'input' can be parsed as a Device Management Group Policy Configuration Id Definition Value ID
func ValidateDeviceManagementGroupPolicyConfigurationIdDefinitionValueID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Configuration Id Definition Value ID
func (id DeviceManagementGroupPolicyConfigurationIdDefinitionValueId) ID() string {
	fmtString := "/deviceManagement/groupPolicyConfigurations/%s/definitionValues/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyConfigurationId, id.GroupPolicyDefinitionValueId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Configuration Id Definition Value ID
func (id DeviceManagementGroupPolicyConfigurationIdDefinitionValueId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyConfigurations", "groupPolicyConfigurations", "groupPolicyConfigurations"),
		resourceids.UserSpecifiedSegment("groupPolicyConfigurationId", "groupPolicyConfigurationId"),
		resourceids.StaticSegment("definitionValues", "definitionValues", "definitionValues"),
		resourceids.UserSpecifiedSegment("groupPolicyDefinitionValueId", "groupPolicyDefinitionValueId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Configuration Id Definition Value ID
func (id DeviceManagementGroupPolicyConfigurationIdDefinitionValueId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Configuration: %q", id.GroupPolicyConfigurationId),
		fmt.Sprintf("Group Policy Definition Value: %q", id.GroupPolicyDefinitionValueId),
	}
	return fmt.Sprintf("Device Management Group Policy Configuration Id Definition Value (%s)", strings.Join(components, "\n"))
}
