package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}

// MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId is a struct representing the Resource ID for a Me Managed Device Id Assignment Filter Evaluation Status Detail
type MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId struct {
	ManagedDeviceId                           string
	AssignmentFilterEvaluationStatusDetailsId string
}

// NewMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailID returns a new MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId struct
func NewMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(managedDeviceId string, assignmentFilterEvaluationStatusDetailsId string) MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId {
	return MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
		ManagedDeviceId: managedDeviceId,
		AssignmentFilterEvaluationStatusDetailsId: assignmentFilterEvaluationStatusDetailsId,
	}
}

// ParseMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailID parses 'input' into a MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId
func ParseMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(input string) (*MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively(input string) (*MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.AssignmentFilterEvaluationStatusDetailsId, ok = input.Parsed["assignmentFilterEvaluationStatusDetailsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "assignmentFilterEvaluationStatusDetailsId", input)
	}

	return nil
}

// ValidateMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailID checks that 'input' can be parsed as a Me Managed Device Id Assignment Filter Evaluation Status Detail ID
func ValidateMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Id Assignment Filter Evaluation Status Detail ID
func (id MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId) ID() string {
	fmtString := "/me/managedDevices/%s/assignmentFilterEvaluationStatusDetails/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.AssignmentFilterEvaluationStatusDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Id Assignment Filter Evaluation Status Detail ID
func (id MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("assignmentFilterEvaluationStatusDetails", "assignmentFilterEvaluationStatusDetails", "assignmentFilterEvaluationStatusDetails"),
		resourceids.UserSpecifiedSegment("assignmentFilterEvaluationStatusDetailsId", "assignmentFilterEvaluationStatusDetailsId"),
	}
}

// String returns a human-readable description of this Me Managed Device Id Assignment Filter Evaluation Status Detail ID
func (id MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Assignment Filter Evaluation Status Details: %q", id.AssignmentFilterEvaluationStatusDetailsId),
	}
	return fmt.Sprintf("Me Managed Device Id Assignment Filter Evaluation Status Detail (%s)", strings.Join(components, "\n"))
}
