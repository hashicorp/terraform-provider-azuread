package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdScopedRoleMemberOfId{}

// UserIdScopedRoleMemberOfId is a struct representing the Resource ID for a User Id Scoped Role Member Of
type UserIdScopedRoleMemberOfId struct {
	UserId                 string
	ScopedRoleMembershipId string
}

// NewUserIdScopedRoleMemberOfID returns a new UserIdScopedRoleMemberOfId struct
func NewUserIdScopedRoleMemberOfID(userId string, scopedRoleMembershipId string) UserIdScopedRoleMemberOfId {
	return UserIdScopedRoleMemberOfId{
		UserId:                 userId,
		ScopedRoleMembershipId: scopedRoleMembershipId,
	}
}

// ParseUserIdScopedRoleMemberOfID parses 'input' into a UserIdScopedRoleMemberOfId
func ParseUserIdScopedRoleMemberOfID(input string) (*UserIdScopedRoleMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdScopedRoleMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdScopedRoleMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdScopedRoleMemberOfIDInsensitively parses 'input' case-insensitively into a UserIdScopedRoleMemberOfId
// note: this method should only be used for API response data and not user input
func ParseUserIdScopedRoleMemberOfIDInsensitively(input string) (*UserIdScopedRoleMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdScopedRoleMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdScopedRoleMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdScopedRoleMemberOfId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ScopedRoleMembershipId, ok = input.Parsed["scopedRoleMembershipId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scopedRoleMembershipId", input)
	}

	return nil
}

// ValidateUserIdScopedRoleMemberOfID checks that 'input' can be parsed as a User Id Scoped Role Member Of ID
func ValidateUserIdScopedRoleMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdScopedRoleMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Scoped Role Member Of ID
func (id UserIdScopedRoleMemberOfId) ID() string {
	fmtString := "/users/%s/scopedRoleMemberOf/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ScopedRoleMembershipId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Scoped Role Member Of ID
func (id UserIdScopedRoleMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("scopedRoleMemberOf", "scopedRoleMemberOf", "scopedRoleMemberOf"),
		resourceids.UserSpecifiedSegment("scopedRoleMembershipId", "scopedRoleMembershipId"),
	}
}

// String returns a human-readable description of this User Id Scoped Role Member Of ID
func (id UserIdScopedRoleMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Scoped Role Membership: %q", id.ScopedRoleMembershipId),
	}
	return fmt.Sprintf("User Id Scoped Role Member Of (%s)", strings.Join(components, "\n"))
}
