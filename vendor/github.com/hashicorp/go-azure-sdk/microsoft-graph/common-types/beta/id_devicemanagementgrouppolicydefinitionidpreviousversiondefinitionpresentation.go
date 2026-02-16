package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId{}

// DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId is a struct representing the Resource ID for a Device Management Group Policy Definition Id Previous Version Definition Presentation
type DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId struct {
	GroupPolicyDefinitionId   string
	GroupPolicyPresentationId string
}

// NewDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationID returns a new DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId struct
func NewDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationID(groupPolicyDefinitionId string, groupPolicyPresentationId string) DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId {
	return DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId{
		GroupPolicyDefinitionId:   groupPolicyDefinitionId,
		GroupPolicyPresentationId: groupPolicyPresentationId,
	}
}

// ParseDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationID parses 'input' into a DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId
func ParseDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationID(input string) (*DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationIDInsensitively(input string) (*DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyDefinitionId, ok = input.Parsed["groupPolicyDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyDefinitionId", input)
	}

	if id.GroupPolicyPresentationId, ok = input.Parsed["groupPolicyPresentationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyPresentationId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationID checks that 'input' can be parsed as a Device Management Group Policy Definition Id Previous Version Definition Presentation ID
func ValidateDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Definition Id Previous Version Definition Presentation ID
func (id DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId) ID() string {
	fmtString := "/deviceManagement/groupPolicyDefinitions/%s/previousVersionDefinition/presentations/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyDefinitionId, id.GroupPolicyPresentationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Definition Id Previous Version Definition Presentation ID
func (id DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyDefinitions", "groupPolicyDefinitions", "groupPolicyDefinitions"),
		resourceids.UserSpecifiedSegment("groupPolicyDefinitionId", "groupPolicyDefinitionId"),
		resourceids.StaticSegment("previousVersionDefinition", "previousVersionDefinition", "previousVersionDefinition"),
		resourceids.StaticSegment("presentations", "presentations", "presentations"),
		resourceids.UserSpecifiedSegment("groupPolicyPresentationId", "groupPolicyPresentationId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Definition Id Previous Version Definition Presentation ID
func (id DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Definition: %q", id.GroupPolicyDefinitionId),
		fmt.Sprintf("Group Policy Presentation: %q", id.GroupPolicyPresentationId),
	}
	return fmt.Sprintf("Device Management Group Policy Definition Id Previous Version Definition Presentation (%s)", strings.Join(components, "\n"))
}
