package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyDefinitionFileId{}

// DeviceManagementGroupPolicyDefinitionFileId is a struct representing the Resource ID for a Device Management Group Policy Definition File
type DeviceManagementGroupPolicyDefinitionFileId struct {
	GroupPolicyDefinitionFileId string
}

// NewDeviceManagementGroupPolicyDefinitionFileID returns a new DeviceManagementGroupPolicyDefinitionFileId struct
func NewDeviceManagementGroupPolicyDefinitionFileID(groupPolicyDefinitionFileId string) DeviceManagementGroupPolicyDefinitionFileId {
	return DeviceManagementGroupPolicyDefinitionFileId{
		GroupPolicyDefinitionFileId: groupPolicyDefinitionFileId,
	}
}

// ParseDeviceManagementGroupPolicyDefinitionFileID parses 'input' into a DeviceManagementGroupPolicyDefinitionFileId
func ParseDeviceManagementGroupPolicyDefinitionFileID(input string) (*DeviceManagementGroupPolicyDefinitionFileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionFileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionFileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyDefinitionFileIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyDefinitionFileId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyDefinitionFileIDInsensitively(input string) (*DeviceManagementGroupPolicyDefinitionFileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyDefinitionFileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyDefinitionFileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyDefinitionFileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyDefinitionFileId, ok = input.Parsed["groupPolicyDefinitionFileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyDefinitionFileId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyDefinitionFileID checks that 'input' can be parsed as a Device Management Group Policy Definition File ID
func ValidateDeviceManagementGroupPolicyDefinitionFileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyDefinitionFileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Definition File ID
func (id DeviceManagementGroupPolicyDefinitionFileId) ID() string {
	fmtString := "/deviceManagement/groupPolicyDefinitionFiles/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyDefinitionFileId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Definition File ID
func (id DeviceManagementGroupPolicyDefinitionFileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyDefinitionFiles", "groupPolicyDefinitionFiles", "groupPolicyDefinitionFiles"),
		resourceids.UserSpecifiedSegment("groupPolicyDefinitionFileId", "groupPolicyDefinitionFileId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Definition File ID
func (id DeviceManagementGroupPolicyDefinitionFileId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Definition File: %q", id.GroupPolicyDefinitionFileId),
	}
	return fmt.Sprintf("Device Management Group Policy Definition File (%s)", strings.Join(components, "\n"))
}
