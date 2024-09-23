package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTodoListIdTaskIdChecklistItemId{}

// UserIdTodoListIdTaskIdChecklistItemId is a struct representing the Resource ID for a User Id Todo List Id Task Id Checklist Item
type UserIdTodoListIdTaskIdChecklistItemId struct {
	UserId          string
	TodoTaskListId  string
	TodoTaskId      string
	ChecklistItemId string
}

// NewUserIdTodoListIdTaskIdChecklistItemID returns a new UserIdTodoListIdTaskIdChecklistItemId struct
func NewUserIdTodoListIdTaskIdChecklistItemID(userId string, todoTaskListId string, todoTaskId string, checklistItemId string) UserIdTodoListIdTaskIdChecklistItemId {
	return UserIdTodoListIdTaskIdChecklistItemId{
		UserId:          userId,
		TodoTaskListId:  todoTaskListId,
		TodoTaskId:      todoTaskId,
		ChecklistItemId: checklistItemId,
	}
}

// ParseUserIdTodoListIdTaskIdChecklistItemID parses 'input' into a UserIdTodoListIdTaskIdChecklistItemId
func ParseUserIdTodoListIdTaskIdChecklistItemID(input string) (*UserIdTodoListIdTaskIdChecklistItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListIdTaskIdChecklistItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListIdTaskIdChecklistItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdTodoListIdTaskIdChecklistItemIDInsensitively parses 'input' case-insensitively into a UserIdTodoListIdTaskIdChecklistItemId
// note: this method should only be used for API response data and not user input
func ParseUserIdTodoListIdTaskIdChecklistItemIDInsensitively(input string) (*UserIdTodoListIdTaskIdChecklistItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListIdTaskIdChecklistItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListIdTaskIdChecklistItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdTodoListIdTaskIdChecklistItemId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ChecklistItemId, ok = input.Parsed["checklistItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "checklistItemId", input)
	}

	return nil
}

// ValidateUserIdTodoListIdTaskIdChecklistItemID checks that 'input' can be parsed as a User Id Todo List Id Task Id Checklist Item ID
func ValidateUserIdTodoListIdTaskIdChecklistItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdTodoListIdTaskIdChecklistItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Todo List Id Task Id Checklist Item ID
func (id UserIdTodoListIdTaskIdChecklistItemId) ID() string {
	fmtString := "/users/%s/todo/lists/%s/tasks/%s/checklistItems/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId, id.TodoTaskId, id.ChecklistItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Todo List Id Task Id Checklist Item ID
func (id UserIdTodoListIdTaskIdChecklistItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("todo", "todo", "todo"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskId"),
		resourceids.StaticSegment("checklistItems", "checklistItems", "checklistItems"),
		resourceids.UserSpecifiedSegment("checklistItemId", "checklistItemId"),
	}
}

// String returns a human-readable description of this User Id Todo List Id Task Id Checklist Item ID
func (id UserIdTodoListIdTaskIdChecklistItemId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Checklist Item: %q", id.ChecklistItemId),
	}
	return fmt.Sprintf("User Id Todo List Id Task Id Checklist Item (%s)", strings.Join(components, "\n"))
}
