package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdScheduleOfferShiftRequestId{}

// UserIdJoinedTeamIdScheduleOfferShiftRequestId is a struct representing the Resource ID for a User Id Joined Team Id Schedule Offer Shift Request
type UserIdJoinedTeamIdScheduleOfferShiftRequestId struct {
	UserId              string
	TeamId              string
	OfferShiftRequestId string
}

// NewUserIdJoinedTeamIdScheduleOfferShiftRequestID returns a new UserIdJoinedTeamIdScheduleOfferShiftRequestId struct
func NewUserIdJoinedTeamIdScheduleOfferShiftRequestID(userId string, teamId string, offerShiftRequestId string) UserIdJoinedTeamIdScheduleOfferShiftRequestId {
	return UserIdJoinedTeamIdScheduleOfferShiftRequestId{
		UserId:              userId,
		TeamId:              teamId,
		OfferShiftRequestId: offerShiftRequestId,
	}
}

// ParseUserIdJoinedTeamIdScheduleOfferShiftRequestID parses 'input' into a UserIdJoinedTeamIdScheduleOfferShiftRequestId
func ParseUserIdJoinedTeamIdScheduleOfferShiftRequestID(input string) (*UserIdJoinedTeamIdScheduleOfferShiftRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleOfferShiftRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleOfferShiftRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdScheduleOfferShiftRequestIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdScheduleOfferShiftRequestId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdScheduleOfferShiftRequestIDInsensitively(input string) (*UserIdJoinedTeamIdScheduleOfferShiftRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleOfferShiftRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleOfferShiftRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdScheduleOfferShiftRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.OfferShiftRequestId, ok = input.Parsed["offerShiftRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "offerShiftRequestId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdScheduleOfferShiftRequestID checks that 'input' can be parsed as a User Id Joined Team Id Schedule Offer Shift Request ID
func ValidateUserIdJoinedTeamIdScheduleOfferShiftRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdScheduleOfferShiftRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Schedule Offer Shift Request ID
func (id UserIdJoinedTeamIdScheduleOfferShiftRequestId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/offerShiftRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.OfferShiftRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Schedule Offer Shift Request ID
func (id UserIdJoinedTeamIdScheduleOfferShiftRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("offerShiftRequests", "offerShiftRequests", "offerShiftRequests"),
		resourceids.UserSpecifiedSegment("offerShiftRequestId", "offerShiftRequestId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Schedule Offer Shift Request ID
func (id UserIdJoinedTeamIdScheduleOfferShiftRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Offer Shift Request: %q", id.OfferShiftRequestId),
	}
	return fmt.Sprintf("User Id Joined Team Id Schedule Offer Shift Request (%s)", strings.Join(components, "\n"))
}
