package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTransitiveMemberOfId{}

// UserIdTransitiveMemberOfId is a struct representing the Resource ID for a User Id Transitive Member Of
type UserIdTransitiveMemberOfId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserIdTransitiveMemberOfID returns a new UserIdTransitiveMemberOfId struct
func NewUserIdTransitiveMemberOfID(userId string, directoryObjectId string) UserIdTransitiveMemberOfId {
	return UserIdTransitiveMemberOfId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserIdTransitiveMemberOfID parses 'input' into a UserIdTransitiveMemberOfId
func ParseUserIdTransitiveMemberOfID(input string) (*UserIdTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTransitiveMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdTransitiveMemberOfIDInsensitively parses 'input' case-insensitively into a UserIdTransitiveMemberOfId
// note: this method should only be used for API response data and not user input
func ParseUserIdTransitiveMemberOfIDInsensitively(input string) (*UserIdTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTransitiveMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdTransitiveMemberOfId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateUserIdTransitiveMemberOfID checks that 'input' can be parsed as a User Id Transitive Member Of ID
func ValidateUserIdTransitiveMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdTransitiveMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Transitive Member Of ID
func (id UserIdTransitiveMemberOfId) ID() string {
	fmtString := "/users/%s/transitiveMemberOf/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Transitive Member Of ID
func (id UserIdTransitiveMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("transitiveMemberOf", "transitiveMemberOf", "transitiveMemberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this User Id Transitive Member Of ID
func (id UserIdTransitiveMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Id Transitive Member Of (%s)", strings.Join(components, "\n"))
}
