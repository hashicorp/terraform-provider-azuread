package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTodoListIdTaskIdLinkedResourceId{}

// UserIdTodoListIdTaskIdLinkedResourceId is a struct representing the Resource ID for a User Id Todo List Id Task Id Linked Resource
type UserIdTodoListIdTaskIdLinkedResourceId struct {
	UserId           string
	TodoTaskListId   string
	TodoTaskId       string
	LinkedResourceId string
}

// NewUserIdTodoListIdTaskIdLinkedResourceID returns a new UserIdTodoListIdTaskIdLinkedResourceId struct
func NewUserIdTodoListIdTaskIdLinkedResourceID(userId string, todoTaskListId string, todoTaskId string, linkedResourceId string) UserIdTodoListIdTaskIdLinkedResourceId {
	return UserIdTodoListIdTaskIdLinkedResourceId{
		UserId:           userId,
		TodoTaskListId:   todoTaskListId,
		TodoTaskId:       todoTaskId,
		LinkedResourceId: linkedResourceId,
	}
}

// ParseUserIdTodoListIdTaskIdLinkedResourceID parses 'input' into a UserIdTodoListIdTaskIdLinkedResourceId
func ParseUserIdTodoListIdTaskIdLinkedResourceID(input string) (*UserIdTodoListIdTaskIdLinkedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListIdTaskIdLinkedResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListIdTaskIdLinkedResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdTodoListIdTaskIdLinkedResourceIDInsensitively parses 'input' case-insensitively into a UserIdTodoListIdTaskIdLinkedResourceId
// note: this method should only be used for API response data and not user input
func ParseUserIdTodoListIdTaskIdLinkedResourceIDInsensitively(input string) (*UserIdTodoListIdTaskIdLinkedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListIdTaskIdLinkedResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListIdTaskIdLinkedResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdTodoListIdTaskIdLinkedResourceId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.LinkedResourceId, ok = input.Parsed["linkedResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "linkedResourceId", input)
	}

	return nil
}

// ValidateUserIdTodoListIdTaskIdLinkedResourceID checks that 'input' can be parsed as a User Id Todo List Id Task Id Linked Resource ID
func ValidateUserIdTodoListIdTaskIdLinkedResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdTodoListIdTaskIdLinkedResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Todo List Id Task Id Linked Resource ID
func (id UserIdTodoListIdTaskIdLinkedResourceId) ID() string {
	fmtString := "/users/%s/todo/lists/%s/tasks/%s/linkedResources/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId, id.TodoTaskId, id.LinkedResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Todo List Id Task Id Linked Resource ID
func (id UserIdTodoListIdTaskIdLinkedResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("todo", "todo", "todo"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("todoTaskId", "todoTaskId"),
		resourceids.StaticSegment("linkedResources", "linkedResources", "linkedResources"),
		resourceids.UserSpecifiedSegment("linkedResourceId", "linkedResourceId"),
	}
}

// String returns a human-readable description of this User Id Todo List Id Task Id Linked Resource ID
func (id UserIdTodoListIdTaskIdLinkedResourceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Todo Task: %q", id.TodoTaskId),
		fmt.Sprintf("Linked Resource: %q", id.LinkedResourceId),
	}
	return fmt.Sprintf("User Id Todo List Id Task Id Linked Resource (%s)", strings.Join(components, "\n"))
}
