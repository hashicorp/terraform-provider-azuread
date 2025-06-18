package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdEmployeeExperienceAssignedRoleIdMemberId{}

// UserIdEmployeeExperienceAssignedRoleIdMemberId is a struct representing the Resource ID for a User Id Employee Experience Assigned Role Id Member
type UserIdEmployeeExperienceAssignedRoleIdMemberId struct {
	UserId                 string
	EngagementRoleId       string
	EngagementRoleMemberId string
}

// NewUserIdEmployeeExperienceAssignedRoleIdMemberID returns a new UserIdEmployeeExperienceAssignedRoleIdMemberId struct
func NewUserIdEmployeeExperienceAssignedRoleIdMemberID(userId string, engagementRoleId string, engagementRoleMemberId string) UserIdEmployeeExperienceAssignedRoleIdMemberId {
	return UserIdEmployeeExperienceAssignedRoleIdMemberId{
		UserId:                 userId,
		EngagementRoleId:       engagementRoleId,
		EngagementRoleMemberId: engagementRoleMemberId,
	}
}

// ParseUserIdEmployeeExperienceAssignedRoleIdMemberID parses 'input' into a UserIdEmployeeExperienceAssignedRoleIdMemberId
func ParseUserIdEmployeeExperienceAssignedRoleIdMemberID(input string) (*UserIdEmployeeExperienceAssignedRoleIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEmployeeExperienceAssignedRoleIdMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEmployeeExperienceAssignedRoleIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdEmployeeExperienceAssignedRoleIdMemberIDInsensitively parses 'input' case-insensitively into a UserIdEmployeeExperienceAssignedRoleIdMemberId
// note: this method should only be used for API response data and not user input
func ParseUserIdEmployeeExperienceAssignedRoleIdMemberIDInsensitively(input string) (*UserIdEmployeeExperienceAssignedRoleIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEmployeeExperienceAssignedRoleIdMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEmployeeExperienceAssignedRoleIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdEmployeeExperienceAssignedRoleIdMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EngagementRoleId, ok = input.Parsed["engagementRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "engagementRoleId", input)
	}

	if id.EngagementRoleMemberId, ok = input.Parsed["engagementRoleMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "engagementRoleMemberId", input)
	}

	return nil
}

// ValidateUserIdEmployeeExperienceAssignedRoleIdMemberID checks that 'input' can be parsed as a User Id Employee Experience Assigned Role Id Member ID
func ValidateUserIdEmployeeExperienceAssignedRoleIdMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdEmployeeExperienceAssignedRoleIdMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Employee Experience Assigned Role Id Member ID
func (id UserIdEmployeeExperienceAssignedRoleIdMemberId) ID() string {
	fmtString := "/users/%s/employeeExperience/assignedRoles/%s/members/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EngagementRoleId, id.EngagementRoleMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Employee Experience Assigned Role Id Member ID
func (id UserIdEmployeeExperienceAssignedRoleIdMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("employeeExperience", "employeeExperience", "employeeExperience"),
		resourceids.StaticSegment("assignedRoles", "assignedRoles", "assignedRoles"),
		resourceids.UserSpecifiedSegment("engagementRoleId", "engagementRoleId"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("engagementRoleMemberId", "engagementRoleMemberId"),
	}
}

// String returns a human-readable description of this User Id Employee Experience Assigned Role Id Member ID
func (id UserIdEmployeeExperienceAssignedRoleIdMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Engagement Role: %q", id.EngagementRoleId),
		fmt.Sprintf("Engagement Role Member: %q", id.EngagementRoleMemberId),
	}
	return fmt.Sprintf("User Id Employee Experience Assigned Role Id Member (%s)", strings.Join(components, "\n"))
}
