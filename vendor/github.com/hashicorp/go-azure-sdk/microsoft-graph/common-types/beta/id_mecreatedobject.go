package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCreatedObjectId{}

// MeCreatedObjectId is a struct representing the Resource ID for a Me Created Object
type MeCreatedObjectId struct {
	DirectoryObjectId string
}

// NewMeCreatedObjectID returns a new MeCreatedObjectId struct
func NewMeCreatedObjectID(directoryObjectId string) MeCreatedObjectId {
	return MeCreatedObjectId{
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeCreatedObjectID parses 'input' into a MeCreatedObjectId
func ParseMeCreatedObjectID(input string) (*MeCreatedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCreatedObjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCreatedObjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCreatedObjectIDInsensitively parses 'input' case-insensitively into a MeCreatedObjectId
// note: this method should only be used for API response data and not user input
func ParseMeCreatedObjectIDInsensitively(input string) (*MeCreatedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCreatedObjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCreatedObjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCreatedObjectId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateMeCreatedObjectID checks that 'input' can be parsed as a Me Created Object ID
func ValidateMeCreatedObjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCreatedObjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Created Object ID
func (id MeCreatedObjectId) ID() string {
	fmtString := "/me/createdObjects/%s"
	return fmt.Sprintf(fmtString, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Created Object ID
func (id MeCreatedObjectId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("createdObjects", "createdObjects", "createdObjects"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Me Created Object ID
func (id MeCreatedObjectId) String() string {
	components := []string{
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Created Object (%s)", strings.Join(components, "\n"))
}
