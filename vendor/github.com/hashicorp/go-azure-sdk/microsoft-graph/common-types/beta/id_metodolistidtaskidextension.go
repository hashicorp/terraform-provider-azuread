package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeTodoListIdTaskIdExtensionId{}

// MeTodoListIdTaskIdExtensionId is a struct representing the Resource ID for a Me Todo List Id Task Id Extension
type MeTodoListIdTaskIdExtensionId struct {
	TodoTaskListId string
	TodoTaskId     string
	ExtensionId    string
}

// NewMeTodoListIdTaskIdExtensionID returns a new MeTodoListIdTaskIdExtensionId struct
func NewMeTodoListIdTaskIdExtensionID(todoTaskListId string, todoTaskId string, extensionId string) MeTodoListIdTaskIdExtensionId {
	return MeTodoListIdTaskIdExtensionId{
		TodoTaskListId: todoTaskListId,
		TodoTaskId:     todoTaskId,
		ExtensionId:    extensionId,
	}
}

// ParseMeTodoListIdTaskIdExtensionID parses 'input' into a MeTodoListIdTaskIdExtensionId
func ParseMeTodoListIdTaskIdExtensionID(input string) (*MeTodoListIdTaskIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTodoListIdTaskIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTodoListIdTaskIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeTodoListIdTaskIdExtensionIDInsensitively parses 'input' case-insensitively into a MeTodoListIdTaskIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeTodoListIdTaskIdExtensionIDInsensitively(input string) (*MeTodoListIdTaskIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTodoListIdTaskIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTodoListIdTaskIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeTodoListIdTaskIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TodoTaskListId, ok = input.Parsed["todoTaskListId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", input)
	}

	if id.TodoTaskId, ok = input.Parsed["todoTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "todoTaskId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeTodoListIdTaskIdExtensionID checks that 'input' can be parsed as a Me Todo List Id Task Id Extension ID
func ValidateMeTodoListIdTaskIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTodoListIdTaskIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Todo List Id Task Id Extension ID
func (id MeTodoListIdTaskIdExtensionId) ID() string {
	fmtString := "/me/todo/lists/%s/tasks/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.TodoTaskListId, id.TodoTaskId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Todo List Id Task Id Extension ID
func (id MeTodoListIdTaskIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("todo", "todo", "todo"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Me Todo List Id Task Id Extension ID
func (id MeTodoListIdTaskIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Todo List Id Task Id Extension (%s)", strings.Join(components, "\n"))
}
