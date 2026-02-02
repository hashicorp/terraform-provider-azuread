package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId{}

// DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId is a struct representing the Resource ID for a Device Management Group Policy Definition Id Previous Version Definition Next Version Definition Presentation
type DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId struct {
	GroupPolicyDefinitionId   string
	GroupPolicyPresentationId string
}

// NewDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationID returns a new DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId struct
func NewDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationID(groupPolicyDefinitionId string, groupPolicyPresentationId string) DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId {
	return DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId{
		GroupPolicyDefinitionId:   groupPolicyDefinitionId,
		GroupPolicyPresentationId: groupPolicyPresentationId,
	}
}

// ParseDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationID parses 'input' into a DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId
func ParseDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationID(input string) (*DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationIDInsensitively(input string) (*DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyDefinitionId, ok = input.Parsed["groupPolicyDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyDefinitionId", input)
	}

	if id.GroupPolicyPresentationId, ok = input.Parsed["groupPolicyPresentationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyPresentationId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationID checks that 'input' can be parsed as a Device Management Group Policy Definition Id Previous Version Definition Next Version Definition Presentation ID
func ValidateDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Definition Id Previous Version Definition Next Version Definition Presentation ID
func (id DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId) ID() string {
	fmtString := "/deviceManagement/groupPolicyDefinitions/%s/previousVersionDefinition/nextVersionDefinition/presentations/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyDefinitionId, id.GroupPolicyPresentationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Definition Id Previous Version Definition Next Version Definition Presentation ID
func (id DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyDefinitions", "groupPolicyDefinitions", "groupPolicyDefinitions"),
		resourceids.UserSpecifiedSegment("groupPolicyDefinitionId", "groupPolicyDefinitionId"),
		resourceids.StaticSegment("previousVersionDefinition", "previousVersionDefinition", "previousVersionDefinition"),
		resourceids.StaticSegment("nextVersionDefinition", "nextVersionDefinition", "nextVersionDefinition"),
		resourceids.StaticSegment("presentations", "presentations", "presentations"),
		resourceids.UserSpecifiedSegment("groupPolicyPresentationId", "groupPolicyPresentationId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Definition Id Previous Version Definition Next Version Definition Presentation ID
func (id DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Definition: %q", id.GroupPolicyDefinitionId),
		fmt.Sprintf("Group Policy Presentation: %q", id.GroupPolicyPresentationId),
	}
	return fmt.Sprintf("Device Management Group Policy Definition Id Previous Version Definition Next Version Definition Presentation (%s)", strings.Join(components, "\n"))
}
