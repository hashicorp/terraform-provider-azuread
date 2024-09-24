package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTodoListIdTaskIdAttachmentId{}

// UserIdTodoListIdTaskIdAttachmentId is a struct representing the Resource ID for a User Id Todo List Id Task Id Attachment
type UserIdTodoListIdTaskIdAttachmentId struct {
	UserId           string
	TodoTaskListId   string
	TodoTaskId       string
	AttachmentBaseId string
}

// NewUserIdTodoListIdTaskIdAttachmentID returns a new UserIdTodoListIdTaskIdAttachmentId struct
func NewUserIdTodoListIdTaskIdAttachmentID(userId string, todoTaskListId string, todoTaskId string, attachmentBaseId string) UserIdTodoListIdTaskIdAttachmentId {
	return UserIdTodoListIdTaskIdAttachmentId{
		UserId:           userId,
		TodoTaskListId:   todoTaskListId,
		TodoTaskId:       todoTaskId,
		AttachmentBaseId: attachmentBaseId,
	}
}

// ParseUserIdTodoListIdTaskIdAttachmentID parses 'input' into a UserIdTodoListIdTaskIdAttachmentId
func ParseUserIdTodoListIdTaskIdAttachmentID(input string) (*UserIdTodoListIdTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListIdTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListIdTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdTodoListIdTaskIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdTodoListIdTaskIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdTodoListIdTaskIdAttachmentIDInsensitively(input string) (*UserIdTodoListIdTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListIdTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListIdTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdTodoListIdTaskIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TodoTaskListId, ok = input.Parsed["todoTaskListId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", input)
	}

	if id.TodoTaskId, ok = input.Parsed["todoTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", input)
	}

	if id.AttachmentBaseId, ok = input.Parsed["attachmentBaseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentBaseId", input)
	}

	return nil
}

// ValidateUserIdTodoListIdTaskIdAttachmentID checks that 'input' can be parsed as a User Id Todo List Id Task Id Attachment ID
func ValidateUserIdTodoListIdTaskIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdTodoListIdTaskIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Todo List Id Task Id Attachment ID
func (id UserIdTodoListIdTaskIdAttachmentId) ID() string {
	fmtString := "/users/%s/todo/lists/%s/tasks/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId, id.TodoTaskId, id.AttachmentBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Todo List Id Task Id Attachment ID
func (id UserIdTodoListIdTaskIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("todo", "todo", "todo"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentBaseId", "attachmentBaseId"),
	}
}

// String returns a human-readable description of this User Id Todo List Id Task Id Attachment ID
func (id UserIdTodoListIdTaskIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Attachment Base: %q", id.AttachmentBaseId),
	}
	return fmt.Sprintf("User Id Todo List Id Task Id Attachment (%s)", strings.Join(components, "\n"))
}
