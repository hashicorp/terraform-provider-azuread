package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMemberOfId{}

// UserIdMemberOfId is a struct representing the Resource ID for a User Id Member Of
type UserIdMemberOfId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserIdMemberOfID returns a new UserIdMemberOfId struct
func NewUserIdMemberOfID(userId string, directoryObjectId string) UserIdMemberOfId {
	return UserIdMemberOfId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserIdMemberOfID parses 'input' into a UserIdMemberOfId
func ParseUserIdMemberOfID(input string) (*UserIdMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMemberOfIDInsensitively parses 'input' case-insensitively into a UserIdMemberOfId
// note: this method should only be used for API response data and not user input
func ParseUserIdMemberOfIDInsensitively(input string) (*UserIdMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMemberOfId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateUserIdMemberOfID checks that 'input' can be parsed as a User Id Member Of ID
func ValidateUserIdMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Member Of ID
func (id UserIdMemberOfId) ID() string {
	fmtString := "/users/%s/memberOf/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Member Of ID
func (id UserIdMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("memberOf", "memberOf", "memberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this User Id Member Of ID
func (id UserIdMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Id Member Of (%s)", strings.Join(components, "\n"))
}
