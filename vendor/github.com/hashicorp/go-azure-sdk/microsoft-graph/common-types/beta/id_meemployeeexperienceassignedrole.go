package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEmployeeExperienceAssignedRoleId{}

// MeEmployeeExperienceAssignedRoleId is a struct representing the Resource ID for a Me Employee Experience Assigned Role
type MeEmployeeExperienceAssignedRoleId struct {
	EngagementRoleId string
}

// NewMeEmployeeExperienceAssignedRoleID returns a new MeEmployeeExperienceAssignedRoleId struct
func NewMeEmployeeExperienceAssignedRoleID(engagementRoleId string) MeEmployeeExperienceAssignedRoleId {
	return MeEmployeeExperienceAssignedRoleId{
		EngagementRoleId: engagementRoleId,
	}
}

// ParseMeEmployeeExperienceAssignedRoleID parses 'input' into a MeEmployeeExperienceAssignedRoleId
func ParseMeEmployeeExperienceAssignedRoleID(input string) (*MeEmployeeExperienceAssignedRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEmployeeExperienceAssignedRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEmployeeExperienceAssignedRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeEmployeeExperienceAssignedRoleIDInsensitively parses 'input' case-insensitively into a MeEmployeeExperienceAssignedRoleId
// note: this method should only be used for API response data and not user input
func ParseMeEmployeeExperienceAssignedRoleIDInsensitively(input string) (*MeEmployeeExperienceAssignedRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEmployeeExperienceAssignedRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEmployeeExperienceAssignedRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeEmployeeExperienceAssignedRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EngagementRoleId, ok = input.Parsed["engagementRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "engagementRoleId", input)
	}

	return nil
}

// ValidateMeEmployeeExperienceAssignedRoleID checks that 'input' can be parsed as a Me Employee Experience Assigned Role ID
func ValidateMeEmployeeExperienceAssignedRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEmployeeExperienceAssignedRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Employee Experience Assigned Role ID
func (id MeEmployeeExperienceAssignedRoleId) ID() string {
	fmtString := "/me/employeeExperience/assignedRoles/%s"
	return fmt.Sprintf(fmtString, id.EngagementRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Employee Experience Assigned Role ID
func (id MeEmployeeExperienceAssignedRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("employeeExperience", "employeeExperience", "employeeExperience"),
		resourceids.StaticSegment("assignedRoles", "assignedRoles", "assignedRoles"),
		resourceids.UserSpecifiedSegment("engagementRoleId", "engagementRoleId"),
	}
}

// String returns a human-readable description of this Me Employee Experience Assigned Role ID
func (id MeEmployeeExperienceAssignedRoleId) String() string {
	components := []string{
		fmt.Sprintf("Engagement Role: %q", id.EngagementRoleId),
	}
	return fmt.Sprintf("Me Employee Experience Assigned Role (%s)", strings.Join(components, "\n"))
}
