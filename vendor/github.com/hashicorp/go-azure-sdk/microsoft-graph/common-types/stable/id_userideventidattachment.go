package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdEventIdAttachmentId{}

// UserIdEventIdAttachmentId is a struct representing the Resource ID for a User Id Event Id Attachment
type UserIdEventIdAttachmentId struct {
	UserId       string
	EventId      string
	AttachmentId string
}

// NewUserIdEventIdAttachmentID returns a new UserIdEventIdAttachmentId struct
func NewUserIdEventIdAttachmentID(userId string, eventId string, attachmentId string) UserIdEventIdAttachmentId {
	return UserIdEventIdAttachmentId{
		UserId:       userId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseUserIdEventIdAttachmentID parses 'input' into a UserIdEventIdAttachmentId
func ParseUserIdEventIdAttachmentID(input string) (*UserIdEventIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEventIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEventIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdEventIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdEventIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdEventIdAttachmentIDInsensitively(input string) (*UserIdEventIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEventIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEventIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdEventIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateUserIdEventIdAttachmentID checks that 'input' can be parsed as a User Id Event Id Attachment ID
func ValidateUserIdEventIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdEventIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Event Id Attachment ID
func (id UserIdEventIdAttachmentId) ID() string {
	fmtString := "/users/%s/events/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Event Id Attachment ID
func (id UserIdEventIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this User Id Event Id Attachment ID
func (id UserIdEventIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Id Event Id Attachment (%s)", strings.Join(components, "\n"))
}
