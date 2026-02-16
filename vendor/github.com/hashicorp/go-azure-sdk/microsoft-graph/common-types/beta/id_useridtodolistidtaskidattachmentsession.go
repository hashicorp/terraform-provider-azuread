package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTodoListIdTaskIdAttachmentSessionId{}

// UserIdTodoListIdTaskIdAttachmentSessionId is a struct representing the Resource ID for a User Id Todo List Id Task Id Attachment Session
type UserIdTodoListIdTaskIdAttachmentSessionId struct {
	UserId              string
	TodoTaskListId      string
	TodoTaskId          string
	AttachmentSessionId string
}

// NewUserIdTodoListIdTaskIdAttachmentSessionID returns a new UserIdTodoListIdTaskIdAttachmentSessionId struct
func NewUserIdTodoListIdTaskIdAttachmentSessionID(userId string, todoTaskListId string, todoTaskId string, attachmentSessionId string) UserIdTodoListIdTaskIdAttachmentSessionId {
	return UserIdTodoListIdTaskIdAttachmentSessionId{
		UserId:              userId,
		TodoTaskListId:      todoTaskListId,
		TodoTaskId:          todoTaskId,
		AttachmentSessionId: attachmentSessionId,
	}
}

// ParseUserIdTodoListIdTaskIdAttachmentSessionID parses 'input' into a UserIdTodoListIdTaskIdAttachmentSessionId
func ParseUserIdTodoListIdTaskIdAttachmentSessionID(input string) (*UserIdTodoListIdTaskIdAttachmentSessionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListIdTaskIdAttachmentSessionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListIdTaskIdAttachmentSessionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdTodoListIdTaskIdAttachmentSessionIDInsensitively parses 'input' case-insensitively into a UserIdTodoListIdTaskIdAttachmentSessionId
// note: this method should only be used for API response data and not user input
func ParseUserIdTodoListIdTaskIdAttachmentSessionIDInsensitively(input string) (*UserIdTodoListIdTaskIdAttachmentSessionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListIdTaskIdAttachmentSessionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListIdTaskIdAttachmentSessionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdTodoListIdTaskIdAttachmentSessionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AttachmentSessionId, ok = input.Parsed["attachmentSessionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentSessionId", input)
	}

	return nil
}

// ValidateUserIdTodoListIdTaskIdAttachmentSessionID checks that 'input' can be parsed as a User Id Todo List Id Task Id Attachment Session ID
func ValidateUserIdTodoListIdTaskIdAttachmentSessionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdTodoListIdTaskIdAttachmentSessionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Todo List Id Task Id Attachment Session ID
func (id UserIdTodoListIdTaskIdAttachmentSessionId) ID() string {
	fmtString := "/users/%s/todo/lists/%s/tasks/%s/attachmentSessions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId, id.TodoTaskId, id.AttachmentSessionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Todo List Id Task Id Attachment Session ID
func (id UserIdTodoListIdTaskIdAttachmentSessionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("todo", "todo", "todo"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskId"),
		resourceids.StaticSegment("attachmentSessions", "attachmentSessions", "attachmentSessions"),
		resourceids.UserSpecifiedSegment("attachmentSessionId", "attachmentSessionId"),
	}
}

// String returns a human-readable description of this User Id Todo List Id Task Id Attachment Session ID
func (id UserIdTodoListIdTaskIdAttachmentSessionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Attachment Session: %q", id.AttachmentSessionId),
	}
	return fmt.Sprintf("User Id Todo List Id Task Id Attachment Session (%s)", strings.Join(components, "\n"))
}
