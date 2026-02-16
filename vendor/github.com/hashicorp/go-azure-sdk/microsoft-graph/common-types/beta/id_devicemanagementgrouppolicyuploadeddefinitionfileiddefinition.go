package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId{}

// DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId is a struct representing the Resource ID for a Device Management Group Policy Uploaded Definition File Id Definition
type DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId struct {
	GroupPolicyUploadedDefinitionFileId string
	GroupPolicyDefinitionId             string
}

// NewDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionID returns a new DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId struct
func NewDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionID(groupPolicyUploadedDefinitionFileId string, groupPolicyDefinitionId string) DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId {
	return DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId{
		GroupPolicyUploadedDefinitionFileId: groupPolicyUploadedDefinitionFileId,
		GroupPolicyDefinitionId:             groupPolicyDefinitionId,
	}
}

// ParseDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionID parses 'input' into a DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId
func ParseDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionID(input string) (*DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionIDInsensitively(input string) (*DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyUploadedDefinitionFileId, ok = input.Parsed["groupPolicyUploadedDefinitionFileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyUploadedDefinitionFileId", input)
	}

	if id.GroupPolicyDefinitionId, ok = input.Parsed["groupPolicyDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionID checks that 'input' can be parsed as a Device Management Group Policy Uploaded Definition File Id Definition ID
func ValidateDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Uploaded Definition File Id Definition ID
func (id DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId) ID() string {
	fmtString := "/deviceManagement/groupPolicyUploadedDefinitionFiles/%s/definitions/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyUploadedDefinitionFileId, id.GroupPolicyDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Uploaded Definition File Id Definition ID
func (id DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyUploadedDefinitionFiles", "groupPolicyUploadedDefinitionFiles", "groupPolicyUploadedDefinitionFiles"),
		resourceids.UserSpecifiedSegment("groupPolicyUploadedDefinitionFileId", "groupPolicyUploadedDefinitionFileId"),
		resourceids.StaticSegment("definitions", "definitions", "definitions"),
		resourceids.UserSpecifiedSegment("groupPolicyDefinitionId", "groupPolicyDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Uploaded Definition File Id Definition ID
func (id DeviceManagementGroupPolicyUploadedDefinitionFileIdDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Uploaded Definition File: %q", id.GroupPolicyUploadedDefinitionFileId),
		fmt.Sprintf("Group Policy Definition: %q", id.GroupPolicyDefinitionId),
	}
	return fmt.Sprintf("Device Management Group Policy Uploaded Definition File Id Definition (%s)", strings.Join(components, "\n"))
}
