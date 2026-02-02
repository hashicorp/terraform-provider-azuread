package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOutlookTaskIdAttachmentId{}

// UserIdOutlookTaskIdAttachmentId is a struct representing the Resource ID for a User Id Outlook Task Id Attachment
type UserIdOutlookTaskIdAttachmentId struct {
	UserId        string
	OutlookTaskId string
	AttachmentId  string
}

// NewUserIdOutlookTaskIdAttachmentID returns a new UserIdOutlookTaskIdAttachmentId struct
func NewUserIdOutlookTaskIdAttachmentID(userId string, outlookTaskId string, attachmentId string) UserIdOutlookTaskIdAttachmentId {
	return UserIdOutlookTaskIdAttachmentId{
		UserId:        userId,
		OutlookTaskId: outlookTaskId,
		AttachmentId:  attachmentId,
	}
}

// ParseUserIdOutlookTaskIdAttachmentID parses 'input' into a UserIdOutlookTaskIdAttachmentId
func ParseUserIdOutlookTaskIdAttachmentID(input string) (*UserIdOutlookTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOutlookTaskIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdOutlookTaskIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdOutlookTaskIdAttachmentIDInsensitively(input string) (*UserIdOutlookTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOutlookTaskIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OutlookTaskId, ok = input.Parsed["outlookTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateUserIdOutlookTaskIdAttachmentID checks that 'input' can be parsed as a User Id Outlook Task Id Attachment ID
func ValidateUserIdOutlookTaskIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOutlookTaskIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Outlook Task Id Attachment ID
func (id UserIdOutlookTaskIdAttachmentId) ID() string {
	fmtString := "/users/%s/outlook/tasks/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Outlook Task Id Attachment ID
func (id UserIdOutlookTaskIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this User Id Outlook Task Id Attachment ID
func (id UserIdOutlookTaskIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Id Outlook Task Id Attachment (%s)", strings.Join(components, "\n"))
}
