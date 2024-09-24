package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAppRoleAssignmentId{}

// MeAppRoleAssignmentId is a struct representing the Resource ID for a Me App Role Assignment
type MeAppRoleAssignmentId struct {
	AppRoleAssignmentId string
}

// NewMeAppRoleAssignmentID returns a new MeAppRoleAssignmentId struct
func NewMeAppRoleAssignmentID(appRoleAssignmentId string) MeAppRoleAssignmentId {
	return MeAppRoleAssignmentId{
		AppRoleAssignmentId: appRoleAssignmentId,
	}
}

// ParseMeAppRoleAssignmentID parses 'input' into a MeAppRoleAssignmentId
func ParseMeAppRoleAssignmentID(input string) (*MeAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAppRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAppRoleAssignmentIDInsensitively parses 'input' case-insensitively into a MeAppRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseMeAppRoleAssignmentIDInsensitively(input string) (*MeAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAppRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAppRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AppRoleAssignmentId, ok = input.Parsed["appRoleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appRoleAssignmentId", input)
	}

	return nil
}

// ValidateMeAppRoleAssignmentID checks that 'input' can be parsed as a Me App Role Assignment ID
func ValidateMeAppRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAppRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me App Role Assignment ID
func (id MeAppRoleAssignmentId) ID() string {
	fmtString := "/me/appRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.AppRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me App Role Assignment ID
func (id MeAppRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("appRoleAssignments", "appRoleAssignments", "appRoleAssignments"),
		resourceids.UserSpecifiedSegment("appRoleAssignmentId", "appRoleAssignmentId"),
	}
}

// String returns a human-readable description of this Me App Role Assignment ID
func (id MeAppRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("App Role Assignment: %q", id.AppRoleAssignmentId),
	}
	return fmt.Sprintf("Me App Role Assignment (%s)", strings.Join(components, "\n"))
}
