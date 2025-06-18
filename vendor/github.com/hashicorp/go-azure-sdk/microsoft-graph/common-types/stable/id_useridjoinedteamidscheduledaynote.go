package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdScheduleDayNoteId{}

// UserIdJoinedTeamIdScheduleDayNoteId is a struct representing the Resource ID for a User Id Joined Team Id Schedule Day Note
type UserIdJoinedTeamIdScheduleDayNoteId struct {
	UserId    string
	TeamId    string
	DayNoteId string
}

// NewUserIdJoinedTeamIdScheduleDayNoteID returns a new UserIdJoinedTeamIdScheduleDayNoteId struct
func NewUserIdJoinedTeamIdScheduleDayNoteID(userId string, teamId string, dayNoteId string) UserIdJoinedTeamIdScheduleDayNoteId {
	return UserIdJoinedTeamIdScheduleDayNoteId{
		UserId:    userId,
		TeamId:    teamId,
		DayNoteId: dayNoteId,
	}
}

// ParseUserIdJoinedTeamIdScheduleDayNoteID parses 'input' into a UserIdJoinedTeamIdScheduleDayNoteId
func ParseUserIdJoinedTeamIdScheduleDayNoteID(input string) (*UserIdJoinedTeamIdScheduleDayNoteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleDayNoteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleDayNoteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdScheduleDayNoteIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdScheduleDayNoteId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdScheduleDayNoteIDInsensitively(input string) (*UserIdJoinedTeamIdScheduleDayNoteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleDayNoteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleDayNoteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdScheduleDayNoteId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.DayNoteId, ok = input.Parsed["dayNoteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dayNoteId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdScheduleDayNoteID checks that 'input' can be parsed as a User Id Joined Team Id Schedule Day Note ID
func ValidateUserIdJoinedTeamIdScheduleDayNoteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdScheduleDayNoteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Schedule Day Note ID
func (id UserIdJoinedTeamIdScheduleDayNoteId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/dayNotes/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.DayNoteId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Schedule Day Note ID
func (id UserIdJoinedTeamIdScheduleDayNoteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("dayNotes", "dayNotes", "dayNotes"),
		resourceids.UserSpecifiedSegment("dayNoteId", "dayNoteId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Schedule Day Note ID
func (id UserIdJoinedTeamIdScheduleDayNoteId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Day Note: %q", id.DayNoteId),
	}
	return fmt.Sprintf("User Id Joined Team Id Schedule Day Note (%s)", strings.Join(components, "\n"))
}
