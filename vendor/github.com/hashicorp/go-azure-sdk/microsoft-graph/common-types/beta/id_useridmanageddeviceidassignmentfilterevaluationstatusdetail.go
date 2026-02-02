package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}

// UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId is a struct representing the Resource ID for a User Id Managed Device Id Assignment Filter Evaluation Status Detail
type UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId struct {
	UserId                                    string
	ManagedDeviceId                           string
	AssignmentFilterEvaluationStatusDetailsId string
}

// NewUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailID returns a new UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId struct
func NewUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(userId string, managedDeviceId string, assignmentFilterEvaluationStatusDetailsId string) UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId {
	return UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
		UserId:          userId,
		ManagedDeviceId: managedDeviceId,
		AssignmentFilterEvaluationStatusDetailsId: assignmentFilterEvaluationStatusDetailsId,
	}
}

// ParseUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailID parses 'input' into a UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId
func ParseUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(input string) (*UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively parses 'input' case-insensitively into a UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId
// note: this method should only be used for API response data and not user input
func ParseUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively(input string) (*UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.AssignmentFilterEvaluationStatusDetailsId, ok = input.Parsed["assignmentFilterEvaluationStatusDetailsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "assignmentFilterEvaluationStatusDetailsId", input)
	}

	return nil
}

// ValidateUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailID checks that 'input' can be parsed as a User Id Managed Device Id Assignment Filter Evaluation Status Detail ID
func ValidateUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Managed Device Id Assignment Filter Evaluation Status Detail ID
func (id UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/assignmentFilterEvaluationStatusDetails/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.AssignmentFilterEvaluationStatusDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Managed Device Id Assignment Filter Evaluation Status Detail ID
func (id UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("assignmentFilterEvaluationStatusDetails", "assignmentFilterEvaluationStatusDetails", "assignmentFilterEvaluationStatusDetails"),
		resourceids.UserSpecifiedSegment("assignmentFilterEvaluationStatusDetailsId", "assignmentFilterEvaluationStatusDetailsId"),
	}
}

// String returns a human-readable description of this User Id Managed Device Id Assignment Filter Evaluation Status Detail ID
func (id UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Assignment Filter Evaluation Status Details: %q", id.AssignmentFilterEvaluationStatusDetailsId),
	}
	return fmt.Sprintf("User Id Managed Device Id Assignment Filter Evaluation Status Detail (%s)", strings.Join(components, "\n"))
}
