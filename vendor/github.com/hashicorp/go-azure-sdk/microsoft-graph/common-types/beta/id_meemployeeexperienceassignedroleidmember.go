package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEmployeeExperienceAssignedRoleIdMemberId{}

// MeEmployeeExperienceAssignedRoleIdMemberId is a struct representing the Resource ID for a Me Employee Experience Assigned Role Id Member
type MeEmployeeExperienceAssignedRoleIdMemberId struct {
	EngagementRoleId       string
	EngagementRoleMemberId string
}

// NewMeEmployeeExperienceAssignedRoleIdMemberID returns a new MeEmployeeExperienceAssignedRoleIdMemberId struct
func NewMeEmployeeExperienceAssignedRoleIdMemberID(engagementRoleId string, engagementRoleMemberId string) MeEmployeeExperienceAssignedRoleIdMemberId {
	return MeEmployeeExperienceAssignedRoleIdMemberId{
		EngagementRoleId:       engagementRoleId,
		EngagementRoleMemberId: engagementRoleMemberId,
	}
}

// ParseMeEmployeeExperienceAssignedRoleIdMemberID parses 'input' into a MeEmployeeExperienceAssignedRoleIdMemberId
func ParseMeEmployeeExperienceAssignedRoleIdMemberID(input string) (*MeEmployeeExperienceAssignedRoleIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEmployeeExperienceAssignedRoleIdMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEmployeeExperienceAssignedRoleIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeEmployeeExperienceAssignedRoleIdMemberIDInsensitively parses 'input' case-insensitively into a MeEmployeeExperienceAssignedRoleIdMemberId
// note: this method should only be used for API response data and not user input
func ParseMeEmployeeExperienceAssignedRoleIdMemberIDInsensitively(input string) (*MeEmployeeExperienceAssignedRoleIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEmployeeExperienceAssignedRoleIdMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEmployeeExperienceAssignedRoleIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeEmployeeExperienceAssignedRoleIdMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EngagementRoleId, ok = input.Parsed["engagementRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "engagementRoleId", input)
	}

	if id.EngagementRoleMemberId, ok = input.Parsed["engagementRoleMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "engagementRoleMemberId", input)
	}

	return nil
}

// ValidateMeEmployeeExperienceAssignedRoleIdMemberID checks that 'input' can be parsed as a Me Employee Experience Assigned Role Id Member ID
func ValidateMeEmployeeExperienceAssignedRoleIdMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEmployeeExperienceAssignedRoleIdMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Employee Experience Assigned Role Id Member ID
func (id MeEmployeeExperienceAssignedRoleIdMemberId) ID() string {
	fmtString := "/me/employeeExperience/assignedRoles/%s/members/%s"
	return fmt.Sprintf(fmtString, id.EngagementRoleId, id.EngagementRoleMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Employee Experience Assigned Role Id Member ID
func (id MeEmployeeExperienceAssignedRoleIdMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("employeeExperience", "employeeExperience", "employeeExperience"),
		resourceids.StaticSegment("assignedRoles", "assignedRoles", "assignedRoles"),
		resourceids.UserSpecifiedSegment("engagementRoleId", "engagementRoleId"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("engagementRoleMemberId", "engagementRoleMemberId"),
	}
}

// String returns a human-readable description of this Me Employee Experience Assigned Role Id Member ID
func (id MeEmployeeExperienceAssignedRoleIdMemberId) String() string {
	components := []string{
		fmt.Sprintf("Engagement Role: %q", id.EngagementRoleId),
		fmt.Sprintf("Engagement Role Member: %q", id.EngagementRoleMemberId),
	}
	return fmt.Sprintf("Me Employee Experience Assigned Role Id Member (%s)", strings.Join(components, "\n"))
}
