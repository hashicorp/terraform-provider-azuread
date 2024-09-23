package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeTransitiveMemberOfId{}

// MeTransitiveMemberOfId is a struct representing the Resource ID for a Me Transitive Member Of
type MeTransitiveMemberOfId struct {
	DirectoryObjectId string
}

// NewMeTransitiveMemberOfID returns a new MeTransitiveMemberOfId struct
func NewMeTransitiveMemberOfID(directoryObjectId string) MeTransitiveMemberOfId {
	return MeTransitiveMemberOfId{
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeTransitiveMemberOfID parses 'input' into a MeTransitiveMemberOfId
func ParseMeTransitiveMemberOfID(input string) (*MeTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTransitiveMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeTransitiveMemberOfIDInsensitively parses 'input' case-insensitively into a MeTransitiveMemberOfId
// note: this method should only be used for API response data and not user input
func ParseMeTransitiveMemberOfIDInsensitively(input string) (*MeTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTransitiveMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeTransitiveMemberOfId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateMeTransitiveMemberOfID checks that 'input' can be parsed as a Me Transitive Member Of ID
func ValidateMeTransitiveMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTransitiveMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Transitive Member Of ID
func (id MeTransitiveMemberOfId) ID() string {
	fmtString := "/me/transitiveMemberOf/%s"
	return fmt.Sprintf(fmtString, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Transitive Member Of ID
func (id MeTransitiveMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("transitiveMemberOf", "transitiveMemberOf", "transitiveMemberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Me Transitive Member Of ID
func (id MeTransitiveMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Transitive Member Of (%s)", strings.Join(components, "\n"))
}
