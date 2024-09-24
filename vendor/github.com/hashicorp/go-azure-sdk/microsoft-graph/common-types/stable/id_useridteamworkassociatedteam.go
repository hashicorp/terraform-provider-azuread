package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTeamworkAssociatedTeamId{}

// UserIdTeamworkAssociatedTeamId is a struct representing the Resource ID for a User Id Teamwork Associated Team
type UserIdTeamworkAssociatedTeamId struct {
	UserId               string
	AssociatedTeamInfoId string
}

// NewUserIdTeamworkAssociatedTeamID returns a new UserIdTeamworkAssociatedTeamId struct
func NewUserIdTeamworkAssociatedTeamID(userId string, associatedTeamInfoId string) UserIdTeamworkAssociatedTeamId {
	return UserIdTeamworkAssociatedTeamId{
		UserId:               userId,
		AssociatedTeamInfoId: associatedTeamInfoId,
	}
}

// ParseUserIdTeamworkAssociatedTeamID parses 'input' into a UserIdTeamworkAssociatedTeamId
func ParseUserIdTeamworkAssociatedTeamID(input string) (*UserIdTeamworkAssociatedTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTeamworkAssociatedTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTeamworkAssociatedTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdTeamworkAssociatedTeamIDInsensitively parses 'input' case-insensitively into a UserIdTeamworkAssociatedTeamId
// note: this method should only be used for API response data and not user input
func ParseUserIdTeamworkAssociatedTeamIDInsensitively(input string) (*UserIdTeamworkAssociatedTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTeamworkAssociatedTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTeamworkAssociatedTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdTeamworkAssociatedTeamId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.AssociatedTeamInfoId, ok = input.Parsed["associatedTeamInfoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "associatedTeamInfoId", input)
	}

	return nil
}

// ValidateUserIdTeamworkAssociatedTeamID checks that 'input' can be parsed as a User Id Teamwork Associated Team ID
func ValidateUserIdTeamworkAssociatedTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdTeamworkAssociatedTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Teamwork Associated Team ID
func (id UserIdTeamworkAssociatedTeamId) ID() string {
	fmtString := "/users/%s/teamwork/associatedTeams/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AssociatedTeamInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Teamwork Associated Team ID
func (id UserIdTeamworkAssociatedTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("teamwork", "teamwork", "teamwork"),
		resourceids.StaticSegment("associatedTeams", "associatedTeams", "associatedTeams"),
		resourceids.UserSpecifiedSegment("associatedTeamInfoId", "associatedTeamInfoId"),
	}
}

// String returns a human-readable description of this User Id Teamwork Associated Team ID
func (id UserIdTeamworkAssociatedTeamId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Associated Team Info: %q", id.AssociatedTeamInfoId),
	}
	return fmt.Sprintf("User Id Teamwork Associated Team (%s)", strings.Join(components, "\n"))
}
