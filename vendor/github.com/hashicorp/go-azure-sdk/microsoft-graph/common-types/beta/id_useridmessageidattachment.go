package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMessageIdAttachmentId{}

// UserIdMessageIdAttachmentId is a struct representing the Resource ID for a User Id Message Id Attachment
type UserIdMessageIdAttachmentId struct {
	UserId       string
	MessageId    string
	AttachmentId string
}

// NewUserIdMessageIdAttachmentID returns a new UserIdMessageIdAttachmentId struct
func NewUserIdMessageIdAttachmentID(userId string, messageId string, attachmentId string) UserIdMessageIdAttachmentId {
	return UserIdMessageIdAttachmentId{
		UserId:       userId,
		MessageId:    messageId,
		AttachmentId: attachmentId,
	}
}

// ParseUserIdMessageIdAttachmentID parses 'input' into a UserIdMessageIdAttachmentId
func ParseUserIdMessageIdAttachmentID(input string) (*UserIdMessageIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMessageIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMessageIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMessageIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdMessageIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdMessageIdAttachmentIDInsensitively(input string) (*UserIdMessageIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMessageIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMessageIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMessageIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.MessageId, ok = input.Parsed["messageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageId", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateUserIdMessageIdAttachmentID checks that 'input' can be parsed as a User Id Message Id Attachment ID
func ValidateUserIdMessageIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMessageIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Message Id Attachment ID
func (id UserIdMessageIdAttachmentId) ID() string {
	fmtString := "/users/%s/messages/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MessageId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Message Id Attachment ID
func (id UserIdMessageIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this User Id Message Id Attachment ID
func (id UserIdMessageIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Id Message Id Attachment (%s)", strings.Join(components, "\n"))
}
