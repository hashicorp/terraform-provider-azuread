package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyCategoryId{}

// DeviceManagementGroupPolicyCategoryId is a struct representing the Resource ID for a Device Management Group Policy Category
type DeviceManagementGroupPolicyCategoryId struct {
	GroupPolicyCategoryId string
}

// NewDeviceManagementGroupPolicyCategoryID returns a new DeviceManagementGroupPolicyCategoryId struct
func NewDeviceManagementGroupPolicyCategoryID(groupPolicyCategoryId string) DeviceManagementGroupPolicyCategoryId {
	return DeviceManagementGroupPolicyCategoryId{
		GroupPolicyCategoryId: groupPolicyCategoryId,
	}
}

// ParseDeviceManagementGroupPolicyCategoryID parses 'input' into a DeviceManagementGroupPolicyCategoryId
func ParseDeviceManagementGroupPolicyCategoryID(input string) (*DeviceManagementGroupPolicyCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyCategoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyCategoryIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyCategoryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyCategoryIDInsensitively(input string) (*DeviceManagementGroupPolicyCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyCategoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyCategoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyCategoryId, ok = input.Parsed["groupPolicyCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyCategoryId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyCategoryID checks that 'input' can be parsed as a Device Management Group Policy Category ID
func ValidateDeviceManagementGroupPolicyCategoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyCategoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Category ID
func (id DeviceManagementGroupPolicyCategoryId) ID() string {
	fmtString := "/deviceManagement/groupPolicyCategories/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyCategoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Category ID
func (id DeviceManagementGroupPolicyCategoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyCategories", "groupPolicyCategories", "groupPolicyCategories"),
		resourceids.UserSpecifiedSegment("groupPolicyCategoryId", "groupPolicyCategoryId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Category ID
func (id DeviceManagementGroupPolicyCategoryId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Category: %q", id.GroupPolicyCategoryId),
	}
	return fmt.Sprintf("Device Management Group Policy Category (%s)", strings.Join(components, "\n"))
}
