package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId{}

// DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId is a struct representing the Resource ID for a Device Management Group Policy Uploaded Definition File Id Group Policy Operation
type DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId struct {
	GroupPolicyUploadedDefinitionFileId string
	GroupPolicyOperationId              string
}

// NewDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationID returns a new DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId struct
func NewDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationID(groupPolicyUploadedDefinitionFileId string, groupPolicyOperationId string) DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId {
	return DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId{
		GroupPolicyUploadedDefinitionFileId: groupPolicyUploadedDefinitionFileId,
		GroupPolicyOperationId:              groupPolicyOperationId,
	}
}

// ParseDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationID parses 'input' into a DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId
func ParseDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationID(input string) (*DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationIDInsensitively(input string) (*DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyUploadedDefinitionFileId, ok = input.Parsed["groupPolicyUploadedDefinitionFileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyUploadedDefinitionFileId", input)
	}

	if id.GroupPolicyOperationId, ok = input.Parsed["groupPolicyOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyOperationId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationID checks that 'input' can be parsed as a Device Management Group Policy Uploaded Definition File Id Group Policy Operation ID
func ValidateDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Uploaded Definition File Id Group Policy Operation ID
func (id DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId) ID() string {
	fmtString := "/deviceManagement/groupPolicyUploadedDefinitionFiles/%s/groupPolicyOperations/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyUploadedDefinitionFileId, id.GroupPolicyOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Uploaded Definition File Id Group Policy Operation ID
func (id DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyUploadedDefinitionFiles", "groupPolicyUploadedDefinitionFiles", "groupPolicyUploadedDefinitionFiles"),
		resourceids.UserSpecifiedSegment("groupPolicyUploadedDefinitionFileId", "groupPolicyUploadedDefinitionFileId"),
		resourceids.StaticSegment("groupPolicyOperations", "groupPolicyOperations", "groupPolicyOperations"),
		resourceids.UserSpecifiedSegment("groupPolicyOperationId", "groupPolicyOperationId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Uploaded Definition File Id Group Policy Operation ID
func (id DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Uploaded Definition File: %q", id.GroupPolicyUploadedDefinitionFileId),
		fmt.Sprintf("Group Policy Operation: %q", id.GroupPolicyOperationId),
	}
	return fmt.Sprintf("Device Management Group Policy Uploaded Definition File Id Group Policy Operation (%s)", strings.Join(components, "\n"))
}
