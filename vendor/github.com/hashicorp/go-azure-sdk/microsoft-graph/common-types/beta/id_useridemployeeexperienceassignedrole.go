package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdEmployeeExperienceAssignedRoleId{}

// UserIdEmployeeExperienceAssignedRoleId is a struct representing the Resource ID for a User Id Employee Experience Assigned Role
type UserIdEmployeeExperienceAssignedRoleId struct {
	UserId           string
	EngagementRoleId string
}

// NewUserIdEmployeeExperienceAssignedRoleID returns a new UserIdEmployeeExperienceAssignedRoleId struct
func NewUserIdEmployeeExperienceAssignedRoleID(userId string, engagementRoleId string) UserIdEmployeeExperienceAssignedRoleId {
	return UserIdEmployeeExperienceAssignedRoleId{
		UserId:           userId,
		EngagementRoleId: engagementRoleId,
	}
}

// ParseUserIdEmployeeExperienceAssignedRoleID parses 'input' into a UserIdEmployeeExperienceAssignedRoleId
func ParseUserIdEmployeeExperienceAssignedRoleID(input string) (*UserIdEmployeeExperienceAssignedRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEmployeeExperienceAssignedRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEmployeeExperienceAssignedRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdEmployeeExperienceAssignedRoleIDInsensitively parses 'input' case-insensitively into a UserIdEmployeeExperienceAssignedRoleId
// note: this method should only be used for API response data and not user input
func ParseUserIdEmployeeExperienceAssignedRoleIDInsensitively(input string) (*UserIdEmployeeExperienceAssignedRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEmployeeExperienceAssignedRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEmployeeExperienceAssignedRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdEmployeeExperienceAssignedRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EngagementRoleId, ok = input.Parsed["engagementRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "engagementRoleId", input)
	}

	return nil
}

// ValidateUserIdEmployeeExperienceAssignedRoleID checks that 'input' can be parsed as a User Id Employee Experience Assigned Role ID
func ValidateUserIdEmployeeExperienceAssignedRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdEmployeeExperienceAssignedRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Employee Experience Assigned Role ID
func (id UserIdEmployeeExperienceAssignedRoleId) ID() string {
	fmtString := "/users/%s/employeeExperience/assignedRoles/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EngagementRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Employee Experience Assigned Role ID
func (id UserIdEmployeeExperienceAssignedRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("employeeExperience", "employeeExperience", "employeeExperience"),
		resourceids.StaticSegment("assignedRoles", "assignedRoles", "assignedRoles"),
		resourceids.UserSpecifiedSegment("engagementRoleId", "engagementRoleId"),
	}
}

// String returns a human-readable description of this User Id Employee Experience Assigned Role ID
func (id UserIdEmployeeExperienceAssignedRoleId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Engagement Role: %q", id.EngagementRoleId),
	}
	return fmt.Sprintf("User Id Employee Experience Assigned Role (%s)", strings.Join(components, "\n"))
}
