package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTermsAndConditionIdAcceptanceStatusId{}

// DeviceManagementTermsAndConditionIdAcceptanceStatusId is a struct representing the Resource ID for a Device Management Terms And Condition Id Acceptance Status
type DeviceManagementTermsAndConditionIdAcceptanceStatusId struct {
	TermsAndConditionsId                 string
	TermsAndConditionsAcceptanceStatusId string
}

// NewDeviceManagementTermsAndConditionIdAcceptanceStatusID returns a new DeviceManagementTermsAndConditionIdAcceptanceStatusId struct
func NewDeviceManagementTermsAndConditionIdAcceptanceStatusID(termsAndConditionsId string, termsAndConditionsAcceptanceStatusId string) DeviceManagementTermsAndConditionIdAcceptanceStatusId {
	return DeviceManagementTermsAndConditionIdAcceptanceStatusId{
		TermsAndConditionsId:                 termsAndConditionsId,
		TermsAndConditionsAcceptanceStatusId: termsAndConditionsAcceptanceStatusId,
	}
}

// ParseDeviceManagementTermsAndConditionIdAcceptanceStatusID parses 'input' into a DeviceManagementTermsAndConditionIdAcceptanceStatusId
func ParseDeviceManagementTermsAndConditionIdAcceptanceStatusID(input string) (*DeviceManagementTermsAndConditionIdAcceptanceStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTermsAndConditionIdAcceptanceStatusId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTermsAndConditionIdAcceptanceStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTermsAndConditionIdAcceptanceStatusIDInsensitively parses 'input' case-insensitively into a DeviceManagementTermsAndConditionIdAcceptanceStatusId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTermsAndConditionIdAcceptanceStatusIDInsensitively(input string) (*DeviceManagementTermsAndConditionIdAcceptanceStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTermsAndConditionIdAcceptanceStatusId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTermsAndConditionIdAcceptanceStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTermsAndConditionIdAcceptanceStatusId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TermsAndConditionsId, ok = input.Parsed["termsAndConditionsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "termsAndConditionsId", input)
	}

	if id.TermsAndConditionsAcceptanceStatusId, ok = input.Parsed["termsAndConditionsAcceptanceStatusId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "termsAndConditionsAcceptanceStatusId", input)
	}

	return nil
}

// ValidateDeviceManagementTermsAndConditionIdAcceptanceStatusID checks that 'input' can be parsed as a Device Management Terms And Condition Id Acceptance Status ID
func ValidateDeviceManagementTermsAndConditionIdAcceptanceStatusID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTermsAndConditionIdAcceptanceStatusID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Terms And Condition Id Acceptance Status ID
func (id DeviceManagementTermsAndConditionIdAcceptanceStatusId) ID() string {
	fmtString := "/deviceManagement/termsAndConditions/%s/acceptanceStatuses/%s"
	return fmt.Sprintf(fmtString, id.TermsAndConditionsId, id.TermsAndConditionsAcceptanceStatusId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Terms And Condition Id Acceptance Status ID
func (id DeviceManagementTermsAndConditionIdAcceptanceStatusId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("termsAndConditions", "termsAndConditions", "termsAndConditions"),
		resourceids.UserSpecifiedSegment("termsAndConditionsId", "termsAndConditionsId"),
		resourceids.StaticSegment("acceptanceStatuses", "acceptanceStatuses", "acceptanceStatuses"),
		resourceids.UserSpecifiedSegment("termsAndConditionsAcceptanceStatusId", "termsAndConditionsAcceptanceStatusId"),
	}
}

// String returns a human-readable description of this Device Management Terms And Condition Id Acceptance Status ID
func (id DeviceManagementTermsAndConditionIdAcceptanceStatusId) String() string {
	components := []string{
		fmt.Sprintf("Terms And Conditions: %q", id.TermsAndConditionsId),
		fmt.Sprintf("Terms And Conditions Acceptance Status: %q", id.TermsAndConditionsAcceptanceStatusId),
	}
	return fmt.Sprintf("Device Management Terms And Condition Id Acceptance Status (%s)", strings.Join(components, "\n"))
}
