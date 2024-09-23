package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyDefinitionIdPresentationId{}

// DeviceManagementGroupPolicyDefinitionIdPresentationId is a struct representing the Resource ID for a Device Management Group Policy Definition Id Presentation
type DeviceManagementGroupPolicyDefinitionIdPresentationId struct {
	GroupPolicyDefinitionId   string
	GroupPolicyPresentationId string
}

// NewDeviceManagementGroupPolicyDefinitionIdPresentationID returns a new DeviceManagementGroupPolicyDefinitionIdPresentationId struct
func NewDeviceManagementGroupPolicyDefinitionIdPresentationID(groupPolicyDefinitionId string, groupPolicyPresentationId string) DeviceManagementGroupPolicyDefinitionIdPresentationId {
	return DeviceManagementGroupPolicyDefinitionIdPresentationId{
		GroupPolicyDefinitionId:   groupPolicyDefinitionId,
		GroupPolicyPresentationId: groupPolicyPresentationId,
	}
}

// ParseDeviceManagementGroupPolicyDefinitionIdPresentationID parses 'input' into a DeviceManagementGroupPolicyDefinitionIdPresentationId
func ParseDeviceManagementGroupPolicyDefinitionIdPresentationID(input string) (*DeviceManagementGroupPolicyDefinitionIdPresentationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionIdPresentationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionIdPresentationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyDefinitionIdPresentationIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyDefinitionIdPresentationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyDefinitionIdPresentationIDInsensitively(input string) (*DeviceManagementGroupPolicyDefinitionIdPresentationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionIdPresentationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionIdPresentationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyDefinitionIdPresentationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyDefinitionId, ok = input.Parsed["groupPolicyDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyDefinitionId", input)
	}

	if id.GroupPolicyPresentationId, ok = input.Parsed["groupPolicyPresentationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyPresentationId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyDefinitionIdPresentationID checks that 'input' can be parsed as a Device Management Group Policy Definition Id Presentation ID
func ValidateDeviceManagementGroupPolicyDefinitionIdPresentationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyDefinitionIdPresentationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Definition Id Presentation ID
func (id DeviceManagementGroupPolicyDefinitionIdPresentationId) ID() string {
	fmtString := "/deviceManagement/groupPolicyDefinitions/%s/presentations/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyDefinitionId, id.GroupPolicyPresentationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Definition Id Presentation ID
func (id DeviceManagementGroupPolicyDefinitionIdPresentationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyDefinitions", "groupPolicyDefinitions", "groupPolicyDefinitions"),
		resourceids.UserSpecifiedSegment("groupPolicyDefinitionId", "groupPolicyDefinitionId"),
		resourceids.StaticSegment("presentations", "presentations", "presentations"),
		resourceids.UserSpecifiedSegment("groupPolicyPresentationId", "groupPolicyPresentationId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Definition Id Presentation ID
func (id DeviceManagementGroupPolicyDefinitionIdPresentationId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Definition: %q", id.GroupPolicyDefinitionId),
		fmt.Sprintf("Group Policy Presentation: %q", id.GroupPolicyPresentationId),
	}
	return fmt.Sprintf("Device Management Group Policy Definition Id Presentation (%s)", strings.Join(components, "\n"))
}
