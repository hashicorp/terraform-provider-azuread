package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyUploadedDefinitionFileId{}

// DeviceManagementGroupPolicyUploadedDefinitionFileId is a struct representing the Resource ID for a Device Management Group Policy Uploaded Definition File
type DeviceManagementGroupPolicyUploadedDefinitionFileId struct {
	GroupPolicyUploadedDefinitionFileId string
}

// NewDeviceManagementGroupPolicyUploadedDefinitionFileID returns a new DeviceManagementGroupPolicyUploadedDefinitionFileId struct
func NewDeviceManagementGroupPolicyUploadedDefinitionFileID(groupPolicyUploadedDefinitionFileId string) DeviceManagementGroupPolicyUploadedDefinitionFileId {
	return DeviceManagementGroupPolicyUploadedDefinitionFileId{
		GroupPolicyUploadedDefinitionFileId: groupPolicyUploadedDefinitionFileId,
	}
}

// ParseDeviceManagementGroupPolicyUploadedDefinitionFileID parses 'input' into a DeviceManagementGroupPolicyUploadedDefinitionFileId
func ParseDeviceManagementGroupPolicyUploadedDefinitionFileID(input string) (*DeviceManagementGroupPolicyUploadedDefinitionFileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyUploadedDefinitionFileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyUploadedDefinitionFileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyUploadedDefinitionFileIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyUploadedDefinitionFileId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyUploadedDefinitionFileIDInsensitively(input string) (*DeviceManagementGroupPolicyUploadedDefinitionFileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyUploadedDefinitionFileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyUploadedDefinitionFileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyUploadedDefinitionFileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyUploadedDefinitionFileId, ok = input.Parsed["groupPolicyUploadedDefinitionFileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyUploadedDefinitionFileId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyUploadedDefinitionFileID checks that 'input' can be parsed as a Device Management Group Policy Uploaded Definition File ID
func ValidateDeviceManagementGroupPolicyUploadedDefinitionFileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyUploadedDefinitionFileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Uploaded Definition File ID
func (id DeviceManagementGroupPolicyUploadedDefinitionFileId) ID() string {
	fmtString := "/deviceManagement/groupPolicyUploadedDefinitionFiles/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyUploadedDefinitionFileId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Uploaded Definition File ID
func (id DeviceManagementGroupPolicyUploadedDefinitionFileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyUploadedDefinitionFiles", "groupPolicyUploadedDefinitionFiles", "groupPolicyUploadedDefinitionFiles"),
		resourceids.UserSpecifiedSegment("groupPolicyUploadedDefinitionFileId", "groupPolicyUploadedDefinitionFileId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Uploaded Definition File ID
func (id DeviceManagementGroupPolicyUploadedDefinitionFileId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Uploaded Definition File: %q", id.GroupPolicyUploadedDefinitionFileId),
	}
	return fmt.Sprintf("Device Management Group Policy Uploaded Definition File (%s)", strings.Join(components, "\n"))
}
