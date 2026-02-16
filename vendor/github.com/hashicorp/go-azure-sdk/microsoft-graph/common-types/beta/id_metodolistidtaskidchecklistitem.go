package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeTodoListIdTaskIdChecklistItemId{}

// MeTodoListIdTaskIdChecklistItemId is a struct representing the Resource ID for a Me Todo List Id Task Id Checklist Item
type MeTodoListIdTaskIdChecklistItemId struct {
	TodoTaskListId  string
	TodoTaskId      string
	ChecklistItemId string
}

// NewMeTodoListIdTaskIdChecklistItemID returns a new MeTodoListIdTaskIdChecklistItemId struct
func NewMeTodoListIdTaskIdChecklistItemID(todoTaskListId string, todoTaskId string, checklistItemId string) MeTodoListIdTaskIdChecklistItemId {
	return MeTodoListIdTaskIdChecklistItemId{
		TodoTaskListId:  todoTaskListId,
		TodoTaskId:      todoTaskId,
		ChecklistItemId: checklistItemId,
	}
}

// ParseMeTodoListIdTaskIdChecklistItemID parses 'input' into a MeTodoListIdTaskIdChecklistItemId
func ParseMeTodoListIdTaskIdChecklistItemID(input string) (*MeTodoListIdTaskIdChecklistItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTodoListIdTaskIdChecklistItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTodoListIdTaskIdChecklistItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeTodoListIdTaskIdChecklistItemIDInsensitively parses 'input' case-insensitively into a MeTodoListIdTaskIdChecklistItemId
// note: this method should only be used for API response data and not user input
func ParseMeTodoListIdTaskIdChecklistItemIDInsensitively(input string) (*MeTodoListIdTaskIdChecklistItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTodoListIdTaskIdChecklistItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTodoListIdTaskIdChecklistItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeTodoListIdTaskIdChecklistItemId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TodoTaskListId, ok = input.Parsed["todoTaskListId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", input)
	}

	if id.TodoTaskId, ok = input.Parsed["todoTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", input)
	}

	if id.ChecklistItemId, ok = input.Parsed["checklistItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "checklistItemId", input)
	}

	return nil
}

// ValidateMeTodoListIdTaskIdChecklistItemID checks that 'input' can be parsed as a Me Todo List Id Task Id Checklist Item ID
func ValidateMeTodoListIdTaskIdChecklistItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTodoListIdTaskIdChecklistItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Todo List Id Task Id Checklist Item ID
func (id MeTodoListIdTaskIdChecklistItemId) ID() string {
	fmtString := "/me/todo/lists/%s/tasks/%s/checklistItems/%s"
	return fmt.Sprintf(fmtString, id.TodoTaskListId, id.TodoTaskId, id.ChecklistItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Todo List Id Task Id Checklist Item ID
func (id MeTodoListIdTaskIdChecklistItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("todo", "todo", "todo"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskId"),
		resourceids.StaticSegment("checklistItems", "checklistItems", "checklistItems"),
		resourceids.UserSpecifiedSegment("checklistItemId", "checklistItemId"),
	}
}

// String returns a human-readable description of this Me Todo List Id Task Id Checklist Item ID
func (id MeTodoListIdTaskIdChecklistItemId) String() string {
	components := []string{
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Checklist Item: %q", id.ChecklistItemId),
	}
	return fmt.Sprintf("Me Todo List Id Task Id Checklist Item (%s)", strings.Join(components, "\n"))
}
