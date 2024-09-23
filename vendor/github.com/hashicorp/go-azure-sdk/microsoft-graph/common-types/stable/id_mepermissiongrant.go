package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePermissionGrantId{}

// MePermissionGrantId is a struct representing the Resource ID for a Me Permission Grant
type MePermissionGrantId struct {
	ResourceSpecificPermissionGrantId string
}

// NewMePermissionGrantID returns a new MePermissionGrantId struct
func NewMePermissionGrantID(resourceSpecificPermissionGrantId string) MePermissionGrantId {
	return MePermissionGrantId{
		ResourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// ParseMePermissionGrantID parses 'input' into a MePermissionGrantId
func ParseMePermissionGrantID(input string) (*MePermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePermissionGrantIDInsensitively parses 'input' case-insensitively into a MePermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseMePermissionGrantIDInsensitively(input string) (*MePermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ResourceSpecificPermissionGrantId, ok = input.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", input)
	}

	return nil
}

// ValidateMePermissionGrantID checks that 'input' can be parsed as a Me Permission Grant ID
func ValidateMePermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Permission Grant ID
func (id MePermissionGrantId) ID() string {
	fmtString := "/me/permissionGrants/%s"
	return fmt.Sprintf(fmtString, id.ResourceSpecificPermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Permission Grant ID
func (id MePermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("permissionGrants", "permissionGrants", "permissionGrants"),
		resourceids.UserSpecifiedSegment("resourceSpecificPermissionGrantId", "resourceSpecificPermissionGrantId"),
	}
}

// String returns a human-readable description of this Me Permission Grant ID
func (id MePermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("Resource Specific Permission Grant: %q", id.ResourceSpecificPermissionGrantId),
	}
	return fmt.Sprintf("Me Permission Grant (%s)", strings.Join(components, "\n"))
}
