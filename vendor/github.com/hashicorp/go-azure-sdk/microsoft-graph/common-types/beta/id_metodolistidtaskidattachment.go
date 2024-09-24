package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeTodoListIdTaskIdAttachmentId{}

// MeTodoListIdTaskIdAttachmentId is a struct representing the Resource ID for a Me Todo List Id Task Id Attachment
type MeTodoListIdTaskIdAttachmentId struct {
	TodoTaskListId   string
	TodoTaskId       string
	AttachmentBaseId string
}

// NewMeTodoListIdTaskIdAttachmentID returns a new MeTodoListIdTaskIdAttachmentId struct
func NewMeTodoListIdTaskIdAttachmentID(todoTaskListId string, todoTaskId string, attachmentBaseId string) MeTodoListIdTaskIdAttachmentId {
	return MeTodoListIdTaskIdAttachmentId{
		TodoTaskListId:   todoTaskListId,
		TodoTaskId:       todoTaskId,
		AttachmentBaseId: attachmentBaseId,
	}
}

// ParseMeTodoListIdTaskIdAttachmentID parses 'input' into a MeTodoListIdTaskIdAttachmentId
func ParseMeTodoListIdTaskIdAttachmentID(input string) (*MeTodoListIdTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTodoListIdTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTodoListIdTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeTodoListIdTaskIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeTodoListIdTaskIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeTodoListIdTaskIdAttachmentIDInsensitively(input string) (*MeTodoListIdTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTodoListIdTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTodoListIdTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeTodoListIdTaskIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeTodoListIdTaskIdAttachmentID checks that 'input' can be parsed as a Me Todo List Id Task Id Attachment ID
func ValidateMeTodoListIdTaskIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTodoListIdTaskIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Todo List Id Task Id Attachment ID
func (id MeTodoListIdTaskIdAttachmentId) ID() string {
	fmtString := "/me/todo/lists/%s/tasks/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.TodoTaskListId, id.TodoTaskId, id.AttachmentBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Todo List Id Task Id Attachment ID
func (id MeTodoListIdTaskIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("todo", "todo", "todo"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentBaseId", "attachmentBaseId"),
	}
}

// String returns a human-readable description of this Me Todo List Id Task Id Attachment ID
func (id MeTodoListIdTaskIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Attachment Base: %q", id.AttachmentBaseId),
	}
	return fmt.Sprintf("Me Todo List Id Task Id Attachment (%s)", strings.Join(components, "\n"))
}
