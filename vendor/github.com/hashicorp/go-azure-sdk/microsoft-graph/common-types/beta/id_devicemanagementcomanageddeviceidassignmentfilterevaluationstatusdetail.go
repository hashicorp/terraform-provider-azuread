package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}

// DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId is a struct representing the Resource ID for a Device Management Comanaged Device Id Assignment Filter Evaluation Status Detail
type DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId struct {
	ManagedDeviceId                           string
	AssignmentFilterEvaluationStatusDetailsId string
}

// NewDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailID returns a new DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId struct
func NewDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailID(managedDeviceId string, assignmentFilterEvaluationStatusDetailsId string) DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId {
	return DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
		ManagedDeviceId: managedDeviceId,
		AssignmentFilterEvaluationStatusDetailsId: assignmentFilterEvaluationStatusDetailsId,
	}
}

// ParseDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailID parses 'input' into a DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId
func ParseDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailID(input string) (*DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively parses 'input' case-insensitively into a DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively(input string) (*DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.AssignmentFilterEvaluationStatusDetailsId, ok = input.Parsed["assignmentFilterEvaluationStatusDetailsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "assignmentFilterEvaluationStatusDetailsId", input)
	}

	return nil
}

// ValidateDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailID checks that 'input' can be parsed as a Device Management Comanaged Device Id Assignment Filter Evaluation Status Detail ID
func ValidateDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Comanaged Device Id Assignment Filter Evaluation Status Detail ID
func (id DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId) ID() string {
	fmtString := "/deviceManagement/comanagedDevices/%s/assignmentFilterEvaluationStatusDetails/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.AssignmentFilterEvaluationStatusDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Comanaged Device Id Assignment Filter Evaluation Status Detail ID
func (id DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("comanagedDevices", "comanagedDevices", "comanagedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("assignmentFilterEvaluationStatusDetails", "assignmentFilterEvaluationStatusDetails", "assignmentFilterEvaluationStatusDetails"),
		resourceids.UserSpecifiedSegment("assignmentFilterEvaluationStatusDetailsId", "assignmentFilterEvaluationStatusDetailsId"),
	}
}

// String returns a human-readable description of this Device Management Comanaged Device Id Assignment Filter Evaluation Status Detail ID
func (id DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Assignment Filter Evaluation Status Details: %q", id.AssignmentFilterEvaluationStatusDetailsId),
	}
	return fmt.Sprintf("Device Management Comanaged Device Id Assignment Filter Evaluation Status Detail (%s)", strings.Join(components, "\n"))
}
