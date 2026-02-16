package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdVirtualEventWebinarId{}

// UserIdVirtualEventWebinarId is a struct representing the Resource ID for a User Id Virtual Event Webinar
type UserIdVirtualEventWebinarId struct {
	UserId                string
	VirtualEventWebinarId string
}

// NewUserIdVirtualEventWebinarID returns a new UserIdVirtualEventWebinarId struct
func NewUserIdVirtualEventWebinarID(userId string, virtualEventWebinarId string) UserIdVirtualEventWebinarId {
	return UserIdVirtualEventWebinarId{
		UserId:                userId,
		VirtualEventWebinarId: virtualEventWebinarId,
	}
}

// ParseUserIdVirtualEventWebinarID parses 'input' into a UserIdVirtualEventWebinarId
func ParseUserIdVirtualEventWebinarID(input string) (*UserIdVirtualEventWebinarId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdVirtualEventWebinarId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdVirtualEventWebinarId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdVirtualEventWebinarIDInsensitively parses 'input' case-insensitively into a UserIdVirtualEventWebinarId
// note: this method should only be used for API response data and not user input
func ParseUserIdVirtualEventWebinarIDInsensitively(input string) (*UserIdVirtualEventWebinarId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdVirtualEventWebinarId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdVirtualEventWebinarId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdVirtualEventWebinarId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.VirtualEventWebinarId, ok = input.Parsed["virtualEventWebinarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "virtualEventWebinarId", input)
	}

	return nil
}

// ValidateUserIdVirtualEventWebinarID checks that 'input' can be parsed as a User Id Virtual Event Webinar ID
func ValidateUserIdVirtualEventWebinarID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdVirtualEventWebinarID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Virtual Event Webinar ID
func (id UserIdVirtualEventWebinarId) ID() string {
	fmtString := "/users/%s/virtualEvents/webinars/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.VirtualEventWebinarId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Virtual Event Webinar ID
func (id UserIdVirtualEventWebinarId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("virtualEvents", "virtualEvents", "virtualEvents"),
		resourceids.StaticSegment("webinars", "webinars", "webinars"),
		resourceids.UserSpecifiedSegment("virtualEventWebinarId", "virtualEventWebinarId"),
	}
}

// String returns a human-readable description of this User Id Virtual Event Webinar ID
func (id UserIdVirtualEventWebinarId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Virtual Event Webinar: %q", id.VirtualEventWebinarId),
	}
	return fmt.Sprintf("User Id Virtual Event Webinar (%s)", strings.Join(components, "\n"))
}
