package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCreatedObjectId{}

// UserIdCreatedObjectId is a struct representing the Resource ID for a User Id Created Object
type UserIdCreatedObjectId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserIdCreatedObjectID returns a new UserIdCreatedObjectId struct
func NewUserIdCreatedObjectID(userId string, directoryObjectId string) UserIdCreatedObjectId {
	return UserIdCreatedObjectId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserIdCreatedObjectID parses 'input' into a UserIdCreatedObjectId
func ParseUserIdCreatedObjectID(input string) (*UserIdCreatedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCreatedObjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCreatedObjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCreatedObjectIDInsensitively parses 'input' case-insensitively into a UserIdCreatedObjectId
// note: this method should only be used for API response data and not user input
func ParseUserIdCreatedObjectIDInsensitively(input string) (*UserIdCreatedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCreatedObjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCreatedObjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCreatedObjectId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateUserIdCreatedObjectID checks that 'input' can be parsed as a User Id Created Object ID
func ValidateUserIdCreatedObjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCreatedObjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Created Object ID
func (id UserIdCreatedObjectId) ID() string {
	fmtString := "/users/%s/createdObjects/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Created Object ID
func (id UserIdCreatedObjectId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("createdObjects", "createdObjects", "createdObjects"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this User Id Created Object ID
func (id UserIdCreatedObjectId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Id Created Object (%s)", strings.Join(components, "\n"))
}
