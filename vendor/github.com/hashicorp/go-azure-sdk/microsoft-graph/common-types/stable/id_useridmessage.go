package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMessageId{}

// UserIdMessageId is a struct representing the Resource ID for a User Id Message
type UserIdMessageId struct {
	UserId    string
	MessageId string
}

// NewUserIdMessageID returns a new UserIdMessageId struct
func NewUserIdMessageID(userId string, messageId string) UserIdMessageId {
	return UserIdMessageId{
		UserId:    userId,
		MessageId: messageId,
	}
}

// ParseUserIdMessageID parses 'input' into a UserIdMessageId
func ParseUserIdMessageID(input string) (*UserIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMessageIDInsensitively parses 'input' case-insensitively into a UserIdMessageId
// note: this method should only be used for API response data and not user input
func ParseUserIdMessageIDInsensitively(input string) (*UserIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMessageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.MessageId, ok = input.Parsed["messageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageId", input)
	}

	return nil
}

// ValidateUserIdMessageID checks that 'input' can be parsed as a User Id Message ID
func ValidateUserIdMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Message ID
func (id UserIdMessageId) ID() string {
	fmtString := "/users/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Message ID
func (id UserIdMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
	}
}

// String returns a human-readable description of this User Id Message ID
func (id UserIdMessageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Message: %q", id.MessageId),
	}
	return fmt.Sprintf("User Id Message (%s)", strings.Join(components, "\n"))
}
