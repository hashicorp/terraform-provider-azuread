package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileProjectId{}

// UserIdProfileProjectId is a struct representing the Resource ID for a User Id Profile Project
type UserIdProfileProjectId struct {
	UserId                 string
	ProjectParticipationId string
}

// NewUserIdProfileProjectID returns a new UserIdProfileProjectId struct
func NewUserIdProfileProjectID(userId string, projectParticipationId string) UserIdProfileProjectId {
	return UserIdProfileProjectId{
		UserId:                 userId,
		ProjectParticipationId: projectParticipationId,
	}
}

// ParseUserIdProfileProjectID parses 'input' into a UserIdProfileProjectId
func ParseUserIdProfileProjectID(input string) (*UserIdProfileProjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileProjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileProjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfileProjectIDInsensitively parses 'input' case-insensitively into a UserIdProfileProjectId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfileProjectIDInsensitively(input string) (*UserIdProfileProjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileProjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileProjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfileProjectId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ProjectParticipationId, ok = input.Parsed["projectParticipationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "projectParticipationId", input)
	}

	return nil
}

// ValidateUserIdProfileProjectID checks that 'input' can be parsed as a User Id Profile Project ID
func ValidateUserIdProfileProjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfileProjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Project ID
func (id UserIdProfileProjectId) ID() string {
	fmtString := "/users/%s/profile/projects/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ProjectParticipationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Project ID
func (id UserIdProfileProjectId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("projects", "projects", "projects"),
		resourceids.UserSpecifiedSegment("projectParticipationId", "projectParticipationId"),
	}
}

// String returns a human-readable description of this User Id Profile Project ID
func (id UserIdProfileProjectId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Project Participation: %q", id.ProjectParticipationId),
	}
	return fmt.Sprintf("User Id Profile Project (%s)", strings.Join(components, "\n"))
}
