package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOutlookTaskId{}

// MeOutlookTaskId is a struct representing the Resource ID for a Me Outlook Task
type MeOutlookTaskId struct {
	OutlookTaskId string
}

// NewMeOutlookTaskID returns a new MeOutlookTaskId struct
func NewMeOutlookTaskID(outlookTaskId string) MeOutlookTaskId {
	return MeOutlookTaskId{
		OutlookTaskId: outlookTaskId,
	}
}

// ParseMeOutlookTaskID parses 'input' into a MeOutlookTaskId
func ParseMeOutlookTaskID(input string) (*MeOutlookTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOutlookTaskIDInsensitively parses 'input' case-insensitively into a MeOutlookTaskId
// note: this method should only be used for API response data and not user input
func ParseMeOutlookTaskIDInsensitively(input string) (*MeOutlookTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOutlookTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OutlookTaskId, ok = input.Parsed["outlookTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", input)
	}

	return nil
}

// ValidateMeOutlookTaskID checks that 'input' can be parsed as a Me Outlook Task ID
func ValidateMeOutlookTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOutlookTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Outlook Task ID
func (id MeOutlookTaskId) ID() string {
	fmtString := "/me/outlook/tasks/%s"
	return fmt.Sprintf(fmtString, id.OutlookTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Outlook Task ID
func (id MeOutlookTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskId"),
	}
}

// String returns a human-readable description of this Me Outlook Task ID
func (id MeOutlookTaskId) String() string {
	components := []string{
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
	}
	return fmt.Sprintf("Me Outlook Task (%s)", strings.Join(components, "\n"))
}
