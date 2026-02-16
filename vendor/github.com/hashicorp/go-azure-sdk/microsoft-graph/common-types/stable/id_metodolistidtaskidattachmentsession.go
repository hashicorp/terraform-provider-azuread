package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeTodoListIdTaskIdAttachmentSessionId{}

// MeTodoListIdTaskIdAttachmentSessionId is a struct representing the Resource ID for a Me Todo List Id Task Id Attachment Session
type MeTodoListIdTaskIdAttachmentSessionId struct {
	TodoTaskListId      string
	TodoTaskId          string
	AttachmentSessionId string
}

// NewMeTodoListIdTaskIdAttachmentSessionID returns a new MeTodoListIdTaskIdAttachmentSessionId struct
func NewMeTodoListIdTaskIdAttachmentSessionID(todoTaskListId string, todoTaskId string, attachmentSessionId string) MeTodoListIdTaskIdAttachmentSessionId {
	return MeTodoListIdTaskIdAttachmentSessionId{
		TodoTaskListId:      todoTaskListId,
		TodoTaskId:          todoTaskId,
		AttachmentSessionId: attachmentSessionId,
	}
}

// ParseMeTodoListIdTaskIdAttachmentSessionID parses 'input' into a MeTodoListIdTaskIdAttachmentSessionId
func ParseMeTodoListIdTaskIdAttachmentSessionID(input string) (*MeTodoListIdTaskIdAttachmentSessionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTodoListIdTaskIdAttachmentSessionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTodoListIdTaskIdAttachmentSessionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeTodoListIdTaskIdAttachmentSessionIDInsensitively parses 'input' case-insensitively into a MeTodoListIdTaskIdAttachmentSessionId
// note: this method should only be used for API response data and not user input
func ParseMeTodoListIdTaskIdAttachmentSessionIDInsensitively(input string) (*MeTodoListIdTaskIdAttachmentSessionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTodoListIdTaskIdAttachmentSessionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTodoListIdTaskIdAttachmentSessionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeTodoListIdTaskIdAttachmentSessionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeTodoListIdTaskIdAttachmentSessionID checks that 'input' can be parsed as a Me Todo List Id Task Id Attachment Session ID
func ValidateMeTodoListIdTaskIdAttachmentSessionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTodoListIdTaskIdAttachmentSessionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Todo List Id Task Id Attachment Session ID
func (id MeTodoListIdTaskIdAttachmentSessionId) ID() string {
	fmtString := "/me/todo/lists/%s/tasks/%s/attachmentSessions/%s"
	return fmt.Sprintf(fmtString, id.TodoTaskListId, id.TodoTaskId, id.AttachmentSessionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Todo List Id Task Id Attachment Session ID
func (id MeTodoListIdTaskIdAttachmentSessionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("todo", "todo", "todo"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskId"),
		resourceids.StaticSegment("attachmentSessions", "attachmentSessions", "attachmentSessions"),
		resourceids.UserSpecifiedSegment("attachmentSessionId", "attachmentSessionId"),
	}
}

// String returns a human-readable description of this Me Todo List Id Task Id Attachment Session ID
func (id MeTodoListIdTaskIdAttachmentSessionId) String() string {
	components := []string{
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Attachment Session: %q", id.AttachmentSessionId),
	}
	return fmt.Sprintf("Me Todo List Id Task Id Attachment Session (%s)", strings.Join(components, "\n"))
}
