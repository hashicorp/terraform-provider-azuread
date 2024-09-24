package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTermsAndConditionIdGroupAssignmentId{}

// DeviceManagementTermsAndConditionIdGroupAssignmentId is a struct representing the Resource ID for a Device Management Terms And Condition Id Group Assignment
type DeviceManagementTermsAndConditionIdGroupAssignmentId struct {
	TermsAndConditionsId                string
	TermsAndConditionsGroupAssignmentId string
}

// NewDeviceManagementTermsAndConditionIdGroupAssignmentID returns a new DeviceManagementTermsAndConditionIdGroupAssignmentId struct
func NewDeviceManagementTermsAndConditionIdGroupAssignmentID(termsAndConditionsId string, termsAndConditionsGroupAssignmentId string) DeviceManagementTermsAndConditionIdGroupAssignmentId {
	return DeviceManagementTermsAndConditionIdGroupAssignmentId{
		TermsAndConditionsId:                termsAndConditionsId,
		TermsAndConditionsGroupAssignmentId: termsAndConditionsGroupAssignmentId,
	}
}

// ParseDeviceManagementTermsAndConditionIdGroupAssignmentID parses 'input' into a DeviceManagementTermsAndConditionIdGroupAssignmentId
func ParseDeviceManagementTermsAndConditionIdGroupAssignmentID(input string) (*DeviceManagementTermsAndConditionIdGroupAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTermsAndConditionIdGroupAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTermsAndConditionIdGroupAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTermsAndConditionIdGroupAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementTermsAndConditionIdGroupAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTermsAndConditionIdGroupAssignmentIDInsensitively(input string) (*DeviceManagementTermsAndConditionIdGroupAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTermsAndConditionIdGroupAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTermsAndConditionIdGroupAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTermsAndConditionIdGroupAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TermsAndConditionsId, ok = input.Parsed["termsAndConditionsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "termsAndConditionsId", input)
	}

	if id.TermsAndConditionsGroupAssignmentId, ok = input.Parsed["termsAndConditionsGroupAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "termsAndConditionsGroupAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementTermsAndConditionIdGroupAssignmentID checks that 'input' can be parsed as a Device Management Terms And Condition Id Group Assignment ID
func ValidateDeviceManagementTermsAndConditionIdGroupAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTermsAndConditionIdGroupAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Terms And Condition Id Group Assignment ID
func (id DeviceManagementTermsAndConditionIdGroupAssignmentId) ID() string {
	fmtString := "/deviceManagement/termsAndConditions/%s/groupAssignments/%s"
	return fmt.Sprintf(fmtString, id.TermsAndConditionsId, id.TermsAndConditionsGroupAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Terms And Condition Id Group Assignment ID
func (id DeviceManagementTermsAndConditionIdGroupAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("termsAndConditions", "termsAndConditions", "termsAndConditions"),
		resourceids.UserSpecifiedSegment("termsAndConditionsId", "termsAndConditionsId"),
		resourceids.StaticSegment("groupAssignments", "groupAssignments", "groupAssignments"),
		resourceids.UserSpecifiedSegment("termsAndConditionsGroupAssignmentId", "termsAndConditionsGroupAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Terms And Condition Id Group Assignment ID
func (id DeviceManagementTermsAndConditionIdGroupAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Terms And Conditions: %q", id.TermsAndConditionsId),
		fmt.Sprintf("Terms And Conditions Group Assignment: %q", id.TermsAndConditionsGroupAssignmentId),
	}
	return fmt.Sprintf("Device Management Terms And Condition Id Group Assignment (%s)", strings.Join(components, "\n"))
}
