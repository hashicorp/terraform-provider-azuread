package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAppRoleAssignmentId{}

// UserIdAppRoleAssignmentId is a struct representing the Resource ID for a User Id App Role Assignment
type UserIdAppRoleAssignmentId struct {
	UserId              string
	AppRoleAssignmentId string
}

// NewUserIdAppRoleAssignmentID returns a new UserIdAppRoleAssignmentId struct
func NewUserIdAppRoleAssignmentID(userId string, appRoleAssignmentId string) UserIdAppRoleAssignmentId {
	return UserIdAppRoleAssignmentId{
		UserId:              userId,
		AppRoleAssignmentId: appRoleAssignmentId,
	}
}

// ParseUserIdAppRoleAssignmentID parses 'input' into a UserIdAppRoleAssignmentId
func ParseUserIdAppRoleAssignmentID(input string) (*UserIdAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAppRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAppRoleAssignmentIDInsensitively parses 'input' case-insensitively into a UserIdAppRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdAppRoleAssignmentIDInsensitively(input string) (*UserIdAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAppRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAppRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.AppRoleAssignmentId, ok = input.Parsed["appRoleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appRoleAssignmentId", input)
	}

	return nil
}

// ValidateUserIdAppRoleAssignmentID checks that 'input' can be parsed as a User Id App Role Assignment ID
func ValidateUserIdAppRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAppRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id App Role Assignment ID
func (id UserIdAppRoleAssignmentId) ID() string {
	fmtString := "/users/%s/appRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AppRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id App Role Assignment ID
func (id UserIdAppRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("appRoleAssignments", "appRoleAssignments", "appRoleAssignments"),
		resourceids.UserSpecifiedSegment("appRoleAssignmentId", "appRoleAssignmentId"),
	}
}

// String returns a human-readable description of this User Id App Role Assignment ID
func (id UserIdAppRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("App Role Assignment: %q", id.AppRoleAssignmentId),
	}
	return fmt.Sprintf("User Id App Role Assignment (%s)", strings.Join(components, "\n"))
}
