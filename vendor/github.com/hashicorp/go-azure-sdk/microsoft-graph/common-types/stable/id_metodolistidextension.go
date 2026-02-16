package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeTodoListIdExtensionId{}

// MeTodoListIdExtensionId is a struct representing the Resource ID for a Me Todo List Id Extension
type MeTodoListIdExtensionId struct {
	TodoTaskListId string
	ExtensionId    string
}

// NewMeTodoListIdExtensionID returns a new MeTodoListIdExtensionId struct
func NewMeTodoListIdExtensionID(todoTaskListId string, extensionId string) MeTodoListIdExtensionId {
	return MeTodoListIdExtensionId{
		TodoTaskListId: todoTaskListId,
		ExtensionId:    extensionId,
	}
}

// ParseMeTodoListIdExtensionID parses 'input' into a MeTodoListIdExtensionId
func ParseMeTodoListIdExtensionID(input string) (*MeTodoListIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTodoListIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTodoListIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeTodoListIdExtensionIDInsensitively parses 'input' case-insensitively into a MeTodoListIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeTodoListIdExtensionIDInsensitively(input string) (*MeTodoListIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTodoListIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTodoListIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeTodoListIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TodoTaskListId, ok = input.Parsed["todoTaskListId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "todoTaskListId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeTodoListIdExtensionID checks that 'input' can be parsed as a Me Todo List Id Extension ID
func ValidateMeTodoListIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTodoListIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Todo List Id Extension ID
func (id MeTodoListIdExtensionId) ID() string {
	fmtString := "/me/todo/lists/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.TodoTaskListId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Todo List Id Extension ID
func (id MeTodoListIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("todo", "todo", "todo"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("todoTaskListId", "todoTaskListId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Me Todo List Id Extension ID
func (id MeTodoListIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Todo Task List: %q", id.TodoTaskListId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Todo List Id Extension (%s)", strings.Join(components, "\n"))
}
