package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleOfferShiftRequestId{}

// MeJoinedTeamIdScheduleOfferShiftRequestId is a struct representing the Resource ID for a Me Joined Team Id Schedule Offer Shift Request
type MeJoinedTeamIdScheduleOfferShiftRequestId struct {
	TeamId              string
	OfferShiftRequestId string
}

// NewMeJoinedTeamIdScheduleOfferShiftRequestID returns a new MeJoinedTeamIdScheduleOfferShiftRequestId struct
func NewMeJoinedTeamIdScheduleOfferShiftRequestID(teamId string, offerShiftRequestId string) MeJoinedTeamIdScheduleOfferShiftRequestId {
	return MeJoinedTeamIdScheduleOfferShiftRequestId{
		TeamId:              teamId,
		OfferShiftRequestId: offerShiftRequestId,
	}
}

// ParseMeJoinedTeamIdScheduleOfferShiftRequestID parses 'input' into a MeJoinedTeamIdScheduleOfferShiftRequestId
func ParseMeJoinedTeamIdScheduleOfferShiftRequestID(input string) (*MeJoinedTeamIdScheduleOfferShiftRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleOfferShiftRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleOfferShiftRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdScheduleOfferShiftRequestIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdScheduleOfferShiftRequestId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdScheduleOfferShiftRequestIDInsensitively(input string) (*MeJoinedTeamIdScheduleOfferShiftRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleOfferShiftRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleOfferShiftRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdScheduleOfferShiftRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.OfferShiftRequestId, ok = input.Parsed["offerShiftRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "offerShiftRequestId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdScheduleOfferShiftRequestID checks that 'input' can be parsed as a Me Joined Team Id Schedule Offer Shift Request ID
func ValidateMeJoinedTeamIdScheduleOfferShiftRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdScheduleOfferShiftRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Schedule Offer Shift Request ID
func (id MeJoinedTeamIdScheduleOfferShiftRequestId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/offerShiftRequests/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.OfferShiftRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Schedule Offer Shift Request ID
func (id MeJoinedTeamIdScheduleOfferShiftRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("offerShiftRequests", "offerShiftRequests", "offerShiftRequests"),
		resourceids.UserSpecifiedSegment("offerShiftRequestId", "offerShiftRequestId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Schedule Offer Shift Request ID
func (id MeJoinedTeamIdScheduleOfferShiftRequestId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Offer Shift Request: %q", id.OfferShiftRequestId),
	}
	return fmt.Sprintf("Me Joined Team Id Schedule Offer Shift Request (%s)", strings.Join(components, "\n"))
}
