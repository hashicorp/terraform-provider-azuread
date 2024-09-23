package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTodoListIdTaskIdExtensionId{}

// UserIdTodoListIdTaskIdExtensionId is a struct representing the Resource ID for a User Id Todo List Id Task Id Extension
type UserIdTodoListIdTaskIdExtensionId struct {
	UserId         string
	TodoTaskListId string
	TodoTaskId     string
	ExtensionId    string
}

// NewUserIdTodoListIdTaskIdExtensionID returns a new UserIdTodoListIdTaskIdExtensionId struct
func NewUserIdTodoListIdTaskIdExtensionID(userId string, todoTaskListId string, todoTaskId string, extensionId string) UserIdTodoListIdTaskIdExtensionId {
	return UserIdTodoListIdTaskIdExtensionId{
		UserId:         userId,
		TodoTaskListId: todoTaskListId,
		TodoTaskId:     todoTaskId,
		ExtensionId:    extensionId,
	}
}

// ParseUserIdTodoListIdTaskIdExtensionID parses 'input' into a UserIdTodoListIdTaskIdExtensionId
func ParseUserIdTodoListIdTaskIdExtensionID(input string) (*UserIdTodoListIdTaskIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListIdTaskIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListIdTaskIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdTodoListIdTaskIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdTodoListIdTaskIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdTodoListIdTaskIdExtensionIDInsensitively(input string) (*UserIdTodoListIdTaskIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListIdTaskIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListIdTaskIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdTodoListIdTaskIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdTodoListIdTaskIdExtensionID checks that 'input' can be parsed as a User Id Todo List Id Task Id Extension ID
func ValidateUserIdTodoListIdTaskIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdTodoListIdTaskIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Todo List Id Task Id Extension ID
func (id UserIdTodoListIdTaskIdExtensionId) ID() string {
	fmtString := "/users/%s/todo/lists/%s/tasks/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId, id.TodoTaskId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Todo List Id Task Id Extension ID
func (id UserIdTodoListIdTaskIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("todo", "todo", "todo"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this User Id Todo List Id Task Id Extension ID
func (id UserIdTodoListIdTaskIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Todo List Id Task Id Extension (%s)", strings.Join(components, "\n"))
}
