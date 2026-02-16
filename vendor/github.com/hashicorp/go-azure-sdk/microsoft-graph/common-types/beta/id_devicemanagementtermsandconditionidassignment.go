package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTermsAndConditionIdAssignmentId{}

// DeviceManagementTermsAndConditionIdAssignmentId is a struct representing the Resource ID for a Device Management Terms And Condition Id Assignment
type DeviceManagementTermsAndConditionIdAssignmentId struct {
	TermsAndConditionsId           string
	TermsAndConditionsAssignmentId string
}

// NewDeviceManagementTermsAndConditionIdAssignmentID returns a new DeviceManagementTermsAndConditionIdAssignmentId struct
func NewDeviceManagementTermsAndConditionIdAssignmentID(termsAndConditionsId string, termsAndConditionsAssignmentId string) DeviceManagementTermsAndConditionIdAssignmentId {
	return DeviceManagementTermsAndConditionIdAssignmentId{
		TermsAndConditionsId:           termsAndConditionsId,
		TermsAndConditionsAssignmentId: termsAndConditionsAssignmentId,
	}
}

// ParseDeviceManagementTermsAndConditionIdAssignmentID parses 'input' into a DeviceManagementTermsAndConditionIdAssignmentId
func ParseDeviceManagementTermsAndConditionIdAssignmentID(input string) (*DeviceManagementTermsAndConditionIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTermsAndConditionIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTermsAndConditionIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTermsAndConditionIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementTermsAndConditionIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTermsAndConditionIdAssignmentIDInsensitively(input string) (*DeviceManagementTermsAndConditionIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTermsAndConditionIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTermsAndConditionIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTermsAndConditionIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TermsAndConditionsId, ok = input.Parsed["termsAndConditionsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "termsAndConditionsId", input)
	}

	if id.TermsAndConditionsAssignmentId, ok = input.Parsed["termsAndConditionsAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "termsAndConditionsAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementTermsAndConditionIdAssignmentID checks that 'input' can be parsed as a Device Management Terms And Condition Id Assignment ID
func ValidateDeviceManagementTermsAndConditionIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTermsAndConditionIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Terms And Condition Id Assignment ID
func (id DeviceManagementTermsAndConditionIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/termsAndConditions/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.TermsAndConditionsId, id.TermsAndConditionsAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Terms And Condition Id Assignment ID
func (id DeviceManagementTermsAndConditionIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("termsAndConditions", "termsAndConditions", "termsAndConditions"),
		resourceids.UserSpecifiedSegment("termsAndConditionsId", "termsAndConditionsId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("termsAndConditionsAssignmentId", "termsAndConditionsAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Terms And Condition Id Assignment ID
func (id DeviceManagementTermsAndConditionIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Terms And Conditions: %q", id.TermsAndConditionsId),
		fmt.Sprintf("Terms And Conditions Assignment: %q", id.TermsAndConditionsAssignmentId),
	}
	return fmt.Sprintf("Device Management Terms And Condition Id Assignment (%s)", strings.Join(components, "\n"))
}
