package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyDefinitionFileIdDefinitionId{}

// DeviceManagementGroupPolicyDefinitionFileIdDefinitionId is a struct representing the Resource ID for a Device Management Group Policy Definition File Id Definition
type DeviceManagementGroupPolicyDefinitionFileIdDefinitionId struct {
	GroupPolicyDefinitionFileId string
	GroupPolicyDefinitionId     string
}

// NewDeviceManagementGroupPolicyDefinitionFileIdDefinitionID returns a new DeviceManagementGroupPolicyDefinitionFileIdDefinitionId struct
func NewDeviceManagementGroupPolicyDefinitionFileIdDefinitionID(groupPolicyDefinitionFileId string, groupPolicyDefinitionId string) DeviceManagementGroupPolicyDefinitionFileIdDefinitionId {
	return DeviceManagementGroupPolicyDefinitionFileIdDefinitionId{
		GroupPolicyDefinitionFileId: groupPolicyDefinitionFileId,
		GroupPolicyDefinitionId:     groupPolicyDefinitionId,
	}
}

// ParseDeviceManagementGroupPolicyDefinitionFileIdDefinitionID parses 'input' into a DeviceManagementGroupPolicyDefinitionFileIdDefinitionId
func ParseDeviceManagementGroupPolicyDefinitionFileIdDefinitionID(input string) (*DeviceManagementGroupPolicyDefinitionFileIdDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionFileIdDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionFileIdDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyDefinitionFileIdDefinitionIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyDefinitionFileIdDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyDefinitionFileIdDefinitionIDInsensitively(input string) (*DeviceManagementGroupPolicyDefinitionFileIdDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionFileIdDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionFileIdDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyDefinitionFileIdDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyDefinitionFileId, ok = input.Parsed["groupPolicyDefinitionFileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyDefinitionFileId", input)
	}

	if id.GroupPolicyDefinitionId, ok = input.Parsed["groupPolicyDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyDefinitionFileIdDefinitionID checks that 'input' can be parsed as a Device Management Group Policy Definition File Id Definition ID
func ValidateDeviceManagementGroupPolicyDefinitionFileIdDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyDefinitionFileIdDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Definition File Id Definition ID
func (id DeviceManagementGroupPolicyDefinitionFileIdDefinitionId) ID() string {
	fmtString := "/deviceManagement/groupPolicyDefinitionFiles/%s/definitions/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyDefinitionFileId, id.GroupPolicyDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Definition File Id Definition ID
func (id DeviceManagementGroupPolicyDefinitionFileIdDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyDefinitionFiles", "groupPolicyDefinitionFiles", "groupPolicyDefinitionFiles"),
		resourceids.UserSpecifiedSegment("groupPolicyDefinitionFileId", "groupPolicyDefinitionFileId"),
		resourceids.StaticSegment("definitions", "definitions", "definitions"),
		resourceids.UserSpecifiedSegment("groupPolicyDefinitionId", "groupPolicyDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Definition File Id Definition ID
func (id DeviceManagementGroupPolicyDefinitionFileIdDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Definition File: %q", id.GroupPolicyDefinitionFileId),
		fmt.Sprintf("Group Policy Definition: %q", id.GroupPolicyDefinitionId),
	}
	return fmt.Sprintf("Device Management Group Policy Definition File Id Definition (%s)", strings.Join(components, "\n"))
}
