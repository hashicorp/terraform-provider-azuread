package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyDefinitionId{}

// DeviceManagementGroupPolicyDefinitionId is a struct representing the Resource ID for a Device Management Group Policy Definition
type DeviceManagementGroupPolicyDefinitionId struct {
	GroupPolicyDefinitionId string
}

// NewDeviceManagementGroupPolicyDefinitionID returns a new DeviceManagementGroupPolicyDefinitionId struct
func NewDeviceManagementGroupPolicyDefinitionID(groupPolicyDefinitionId string) DeviceManagementGroupPolicyDefinitionId {
	return DeviceManagementGroupPolicyDefinitionId{
		GroupPolicyDefinitionId: groupPolicyDefinitionId,
	}
}

// ParseDeviceManagementGroupPolicyDefinitionID parses 'input' into a DeviceManagementGroupPolicyDefinitionId
func ParseDeviceManagementGroupPolicyDefinitionID(input string) (*DeviceManagementGroupPolicyDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyDefinitionIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyDefinitionIDInsensitively(input string) (*DeviceManagementGroupPolicyDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyDefinitionId, ok = input.Parsed["groupPolicyDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyDefinitionID checks that 'input' can be parsed as a Device Management Group Policy Definition ID
func ValidateDeviceManagementGroupPolicyDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Definition ID
func (id DeviceManagementGroupPolicyDefinitionId) ID() string {
	fmtString := "/deviceManagement/groupPolicyDefinitions/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Definition ID
func (id DeviceManagementGroupPolicyDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyDefinitions", "groupPolicyDefinitions", "groupPolicyDefinitions"),
		resourceids.UserSpecifiedSegment("groupPolicyDefinitionId", "groupPolicyDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Definition ID
func (id DeviceManagementGroupPolicyDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Definition: %q", id.GroupPolicyDefinitionId),
	}
	return fmt.Sprintf("Device Management Group Policy Definition (%s)", strings.Join(components, "\n"))
}
