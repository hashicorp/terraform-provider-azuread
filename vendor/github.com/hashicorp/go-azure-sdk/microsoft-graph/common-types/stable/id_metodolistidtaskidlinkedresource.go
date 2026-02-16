package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeTodoListIdTaskIdLinkedResourceId{}

// MeTodoListIdTaskIdLinkedResourceId is a struct representing the Resource ID for a Me Todo List Id Task Id Linked Resource
type MeTodoListIdTaskIdLinkedResourceId struct {
	TodoTaskListId   string
	TodoTaskId       string
	LinkedResourceId string
}

// NewMeTodoListIdTaskIdLinkedResourceID returns a new MeTodoListIdTaskIdLinkedResourceId struct
func NewMeTodoListIdTaskIdLinkedResourceID(todoTaskListId string, todoTaskId string, linkedResourceId string) MeTodoListIdTaskIdLinkedResourceId {
	return MeTodoListIdTaskIdLinkedResourceId{
		TodoTaskListId:   todoTaskListId,
		TodoTaskId:       todoTaskId,
		LinkedResourceId: linkedResourceId,
	}
}

// ParseMeTodoListIdTaskIdLinkedResourceID parses 'input' into a MeTodoListIdTaskIdLinkedResourceId
func ParseMeTodoListIdTaskIdLinkedResourceID(input string) (*MeTodoListIdTaskIdLinkedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTodoListIdTaskIdLinkedResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTodoListIdTaskIdLinkedResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeTodoListIdTaskIdLinkedResourceIDInsensitively parses 'input' case-insensitively into a MeTodoListIdTaskIdLinkedResourceId
// note: this method should only be used for API response data and not user input
func ParseMeTodoListIdTaskIdLinkedResourceIDInsensitively(input string) (*MeTodoListIdTaskIdLinkedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTodoListIdTaskIdLinkedResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTodoListIdTaskIdLinkedResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeTodoListIdTaskIdLinkedResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TodoTaskListId, ok = input.Parsed["todoTaskListId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", input)
	}

	if id.TodoTaskId, ok = input.Parsed["todoTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", input)
	}

	if id.LinkedResourceId, ok = input.Parsed["linkedResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "linkedResourceId", input)
	}

	return nil
}

// ValidateMeTodoListIdTaskIdLinkedResourceID checks that 'input' can be parsed as a Me Todo List Id Task Id Linked Resource ID
func ValidateMeTodoListIdTaskIdLinkedResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTodoListIdTaskIdLinkedResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Todo List Id Task Id Linked Resource ID
func (id MeTodoListIdTaskIdLinkedResourceId) ID() string {
	fmtString := "/me/todo/lists/%s/tasks/%s/linkedResources/%s"
	return fmt.Sprintf(fmtString, id.TodoTaskListId, id.TodoTaskId, id.LinkedResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Todo List Id Task Id Linked Resource ID
func (id MeTodoListIdTaskIdLinkedResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("todo", "todo", "todo"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskId"),
		resourceids.StaticSegment("linkedResources", "linkedResources", "linkedResources"),
		resourceids.UserSpecifiedSegment("linkedResourceId", "linkedResourceId"),
	}
}

// String returns a human-readable description of this Me Todo List Id Task Id Linked Resource ID
func (id MeTodoListIdTaskIdLinkedResourceId) String() string {
	components := []string{
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Linked Resource: %q", id.LinkedResourceId),
	}
	return fmt.Sprintf("Me Todo List Id Task Id Linked Resource (%s)", strings.Join(components, "\n"))
}
