package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOwnedObjectId{}

// UserIdOwnedObjectId is a struct representing the Resource ID for a User Id Owned Object
type UserIdOwnedObjectId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserIdOwnedObjectID returns a new UserIdOwnedObjectId struct
func NewUserIdOwnedObjectID(userId string, directoryObjectId string) UserIdOwnedObjectId {
	return UserIdOwnedObjectId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserIdOwnedObjectID parses 'input' into a UserIdOwnedObjectId
func ParseUserIdOwnedObjectID(input string) (*UserIdOwnedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOwnedObjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOwnedObjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOwnedObjectIDInsensitively parses 'input' case-insensitively into a UserIdOwnedObjectId
// note: this method should only be used for API response data and not user input
func ParseUserIdOwnedObjectIDInsensitively(input string) (*UserIdOwnedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOwnedObjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOwnedObjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOwnedObjectId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateUserIdOwnedObjectID checks that 'input' can be parsed as a User Id Owned Object ID
func ValidateUserIdOwnedObjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOwnedObjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Owned Object ID
func (id UserIdOwnedObjectId) ID() string {
	fmtString := "/users/%s/ownedObjects/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Owned Object ID
func (id UserIdOwnedObjectId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("ownedObjects", "ownedObjects", "ownedObjects"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this User Id Owned Object ID
func (id UserIdOwnedObjectId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Id Owned Object (%s)", strings.Join(components, "\n"))
}
