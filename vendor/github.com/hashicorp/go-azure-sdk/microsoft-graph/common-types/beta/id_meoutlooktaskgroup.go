package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOutlookTaskGroupId{}

// MeOutlookTaskGroupId is a struct representing the Resource ID for a Me Outlook Task Group
type MeOutlookTaskGroupId struct {
	OutlookTaskGroupId string
}

// NewMeOutlookTaskGroupID returns a new MeOutlookTaskGroupId struct
func NewMeOutlookTaskGroupID(outlookTaskGroupId string) MeOutlookTaskGroupId {
	return MeOutlookTaskGroupId{
		OutlookTaskGroupId: outlookTaskGroupId,
	}
}

// ParseMeOutlookTaskGroupID parses 'input' into a MeOutlookTaskGroupId
func ParseMeOutlookTaskGroupID(input string) (*MeOutlookTaskGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOutlookTaskGroupIDInsensitively parses 'input' case-insensitively into a MeOutlookTaskGroupId
// note: this method should only be used for API response data and not user input
func ParseMeOutlookTaskGroupIDInsensitively(input string) (*MeOutlookTaskGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOutlookTaskGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OutlookTaskGroupId, ok = input.Parsed["outlookTaskGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskGroupId", input)
	}

	return nil
}

// ValidateMeOutlookTaskGroupID checks that 'input' can be parsed as a Me Outlook Task Group ID
func ValidateMeOutlookTaskGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOutlookTaskGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Outlook Task Group ID
func (id MeOutlookTaskGroupId) ID() string {
	fmtString := "/me/outlook/taskGroups/%s"
	return fmt.Sprintf(fmtString, id.OutlookTaskGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Outlook Task Group ID
func (id MeOutlookTaskGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("taskGroups", "taskGroups", "taskGroups"),
		resourceids.UserSpecifiedSegment("outlookTaskGroupId", "outlookTaskGroupId"),
	}
}

// String returns a human-readable description of this Me Outlook Task Group ID
func (id MeOutlookTaskGroupId) String() string {
	components := []string{
		fmt.Sprintf("Outlook Task Group: %q", id.OutlookTaskGroupId),
	}
	return fmt.Sprintf("Me Outlook Task Group (%s)", strings.Join(components, "\n"))
}
