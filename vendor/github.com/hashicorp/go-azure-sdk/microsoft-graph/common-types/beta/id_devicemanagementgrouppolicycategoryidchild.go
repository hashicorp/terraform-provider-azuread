package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyCategoryIdChildId{}

// DeviceManagementGroupPolicyCategoryIdChildId is a struct representing the Resource ID for a Device Management Group Policy Category Id Child
type DeviceManagementGroupPolicyCategoryIdChildId struct {
	GroupPolicyCategoryId  string
	GroupPolicyCategoryId1 string
}

// NewDeviceManagementGroupPolicyCategoryIdChildID returns a new DeviceManagementGroupPolicyCategoryIdChildId struct
func NewDeviceManagementGroupPolicyCategoryIdChildID(groupPolicyCategoryId string, groupPolicyCategoryId1 string) DeviceManagementGroupPolicyCategoryIdChildId {
	return DeviceManagementGroupPolicyCategoryIdChildId{
		GroupPolicyCategoryId:  groupPolicyCategoryId,
		GroupPolicyCategoryId1: groupPolicyCategoryId1,
	}
}

// ParseDeviceManagementGroupPolicyCategoryIdChildID parses 'input' into a DeviceManagementGroupPolicyCategoryIdChildId
func ParseDeviceManagementGroupPolicyCategoryIdChildID(input string) (*DeviceManagementGroupPolicyCategoryIdChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyCategoryIdChildId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyCategoryIdChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyCategoryIdChildIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyCategoryIdChildId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyCategoryIdChildIDInsensitively(input string) (*DeviceManagementGroupPolicyCategoryIdChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyCategoryIdChildId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyCategoryIdChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyCategoryIdChildId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyCategoryId, ok = input.Parsed["groupPolicyCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyCategoryId", input)
	}

	if id.GroupPolicyCategoryId1, ok = input.Parsed["groupPolicyCategoryId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyCategoryId1", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyCategoryIdChildID checks that 'input' can be parsed as a Device Management Group Policy Category Id Child ID
func ValidateDeviceManagementGroupPolicyCategoryIdChildID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyCategoryIdChildID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Category Id Child ID
func (id DeviceManagementGroupPolicyCategoryIdChildId) ID() string {
	fmtString := "/deviceManagement/groupPolicyCategories/%s/children/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyCategoryId, id.GroupPolicyCategoryId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Category Id Child ID
func (id DeviceManagementGroupPolicyCategoryIdChildId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyCategories", "groupPolicyCategories", "groupPolicyCategories"),
		resourceids.UserSpecifiedSegment("groupPolicyCategoryId", "groupPolicyCategoryId"),
		resourceids.StaticSegment("children", "children", "children"),
		resourceids.UserSpecifiedSegment("groupPolicyCategoryId1", "groupPolicyCategoryId1"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Category Id Child ID
func (id DeviceManagementGroupPolicyCategoryIdChildId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Category: %q", id.GroupPolicyCategoryId),
		fmt.Sprintf("Group Policy Category Id 1: %q", id.GroupPolicyCategoryId1),
	}
	return fmt.Sprintf("Device Management Group Policy Category Id Child (%s)", strings.Join(components, "\n"))
}
