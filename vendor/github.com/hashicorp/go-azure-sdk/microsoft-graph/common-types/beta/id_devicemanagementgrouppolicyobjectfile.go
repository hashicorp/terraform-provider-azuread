package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyObjectFileId{}

// DeviceManagementGroupPolicyObjectFileId is a struct representing the Resource ID for a Device Management Group Policy Object File
type DeviceManagementGroupPolicyObjectFileId struct {
	GroupPolicyObjectFileId string
}

// NewDeviceManagementGroupPolicyObjectFileID returns a new DeviceManagementGroupPolicyObjectFileId struct
func NewDeviceManagementGroupPolicyObjectFileID(groupPolicyObjectFileId string) DeviceManagementGroupPolicyObjectFileId {
	return DeviceManagementGroupPolicyObjectFileId{
		GroupPolicyObjectFileId: groupPolicyObjectFileId,
	}
}

// ParseDeviceManagementGroupPolicyObjectFileID parses 'input' into a DeviceManagementGroupPolicyObjectFileId
func ParseDeviceManagementGroupPolicyObjectFileID(input string) (*DeviceManagementGroupPolicyObjectFileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyObjectFileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyObjectFileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyObjectFileIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyObjectFileId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyObjectFileIDInsensitively(input string) (*DeviceManagementGroupPolicyObjectFileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyObjectFileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyObjectFileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyObjectFileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyObjectFileId, ok = input.Parsed["groupPolicyObjectFileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyObjectFileId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyObjectFileID checks that 'input' can be parsed as a Device Management Group Policy Object File ID
func ValidateDeviceManagementGroupPolicyObjectFileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyObjectFileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Object File ID
func (id DeviceManagementGroupPolicyObjectFileId) ID() string {
	fmtString := "/deviceManagement/groupPolicyObjectFiles/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyObjectFileId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Object File ID
func (id DeviceManagementGroupPolicyObjectFileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyObjectFiles", "groupPolicyObjectFiles", "groupPolicyObjectFiles"),
		resourceids.UserSpecifiedSegment("groupPolicyObjectFileId", "groupPolicyObjectFileId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Object File ID
func (id DeviceManagementGroupPolicyObjectFileId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Object File: %q", id.GroupPolicyObjectFileId),
	}
	return fmt.Sprintf("Device Management Group Policy Object File (%s)", strings.Join(components, "\n"))
}
