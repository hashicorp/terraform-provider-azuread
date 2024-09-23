package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId{}

// DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId is a struct representing the Resource ID for a Device Management Group Policy Definition Id Next Version Definition Previous Version Definition Presentation
type DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId struct {
	GroupPolicyDefinitionId   string
	GroupPolicyPresentationId string
}

// NewDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID returns a new DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId struct
func NewDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID(groupPolicyDefinitionId string, groupPolicyPresentationId string) DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId {
	return DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId{
		GroupPolicyDefinitionId:   groupPolicyDefinitionId,
		GroupPolicyPresentationId: groupPolicyPresentationId,
	}
}

// ParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID parses 'input' into a DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId
func ParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID(input string) (*DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationIDInsensitively(input string) (*DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyDefinitionId, ok = input.Parsed["groupPolicyDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyDefinitionId", input)
	}

	if id.GroupPolicyPresentationId, ok = input.Parsed["groupPolicyPresentationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyPresentationId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID checks that 'input' can be parsed as a Device Management Group Policy Definition Id Next Version Definition Previous Version Definition Presentation ID
func ValidateDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Definition Id Next Version Definition Previous Version Definition Presentation ID
func (id DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId) ID() string {
	fmtString := "/deviceManagement/groupPolicyDefinitions/%s/nextVersionDefinition/previousVersionDefinition/presentations/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyDefinitionId, id.GroupPolicyPresentationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Definition Id Next Version Definition Previous Version Definition Presentation ID
func (id DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyDefinitions", "groupPolicyDefinitions", "groupPolicyDefinitions"),
		resourceids.UserSpecifiedSegment("groupPolicyDefinitionId", "groupPolicyDefinitionId"),
		resourceids.StaticSegment("nextVersionDefinition", "nextVersionDefinition", "nextVersionDefinition"),
		resourceids.StaticSegment("previousVersionDefinition", "previousVersionDefinition", "previousVersionDefinition"),
		resourceids.StaticSegment("presentations", "presentations", "presentations"),
		resourceids.UserSpecifiedSegment("groupPolicyPresentationId", "groupPolicyPresentationId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Definition Id Next Version Definition Previous Version Definition Presentation ID
func (id DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Definition: %q", id.GroupPolicyDefinitionId),
		fmt.Sprintf("Group Policy Presentation: %q", id.GroupPolicyPresentationId),
	}
	return fmt.Sprintf("Device Management Group Policy Definition Id Next Version Definition Previous Version Definition Presentation (%s)", strings.Join(components, "\n"))
}
