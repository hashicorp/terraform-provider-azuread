package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOwnedObjectId{}

// MeOwnedObjectId is a struct representing the Resource ID for a Me Owned Object
type MeOwnedObjectId struct {
	DirectoryObjectId string
}

// NewMeOwnedObjectID returns a new MeOwnedObjectId struct
func NewMeOwnedObjectID(directoryObjectId string) MeOwnedObjectId {
	return MeOwnedObjectId{
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeOwnedObjectID parses 'input' into a MeOwnedObjectId
func ParseMeOwnedObjectID(input string) (*MeOwnedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOwnedObjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOwnedObjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOwnedObjectIDInsensitively parses 'input' case-insensitively into a MeOwnedObjectId
// note: this method should only be used for API response data and not user input
func ParseMeOwnedObjectIDInsensitively(input string) (*MeOwnedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOwnedObjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOwnedObjectId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOwnedObjectId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateMeOwnedObjectID checks that 'input' can be parsed as a Me Owned Object ID
func ValidateMeOwnedObjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOwnedObjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Owned Object ID
func (id MeOwnedObjectId) ID() string {
	fmtString := "/me/ownedObjects/%s"
	return fmt.Sprintf(fmtString, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Owned Object ID
func (id MeOwnedObjectId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("ownedObjects", "ownedObjects", "ownedObjects"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Me Owned Object ID
func (id MeOwnedObjectId) String() string {
	components := []string{
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Owned Object (%s)", strings.Join(components, "\n"))
}
