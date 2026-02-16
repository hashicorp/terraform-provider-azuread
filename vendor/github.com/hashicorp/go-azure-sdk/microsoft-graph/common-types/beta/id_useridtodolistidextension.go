package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTodoListIdExtensionId{}

// UserIdTodoListIdExtensionId is a struct representing the Resource ID for a User Id Todo List Id Extension
type UserIdTodoListIdExtensionId struct {
	UserId         string
	TodoTaskListId string
	ExtensionId    string
}

// NewUserIdTodoListIdExtensionID returns a new UserIdTodoListIdExtensionId struct
func NewUserIdTodoListIdExtensionID(userId string, todoTaskListId string, extensionId string) UserIdTodoListIdExtensionId {
	return UserIdTodoListIdExtensionId{
		UserId:         userId,
		TodoTaskListId: todoTaskListId,
		ExtensionId:    extensionId,
	}
}

// ParseUserIdTodoListIdExtensionID parses 'input' into a UserIdTodoListIdExtensionId
func ParseUserIdTodoListIdExtensionID(input string) (*UserIdTodoListIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdTodoListIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdTodoListIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdTodoListIdExtensionIDInsensitively(input string) (*UserIdTodoListIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTodoListIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTodoListIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdTodoListIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TodoTaskListId, ok = input.Parsed["todoTaskListId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdTodoListIdExtensionID checks that 'input' can be parsed as a User Id Todo List Id Extension ID
func ValidateUserIdTodoListIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdTodoListIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Todo List Id Extension ID
func (id UserIdTodoListIdExtensionId) ID() string {
	fmtString := "/users/%s/todo/lists/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TodoTaskListId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Todo List Id Extension ID
func (id UserIdTodoListIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("todo", "todo", "todo"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this User Id Todo List Id Extension ID
func (id UserIdTodoListIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Todo List Id Extension (%s)", strings.Join(components, "\n"))
}
