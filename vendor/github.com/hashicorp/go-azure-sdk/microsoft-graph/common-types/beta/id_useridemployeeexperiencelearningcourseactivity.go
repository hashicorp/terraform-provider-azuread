package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdEmployeeExperienceLearningCourseActivityId{}

// UserIdEmployeeExperienceLearningCourseActivityId is a struct representing the Resource ID for a User Id Employee Experience Learning Course Activity
type UserIdEmployeeExperienceLearningCourseActivityId struct {
	UserId                   string
	LearningCourseActivityId string
}

// NewUserIdEmployeeExperienceLearningCourseActivityID returns a new UserIdEmployeeExperienceLearningCourseActivityId struct
func NewUserIdEmployeeExperienceLearningCourseActivityID(userId string, learningCourseActivityId string) UserIdEmployeeExperienceLearningCourseActivityId {
	return UserIdEmployeeExperienceLearningCourseActivityId{
		UserId:                   userId,
		LearningCourseActivityId: learningCourseActivityId,
	}
}

// ParseUserIdEmployeeExperienceLearningCourseActivityID parses 'input' into a UserIdEmployeeExperienceLearningCourseActivityId
func ParseUserIdEmployeeExperienceLearningCourseActivityID(input string) (*UserIdEmployeeExperienceLearningCourseActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEmployeeExperienceLearningCourseActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEmployeeExperienceLearningCourseActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdEmployeeExperienceLearningCourseActivityIDInsensitively parses 'input' case-insensitively into a UserIdEmployeeExperienceLearningCourseActivityId
// note: this method should only be used for API response data and not user input
func ParseUserIdEmployeeExperienceLearningCourseActivityIDInsensitively(input string) (*UserIdEmployeeExperienceLearningCourseActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEmployeeExperienceLearningCourseActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEmployeeExperienceLearningCourseActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdEmployeeExperienceLearningCourseActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.LearningCourseActivityId, ok = input.Parsed["learningCourseActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "learningCourseActivityId", input)
	}

	return nil
}

// ValidateUserIdEmployeeExperienceLearningCourseActivityID checks that 'input' can be parsed as a User Id Employee Experience Learning Course Activity ID
func ValidateUserIdEmployeeExperienceLearningCourseActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdEmployeeExperienceLearningCourseActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Employee Experience Learning Course Activity ID
func (id UserIdEmployeeExperienceLearningCourseActivityId) ID() string {
	fmtString := "/users/%s/employeeExperience/learningCourseActivities/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.LearningCourseActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Employee Experience Learning Course Activity ID
func (id UserIdEmployeeExperienceLearningCourseActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("employeeExperience", "employeeExperience", "employeeExperience"),
		resourceids.StaticSegment("learningCourseActivities", "learningCourseActivities", "learningCourseActivities"),
		resourceids.UserSpecifiedSegment("learningCourseActivityId", "learningCourseActivityId"),
	}
}

// String returns a human-readable description of this User Id Employee Experience Learning Course Activity ID
func (id UserIdEmployeeExperienceLearningCourseActivityId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Learning Course Activity: %q", id.LearningCourseActivityId),
	}
	return fmt.Sprintf("User Id Employee Experience Learning Course Activity (%s)", strings.Join(components, "\n"))
}
