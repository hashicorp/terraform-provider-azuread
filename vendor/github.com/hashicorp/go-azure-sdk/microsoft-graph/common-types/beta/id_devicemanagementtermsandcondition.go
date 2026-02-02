package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTermsAndConditionId{}

// DeviceManagementTermsAndConditionId is a struct representing the Resource ID for a Device Management Terms And Condition
type DeviceManagementTermsAndConditionId struct {
	TermsAndConditionsId string
}

// NewDeviceManagementTermsAndConditionID returns a new DeviceManagementTermsAndConditionId struct
func NewDeviceManagementTermsAndConditionID(termsAndConditionsId string) DeviceManagementTermsAndConditionId {
	return DeviceManagementTermsAndConditionId{
		TermsAndConditionsId: termsAndConditionsId,
	}
}

// ParseDeviceManagementTermsAndConditionID parses 'input' into a DeviceManagementTermsAndConditionId
func ParseDeviceManagementTermsAndConditionID(input string) (*DeviceManagementTermsAndConditionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTermsAndConditionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTermsAndConditionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTermsAndConditionIDInsensitively parses 'input' case-insensitively into a DeviceManagementTermsAndConditionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTermsAndConditionIDInsensitively(input string) (*DeviceManagementTermsAndConditionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTermsAndConditionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTermsAndConditionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTermsAndConditionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TermsAndConditionsId, ok = input.Parsed["termsAndConditionsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "termsAndConditionsId", input)
	}

	return nil
}

// ValidateDeviceManagementTermsAndConditionID checks that 'input' can be parsed as a Device Management Terms And Condition ID
func ValidateDeviceManagementTermsAndConditionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTermsAndConditionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Terms And Condition ID
func (id DeviceManagementTermsAndConditionId) ID() string {
	fmtString := "/deviceManagement/termsAndConditions/%s"
	return fmt.Sprintf(fmtString, id.TermsAndConditionsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Terms And Condition ID
func (id DeviceManagementTermsAndConditionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("termsAndConditions", "termsAndConditions", "termsAndConditions"),
		resourceids.UserSpecifiedSegment("termsAndConditionsId", "termsAndConditionsId"),
	}
}

// String returns a human-readable description of this Device Management Terms And Condition ID
func (id DeviceManagementTermsAndConditionId) String() string {
	components := []string{
		fmt.Sprintf("Terms And Conditions: %q", id.TermsAndConditionsId),
	}
	return fmt.Sprintf("Device Management Terms And Condition (%s)", strings.Join(components, "\n"))
}
