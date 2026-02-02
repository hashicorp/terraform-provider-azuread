package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeTodoListIdTaskId{}

// MeTodoListIdTaskId is a struct representing the Resource ID for a Me Todo List Id Task
type MeTodoListIdTaskId struct {
	TodoTaskListId string
	TodoTaskId     string
}

// NewMeTodoListIdTaskID returns a new MeTodoListIdTaskId struct
func NewMeTodoListIdTaskID(todoTaskListId string, todoTaskId string) MeTodoListIdTaskId {
	return MeTodoListIdTaskId{
		TodoTaskListId: todoTaskListId,
		TodoTaskId:     todoTaskId,
	}
}

// ParseMeTodoListIdTaskID parses 'input' into a MeTodoListIdTaskId
func ParseMeTodoListIdTaskID(input string) (*MeTodoListIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTodoListIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTodoListIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeTodoListIdTaskIDInsensitively parses 'input' case-insensitively into a MeTodoListIdTaskId
// note: this method should only be used for API response data and not user input
func ParseMeTodoListIdTaskIDInsensitively(input string) (*MeTodoListIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTodoListIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTodoListIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeTodoListIdTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TodoTaskListId, ok = input.Parsed["todoTaskListId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", input)
	}

	if id.TodoTaskId, ok = input.Parsed["todoTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", input)
	}

	return nil
}

// ValidateMeTodoListIdTaskID checks that 'input' can be parsed as a Me Todo List Id Task ID
func ValidateMeTodoListIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTodoListIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Todo List Id Task ID
func (id MeTodoListIdTaskId) ID() string {
	fmtString := "/me/todo/lists/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.TodoTaskListId, id.TodoTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Todo List Id Task ID
func (id MeTodoListIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("todo", "todo", "todo"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskId"),
	}
}

// String returns a human-readable description of this Me Todo List Id Task ID
func (id MeTodoListIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
	}
	return fmt.Sprintf("Me Todo List Id Task (%s)", strings.Join(components, "\n"))
}
