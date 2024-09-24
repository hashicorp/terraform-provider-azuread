package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdPermissionGrantId{}

// UserIdJoinedTeamIdPermissionGrantId is a struct representing the Resource ID for a User Id Joined Team Id Permission Grant
type UserIdJoinedTeamIdPermissionGrantId struct {
	UserId                            string
	TeamId                            string
	ResourceSpecificPermissionGrantId string
}

// NewUserIdJoinedTeamIdPermissionGrantID returns a new UserIdJoinedTeamIdPermissionGrantId struct
func NewUserIdJoinedTeamIdPermissionGrantID(userId string, teamId string, resourceSpecificPermissionGrantId string) UserIdJoinedTeamIdPermissionGrantId {
	return UserIdJoinedTeamIdPermissionGrantId{
		UserId:                            userId,
		TeamId:                            teamId,
		ResourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// ParseUserIdJoinedTeamIdPermissionGrantID parses 'input' into a UserIdJoinedTeamIdPermissionGrantId
func ParseUserIdJoinedTeamIdPermissionGrantID(input string) (*UserIdJoinedTeamIdPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdPermissionGrantIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdPermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdPermissionGrantIDInsensitively(input string) (*UserIdJoinedTeamIdPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdPermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ResourceSpecificPermissionGrantId, ok = input.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdPermissionGrantID checks that 'input' can be parsed as a User Id Joined Team Id Permission Grant ID
func ValidateUserIdJoinedTeamIdPermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdPermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Permission Grant ID
func (id UserIdJoinedTeamIdPermissionGrantId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/permissionGrants/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ResourceSpecificPermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Permission Grant ID
func (id UserIdJoinedTeamIdPermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("permissionGrants", "permissionGrants", "permissionGrants"),
		resourceids.UserSpecifiedSegment("resourceSpecificPermissionGrantId", "resourceSpecificPermissionGrantId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Permission Grant ID
func (id UserIdJoinedTeamIdPermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Resource Specific Permission Grant: %q", id.ResourceSpecificPermissionGrantId),
	}
	return fmt.Sprintf("User Id Joined Team Id Permission Grant (%s)", strings.Join(components, "\n"))
}
