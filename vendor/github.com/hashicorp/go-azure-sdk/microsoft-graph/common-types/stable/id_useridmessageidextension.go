package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMessageIdExtensionId{}

// UserIdMessageIdExtensionId is a struct representing the Resource ID for a User Id Message Id Extension
type UserIdMessageIdExtensionId struct {
	UserId      string
	MessageId   string
	ExtensionId string
}

// NewUserIdMessageIdExtensionID returns a new UserIdMessageIdExtensionId struct
func NewUserIdMessageIdExtensionID(userId string, messageId string, extensionId string) UserIdMessageIdExtensionId {
	return UserIdMessageIdExtensionId{
		UserId:      userId,
		MessageId:   messageId,
		ExtensionId: extensionId,
	}
}

// ParseUserIdMessageIdExtensionID parses 'input' into a UserIdMessageIdExtensionId
func ParseUserIdMessageIdExtensionID(input string) (*UserIdMessageIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMessageIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMessageIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMessageIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdMessageIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdMessageIdExtensionIDInsensitively(input string) (*UserIdMessageIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMessageIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMessageIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMessageIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.MessageId, ok = input.Parsed["messageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdMessageIdExtensionID checks that 'input' can be parsed as a User Id Message Id Extension ID
func ValidateUserIdMessageIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMessageIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Message Id Extension ID
func (id UserIdMessageIdExtensionId) ID() string {
	fmtString := "/users/%s/messages/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MessageId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Message Id Extension ID
func (id UserIdMessageIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this User Id Message Id Extension ID
func (id UserIdMessageIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Message Id Extension (%s)", strings.Join(components, "\n"))
}
