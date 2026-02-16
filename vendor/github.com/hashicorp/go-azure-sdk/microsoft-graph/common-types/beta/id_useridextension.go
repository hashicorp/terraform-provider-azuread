package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdExtensionId{}

// UserIdExtensionId is a struct representing the Resource ID for a User Id Extension
type UserIdExtensionId struct {
	UserId      string
	ExtensionId string
}

// NewUserIdExtensionID returns a new UserIdExtensionId struct
func NewUserIdExtensionID(userId string, extensionId string) UserIdExtensionId {
	return UserIdExtensionId{
		UserId:      userId,
		ExtensionId: extensionId,
	}
}

// ParseUserIdExtensionID parses 'input' into a UserIdExtensionId
func ParseUserIdExtensionID(input string) (*UserIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdExtensionIDInsensitively(input string) (*UserIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdExtensionID checks that 'input' can be parsed as a User Id Extension ID
func ValidateUserIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Extension ID
func (id UserIdExtensionId) ID() string {
	fmtString := "/users/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Extension ID
func (id UserIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this User Id Extension ID
func (id UserIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Extension (%s)", strings.Join(components, "\n"))
}
