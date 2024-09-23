package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileProjectId{}

// MeProfileProjectId is a struct representing the Resource ID for a Me Profile Project
type MeProfileProjectId struct {
	ProjectParticipationId string
}

// NewMeProfileProjectID returns a new MeProfileProjectId struct
func NewMeProfileProjectID(projectParticipationId string) MeProfileProjectId {
	return MeProfileProjectId{
		ProjectParticipationId: projectParticipationId,
	}
}

// ParseMeProfileProjectID parses 'input' into a MeProfileProjectId
func ParseMeProfileProjectID(input string) (*MeProfileProjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileProjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileProjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeProfileProjectIDInsensitively parses 'input' case-insensitively into a MeProfileProjectId
// note: this method should only be used for API response data and not user input
func ParseMeProfileProjectIDInsensitively(input string) (*MeProfileProjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeProfileProjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeProfileProjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeProfileProjectId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ProjectParticipationId, ok = input.Parsed["projectParticipationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "projectParticipationId", input)
	}

	return nil
}

// ValidateMeProfileProjectID checks that 'input' can be parsed as a Me Profile Project ID
func ValidateMeProfileProjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileProjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Project ID
func (id MeProfileProjectId) ID() string {
	fmtString := "/me/profile/projects/%s"
	return fmt.Sprintf(fmtString, id.ProjectParticipationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Project ID
func (id MeProfileProjectId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("projects", "projects", "projects"),
		resourceids.UserSpecifiedSegment("projectParticipationId", "projectParticipationId"),
	}
}

// String returns a human-readable description of this Me Profile Project ID
func (id MeProfileProjectId) String() string {
	components := []string{
		fmt.Sprintf("Project Participation: %q", id.ProjectParticipationId),
	}
	return fmt.Sprintf("Me Profile Project (%s)", strings.Join(components, "\n"))
}
