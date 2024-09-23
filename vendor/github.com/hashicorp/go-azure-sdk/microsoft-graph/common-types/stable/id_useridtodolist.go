package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTodoListId{}

// UserIdTodoListId is a struct representing the Resource ID for a User Id Todo List
type UserIdTodoListId struct {
	UserId         string
	TodoTaskListId string
}

// NewUserIdTodoListID returns a new UserIdTodoListId struct
func NewUserIdTodoListID(userId string, todoTaskListId string) UserIdTodoListId {
	return UserIdTodoListId{
		UserId:         userId,
		TodoTaskListId: todoTaskListId,
	}
}

// ParseUserIdTodoListID parses 'input' into a UserIdTodoListId
func ParseUserIdTodoListID(input string) (*UserIdTodoListId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdTodoListIDInsensitively parses 'input' case-insensitively into a UserIdTodoListId
// note: this method should only be used for API response data and not user input
func ParseUserIdTodoListIDInsensitively(input string) (*UserIdTodoListId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdTodoListId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TodoTaskListId, ok = input.Parsed["todoTaskListId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", input)
	}

	return nil
}

// ValidateUserIdTodoListID checks that 'input' can be parsed as a User Id Todo List ID
func ValidateUserIdTodoListID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdTodoListID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Todo List ID
func (id UserIdTodoListId) ID() string {
	fmtString := "/users/%s/todo/lists/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Todo List ID
func (id UserIdTodoListId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("todo", "todo", "todo"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListId"),
	}
}

// String returns a human-readable description of this User Id Todo List ID
func (id UserIdTodoListId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
	}
	return fmt.Sprintf("User Id Todo List (%s)", strings.Join(components, "\n"))
}
