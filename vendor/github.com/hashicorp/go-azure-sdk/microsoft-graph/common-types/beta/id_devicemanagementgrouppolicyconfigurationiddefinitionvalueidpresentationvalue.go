package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId{}

// DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId is a struct representing the Resource ID for a Device Management Group Policy Configuration Id Definition Value Id Presentation Value
type DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId struct {
	GroupPolicyConfigurationId     string
	GroupPolicyDefinitionValueId   string
	GroupPolicyPresentationValueId string
}

// NewDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueID returns a new DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId struct
func NewDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueID(groupPolicyConfigurationId string, groupPolicyDefinitionValueId string, groupPolicyPresentationValueId string) DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId {
	return DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId{
		GroupPolicyConfigurationId:     groupPolicyConfigurationId,
		GroupPolicyDefinitionValueId:   groupPolicyDefinitionValueId,
		GroupPolicyPresentationValueId: groupPolicyPresentationValueId,
	}
}

// ParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueID parses 'input' into a DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId
func ParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueID(input string) (*DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueIDInsensitively(input string) (*DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyConfigurationId, ok = input.Parsed["groupPolicyConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyConfigurationId", input)
	}

	if id.GroupPolicyDefinitionValueId, ok = input.Parsed["groupPolicyDefinitionValueId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyDefinitionValueId", input)
	}

	if id.GroupPolicyPresentationValueId, ok = input.Parsed["groupPolicyPresentationValueId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyPresentationValueId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueID checks that 'input' can be parsed as a Device Management Group Policy Configuration Id Definition Value Id Presentation Value ID
func ValidateDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Configuration Id Definition Value Id Presentation Value ID
func (id DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId) ID() string {
	fmtString := "/deviceManagement/groupPolicyConfigurations/%s/definitionValues/%s/presentationValues/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyConfigurationId, id.GroupPolicyDefinitionValueId, id.GroupPolicyPresentationValueId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Configuration Id Definition Value Id Presentation Value ID
func (id DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyConfigurations", "groupPolicyConfigurations", "groupPolicyConfigurations"),
		resourceids.UserSpecifiedSegment("groupPolicyConfigurationId", "groupPolicyConfigurationId"),
		resourceids.StaticSegment("definitionValues", "definitionValues", "definitionValues"),
		resourceids.UserSpecifiedSegment("groupPolicyDefinitionValueId", "groupPolicyDefinitionValueId"),
		resourceids.StaticSegment("presentationValues", "presentationValues", "presentationValues"),
		resourceids.UserSpecifiedSegment("groupPolicyPresentationValueId", "groupPolicyPresentationValueId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Configuration Id Definition Value Id Presentation Value ID
func (id DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Configuration: %q", id.GroupPolicyConfigurationId),
		fmt.Sprintf("Group Policy Definition Value: %q", id.GroupPolicyDefinitionValueId),
		fmt.Sprintf("Group Policy Presentation Value: %q", id.GroupPolicyPresentationValueId),
	}
	return fmt.Sprintf("Device Management Group Policy Configuration Id Definition Value Id Presentation Value (%s)", strings.Join(components, "\n"))
}
