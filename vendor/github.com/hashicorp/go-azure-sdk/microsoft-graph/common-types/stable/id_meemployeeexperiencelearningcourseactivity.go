package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEmployeeExperienceLearningCourseActivityId{}

// MeEmployeeExperienceLearningCourseActivityId is a struct representing the Resource ID for a Me Employee Experience Learning Course Activity
type MeEmployeeExperienceLearningCourseActivityId struct {
	LearningCourseActivityId string
}

// NewMeEmployeeExperienceLearningCourseActivityID returns a new MeEmployeeExperienceLearningCourseActivityId struct
func NewMeEmployeeExperienceLearningCourseActivityID(learningCourseActivityId string) MeEmployeeExperienceLearningCourseActivityId {
	return MeEmployeeExperienceLearningCourseActivityId{
		LearningCourseActivityId: learningCourseActivityId,
	}
}

// ParseMeEmployeeExperienceLearningCourseActivityID parses 'input' into a MeEmployeeExperienceLearningCourseActivityId
func ParseMeEmployeeExperienceLearningCourseActivityID(input string) (*MeEmployeeExperienceLearningCourseActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEmployeeExperienceLearningCourseActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEmployeeExperienceLearningCourseActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeEmployeeExperienceLearningCourseActivityIDInsensitively parses 'input' case-insensitively into a MeEmployeeExperienceLearningCourseActivityId
// note: this method should only be used for API response data and not user input
func ParseMeEmployeeExperienceLearningCourseActivityIDInsensitively(input string) (*MeEmployeeExperienceLearningCourseActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEmployeeExperienceLearningCourseActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEmployeeExperienceLearningCourseActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeEmployeeExperienceLearningCourseActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.LearningCourseActivityId, ok = input.Parsed["learningCourseActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "learningCourseActivityId", input)
	}

	return nil
}

// ValidateMeEmployeeExperienceLearningCourseActivityID checks that 'input' can be parsed as a Me Employee Experience Learning Course Activity ID
func ValidateMeEmployeeExperienceLearningCourseActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEmployeeExperienceLearningCourseActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Employee Experience Learning Course Activity ID
func (id MeEmployeeExperienceLearningCourseActivityId) ID() string {
	fmtString := "/me/employeeExperience/learningCourseActivities/%s"
	return fmt.Sprintf(fmtString, id.LearningCourseActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Employee Experience Learning Course Activity ID
func (id MeEmployeeExperienceLearningCourseActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("employeeExperience", "employeeExperience", "employeeExperience"),
		resourceids.StaticSegment("learningCourseActivities", "learningCourseActivities", "learningCourseActivities"),
		resourceids.UserSpecifiedSegment("learningCourseActivityId", "learningCourseActivityId"),
	}
}

// String returns a human-readable description of this Me Employee Experience Learning Course Activity ID
func (id MeEmployeeExperienceLearningCourseActivityId) String() string {
	components := []string{
		fmt.Sprintf("Learning Course Activity: %q", id.LearningCourseActivityId),
	}
	return fmt.Sprintf("Me Employee Experience Learning Course Activity (%s)", strings.Join(components, "\n"))
}
