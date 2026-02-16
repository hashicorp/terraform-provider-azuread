package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTodoListIdTaskId{}

// UserIdTodoListIdTaskId is a struct representing the Resource ID for a User Id Todo List Id Task
type UserIdTodoListIdTaskId struct {
	UserId         string
	TodoTaskListId string
	TodoTaskId     string
}

// NewUserIdTodoListIdTaskID returns a new UserIdTodoListIdTaskId struct
func NewUserIdTodoListIdTaskID(userId string, todoTaskListId string, todoTaskId string) UserIdTodoListIdTaskId {
	return UserIdTodoListIdTaskId{
		UserId:         userId,
		TodoTaskListId: todoTaskListId,
		TodoTaskId:     todoTaskId,
	}
}

// ParseUserIdTodoListIdTaskID parses 'input' into a UserIdTodoListIdTaskId
func ParseUserIdTodoListIdTaskID(input string) (*UserIdTodoListIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdTodoListIdTaskIDInsensitively parses 'input' case-insensitively into a UserIdTodoListIdTaskId
// note: this method should only be used for API response data and not user input
func ParseUserIdTodoListIdTaskIDInsensitively(input string) (*UserIdTodoListIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdTodoListIdTaskId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateUserIdTodoListIdTaskID checks that 'input' can be parsed as a User Id Todo List Id Task ID
func ValidateUserIdTodoListIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdTodoListIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Todo List Id Task ID
func (id UserIdTodoListIdTaskId) ID() string {
	fmtString := "/users/%s/todo/lists/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId, id.TodoTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Todo List Id Task ID
func (id UserIdTodoListIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("todo", "todo", "todo"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskId"),
	}
}

// String returns a human-readable description of this User Id Todo List Id Task ID
func (id UserIdTodoListIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
	}
	return fmt.Sprintf("User Id Todo List Id Task (%s)", strings.Join(components, "\n"))
}
