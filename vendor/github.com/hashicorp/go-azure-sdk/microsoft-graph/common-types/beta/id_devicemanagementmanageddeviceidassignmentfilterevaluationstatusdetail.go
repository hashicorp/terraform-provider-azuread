package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}

// DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId is a struct representing the Resource ID for a Device Management Managed Device Id Assignment Filter Evaluation Status Detail
type DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId struct {
	ManagedDeviceId                           string
	AssignmentFilterEvaluationStatusDetailsId string
}

// NewDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailID returns a new DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId struct
func NewDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(managedDeviceId string, assignmentFilterEvaluationStatusDetailsId string) DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId {
	return DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
		ManagedDeviceId: managedDeviceId,
		AssignmentFilterEvaluationStatusDetailsId: assignmentFilterEvaluationStatusDetailsId,
	}
}

// ParseDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailID parses 'input' into a DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId
func ParseDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(input string) (*DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively parses 'input' case-insensitively into a DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively(input string) (*DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.AssignmentFilterEvaluationStatusDetailsId, ok = input.Parsed["assignmentFilterEvaluationStatusDetailsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "assignmentFilterEvaluationStatusDetailsId", input)
	}

	return nil
}

// ValidateDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailID checks that 'input' can be parsed as a Device Management Managed Device Id Assignment Filter Evaluation Status Detail ID
func ValidateDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Managed Device Id Assignment Filter Evaluation Status Detail ID
func (id DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId) ID() string {
	fmtString := "/deviceManagement/managedDevices/%s/assignmentFilterEvaluationStatusDetails/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.AssignmentFilterEvaluationStatusDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Managed Device Id Assignment Filter Evaluation Status Detail ID
func (id DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("assignmentFilterEvaluationStatusDetails", "assignmentFilterEvaluationStatusDetails", "assignmentFilterEvaluationStatusDetails"),
		resourceids.UserSpecifiedSegment("assignmentFilterEvaluationStatusDetailsId", "assignmentFilterEvaluationStatusDetailsId"),
	}
}

// String returns a human-readable description of this Device Management Managed Device Id Assignment Filter Evaluation Status Detail ID
func (id DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Assignment Filter Evaluation Status Details: %q", id.AssignmentFilterEvaluationStatusDetailsId),
	}
	return fmt.Sprintf("Device Management Managed Device Id Assignment Filter Evaluation Status Detail (%s)", strings.Join(components, "\n"))
}
